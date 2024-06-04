import {
  BesuConfig,
  TransactionResponse,
} from "../../../types/blockchain-interaction";
import {
  DefaultApi as BesuApi,
  EthContractInvocationType,
} from "@hyperledger/cactus-plugin-ledger-connector-besu";
import { NetworkBridge } from "./network-bridge";
import { PluginBungeeHermes } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/plugin-bungee-hermes";
import { StrategyBesu } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/strategy/strategy-besu";
import { PrivacyPolicyOpts } from "@hyperledger/cactus-plugin-bungee-hermes/src/main/typescript/generated/openapi/typescript-axios";

export class BesuBridge implements NetworkBridge {
  network: string = "BESU";

  // connector: PluginLedgerConnectorBesu;
  api: BesuApi;
  bungee: PluginBungeeHermes;
  // options: IPluginLedgerConnectorBesuOptions;
  config: BesuConfig;

  constructor(besuConfig: BesuConfig) {
    this.config = besuConfig;
    // this.options = besuConfig.options;
    // this.connector = new PluginLedgerConnectorBesu(besuConfig.options);
    this.api = besuConfig.api;
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
    const response = await this.api.invokeContractV1({
      contractName: this.config.contractName,
      invocationType: EthContractInvocationType.Send,
      methodName: methodName,
      params: params,
      signingCredential: this.config.signingCredential,
    });
    // const response = await this.connector.invokeContract({
    //   contractName: this.config.contractName,
    //   invocationType: EthContractInvocationType.Send,
    //   methodName: methodName,
    //   params: params,
    //   signingCredential: this.config.signingCredential,
    // });

    if (response.status != 200) {
      throw new Error(`Error invoking contract: ${response.status}`);
    }

    return {
      transactionId: response.data.transactionReceipt?.transactionHash ?? "",
      output: response.data.callOutput.toString(),
    };
  }

  public async getReceipt(
    assetId: string,
    transactionHash: string,
  ): Promise<string> {
    const networkDetails = {
      connectorApiPath: this.config.apiPath,
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
      [assetId, transactionHash],
    );

    return view.getViewStr();
  }
}
