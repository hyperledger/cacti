/*
 * Hyperledger Cacti Plugin - Connector Sawtooth
 * Can perform basic tasks on a Sawtooth ledger
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
import org.openapitools.client.model.WatchBlocksV1ListenerType;
import org.openapitools.client.model.WatchBlocksV1TransactionFilter;

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

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * WatchBlocksV1Options
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class WatchBlocksV1Options {
  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private WatchBlocksV1ListenerType type;

  public static final String SERIALIZED_NAME_TX_FILTER_BY = "txFilterBy";
  @SerializedName(SERIALIZED_NAME_TX_FILTER_BY)
  private WatchBlocksV1TransactionFilter txFilterBy;

  public WatchBlocksV1Options() {
  }

  public WatchBlocksV1Options type(WatchBlocksV1ListenerType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable

  public WatchBlocksV1ListenerType getType() {
    return type;
  }


  public void setType(WatchBlocksV1ListenerType type) {
    this.type = type;
  }


  public WatchBlocksV1Options txFilterBy(WatchBlocksV1TransactionFilter txFilterBy) {
    
    this.txFilterBy = txFilterBy;
    return this;
  }

   /**
   * Get txFilterBy
   * @return txFilterBy
  **/
  @javax.annotation.Nullable

  public WatchBlocksV1TransactionFilter getTxFilterBy() {
    return txFilterBy;
  }


  public void setTxFilterBy(WatchBlocksV1TransactionFilter txFilterBy) {
    this.txFilterBy = txFilterBy;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WatchBlocksV1Options watchBlocksV1Options = (WatchBlocksV1Options) o;
    return Objects.equals(this.type, watchBlocksV1Options.type) &&
        Objects.equals(this.txFilterBy, watchBlocksV1Options.txFilterBy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, txFilterBy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WatchBlocksV1Options {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    txFilterBy: ").append(toIndentedString(txFilterBy)).append("\n");
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
    openapiFields.add("type");
    openapiFields.add("txFilterBy");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to WatchBlocksV1Options
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!WatchBlocksV1Options.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in WatchBlocksV1Options is not found in the empty JSON string", WatchBlocksV1Options.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!WatchBlocksV1Options.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `WatchBlocksV1Options` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `txFilterBy`
      if (jsonObj.get("txFilterBy") != null && !jsonObj.get("txFilterBy").isJsonNull()) {
        WatchBlocksV1TransactionFilter.validateJsonObject(jsonObj.getAsJsonObject("txFilterBy"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!WatchBlocksV1Options.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'WatchBlocksV1Options' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<WatchBlocksV1Options> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(WatchBlocksV1Options.class));

       return (TypeAdapter<T>) new TypeAdapter<WatchBlocksV1Options>() {
           @Override
           public void write(JsonWriter out, WatchBlocksV1Options value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public WatchBlocksV1Options read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of WatchBlocksV1Options given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of WatchBlocksV1Options
  * @throws IOException if the JSON string is invalid with respect to WatchBlocksV1Options
  */
  public static WatchBlocksV1Options fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, WatchBlocksV1Options.class);
  }

 /**
  * Convert an instance of WatchBlocksV1Options to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

