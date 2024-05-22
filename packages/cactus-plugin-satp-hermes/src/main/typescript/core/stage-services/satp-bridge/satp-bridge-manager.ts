// this file contains a class that encapsulates the logic for managing the SATP bridge (lock, unlock, etc).
// should inject satp gateway session data (having parameters/chains for transactions), and processes smart contract output
import { BridgeManager } from "./bridge-manager";
import { Logger, LoggerProvider } from "@hyperledger/cactus-common";
import { SATPBridgeConfig } from "../../types";

export class SatpBridgeManager implements BridgeManager {
  public static readonly CLASS_NAME = "FabricBridgeManager";

  private _log: Logger;

  public get log(): Logger {
    return this._log;
  }

  constructor(private config: SATPBridgeConfig) {
    const level = "INFO";
    const label = SatpBridgeManager.CLASS_NAME;
    this._log = LoggerProvider.getOrCreate({ level, label });
  }

  public get className(): string {
    return SatpBridgeManager.CLASS_NAME;
  }

  public async lockAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#lockAsset()`;

    const transactionId = (
      await this.config.network.runTransaction("LockAsset", [assetId])
    ).transactionId;

    if (transactionId == undefined) {
      throw new Error(`${fnTag}, transactionId is undefined`);
    }

    const receipt = await this.config.network.getReceipt(
      assetId,
      transactionId,
    );

    this.log.info(`${fnTag}, proof of the asset lock: ${receipt}`);

    return receipt;
  }

  public async unlockAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#unlockAsset()`;

    const transactionId = (
      await this.config.network.runTransaction("UnlockAsset", [assetId])
    ).transactionId;

    if (transactionId == undefined) {
      throw new Error(`${fnTag}, transactionId is undefined`);
    }

    const receipt = await this.config.network.getReceipt(
      assetId,
      transactionId,
    );

    this.log.info(`${fnTag}, proof of the asset unlock: ${receipt}`);

    return receipt;
  }

  public async mintAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#mintAsset()`;

    const transactionId = (
      await this.config.network.runTransaction("MintAsset", [assetId])
    ).transactionId;

    if (transactionId == undefined) {
      throw new Error(`${fnTag}, transactionId is undefined`);
    }

    const receipt = await this.config.network.getReceipt(
      assetId,
      transactionId,
    );

    this.log.info(`${fnTag}, proof of the asset creation: ${receipt}`);

    return receipt;
  }

  public async burnAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#burnAsset()`;

    const transactionId = (
      await this.config.network.runTransaction("BurnAsset", [assetId])
    ).transactionId;

    if (transactionId == undefined) {
      throw new Error(`${fnTag}, transactionId is undefined`);
    }

    const receipt = await this.config.network.getReceipt(
      assetId,
      transactionId,
    );

    this.log.info(`${fnTag}, proof of the asset deletion: ${receipt}`);

    return receipt;
  }

  public async assignAsset(
    assetId: string,
    recipient: string,
  ): Promise<string> {
    const fnTag = `${this.className}#assignAsset()`;

    const transactionId = (
      await this.config.network.runTransaction("AssignAsset", [
        assetId,
        recipient,
      ])
    ).transactionId;

    if (transactionId == undefined) {
      throw new Error(`${fnTag}, transactionId is undefined`);
    }

    const receipt = await this.config.network.getReceipt(
      assetId,
      transactionId,
    );

    this.log.info(`${fnTag}, proof of the asset assignment: ${receipt}`);

    return receipt;
  }
  public async verifyAssetExistence(
    assetId: string,
  ): Promise<boolean | undefined> {
    //todo: implement this
    const assetExists = await this.config.network.runTransaction(
      "AssetExists",
      [assetId],
    );

    if (assetExists == undefined) {
      return false;
    }

    return true;
  }
  public async verifyLockAsset(assetId: string): Promise<boolean | undefined> {
    //todo: implement this
    const lockAsset = await this.config.network.runTransaction("LockAsset", [
      assetId,
    ]);

    if (lockAsset.output == undefined) {
      return false;
    }

    return true;
  }
}
