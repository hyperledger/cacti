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
import org.openapitools.client.model.QuorumTransactionConfigTo;

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
 * QuorumTransactionConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class QuorumTransactionConfig {
  public static final String SERIALIZED_NAME_RAW_TRANSACTION = "rawTransaction";
  @SerializedName(SERIALIZED_NAME_RAW_TRANSACTION)
  private String rawTransaction;

  public static final String SERIALIZED_NAME_FROM = "from";
  @SerializedName(SERIALIZED_NAME_FROM)
  private QuorumTransactionConfigFrom from;

  public static final String SERIALIZED_NAME_TO = "to";
  @SerializedName(SERIALIZED_NAME_TO)
  private QuorumTransactionConfigTo to;

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

  public static final String SERIALIZED_NAME_DATA = "data";
  @SerializedName(SERIALIZED_NAME_DATA)
  private QuorumTransactionConfigTo data;

  public QuorumTransactionConfig() {
  }

  public QuorumTransactionConfig rawTransaction(String rawTransaction) {
    
    this.rawTransaction = rawTransaction;
    return this;
  }

   /**
   * Get rawTransaction
   * @return rawTransaction
  **/
  @javax.annotation.Nullable
  public String getRawTransaction() {
    return rawTransaction;
  }


  public void setRawTransaction(String rawTransaction) {
    this.rawTransaction = rawTransaction;
  }


  public QuorumTransactionConfig from(QuorumTransactionConfigFrom from) {
    
    this.from = from;
    return this;
  }

   /**
   * Get from
   * @return from
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigFrom getFrom() {
    return from;
  }


  public void setFrom(QuorumTransactionConfigFrom from) {
    this.from = from;
  }


  public QuorumTransactionConfig to(QuorumTransactionConfigTo to) {
    
    this.to = to;
    return this;
  }

   /**
   * Get to
   * @return to
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigTo getTo() {
    return to;
  }


  public void setTo(QuorumTransactionConfigTo to) {
    this.to = to;
  }


  public QuorumTransactionConfig value(QuorumTransactionConfigFrom value) {
    
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


  public QuorumTransactionConfig gas(QuorumTransactionConfigFrom gas) {
    
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


  public QuorumTransactionConfig gasPrice(QuorumTransactionConfigFrom gasPrice) {
    
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


  public QuorumTransactionConfig nonce(BigDecimal nonce) {
    
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


  public QuorumTransactionConfig data(QuorumTransactionConfigTo data) {
    
    this.data = data;
    return this;
  }

   /**
   * Get data
   * @return data
  **/
  @javax.annotation.Nullable
  public QuorumTransactionConfigTo getData() {
    return data;
  }


  public void setData(QuorumTransactionConfigTo data) {
    this.data = data;
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
   * @return the QuorumTransactionConfig instance itself
   */
  public QuorumTransactionConfig putAdditionalProperty(String key, Object value) {
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
    QuorumTransactionConfig quorumTransactionConfig = (QuorumTransactionConfig) o;
    return Objects.equals(this.rawTransaction, quorumTransactionConfig.rawTransaction) &&
        Objects.equals(this.from, quorumTransactionConfig.from) &&
        Objects.equals(this.to, quorumTransactionConfig.to) &&
        Objects.equals(this.value, quorumTransactionConfig.value) &&
        Objects.equals(this.gas, quorumTransactionConfig.gas) &&
        Objects.equals(this.gasPrice, quorumTransactionConfig.gasPrice) &&
        Objects.equals(this.nonce, quorumTransactionConfig.nonce) &&
        Objects.equals(this.data, quorumTransactionConfig.data)&&
        Objects.equals(this.additionalProperties, quorumTransactionConfig.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(rawTransaction, from, to, value, gas, gasPrice, nonce, data, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class QuorumTransactionConfig {\n");
    sb.append("    rawTransaction: ").append(toIndentedString(rawTransaction)).append("\n");
    sb.append("    from: ").append(toIndentedString(from)).append("\n");
    sb.append("    to: ").append(toIndentedString(to)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    gas: ").append(toIndentedString(gas)).append("\n");
    sb.append("    gasPrice: ").append(toIndentedString(gasPrice)).append("\n");
    sb.append("    nonce: ").append(toIndentedString(nonce)).append("\n");
    sb.append("    data: ").append(toIndentedString(data)).append("\n");
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
    openapiFields.add("rawTransaction");
    openapiFields.add("from");
    openapiFields.add("to");
    openapiFields.add("value");
    openapiFields.add("gas");
    openapiFields.add("gasPrice");
    openapiFields.add("nonce");
    openapiFields.add("data");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to QuorumTransactionConfig
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!QuorumTransactionConfig.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in QuorumTransactionConfig is not found in the empty JSON string", QuorumTransactionConfig.openapiRequiredFields.toString()));
        }
      }
      if ((jsonObj.get("rawTransaction") != null && !jsonObj.get("rawTransaction").isJsonNull()) && !jsonObj.get("rawTransaction").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `rawTransaction` to be a primitive type in the JSON string but got `%s`", jsonObj.get("rawTransaction").toString()));
      }
      // validate the optional field `from`
      if (jsonObj.get("from") != null && !jsonObj.get("from").isJsonNull()) {
        QuorumTransactionConfigFrom.validateJsonObject(jsonObj.getAsJsonObject("from"));
      }
      // validate the optional field `to`
      if (jsonObj.get("to") != null && !jsonObj.get("to").isJsonNull()) {
        QuorumTransactionConfigTo.validateJsonObject(jsonObj.getAsJsonObject("to"));
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
      // validate the optional field `data`
      if (jsonObj.get("data") != null && !jsonObj.get("data").isJsonNull()) {
        QuorumTransactionConfigTo.validateJsonObject(jsonObj.getAsJsonObject("data"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!QuorumTransactionConfig.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'QuorumTransactionConfig' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<QuorumTransactionConfig> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(QuorumTransactionConfig.class));

       return (TypeAdapter<T>) new TypeAdapter<QuorumTransactionConfig>() {
           @Override
           public void write(JsonWriter out, QuorumTransactionConfig value) throws IOException {
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
           public QuorumTransactionConfig read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             QuorumTransactionConfig instance = thisAdapter.fromJsonTree(jsonObj);
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
  * Create an instance of QuorumTransactionConfig given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of QuorumTransactionConfig
  * @throws IOException if the JSON string is invalid with respect to QuorumTransactionConfig
  */
  public static QuorumTransactionConfig fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, QuorumTransactionConfig.class);
  }

 /**
  * Convert an instance of QuorumTransactionConfig to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

