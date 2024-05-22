// /*
//  * Copyright IBM Corp. All Rights Reserved.
//  *
//  * SPDX-License-Identifier: Apache-2.0
//  */

// import { Context } from "fabric-contract-api";
// import { ChaincodeStub, ClientIdentity } from "fabric-shim";
// import { SATPContract } from "./../src/satp-contract";

// import * as winston from "winston";
// import * as chai from "chai";
// import * as chaiAsPromised from "chai-as-promised";
// import * as sinon from "sinon";
// import * as sinonChai from "sinon-chai";

// const bridgedOutAmountKey = "amountBridgedOut";

// const USER_A_FABRIC_ID =
//   "x509::/OU=org1/OU=client/OU=department1/CN=userA::/C=US/ST=North Carolina/L=Durham/O=org1.example.com/CN=ca.org1.example.com";

// const USER_A_ETH_ADDRESS =
//   "x509::/OU=org1/OU=client/OU=department1/CN=userA::/C=US/ST=North Carolina/L=Durham/O=org1.example.com/CN=ca.org1.example.com";

// chai.should();
// chai.use(chaiAsPromised);
// chai.use(sinonChai);

// class TestContext extends Context {
//   public stub: sinon.SinonStubbedInstance<ChaincodeStub> =
//     sinon.createStubInstance(ChaincodeStub);
//   public clientIdentity: sinon.SinonStubbedInstance<ClientIdentity> =
//     sinon.createStubInstance(ClientIdentity);
//   public logging = {
//     getLogger: sinon
//       .stub()
//       .returns(sinon.createStubInstance(winston.createLogger().constructor)),
//     setLevel: sinon.stub(),
//   };
// }

// describe("SATPContract", () => {
//   let contract: SATPContract;
//   let ctx: TestContext;

//   beforeEach(() => {
//     contract = new SATPContract();
//     ctx = new TestContext();
//     ctx.clientIdentity.getMSPID.resolves("Org2MSP");
//     ctx.stub.getState.withArgs(bridgedOutAmountKey).resolves(Buffer.from("50"));
//   });

//   describe("#mint", () => {
//     it("should mint asset", async () => {
//       ctx.clientIdentity.getMSPID.returns("Org1MSP");
//       ctx.clientIdentity.getID.returns("Alice");
//       ctx.stub.createCompositeKey.returns("balance_Alice");
//       const response = await contract.mint(ctx, 100);
//       sinon.assert.calledWith(
//         ctx.stub.putState.getCall(0),
//         "balance_Alice",
//         Buffer.from("100"),
//       );
//       ctx.stub.putState.should.have.been.calledWith(
//         bridgedOutAmountKey,
//         Buffer.from("100"),
//       );
//       chai.expect(response).to.equals(true);
//     });
//   });
// });
