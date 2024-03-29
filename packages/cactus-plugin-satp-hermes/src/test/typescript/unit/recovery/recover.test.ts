import "jest-extended";
import { Secp256k1Keys } from "@hyperledger/cactus-common";
import { v4 as uuidV4 } from "uuid";
import { PluginSatpGateway } from "../../../../main/typescript/gateway/plugin-satp-gateway";

import { RecoverV1Message } from "../../../../main/typescript/public-api";
import { randomInt } from "crypto";
import { checkValidRecoverMessage } from "../../../../main/typescript/gateway/recovery/recover";

import { BesuSatpGateway } from "../../../../main/typescript/gateway/besu-satp-gateway";
import { FabricSatpGateway } from "../../../../main/typescript/gateway/fabric-satp-gateway";
import { ClientGatewayHelper } from "../../../../main/typescript/gateway/client/client-helper";
import { ServerGatewayHelper } from "../../../../main/typescript/gateway/server/server-helper";

import {
  knexClientConnection,
  knexRemoteConnection,
  knexServerConnection,
} from "../../knex.config";

let pluginSourceGateway: PluginSatpGateway;
let pluginRecipientGateway: PluginSatpGateway;
let sessionID: string;

let sequenceNumber: number;

beforeEach(async () => {
  sessionID = uuidV4();
  sequenceNumber = randomInt(100);

  const sourceGatewayConstructor = {
    name: "plugin-satp-gateway#sourceGateway",
    dltIDs: ["DLT2"],
    instanceId: uuidV4(),
    keyPair: Secp256k1Keys.generateKeyPairsBuffer(),
    clientHelper: new ClientGatewayHelper(),
    serverHelper: new ServerGatewayHelper(),
    knexLocalConfig: knexClientConnection,
    knexRemoteConfig: knexRemoteConnection,
  };
  const recipientGatewayConstructor = {
    name: "plugin-satp-gateway#recipientGateway",
    dltIDs: ["DLT1"],
    instanceId: uuidV4(),
    keyPair: Secp256k1Keys.generateKeyPairsBuffer(),
    clientHelper: new ClientGatewayHelper(),
    serverHelper: new ServerGatewayHelper(),
    knexLocalConfig: knexServerConnection,
    knexRemoteConfig: knexRemoteConnection,
  };

  pluginSourceGateway = new FabricSatpGateway(sourceGatewayConstructor);
  pluginRecipientGateway = new BesuSatpGateway(recipientGatewayConstructor);

  const sessionData = {
    lastSequenceNumber: sequenceNumber,
    sourceGatewayPubkey: pluginSourceGateway.pubKey,
    recipientGatewayPubkey: pluginRecipientGateway.pubKey,
    step: 0,
    maxTimeout: 0,
    maxRetries: 0,
    rollbackProofs: [],
    sourceBasePath: "",
    recipientBasePath: "",
    rollbackActionsPerformed: [],
  };

  pluginSourceGateway.sessions.set(sessionID, sessionData);
  pluginRecipientGateway.sessions.set(sessionID, sessionData);

  if (
    pluginSourceGateway.localRepository?.database == undefined ||
    pluginRecipientGateway.localRepository?.database == undefined
  ) {
    throw new Error("Database is not correctly initialized");
  }

  await pluginSourceGateway.localRepository?.reset();
  await pluginRecipientGateway.localRepository?.reset();
});

test("valid recover message request from client", async () => {
  const recoverMessage: RecoverV1Message = {
    sessionID: sessionID,
    odapPhase: "1",
    sequenceNumber: sequenceNumber,
    lastLogEntryTimestamp: "sometimestamp",
    signature: "",
    isBackup: false,
    newBasePath: "",
  };

  recoverMessage.signature = PluginSatpGateway.bufArray2HexStr(
    pluginSourceGateway.sign(JSON.stringify(recoverMessage)),
  );

  await checkValidRecoverMessage(recoverMessage, pluginRecipientGateway);
});

test("valid recover message request from server", async () => {
  const recoverMessage: RecoverV1Message = {
    sessionID: sessionID,
    odapPhase: "1",
    sequenceNumber: sequenceNumber,
    lastLogEntryTimestamp: "sometimestamp",
    signature: "",
    isBackup: false,
    newBasePath: "",
  };

  recoverMessage.signature = PluginSatpGateway.bufArray2HexStr(
    pluginRecipientGateway.sign(JSON.stringify(recoverMessage)),
  );

  await checkValidRecoverMessage(recoverMessage, pluginSourceGateway).catch(
    () => {
      throw new Error("Test failed");
    },
  );
});

test("recover message request from client with wrong signature", async () => {
  const recoverMessage: RecoverV1Message = {
    sessionID: sessionID,
    odapPhase: "1",
    sequenceNumber: sequenceNumber,
    lastLogEntryTimestamp: "sometimestamp",
    signature: "",
    isBackup: false,
    newBasePath: "",
  };

  recoverMessage.signature = PluginSatpGateway.bufArray2HexStr(
    pluginRecipientGateway.sign(JSON.stringify("wrongRecoverMessage")),
  );

  await checkValidRecoverMessage(recoverMessage, pluginSourceGateway)
    .then(() => {
      throw new Error("Test Failed");
    })
    .catch((ex: Error) =>
      expect(ex.message).toMatch("message signature verification failed"),
    );
});

afterEach(() => {
  pluginSourceGateway.localRepository?.destroy();
  pluginRecipientGateway.localRepository?.destroy();
  pluginSourceGateway.remoteRepository?.destroy();
  pluginRecipientGateway.remoteRepository?.destroy();
});
