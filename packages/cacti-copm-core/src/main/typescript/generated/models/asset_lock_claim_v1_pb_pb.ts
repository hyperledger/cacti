//
//Hyperledger Cacti Plugin - Common Operator Primitives
//
//Contains/describes the Hyperledger Cacti Common Operator Primitives plugin.  These primitives require specific chaincode and weaver relays to be deployed on the network as described at https://hyperledger.github.io/cacti/weaver/introduction/.
//
//The version of the OpenAPI document: 2.0.0-rc.2
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file models/asset_lock_claim_v1_pb.proto (package org.hyperledger.cacti.plugin.cacti.plugin.copm.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AssetAccountV1PB } from "./asset_account_v1_pb_pb.js";
import { TransferrableAssetV1PB } from "./transferrable_asset_v1_pb_pb.js";
import { HashInfoV1PB } from "./hash_info_v1_pb_pb.js";

/**
 * @generated from message org.hyperledger.cacti.plugin.cacti.plugin.copm.core.AssetLockClaimV1PB
 */
export class AssetLockClaimV1PB extends Message<AssetLockClaimV1PB> {
  /**
   * @generated from field: string lock_id = 338696367;
   */
  lockId = "";

  /**
   * @generated from field: org.hyperledger.cacti.plugin.cacti.plugin.copm.core.AssetAccountV1PB source = 359634918;
   */
  source?: AssetAccountV1PB;

  /**
   * @generated from field: org.hyperledger.cacti.plugin.cacti.plugin.copm.core.AssetAccountV1PB destination = 356105204;
   */
  destination?: AssetAccountV1PB;

  /**
   * @generated from field: org.hyperledger.cacti.plugin.cacti.plugin.copm.core.TransferrableAssetV1PB asset = 93121264;
   */
  asset?: TransferrableAssetV1PB;

  /**
   * @generated from field: org.hyperledger.cacti.plugin.cacti.plugin.copm.core.HashInfoV1PB hash_info = 309318975;
   */
  hashInfo?: HashInfoV1PB;

  /**
   * @generated from field: string source_certificate = 72464244;
   */
  sourceCertificate = "";

  /**
   * @generated from field: string dest_certificate = 165133181;
   */
  destCertificate = "";

  constructor(data?: PartialMessage<AssetLockClaimV1PB>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "org.hyperledger.cacti.plugin.cacti.plugin.copm.core.AssetLockClaimV1PB";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 338696367, name: "lock_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 359634918, name: "source", kind: "message", T: AssetAccountV1PB },
    { no: 356105204, name: "destination", kind: "message", T: AssetAccountV1PB },
    { no: 93121264, name: "asset", kind: "message", T: TransferrableAssetV1PB },
    { no: 309318975, name: "hash_info", kind: "message", T: HashInfoV1PB },
    { no: 72464244, name: "source_certificate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 165133181, name: "dest_certificate", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssetLockClaimV1PB {
    return new AssetLockClaimV1PB().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssetLockClaimV1PB {
    return new AssetLockClaimV1PB().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssetLockClaimV1PB {
    return new AssetLockClaimV1PB().fromJsonString(jsonString, options);
  }

  static equals(a: AssetLockClaimV1PB | PlainMessage<AssetLockClaimV1PB> | undefined, b: AssetLockClaimV1PB | PlainMessage<AssetLockClaimV1PB> | undefined): boolean {
    return proto3.util.equals(AssetLockClaimV1PB, a, b);
  }
}

