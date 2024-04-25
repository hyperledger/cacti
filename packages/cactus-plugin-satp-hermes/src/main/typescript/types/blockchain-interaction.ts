// TODO: define alias types for SATPLedgerConnector, which encapsulates IPluginLedgerConnector

import { FabricSigningCredential } from "@hyperledger/cactus-plugin-ledger-connector-fabric";

import { DefaultApi as FabricApi } from "@hyperledger/cactus-plugin-ledger-connector-fabric";
import { DefaultApi as BesuApi, Web3SigningCredential } from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { BesuSignTransactionEndpointV1 } from "@hyperledger/cactus-plugin-ledger-connector-besu/src/main/typescript/web-services/sign-transaction-endpoint-v1";

// inject gateway, get connectors
export type SATPLedgerConnector = string;

// TODO Define lock interfaces and strategy pattern for locking (as function of locking blockchain) (see what smart contract implementations return)

export interface FabricConfig {
  fabricApi: FabricApi;
  signingCredential: FabricSigningCredential;
  channelName: string;
  contractName: string;
}
export interface BesuConfig {
  besuApi: BesuApi;
  signingCredential: Web3SigningCredential;
  contractName: string;
  keyChainId: string;
}
