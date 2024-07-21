//
//Hyperledger Cactus Plugin - Keychain Memory 
//
//Contains/describes the Hyperledger Cacti Keychain Memory plugin.
//
//The version of the OpenAPI document: 2.0.0-rc.3
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file models/get_keychain_entry_request_v1_pb.proto (package org.hyperledger.cacti.plugin.keychain.memory, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message org.hyperledger.cacti.plugin.keychain.memory.GetKeychainEntryRequestV1PB
 */
export class GetKeychainEntryRequestV1PB extends Message<GetKeychainEntryRequestV1PB> {
  /**
   * The key for the entry to get from the keychain.
   *
   * @generated from field: string key = 106079;
   */
  key = "";

  constructor(data?: PartialMessage<GetKeychainEntryRequestV1PB>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "org.hyperledger.cacti.plugin.keychain.memory.GetKeychainEntryRequestV1PB";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 106079, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetKeychainEntryRequestV1PB {
    return new GetKeychainEntryRequestV1PB().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetKeychainEntryRequestV1PB {
    return new GetKeychainEntryRequestV1PB().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetKeychainEntryRequestV1PB {
    return new GetKeychainEntryRequestV1PB().fromJsonString(jsonString, options);
  }

  static equals(a: GetKeychainEntryRequestV1PB | PlainMessage<GetKeychainEntryRequestV1PB> | undefined, b: GetKeychainEntryRequestV1PB | PlainMessage<GetKeychainEntryRequestV1PB> | undefined): boolean {
    return proto3.util.equals(GetKeychainEntryRequestV1PB, a, b);
  }
}

