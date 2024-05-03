import {
  BesuConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";
import {
  EthContractInvocationType,
  IPluginLedgerConnectorBesuOptions,
  PluginLedgerConnectorBesu,
} from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { NetworkBridge } from "../../types";

export class BesuBridge implements NetworkBridge {
  network: string = "besu";

  connector: PluginLedgerConnectorBesu;
  options: IPluginLedgerConnectorBesuOptions;
  config: BesuConfig;

  constructor(
    besuOptions: IPluginLedgerConnectorBesuOptions,
    besuConfig: BesuConfig,
  ) {
    this.options = besuOptions;
    this.config = besuConfig;
    this.connector = new PluginLedgerConnectorBesu(besuOptions);
  }
  public networkName(): string {
    return this.network;
  }

  public async runTransaction(
    methodName: string,
    params: string[],
  ): Promise<TransactionResponse> {
    const response = await this.connector.invokeContract({
      contractName: this.config.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: methodName,
      params: params,
      signingCredential: this.config.signingCredential,
    });
    return {
      transaction: response.transactionReceipt?.transactionHash ?? "",
      receipt: JSON.stringify(response.transactionReceipt),
      output: JSON.stringify(response),
    };
  }
  public async getReceipt(transactionId: string): Promise<string> {
    const response = await this.connector.invokeContract({
      contractName: this.config.contractName,
      invocationType: EthContractInvocationType.Call,
      methodName: "GetBlockByTxID",
      params: [transactionId],
      signingCredential: this.config.signingCredential,
    });
    return response.callOutput;
  }
}
