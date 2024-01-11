/*
 * Hyperledger Cactus Plugin - Connector Besu
 * Can perform basic tasks on a Besu ledger
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
 * Gets or Sets EthContractInvocationType
 */
@JsonAdapter(EthContractInvocationType.Adapter.class)
public enum EthContractInvocationType {
  
  SEND("SEND"),
  
  CALL("CALL");

  private String value;

  EthContractInvocationType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static EthContractInvocationType fromValue(String value) {
    for (EthContractInvocationType b : EthContractInvocationType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<EthContractInvocationType> {
    @Override
    public void write(final JsonWriter jsonWriter, final EthContractInvocationType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public EthContractInvocationType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return EthContractInvocationType.fromValue(value);
    }
  }
}

