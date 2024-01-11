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
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
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
 * DeployContractV1Request
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class DeployContractV1Request {
  public static final String SERIALIZED_NAME_CONTRACT_NAME = "contractName";
  @SerializedName(SERIALIZED_NAME_CONTRACT_NAME)
  private String contractName;

  public static final String SERIALIZED_NAME_CONSTRUCTOR_ARGS = "constructorArgs";
  @SerializedName(SERIALIZED_NAME_CONSTRUCTOR_ARGS)
  private List<Object> constructorArgs = null;

  public static final String SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL = "web3SigningCredential";
  @SerializedName(SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL)
  private Web3SigningCredential web3SigningCredential;

  public static final String SERIALIZED_NAME_KEYCHAIN_ID = "keychainId";
  @SerializedName(SERIALIZED_NAME_KEYCHAIN_ID)
  private String keychainId;

  public static final String SERIALIZED_NAME_GAS = "gas";
  @SerializedName(SERIALIZED_NAME_GAS)
  private BigDecimal gas;

  public static final String SERIALIZED_NAME_GAS_PRICE = "gasPrice";
  @SerializedName(SERIALIZED_NAME_GAS_PRICE)
  private String gasPrice;

  public static final String SERIALIZED_NAME_TIMEOUT_MS = "timeoutMs";
  @SerializedName(SERIALIZED_NAME_TIMEOUT_MS)
  private BigDecimal timeoutMs = new BigDecimal("60000");

  public DeployContractV1Request() {
  }

  public DeployContractV1Request contractName(String contractName) {
    
    this.contractName = contractName;
    return this;
  }

   /**
   * The contract name for retrieve the contracts json on the keychain.
   * @return contractName
  **/
  @javax.annotation.Nonnull
  public String getContractName() {
    return contractName;
  }


  public void setContractName(String contractName) {
    this.contractName = contractName;
  }


  public DeployContractV1Request constructorArgs(List<Object> constructorArgs) {
    
    this.constructorArgs = constructorArgs;
    return this;
  }

  public DeployContractV1Request addConstructorArgsItem(Object constructorArgsItem) {
    if (this.constructorArgs == null) {
      this.constructorArgs = null;
    }
    this.constructorArgs.add(constructorArgsItem);
    return this;
  }

   /**
   * Get constructorArgs
   * @return constructorArgs
  **/
  @javax.annotation.Nullable
  public List<Object> getConstructorArgs() {
    return constructorArgs;
  }


  public void setConstructorArgs(List<Object> constructorArgs) {
    this.constructorArgs = constructorArgs;
  }


  public DeployContractV1Request web3SigningCredential(Web3SigningCredential web3SigningCredential) {
    
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


  public DeployContractV1Request keychainId(String keychainId) {
    
    this.keychainId = keychainId;
    return this;
  }

   /**
   * The keychainId for retrieve the contracts json.
   * @return keychainId
  **/
  @javax.annotation.Nonnull
  public String getKeychainId() {
    return keychainId;
  }


  public void setKeychainId(String keychainId) {
    this.keychainId = keychainId;
  }


  public DeployContractV1Request gas(BigDecimal gas) {
    
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


  public DeployContractV1Request gasPrice(String gasPrice) {
    
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


  public DeployContractV1Request timeoutMs(BigDecimal timeoutMs) {
    
    this.timeoutMs = timeoutMs;
    return this;
  }

   /**
   * The amount of milliseconds to wait for a transaction receipt with theaddress of the contract(which indicates successful deployment) beforegiving up and crashing.
   * minimum: 0
   * @return timeoutMs
  **/
  @javax.annotation.Nullable
  public BigDecimal getTimeoutMs() {
    return timeoutMs;
  }


  public void setTimeoutMs(BigDecimal timeoutMs) {
    this.timeoutMs = timeoutMs;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    DeployContractV1Request deployContractV1Request = (DeployContractV1Request) o;
    return Objects.equals(this.contractName, deployContractV1Request.contractName) &&
        Objects.equals(this.constructorArgs, deployContractV1Request.constructorArgs) &&
        Objects.equals(this.web3SigningCredential, deployContractV1Request.web3SigningCredential) &&
        Objects.equals(this.keychainId, deployContractV1Request.keychainId) &&
        Objects.equals(this.gas, deployContractV1Request.gas) &&
        Objects.equals(this.gasPrice, deployContractV1Request.gasPrice) &&
        Objects.equals(this.timeoutMs, deployContractV1Request.timeoutMs);
  }

  @Override
  public int hashCode() {
    return Objects.hash(contractName, constructorArgs, web3SigningCredential, keychainId, gas, gasPrice, timeoutMs);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class DeployContractV1Request {\n");
    sb.append("    contractName: ").append(toIndentedString(contractName)).append("\n");
    sb.append("    constructorArgs: ").append(toIndentedString(constructorArgs)).append("\n");
    sb.append("    web3SigningCredential: ").append(toIndentedString(web3SigningCredential)).append("\n");
    sb.append("    keychainId: ").append(toIndentedString(keychainId)).append("\n");
    sb.append("    gas: ").append(toIndentedString(gas)).append("\n");
    sb.append("    gasPrice: ").append(toIndentedString(gasPrice)).append("\n");
    sb.append("    timeoutMs: ").append(toIndentedString(timeoutMs)).append("\n");
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
    openapiFields.add("contractName");
    openapiFields.add("constructorArgs");
    openapiFields.add("web3SigningCredential");
    openapiFields.add("keychainId");
    openapiFields.add("gas");
    openapiFields.add("gasPrice");
    openapiFields.add("timeoutMs");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("contractName");
    openapiRequiredFields.add("web3SigningCredential");
    openapiRequiredFields.add("keychainId");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to DeployContractV1Request
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!DeployContractV1Request.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in DeployContractV1Request is not found in the empty JSON string", DeployContractV1Request.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!DeployContractV1Request.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `DeployContractV1Request` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : DeployContractV1Request.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("contractName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `contractName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("contractName").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("constructorArgs") != null && !jsonObj.get("constructorArgs").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `constructorArgs` to be an array in the JSON string but got `%s`", jsonObj.get("constructorArgs").toString()));
      }
      // validate the required field `web3SigningCredential`
      Web3SigningCredential.validateJsonObject(jsonObj.getAsJsonObject("web3SigningCredential"));
      if (!jsonObj.get("keychainId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `keychainId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("keychainId").toString()));
      }
      if ((jsonObj.get("gasPrice") != null && !jsonObj.get("gasPrice").isJsonNull()) && !jsonObj.get("gasPrice").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `gasPrice` to be a primitive type in the JSON string but got `%s`", jsonObj.get("gasPrice").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!DeployContractV1Request.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'DeployContractV1Request' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<DeployContractV1Request> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(DeployContractV1Request.class));

       return (TypeAdapter<T>) new TypeAdapter<DeployContractV1Request>() {
           @Override
           public void write(JsonWriter out, DeployContractV1Request value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public DeployContractV1Request read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of DeployContractV1Request given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of DeployContractV1Request
  * @throws IOException if the JSON string is invalid with respect to DeployContractV1Request
  */
  public static DeployContractV1Request fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, DeployContractV1Request.class);
  }

 /**
  * Convert an instance of DeployContractV1Request to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

