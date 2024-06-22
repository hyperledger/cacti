import test, { Test } from "tape-promise/tape";
import { v4 as internalIpV4 } from "internal-ip";

import {
  Containers,
  CordaTestLedger,
  pruneDockerAllIfGithubAction,
} from "@hyperledger/cactus-test-tooling";
import { LogLevelDesc } from "@hyperledger/cactus-common";
import {
  SampleCordappEnum,
  CordaConnectorContainer,
} from "@hyperledger/cactus-test-tooling";

import {
  CordappDeploymentConfig,
  DefaultApi as CordaApi,
  DeployContractJarsV1Request,
  FlowInvocationType,
  InvokeContractV1Request,
  JvmTypeKind,
} from "../../../main/typescript/generated/openapi/typescript-axios/index";
import { Configuration } from "@hyperledger/cactus-core-api";
import {
  createJvmCordaAmount,
  createJvmCordaIdentityParty,
  createJvmCordaUniqueIdentifier,
} from "../../../main/typescript/public-api";

const testCase = "Tests are passing on the JVM side";
const logLevel: LogLevelDesc = "INFO";

test.onFailure(async () => {
  await Containers.logDiagnostics({ logLevel });
});

test("BEFORE " + testCase, async (t: Test) => {
  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning didn't throw OK");
  t.end();
});

