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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.EthContractInvocationType;
import org.openapitools.client.model.QuorumTransactionConfigFrom;
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
 * InvokeContractV1Request
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class InvokeContractV1Request {
  public static final String SERIALIZED_NAME_CONTRACT_NAME = "contractName";
  @SerializedName(SERIALIZED_NAME_CONTRACT_NAME)
  private String contractName;

  public static final String SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL = "web3SigningCredential";
  @SerializedName(SERIALIZED_NAME_WEB3_SIGNING_CREDENTIAL)
  private Web3SigningCredential web3SigningCredential;

  public static final String SERIALIZED_NAME_INVOCATION_TYPE = "invocationType";
  @SerializedName(SERIALIZED_NAME_INVOCATION_TYPE)
  private EthContractInvocationType invocationType;

  public static final String SERIALIZED_NAME_METHOD_NAME = "methodName";
  @SerializedName(SERIALIZED_NAME_METHOD_NAME)
  private String methodName;

  public static final String SERIALIZED_NAME_PARAMS = "params";
  @SerializedName(SERIALIZED_NAME_PARAMS)
  private List<Object> params = null;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private QuorumTransactionConfigFrom value;

  public static final String SERIALIZED_NAME_GAS = "gas";
  @SerializedName(SERIALIZED_NAME_GAS)
  private QuorumTransactionConfigFrom gas;

  public static final String SERIALIZED_NAME_GAS_PRICE = "gasPrice";
  @SerializedName(SERIALIZED_NAME_GAS_PRICE)
  private QuorumTransactionConfigFrom gasPrice;

  public static final String SERIALIZED_NAME_NONCE = "nonce";
  @SerializedName(SERIALIZED_NAME_NONCE)
  private BigDecimal nonce;

  public static final String SERIALIZED_NAME_TIMEOUT_MS = "timeoutMs";
  @SerializedName(SERIALIZED_NAME_TIMEOUT_MS)
  private BigDecimal timeoutMs = new BigDecimal("60000");

  public static final String SERIALIZED_NAME_KEYCHAIN_ID = "keychainId";
  @SerializedName(SERIALIZED_NAME_KEYCHAIN_ID)
  private String keychainId;

  public InvokeContractV1Request() {
  }

  public InvokeContractV1Request contractName(String contractName) {
    
    this.contractName = contractName;
    return this;
  }

   /**
   * The contract name to find it in the keychain plugin
   * @return contractName
  **/
  @javax.annotation.Nonnull
  public String getContractName() {
    return contractName;
  }


  public void setContractName(String contractName) {
    this.contractName = contractName;
  }


  public InvokeContractV1Request web3SigningCredential(Web3SigningCredential web3SigningCredential) {
    
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


  public InvokeContractV1Request invocationType(EthContractInvocationType invocationType) {
    
    this.invocationType = invocationType;
    return this;
  }

   /**
   * Get invocationType
   * @return invocationType
  **/
  @javax.annotation.Nonnull
  public EthContractInvocationType getInvocationType() {
    return invocationType;
  }


  public void setInvocationType(EthContractInvocationType invocationType) {
    this.invocationType = invocationType;
  }


  public InvokeContractV1Request methodName(String methodName) {
    
    this.methodName = methodName;
    return this;
  }

   /**
   * The name of the contract method to invoke.
   * @return methodName
  **/
  @javax.annotation.Nonnull
  public String getMethodName() {
    return methodName;
  }


  public void setMethodName(String methodName) {
    this.methodName = methodName;
  }


  public InvokeContractV1Request params(List<Object> params) {
    
    this.params = params;
    return this;
  }

  public InvokeContractV1Request addParamsItem(Object paramsItem) {
    if (this.params == null) {
      this.params = null;
    }
    this.params.add(paramsItem);
    return this;
  }

   /**
   * The list of arguments to pass in to the contract method being invoked.
   * @return params
  **/
  @javax.annotation.Nonnull
  public List<Object> getParams() {
    return params;
  }


  public void setParams(List<Object> params) {
    this.params = params;
  }


  public InvokeContractV1Request value(QuorumTransactionConfigFrom value) {
    
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigFrom getValue() {
    return value;
  }


  public void setValue(QuorumTransactionConfigFrom value) {
    this.value = value;
  }


  public InvokeContractV1Request gas(QuorumTransactionConfigFrom gas) {
    
    this.gas = gas;
    return this;
  }

   /**
   * Get gas
   * @return gas
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigFrom getGas() {
    return gas;
  }


  public void setGas(QuorumTransactionConfigFrom gas) {
    this.gas = gas;
  }


  public InvokeContractV1Request gasPrice(QuorumTransactionConfigFrom gasPrice) {
    
    this.gasPrice = gasPrice;
    return this;
  }

   /**
   * Get gasPrice
   * @return gasPrice
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigFrom getGasPrice() {
    return gasPrice;
  }


  public void setGasPrice(QuorumTransactionConfigFrom gasPrice) {
    this.gasPrice = gasPrice;
  }


  public InvokeContractV1Request nonce(BigDecimal nonce) {
    
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


  public InvokeContractV1Request timeoutMs(BigDecimal timeoutMs) {
    
    this.timeoutMs = timeoutMs;
    return this;
  }

   /**
   * The amount of milliseconds to wait for a transaction receipt beforegiving up and crashing. Only has any effect if the invocation type is SEND
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


  public InvokeContractV1Request keychainId(String keychainId) {
    
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



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InvokeContractV1Request invokeContractV1Request = (InvokeContractV1Request) o;
    return Objects.equals(this.contractName, invokeContractV1Request.contractName) &&
        Objects.equals(this.web3SigningCredential, invokeContractV1Request.web3SigningCredential) &&
        Objects.equals(this.invocationType, invokeContractV1Request.invocationType) &&
        Objects.equals(this.methodName, invokeContractV1Request.methodName) &&
        Objects.equals(this.params, invokeContractV1Request.params) &&
        Objects.equals(this.value, invokeContractV1Request.value) &&
        Objects.equals(this.gas, invokeContractV1Request.gas) &&
        Objects.equals(this.gasPrice, invokeContractV1Request.gasPrice) &&
        Objects.equals(this.nonce, invokeContractV1Request.nonce) &&
        Objects.equals(this.timeoutMs, invokeContractV1Request.timeoutMs) &&
        Objects.equals(this.keychainId, invokeContractV1Request.keychainId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(contractName, web3SigningCredential, invocationType, methodName, params, value, gas, gasPrice, nonce, timeoutMs, keychainId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InvokeContractV1Request {\n");
    sb.append("    contractName: ").append(toIndentedString(contractName)).append("\n");
    sb.append("    web3SigningCredential: ").append(toIndentedString(web3SigningCredential)).append("\n");
    sb.append("    invocationType: ").append(toIndentedString(invocationType)).append("\n");
    sb.append("    methodName: ").append(toIndentedString(methodName)).append("\n");
    sb.append("    params: ").append(toIndentedString(params)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    gas: ").append(toIndentedString(gas)).append("\n");
    sb.append("    gasPrice: ").append(toIndentedString(gasPrice)).append("\n");
    sb.append("    nonce: ").append(toIndentedString(nonce)).append("\n");
    sb.append("    timeoutMs: ").append(toIndentedString(timeoutMs)).append("\n");
    sb.append("    keychainId: ").append(toIndentedString(keychainId)).append("\n");
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
    openapiFields.add("web3SigningCredential");
    openapiFields.add("invocationType");
    openapiFields.add("methodName");
    openapiFields.add("params");
    openapiFields.add("value");
    openapiFields.add("gas");
    openapiFields.add("gasPrice");
    openapiFields.add("nonce");
    openapiFields.add("timeoutMs");
    openapiFields.add("keychainId");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("contractName");
    openapiRequiredFields.add("web3SigningCredential");
    openapiRequiredFields.add("invocationType");
    openapiRequiredFields.add("methodName");
    openapiRequiredFields.add("params");
    openapiRequiredFields.add("keychainId");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to InvokeContractV1Request
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!InvokeContractV1Request.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in InvokeContractV1Request is not found in the empty JSON string", InvokeContractV1Request.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!InvokeContractV1Request.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `InvokeContractV1Request` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : InvokeContractV1Request.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("contractName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `contractName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("contractName").toString()));
      }
      // validate the required field `web3SigningCredential`
      Web3SigningCredential.validateJsonObject(jsonObj.getAsJsonObject("web3SigningCredential"));
      if (!jsonObj.get("methodName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `methodName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("methodName").toString()));
      }
      // ensure the required json array is present
      if (jsonObj.get("params") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("params").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `params` to be an array in the JSON string but got `%s`", jsonObj.get("params").toString()));
      }
      // validate the optional field `value`
      if (jsonObj.get("value") != null && !jsonObj.get("value").isJsonNull()) {
        QuorumTransactionConfigFrom.validateJsonObject(jsonObj.getAsJsonObject("value"));
      }
      // validate the optional field `gas`
      if (jsonObj.get("gas") != null && !jsonObj.get("gas").isJsonNull()) {
        QuorumTransactionConfigFrom.validateJsonObject(jsonObj.getAsJsonObject("gas"));
      }
      // validate the optional field `gasPrice`
      if (jsonObj.get("gasPrice") != null && !jsonObj.get("gasPrice").isJsonNull()) {
        QuorumTransactionConfigFrom.validateJsonObject(jsonObj.getAsJsonObject("gasPrice"));
      }
      if (!jsonObj.get("keychainId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `keychainId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("keychainId").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!InvokeContractV1Request.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'InvokeContractV1Request' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<InvokeContractV1Request> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(InvokeContractV1Request.class));

       return (TypeAdapter<T>) new TypeAdapter<InvokeContractV1Request>() {
           @Override
           public void write(JsonWriter out, InvokeContractV1Request value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public InvokeContractV1Request read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of InvokeContractV1Request given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of InvokeContractV1Request
  * @throws IOException if the JSON string is invalid with respect to InvokeContractV1Request
  */
  public static InvokeContractV1Request fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, InvokeContractV1Request.class);
  }

 /**
  * Convert an instance of InvokeContractV1Request to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

