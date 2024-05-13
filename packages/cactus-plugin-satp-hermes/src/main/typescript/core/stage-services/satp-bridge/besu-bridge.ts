import {
  BesuConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";
import {
  EthContractInvocationType,
  IPluginLedgerConnectorBesuOptions,
  PluginLedgerConnectorBesu,
} from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { NetworkBridge } from "./network-bridge";
import { PluginBungeeHermes } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/plugin-bungee-hermes";
import { StrategyBesu } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/strategy/strategy-besu";

export class BesuBridge implements NetworkBridge {
  network: string = "BESU";

  connector: PluginLedgerConnectorBesu;
  bungee: PluginBungeeHermes;
  options: IPluginLedgerConnectorBesuOptions;
  config: BesuConfig;

  constructor(besuConfig: BesuConfig) {
    this.config = besuConfig;
    this.options = besuConfig.options;
    this.connector = new PluginLedgerConnectorBesu(besuConfig.options);
    this.bungee = new PluginBungeeHermes(besuConfig.bungeeOptions);
    this.bungee.addStrategy(this.network, new StrategyBesu("INFO"));
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
      transactionId: response.transactionReceipt?.transactionHash ?? "",
      output: response.callOutput.toString(),
    };
  }

  public async getReceipt(transactionHash: string): Promise<string> {
    const networkDetails = {
      connectorApiPath: "", //todo check this to not use api
      signingCredential: this.config.signingCredential,
      contractName: this.config.contractName,
      contractAddress: this.config.contractAddress,
      keychainId: this.config.keychainId,
      participant: "Org1MSP",
    };

    const snapshot = await this.bungee.generateSnapshot(
      [],
      this.network,
      networkDetails,
    );

    const view = this.bungee.generateView(
      snapshot,
      "0",
      Number.MAX_SAFE_INTEGER.toString(),
      undefined,
    );

    // process view

    if (view.view == undefined) {
      throw new Error("View is undefined");
    }

    return view.view.getViewStr();
  }
}
