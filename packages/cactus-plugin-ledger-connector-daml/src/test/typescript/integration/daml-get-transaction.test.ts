import "jest-extended";
import { v4 as uuidv4 } from "uuid";
import { DamlTestLedger } from "@hyperledger/cactus-test-tooling";

describe("PluginLedgerConnectorDAML", () => {
  const damlTestLedger = new DamlTestLedger();
  let rpcApiHttpHost: string;

  beforeAll(async () => {
    await damlTestLedger.start();
    rpcApiHttpHost = await damlTestLedger.getDamlApiHost();
  });

  afterAll(async () => {
    await damlTestLedger.stop();
    await damlTestLedger.destroy();
  });

});
