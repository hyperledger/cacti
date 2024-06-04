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

import { Token, TokenType } from "./token";

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

  @Transaction(false)
  public async Initialize(ctx: Context, owner: string): Promise<boolean> {
    await ctx.stub.putState("owner", Buffer.from(owner));
    return true;
  }

  @Transaction(false)
  public async setBridge(ctx: Context, bridge: string): Promise<boolean> {
    this.checkPermission(ctx);
    await ctx.stub.putState("bridge", Buffer.from(bridge));
    return true;
  }

  @Transaction()
  public async getToken(ctx: Context, tokenId: string): Promise<Token> {
    const valueBytes: Uint8Array = await ctx.stub.getState(tokenId);

    if (!valueBytes) {
      throw new Error(`Asset with ID ${tokenId} does not exist`);
    }

    return JSON.parse(valueBytes.toString()) as Token;
  }

  @Transaction()
  @Returns("boolean")
  public async wrap(
    ctx: Context,
    assetContractID: string,
    tokenType: TokenType,
    tokenId: string,
    owner: string,
    channelName: string,
    contractName: string,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const valueBytes = await ctx.stub.getState(tokenId);
    if (valueBytes) {
      throw new Error(`Asset with ID ${tokenId} is already wrapped`);
    }

    const token: Token = {
      address: assetContractID,
      tokenType: tokenType,
      tokenId: tokenId,
      owner: owner,
      channelName: channelName,
      contractName: contractName,
      amount: 0,
    };

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
  }

  @Transaction()
  @Returns("boolean")
  public async unwrap(ctx: Context, tokenId: string): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    if (token.amount > 0) {
      throw new Error("Token has locked amount");
    }

    await ctx.stub.deleteState(tokenId);
    return true;
  }

  @Transaction()
  @Returns("number")
  public async lockedAmount(ctx: Context, tokenId: string): Promise<number> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    return token.amount;
  }

  @Transaction()
  @Returns("boolean")
  public async lock(
    ctx: Context,
    tokenId: string,
    amount: number,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    const to = await ctx.clientIdentity.getMSPID();

    const lockResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["transfer", token.owner, to, amount.toString()],
      token.channelName,
    );

    if (lockResponse.status !== 200) {
      throw new Error("Lock failed");
    }

    token.amount += amount;

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
  }

  @Transaction()
  public async unlock(
    ctx: Context,
    tokenId: string,
    amount: number,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    if (token.amount < amount) {
      throw new Error(
        "No sufficient amount locked, total tried to unlock: " +
          amount +
          " total locked: " +
          token.amount,
      );
    }

    const spender = await ctx.clientIdentity.getMSPID();

    const approveResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["Approve", spender, amount.toString()],
      token.channelName,
    );

    if (approveResponse.status !== 200) {
      throw new Error("Approve failed");
    }

    const unlockResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["transfer", token.owner, spender, amount.toString()],
      token.channelName,
    );

    if (unlockResponse.status !== 200) {
      throw new Error("Unlock failed");
    }

    token.amount -= amount;

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
  }

  @Transaction()
  @Returns("boolean")
  public async mint(
    ctx: Context,
    tokenId: string,
    amount: number,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    const mintResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["mint", amount.toString()],
      token.channelName,
    );

    if (mintResponse.status !== 200) {
      throw new Error("Mint failed");
    }

    token.amount += amount;

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
  }

  @Transaction()
  @Returns("boolean")
  public async burn(
    ctx: Context,
    tokenId: string,
    amount: number,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

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

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
  }

  @Transaction()
  @Returns("boolean")
  public async assign(
    ctx: Context,
    tokenId: string,
    to: string,
    amount: number,
  ): Promise<boolean> {
    await this.checkPermission(ctx);

    const token = await this.getToken(ctx, tokenId);

    if (token.amount < amount) {
      throw new Error("No sufficient amount locked");
    }

    const from = await ctx.clientIdentity.getMSPID();

    const assignResponse = await ctx.stub.invokeChaincode(
      token.contractName,
      ["assign", from, to, amount.toString()],
      token.channelName,
    );

    if (assignResponse.status !== 200) {
      throw new Error("Assign failed");
    }

    token.amount -= amount;

    await ctx.stub.putState(tokenId, Buffer.from(JSON.stringify(token)));
    return true;
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

  // add two number checking for overflow
  add(a, b) {
    const c = a + b;
    if (a !== c - b || b !== c - a) {
      throw new Error(`Math: addition overflow occurred ${a} + ${b}`);
    }
    return c;
  }

  // add two number checking for overflow
  sub(a, b) {
    const c = a - b;
    if (a !== c + b || b !== a - c) {
      throw new Error(`Math: subtraction overflow occurred ${a} - ${b}`);
    }
    return c;
  }

  private async checkPermission(ctx: Context) {
    let owner = await ctx.stub.getState("owner");
    let bridge = await ctx.stub.getState("bridge");

    owner = owner ? owner : Buffer.from("");

    bridge = bridge ? bridge : Buffer.from("");

    const clientMSPID = await ctx.clientIdentity.getMSPID();
    if (
      !(clientMSPID == owner.toString() || clientMSPID == bridge.toString())
    ) {
      throw new Error(
        `client is not authorized to perform the operation. ${clientMSPID}`,
      );
    }
  }
}
