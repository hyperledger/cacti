/*
 * Hyperledger Cactus API
 * Interact with a Cactus deployment through HTTP.
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
 * MemoryUsage
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class MemoryUsage {
  public static final String SERIALIZED_NAME_RSS = "rss";
  @SerializedName(SERIALIZED_NAME_RSS)
  private BigDecimal rss;

  public static final String SERIALIZED_NAME_HEAP_TOTAL = "heapTotal";
  @SerializedName(SERIALIZED_NAME_HEAP_TOTAL)
  private BigDecimal heapTotal;

  public static final String SERIALIZED_NAME_HEAP_USED = "heapUsed";
  @SerializedName(SERIALIZED_NAME_HEAP_USED)
  private BigDecimal heapUsed;

  public static final String SERIALIZED_NAME_EXTERNAL = "external";
  @SerializedName(SERIALIZED_NAME_EXTERNAL)
  private BigDecimal external;

  public static final String SERIALIZED_NAME_ARRAY_BUFFERS = "arrayBuffers";
  @SerializedName(SERIALIZED_NAME_ARRAY_BUFFERS)
  private BigDecimal arrayBuffers;

  public MemoryUsage() {
  }

  public MemoryUsage rss(BigDecimal rss) {
    
    this.rss = rss;
    return this;
  }

   /**
   * Get rss
   * @return rss
  **/
  @javax.annotation.Nullable
  public BigDecimal getRss() {
    return rss;
  }


  public void setRss(BigDecimal rss) {
    this.rss = rss;
  }


  public MemoryUsage heapTotal(BigDecimal heapTotal) {
    
    this.heapTotal = heapTotal;
    return this;
  }

   /**
   * Get heapTotal
   * @return heapTotal
  **/
  @javax.annotation.Nullable
  public BigDecimal getHeapTotal() {
    return heapTotal;
  }


  public void setHeapTotal(BigDecimal heapTotal) {
    this.heapTotal = heapTotal;
  }


  public MemoryUsage heapUsed(BigDecimal heapUsed) {
    
    this.heapUsed = heapUsed;
    return this;
  }

   /**
   * Get heapUsed
   * @return heapUsed
  **/
  @javax.annotation.Nullable
  public BigDecimal getHeapUsed() {
    return heapUsed;
  }


  public void setHeapUsed(BigDecimal heapUsed) {
    this.heapUsed = heapUsed;
  }


  public MemoryUsage external(BigDecimal external) {
    
    this.external = external;
    return this;
  }

   /**
   * Get external
   * @return external
  **/
  @javax.annotation.Nullable
  public BigDecimal getExternal() {
    return external;
  }


  public void setExternal(BigDecimal external) {
    this.external = external;
  }


  public MemoryUsage arrayBuffers(BigDecimal arrayBuffers) {
    
    this.arrayBuffers = arrayBuffers;
    return this;
  }

   /**
   * Get arrayBuffers
   * @return arrayBuffers
  **/
  @javax.annotation.Nullable
  public BigDecimal getArrayBuffers() {
    return arrayBuffers;
  }


  public void setArrayBuffers(BigDecimal arrayBuffers) {
    this.arrayBuffers = arrayBuffers;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    MemoryUsage memoryUsage = (MemoryUsage) o;
    return Objects.equals(this.rss, memoryUsage.rss) &&
        Objects.equals(this.heapTotal, memoryUsage.heapTotal) &&
        Objects.equals(this.heapUsed, memoryUsage.heapUsed) &&
        Objects.equals(this.external, memoryUsage.external) &&
        Objects.equals(this.arrayBuffers, memoryUsage.arrayBuffers);
  }

  @Override
  public int hashCode() {
    return Objects.hash(rss, heapTotal, heapUsed, external, arrayBuffers);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class MemoryUsage {\n");
    sb.append("    rss: ").append(toIndentedString(rss)).append("\n");
    sb.append("    heapTotal: ").append(toIndentedString(heapTotal)).append("\n");
    sb.append("    heapUsed: ").append(toIndentedString(heapUsed)).append("\n");
    sb.append("    external: ").append(toIndentedString(external)).append("\n");
    sb.append("    arrayBuffers: ").append(toIndentedString(arrayBuffers)).append("\n");
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
    openapiFields.add("rss");
    openapiFields.add("heapTotal");
    openapiFields.add("heapUsed");
    openapiFields.add("external");
    openapiFields.add("arrayBuffers");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to MemoryUsage
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!MemoryUsage.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in MemoryUsage is not found in the empty JSON string", MemoryUsage.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!MemoryUsage.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `MemoryUsage` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!MemoryUsage.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'MemoryUsage' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<MemoryUsage> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(MemoryUsage.class));

       return (TypeAdapter<T>) new TypeAdapter<MemoryUsage>() {
           @Override
           public void write(JsonWriter out, MemoryUsage value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public MemoryUsage read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of MemoryUsage given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of MemoryUsage
  * @throws IOException if the JSON string is invalid with respect to MemoryUsage
  */
  public static MemoryUsage fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, MemoryUsage.class);
  }

 /**
  * Convert an instance of MemoryUsage to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

