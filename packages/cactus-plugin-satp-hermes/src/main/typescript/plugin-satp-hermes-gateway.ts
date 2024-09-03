import {
  Secp256k1Keys,
  Logger,
  Checks,
  LoggerProvider,
  ILoggerOptions,
  JsObjectSigner,
  IJsObjectSignerOptions,
  LogLevelDesc,
} from "@hyperledger/cactus-common";
import { v4 as uuidv4 } from "uuid";

import {
  IsDefined,
  IsNotEmptyObject,
  IsObject,
  IsString,
  Contains,
  ValidatorOptions,
} from "class-validator";

import path from "path";

import {
  SATPGatewayConfig,
  GatewayIdentity,
  ShutdownHook,
  SupportedChain,
  Address,
  DraftVersions,
} from "./core/types";
import {
  GatewayOrchestrator,
  IGatewayOrchestratorOptions,
} from "./gol/gateway-orchestrator";
export { SATPGatewayConfig };
import express, { Express } from "express";
import http from "http";
import {
  DEFAULT_PORT_GATEWAY_CLIENT,
  DEFAULT_PORT_GATEWAY_SERVER,
  SATP_VERSION,
} from "./core/constants";
import { bufArray2HexStr } from "./gateway-utils";
import {
  ILocalLogRepository,
  IRemoteLogRepository,
} from "./repository/interfaces/repository";
import { BLODispatcher, BLODispatcherOptions } from "./blo/dispatcher";
import fs from "fs";
import swaggerUi, { JsonObject } from "swagger-ui-express";
import {
  IPluginWebService,
  ICactusPlugin,
  IWebServiceEndpoint,
} from "@hyperledger/cactus-core-api";
import {
  ISATPBridgesOptions,
  SATPBridgesManager,
} from "./gol/satp-bridges-manager";

import dotenv from "dotenv";
import { IPrivacyPolicyValue } from "@hyperledger/cactus-plugin-bungee-hermes/dist/lib/main/typescript/view-creation/privacy-policies";
import {
  MergePolicyOpts,
  PrivacyPolicyOpts,
} from "@hyperledger/cactus-plugin-bungee-hermes/dist/lib/main/typescript/generated/openapi/typescript-axios";
import { IMergePolicyValue } from "@hyperledger/cactus-plugin-bungee-hermes/dist/lib/main/typescript/view-merging/merge-policies";
import { ISignerKeyPairs } from "@hyperledger/cactus-common/src/main/typescript/signer-key-pairs";
dotenv.config({ path: path.resolve(__dirname, "../../../.env.example") });

export class SATPGateway implements IPluginWebService, ICactusPlugin {
  // todo more checks; example port from config is between 3000 and 9000
  @IsDefined()
  @IsNotEmptyObject()
  @IsObject()
  private readonly logger: Logger;

  @IsDefined()
  @IsNotEmptyObject()
  @IsObject()
  private readonly config: SATPGatewayConfig;

  @IsString()
  @Contains("Gateway")
  public readonly className = "SATPGateway";

  @IsString()
  public readonly instanceId: string;
  private supportedDltIDs: SupportedChain[];
  private gatewayOrchestrator: GatewayOrchestrator;
  private bridgesManager: SATPBridgesManager;

  private BLOApplication?: Express;
  private BLOServer?: http.Server;
  private BLODispatcher?: BLODispatcher;
  private GOLApplication?: Express;
  private GOLServer?: http.Server;
  private readonly OAS: JsonObject;
  public OAPIServerEnabled: boolean = false;

  private signer: JsObjectSigner;
  private _pubKey: string;

  public localRepository?: ILocalLogRepository;
  public remoteRepository?: IRemoteLogRepository;
  private readonly shutdownHooks: ShutdownHook[];

  constructor(public readonly options: SATPGatewayConfig) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    this.config = SATPGateway.ProcessGatewayCoordinatorConfig(options);
    this.shutdownHooks = [];
    const level = options.logLevel || "INFO";
    const logOptions: ILoggerOptions = {
      level: level,
      label: this.className,
    };
    this.logger = LoggerProvider.getOrCreate(logOptions);
    this.logger.info("Initializing Gateway Coordinator");

    if (this.config.keyPair == undefined) {
      throw new Error("Key pair is undefined");
    }

    this._pubKey = bufArray2HexStr(this.config.keyPair.publicKey);

    this.logger.info(`Gateway's public key: ${this._pubKey}`);

