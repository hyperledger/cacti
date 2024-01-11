/*
 * Hyperledger Cactus Plugin - Connector Iroha V2
 * Can perform basic tasks on a Iroha V2 ledger
 *
 * The version of the OpenAPI document: v2.0.0-alpha.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * Command names that correspond to Iroha Special Instructions (https://hyperledger.github.io/iroha-2-docs/guide/advanced/isi.html)
 */
@JsonAdapter(IrohaInstruction.Adapter.class)
public enum IrohaInstruction {
  
  /**
   * Register new domain
   */
  RegisterDomain("registerDomain"),
  
  /**
   * Register new asset definition
   */
  RegisterAssetDefinition("registerAssetDefinition"),
  
  /**
   * Register new asset
   */
  RegisterAsset("registerAsset"),
  
  /**
   * Mint asset value
   */
  MintAsset("mintAsset"),
  
  /**
   * Burn asset value
   */
  BurnAsset("burnAsset"),
  
  /**
   * Transfer asset between accounts
   */
  TransferAsset("transferAsset"),
  
  /**
   * Register new account
   */
  RegisterAccount("registerAccount");

  private String value;

  IrohaInstruction(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static IrohaInstruction fromValue(String value) {
    for (IrohaInstruction b : IrohaInstruction.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<IrohaInstruction> {
    @Override
    public void write(final JsonWriter jsonWriter, final IrohaInstruction enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public IrohaInstruction read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return IrohaInstruction.fromValue(value);
    }
  }
}

