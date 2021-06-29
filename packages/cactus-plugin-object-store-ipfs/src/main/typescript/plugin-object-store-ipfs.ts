import path from "path";
import type { Server } from "http";
import type { Server as SecureServer } from "https";
import type { Express } from "express";
import { create } from "ipfs-http-client";
import type { Options } from "ipfs-http-client";
import { Optional } from "typescript-optional";
import { RuntimeError } from "run-time-error";
import { Logger, Checks, LoggerProvider } from "@hyperledger/cactus-common";
import type { LogLevelDesc } from "@hyperledger/cactus-common";
import {
  IPluginObjectStore,
  ICactusPluginOptions,
  IWebServiceEndpoint,
  GetObjectRequestV1,
  GetObjectResponseV1,
  HasObjectRequestV1,
  HasObjectResponseV1,
  SetObjectRequestV1,
  SetObjectResponseV1,
  RemoveObjectRequestV1,
  RemoveObjectResponseV1,
  // IPluginKeychain,
} from "@hyperledger/cactus-core-api";

import { GetObjectEndpointV1 } from "./web-services/get-object-endpoint-v1";
import { SetObjectEndpointV1 } from "./web-services/set-object-endpoint-v1";
import { HasObjectEndpointV1 } from "./web-services/has-object-endpoint-v1";
import { RemoveObjectEndpointV1 } from "./web-services/remove-object-endpoint-v1";
import type { IIpfsHttpClient } from "./i-ipfs-http-client";
import { isIpfsHttpClientOptions } from "./i-ipfs-http-client";
import { PluginRegistry } from "@hyperledger/cactus-core";
import { PluginKeychainMemory } from "@hyperledger/cactus-plugin-keychain-memory";
// import { v4 as uuidv4 } from "uuid";

export const K_IPFS_JS_HTTP_ERROR_FILE_DOES_NOT_EXIST =
  "HTTPError: file does not exist";

export interface IPluginObjectStoreIpfsOptions extends ICactusPluginOptions {
  readonly logLevel?: LogLevelDesc;
  readonly parentDir: string;
  readonly ipfsClientOrOptions: Options | IIpfsHttpClient;
  readonly pluginRegistry?: PluginRegistry;
  readonly keychainId?: string;
  readonly keychain?: PluginKeychainMemory;
}

export class PluginObjectStoreIpfs implements IPluginObjectStore {
  public static readonly CLASS_NAME = "PluginObjectStoreIpfs";

  private readonly ipfs: IIpfsHttpClient;
  private readonly log: Logger;
  private readonly instanceId: string;
  private readonly parentDir: string;

  public get className(): string {
    return PluginObjectStoreIpfs.CLASS_NAME;
  }

