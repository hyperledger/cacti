/*
 * Hyperledger Cactus Plugin - HTLC-ETH Besu
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
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
 * Gets or Sets Web3SigningCredentialType
 */
@JsonAdapter(Web3SigningCredentialType.Adapter.class)
public enum Web3SigningCredentialType {
  
  CACTUS_KEYCHAIN_REF("CACTUS_KEYCHAIN_REF"),
  
  GETH_KEYCHAIN_PASSWORD("GETH_KEYCHAIN_PASSWORD"),
  
  PRIVATE_KEY_HEX("PRIVATE_KEY_HEX"),
  
  NONE("NONE");

  private String value;

  Web3SigningCredentialType(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static Web3SigningCredentialType fromValue(String value) {
    for (Web3SigningCredentialType b : Web3SigningCredentialType.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<Web3SigningCredentialType> {
    @Override
    public void write(final JsonWriter jsonWriter, final Web3SigningCredentialType enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public Web3SigningCredentialType read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return Web3SigningCredentialType.fromValue(value);
    }
  }
}

