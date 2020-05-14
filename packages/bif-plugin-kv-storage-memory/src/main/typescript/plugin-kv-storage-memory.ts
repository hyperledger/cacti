import { IPluginKVStorage } from '@hyperledger-labs/bif-core-api';

export interface IPluginKVStorageOptions {
  backend: Map<string, any>;
}

export class PluginKVStorageMemory implements IPluginKVStorage {
  constructor(public readonly options: IPluginKVStorageOptions) {
    if (!options) {
      throw new Error(`PluginKVStorageMemory#ctor options falsy.`);
    }
    if (!options.backend) {
      options.backend = new Map();
    }
  }

  async get<T>(key: string): Promise<T> {
    return this.options.backend.get(key);
  }

  async has(key: string): Promise<boolean> {
    return this.options.backend.has(key);
  }

  async set<T>(key: string, value: T): Promise<void> {
    this.options.backend.set(key, value);
  }

  async delete<T>(key: string): Promise<void> {
    this.options.backend.delete(key);
  }
}
