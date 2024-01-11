/*
 * Hyperledger Cactus Plugin - Connector Quorum
 * Can perform basic tasks on a Quorum ledger
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
import org.openapitools.client.model.QuorumTransactionConfigFrom;

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
 * Web3BlockHeader
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class Web3BlockHeader {
  public static final String SERIALIZED_NAME_NUMBER = "number";
  @SerializedName(SERIALIZED_NAME_NUMBER)
  private BigDecimal number;

  public static final String SERIALIZED_NAME_HASH = "hash";
  @SerializedName(SERIALIZED_NAME_HASH)
  private String hash;

  public static final String SERIALIZED_NAME_PARENT_HASH = "parentHash";
  @SerializedName(SERIALIZED_NAME_PARENT_HASH)
  private String parentHash;

  public static final String SERIALIZED_NAME_NONCE = "nonce";
  @SerializedName(SERIALIZED_NAME_NONCE)
  private String nonce;

  public static final String SERIALIZED_NAME_SHA3_UNCLES = "sha3Uncles";
  @SerializedName(SERIALIZED_NAME_SHA3_UNCLES)
  private String sha3Uncles;

  public static final String SERIALIZED_NAME_LOGS_BLOOM = "logsBloom";
  @SerializedName(SERIALIZED_NAME_LOGS_BLOOM)
  private String logsBloom;

  public static final String SERIALIZED_NAME_TRANSACTIONS_ROOT = "transactionsRoot";
  @SerializedName(SERIALIZED_NAME_TRANSACTIONS_ROOT)
  private String transactionsRoot;

  public static final String SERIALIZED_NAME_STATE_ROOT = "stateRoot";
  @SerializedName(SERIALIZED_NAME_STATE_ROOT)
  private String stateRoot;

  public static final String SERIALIZED_NAME_RECEIPTS_ROOT = "receiptsRoot";
  @SerializedName(SERIALIZED_NAME_RECEIPTS_ROOT)
  private String receiptsRoot;

  public static final String SERIALIZED_NAME_DIFFICULTY = "difficulty";
  @SerializedName(SERIALIZED_NAME_DIFFICULTY)
  private String difficulty;

  public static final String SERIALIZED_NAME_MIX_HASH = "mixHash";
  @SerializedName(SERIALIZED_NAME_MIX_HASH)
  private String mixHash;

  public static final String SERIALIZED_NAME_MINER = "miner";
  @SerializedName(SERIALIZED_NAME_MINER)
  private String miner;

  public static final String SERIALIZED_NAME_EXTRA_DATA = "extraData";
  @SerializedName(SERIALIZED_NAME_EXTRA_DATA)
  private String extraData;

  public static final String SERIALIZED_NAME_GAS_LIMIT = "gasLimit";
  @SerializedName(SERIALIZED_NAME_GAS_LIMIT)
  private Integer gasLimit;

  public static final String SERIALIZED_NAME_GAS_USED = "gasUsed";
  @SerializedName(SERIALIZED_NAME_GAS_USED)
  private Integer gasUsed;

  public static final String SERIALIZED_NAME_TIMESTAMP = "timestamp";
  @SerializedName(SERIALIZED_NAME_TIMESTAMP)
  private QuorumTransactionConfigFrom timestamp;

  public Web3BlockHeader() {
  }

  public Web3BlockHeader number(BigDecimal number) {
    
    this.number = number;
    return this;
  }

   /**
   * Get number
   * @return number
  **/
  @javax.annotation.Nonnull
  public BigDecimal getNumber() {
    return number;
  }


  public void setNumber(BigDecimal number) {
    this.number = number;
  }


  public Web3BlockHeader hash(String hash) {
    
    this.hash = hash;
    return this;
  }

   /**
   * Get hash
   * @return hash
  **/
  @javax.annotation.Nonnull
  public String getHash() {
    return hash;
  }


  public void setHash(String hash) {
    this.hash = hash;
  }


  public Web3BlockHeader parentHash(String parentHash) {
    
    this.parentHash = parentHash;
    return this;
  }

   /**
   * Get parentHash
   * @return parentHash
  **/
  @javax.annotation.Nonnull
  public String getParentHash() {
    return parentHash;
  }


  public void setParentHash(String parentHash) {
    this.parentHash = parentHash;
  }


  public Web3BlockHeader nonce(String nonce) {
    
    this.nonce = nonce;
    return this;
  }

   /**
   * Get nonce
   * @return nonce
  **/
  @javax.annotation.Nonnull
  public String getNonce() {
    return nonce;
  }


  public void setNonce(String nonce) {
    this.nonce = nonce;
  }


  public Web3BlockHeader sha3Uncles(String sha3Uncles) {
    
    this.sha3Uncles = sha3Uncles;
    return this;
  }

   /**
   * Get sha3Uncles
   * @return sha3Uncles
  **/
  @javax.annotation.Nonnull
  public String getSha3Uncles() {
    return sha3Uncles;
  }


  public void setSha3Uncles(String sha3Uncles) {
    this.sha3Uncles = sha3Uncles;
  }


  public Web3BlockHeader logsBloom(String logsBloom) {
    
    this.logsBloom = logsBloom;
    return this;
  }

   /**
   * Get logsBloom
   * @return logsBloom
  **/
  @javax.annotation.Nonnull
  public String getLogsBloom() {
    return logsBloom;
  }


  public void setLogsBloom(String logsBloom) {
    this.logsBloom = logsBloom;
  }


  public Web3BlockHeader transactionsRoot(String transactionsRoot) {
    
    this.transactionsRoot = transactionsRoot;
    return this;
  }

   /**
   * Get transactionsRoot
   * @return transactionsRoot
  **/
  @javax.annotation.Nullable
  public String getTransactionsRoot() {
    return transactionsRoot;
  }


  public void setTransactionsRoot(String transactionsRoot) {
    this.transactionsRoot = transactionsRoot;
  }


  public Web3BlockHeader stateRoot(String stateRoot) {
    
    this.stateRoot = stateRoot;
    return this;
  }

   /**
   * Get stateRoot
   * @return stateRoot
  **/
  @javax.annotation.Nonnull
  public String getStateRoot() {
    return stateRoot;
  }


  public void setStateRoot(String stateRoot) {
    this.stateRoot = stateRoot;
  }


  public Web3BlockHeader receiptsRoot(String receiptsRoot) {
    
    this.receiptsRoot = receiptsRoot;
    return this;
  }

   /**
   * Get receiptsRoot
   * @return receiptsRoot
  **/
  @javax.annotation.Nullable
  public String getReceiptsRoot() {
    return receiptsRoot;
  }


  public void setReceiptsRoot(String receiptsRoot) {
    this.receiptsRoot = receiptsRoot;
  }


  public Web3BlockHeader difficulty(String difficulty) {
    
    this.difficulty = difficulty;
    return this;
  }

   /**
   * Get difficulty
   * @return difficulty
  **/
  @javax.annotation.Nullable
  public String getDifficulty() {
    return difficulty;
  }


  public void setDifficulty(String difficulty) {
    this.difficulty = difficulty;
  }


  public Web3BlockHeader mixHash(String mixHash) {
    
    this.mixHash = mixHash;
    return this;
  }

   /**
   * Get mixHash
   * @return mixHash
  **/
  @javax.annotation.Nullable
  public String getMixHash() {
    return mixHash;
  }


  public void setMixHash(String mixHash) {
    this.mixHash = mixHash;
  }


  public Web3BlockHeader miner(String miner) {
    
    this.miner = miner;
    return this;
  }

   /**
   * Get miner
   * @return miner
  **/
  @javax.annotation.Nonnull
  public String getMiner() {
    return miner;
  }


  public void setMiner(String miner) {
    this.miner = miner;
  }


  public Web3BlockHeader extraData(String extraData) {
    
    this.extraData = extraData;
    return this;
  }

   /**
   * Get extraData
   * @return extraData
  **/
  @javax.annotation.Nonnull
  public String getExtraData() {
    return extraData;
  }


  public void setExtraData(String extraData) {
    this.extraData = extraData;
  }


  public Web3BlockHeader gasLimit(Integer gasLimit) {
    
    this.gasLimit = gasLimit;
    return this;
  }

   /**
   * Get gasLimit
   * @return gasLimit
  **/
  @javax.annotation.Nonnull
  public Integer getGasLimit() {
    return gasLimit;
  }


  public void setGasLimit(Integer gasLimit) {
    this.gasLimit = gasLimit;
  }


  public Web3BlockHeader gasUsed(Integer gasUsed) {
    
    this.gasUsed = gasUsed;
    return this;
  }

   /**
   * Get gasUsed
   * @return gasUsed
  **/
  @javax.annotation.Nonnull
  public Integer getGasUsed() {
    return gasUsed;
  }


  public void setGasUsed(Integer gasUsed) {
    this.gasUsed = gasUsed;
  }


  public Web3BlockHeader timestamp(QuorumTransactionConfigFrom timestamp) {
    
    this.timestamp = timestamp;
    return this;
  }

   /**
   * Get timestamp
   * @return timestamp
  **/
  @javax.annotation.Nonnull
  public QuorumTransactionConfigFrom getTimestamp() {
    return timestamp;
  }


  public void setTimestamp(QuorumTransactionConfigFrom timestamp) {
    this.timestamp = timestamp;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Web3BlockHeader web3BlockHeader = (Web3BlockHeader) o;
    return Objects.equals(this.number, web3BlockHeader.number) &&
        Objects.equals(this.hash, web3BlockHeader.hash) &&
        Objects.equals(this.parentHash, web3BlockHeader.parentHash) &&
        Objects.equals(this.nonce, web3BlockHeader.nonce) &&
        Objects.equals(this.sha3Uncles, web3BlockHeader.sha3Uncles) &&
        Objects.equals(this.logsBloom, web3BlockHeader.logsBloom) &&
        Objects.equals(this.transactionsRoot, web3BlockHeader.transactionsRoot) &&
        Objects.equals(this.stateRoot, web3BlockHeader.stateRoot) &&
        Objects.equals(this.receiptsRoot, web3BlockHeader.receiptsRoot) &&
        Objects.equals(this.difficulty, web3BlockHeader.difficulty) &&
        Objects.equals(this.mixHash, web3BlockHeader.mixHash) &&
        Objects.equals(this.miner, web3BlockHeader.miner) &&
        Objects.equals(this.extraData, web3BlockHeader.extraData) &&
        Objects.equals(this.gasLimit, web3BlockHeader.gasLimit) &&
        Objects.equals(this.gasUsed, web3BlockHeader.gasUsed) &&
        Objects.equals(this.timestamp, web3BlockHeader.timestamp);
  }

  @Override
  public int hashCode() {
    return Objects.hash(number, hash, parentHash, nonce, sha3Uncles, logsBloom, transactionsRoot, stateRoot, receiptsRoot, difficulty, mixHash, miner, extraData, gasLimit, gasUsed, timestamp);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Web3BlockHeader {\n");
    sb.append("    number: ").append(toIndentedString(number)).append("\n");
    sb.append("    hash: ").append(toIndentedString(hash)).append("\n");
    sb.append("    parentHash: ").append(toIndentedString(parentHash)).append("\n");
    sb.append("    nonce: ").append(toIndentedString(nonce)).append("\n");
    sb.append("    sha3Uncles: ").append(toIndentedString(sha3Uncles)).append("\n");
    sb.append("    logsBloom: ").append(toIndentedString(logsBloom)).append("\n");
    sb.append("    transactionsRoot: ").append(toIndentedString(transactionsRoot)).append("\n");
    sb.append("    stateRoot: ").append(toIndentedString(stateRoot)).append("\n");
    sb.append("    receiptsRoot: ").append(toIndentedString(receiptsRoot)).append("\n");
    sb.append("    difficulty: ").append(toIndentedString(difficulty)).append("\n");
    sb.append("    mixHash: ").append(toIndentedString(mixHash)).append("\n");
    sb.append("    miner: ").append(toIndentedString(miner)).append("\n");
    sb.append("    extraData: ").append(toIndentedString(extraData)).append("\n");
    sb.append("    gasLimit: ").append(toIndentedString(gasLimit)).append("\n");
    sb.append("    gasUsed: ").append(toIndentedString(gasUsed)).append("\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
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
    openapiFields.add("number");
    openapiFields.add("hash");
    openapiFields.add("parentHash");
    openapiFields.add("nonce");
    openapiFields.add("sha3Uncles");
    openapiFields.add("logsBloom");
    openapiFields.add("transactionsRoot");
    openapiFields.add("stateRoot");
    openapiFields.add("receiptsRoot");
    openapiFields.add("difficulty");
    openapiFields.add("mixHash");
    openapiFields.add("miner");
    openapiFields.add("extraData");
    openapiFields.add("gasLimit");
    openapiFields.add("gasUsed");
    openapiFields.add("timestamp");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("number");
    openapiRequiredFields.add("hash");
    openapiRequiredFields.add("parentHash");
    openapiRequiredFields.add("nonce");
    openapiRequiredFields.add("sha3Uncles");
    openapiRequiredFields.add("logsBloom");
    openapiRequiredFields.add("stateRoot");
    openapiRequiredFields.add("miner");
    openapiRequiredFields.add("extraData");
    openapiRequiredFields.add("gasLimit");
    openapiRequiredFields.add("gasUsed");
    openapiRequiredFields.add("timestamp");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Web3BlockHeader
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!Web3BlockHeader.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Web3BlockHeader is not found in the empty JSON string", Web3BlockHeader.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!Web3BlockHeader.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `Web3BlockHeader` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : Web3BlockHeader.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("hash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hash").toString()));
      }
      if (!jsonObj.get("parentHash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `parentHash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("parentHash").toString()));
      }
      if (!jsonObj.get("nonce").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `nonce` to be a primitive type in the JSON string but got `%s`", jsonObj.get("nonce").toString()));
      }
      if (!jsonObj.get("sha3Uncles").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sha3Uncles` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sha3Uncles").toString()));
      }
      if (!jsonObj.get("logsBloom").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `logsBloom` to be a primitive type in the JSON string but got `%s`", jsonObj.get("logsBloom").toString()));
      }
      if ((jsonObj.get("transactionsRoot") != null && !jsonObj.get("transactionsRoot").isJsonNull()) && !jsonObj.get("transactionsRoot").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `transactionsRoot` to be a primitive type in the JSON string but got `%s`", jsonObj.get("transactionsRoot").toString()));
      }
      if (!jsonObj.get("stateRoot").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `stateRoot` to be a primitive type in the JSON string but got `%s`", jsonObj.get("stateRoot").toString()));
      }
      if ((jsonObj.get("receiptsRoot") != null && !jsonObj.get("receiptsRoot").isJsonNull()) && !jsonObj.get("receiptsRoot").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `receiptsRoot` to be a primitive type in the JSON string but got `%s`", jsonObj.get("receiptsRoot").toString()));
      }
      if ((jsonObj.get("difficulty") != null && !jsonObj.get("difficulty").isJsonNull()) && !jsonObj.get("difficulty").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `difficulty` to be a primitive type in the JSON string but got `%s`", jsonObj.get("difficulty").toString()));
      }
      if ((jsonObj.get("mixHash") != null && !jsonObj.get("mixHash").isJsonNull()) && !jsonObj.get("mixHash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `mixHash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("mixHash").toString()));
      }
      if (!jsonObj.get("miner").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `miner` to be a primitive type in the JSON string but got `%s`", jsonObj.get("miner").toString()));
      }
      if (!jsonObj.get("extraData").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `extraData` to be a primitive type in the JSON string but got `%s`", jsonObj.get("extraData").toString()));
      }
      // validate the required field `timestamp`
      QuorumTransactionConfigFrom.validateJsonObject(jsonObj.getAsJsonObject("timestamp"));
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Web3BlockHeader.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Web3BlockHeader' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Web3BlockHeader> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Web3BlockHeader.class));

       return (TypeAdapter<T>) new TypeAdapter<Web3BlockHeader>() {
           @Override
           public void write(JsonWriter out, Web3BlockHeader value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public Web3BlockHeader read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of Web3BlockHeader given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Web3BlockHeader
  * @throws IOException if the JSON string is invalid with respect to Web3BlockHeader
  */
  public static Web3BlockHeader fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Web3BlockHeader.class);
  }

 /**
  * Convert an instance of Web3BlockHeader to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

