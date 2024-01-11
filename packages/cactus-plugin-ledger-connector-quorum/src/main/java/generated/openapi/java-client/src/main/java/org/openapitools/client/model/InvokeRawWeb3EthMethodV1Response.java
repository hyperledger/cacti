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
import org.openapitools.jackson.nullable.JsonNullable;

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
 * InvokeRawWeb3EthMethodV1Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class InvokeRawWeb3EthMethodV1Response {
  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private BigDecimal status;

  public static final String SERIALIZED_NAME_DATA = "data";
  @SerializedName(SERIALIZED_NAME_DATA)
  private Object data = null;

  public static final String SERIALIZED_NAME_ERROR_DETAIL = "errorDetail";
  @SerializedName(SERIALIZED_NAME_ERROR_DETAIL)
  private String errorDetail;

  public InvokeRawWeb3EthMethodV1Response() {
  }

  public InvokeRawWeb3EthMethodV1Response status(BigDecimal status) {
    
    this.status = status;
    return this;
  }

   /**
   * Status code of the operation
   * @return status
  **/
  @javax.annotation.Nonnull
  public BigDecimal getStatus() {
    return status;
  }


  public void setStatus(BigDecimal status) {
    this.status = status;
  }


  public InvokeRawWeb3EthMethodV1Response data(Object data) {
    
    this.data = data;
    return this;
  }

   /**
   * Output of requested web3.eth method
   * @return data
  **/
  @javax.annotation.Nullable
  public Object getData() {
    return data;
  }


  public void setData(Object data) {
    this.data = data;
  }


  public InvokeRawWeb3EthMethodV1Response errorDetail(String errorDetail) {
    
    this.errorDetail = errorDetail;
    return this;
  }

   /**
   * Error details
   * @return errorDetail
  **/
  @javax.annotation.Nullable
  public String getErrorDetail() {
    return errorDetail;
  }


  public void setErrorDetail(String errorDetail) {
    this.errorDetail = errorDetail;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InvokeRawWeb3EthMethodV1Response invokeRawWeb3EthMethodV1Response = (InvokeRawWeb3EthMethodV1Response) o;
    return Objects.equals(this.status, invokeRawWeb3EthMethodV1Response.status) &&
        Objects.equals(this.data, invokeRawWeb3EthMethodV1Response.data) &&
        Objects.equals(this.errorDetail, invokeRawWeb3EthMethodV1Response.errorDetail);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(status, data, errorDetail);
  }

  private static <T> int hashCodeNullable(JsonNullable<T> a) {
    if (a == null) {
      return 1;
    }
    return a.isPresent() ? Arrays.deepHashCode(new Object[]{a.get()}) : 31;
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InvokeRawWeb3EthMethodV1Response {\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    data: ").append(toIndentedString(data)).append("\n");
    sb.append("    errorDetail: ").append(toIndentedString(errorDetail)).append("\n");
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
    openapiFields.add("status");
    openapiFields.add("data");
    openapiFields.add("errorDetail");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("status");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to InvokeRawWeb3EthMethodV1Response
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!InvokeRawWeb3EthMethodV1Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in InvokeRawWeb3EthMethodV1Response is not found in the empty JSON string", InvokeRawWeb3EthMethodV1Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!InvokeRawWeb3EthMethodV1Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `InvokeRawWeb3EthMethodV1Response` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : InvokeRawWeb3EthMethodV1Response.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if ((jsonObj.get("errorDetail") != null && !jsonObj.get("errorDetail").isJsonNull()) && !jsonObj.get("errorDetail").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `errorDetail` to be a primitive type in the JSON string but got `%s`", jsonObj.get("errorDetail").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!InvokeRawWeb3EthMethodV1Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'InvokeRawWeb3EthMethodV1Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<InvokeRawWeb3EthMethodV1Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(InvokeRawWeb3EthMethodV1Response.class));

       return (TypeAdapter<T>) new TypeAdapter<InvokeRawWeb3EthMethodV1Response>() {
           @Override
           public void write(JsonWriter out, InvokeRawWeb3EthMethodV1Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public InvokeRawWeb3EthMethodV1Response read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of InvokeRawWeb3EthMethodV1Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of InvokeRawWeb3EthMethodV1Response
  * @throws IOException if the JSON string is invalid with respect to InvokeRawWeb3EthMethodV1Response
  */
  public static InvokeRawWeb3EthMethodV1Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, InvokeRawWeb3EthMethodV1Response.class);
  }

 /**
  * Convert an instance of InvokeRawWeb3EthMethodV1Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

