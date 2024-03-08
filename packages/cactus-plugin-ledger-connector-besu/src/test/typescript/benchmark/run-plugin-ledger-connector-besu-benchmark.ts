import path from "path";
import { EOL } from "os";
import * as Benchmark from "benchmark";

import test, { Test } from "tape-promise/tape";
import { v4 as uuidv4 } from "uuid";
import { Server as SocketIoServer } from "socket.io";
import { PluginRegistry } from "@hyperledger/cactus-core";
import fse from "fs-extra";
import {
  PluginLedgerConnectorBesu,
  BesuApiClient,
  IPluginLedgerConnectorBesuOptions,
} from "../../../main/typescript/public-api";
import { PluginKeychainMemory } from "@hyperledger/cactus-plugin-keychain-memory";
import {
  BesuTestLedger,
  pruneDockerAllIfGithubAction,
} from "@hyperledger/cactus-test-tooling";
import KeyEncoder from "key-encoder";
import {
  IListenOptions,
  KeyFormat,
  LogLevelDesc,
  Secp256k1Keys,
  Servers,
} from "@hyperledger/cactus-common";
import { Constants } from "@hyperledger/cactus-core-api";
import express from "express";
import bodyParser from "body-parser";
import http from "http";
import HelloWorldContractJson from "../../solidity/hello-world-contract/HelloWorld.json";
import { AddressInfo } from "net";
import { BesuApiClientOptions } from "../../../main/typescript/api-client/besu-api-client";

import { installOpenapiValidationMiddleware } from "@hyperledger/cactus-core";
import OAS from "../../../main/json/openapi.json";

const LOG_TAG =
  "[packages/cactus-plugin-ledger-connector-besu/src/test/typescript/benchmark/run-plugin-ledger-connector-besu-benchmark.ts]";

const logLevel: LogLevelDesc = "TRACE";
const testCase = "Besu API";

test("BEFORE " + testCase, async (t: Test) => {
  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning did not throw OK");
  t.end();
});

const createTestInfrastructure = async (t: {
  readonly logLevel: LogLevelDesc;
}) => {
  const keyEncoder: KeyEncoder = new KeyEncoder("secp256k1");
  const keychainIdForSigned = uuidv4();
  const keychainIdForUnsigned = uuidv4();
  const keychainRefForSigned = uuidv4();
  const keychainRefForUnsigned = uuidv4();

  const besuTestLedger = new BesuTestLedger();
  await besuTestLedger.start();
  test.onFinish(async () => {
    await besuTestLedger.stop();
    await besuTestLedger.destroy();
  });

  const rpcApiHttpHost = await besuTestLedger.getRpcApiHttpHost();
  const rpcApiWsHost = await besuTestLedger.getRpcApiWsHost();

  const testEthAccount1 = await besuTestLedger.createEthTestAccount();
  // const testEthAccount2 = await besuTestLedger.createEthTestAccount();
  // const firstHighNetWorthAccount = besuTestLedger.getGenesisAccountPubKey();

  // keychainPlugin for signed transactions
  const { privateKey } = Secp256k1Keys.generateKeyPairsBuffer();
  const keyHex = privateKey.toString("hex");
  const pem = keyEncoder.encodePrivate(keyHex, KeyFormat.Raw, KeyFormat.PEM);
  const signedKeychainPlugin = new PluginKeychainMemory({
    instanceId: uuidv4(),
    keychainId: keychainIdForSigned,
    backend: new Map([[keychainRefForSigned, pem]]),
    logLevel,
  });

  // keychainPlugin for unsigned transactions
  const keychainEntryValue = testEthAccount1.privateKey;
  const unsignedKeychainPlugin = new PluginKeychainMemory({
    instanceId: uuidv4(),
    keychainId: keychainIdForUnsigned,
    backend: new Map([[keychainRefForUnsigned, keychainEntryValue]]),
    logLevel,
  });
  unsignedKeychainPlugin.set(
    HelloWorldContractJson.contractName,
    JSON.stringify(HelloWorldContractJson),
  );

  const pluginRegistry = new PluginRegistry({
    plugins: [signedKeychainPlugin, unsignedKeychainPlugin],
  });

  const options: IPluginLedgerConnectorBesuOptions = {
    instanceId: uuidv4(),
    rpcApiHttpHost,
    rpcApiWsHost,
    pluginRegistry,
    logLevel,
  };
  const connector = new PluginLedgerConnectorBesu(options);
  pluginRegistry.add(connector);

  const expressApp = express();
  expressApp.use(bodyParser.json({ limit: "250mb" }));
  const server = http.createServer(expressApp);

  const wsApi = new SocketIoServer(server, {
    path: Constants.SocketIoConnectionPathV1,
  });

  const listenOptions: IListenOptions = {
    hostname: "127.0.0.1",
    port: 0,
    server,
  };
  const addressInfo = (await Servers.listen(listenOptions)) as AddressInfo;
  test.onFinish(async () => await Servers.shutdown(server));
  const { address, port } = addressInfo;
  const apiHost = `http://${address}:${port}`;
  // t.comment(
  //   `Metrics URL: ${apiHost}/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-besu/get-prometheus-exporter-metrics`,
  // );

  // const wsBasePath = apiHost + Constants.SocketIoConnectionPathV1;
  // t.comment("WS base path: " + wsBasePath);
  const besuApiClientOptions = new BesuApiClientOptions({ basePath: apiHost });
  const apiClient = new BesuApiClient(besuApiClientOptions);

  await installOpenapiValidationMiddleware({
    logLevel,
    app: expressApp,
    apiSpec: OAS,
  });

  await connector.getOrCreateWebServices();
  await connector.registerWebServices(expressApp, wsApi);

  return {
    httpApi: apiClient,
    apiServer: connector,
  };
};

