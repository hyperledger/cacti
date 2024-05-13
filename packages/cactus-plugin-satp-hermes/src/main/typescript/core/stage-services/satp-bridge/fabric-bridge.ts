import {
  FabricContractInvocationType,
  IPluginLedgerConnectorFabricOptions,
  PluginLedgerConnectorFabric,
} from "@hyperledger/cactus-plugin-ledger-connector-fabric";
import { PluginBungeeHermes } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/plugin-bungee-hermes";
import { StrategyFabric } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/strategy/strategy-fabric";
import { NetworkBridge } from "./network-bridge";
import {
  FabricConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";

export class FabricBridge implements NetworkBridge {
  network: string = "FABRIC";

  connector: PluginLedgerConnectorFabric;
  bungee: PluginBungeeHermes;
  options: IPluginLedgerConnectorFabricOptions;
  config: FabricConfig;

  constructor(fabricConfig: FabricConfig) {
    this.config = fabricConfig;
    this.options = fabricConfig.options;
    this.connector = new PluginLedgerConnectorFabric(fabricConfig.options);
    this.bungee = new PluginBungeeHermes(fabricConfig.bungeeOptions);
    this.bungee.addStrategy(this.network, new StrategyFabric("INFO"));
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
      transactionId: response.transactionId,
      output: response.functionOutput,
    };
  }

  public async getReceipt(transactionId: string): Promise<string> {
    const networkDetails = {
      connectorApiPath: "", //todo check this to not use api
      signingCredential: this.config.signingCredential,
      contractName: this.config.contractName,
      channelName: this.config.channelName,
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
