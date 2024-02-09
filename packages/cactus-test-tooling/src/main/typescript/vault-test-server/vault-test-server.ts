import { EventEmitter } from "events";

import Docker, { Container } from "dockerode";

import {
  Logger,
  Checks,
  LogLevelDesc,
  LoggerProvider,
} from "@hyperledger/cactus-common";

import { Containers } from "../common/containers.js";

export interface IVaultTestServerOptions {
  envVars?: string[];
  imageVersion?: string;
  imageName?: string;
  logLevel?: LogLevelDesc;
}

export const K_DEFAULT_VAULT_IMAGE_NAME = "vault";
export const K_DEFAULT_VAULT_IMAGE_VERSION = "1.6.1";
export const K_DEFAULT_VAULT_HTTP_PORT = 8200;
export const K_DEFAULT_VAULT_DEV_ROOT_TOKEN = "myroot";

/**
 * > **Do not use this image in production.**
 *
 * Class responsible for programmatically managing a container that is running
 * the official Hashicorp Vault image in development mode.
 */
export class VaultTestServer {
  public static readonly CLASS_NAME = "VaultTestServer";

  private readonly log: Logger;
  private readonly imageName: string;
  private readonly imageVersion: string;
  private readonly envVars: string[];

  private _container: Container | undefined;
  private _containerId: string | undefined;

  public get containerId(): string {
    Checks.nonBlankString(this._containerId, `${this.className}:_containerId`);
    return this._containerId as string;
  }

  public get container(): Container {
    Checks.nonBlankString(this._container, `${this.className}:_container`);
    return this._container as Container;
  }

  public get className(): string {
    return VaultTestServer.CLASS_NAME;
  }

  public get imageFqn(): string {
    return `${this.imageName}:${this.imageVersion}`;
  }

  constructor(public readonly options: IVaultTestServerOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });

    this.imageName = this.options.imageName || K_DEFAULT_VAULT_IMAGE_NAME;
    this.imageVersion =
      this.options.imageVersion || K_DEFAULT_VAULT_IMAGE_VERSION;
    this.envVars = this.options.envVars || [];

    this.log.info(`Created ${this.className} OK. Image FQN: ${this.imageFqn}`);
  }

  public async start(): Promise<Container> {
    if (this._container) {
      await this._container.stop();
      await this._container.remove();
    }
    const docker = new Docker();

    await Containers.pullImage(this.imageFqn, {}, this.options.logLevel);

    return new Promise<Container>((resolve, reject) => {
      const eventEmitter: EventEmitter = docker.run(
        this.imageFqn,
        [],
        [],
        {
          ExposedPorts: {
            [`${K_DEFAULT_VAULT_HTTP_PORT}/tcp`]: {},
          },
          // This is a workaround needed for macOS which has issues with routing
          // to docker container's IP addresses directly...
          // https://stackoverflow.com/a/39217691
          PublishAllPorts: true,
          Env: [
            `VAULT_DEV_ROOT_TOKEN_ID=${K_DEFAULT_VAULT_DEV_ROOT_TOKEN}`,
            ...this.envVars,
          ],
          HostConfig: {
            // NetworkMode: "host",
            CapAdd: ["IPC_LOCK"],
            PublishAllPorts: true,
          },
          Healthcheck: {
            Test: ["CMD-SHELL", "wget -O- http://127.0.0.1:8200/v1/sys/health"],
            Interval: 100 * 1000000,
          },
        },
        {},
        (err: Error) => {
          if (err) {
            reject(err);
          }
        },
      );

      eventEmitter.once("start", async (container: Container) => {
        this._container = container;
        this._containerId = container.id;
        try {
          await Containers.waitForHealthCheck(this._containerId);
          resolve(container);
        } catch (ex) {
          reject(ex);
        }
      });
    });
  }

  public async stop(): Promise<void> {
    await Containers.stop(this._container as Docker.Container);
  }

  public destroy(): Promise<void> {
    const fnTag = `${this.className}#destroy()`;
    if (this._container) {
      return this._container.remove();
    } else {
      const ex = new Error(`${fnTag} Container not found, nothing to destroy.`);
      return Promise.reject(ex);
    }
  }

  public async getHostPortHttp(): Promise<number> {
    const fnTag = `${this.className}#getHostPortHttp()`;
    if (this._containerId) {
      const cInfo = await Containers.getById(this._containerId);
      return Containers.getPublicPort(K_DEFAULT_VAULT_HTTP_PORT, cInfo);
    } else {
      throw new Error(`${fnTag} Container ID not set. Did you call start()?`);
    }
  }
}
