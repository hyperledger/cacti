export * from "./generated/openapi/typescript-axios/index.js";

import { IPluginFactoryOptions } from "@hyperledger/cactus-core-api";

import { PluginFactoryKeychain } from "./plugin-factory-keychain.js";
export { PluginFactoryKeychain } from "./plugin-factory-keychain.js";
export {
  IPluginKeychainVaultOptions,
  PluginKeychainVault,
} from "./plugin-keychain-vault.js";

export async function createPluginFactory(
  pluginFactoryOptions: IPluginFactoryOptions,
): Promise<PluginFactoryKeychain> {
  return new PluginFactoryKeychain(pluginFactoryOptions);
}