    const signerOptions: IJsObjectSignerOptions = {
      privateKey: bufArray2HexStr(this.config.keyPair.privateKey),
      logLevel: "debug",
    };
    this.signer = new JsObjectSigner(signerOptions);

    const gatewayOrchestratorOptions: IGatewayOrchestratorOptions = {
      logLevel: this.config.logLevel,
      localGateway: this.config.gid!,
      counterPartyGateways: this.config.counterPartyGateways,
      signer: this.signer!,
    };

    const bridgesManagerOptions: ISATPBridgesOptions = {
      logLevel: this.config.logLevel,
      supportedDLTs: this.config.gid!.supportedDLTs,
      networks: options.bridgesConfig ? options.bridgesConfig : [],
    };

    this.bridgesManager = new SATPBridgesManager(bridgesManagerOptions);

    if (!this.bridgesManager) {
      throw new Error("BridgesManager is not defined");
    }

    if (this.config.gid) {
      this.logger.info(
        "Initializing gateway connection manager with seed gateways",
      );
      this.gatewayOrchestrator = new GatewayOrchestrator(
        gatewayOrchestratorOptions,
      );
    } else {
      throw new Error("GatewayIdentity is not defined");
    }

    this.instanceId = uuidv4();
    const dispatcherOps: BLODispatcherOptions = {
      logger: this.logger,
      logLevel: this.config.logLevel,
      instanceId: this.config.gid!.id,
      orchestrator: this.gatewayOrchestrator,
      signer: this.signer,
      bridgesManager: this.bridgesManager,
      pubKey: this.pubKey,
    };

    this.supportedDltIDs = this.config.gid!.supportedDLTs;

    if (!this.config.gid || !dispatcherOps.instanceId) {
      throw new Error("Invalid configuration");
    }

    this.BLODispatcher = new BLODispatcher(dispatcherOps);
    this.OAPIServerEnabled = this.config.enableOpenAPI ?? true;

