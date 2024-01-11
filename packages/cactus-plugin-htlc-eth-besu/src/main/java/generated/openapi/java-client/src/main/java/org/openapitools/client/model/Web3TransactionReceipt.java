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
 * Web3TransactionReceipt
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class Web3TransactionReceipt {
  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private Boolean status;

  public static final String SERIALIZED_NAME_TRANSACTION_HASH = "transactionHash";
  @SerializedName(SERIALIZED_NAME_TRANSACTION_HASH)
  private String transactionHash;

  public static final String SERIALIZED_NAME_TRANSACTION_INDEX = "transactionIndex";
  @SerializedName(SERIALIZED_NAME_TRANSACTION_INDEX)
  private BigDecimal transactionIndex;

  public static final String SERIALIZED_NAME_BLOCK_HASH = "blockHash";
  @SerializedName(SERIALIZED_NAME_BLOCK_HASH)
  private String blockHash;

  public static final String SERIALIZED_NAME_BLOCK_NUMBER = "blockNumber";
  @SerializedName(SERIALIZED_NAME_BLOCK_NUMBER)
  private BigDecimal blockNumber;

  public static final String SERIALIZED_NAME_GAS_USED = "gasUsed";
  @SerializedName(SERIALIZED_NAME_GAS_USED)
  private BigDecimal gasUsed;

  public static final String SERIALIZED_NAME_CONTRACT_ADDRESS = "contractAddress";
  @SerializedName(SERIALIZED_NAME_CONTRACT_ADDRESS)
  private String contractAddress;

  public static final String SERIALIZED_NAME_FROM = "from";
  @SerializedName(SERIALIZED_NAME_FROM)
  private String from;

  public static final String SERIALIZED_NAME_TO = "to";
  @SerializedName(SERIALIZED_NAME_TO)
  private String to;

  public Web3TransactionReceipt() {
  }

  public Web3TransactionReceipt status(Boolean status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nonnull
  public Boolean getStatus() {
    return status;
  }


  public void setStatus(Boolean status) {
    this.status = status;
  }


  public Web3TransactionReceipt transactionHash(String transactionHash) {
    
    this.transactionHash = transactionHash;
    return this;
  }

   /**
   * Get transactionHash
   * @return transactionHash
  **/
  @javax.annotation.Nonnull
  public String getTransactionHash() {
    return transactionHash;
  }


  public void setTransactionHash(String transactionHash) {
    this.transactionHash = transactionHash;
  }


  public Web3TransactionReceipt transactionIndex(BigDecimal transactionIndex) {
    
    this.transactionIndex = transactionIndex;
    return this;
  }

   /**
   * Get transactionIndex
   * @return transactionIndex
  **/
  @javax.annotation.Nonnull
  public BigDecimal getTransactionIndex() {
    return transactionIndex;
  }


  public void setTransactionIndex(BigDecimal transactionIndex) {
    this.transactionIndex = transactionIndex;
  }


  public Web3TransactionReceipt blockHash(String blockHash) {
    
    this.blockHash = blockHash;
    return this;
  }

   /**
   * Get blockHash
   * @return blockHash
  **/
  @javax.annotation.Nonnull
  public String getBlockHash() {
    return blockHash;
  }


  public void setBlockHash(String blockHash) {
    this.blockHash = blockHash;
  }


  public Web3TransactionReceipt blockNumber(BigDecimal blockNumber) {
    
    this.blockNumber = blockNumber;
    return this;
  }

   /**
   * Get blockNumber
   * @return blockNumber
  **/
  @javax.annotation.Nonnull
  public BigDecimal getBlockNumber() {
    return blockNumber;
  }


  public void setBlockNumber(BigDecimal blockNumber) {
    this.blockNumber = blockNumber;
  }


  public Web3TransactionReceipt gasUsed(BigDecimal gasUsed) {
    
    this.gasUsed = gasUsed;
    return this;
  }

   /**
   * Get gasUsed
   * @return gasUsed
  **/
  @javax.annotation.Nonnull
  public BigDecimal getGasUsed() {
    return gasUsed;
  }


  public void setGasUsed(BigDecimal gasUsed) {
    this.gasUsed = gasUsed;
  }


  public Web3TransactionReceipt contractAddress(String contractAddress) {
    
    this.contractAddress = contractAddress;
    return this;
  }

   /**
   * Get contractAddress
   * @return contractAddress
  **/
  @javax.annotation.Nullable
  public String getContractAddress() {
    return contractAddress;
  }


  public void setContractAddress(String contractAddress) {
    this.contractAddress = contractAddress;
  }


  public Web3TransactionReceipt from(String from) {
    
    this.from = from;
    return this;
  }

   /**
   * Get from
   * @return from
  **/
  @javax.annotation.Nonnull
  public String getFrom() {
    return from;
  }


  public void setFrom(String from) {
    this.from = from;
  }


  public Web3TransactionReceipt to(String to) {
    
    this.to = to;
    return this;
  }

   /**
   * Get to
   * @return to
  **/
  @javax.annotation.Nonnull
  public String getTo() {
    return to;
  }


  public void setTo(String to) {
    this.to = to;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the Web3TransactionReceipt instance itself
   */
  public Web3TransactionReceipt putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Web3TransactionReceipt web3TransactionReceipt = (Web3TransactionReceipt) o;
    return Objects.equals(this.status, web3TransactionReceipt.status) &&
        Objects.equals(this.transactionHash, web3TransactionReceipt.transactionHash) &&
        Objects.equals(this.transactionIndex, web3TransactionReceipt.transactionIndex) &&
        Objects.equals(this.blockHash, web3TransactionReceipt.blockHash) &&
        Objects.equals(this.blockNumber, web3TransactionReceipt.blockNumber) &&
        Objects.equals(this.gasUsed, web3TransactionReceipt.gasUsed) &&
        Objects.equals(this.contractAddress, web3TransactionReceipt.contractAddress) &&
        Objects.equals(this.from, web3TransactionReceipt.from) &&
        Objects.equals(this.to, web3TransactionReceipt.to)&&
        Objects.equals(this.additionalProperties, web3TransactionReceipt.additionalProperties);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(status, transactionHash, transactionIndex, blockHash, blockNumber, gasUsed, contractAddress, from, to, additionalProperties);
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
    sb.append("class Web3TransactionReceipt {\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    transactionHash: ").append(toIndentedString(transactionHash)).append("\n");
    sb.append("    transactionIndex: ").append(toIndentedString(transactionIndex)).append("\n");
    sb.append("    blockHash: ").append(toIndentedString(blockHash)).append("\n");
    sb.append("    blockNumber: ").append(toIndentedString(blockNumber)).append("\n");
    sb.append("    gasUsed: ").append(toIndentedString(gasUsed)).append("\n");
    sb.append("    contractAddress: ").append(toIndentedString(contractAddress)).append("\n");
    sb.append("    from: ").append(toIndentedString(from)).append("\n");
    sb.append("    to: ").append(toIndentedString(to)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
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
    openapiFields.add("status");
    openapiFields.add("transactionHash");
    openapiFields.add("transactionIndex");
    openapiFields.add("blockHash");
    openapiFields.add("blockNumber");
    openapiFields.add("gasUsed");
    openapiFields.add("contractAddress");
    openapiFields.add("from");
    openapiFields.add("to");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("status");
    openapiRequiredFields.add("transactionHash");
    openapiRequiredFields.add("transactionIndex");
    openapiRequiredFields.add("blockHash");
    openapiRequiredFields.add("blockNumber");
    openapiRequiredFields.add("gasUsed");
    openapiRequiredFields.add("from");
    openapiRequiredFields.add("to");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Web3TransactionReceipt
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!Web3TransactionReceipt.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Web3TransactionReceipt is not found in the empty JSON string", Web3TransactionReceipt.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : Web3TransactionReceipt.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("transactionHash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `transactionHash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("transactionHash").toString()));
      }
      if (!jsonObj.get("blockHash").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `blockHash` to be a primitive type in the JSON string but got `%s`", jsonObj.get("blockHash").toString()));
      }
      if ((jsonObj.get("contractAddress") != null && !jsonObj.get("contractAddress").isJsonNull()) && !jsonObj.get("contractAddress").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `contractAddress` to be a primitive type in the JSON string but got `%s`", jsonObj.get("contractAddress").toString()));
      }
      if (!jsonObj.get("from").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `from` to be a primitive type in the JSON string but got `%s`", jsonObj.get("from").toString()));
      }
      if (!jsonObj.get("to").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `to` to be a primitive type in the JSON string but got `%s`", jsonObj.get("to").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Web3TransactionReceipt.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Web3TransactionReceipt' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Web3TransactionReceipt> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Web3TransactionReceipt.class));

       return (TypeAdapter<T>) new TypeAdapter<Web3TransactionReceipt>() {
           @Override
           public void write(JsonWriter out, Web3TransactionReceipt value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additional properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public Web3TransactionReceipt read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             Web3TransactionReceipt instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of Web3TransactionReceipt given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Web3TransactionReceipt
  * @throws IOException if the JSON string is invalid with respect to Web3TransactionReceipt
  */
  public static Web3TransactionReceipt fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Web3TransactionReceipt.class);
  }

 /**
  * Convert an instance of Web3TransactionReceipt to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

