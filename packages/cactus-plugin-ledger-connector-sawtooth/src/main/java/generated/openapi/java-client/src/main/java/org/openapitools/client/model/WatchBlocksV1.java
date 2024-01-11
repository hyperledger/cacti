/*
 * Hyperledger Cacti Plugin - Connector Sawtooth
 * Can perform basic tasks on a Sawtooth ledger
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
 * Gets or Sets WatchBlocksV1
 */
@JsonAdapter(WatchBlocksV1.Adapter.class)
public enum WatchBlocksV1 {
  
  Subscribe("org.hyperledger.cacti.api.async.sawtooth.WatchBlocksV1.Subscribe"),
  
  Next("org.hyperledger.cacti.api.async.sawtooth.WatchBlocksV1.Next"),
  
  Unsubscribe("org.hyperledger.cacti.api.async.sawtooth.WatchBlocksV1.Unsubscribe"),
  
  Error("org.hyperledger.cacti.api.async.sawtooth.WatchBlocksV1.Error"),
  
  Complete("org.hyperledger.cacti.api.async.sawtooth.WatchBlocksV1.Complete");

  private String value;

  WatchBlocksV1(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static WatchBlocksV1 fromValue(String value) {
    for (WatchBlocksV1 b : WatchBlocksV1.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<WatchBlocksV1> {
    @Override
    public void write(final JsonWriter jsonWriter, final WatchBlocksV1 enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public WatchBlocksV1 read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return WatchBlocksV1.fromValue(value);
    }
  }
}

