import {
  DefaultApi as FabricApi,
  FabricContractInvocationType,
} from "@hyperledger/cactus-plugin-ledger-connector-fabric";
import { PluginBungeeHermes } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/plugin-bungee-hermes";
import { StrategyFabric } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/strategy/strategy-fabric";
import { NetworkBridge } from "./network-bridge";
import {
  FabricConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";
import { PrivacyPolicyOpts } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/generated/openapi/typescript-axios";

export class FabricBridge implements NetworkBridge {
  network: string = "FABRIC";

  // connector: PluginLedgerConnectorFabric;
  bungee: PluginBungeeHermes;
  // options: IPluginLedgerConnectorFabricOptions;
  config: FabricConfig;
  api: FabricApi;

  constructor(fabricConfig: FabricConfig) {
    this.config = fabricConfig;
    // this.options = fabricConfig.options;
    // this.connector = new PluginLedgerConnectorFabric(fabricConfig.options);
    this.api = fabricConfig.api;
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
    const response = await this.api.runTransactionV1({
      signingCredential: this.config.signingCredential,
      channelName: this.config.channelName,
      contractName: this.config.contractName,
      invocationType: FabricContractInvocationType.Send,
      methodName: methodName,
      params: params,
    });
    // const response = await this.connector.transact({
    //   signingCredential: this.config.signingCredential,
    //   channelName: this.config.channelName,
    //   methodName: methodName,
    //   params: params,
    //   contractName: this.config.contractName,
    //   invocationType: FabricContractInvocationType.Send,
    // });

    if (response.status != 200) {
      throw new Error(`Error invoking contract: ${response.status}`);
    }
    return {
      transactionId: response.data.transactionId,
      output: response.data.functionOutput,
    };
  }

  public async getReceipt(
    assetId: string,
    transactionId: string,
  ): Promise<string> {
    const networkDetails = {
      connectorApiPath: this.config.apiPath, //todo check this to not use api
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

    const generated = this.bungee.generateView(
      snapshot,
      "0",
      Number.MAX_SAFE_INTEGER.toString(),
      undefined,
    );

    if (generated.view == undefined) {
      throw new Error("View is undefined");
    }

    const view = await this.bungee.processView(
      generated.view,
      PrivacyPolicyOpts.SingleTransaction,
      [assetId, transactionId],
    );

    return view.getViewStr();
  }
}