const main = async (opts: { readonly argv: Readonly<Array<string>> }) => {
  const logLevel: LogLevelDesc = "WARN";

  const gitRootPath = path.join(
    __dirname,
    "../../../../../../", // walk back up to the project root
  );

  console.log("%s gitRootPath=%s", LOG_TAG, gitRootPath);

  const DEFAULT_OUTPUT_FILE_RELATIVE_PATH =
    ".tmp/benchmark-results/plugin-ledger-connector-besu/run-plugin-ledger-connector-besu-benchmark.ts.log";

  const relativeOutputFilePath =
    opts.argv[2] === undefined
      ? DEFAULT_OUTPUT_FILE_RELATIVE_PATH
      : opts.argv[2];

  console.log(
    "%s DEFAULT_OUTPUT_FILE_RELATIVE_PATH=%s",
    LOG_TAG,
    DEFAULT_OUTPUT_FILE_RELATIVE_PATH,
  );

  console.log("%s opts.argv[2]=%s", LOG_TAG, opts.argv[2]);

  console.log("%s relativeOutputFilePath=%s", LOG_TAG, relativeOutputFilePath);

  const absoluteOutputFilePath = path.join(gitRootPath, relativeOutputFilePath);

  console.log("%s absoluteOutputFilePath=%s", LOG_TAG, absoluteOutputFilePath);

  const absoluteOutputDirPath = path.dirname(absoluteOutputFilePath);
  console.log("%s absoluteOutputDirPath=%s", LOG_TAG, absoluteOutputDirPath);

  await fse.mkdirp(absoluteOutputDirPath);
  console.log("%s mkdir -p OK: %s", LOG_TAG, absoluteOutputDirPath);

  const { apiServer, httpApi } = await createTestInfrastructure({ logLevel });

  const minSamples = 100;
  const suite = new Benchmark.Suite({});

  const cycles: string[] = [];

  await new Promise((resolve, reject) => {
    suite
      .add("plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1", {
        defer: true,
        minSamples,
        fn: async function (deferred: Benchmark.Deferred) {
          await httpApi.getOpenApiSpecV1();
          deferred.resolve();
        },
      })
      .on("cycle", (event: { target: unknown }) => {
        // Output benchmark result by converting benchmark result to string
        // Example line on stdout:
        // plugin-ledger-connector-besu_HTTP_GET_getOpenApiSpecV1 x 1,020 ops/sec ±2.25% (177 runs sampled)
        const cycle = String(event.target);
        console.log("%s Benchmark.js CYCLE: %s", LOG_TAG, cycle);
        cycles.push(cycle);
      })
      .on("complete", function () {
        console.log("%s Benchmark.js COMPLETE.", LOG_TAG);
        resolve(suite);
      })
      .on("error", (ex: unknown) => {
        console.log("%s Benchmark.js ERROR: %o", LOG_TAG, ex);
        reject(ex);
      })
      .run();
  });

  const data = cycles.join(EOL);
  console.log("%s Writing results...", LOG_TAG);
  await fse.writeFile(absoluteOutputFilePath, data, { encoding: "utf-8" });
  console.log("%s Wrote results to %s", LOG_TAG, absoluteOutputFilePath);

  await apiServer.shutdown();
  console.log("%s Shut down API server OK", LOG_TAG);
};

main({ argv: process.argv }).catch((ex: unknown) => {
  console.error("%s process crashed with:", LOG_TAG, ex);
  process.exit(1);
});