test(testCase, async (t: Test) => {
  const ledger = new CordaTestLedger({
    imageName: "ghcr.io/hyperledger/cactus-corda-4-8-all-in-one-obligation",
    imageVersion: "2021-08-31--feat-889",
    logLevel,
  });
  t.ok(ledger, "CordaTestLedger v4.8 instantaited OK");

  test.onFinish(async () => {
    await ledger.stop();
    await ledger.destroy();
    await pruneDockerAllIfGithubAction({ logLevel });
  });
  const ledgerContainer = await ledger.start();
  t.ok(
    ledgerContainer,
    "CordaTestLedger v4.8 container truthy post-start() OK",
  );

  await ledger.logDebugPorts();
  const partyARpcPort = await ledger.getRpcAPublicPort();

  const jarFiles = await ledger.pullCordappJars(
    SampleCordappEnum.ADVANCED_OBLIGATION,
  );
  t.comment(`Fetched ${jarFiles.length} cordapp jars OK`);

  const internalIpOrUndefined = await internalIpV4();
  t.ok(internalIpOrUndefined, "Determined LAN IPv4 address successfully OK");
  const internalIp = internalIpOrUndefined as string;
  t.comment(`Internal IP (based on default gateway): ${internalIp}`);

  const springAppConfig = {
    logging: {
      level: {
        root: "INFO",
        "org.hyperledger.cactus": "DEBUG",
      },
    },
    cactus: {
      corda: {
        node: { host: internalIp },
        // TODO: parse the gradle build files to extract the credentials?
        rpc: { port: partyARpcPort, username: "user1", password: "password" },
      },
    },
  };
  const springApplicationJson = JSON.stringify(springAppConfig);
  const envVarSpringAppJson = `SPRING_APPLICATION_JSON=${springApplicationJson}`;
  t.comment(envVarSpringAppJson);

  const connector = new CordaConnectorContainer({
    logLevel,
    imageName: "ghcr.io/hyperledger/cactus-connector-corda-server",
    imageVersion: "2021-11-23--feat-1493",
    envVars: [envVarSpringAppJson],
  });
  t.ok(CordaConnectorContainer, "CordaConnectorContainer instantiated OK");

  test.onFinish(async () => {
    try {
      await connector.stop();
    } finally {
      await connector.destroy();
    }
  });

  const connectorContainer = await connector.start();
  t.ok(connectorContainer, "CordaConnectorContainer started OK");

  await connector.logDebugPorts();
  const apiUrl = await connector.getApiLocalhostUrl();
  const config = new Configuration({ basePath: apiUrl });
  const apiClient = new CordaApi(config);

  const flowsRes = await apiClient.listFlowsV1();
  t.ok(flowsRes.status === 200, "flowsRes.status === 200 OK");
  t.ok(flowsRes.data, "flowsRes.data truthy OK");
  t.ok(flowsRes.data.flowNames, "flowsRes.data.flowNames truthy OK");
  t.comment(`apiClient.listFlowsV1() => ${JSON.stringify(flowsRes.data)}`);

  const diagRes = await apiClient.diagnoseNodeV1();
  t.ok(diagRes.status === 200, "diagRes.status === 200 OK");
  t.ok(diagRes.data, "diagRes.data truthy OK");
  t.ok(diagRes.data.nodeDiagnosticInfo, "nodeDiagnosticInfo truthy OK");
  const ndi = diagRes.data.nodeDiagnosticInfo;
  t.ok(ndi.cordapps, "ndi.cordapps truthy OK");
  t.ok(Array.isArray(ndi.cordapps), "ndi.cordapps is Array truthy OK");
  t.true((ndi.cordapps as []).length > 0, "ndi.cordapps non-empty true OK");
  t.ok(ndi.vendor, "ndi.vendor truthy OK");
  t.ok(ndi.version, "ndi.version truthy OK");
  t.ok(ndi.revision, "ndi.revision truthy OK");
  t.ok(ndi.platformVersion, "ndi.platformVersion truthy OK");

  t.comment(`apiClient.diagnoseNodeV1() => ${JSON.stringify(diagRes.data)}`);

  const cordappDeploymentConfigs: CordappDeploymentConfig[] = [];
  const depReq: DeployContractJarsV1Request = {
    jarFiles,
    cordappDeploymentConfigs,
  };
  const depRes = await apiClient.deployContractJarsV1(depReq);
  t.ok(depRes, "Jar deployment response truthy OK");
  t.equal(depRes.status, 200, "Jar deployment status code === 200 OK");
  t.ok(depRes.data, "Jar deployment response body truthy OK");
  t.ok(depRes.data.deployedJarFiles, "Jar deployment body deployedJarFiles OK");
  t.equal(
    depRes.data.deployedJarFiles.length,
    jarFiles.length,
    "Deployed jar file count equals count in request OK",
  );

  const networkMapRes = await apiClient.networkMapV1();
  const partyA = networkMapRes.data.find((it) =>
    it.legalIdentities.some((it2) => it2.name.organisation === "ParticipantA"),
  );

  const partyB = networkMapRes.data.find((it) =>
    it.legalIdentities.some((it2) => it2.name.organisation === "ParticipantB"),
  );

  const req: InvokeContractV1Request = {
    flowFullClassName: "net.corda.samples.obligation.flows.IOUIssueFlow",
    flowInvocationType: FlowInvocationType.TrackedFlowDynamic,
    params: [
      {
        jvmTypeKind: JvmTypeKind.Reference,
        jvmType: {
          fqClassName: "net.corda.samples.obligation.states.IOUState",
        },

        jvmCtorArgs: [
          createJvmCordaAmount({ currencyCode: "USD", amount: 42 }),
          createJvmCordaIdentityParty({
            party: partyA!.legalIdentities[0],
          }),
          createJvmCordaIdentityParty({
            party: partyB!.legalIdentities[0],
          }),
          createJvmCordaAmount({ currencyCode: "USD", amount: 1 }),
          createJvmCordaUniqueIdentifier({
            uniqueidentifier: "7fc2161e-f8d0-4c86-a596-08326bdafd56",
          }),
        ],
      },
    ],
    timeoutMs: 60000,
  } as unknown as InvokeContractV1Request;

  const res = await apiClient.invokeContractV1(req);
  t.ok(res, "InvokeContractV1Request truthy OK");
  t.equal(res.status, 200, "InvokeContractV1Request status code === 200 OK");

  t.end();
});
