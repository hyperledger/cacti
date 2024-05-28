/*
  SPDX-License-Identifier: Apache-2.0
*/

import { Object, Property } from "fabric-contract-api";

export enum TokenType {
  ERC20 = "ERC20",
  ERC721 = "ERC721",
  ERC1155 = "ERC1155",
}

@Object()
export class Token {
  @Property()
  public address: string;

  @Property()
  public tokenType: TokenType;

  @Property()
  public tokenId: string;

  @Property()
  public owner: string;

  @Property()
  channelName: string;

  @Property()
  contractName: string;

  @Property()
  amount: number;
}
