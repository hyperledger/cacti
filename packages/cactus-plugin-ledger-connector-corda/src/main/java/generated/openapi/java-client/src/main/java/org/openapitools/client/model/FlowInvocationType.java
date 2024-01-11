/*
 * Hyperledger Cactus Plugin - Connector Corda
 * Can perform basic tasks on a Corda ledger
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
 * Determines which flow starting method will be used on the back-end when invoking the flow. Based on the value here the plugin back-end might invoke the rpc.startFlowDynamic() method or the rpc.startTrackedFlowDynamic() method. Streamed responses are aggregated and returned in a single response to HTTP callers who are not equipped to handle streams like WebSocket/gRPC/etc. do.
 */
@JsonAdapter(FlowInvocationType.Adapter.class)
public enum FlowInvocationType {
  
  TRACKED_FLOW_DYNAMIC("TRACKED_FLOW_DYNAMIC"),
  
  FLOW_DYNAMIC("FLOW_DYNAMIC");

  private String value;

  FlowInvocationType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static FlowInvocationType fromValue(String value) {
    for (FlowInvocationType b : FlowInvocationType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<FlowInvocationType> {
    @Override
    public void write(final JsonWriter jsonWriter, final FlowInvocationType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public FlowInvocationType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return FlowInvocationType.fromValue(value);
    }
  }
}

