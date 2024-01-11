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
 * ListFlowsV1Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class ListFlowsV1Response {
  public static final String SERIALIZED_NAME_FLOW_NAMES = "flowNames";
  @SerializedName(SERIALIZED_NAME_FLOW_NAMES)
  private List<String> flowNames = null;

  public ListFlowsV1Response() {
  }

  public ListFlowsV1Response flowNames(List<String> flowNames) {
    
    this.flowNames = flowNames;
    return this;
  }

  public ListFlowsV1Response addFlowNamesItem(String flowNamesItem) {
    if (this.flowNames == null) {
      this.flowNames = null;
    }
    this.flowNames.add(flowNamesItem);
    return this;
  }

   /**
   * An array of strings storing the names of the flows as returned by the flowList Corda RPC.
   * @return flowNames
  **/
  @javax.annotation.Nonnull
  public List<String> getFlowNames() {
    return flowNames;
  }


  public void setFlowNames(List<String> flowNames) {
    this.flowNames = flowNames;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ListFlowsV1Response listFlowsV1Response = (ListFlowsV1Response) o;
    return Objects.equals(this.flowNames, listFlowsV1Response.flowNames);
  }

  @Override
  public int hashCode() {
    return Objects.hash(flowNames);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ListFlowsV1Response {\n");
    sb.append("    flowNames: ").append(toIndentedString(flowNames)).append("\n");
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
    openapiFields.add("flowNames");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("flowNames");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to ListFlowsV1Response
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!ListFlowsV1Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in ListFlowsV1Response is not found in the empty JSON string", ListFlowsV1Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!ListFlowsV1Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `ListFlowsV1Response` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : ListFlowsV1Response.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      // ensure the required json array is present
      if (jsonObj.get("flowNames") == null) {
        throw new IllegalArgumentException("Expected the field `linkedContent` to be an array in the JSON string but got `null`");
      } else if (!jsonObj.get("flowNames").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `flowNames` to be an array in the JSON string but got `%s`", jsonObj.get("flowNames").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!ListFlowsV1Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'ListFlowsV1Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<ListFlowsV1Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(ListFlowsV1Response.class));

       return (TypeAdapter<T>) new TypeAdapter<ListFlowsV1Response>() {
           @Override
           public void write(JsonWriter out, ListFlowsV1Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public ListFlowsV1Response read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of ListFlowsV1Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of ListFlowsV1Response
  * @throws IOException if the JSON string is invalid with respect to ListFlowsV1Response
  */
  public static ListFlowsV1Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, ListFlowsV1Response.class);
  }

 /**
  * Convert an instance of ListFlowsV1Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