    const specPath = path.join(__dirname, "../json/openapi-blo-bundled.json");
    this.OAS = JSON.parse(fs.readFileSync(specPath, "utf8"));
    if (!this.OAS) {
      this.logger.warn("Error loading OAS");
    }
  }

  /* ICactus Plugin methods */

  public getInstanceId(): string {
    return this.instanceId;
  }

  public getPackageName(): string {
    return `@hyperledger/cactus-plugin-satp-hermes`;
  }

  //for testing
  public getBLODispatcher(): BLODispatcher | undefined {
    return this.BLODispatcher;
  }

  public async onPluginInit(): Promise<unknown> {
    const fnTag = `${this.className}#onPluginInit()`;
    this.logger.trace(`Entering ${fnTag}`);
    // resolve gateways on init
    throw new Error("Not implemented");
  }

  /* IPluginWebService methods */
  async registerWebServices(app: Express): Promise<IWebServiceEndpoint[]> {
    const webServices = await this.getOrCreateWebServices();
    webServices.forEach((ws) => {
      this.logger.debug(`Registering service ${ws.getPath()}`);
      ws.registerExpress(app);
    });
    return webServices;
  }

  public async getOrCreateWebServices(): Promise<IWebServiceEndpoint[]> {
    const fnTag = `${this.className}#getOrCreateWebServices()`;
    this.logger.trace(`Entering ${fnTag}`);
    if (!this.BLODispatcher) {
      throw new Error(`Cannot ${fnTag} because BLODispatcher is erroneous`);
    }
    return this.BLODispatcher?.getOrCreateWebServices();
  }

  /* Getters */

  public get Signer(): JsObjectSigner {
    return this.signer;
  }

  public getSupportedDltIDs(): string[] {
    return this.supportedDltIDs;
  }

  public get gatewaySigner(): JsObjectSigner {
    return this.signer;
  }

  public get pubKey(): string {
    return this._pubKey;
  }

  public getOpenApiSpec(): unknown {
    return this.OAS;
  }

  // TODO: keep getter; add an admin endpoint to get identity of connected gateway to BLO
  public get Identity(): GatewayIdentity {
    const fnTag = `${this.className}#getIdentity()`;
    this.logger.trace(`Entering ${fnTag}`);
    if (!this.config.gid) {
      throw new Error("GatewayIdentity is not defined");
    }
    return this.config.gid!;
  }

  /* Gateway configuration helpers */

  private static processGatewayId(): string {
    return process.env.SATP_GATEWAY_ID || uuidv4();
  }

  private static processGatewayName(): string {
    return process.env.SATP_GATEWAY_NAME || uuidv4();
  }

  private static processGatewayVersion(): DraftVersions[] {
    return [
      {
        Core: process.env.SATP_GATEWAY_VERSION_CORE || SATP_VERSION,
        Architecture:
          process.env.SATP_GATEWAY_VERSION_ARCHITECTURE || SATP_VERSION,
        Crash: process.env.SATP_GATEWAY_VERSION_CRASH || SATP_VERSION,
      },
    ];
  }

  private static processGatewaySupportedDLTs(): SupportedChain[] {
    if (process.env.SATP_SUPPORTED_DLTS) {
      const dlts = process.env.SATP_SUPPORTED_DLTS.split(",").filter((dlt) =>
        Object.values(SupportedChain).includes(dlt as SupportedChain),
      ) as SupportedChain[];
      if (dlts.length > 0) {
        return dlts;
      }
      console.warn(
        "SATP_SUPPORTED_DLTS is empty or contains no valid DLTs. Using default values.",
      );
    }
    return [];
  }

  private static processGatewayProofID(): string {
    return process.env.SATP_PROOF_ID || uuidv4();
  }

  private static processGatewayServerPort(): number {
    const port = Number(process.env.SATP_GATEWAY_SERVER_PORT);
    if (process.env.SATP_GATEWAY_SERVER_PORT && !isNaN(Number(port))) {
      return parseInt(process.env.SATP_GATEWAY_SERVER_PORT);
    }
    return DEFAULT_PORT_GATEWAY_SERVER;
  }

  private static processGatewayClientPort(): number {
    const port = Number(process.env.SATP_GATEWAY_CLIENT_PORT);
    if (process.env.SATP_GATEWAY_CLIENT_PORT && !isNaN(Number(port))) {
      return parseInt(process.env.SATP_GATEWAY_CLIENT_PORT);
    }
    return DEFAULT_PORT_GATEWAY_CLIENT;
  }

  private static processGatewayAddress(): Address {
    return (
      (process.env.SATP_GATEWAY_ADDRESS as Address) ||
      `http://localhost:${DEFAULT_PORT_GATEWAY_CLIENT}`
    );
  }

  private static processGatewayIdentity(
    pluginOptions: SATPGatewayConfig,
  ): GatewayIdentity {
    if (!pluginOptions.gid) {
      return {
        id: this.processGatewayId(),
        pubKey: bufArray2HexStr(pluginOptions.keyPair!.publicKey),
        name: this.processGatewayName(),
        version: this.processGatewayVersion(),
        supportedDLTs: this.processGatewaySupportedDLTs(),
        proofID: this.processGatewayProofID(),
        gatewayServerPort: this.processGatewayServerPort(),
        gatewayClientPort: this.processGatewayClientPort(),
        address: this.processGatewayAddress(),
      };
    } else {
      const gid = pluginOptions.gid;
      return {
        id: gid.id, // || this.processGatewayId(),
        pubKey: bufArray2HexStr(pluginOptions.keyPair!.publicKey),
        name: gid.name || this.processGatewayName(),
        version: gid.version, // || this.processGatewayVersion(),
        supportedDLTs: gid.supportedDLTs,
        // gid.supportedDLTs && gid.supportedDLTs.length > 0
        // ? gid.supportedDLTs
        // : this.processGatewaySupportedDLTs(),
        proofID: gid.proofID || this.processGatewayProofID(),
        gatewayServerPort:
          gid.gatewayServerPort && gid.gatewayServerPort !== 0
            ? gid.gatewayServerPort
            : this.processGatewayServerPort(),
        gatewayClientPort:
          gid.gatewayClientPort && gid.gatewayClientPort !== 0
            ? gid.gatewayClientPort
            : this.processGatewayClientPort(),
        address: gid.address || this.processGatewayAddress(),
      };
    }
  }

  private static processCounterPartyGateways(): GatewayIdentity[] {
    if (process.env.SATP_COUNTER_PARTY_GATEWAYS) {
      try {
        const parsedGateways = JSON.parse(
          process.env.SATP_COUNTER_PARTY_GATEWAYS,
        ) as GatewayIdentity[];

        const validGateway = (gateway: unknown): gateway is GatewayIdentity => {
          if (!gateway) {
            return false;
          }
          const gw = gateway as Record<string, unknown>;
          if (
            !("id" in gw) ||
            !("version" in gw) ||
            !("supportedDLTs" in gw) ||
            typeof gw.id !== "string" ||
            typeof gw.version !== "string" ||
            typeof gw.supportedDLTs !== "string"
          ) {
            return false;
          }
          const [Core, Architecture, Crash] = gw.version.split(",");
          if (!Core || !Architecture || !Crash) {
            return false;
          }
          const dlts = gw.supportedDLTs.split(",") as SupportedChain[];
          if (
            !dlts.every((dlt) => Object.values(SupportedChain).includes(dlt)) ||
            dlts.length === 0
          ) {
            return false;
          }
          // validate optional fields if provided
          if (
            ("name" in gw && typeof gw.name !== "string") ||
            ("proofID" in gw && typeof gw.proofID !== "string") ||
            ("gatewayServerPort" in gw &&
              (typeof gw.gatewayServerPort !== "string" ||
                isNaN(Number(gw.gatewayServerPort)))) ||
            ("gatewayClientPort" in gw &&
              (typeof gw.gatewayClientPort !== "string" ||
                isNaN(Number(gw.gatewayClientPort)))) ||
            ("address" in gw && typeof gw.address !== "string")
          ) {
            return false;
          }
          return true;
        };

        if (
          !Array.isArray(parsedGateways) ||
          !parsedGateways.every(validGateway)
        ) {
          throw new Error(
            "SATP_COUNTER_PARTY_GATEWAYS must be an array of valid gateway identities",
          );
        } else {
          return parsedGateways;
        }
      } catch (error) {
        console.warn(
          `Failed to parse SATP_COUNTER_PARTY_GATEWAYS: ${error.message}. Using default.`,
        );
      }
    }
    return [];
  }

  private static processLogLevel(): LogLevelDesc {
    return (
      (process.env.SATP_LOG_LEVEL?.toUpperCase() as LogLevelDesc) || "DEBUG"
    );
  }

  private static processKeyPair(): ISignerKeyPairs {
    if (process.env.SATP_PUBLIC_KEY && process.env.SATP_PRIVATE_KEY) {
      return {
        publicKey: Buffer.from(process.env.SATP_PUBLIC_KEY, "hex"),
        privateKey: Buffer.from(process.env.SATP_PRIVATE_KEY, "hex"),
      };
    }
    return Secp256k1Keys.generateKeyPairsBuffer();
  }

  private static processEnvironment(): "development" | "production" {
    return (
      (process.env.SATP_NODE_ENV as "development" | "production") ||
      "development"
    );
  }

  private static processEnableOpenAPI(): boolean {
    if (process.env.SATP_ENABLE_OPEN_API === "false") {
      return false;
    }
    return true;
  }

  private static processValidationOptions(): ValidatorOptions {
    if (process.env.SATP_VALIDATION_OPTIONS) {
      try {
        const envValidationOptions = JSON.parse(
          process.env.SATP_VALIDATION_OPTIONS,
        ) as ValidatorOptions;

        if (
          typeof envValidationOptions.skipMissingProperties !== "boolean" &&
          envValidationOptions.skipMissingProperties !== undefined
        ) {
          throw new Error(
            "skipMissingProperties must be a boolean if provided",
          );
        } else {
          return envValidationOptions;
        }
      } catch (error) {
        console.warn(
          `Failed to parse SATP_VALIDATION_OPTIONS: ${error}. Using default.`,
        );
      }
    }
    return {};
  }

  private static processPrivacyPolicies(): IPrivacyPolicyValue[] {
    if (process.env.SATP_PRIVACY_POLICIES) {
      try {
        const parsedPolicies = JSON.parse(
          process.env.SATP_PRIVACY_POLICIES,
        ) as IPrivacyPolicyValue[];

        const validPolicies = (
          eachPolicy: unknown,
        ): eachPolicy is IPrivacyPolicyValue => {
          if (!eachPolicy) {
            return false;
          }
          const policy = eachPolicy as Record<string, unknown>;
          return (
            "policy" in policy &&
            "policyHash" in policy &&
            typeof policy.policy === "string" &&
            typeof policy.policyHash === "string" &&
            (policy.policy === PrivacyPolicyOpts.PruneState ||
              policy.policy === PrivacyPolicyOpts.SingleTransaction)
          );
        };

        if (
          !Array.isArray(parsedPolicies) ||
          !parsedPolicies.every(validPolicies)
        ) {
          throw new Error(
            "SATP_PRIVACY_POLICIES must be an array of valid privacy policies",
          );
        } else {
          return parsedPolicies;
        }
      } catch (error) {
        console.warn(
          `Failed to parse SATP_PRIVACY_POLICIES: ${error.message}. Using default.`,
        );
      }
    }
    return [];
  }

  private static processMergePolicies(): IMergePolicyValue[] {
    if (process.env.SATP_MERGE_POLICIES) {
      try {
        const parsedPolicies = JSON.parse(
          process.env.SATP_MERGE_POLICIES,
        ) as IMergePolicyValue[];

        const validPolicies = (
          eachPolicy: unknown,
        ): eachPolicy is IMergePolicyValue => {
          if (!eachPolicy) {
            return false;
          }
          const policy = eachPolicy as Record<string, unknown>;
          return (
            "policy" in policy &&
            "policyHash" in policy &&
            typeof policy.policy === "string" &&
            typeof policy.policyHash === "string" &&
            (policy.policy === MergePolicyOpts.PruneState ||
              policy.policy === MergePolicyOpts.PruneStateFromView ||
              policy.policy === MergePolicyOpts.NONE)
          );
        };

        if (
          !Array.isArray(parsedPolicies) ||
          !parsedPolicies.every(validPolicies)
        ) {
          throw new Error(
            "SATP_MERGE_POLICIES must be an array of valid merge policies if provided",
          );
        } else {
          return parsedPolicies;
        }
      } catch (error) {
        console.warn(
          `Failed to parse SATP_MERGE_POLICIES: ${error.message}. Using default.`,
        );
      }
    }
    return [];
  }

  static ProcessGatewayCoordinatorConfig(
    pluginOptions: SATPGatewayConfig,
  ): SATPGatewayConfig {
    if (!pluginOptions.keyPair) {
      pluginOptions.keyPair = this.processKeyPair();
    }

    pluginOptions.gid = this.processGatewayIdentity(pluginOptions);

    if (!pluginOptions.counterPartyGateways) {
      pluginOptions.counterPartyGateways = this.processCounterPartyGateways();
    }

    if (!pluginOptions.logLevel) {
      pluginOptions.logLevel = this.processLogLevel();
    }

    if (!pluginOptions.environment) {
      pluginOptions.environment = this.processEnvironment();
    }

    if (!pluginOptions.enableOpenAPI) {
      pluginOptions.enableOpenAPI = this.processEnableOpenAPI();
    }

    if (!pluginOptions.validationOptions) {
      pluginOptions.validationOptions = this.processValidationOptions();
    }

    if (!pluginOptions.privacyPolicies) {
      pluginOptions.privacyPolicies = this.processPrivacyPolicies();
    }

    if (!pluginOptions.mergePolicies) {
      pluginOptions.mergePolicies = this.processMergePolicies();
    }

    return pluginOptions;
  }

  /**
   * Startup Methods
   * ----------------
   * This section includes methods responsible for starting up the server and its associated services independently of the existence of a Hyperledger Cacti Node.
   * It ensures that both the GatewayServer and BLOServer are initiated concurrently for efficient launch.
   */
  public async startup(): Promise<void> {
    const fnTag = `${this.className}#startup()`;
    this.logger.trace(`Entering ${fnTag}`);

    await Promise.all([this.startupBLOServer(), this.setupOpenAPIServer()]);

    await Promise.all([this.startupGOLServer()]);
  }

  protected async startupBLOServer(): Promise<void> {
    // starts BOL
    const fnTag = `${this.className}#startupBLOServer()`;
    this.logger.trace(`Entering ${fnTag}`);
    this.logger.info("Starting BOL server");
    const port =
      this.options.gid?.gatewayClientPort ?? DEFAULT_PORT_GATEWAY_CLIENT;

    return new Promise(async (resolve, reject) => {
      if (!this.BLOApplication || !this.BLOServer) {
        if (!this.BLODispatcher) {
          throw new Error("BLODispatcher is not defined");
        }
        this.BLOApplication = express();
        try {
          const webServices = await this.BLODispatcher.getOrCreateWebServices();
          for (const service of webServices) {
            this.logger.debug(`Registering web service: ${service.getPath()}`);
            await service.registerExpress(this.BLOApplication);
          }
        } catch (error) {
          throw new Error(`Failed to register web services: ${error}`);
        }

        this.BLOServer = http.createServer(this.BLOApplication);

        this.BLOServer.listen(port, () => {
          this.logger.info(`BOL server started and listening on port ${port}`);
          resolve();
        });

        this.BLOServer.on("error", (error) => {
          this.logger.error(`BOL server failed to start: ${error}`);
          reject(error);
        });
      } else {
        this.logger.warn("BOL Server already running.");
        resolve();
      }
    });
  }

  protected async startupGOLServer(): Promise<void> {
    // starts GOL
    const fnTag = `${this.className}#startupGOLServer()`;
    this.logger.trace(`Entering ${fnTag}`);
    this.logger.info("Starting GOL server");

    const port =
      this.options.gid?.gatewayServerPort ?? DEFAULT_PORT_GATEWAY_SERVER;

    return new Promise(async (resolve, reject) => {
      if (!this.GOLServer) {
        this.GOLApplication = express();

        this.gatewayOrchestrator.addGOLServer(this.GOLApplication!);
        this.gatewayOrchestrator.startServices();

        this.GOLServer = http.createServer(this.GOLApplication);

        this.GOLServer.listen(port, () => {
          this.logger.info(`GOL server started and listening on port ${port}`);
          resolve();
        });

        this.GOLServer.on("error", (error) => {
          this.logger.error(`GOL server failed to start: ${error}`);
          reject(error);
        });
      } else {
        this.logger.warn("GOL Server already running.");
        resolve();
      }
    });
  }

  public setupOpenAPIServer(): void {
    if (!this.OAPIServerEnabled) {
      this.logger.debug("OpenAPI server is disabled");
      return;
    }

    if (!this.BLOApplication) {
      this.logger.debug(
        "BLOApplication is not defined. Not initializing OpenAPI server",
      );
      return;
    }

    if (!this.OAS) {
      throw new Error("OpenAPI spec is not set");
    }
    // Type assertion here
    this.BLOApplication.use(
      "/api-docs",
      swaggerUi.serve as express.RequestHandler[],
      swaggerUi.setup(this.OAS) as express.RequestHandler,
    );
  }

  /**
   * Gateway Connection Methods
   * --------------------------
   * This section encompasses methods dedicated to establishing connections with gateways.
   * It includes functionalities to add gateways based on provided IDs and resolve specific gateway identities.
   * These operations are fundamental for setting up and managing gateway connections within the system.
   */

  // TODO: addGateways as an admin endpoint, simply calls orchestrator
  public async resolveAndAddGateways(IDs: string[]): Promise<void> {
    const fnTag = `${this.className}#resolveAndAddGateways()`;
    this.logger.trace(`Entering ${fnTag}`);
    this.logger.info("Connecting to gateway");
    this.gatewayOrchestrator.resolveAndAddGateways(IDs);

    // todo connect to gateway
  }

  public async addGateways(gateways: GatewayIdentity[]): Promise<void> {
    const fnTag = `${this.className}#addGateways()`;
    this.logger.trace(`Entering ${fnTag}`);
    this.logger.info("Connecting to gateway");
    this.gatewayOrchestrator.addGateways(gateways);

    // todo connect to gateway
  }

  /**
   * Shutdown Methods
   * -----------------
   * This section includes methods responsible for cleanly shutting down the server and its associated services.
   */
  public onShutdown(hook: ShutdownHook): void {
    const fnTag = `${this.className}#onShutdown()`;
    this.logger.trace(`Entering ${fnTag}`);
    this.logger.debug(`Adding shutdown hook: ${hook.name}`);
    this.shutdownHooks.push(hook);
  }

  public async shutdown(): Promise<void> {
    const fnTag = `${this.className}#getGatewaySeeds()`;
    this.logger.debug(`Entering ${fnTag}`);

    this.logger.info("Shutting down Node server - BOL");
    await this.shutdownBLOServer();

    this.logger.debug("Running shutdown hooks");
    for (const hook of this.shutdownHooks) {
      this.logger.debug(`Running shutdown hook: ${hook.name}`);
      await hook.hook();
    }

    this.logger.info("Shutting down Gateway Connection Manager");
    const connectionsClosed = await this.gatewayOrchestrator.disconnectAll();

    this.logger.info(`Closed ${connectionsClosed} connections`);
    this.logger.info("Gateway Coordinator shut down");
    return;
  }

  private async shutdownBLOServer(): Promise<void> {
    const fnTag = `${this.className}#shutdownBLOServer()`;
    this.logger.debug(`Entering ${fnTag}`);
    if (this.BLOServer) {
      try {
        await this.BLOServer.close();
        this.BLOServer = undefined;
        this.logger.info("Server shut down");
      } catch (error) {
        this.logger.error(
          `Error shutting down the gatewayApplication: ${error}`,
        );
      }
    } else {
      this.logger.warn("Server is not running.");
    }
  }
}
