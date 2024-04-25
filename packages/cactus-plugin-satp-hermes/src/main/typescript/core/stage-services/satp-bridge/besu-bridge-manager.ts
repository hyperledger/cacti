import { Logger, LoggerProvider } from "@hyperledger/cactus-common";
import { SATPBridgeManager } from "./satp-bridge-manager";
import { SATPGateway } from "../../../gateway-refactor";
import { BesuConfig } from "../../../types/blockchain-interaction";
import { EthContractInvocationType } from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { InvokeContractV1Request as BesuInvokeContractV1Request } from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { storeLog } from "../../../gateway-utils";

export class BesuBridgeManager implements SATPBridgeManager {
  public static readonly CLASS_NAME = "BesuBridgeManager";

  private _log: Logger;

  public get log(): Logger {
    return this._log;
  }
  private gateway: SATPGateway;

  private besuConfig: BesuConfig;

  public get className(): string {
    return BesuBridgeManager.CLASS_NAME;
  }

  public constructor(gateway: SATPGateway, besuConfig: BesuConfig) {
    this.gateway = gateway;
    this.besuConfig = besuConfig;
    const level = "INFO";
    const label = BesuBridgeManager.CLASS_NAME;
    this._log = LoggerProvider.getOrCreate({ level, label });
  }