  constructor(public readonly opts: IPluginObjectStoreIpfsOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(opts, `${fnTag} arg options`);
    Checks.nonBlankString(opts.instanceId, `${fnTag} options.instanceId`);
    Checks.nonBlankString(opts.parentDir, `${fnTag} options.parentDir`);
    Checks.truthy(opts.ipfsClientOrOptions, `${fnTag} ipfsClientOrOptions`);

    if (isIpfsHttpClientOptions(opts.ipfsClientOrOptions)) {
      this.ipfs = opts.ipfsClientOrOptions;
    } else if (opts.ipfsClientOrOptions) {
      this.ipfs = create({
        ...(this.opts.ipfsClientOrOptions as Options),
      });
    } else {
      const errorMessage = `${fnTag} Need either "ipfsClient" or "ipfsClientOptions" to construct ${this.className} Neither was provided.`;
      throw new RuntimeError(errorMessage);
    }
    Checks.truthy(this.ipfs, `${fnTag} arg options.backend`);

    const level = this.opts.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });

    this.parentDir = this.opts.parentDir;
    this.instanceId = this.opts.instanceId;

    this.log.info(`Created ${this.className}. InstanceID=${opts.instanceId}`);
  }

  public async onPluginInit(): Promise<unknown> {
    return;
  }

  public async registerWebServices(
    app: Express,
  ): Promise<IWebServiceEndpoint[]> {
    const webServices = await this.getOrCreateWebServices();
    await Promise.all(webServices.map((ws) => ws.registerExpress(app)));
    return webServices;
  }

  public async getOrCreateWebServices(): Promise<IWebServiceEndpoint[]> {
    const { log, opts } = this;
    const { logLevel } = opts;

    log.info(`Installing web services for plugin ${this.getPackageName()}...`);

    const endpoints: IWebServiceEndpoint[] = [
      new GetObjectEndpointV1({
        logLevel,
        plugin: this,
      }),
      new SetObjectEndpointV1({
        logLevel,
        plugin: this,
      }),
      new HasObjectEndpointV1({
        logLevel,
        plugin: this,
      }),
      new RemoveObjectEndpointV1({
        logLevel,
        plugin: this,
      }),
    ];
    const pkg = this.getPackageName();
    log.info(`Installed web services for plugin ${pkg} OK`, { endpoints });

    return endpoints;
  }

  public getInstanceId(): string {
    return this.instanceId;
  }

  public getPackageName(): string {
    return `@hyperledger/cactus-plugin-object-store-ipfs`;
  }

  public getKeyPath(req: { key: string }): string {
    return path.join(this.parentDir, req.key);
  }

  public async get(req: GetObjectRequestV1): Promise<GetObjectResponseV1> {
    const keyPath = this.getKeyPath(req);
    const chunksIterable = this.ipfs.files.read(keyPath);
    const chunks = [];
    for await (const chunk of chunksIterable) {
      chunks.push(chunk);
    }
    const totalLength = chunks.reduce((sum, it) => sum + it.length, 0);
    const array = new Uint8Array(totalLength);

    chunks.reduce((sum, it) => {
      array.set(it, sum);
      return sum + it.length;
    }, 0);

    const value = Buffer.from(array).toString("base64");
    return { key: req.key, value };
  }

  public async has(req: HasObjectRequestV1): Promise<HasObjectResponseV1> {
    const checkedAt = new Date().toJSON();
    const keyPath = this.getKeyPath(req);
    try {
      const statResult = await this.ipfs.files.stat(keyPath);
      this.log.debug(`StatResult for ${req.key}: %o`, statResult);
      this.log.debug(`log for statresult.cid: %o`, statResult.cid.toString());
      return { key: req.key, checkedAt, isPresent: true };
    } catch (ex) {
      if (ex?.stack?.includes(K_IPFS_JS_HTTP_ERROR_FILE_DOES_NOT_EXIST)) {
        const msg = `Stat ${req.key} failed with error message containing phrase "${K_IPFS_JS_HTTP_ERROR_FILE_DOES_NOT_EXIST}" Returning isPresent=false ...`;
        this.log.debug(msg);
        return { key: req.key, checkedAt, isPresent: false };
      } else {
        throw new RuntimeError(`Checking presence of ${req.key} crashed:`, ex);
      }
    }
  }

  public async set(req: SetObjectRequestV1): Promise<SetObjectResponseV1> {
    const keyPath = this.getKeyPath(req);
    try {
      this.log.debug(`Setting object ${keyPath} in IPFS...`);
      const buffer = Buffer.from(req.value, "base64");
      await this.ipfs.files.write(keyPath, buffer, {
        create: true,
        parents: true,
      });
    } catch (ex) {
      throw new RuntimeError(`Can't set object ${keyPath}. Write failed:`, ex);
    }

    const statResult = await this.ipfs.files.stat(keyPath);

    return {
      key: req.key,
      cidHash: statResult.cid.toString(),
    } as SetObjectResponseV1;
  }

  // eslint-disable-next-line prettier/prettier
  public async remove( req: RemoveObjectRequestV1 ): Promise<RemoveObjectResponseV1> {
    const keyPath = this.getKeyPath(req);
    try {
      const removeResult = await this.ipfs.files.rm(keyPath);
      this.log.debug(`RemoveResult for ${req.key}: %o`, removeResult);
      return { key: req.key, isRemoved: true };
    } catch (ex) {}
    return { key: req.key, isRemoved: false };
  }

  public getHttpServer(): Optional<Server | SecureServer> {
    return Optional.empty();
  }

  public async shutdown(): Promise<void> {
    this.log.info(`Shutting down ${this.className}...`);
  }
}
/*A few things I'd change to make the keychain an optional dependency not a hard one (otherwise it's not really a plugin).
	
1) Don't depend on the keychain memory plugin in package.json,
2) use the keychain interface from the core api package instead of the in-memory implementation
3) Don't have a keychain instance as a constructor arg's property, use an optional keychainId request parameter to look up keychain instances in the plugin registry
*/
