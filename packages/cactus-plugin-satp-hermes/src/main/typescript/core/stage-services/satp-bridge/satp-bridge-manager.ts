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

    const response = await this.config.network.runTransaction("LockAsset", [
      assetId,
    ]);

    let receipt = response.receipt;

    if (receipt == undefined) {
      receipt = await this.config.network.getReceipt(response.transaction);
    }

    this.log.info(`${fnTag}, proof of the asset lock: ${receipt}`);

    return receipt;
  }

  public async unlockAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#unlockAsset()`;

    const response = await this.config.network.runTransaction("UnlockAsset", [
      assetId,
    ]);

    let receipt = response.receipt;

    if (receipt == undefined) {
      receipt = await this.config.network.getReceipt(response.transaction);
    }

    this.log.info(`${fnTag}, proof of the asset unlock: ${receipt}`);

    return receipt;
  }

  public async mintAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#mintAsset()`;

    const response = await this.config.network.runTransaction("MintAsset", [
      assetId,
    ]);

    let receipt = response.receipt;

    if (receipt == undefined) {
      receipt = await this.config.network.getReceipt(response.transaction);
    }

    this.log.info(`${fnTag}, proof of the asset creation: ${receipt}`);

    return receipt;
  }

  public async burnAsset(assetId: string): Promise<string> {
    const fnTag = `${this.className}#burnAsset()`;

    const response = await this.config.network.runTransaction("BurnAsset", [
      assetId,
    ]);

    let receipt = response.receipt;

    if (receipt == undefined) {
      receipt = await this.config.network.getReceipt(response.transaction);
    }

    this.log.info(`${fnTag}, proof of the asset deletion: ${receipt}`);

    return receipt;
  }

  public async assignAsset(
    assetId: string,
    recipient: string,
  ): Promise<string> {
    const fnTag = `${this.className}#assignAsset()`;

    const response = await this.config.network.runTransaction("AssignAsset", [
      assetId,
      recipient,
    ]);

    let receipt = response.receipt;

    if (receipt == undefined) {
      receipt = await this.config.network.getReceipt(response.transaction);
    }

    this.log.info(`${fnTag}, proof of the asset assignment: ${receipt}`);

    return receipt;
  }
  public async verifyAssetExistence(
    assetId: string,
  ): Promise<boolean | undefined> {
    const assetExists = await this.config.network.runTransaction(
      "AssetExists",
      [assetId],
    );

    if (assetExists.transaction == undefined) {
      return false;
    }

    return true;
  }
  public async verifyLockAsset(assetId: string): Promise<boolean | undefined> {
    const lockAsset = await this.config.network.runTransaction("LockAsset", [
      assetId,
    ]);

    if (lockAsset.output == undefined) {
      return false;
    }

    return true;
  }
}