  public async lockAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#lockAssetBesu()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }
    let besuLockAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "lock-asset",
      data: JSON.stringify(sessionData),
    });

    const besuAssetLock = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: "lockAssetReference",
      gas: 1000000,
      params: [assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (besuAssetLock.status != 200) {
      throw new Error(`${fnTag}, besu lock asset error`);
    }

    const besuAssetLockDataJson = JSON.parse(
      JSON.stringify(besuAssetLock.data),
    );

    if (besuAssetLockDataJson.out == undefined) {
      throw new Error(`${fnTag}, besu res data out undefined`);
    }

    if (besuAssetLockDataJson.out.transactionReceipt == undefined) {
      throw new Error(`${fnTag}, undefined besu transact receipt`);
    }

    const besuAssetLockReceipt = besuAssetLockDataJson.out.transactionReceipt;
    besuLockAssetProof = JSON.stringify(besuAssetLockReceipt);

    sessionData.lockAssertionClaim = besuLockAssetProof;

    this.log.info(`${fnTag}, proof of the asset lock: ${besuLockAssetProof}`);

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "lock",
      data: besuLockAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "lock-asset",
      data: JSON.stringify(sessionData),
    });

    return besuLockAssetProof;
  }
  public async unlockAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string> {
    const fnTag = `${this.className}#unlockAssetBesu()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }

    let besuUnlockAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec-rollback",
      operation: "unlock-asset",
      data: JSON.stringify(sessionData),
    });

    const assetUnlockResponse = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: "unLockAssetReference",
      gas: 1000000,
      params: [assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (assetUnlockResponse.status != 200) {
      throw new Error(`${fnTag}, besu unlock asset error`);
    }

    const assetUnlockResponseDataJson = JSON.parse(
      JSON.stringify(assetUnlockResponse.data),
    );

    if (assetUnlockResponseDataJson.out == undefined) {
      throw new Error(`${fnTag}, besu res data out undefined`);
    }

    if (assetUnlockResponseDataJson.out.transactionReceipt == undefined) {
      throw new Error(`${fnTag}, undefined besu transact receipt`);
    }

    const besuCreateAssetReceipt =
      assetUnlockResponseDataJson.out.transactionReceipt;
    besuUnlockAssetProof = JSON.stringify(besuCreateAssetReceipt);

    // sessionData.rollbackActionsPerformed.push(
    //   SessionDataRollbackActionsPerformedEnum.Lock,
    // );
    // sessionData.rollbackProofs.push(besuUnlockAssetProof);

    this.log.info(
      `${fnTag}, proof of the asset unlock: ${besuUnlockAssetProof}`,
    );

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "proof-rollback",
      operation: "unlock",
      data: besuUnlockAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done-rollback",
      operation: "unlock-asset",
      data: JSON.stringify(sessionData),
    });

    return besuUnlockAssetProof;
  }
  public async mintAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#mintAssetBesu()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined || sessionData.assetProfile == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }

    let besuCreateAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "mint-asset",
      data: JSON.stringify(sessionData),
    });

    const amount = sessionData.assetProfile.keyInformationLink[0].toString();
    const userEthAddress =
      sessionData.assetProfile.keyInformationLink[2].toString();

    const besuCreateRes = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: "mint",
      gas: 1000000,
      params: [userEthAddress, amount, assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (besuCreateRes.status != 200) {
      throw new Error(`${fnTag}, besu create asset error`);
    }

    const besuCreateResDataJson = JSON.parse(
      JSON.stringify(besuCreateRes.data),
    );

    if (besuCreateResDataJson.out == undefined) {
      throw new Error(`${fnTag}, besu res data out undefined`);
    }

    if (besuCreateResDataJson.out.transactionReceipt == undefined) {
      throw new Error(`${fnTag}, undefined besu transact receipt`);
    }

    const besuCreateAssetReceipt = besuCreateResDataJson.out.transactionReceipt;
    besuCreateAssetProof = JSON.stringify(besuCreateAssetReceipt);

    sessionData.mintAssertionClaims = besuCreateAssetProof;

    this.log.info(
      `${fnTag}, proof of the asset creation: ${besuCreateAssetProof}`,
    );

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "create",
      data: besuCreateAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "create-asset",
      data: JSON.stringify(sessionData),
    });

    return besuCreateAssetProof;
  }
  public async burnAsset(sessionId: string, assetId: string): Promise<string> {
    const fnTag = `${this.className}#burnAssetBesu()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }

    let besuDeleteAssetProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "delete-asset",
      data: JSON.stringify(sessionData),
    });

    const besuAssetDeletion = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: "deleteAssetReference",
      gas: 1000000,
      params: [assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (besuAssetDeletion.status != 200) {
      throw new Error(`${fnTag}, besu delete asset error`);
    }

    const besuAssetDeletionDataJson = JSON.parse(
      JSON.stringify(besuAssetDeletion.data),
    );

    if (besuAssetDeletionDataJson.out == undefined) {
      throw new Error(`${fnTag}, besu res data out undefined`);
    }

    if (besuAssetDeletionDataJson.out.transactionReceipt == undefined) {
      throw new Error(`${fnTag}, undefined besu transact receipt`);
    }

    const besuAssetDeletionReceipt =
      besuAssetDeletionDataJson.out.transactionReceipt;
    besuDeleteAssetProof = JSON.stringify(besuAssetDeletionReceipt);

    sessionData.burnAssertionClaim = besuDeleteAssetProof;

    this.log.info(
      `${fnTag}, proof of the asset deletion: ${besuDeleteAssetProof}`,
    );

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "delete",
      data: besuDeleteAssetProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "delete-asset",
      data: JSON.stringify(sessionData),
    });

    return besuDeleteAssetProof;
  }
  public async assignAsset(
    sessionId: string,
    assetId: string,
    recipient: string,
  ): Promise<string> {
    const fnTag = `${this.className}#burnAssetBesu()`;

    const sessionData = this.gateway.getSession(sessionId);

    if (sessionData == undefined) {
      throw new Error(`${fnTag}, session data is not correctly initialized`);
    }

    let besuAssignProof = "";

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "exec",
      operation: "assign-asset",
      data: JSON.stringify(sessionData),
    });

    const besuAssetDeletion = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: "assignAsset",
      gas: 1000000,
      params: [assetId, recipient],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (besuAssetDeletion.status != 200) {
      throw new Error(`${fnTag}, besu delete asset error`);
    }

    const besuAssetDeletionDataJson = JSON.parse(
      JSON.stringify(besuAssetDeletion.data),
    );

    if (besuAssetDeletionDataJson.out == undefined) {
      throw new Error(`${fnTag}, besu res data out undefined`);
    }

    if (besuAssetDeletionDataJson.out.transactionReceipt == undefined) {
      throw new Error(`${fnTag}, undefined besu transact receipt`);
    }

    const besuAssetAssignReceipt =
      besuAssetDeletionDataJson.out.transactionReceipt;
    besuAssignProof = JSON.stringify(besuAssetAssignReceipt);

    sessionData.burnAssertionClaim = besuAssignProof;

    this.log.info(`${fnTag}, proof of the asset deletion: ${besuAssignProof}`);

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "proof",
      operation: "delete",
      data: besuAssignProof,
    });

    await storeLog(this.gateway, {
      sessionID: sessionId,
      type: "done",
      operation: "delete-asset",
      data: JSON.stringify(sessionData),
    });

    return besuAssignProof;
  }
  public async verifyAssetExistence(
    assetId: string,
  ): Promise<boolean | undefined> {
    const fnTag = `${this.className}#verifyAssetExistenceBesu()`;

    const assetExists = await this.besuConfig.besuApi?.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Call,
      methodName: "isPresent",
      gas: 1000000,
      params: [assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as BesuInvokeContractV1Request);

    if (assetExists == undefined) {
      throw new Error(`${fnTag} the asset does not exist`);
    }

    return assetExists?.data.callOutput == true;
  }
  public async verifyLockAsset(assetId: string): Promise<boolean | undefined> {
    const fnTag = `${this.className}#verifyLockAssetBesu()`;

    const assetIsLocked = await this.besuConfig.besuApi.invokeContractV1({
      contractName: this.besuConfig.contractName,
      invocationType: EthContractInvocationType.Call,
      methodName: "isAssetLocked",
      gas: 1000000,
      params: [assetId],
      signingCredential: this.besuConfig.signingCredential,
      keychainId: this.besuConfig.keyChainId,
    } as unknown as BesuInvokeContractV1Request);

    if (assetIsLocked == undefined) {
      throw new Error(`${fnTag} the asset does not exist`);
    }

    return assetIsLocked?.data.callOutput == true;
  }
}
