/*
 * Hyperledger Cactus Plugin - Connector Iroha
 * Can perform basic tasks on a Iroha ledger
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
 * The list of arguments to pass in to the transaction request to Create Asset.
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class CreateAssetRequestParameters {
  public static final String SERIALIZED_NAME_ASSET_NAME = "assetName";
  @SerializedName(SERIALIZED_NAME_ASSET_NAME)
  private String assetName;

  public static final String SERIALIZED_NAME_DOMAIN_ID = "domainId";
  @SerializedName(SERIALIZED_NAME_DOMAIN_ID)
  private String domainId;

  public static final String SERIALIZED_NAME_PRECISION = "precision";
  @SerializedName(SERIALIZED_NAME_PRECISION)
  private Integer precision;

  public CreateAssetRequestParameters() {
  }

  public CreateAssetRequestParameters assetName(String assetName) {
    
    this.assetName = assetName;
    return this;
  }

   /**
   * Get assetName
   * @return assetName
  **/
  @javax.annotation.Nonnull
  public String getAssetName() {
    return assetName;
  }


  public void setAssetName(String assetName) {
    this.assetName = assetName;
  }


  public CreateAssetRequestParameters domainId(String domainId) {
    
    this.domainId = domainId;
    return this;
  }

   /**
   * Get domainId
   * @return domainId
  **/
  @javax.annotation.Nonnull
  public String getDomainId() {
    return domainId;
  }


  public void setDomainId(String domainId) {
    this.domainId = domainId;
  }


  public CreateAssetRequestParameters precision(Integer precision) {
    
    this.precision = precision;
    return this;
  }

   /**
   * Get precision
   * @return precision
  **/
  @javax.annotation.Nonnull
  public Integer getPrecision() {
    return precision;
  }


  public void setPrecision(Integer precision) {
    this.precision = precision;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateAssetRequestParameters createAssetRequestParameters = (CreateAssetRequestParameters) o;
    return Objects.equals(this.assetName, createAssetRequestParameters.assetName) &&
        Objects.equals(this.domainId, createAssetRequestParameters.domainId) &&
        Objects.equals(this.precision, createAssetRequestParameters.precision);
  }

  @Override
  public int hashCode() {
    return Objects.hash(assetName, domainId, precision);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateAssetRequestParameters {\n");
    sb.append("    assetName: ").append(toIndentedString(assetName)).append("\n");
    sb.append("    domainId: ").append(toIndentedString(domainId)).append("\n");
    sb.append("    precision: ").append(toIndentedString(precision)).append("\n");
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
    openapiFields.add("assetName");
    openapiFields.add("domainId");
    openapiFields.add("precision");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("assetName");
    openapiRequiredFields.add("domainId");
    openapiRequiredFields.add("precision");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to CreateAssetRequestParameters
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!CreateAssetRequestParameters.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in CreateAssetRequestParameters is not found in the empty JSON string", CreateAssetRequestParameters.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!CreateAssetRequestParameters.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `CreateAssetRequestParameters` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : CreateAssetRequestParameters.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("assetName").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `assetName` to be a primitive type in the JSON string but got `%s`", jsonObj.get("assetName").toString()));
      }
      if (!jsonObj.get("domainId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `domainId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("domainId").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!CreateAssetRequestParameters.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'CreateAssetRequestParameters' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<CreateAssetRequestParameters> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(CreateAssetRequestParameters.class));

       return (TypeAdapter<T>) new TypeAdapter<CreateAssetRequestParameters>() {
           @Override
           public void write(JsonWriter out, CreateAssetRequestParameters value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public CreateAssetRequestParameters read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of CreateAssetRequestParameters given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CreateAssetRequestParameters
  * @throws IOException if the JSON string is invalid with respect to CreateAssetRequestParameters
  */
  public static CreateAssetRequestParameters fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CreateAssetRequestParameters.class);
  }

 /**
  * Convert an instance of CreateAssetRequestParameters to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

