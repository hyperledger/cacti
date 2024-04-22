import {
  FabricContractInvocationType,
  RunTransactionRequest,
} from "@hyperledger/cactus-plugin-ledger-connector-fabric";
import { SATPGateway } from "../../../gateway-refactor";
import { storeLog, storeProof } from "../../../gateway-utils";
import { FabricConfig } from "../../../types/blockchain-interaction";
import { SATPBridgeManager } from "./satp-bridge-manager";
import { Logger, LoggerProvider } from "@hyperledger/cactus-common";

export class FabricBridgeManager implements SATPBridgeManager {
  public static readonly CLASS_NAME = "FabricBridgeManager";

  private _log: Logger;

  public get log(): Logger {
    return this._log;
  }
  private gateway: SATPGateway;

  private fabricConfig: FabricConfig;

  public get className(): string {
    return FabricBridgeManager.CLASS_NAME;
  }

  public constructor(gateway: SATPGateway, fabricConfig: FabricConfig) {
    this.gateway = gateway;
    this.fabricConfig = fabricConfig;
    const level = "INFO";
    const label = FabricBridgeManager.CLASS_NAME;
    this._log = LoggerProvider.getOrCreate({ level, label });
  }
  public async lockAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#defineFabricConnection()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == null) {
      throw new Error(`Session not found for ID ${sessionId}`);
    }

    let fabricLockAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "lock-asset",
      data: JSON.stringify(sessionData),
    });

    //todo: handle errors here
    const response = await this.fabricConfig.fabricApi.runTransactionV1({
      signingCredential: this.fabricConfig.signingCredential,
      channelName: this.fabricConfig.channelName,
      contractName: this.fabricConfig.contractName,
      invocationType: FabricContractInvocationType.Send,
      methodName: "LockAsset",
      params: [assetId],
    } as RunTransactionRequest);

    const receipt =
      await this.fabricConfig.fabricApi.getTransactionReceiptByTxIDV1({
        signingCredential: this.fabricConfig.signingCredential,
        channelName: this.fabricConfig.channelName,
        methodName: "GetBlockByTxID",
        params: [this.fabricConfig.channelName, response.data.transactionId],
      } as RunTransactionRequest);

    this.log.warn(receipt.data);

    fabricLockAssetProof = JSON.stringify(receipt.data);

    sessionData.lockAssertionClaim = fabricLockAssetProof;

    this.log.info(`${fnTag}, proof of the asset lock: ${fabricLockAssetProof}`);

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "lock",
      data: fabricLockAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "lock-asset",
      data: JSON.stringify(sessionData),
    });

    return fabricLockAssetProof;
  }

  public async unlockAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string> {
    const fnTag = `${this.className}#defineFabricConnection()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == null) {
      throw new Error(`Session not found for ID ${sessionId}`);
    }

    let fabricUnlockAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec-rollback",
      operation: "unlock-asset",
      data: JSON.stringify(sessionData),
    });

    const response = await this.fabricConfig.fabricApi.runTransactionV1({
      signingCredential: this.fabricConfig.signingCredential,
      channelName: this.fabricConfig.channelName,
      contractName: this.fabricConfig.contractName,
      invocationType: FabricContractInvocationType.Send,
      methodName: "UnlockAsset",
      params: [assetId],
    } as RunTransactionRequest);

    const receipt =
      await this.fabricConfig.fabricApi.getTransactionReceiptByTxIDV1({
        signingCredential: this.fabricConfig.signingCredential,
        channelName: this.fabricConfig.channelName,
        methodName: "GetBlockByTxID",
        params: [this.fabricConfig.channelName, response.data.transactionId],
      } as RunTransactionRequest);

    this.log.warn(receipt.data);

    fabricUnlockAssetProof = JSON.stringify(receipt.data);

    //sessionData.lockRollBackClaim = fabricUnlockAssetProof;
    //todo when implementing the rollback

    this.log.info(
      `${fnTag}, proof of the asset unlock: ${fabricUnlockAssetProof}`,
    );

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "unlock",
      data: fabricUnlockAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "unlock-asset",
      data: JSON.stringify(sessionData),
    });

    return fabricUnlockAssetProof;
  }

  public async mintAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#mintAsset()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == null) {
      throw new Error(`Session not found for ID ${sessionId}`);
    }

    let fabricMintAssetProof = "";

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "mint-asset",
      data: JSON.stringify(sessionData),
    });

    const response = await this.fabricConfig.fabricApi.runTransactionV1({
      contractName: this.fabricConfig.contractName,
      channelName: this.fabricConfig.channelName,
      params: [assetId],
      methodName: "CreateAsset",
      invocationType: FabricContractInvocationType.Send,
      signingCredential: this.fabricConfig.signingCredential,
    });

    const receiptCreate =
      await this.fabricConfig.fabricApi.getTransactionReceiptByTxIDV1({
        signingCredential: this.fabricConfig.signingCredential,
        channelName: this.fabricConfig.channelName,
        contractName: "qscc",
        invocationType: FabricContractInvocationType.Call,
        methodName: "GetBlockByTxID",
        params: [this.fabricConfig.channelName, response.data.transactionId],
      } as RunTransactionRequest);

    this.log.warn(receiptCreate.data);
    fabricMintAssetProof = JSON.stringify(receiptCreate.data);

    sessionData.mintAssertionClaims = fabricMintAssetProof;

    this.log.info(
      `${fnTag}, proof of the asset creation: ${fabricMintAssetProof}`,
    );

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "mint-asset",
      data: fabricMintAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "mint-asset",
      data: JSON.stringify(sessionData),
    });

    return fabricMintAssetProof;
  }

  public async burnAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#burnAsset()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }

    let fabricBurnAssetProof = "";

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "burn-asset",
      data: JSON.stringify(sessionData),
    });

    const burnRes = await this.fabricConfig.fabricApi.runTransactionV1({
      signingCredential: this.fabricConfig.signingCredential,
      channelName: this.fabricConfig.channelName,
      contractName: this.fabricConfig.contractName,
      invocationType: FabricContractInvocationType.Send,
      methodName: "DeleteAsset",
      params: [assetId],
    } as RunTransactionRequest);

    const receiptBurnRes =
      await this.fabricConfig.fabricApi.getTransactionReceiptByTxIDV1({
        signingCredential: this.fabricConfig.signingCredential,
        channelName: this.fabricConfig.channelName,
        contractName: "qscc",
        invocationType: FabricContractInvocationType.Call,
        methodName: "GetBlockByTxID",
        params: [this.fabricConfig.channelName, burnRes.data.transactionId],
      } as RunTransactionRequest);

    this.log.warn(receiptBurnRes.data);
    fabricBurnAssetProof = JSON.stringify(receiptBurnRes.data);

    sessionData.burnAssertionClaim = fabricBurnAssetProof;

    this.log.info(
      `${fnTag}, proof of the asset deletion: ${fabricBurnAssetProof}`,
    );

    await storeProof(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "burn-asset",
      data: fabricBurnAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "burn-asset",
      data: JSON.stringify(sessionData),
    });

    return fabricBurnAssetProof;
  }
  public async assignAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string> {
    throw new Error("Method not implemented.");
  }
  public verifyAssetExistence(assetId: string): Promise<string> {
    throw new Error("Method not implemented.");
  }
  public verifyLockAsset(assetId: string): Promise<string> {
    throw new Error("Method not implemented.");
  }
}
