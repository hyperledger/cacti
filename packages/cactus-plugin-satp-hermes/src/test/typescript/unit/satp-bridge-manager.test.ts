import "jest-extended";
import {
  Containers,
  pruneDockerAllIfGithubAction,
} from "@hyperledger/cactus-test-tooling";
import { LogLevelDesc, LoggerProvider } from "@hyperledger/cactus-common";

import { SATPBridgeConfig } from "../../../main/typescript/core/types";
import { SatpBridgeManager } from "../../../main/typescript/core/stage-services/satp-bridge/satp-bridge-manager";

const logLevel: LogLevelDesc = "DEBUG";
const logger = LoggerProvider.getOrCreate({
  level: logLevel,
  label: "satp-bridge-manager-test",
});

beforeAll(async () => {
  pruneDockerAllIfGithubAction({ logLevel })
    .then(() => {
      logger.info("Pruning throw OK");
    })
    .catch(async () => {
      await Containers.logDiagnostics({ logLevel });
      fail("Pruning didn't throw OK");
    });
});

describe("SATP bridge function testing", () => {
  test("Lock asset function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.lockAsset(assetId);
    expect(receipt).toBeDefined();
  });
  test("Unlock asset function  works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.unlockAsset(assetId);
    expect(receipt).toBeDefined();
  });
  test("Mint asset function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.mintAsset(assetId);
    expect(receipt).toBeDefined();
  });
  test("Burn asset function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.burnAsset(assetId);
    expect(receipt).toBeDefined();
  });
  test("Assign asset function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const recipient = "testRecipient";
    const receipt = await bridgeManager.assignAsset(assetId, recipient);
    expect(receipt).toBeDefined();
  });
  test("Verify asset existence function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.verifyAssetExistence(assetId);
    expect(receipt).toBeDefined();
  });
  test("Verify lock asset function works", async () => {
    const bridgeConfig: SATPBridgeConfig = {
      network: {
        runTransaction: jest.fn().mockReturnValue({
          transactionId: "testTransaction",
          output: "testOutput",
        }),
        getReceipt: jest.fn().mockReturnValue("testReceipt"),
        network: "testNetwork",
        networkName: function (): string {
          return this.network;
        },
      },
    };
    const bridgeManager = new SatpBridgeManager(bridgeConfig);
    const assetId = "testAssetId";
    const receipt = await bridgeManager.verifyLockAsset(assetId);
    expect(receipt).toBeDefined();
  });
});

afterAll(async () => {
  await pruneDockerAllIfGithubAction({ logLevel })
    .then(() => {
      logger.info("Pruning throw OK");
    })
    .catch(async () => {
      await Containers.logDiagnostics({ logLevel });
      fail("Pruning didn't throw OK");
    });
});
