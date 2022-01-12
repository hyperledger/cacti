import { Express, Request, Response } from "express";

import {
  Logger,
  Checks,
  LogLevelDesc,
  LoggerProvider,
  IAsyncProvider,
  LogHelper,
} from "@hyperledger/cactus-common";
import {
  IEndpointAuthzOptions,
  IExpressRequestHandler,
  IWebServiceEndpoint,
} from "@hyperledger/cactus-core-api";
import { registerWebServiceEndpoint } from "@hyperledger/cactus-core";

import { PluginLedgerConnectorXdai } from "../plugin-ledger-connector-xdai";

import OAS from "../../json/openapi.json";
import axios from "axios";

export interface IRunTransactionEndpointOptions {
  logLevel?: LogLevelDesc;
  connector: PluginLedgerConnectorXdai;
}

export class RunTransactionEndpoint implements IWebServiceEndpoint {
  public static readonly CLASS_NAME = "RunTransactionEndpoint";

  private readonly log: Logger;

  public get className(): string {
    return RunTransactionEndpoint.CLASS_NAME;
  }

  constructor(public readonly options: IRunTransactionEndpointOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    Checks.truthy(options.connector, `${fnTag} arg options.connector`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });
  }

  public get oasPath(): typeof OAS.paths["/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-xdai/run-transaction"] {
    return OAS.paths[
      "/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-xdai/run-transaction"
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
    const reqBody = req.body;
    try {
      const resBody = await this.options.connector.transact(reqBody);
      res.json({ success: true, data: resBody });
    } catch (ex: unknown) {
      const stack = LogHelper.getExceptionStack(ex);
      const messages = LogHelper.getExceptionMessage(ex);
      this.log.error(`Crash while serving ${reqTag}`, ex);
      if (axios.isAxiosError(ex)) {
        res.status(500).json({
          message: "Internal Server Error",
          error: stack || messages,
        });
      } else {
        res.status(500).json({
          message: "Internal Server Error",
          error: stack || messages,
        });
      }
    }
  }
}
