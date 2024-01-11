/*
 * Hyperledger Cactus Plugin - Connector Corda
 * Can perform basic tasks on a Corda ledger
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
import java.util.ArrayList;
import java.util.List;

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
 * ClearMonitorTransactionsV1Request
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ClearMonitorTransactionsV1Request {
  public static final String SERIALIZED_NAME_CLIENT_APP_ID = "clientAppId";
  @SerializedName(SERIALIZED_NAME_CLIENT_APP_ID)
  private String clientAppId;

  public static final String SERIALIZED_NAME_STATE_FULL_CLASS_NAME = "stateFullClassName";
  @SerializedName(SERIALIZED_NAME_STATE_FULL_CLASS_NAME)
  private String stateFullClassName;

  public static final String SERIALIZED_NAME_TX_INDEXES = "txIndexes";
  @SerializedName(SERIALIZED_NAME_TX_INDEXES)
  private List<String> txIndexes = null;

  public ClearMonitorTransactionsV1Request() {
  }

  public ClearMonitorTransactionsV1Request clientAppId(String clientAppId) {
    
    this.clientAppId = clientAppId;
    return this;
  }

   /**
   * ID of a client application that wants to monitor the state changes
   * @return clientAppId
  **/
  @javax.annotation.Nonnull
  public String getClientAppId() {
    return clientAppId;
  }


  public void setClientAppId(String clientAppId) {
    this.clientAppId = clientAppId;
  }


  public ClearMonitorTransactionsV1Request stateFullClassName(String stateFullClassName) {
    
    this.stateFullClassName = stateFullClassName;
    return this;
  }

   /**
   * The fully qualified name of the Corda state to monitor
   * @return stateFullClassName
  **/
  @javax.annotation.Nonnull
  public String getStateFullClassName() {
    return stateFullClassName;
  }


  public void setStateFullClassName(String stateFullClassName) {
    this.stateFullClassName = stateFullClassName;
  }


  public ClearMonitorTransactionsV1Request txIndexes(List<String> txIndexes) {
    
    this.txIndexes = txIndexes;
    return this;
  }

  public ClearMonitorTransactionsV1Request addTxIndexesItem(String txIndexesItem) {
    if (this.txIndexes == null) {
      this.txIndexes = null;
    }
    this.txIndexes.add(txIndexesItem);
    return this;
  }

   /**
   * Get txIndexes
   * @return txIndexes
  **/
  @javax.annotation.Nonnull
  public List<String> getTxIndexes() {
    return txIndexes;
  }


  public void setTxIndexes(List<String> txIndexes) {
    this.txIndexes = txIndexes;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ClearMonitorTransactionsV1Request clearMonitorTransactionsV1Request = (ClearMonitorTransactionsV1Request) o;
    return Objects.equals(this.clientAppId, clearMonitorTransactionsV1Request.clientAppId) &&
        Objects.equals(this.stateFullClassName, clearMonitorTransactionsV1Request.stateFullClassName) &&
        Objects.equals(this.txIndexes, clearMonitorTransactionsV1Request.txIndexes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(clientAppId, stateFullClassName, txIndexes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ClearMonitorTransactionsV1Request {\n");
    sb.append("    clientAppId: ").append(toIndentedString(clientAppId)).append("\n");
    sb.append("    stateFullClassName: ").append(toIndentedString(stateFullClassName)).append("\n");
    sb.append("    txIndexes: ").append(toIndentedString(txIndexes)).append("\n");
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
    openapiFields.add("clientAppId");
    openapiFields.add("stateFullClassName");
    openapiFields.add("txIndexes");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("clientAppId");
    openapiRequiredFields.add("stateFullClassName");
    openapiRequiredFields.add("txIndexes");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ClearMonitorTransactionsV1Request
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ClearMonitorTransactionsV1Request.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ClearMonitorTransactionsV1Request is not found in the empty JSON string", ClearMonitorTransactionsV1Request.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ClearMonitorTransactionsV1Request.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ClearMonitorTransactionsV1Request` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : ClearMonitorTransactionsV1Request.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("clientAppId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `clientAppId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("clientAppId").toString()));
      }
      if (!jsonObj.get("stateFullClassName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `stateFullClassName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("stateFullClassName").toString()));
      }
      // ensure the required json array is present
      if (jsonObj.get("txIndexes") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("txIndexes").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `txIndexes` to be an array in the JSON string but got `%s`", jsonObj.get("txIndexes").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ClearMonitorTransactionsV1Request.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ClearMonitorTransactionsV1Request' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ClearMonitorTransactionsV1Request> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ClearMonitorTransactionsV1Request.class));

       return (TypeAdapter<T>) new TypeAdapter<ClearMonitorTransactionsV1Request>() {
           @Override
           public void write(JsonWriter out, ClearMonitorTransactionsV1Request value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ClearMonitorTransactionsV1Request read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ClearMonitorTransactionsV1Request given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ClearMonitorTransactionsV1Request
  * @throws IOException if the JSON string is invalid with respect to ClearMonitorTransactionsV1Request
  */
  public static ClearMonitorTransactionsV1Request fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ClearMonitorTransactionsV1Request.class);
  }

 /**
  * Convert an instance of ClearMonitorTransactionsV1Request to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

