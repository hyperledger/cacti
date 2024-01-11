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
 * CactiTransactionV1AllOf
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class CactiTransactionV1AllOf {
  public static final String SERIALIZED_NAME_PAYLOAD_DECODED = "payload_decoded";
  @SerializedName(SERIALIZED_NAME_PAYLOAD_DECODED)
  private Object payloadDecoded = null;

  public CactiTransactionV1AllOf() {
  }

  public CactiTransactionV1AllOf payloadDecoded(Object payloadDecoded) {
    
    this.payloadDecoded = payloadDecoded;
    return this;
  }

   /**
   * Decoded payload of sawtooth transaction.
   * @return payloadDecoded
  **/
  @javax.annotation.Nullable

  public Object getPayloadDecoded() {
    return payloadDecoded;
  }


  public void setPayloadDecoded(Object payloadDecoded) {
    this.payloadDecoded = payloadDecoded;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CactiTransactionV1AllOf cactiTransactionV1AllOf = (CactiTransactionV1AllOf) o;
    return Objects.equals(this.payloadDecoded, cactiTransactionV1AllOf.payloadDecoded);
  }

  @Override
  public int hashCode() {
    return Objects.hash(payloadDecoded);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CactiTransactionV1AllOf {\n");
    sb.append("    payloadDecoded: ").append(toIndentedString(payloadDecoded)).append("\n");
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
    openapiFields.add("payload_decoded");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("payload_decoded");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to CactiTransactionV1AllOf
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!CactiTransactionV1AllOf.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in CactiTransactionV1AllOf is not found in the empty JSON string", CactiTransactionV1AllOf.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!CactiTransactionV1AllOf.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `CactiTransactionV1AllOf` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : CactiTransactionV1AllOf.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!CactiTransactionV1AllOf.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'CactiTransactionV1AllOf' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<CactiTransactionV1AllOf> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(CactiTransactionV1AllOf.class));

       return (TypeAdapter<T>) new TypeAdapter<CactiTransactionV1AllOf>() {
           @Override
           public void write(JsonWriter out, CactiTransactionV1AllOf value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public CactiTransactionV1AllOf read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of CactiTransactionV1AllOf given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CactiTransactionV1AllOf
  * @throws IOException if the JSON string is invalid with respect to CactiTransactionV1AllOf
  */
  public static CactiTransactionV1AllOf fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CactiTransactionV1AllOf.class);
  }

 /**
  * Convert an instance of CactiTransactionV1AllOf to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

