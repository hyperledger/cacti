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
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.math.BigDecimal;
import org.openapitools.jackson.nullable.JsonNullable;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.TypeAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * EvmTransaction
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class EvmTransaction {
  public static final String SERIALIZED_NAME_HASH = "hash";
  @SerializedName(SERIALIZED_NAME_HASH)
  private String hash;

  public static final String SERIALIZED_NAME_NONCE = "nonce";
  @SerializedName(SERIALIZED_NAME_NONCE)
  private BigDecimal nonce;

  public static final String SERIALIZED_NAME_BLOCK_HASH = "blockHash";
  @SerializedName(SERIALIZED_NAME_BLOCK_HASH)
  private Object blockHash = null;

  public static final String SERIALIZED_NAME_BLOCK_NUMBER = "blockNumber";
  @SerializedName(SERIALIZED_NAME_BLOCK_NUMBER)
  private Object blockNumber = null;

  public static final String SERIALIZED_NAME_TRANSACTION_INDEX = "transactionIndex";
  @SerializedName(SERIALIZED_NAME_TRANSACTION_INDEX)
  private Object transactionIndex = null;

  public static final String SERIALIZED_NAME_FROM = "from";
  @SerializedName(SERIALIZED_NAME_FROM)
  private String from;

  public static final String SERIALIZED_NAME_TO = "to";
  @SerializedName(SERIALIZED_NAME_TO)
  private Object to = null;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private String value;

  public static final String SERIALIZED_NAME_GAS_PRICE = "gasPrice";
  @SerializedName(SERIALIZED_NAME_GAS_PRICE)
  private String gasPrice;

  public static final String SERIALIZED_NAME_GAS = "gas";
  @SerializedName(SERIALIZED_NAME_GAS)
  private BigDecimal gas;

  public static final String SERIALIZED_NAME_INPUT = "input";
  @SerializedName(SERIALIZED_NAME_INPUT)
  private String input;

  public EvmTransaction() {
  }

  public EvmTransaction hash(String hash) {
    
    this.hash = hash;
    return this;
  }

   /**
   * Get hash
   * @return hash
  **/
  @javax.annotation.Nullable
  public String getHash() {
    return hash;
  }


  public void setHash(String hash) {
    this.hash = hash;
  }


  public EvmTransaction nonce(BigDecimal nonce) {
    
    this.nonce = nonce;
    return this;
  }

   /**
   * Get nonce
   * @return nonce
  **/
  @javax.annotation.Nullable
  public BigDecimal getNonce() {
    return nonce;
  }


  public void setNonce(BigDecimal nonce) {
    this.nonce = nonce;
  }


  public EvmTransaction blockHash(Object blockHash) {
    
    this.blockHash = blockHash;
    return this;
  }

   /**
   * Get blockHash
   * @return blockHash
  **/
  @javax.annotation.Nullable
  public Object getBlockHash() {
    return blockHash;
  }


  public void setBlockHash(Object blockHash) {
    this.blockHash = blockHash;
  }


  public EvmTransaction blockNumber(Object blockNumber) {
    
    this.blockNumber = blockNumber;
    return this;
  }

   /**
   * Get blockNumber
   * @return blockNumber
  **/
  @javax.annotation.Nullable
  public Object getBlockNumber() {
    return blockNumber;
  }


  public void setBlockNumber(Object blockNumber) {
    this.blockNumber = blockNumber;
  }


  public EvmTransaction transactionIndex(Object transactionIndex) {
    
    this.transactionIndex = transactionIndex;
    return this;
  }

   /**
   * Get transactionIndex
   * @return transactionIndex
  **/
  @javax.annotation.Nullable
  public Object getTransactionIndex() {
    return transactionIndex;
  }


  public void setTransactionIndex(Object transactionIndex) {
    this.transactionIndex = transactionIndex;
  }


  public EvmTransaction from(String from) {
    
    this.from = from;
    return this;
  }

   /**
   * Get from
   * @return from
  **/
  @javax.annotation.Nullable
  public String getFrom() {
    return from;
  }


  public void setFrom(String from) {
    this.from = from;
  }


  public EvmTransaction to(Object to) {
    
    this.to = to;
    return this;
  }

   /**
   * Get to
   * @return to
  **/
  @javax.annotation.Nullable
  public Object getTo() {
    return to;
  }


  public void setTo(Object to) {
    this.to = to;
  }


  public EvmTransaction value(String value) {
    
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @javax.annotation.Nullable
  public String getValue() {
    return value;
  }


  public void setValue(String value) {
    this.value = value;
  }


  public EvmTransaction gasPrice(String gasPrice) {
    
    this.gasPrice = gasPrice;
    return this;
  }

   /**
   * Get gasPrice
   * @return gasPrice
  **/
  @javax.annotation.Nullable
  public String getGasPrice() {
    return gasPrice;
  }


  public void setGasPrice(String gasPrice) {
    this.gasPrice = gasPrice;
  }


  public EvmTransaction gas(BigDecimal gas) {
    
    this.gas = gas;
    return this;
  }

   /**
   * Get gas
   * @return gas
  **/
  @javax.annotation.Nullable
  public BigDecimal getGas() {
    return gas;
  }


  public void setGas(BigDecimal gas) {
    this.gas = gas;
  }


  public EvmTransaction input(String input) {
    
    this.input = input;
    return this;
  }

   /**
   * Get input
   * @return input
  **/
  @javax.annotation.Nullable
  public String getInput() {
    return input;
  }


  public void setInput(String input) {
    this.input = input;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    EvmTransaction evmTransaction = (EvmTransaction) o;
    return Objects.equals(this.hash, evmTransaction.hash) &&
        Objects.equals(this.nonce, evmTransaction.nonce) &&
        Objects.equals(this.blockHash, evmTransaction.blockHash) &&
        Objects.equals(this.blockNumber, evmTransaction.blockNumber) &&
        Objects.equals(this.transactionIndex, evmTransaction.transactionIndex) &&
        Objects.equals(this.from, evmTransaction.from) &&
        Objects.equals(this.to, evmTransaction.to) &&
        Objects.equals(this.value, evmTransaction.value) &&
        Objects.equals(this.gasPrice, evmTransaction.gasPrice) &&
        Objects.equals(this.gas, evmTransaction.gas) &&
        Objects.equals(this.input, evmTransaction.input);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(hash, nonce, blockHash, blockNumber, transactionIndex, from, to, value, gasPrice, gas, input);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class EvmTransaction {\n");
    sb.append("    hash: ").append(toIndentedString(hash)).append("\n");
    sb.append("    nonce: ").append(toIndentedString(nonce)).append("\n");
    sb.append("    blockHash: ").append(toIndentedString(blockHash)).append("\n");
    sb.append("    blockNumber: ").append(toIndentedString(blockNumber)).append("\n");
    sb.append("    transactionIndex: ").append(toIndentedString(transactionIndex)).append("\n");
    sb.append("    from: ").append(toIndentedString(from)).append("\n");
    sb.append("    to: ").append(toIndentedString(to)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    gasPrice: ").append(toIndentedString(gasPrice)).append("\n");
    sb.append("    gas: ").append(toIndentedString(gas)).append("\n");
    sb.append("    input: ").append(toIndentedString(input)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("hash");
    openapiFields.add("nonce");
    openapiFields.add("blockHash");
    openapiFields.add("blockNumber");
    openapiFields.add("transactionIndex");
    openapiFields.add("from");
    openapiFields.add("to");
    openapiFields.add("value");
    openapiFields.add("gasPrice");
    openapiFields.add("gas");
    openapiFields.add("input");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to EvmTransaction
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!EvmTransaction.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in EvmTransaction is not found in the empty JSON string", EvmTransaction.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!EvmTransaction.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `EvmTransaction` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("hash") != null && !jsonObj.get("hash").isJsonNull()) && !jsonObj.get("hash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hash").toString()));
      }
      if ((jsonObj.get("from") != null && !jsonObj.get("from").isJsonNull()) && !jsonObj.get("from").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `from` to be a primitive type in the JSON string but got `%s`", jsonObj.get("from").toString()));
      }
      if ((jsonObj.get("value") != null && !jsonObj.get("value").isJsonNull()) && !jsonObj.get("value").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `value` to be a primitive type in the JSON string but got `%s`", jsonObj.get("value").toString()));
      }
      if ((jsonObj.get("gasPrice") != null && !jsonObj.get("gasPrice").isJsonNull()) && !jsonObj.get("gasPrice").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `gasPrice` to be a primitive type in the JSON string but got `%s`", jsonObj.get("gasPrice").toString()));
      }
      if ((jsonObj.get("input") != null && !jsonObj.get("input").isJsonNull()) && !jsonObj.get("input").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `input` to be a primitive type in the JSON string but got `%s`", jsonObj.get("input").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!EvmTransaction.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'EvmTransaction' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<EvmTransaction> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(EvmTransaction.class));

       return (TypeAdapter<T>) new TypeAdapter<EvmTransaction>() {
           @Override
           public void write(JsonWriter out, EvmTransaction value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public EvmTransaction read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of EvmTransaction given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of EvmTransaction
  * @throws IOException if the JSON string is invalid with respect to EvmTransaction
  */
  public static EvmTransaction fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, EvmTransaction.class);
  }

 /**
  * Convert an instance of EvmTransaction to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

