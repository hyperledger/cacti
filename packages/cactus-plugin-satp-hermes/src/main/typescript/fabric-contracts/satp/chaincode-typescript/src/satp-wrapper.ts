/*
 * SPDX-License-Identifier: Apache-2.0
 */

import {
  Context,
  Info,
  Returns,
  Transaction,
  Contract,
} from "fabric-contract-api";
import { ITraceableContract } from "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/fabric-contracts/simple-asset/chaincode-typescript/src/ITraceableContract";

enum TokenType {
  ERC20 = "ERC20",
  ERC721 = "ERC721",
  ERC1155 = "ERC1155",
}

interface Token {
  address: string;
  tokenType: TokenType;
  tokenId: string;
  owner: string;
  channelName: string;
  contractName: string;
  amount: number;
}

@Info({
  title: "SATPContractWrapper",
  description: "SATP Wrapper contract for trading assets",
})
export class SATPContractWrapper
  extends Contract
  implements ITraceableContract
{
  public bridge_address: string;

  constructor(_bridge_address: string) {
    super();
    this.bridge_address = _bridge_address;
  }

  public async getToken(ctx: Context, assetContract: string): Promise<Token> {
    const valueBytes = await ctx.stub.getState(assetContract);

    if (!valueBytes) {
      throw new Error("Asset not available");
    }

    return JSON.parse(valueBytes.toString());
  }

  @Transaction()
  public async wrap(
    ctx: Context,
    token: Token,
    assetContract: string,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const valueBytes = await ctx.stub.getState(assetContract);
    if (valueBytes) {
      throw new Error("Asset already wrapped");
    }

    await ctx.stub.putState(assetContract, Buffer.from(JSON.stringify(token)));
  }

  @Transaction()
  public async wrap2(
    ctx: Context,
    assetContract: string,
    tokenType: TokenType,
    tokenId: string,
    owner: string,
    channelName: string,
    contractName: string,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const valueBytes = await ctx.stub.getState(assetContract);
    if (valueBytes) {
      throw new Error("Asset already wrapped");
    }

    const token: Token = {
      address: assetContract,
      tokenType: tokenType,
      tokenId: tokenId,
      owner: owner,
      channelName: channelName,
      contractName: contractName,
      amount: 0,
    };

    await ctx.stub.putState(assetContract, Buffer.from(JSON.stringify(token)));
  }

  @Transaction()
  public async unwrap(ctx: Context, assetContract: string): Promise<void> {
    await ctx.stub.deleteState(assetContract);
  }

  @Transaction()
  public async lock(
    ctx: Context,
    assetContract: string,
    amount: number,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const token = await this.getToken(ctx, assetContract);

    const lockResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["transfer", token.owner, ctx.clientIdentity.getID(), amount.toString()],
      token.channelName,
    );

    if (lockResponse.status !== 200) {
      throw new Error("Lock failed");
    }

    token.amount += amount;

    await ctx.stub.putState(assetContract, Buffer.from(JSON.stringify(token)));
  }

  @Transaction()
  public async unlock(
    ctx: Context,
    assetContract: string,
    amount: number,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const token = await this.getToken(ctx, assetContract);

    if (token.amount < amount) {
      throw new Error("No sufficient amount locked");
    }

    const approveResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["approve", ctx.clientIdentity.getID(), amount.toString()],
      token.channelName,
    );

    if (approveResponse.status !== 200) {
      throw new Error("Approve failed");
    }

    const unlockResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["transfer", token.owner, ctx.clientIdentity.getID(), amount.toString()],
      token.channelName,
    );

    if (unlockResponse.status !== 200) {
      throw new Error("Unlock failed");
    }

    token.amount -= amount;

    await ctx.stub.putState(assetContract, Buffer.from(JSON.stringify(token)));
  }

  @Transaction()
  public async mint(
    ctx: Context,
    assetContract: string,
    amount: number,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const token = await this.getToken(ctx, assetContract);

    const mintResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["mint", amount.toString()],
      token.channelName,
    );

    if (mintResponse.status !== 200) {
      throw new Error("Mint failed");
    }

    token.amount += amount;
  }

  @Transaction()
  public async burn(
    ctx: Context,
    assetContract: string,
    amount: number,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const token = await this.getToken(ctx, assetContract);

    if (token.amount < amount) {
      throw new Error("No sufficient amount locked");
    }

    const burnResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["burn", amount.toString()],
      token.channelName,
    );

    if (burnResponse.status !== 200) {
      throw new Error("Burn failed");
    }

    token.amount -= amount;
  }

  @Transaction()
  public async assign(
    ctx: Context,
    assetContract: string,
    to: string,
    amount: number,
  ): Promise<void> {
    await this.CheckPermission(ctx);

    const token = await this.getToken(ctx, assetContract);

    if (token.amount < amount) {
      throw new Error("No sufficient amount locked");
    }

    const assignResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["assign", ctx.clientIdentity.getID(), to, amount.toString()],
      token.channelName,
    );

    if (assignResponse.status !== 200) {
      throw new Error("Assign failed");
    }

    token.amount -= amount;
  }

  // IsLocked returns true when asset with given ID is locked in world state.
  @Transaction(false)
  @Returns("number")
  public async IsLocked(ctx: Context, assetContract: string): Promise<number> {
    const token = await this.getToken(ctx, assetContract);

    return token.amount;
  }

  // GetAllAssetsKey returns all assets key found in the world state.
  @Transaction(false)
  @Returns("string")
  public async GetAllAssetsKey(ctx: Context): Promise<string> {
    const allResults = [];
    // range query with empty string for startKey and endKey does an open-ended query of all assets in the chaincode namespace.
    const iterator = await ctx.stub.getStateByRange("", "");
    let result = await iterator.next();
    while (!result.done) {
      allResults.push(result.value.key);
      result = await iterator.next();
    }
    return allResults.toString();
  }

  // GetAllTxByKey returns all transations for a specific key.
  @Transaction(false)
  @Returns("string")
  public async GetAllTxByKey(ctx: Context, key: string): Promise<string> {
    const allResults = [];
    const iterator = await ctx.stub.getHistoryForKey(key);
    let result = await iterator.next();
    while (!result.done) {
      const strValue = JSON.stringify(result);
      let record;
      try {
        record = JSON.parse(strValue);
      } catch (err) {
        console.log(err);
        record = strValue;
      }
      allResults.push(record);
      result = await iterator.next();
    }
    return JSON.stringify(allResults);
  }

  private async CheckPermission(ctx: Context) {
    // this needs to be called by entity2 (the bridging entity)
    const clientMSPID = await ctx.clientIdentity.getMSPID();
    if (clientMSPID !== "Org2MSP") {
      throw new Error(
        `client is not authorized to perform the operation. ${clientMSPID} != "Org2MSP"`,
      );
    }
  }
}
