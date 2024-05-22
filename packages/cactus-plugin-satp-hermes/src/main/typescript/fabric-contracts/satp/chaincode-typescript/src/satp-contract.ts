/*
 * SPDX-License-Identifier: Apache-2.0
 */

import { Context, Info, Returns, Transaction } from "fabric-contract-api";
import { ITraceableContract } from "@hyperledger/cactus-plugin-bungee-hermes/src/test/typescript/fabric-contracts/simple-asset/chaincode-typescript/src/ITraceableContract";

import TokenERC20Contract from "../lib/tokenERC20";

@Info({
  title: "SATPContract",
  description: "SATP ERC20 contract for trading asset",
})
export class SATPContract
  extends TokenERC20Contract
  implements ITraceableContract
{
  // GetAllAssetsKey returns all assets key found in the world state.

  constructor() {
    super("SATP");
  }

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

  @Transaction()
  public async mint(ctx: Context, amount: number): Promise<boolean> {
    await super.mint(ctx, amount);
    return true;
  }

  @Transaction()
  public async burn(ctx: Context, amount: number): Promise<boolean> {
    await super.burn(ctx, amount);
    return true;
  }

  @Transaction()
  public async assign(
    ctx: Context,
    from: string,
    to: string,
    amount: number,
  ): Promise<boolean> {
    await super._transfer(ctx, from, to, amount);
    return true;
  }

  @Transaction()
  public async transfer(
    ctx: Context,
    from: string,
    to: string,
    amount: number,
  ): Promise<boolean> {
    await super.TransferFrom(ctx, from, to, amount);
    return true;
  }
}
