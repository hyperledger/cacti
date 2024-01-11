/*
 * Hyperledger Core API
 * Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..
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
 * HasKeychainEntryResponseV1
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class HasKeychainEntryResponseV1 {
  public static final String SERIALIZED_NAME_KEY = "key";
  @SerializedName(SERIALIZED_NAME_KEY)
  private String key;

  public static final String SERIALIZED_NAME_CHECKED_AT = "checkedAt";
  @SerializedName(SERIALIZED_NAME_CHECKED_AT)
  private String checkedAt;

  public static final String SERIALIZED_NAME_IS_PRESENT = "isPresent";
  @SerializedName(SERIALIZED_NAME_IS_PRESENT)
  private Boolean isPresent;

  public HasKeychainEntryResponseV1() {
  }

  public HasKeychainEntryResponseV1 key(String key) {
    
    this.key = key;
    return this;
  }

   /**
   * The key that was used to check the presence of the value in the entry store.
   * @return key
  **/
  @javax.annotation.Nonnull
  public String getKey() {
    return key;
  }


  public void setKey(String key) {
    this.key = key;
  }


  public HasKeychainEntryResponseV1 checkedAt(String checkedAt) {
    
    this.checkedAt = checkedAt;
    return this;
  }

   /**
   * Date and time encoded as JSON when the presence check was performed by the plugin backend.
   * @return checkedAt
  **/
  @javax.annotation.Nonnull
  public String getCheckedAt() {
    return checkedAt;
  }


  public void setCheckedAt(String checkedAt) {
    this.checkedAt = checkedAt;
  }


  public HasKeychainEntryResponseV1 isPresent(Boolean isPresent) {
    
    this.isPresent = isPresent;
    return this;
  }

   /**
   * The boolean true or false indicating the presence or absence of an entry under &#39;key&#39;.
   * @return isPresent
  **/
  @javax.annotation.Nonnull
  public Boolean getIsPresent() {
    return isPresent;
  }


  public void setIsPresent(Boolean isPresent) {
    this.isPresent = isPresent;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HasKeychainEntryResponseV1 hasKeychainEntryResponseV1 = (HasKeychainEntryResponseV1) o;
    return Objects.equals(this.key, hasKeychainEntryResponseV1.key) &&
        Objects.equals(this.checkedAt, hasKeychainEntryResponseV1.checkedAt) &&
        Objects.equals(this.isPresent, hasKeychainEntryResponseV1.isPresent);
  }

  @Override
  public int hashCode() {
    return Objects.hash(key, checkedAt, isPresent);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HasKeychainEntryResponseV1 {\n");
    sb.append("    key: ").append(toIndentedString(key)).append("\n");
    sb.append("    checkedAt: ").append(toIndentedString(checkedAt)).append("\n");
    sb.append("    isPresent: ").append(toIndentedString(isPresent)).append("\n");
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
    openapiFields.add("key");
    openapiFields.add("checkedAt");
    openapiFields.add("isPresent");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("key");
    openapiRequiredFields.add("checkedAt");
    openapiRequiredFields.add("isPresent");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to HasKeychainEntryResponseV1
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!HasKeychainEntryResponseV1.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in HasKeychainEntryResponseV1 is not found in the empty JSON string", HasKeychainEntryResponseV1.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!HasKeychainEntryResponseV1.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `HasKeychainEntryResponseV1` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : HasKeychainEntryResponseV1.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("key").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `key` to be a primitive type in the JSON string but got `%s`", jsonObj.get("key").toString()));
      }
      if (!jsonObj.get("checkedAt").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `checkedAt` to be a primitive type in the JSON string but got `%s`", jsonObj.get("checkedAt").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!HasKeychainEntryResponseV1.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'HasKeychainEntryResponseV1' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<HasKeychainEntryResponseV1> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(HasKeychainEntryResponseV1.class));

       return (TypeAdapter<T>) new TypeAdapter<HasKeychainEntryResponseV1>() {
           @Override
           public void write(JsonWriter out, HasKeychainEntryResponseV1 value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public HasKeychainEntryResponseV1 read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of HasKeychainEntryResponseV1 given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of HasKeychainEntryResponseV1
  * @throws IOException if the JSON string is invalid with respect to HasKeychainEntryResponseV1
  */
  public static HasKeychainEntryResponseV1 fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, HasKeychainEntryResponseV1.class);
  }

 /**
  * Convert an instance of HasKeychainEntryResponseV1 to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

