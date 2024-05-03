import {
  FabricContractInvocationType,
  IPluginLedgerConnectorFabricOptions,
  PluginLedgerConnectorFabric,
} from "@hyperledger/cactus-plugin-ledger-connector-fabric";
import { NetworkBridge } from "../../types";
import {
  FabricConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";

export class FabricBridge implements NetworkBridge {
  network: string = "fabric";

  connector: PluginLedgerConnectorFabric;
  options: IPluginLedgerConnectorFabricOptions;
  config: FabricConfig;

  constructor(
    fabricOptions: IPluginLedgerConnectorFabricOptions,
    fabricConfig: FabricConfig,
  ) {
    this.options = fabricOptions;
    this.config = fabricConfig;
    this.connector = new PluginLedgerConnectorFabric(fabricOptions);
  }

  public networkName(): string {
    return this.network;
  }

  public async runTransaction(
    methodName: string,
    params: string[],
  ): Promise<TransactionResponse> {
    const response = await this.connector.transact({
      signingCredential: this.config.signingCredential,
      channelName: this.config.channelName,
      methodName: methodName,
      params: params,
      contractName: this.config.contractName,
      invocationType: FabricContractInvocationType.Send,
    });
    return {
      transaction: response.transactionId,
      receipt: undefined,
      output: JSON.stringify(response),
    };
  }

  public async getReceipt(transactionId: string): Promise<string> {
    const response = await this.connector.transact({
      signingCredential: this.config.signingCredential,
      channelName: this.config.channelName,
      methodName: "GetBlockByTxID",
      params: [this.config.channelName, transactionId],
      contractName: this.config.contractName,
      invocationType: FabricContractInvocationType.Call,
    });
    return response.functionOutput;
  }
}
