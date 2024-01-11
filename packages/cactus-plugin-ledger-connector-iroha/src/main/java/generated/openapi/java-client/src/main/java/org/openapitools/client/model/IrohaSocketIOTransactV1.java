/*
 * Hyperledger Cactus Plugin - Connector Iroha
 * Can perform basic tasks on a Iroha ledger
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
 * Gets or Sets IrohaSocketIOTransactV1
 */
@JsonAdapter(IrohaSocketIOTransactV1.Adapter.class)
public enum IrohaSocketIOTransactV1 {
  
  SendAsyncRequest("org.hyperledger.cactus.api.async.iroha.IrohaSocketIOTransactV1.SendAsyncRequest"),
  
  SendSyncRequest("org.hyperledger.cactus.api.async.iroha.IrohaSocketIOTransactV1.SendSyncRequest");

  private String value;

  IrohaSocketIOTransactV1(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static IrohaSocketIOTransactV1 fromValue(String value) {
    for (IrohaSocketIOTransactV1 b : IrohaSocketIOTransactV1.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<IrohaSocketIOTransactV1> {
    @Override
    public void write(final JsonWriter jsonWriter, final IrohaSocketIOTransactV1 enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public IrohaSocketIOTransactV1 read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return IrohaSocketIOTransactV1.fromValue(value);
    }
  }
}

