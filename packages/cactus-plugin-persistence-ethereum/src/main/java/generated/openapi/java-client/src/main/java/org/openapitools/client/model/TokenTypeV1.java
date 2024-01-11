/*
 * Hyperledger Cactus Plugin - Persistence Ethereum
 * Synchronizes state of an ethereum ledger into a DB that can later be viewed in GUI
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
 * Gets or Sets TokenTypeV1
 */
@JsonAdapter(TokenTypeV1.Adapter.class)
public enum TokenTypeV1 {
  
  /**
   * EIP-20: Token Standard
   */
  ERC20("erc20"),
  
  /**
   * EIP-721: Non-Fungible Token Standard
   */
  ERC721("erc721");

  private String value;

  TokenTypeV1(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static TokenTypeV1 fromValue(String value) {
    for (TokenTypeV1 b : TokenTypeV1.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<TokenTypeV1> {
    @Override
    public void write(final JsonWriter jsonWriter, final TokenTypeV1 enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public TokenTypeV1 read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return TokenTypeV1.fromValue(value);
    }
  }
}

