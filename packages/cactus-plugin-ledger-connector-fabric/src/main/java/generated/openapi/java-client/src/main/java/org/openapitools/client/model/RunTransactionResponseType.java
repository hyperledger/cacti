/*
 * Hyperledger Cactus Plugin - Connector Fabric
 * Can perform basic tasks on a fabric ledger
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
 * Response format from transaction / query execution
 */
@JsonAdapter(RunTransactionResponseType.Adapter.class)
public enum RunTransactionResponseType {
  
  JSON("org.hyperledger.cacti.api.hlfabric.RunTransactionResponseType.JSON"),
  
  UTF8("org.hyperledger.cacti.api.hlfabric.RunTransactionResponseType.UTF8");

  private String value;

  RunTransactionResponseType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static RunTransactionResponseType fromValue(String value) {
    for (RunTransactionResponseType b : RunTransactionResponseType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<RunTransactionResponseType> {
    @Override
    public void write(final JsonWriter jsonWriter, final RunTransactionResponseType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public RunTransactionResponseType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return RunTransactionResponseType.fromValue(value);
    }
  }
}

