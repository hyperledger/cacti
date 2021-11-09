import { Express, Request, Response } from "express";

import {
  Logger,
  Checks,
  LogLevelDesc,
  LoggerProvider,
  IAsyncProvider,
} from "@hyperledger/cactus-common";
import {
  IEndpointAuthzOptions,
  IExpressRequestHandler,
  IWebServiceEndpoint,
} from "@hyperledger/cactus-core-api";
import { registerWebServiceEndpoint } from "@hyperledger/cactus-core";

import { PluginLedgerConnectorBesu } from "../plugin-ledger-connector-besu";

import OAS from "../../json/openapi.json";
import { InvokeContractJsonObjectV1Request } from "../generated/openapi/typescript-axios";

export interface IInvokeContractJsonObjectEndpointOptions {
  logLevel?: LogLevelDesc;
  connector: PluginLedgerConnectorBesu;
}

export class InvokeContractJsonObjectEndpoint implements IWebServiceEndpoint {
  public static readonly CLASS_NAME = "InvokeContractJsonObjectEndpoint";

  private readonly log: Logger;

  public get className(): string {
    return InvokeContractJsonObjectEndpoint.CLASS_NAME;
  }

  constructor(
    public readonly options: IInvokeContractJsonObjectEndpointOptions,
  ) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    Checks.truthy(options.connector, `${fnTag} arg options.connector`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });
  }

  public get oasPath(): typeof OAS.paths["/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-besu/invoke-contract-json-object"] {
    return OAS.paths[
      "/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-besu/invoke-contract-json-object"
    ];
  }

  public getPath(): string {
    return this.oasPath.post["x-hyperledger-cactus"].http.path;
  }

  public getVerbLowerCase(): string {
    return this.oasPath.post["x-hyperledger-cactus"].http.verbLowerCase;
  }

  public getOperationId(): string {
    return this.oasPath.post.operationId;
  }

  getAuthorizationOptionsProvider(): IAsyncProvider<IEndpointAuthzOptions> {
    // TODO: make this an injectable dependency in the constructor
    return {
      get: async () => ({
        isProtected: true,
        requiredRoles: [],
      }),
    };
  }

  public async registerExpress(
    expressApp: Express,
  ): Promise<IWebServiceEndpoint> {
    await registerWebServiceEndpoint(expressApp, this);
    return this;
  }

  public getExpressRequestHandler(): IExpressRequestHandler {
    return this.handleRequest.bind(this);
  }

  public async handleRequest(req: Request, res: Response): Promise<void> {
    const reqTag = `${this.getVerbLowerCase()} - ${this.getPath()}`;
    this.log.debug(reqTag);
    const reqBody: InvokeContractJsonObjectV1Request = req.body;
    try {
      const resBody = await this.options.connector.getContractInfoJsonObject(
        reqBody,
      );
      res.json(resBody);
    } catch (ex) {
      this.log.error(`Crash while serving ${reqTag}`, ex);
      res.status(500).json({
        message: "Internal Server Error",
        error: ex?.stack || ex?.message,
      });
    }
  }
}
