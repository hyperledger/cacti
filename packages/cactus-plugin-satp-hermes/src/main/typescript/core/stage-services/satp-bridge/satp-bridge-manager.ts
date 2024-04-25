// this file contains a class that encapsulates the logic for managing the SATP bridge (lock, unlock, etc).
// should inject satp gateway session data (having parameters/chains for transactions), and processes smart contract output

import { SATPLedgerConnector } from "../../../types/blockchain-interaction";

export abstract class SATPBridgeManager {
  // Instantiate connectors based on supported implementations, at bridge level
  private connectors: SATPLedgerConnector[] = [];

  /**
   * Locks an asset on the blockchain
   * @param sessionId
   * @param assetId
   * @returns transaction hash
   **/
  public abstract lockAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string>;
  /**
   * Unlocks an asset on the blockchain
   * @param sessionId
   * @param assetId
   * @returns transaction hash
   **/
  public abstract unlockAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string>;
  /**
   * Mints an asset on the blockchain
   * @param sessionId
   * @param assetId
   * @returns transaction hash
   **/
  public abstract mintAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string>;
  /**
   * Burns an asset on the blockchain
   * @param sessionId
   * @param assetId
   * @returns transaction hash
   **/
  public abstract burnAsset(
    sessionId: string,
    assetId: string,
  ): Promise<string>;
  /**
   * Assigns an asset to a recipient on the blockchain
   * @param sessionId
   * @param assetId
   * @param recipient
   * @returns transaction hash
   **/
  public abstract assignAsset(
    sessionId: string,
    assetId: string,
    recipient: string,
  ): Promise<string>;
  /**
   * Verifies the existence of an asset on the blockchain
   * @param assetId
   * @returns boolean
   **/

  //create-rollback
  public abstract verifyAssetExistence(
    assetId: string,
  ): Promise<boolean | undefined>;
  /**
   * Verifies if an asset is locked on the blockchain
   * @param assetId
   * @returns boolean
   **/
  public abstract verifyLockAsset(
    assetId: string,
  ): Promise<boolean | undefined>;
}
