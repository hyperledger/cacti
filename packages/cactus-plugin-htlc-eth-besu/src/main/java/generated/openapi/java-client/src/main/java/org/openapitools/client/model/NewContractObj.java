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
import org.openapitools.client.model.NewContractObjGas;
import org.openapitools.client.model.Web3SigningCredential;

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
 * NewContractObj
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class NewContractObj {
  public static final String SERIALIZED_NAME_CONTRACT_ADDRESS = "contractAddress";
  @SerializedName(SERIALIZED_NAME_CONTRACT_ADDRESS)
  private String contractAddress;

  public static final String SERIALIZED_NAME_INPUT_AMOUNT = "inputAmount";
  @SerializedName(SERIALIZED_NAME_INPUT_AMOUNT)
  private BigDecimal inputAmount;

  public static final String SERIALIZED_NAME_OUTPUT_AMOUNT = "outputAmount";
  @SerializedName(SERIALIZED_NAME_OUTPUT_AMOUNT)
  private BigDecimal outputAmount;

  public static final String SERIALIZED_NAME_EXPIRATION = "expiration";
  @SerializedName(SERIALIZED_NAME_EXPIRATION)
  private BigDecimal expiration;

  public static final String SERIALIZED_NAME_HASH_LOCK = "hashLock";
  @SerializedName(SERIALIZED_NAME_HASH_LOCK)
  private String hashLock;

  public static final String SERIALIZED_NAME_RECEIVER = "receiver";
  @SerializedName(SERIALIZED_NAME_RECEIVER)
  private String receiver;

  public static final String SERIALIZED_NAME_OUTPUT_NETWORK = "outputNetwork";
  @SerializedName(SERIALIZED_NAME_OUTPUT_NETWORK)
  private String outputNetwork;

  public static final String SERIALIZED_NAME_OUTPUT_ADDRESS = "outputAddress";
  @SerializedName(SERIALIZED_NAME_OUTPUT_ADDRESS)
  private String outputAddress;

  public static final String SERIALIZED_NAME_CONNECTOR_ID = "connectorId";
  @SerializedName(SERIALIZED_NAME_CONNECTOR_ID)
  private String connectorId;

  public static final String SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL = "web3SigningCredential";
  @SerializedName(SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL)
  private Web3SigningCredential web3SigningCredential;

  public static final String SERIALIZED_NAME_KEYCHAIN_ID = "keychainId";
  @SerializedName(SERIALIZED_NAME_KEYCHAIN_ID)
  private String keychainId;

  public static final String SERIALIZED_NAME_GAS = "gas";
  @SerializedName(SERIALIZED_NAME_GAS)
  private NewContractObjGas gas;

  public NewContractObj() {
  }

  public NewContractObj contractAddress(String contractAddress) {
    
    this.contractAddress = contractAddress;
    return this;
  }

   /**
   * Contract address
   * @return contractAddress
  **/
  @javax.annotation.Nonnull
  public String getContractAddress() {
    return contractAddress;
  }


  public void setContractAddress(String contractAddress) {
    this.contractAddress = contractAddress;
  }


  public NewContractObj inputAmount(BigDecimal inputAmount) {
    
    this.inputAmount = inputAmount;
    return this;
  }

   /**
   * Get inputAmount
   * @return inputAmount
  **/
  @javax.annotation.Nullable
  public BigDecimal getInputAmount() {
    return inputAmount;
  }


  public void setInputAmount(BigDecimal inputAmount) {
    this.inputAmount = inputAmount;
  }


  public NewContractObj outputAmount(BigDecimal outputAmount) {
    
    this.outputAmount = outputAmount;
    return this;
  }

   /**
   * Get outputAmount
   * @return outputAmount
  **/
  @javax.annotation.Nonnull
  public BigDecimal getOutputAmount() {
    return outputAmount;
  }


  public void setOutputAmount(BigDecimal outputAmount) {
    this.outputAmount = outputAmount;
  }


  public NewContractObj expiration(BigDecimal expiration) {
    
    this.expiration = expiration;
    return this;
  }

   /**
   * Get expiration
   * @return expiration
  **/
  @javax.annotation.Nonnull
  public BigDecimal getExpiration() {
    return expiration;
  }


  public void setExpiration(BigDecimal expiration) {
    this.expiration = expiration;
  }


  public NewContractObj hashLock(String hashLock) {
    
    this.hashLock = hashLock;
    return this;
  }

   /**
   * Get hashLock
   * @return hashLock
  **/
  @javax.annotation.Nonnull
  public String getHashLock() {
    return hashLock;
  }


  public void setHashLock(String hashLock) {
    this.hashLock = hashLock;
  }


  public NewContractObj receiver(String receiver) {
    
    this.receiver = receiver;
    return this;
  }

   /**
   * Get receiver
   * @return receiver
  **/
  @javax.annotation.Nullable
  public String getReceiver() {
    return receiver;
  }


  public void setReceiver(String receiver) {
    this.receiver = receiver;
  }


  public NewContractObj outputNetwork(String outputNetwork) {
    
    this.outputNetwork = outputNetwork;
    return this;
  }

   /**
   * Get outputNetwork
   * @return outputNetwork
  **/
  @javax.annotation.Nonnull
  public String getOutputNetwork() {
    return outputNetwork;
  }


  public void setOutputNetwork(String outputNetwork) {
    this.outputNetwork = outputNetwork;
  }


  public NewContractObj outputAddress(String outputAddress) {
    
    this.outputAddress = outputAddress;
    return this;
  }

   /**
   * Get outputAddress
   * @return outputAddress
  **/
  @javax.annotation.Nonnull
  public String getOutputAddress() {
    return outputAddress;
  }


  public void setOutputAddress(String outputAddress) {
    this.outputAddress = outputAddress;
  }


  public NewContractObj connectorId(String connectorId) {
    
    this.connectorId = connectorId;
    return this;
  }

   /**
   * connectorId for the connector besu plugin
   * @return connectorId
  **/
  @javax.annotation.Nonnull
  public String getConnectorId() {
    return connectorId;
  }


  public void setConnectorId(String connectorId) {
    this.connectorId = connectorId;
  }


  public NewContractObj web3SigningCredential(Web3SigningCredential web3SigningCredential) {
    
    this.web3SigningCredential = web3SigningCredential;
    return this;
  }

   /**
   * Get web3SigningCredential
   * @return web3SigningCredential
  **/
  @javax.annotation.Nonnull
  public Web3SigningCredential getWeb3SigningCredential() {
    return web3SigningCredential;
  }


  public void setWeb3SigningCredential(Web3SigningCredential web3SigningCredential) {
    this.web3SigningCredential = web3SigningCredential;
  }


  public NewContractObj keychainId(String keychainId) {
    
    this.keychainId = keychainId;
    return this;
  }

   /**
   * keychainId for the keychian plugin
   * @return keychainId
  **/
  @javax.annotation.Nonnull
  public String getKeychainId() {
    return keychainId;
  }


  public void setKeychainId(String keychainId) {
    this.keychainId = keychainId;
  }


  public NewContractObj gas(NewContractObjGas gas) {
    
    this.gas = gas;
    return this;
  }

   /**
   * Get gas
   * @return gas
  **/
  @javax.annotation.Nullable
  public NewContractObjGas getGas() {
    return gas;
  }


  public void setGas(NewContractObjGas gas) {
    this.gas = gas;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    NewContractObj newContractObj = (NewContractObj) o;
    return Objects.equals(this.contractAddress, newContractObj.contractAddress) &&
        Objects.equals(this.inputAmount, newContractObj.inputAmount) &&
        Objects.equals(this.outputAmount, newContractObj.outputAmount) &&
        Objects.equals(this.expiration, newContractObj.expiration) &&
        Objects.equals(this.hashLock, newContractObj.hashLock) &&
        Objects.equals(this.receiver, newContractObj.receiver) &&
        Objects.equals(this.outputNetwork, newContractObj.outputNetwork) &&
        Objects.equals(this.outputAddress, newContractObj.outputAddress) &&
        Objects.equals(this.connectorId, newContractObj.connectorId) &&
        Objects.equals(this.web3SigningCredential, newContractObj.web3SigningCredential) &&
        Objects.equals(this.keychainId, newContractObj.keychainId) &&
        Objects.equals(this.gas, newContractObj.gas);
  }

  @Override
  public int hashCode() {
    return Objects.hash(contractAddress, inputAmount, outputAmount, expiration, hashLock, receiver, outputNetwork, outputAddress, connectorId, web3SigningCredential, keychainId, gas);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class NewContractObj {\n");
    sb.append("    contractAddress: ").append(toIndentedString(contractAddress)).append("\n");
    sb.append("    inputAmount: ").append(toIndentedString(inputAmount)).append("\n");
    sb.append("    outputAmount: ").append(toIndentedString(outputAmount)).append("\n");
    sb.append("    expiration: ").append(toIndentedString(expiration)).append("\n");
    sb.append("    hashLock: ").append(toIndentedString(hashLock)).append("\n");
    sb.append("    receiver: ").append(toIndentedString(receiver)).append("\n");
    sb.append("    outputNetwork: ").append(toIndentedString(outputNetwork)).append("\n");
    sb.append("    outputAddress: ").append(toIndentedString(outputAddress)).append("\n");
    sb.append("    connectorId: ").append(toIndentedString(connectorId)).append("\n");
    sb.append("    web3SigningCredential: ").append(toIndentedString(web3SigningCredential)).append("\n");
    sb.append("    keychainId: ").append(toIndentedString(keychainId)).append("\n");
    sb.append("    gas: ").append(toIndentedString(gas)).append("\n");
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
    openapiFields.add("contractAddress");
    openapiFields.add("inputAmount");
    openapiFields.add("outputAmount");
    openapiFields.add("expiration");
    openapiFields.add("hashLock");
    openapiFields.add("receiver");
    openapiFields.add("outputNetwork");
    openapiFields.add("outputAddress");
    openapiFields.add("connectorId");
    openapiFields.add("web3SigningCredential");
    openapiFields.add("keychainId");
    openapiFields.add("gas");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("contractAddress");
    openapiRequiredFields.add("outputAmount");
    openapiRequiredFields.add("expiration");
    openapiRequiredFields.add("hashLock");
    openapiRequiredFields.add("outputNetwork");
    openapiRequiredFields.add("outputAddress");
    openapiRequiredFields.add("connectorId");
    openapiRequiredFields.add("web3SigningCredential");
    openapiRequiredFields.add("keychainId");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to NewContractObj
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!NewContractObj.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in NewContractObj is not found in the empty JSON string", NewContractObj.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!NewContractObj.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `NewContractObj` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : NewContractObj.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("contractAddress").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `contractAddress` to be a primitive type in the JSON string but got `%s`", jsonObj.get("contractAddress").toString()));
      }
      if (!jsonObj.get("hashLock").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hashLock` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hashLock").toString()));
      }
      if ((jsonObj.get("receiver") != null && !jsonObj.get("receiver").isJsonNull()) && !jsonObj.get("receiver").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `receiver` to be a primitive type in the JSON string but got `%s`", jsonObj.get("receiver").toString()));
      }
      if (!jsonObj.get("outputNetwork").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `outputNetwork` to be a primitive type in the JSON string but got `%s`", jsonObj.get("outputNetwork").toString()));
      }
      if (!jsonObj.get("outputAddress").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `outputAddress` to be a primitive type in the JSON string but got `%s`", jsonObj.get("outputAddress").toString()));
      }
      if (!jsonObj.get("connectorId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `connectorId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("connectorId").toString()));
      }
      // validate the required field `web3SigningCredential`
      Web3SigningCredential.validateJsonObject(jsonObj.getAsJsonObject("web3SigningCredential"));
      if (!jsonObj.get("keychainId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `keychainId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("keychainId").toString()));
      }
      // validate the optional field `gas`
      if (jsonObj.get("gas") != null && !jsonObj.get("gas").isJsonNull()) {
        NewContractObjGas.validateJsonObject(jsonObj.getAsJsonObject("gas"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!NewContractObj.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'NewContractObj' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<NewContractObj> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(NewContractObj.class));

       return (TypeAdapter<T>) new TypeAdapter<NewContractObj>() {
           @Override
           public void write(JsonWriter out, NewContractObj value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public NewContractObj read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of NewContractObj given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of NewContractObj
  * @throws IOException if the JSON string is invalid with respect to NewContractObj
  */
  public static NewContractObj fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, NewContractObj.class);
  }

 /**
  * Convert an instance of NewContractObj to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

