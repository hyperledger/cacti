/*
 * Hyperledger Cactus Plugin - Connector Xdai
 * Can perform basic tasks on a Xdai ledger
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
 * Enumerates the possible types of receipts that can be waited for by someone or something that has requested the execution of a transaction on a ledger.
 */
@JsonAdapter(ReceiptType.Adapter.class)
public enum ReceiptType {
  
  NODE_TX_POOL_ACK("NODE_TX_POOL_ACK"),
  
  LEDGER_BLOCK_ACK("LEDGER_BLOCK_ACK");

  private String value;

  ReceiptType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ReceiptType fromValue(String value) {
    for (ReceiptType b : ReceiptType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ReceiptType> {
    @Override
    public void write(final JsonWriter jsonWriter, final ReceiptType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ReceiptType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ReceiptType.fromValue(value);
    }
  }
}

