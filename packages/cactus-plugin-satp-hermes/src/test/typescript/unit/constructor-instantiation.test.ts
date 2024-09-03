import "jest-extended";
import { LogLevelDesc } from "@hyperledger/cactus-common";
import {
  ISATPGatewayRunnerConstructorOptions,
  pruneDockerAllIfGithubAction,
  SATPGatewayRunner,
} from "@hyperledger/cactus-test-tooling";
import {
  DEFAULT_PORT_GATEWAY_CLIENT,
  DEFAULT_PORT_GATEWAY_SERVER,
} from "../../../main/typescript/core/constants";

const testCase = "Instantiate SATP Gateway Runner";
const logLevel: LogLevelDesc = "TRACE";

describe(testCase, () => {
  let gatewayRunner: SATPGatewayRunner;

  const gatewayRunnerOptions: ISATPGatewayRunnerConstructorOptions = {
    containerImageVersion: "2024-09-02-3117",
    containerImageName:
      "ghcr.io/brunoffmateus/cactus-plugin-satp-hermes-gateway",
    serverPort: DEFAULT_PORT_GATEWAY_SERVER,
    clientPort: DEFAULT_PORT_GATEWAY_CLIENT,
    logLevel,
    emitContainerLogs: true,
  };

  beforeAll(async () => {
    const pruning = pruneDockerAllIfGithubAction({ logLevel });
    await expect(pruning).toResolve();
  });

  afterAll(async () => {
    await gatewayRunner.stop();
    await gatewayRunner.destroy();
    await pruneDockerAllIfGithubAction({ logLevel });
  });

  test(testCase, async () => {
    gatewayRunner = new SATPGatewayRunner(gatewayRunnerOptions);
    await gatewayRunner.start();
    expect(gatewayRunner).toBeTruthy();
    expect(gatewayRunner.getContainer()).toBeTruthy();

    await expect(
      gatewayRunner.waitForHealthCheck(60000),
    ).resolves.not.toThrow();

    const serverHost = await gatewayRunner.getServerHost();
    expect(serverHost).toBeTruthy();
    expect(serverHost).toMatch(/^http:\/\/localhost:\d+$/);

    const clientHost = await gatewayRunner.getClientHost();
    expect(clientHost).toBeTruthy();
    expect(clientHost).toMatch(/^http:\/\/localhost:\d+$/);
  });
});
