import test, { Test } from "tape-promise/tape";
import { v4 as uuidv4 } from "uuid";

import { LogLevelDesc } from "@hyperledger/cactus-common";

import {
  Configuration,
  PluginImportAction,
  PluginImportType,
} from "@hyperledger/cactus-core-api";

import {
  ApiServer,
  AuthorizationProtocol,
  ConfigService,
} from "../../../../main/typescript/public-api";

import { K_CACTUS_API_SERVER_TOTAL_PLUGIN_IMPORTS } from "../../../../main/typescript/prometheus-exporter/metrics";

import { DefaultApi as ApiServerApi } from "../../../../main/typescript/public-api";
import path from "path";

const logLevel: LogLevelDesc = "TRACE";

test("can import plugins at runtime (CLI)", async (t: Test) => {
  // const pluginsPath = path.join(
  //   "/tmp/org/hyperledger/cactus/cmd-api-server/install-basic-plugin-keychain-memory_test", // the dir path from the root
  //   uuidv4(), // then a random directory to ensure proper isolation
  // );
  const pluginsPath = path.join(
    __dirname, // start at the current file's path
    "../../../../../../../", // walk back up to the project root
    ".tmp/test/cmd-api-server/runtime-plugin-imports_test", // the dir path from the root
    uuidv4(), // then a random directory to ensure proper isolation
  );
  const pluginManagerOptionsJson = JSON.stringify({ pluginsPath });

  const configService = new ConfigService();
  const apiServerOptions = await configService.newExampleConfig();
  apiServerOptions.authorizationProtocol = AuthorizationProtocol.NONE;
  apiServerOptions.pluginManagerOptionsJson = pluginManagerOptionsJson;
  apiServerOptions.configFile = "";
  apiServerOptions.apiCorsDomainCsv = "*";
  apiServerOptions.apiPort = 0;
  apiServerOptions.cockpitPort = 0;
  apiServerOptions.grpcPort = 0;
  apiServerOptions.apiTlsEnabled = false;
  apiServerOptions.plugins = [
    {
      packageName: "@hyperledger/cactus-plugin-keychain-memory",
      type: PluginImportType.Local,
      action: PluginImportAction.Install,
      options: {
        instanceId: uuidv4(),
        keychainId: uuidv4(),
        logLevel,
      },
    },
  ];
  const config = await configService.newExampleConfigConvict(apiServerOptions);

  const apiServer = new ApiServer({
    config: config.getProperties(),
  });

  const startResponse = apiServer.start();
  await t.doesNotReject(startResponse, "started API server dynamic imports OK");
  t.ok(startResponse, "startResponse truthy OK");

  const addressInfoApi = (await startResponse).addressInfoApi;
  const protocol = apiServerOptions.apiTlsEnabled ? "https" : "http";
  const { address, port } = addressInfoApi;
  const apiHost = `${protocol}://${address}:${port}`;
  t.comment(
    `Metrics URL: ${apiHost}/api/v1/api-server/get-prometheus-exporter-metrics`,
  );

  const apiConfig = new Configuration({ basePath: apiHost });
  const apiClient = new ApiServerApi(apiConfig);

  {
    const res = await apiClient.getPrometheusMetricsV1();
    const promMetricsOutput =
      "# HELP " +
      K_CACTUS_API_SERVER_TOTAL_PLUGIN_IMPORTS +
      " Total number of plugins imported\n" +
      "# TYPE " +
      K_CACTUS_API_SERVER_TOTAL_PLUGIN_IMPORTS +
      " gauge\n" +
      K_CACTUS_API_SERVER_TOTAL_PLUGIN_IMPORTS +
      '{type="' +
      K_CACTUS_API_SERVER_TOTAL_PLUGIN_IMPORTS +
      '"} 1';
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
    t.true(
      res.data.includes(promMetricsOutput),
      "Total 1 plugins imported as expected. RESULT OK",
    );
  }

  test.onFinish(() => apiServer.shutdown());
});
