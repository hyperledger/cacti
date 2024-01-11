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
import org.openapitools.client.model.WatchBlocksV1BlockData;
import org.openapitools.client.model.Web3BlockHeader;

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
 * WatchBlocksV1Progress
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class WatchBlocksV1Progress {
  public static final String SERIALIZED_NAME_BLOCK_HEADER = "blockHeader";
  @SerializedName(SERIALIZED_NAME_BLOCK_HEADER)
  private Web3BlockHeader blockHeader;

  public static final String SERIALIZED_NAME_BLOCK_DATA = "blockData";
  @SerializedName(SERIALIZED_NAME_BLOCK_DATA)
  private WatchBlocksV1BlockData blockData;

  public WatchBlocksV1Progress() {
  }

  public WatchBlocksV1Progress blockHeader(Web3BlockHeader blockHeader) {
    
    this.blockHeader = blockHeader;
    return this;
  }

   /**
   * Get blockHeader
   * @return blockHeader
  **/
  @javax.annotation.Nullable
  public Web3BlockHeader getBlockHeader() {
    return blockHeader;
  }


  public void setBlockHeader(Web3BlockHeader blockHeader) {
    this.blockHeader = blockHeader;
  }


  public WatchBlocksV1Progress blockData(WatchBlocksV1BlockData blockData) {
    
    this.blockData = blockData;
    return this;
  }

   /**
   * Get blockData
   * @return blockData
  **/
  @javax.annotation.Nullable
  public WatchBlocksV1BlockData getBlockData() {
    return blockData;
  }


  public void setBlockData(WatchBlocksV1BlockData blockData) {
    this.blockData = blockData;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WatchBlocksV1Progress watchBlocksV1Progress = (WatchBlocksV1Progress) o;
    return Objects.equals(this.blockHeader, watchBlocksV1Progress.blockHeader) &&
        Objects.equals(this.blockData, watchBlocksV1Progress.blockData);
  }

  @Override
  public int hashCode() {
    return Objects.hash(blockHeader, blockData);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WatchBlocksV1Progress {\n");
    sb.append("    blockHeader: ").append(toIndentedString(blockHeader)).append("\n");
    sb.append("    blockData: ").append(toIndentedString(blockData)).append("\n");
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
    openapiFields.add("blockHeader");
    openapiFields.add("blockData");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to WatchBlocksV1Progress
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!WatchBlocksV1Progress.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in WatchBlocksV1Progress is not found in the empty JSON string", WatchBlocksV1Progress.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!WatchBlocksV1Progress.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `WatchBlocksV1Progress` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `blockHeader`
      if (jsonObj.get("blockHeader") != null && !jsonObj.get("blockHeader").isJsonNull()) {
        Web3BlockHeader.validateJsonObject(jsonObj.getAsJsonObject("blockHeader"));
      }
      // validate the optional field `blockData`
      if (jsonObj.get("blockData") != null && !jsonObj.get("blockData").isJsonNull()) {
        WatchBlocksV1BlockData.validateJsonObject(jsonObj.getAsJsonObject("blockData"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!WatchBlocksV1Progress.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'WatchBlocksV1Progress' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<WatchBlocksV1Progress> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(WatchBlocksV1Progress.class));

       return (TypeAdapter<T>) new TypeAdapter<WatchBlocksV1Progress>() {
           @Override
           public void write(JsonWriter out, WatchBlocksV1Progress value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public WatchBlocksV1Progress read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of WatchBlocksV1Progress given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of WatchBlocksV1Progress
  * @throws IOException if the JSON string is invalid with respect to WatchBlocksV1Progress
  */
  public static WatchBlocksV1Progress fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, WatchBlocksV1Progress.class);
  }

 /**
  * Convert an instance of WatchBlocksV1Progress to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

