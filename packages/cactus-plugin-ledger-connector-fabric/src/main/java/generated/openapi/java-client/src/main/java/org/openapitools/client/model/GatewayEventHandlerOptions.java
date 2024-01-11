/*
 * Hyperledger Cactus Plugin - Connector Fabric
 * Can perform basic tasks on a fabric ledger
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
import org.openapitools.client.model.DefaultEventHandlerStrategy;

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
 * GatewayEventHandlerOptions
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class GatewayEventHandlerOptions {
  public static final String SERIALIZED_NAME_COMMIT_TIMEOUT = "commitTimeout";
  @SerializedName(SERIALIZED_NAME_COMMIT_TIMEOUT)
  private BigDecimal commitTimeout;

  public static final String SERIALIZED_NAME_ENDORSE_TIMEOUT = "endorseTimeout";
  @SerializedName(SERIALIZED_NAME_ENDORSE_TIMEOUT)
  private BigDecimal endorseTimeout;

  public static final String SERIALIZED_NAME_STRATEGY = "strategy";
  @SerializedName(SERIALIZED_NAME_STRATEGY)
  private DefaultEventHandlerStrategy strategy;

  public GatewayEventHandlerOptions() {
  }

  public GatewayEventHandlerOptions commitTimeout(BigDecimal commitTimeout) {
    
    this.commitTimeout = commitTimeout;
    return this;
  }

   /**
   * Get commitTimeout
   * @return commitTimeout
  **/
  @javax.annotation.Nullable
  public BigDecimal getCommitTimeout() {
    return commitTimeout;
  }


  public void setCommitTimeout(BigDecimal commitTimeout) {
    this.commitTimeout = commitTimeout;
  }


  public GatewayEventHandlerOptions endorseTimeout(BigDecimal endorseTimeout) {
    
    this.endorseTimeout = endorseTimeout;
    return this;
  }

   /**
   * Get endorseTimeout
   * @return endorseTimeout
  **/
  @javax.annotation.Nullable
  public BigDecimal getEndorseTimeout() {
    return endorseTimeout;
  }


  public void setEndorseTimeout(BigDecimal endorseTimeout) {
    this.endorseTimeout = endorseTimeout;
  }


  public GatewayEventHandlerOptions strategy(DefaultEventHandlerStrategy strategy) {
    
    this.strategy = strategy;
    return this;
  }

   /**
   * Get strategy
   * @return strategy
  **/
  @javax.annotation.Nonnull
  public DefaultEventHandlerStrategy getStrategy() {
    return strategy;
  }


  public void setStrategy(DefaultEventHandlerStrategy strategy) {
    this.strategy = strategy;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    GatewayEventHandlerOptions gatewayEventHandlerOptions = (GatewayEventHandlerOptions) o;
    return Objects.equals(this.commitTimeout, gatewayEventHandlerOptions.commitTimeout) &&
        Objects.equals(this.endorseTimeout, gatewayEventHandlerOptions.endorseTimeout) &&
        Objects.equals(this.strategy, gatewayEventHandlerOptions.strategy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(commitTimeout, endorseTimeout, strategy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class GatewayEventHandlerOptions {\n");
    sb.append("    commitTimeout: ").append(toIndentedString(commitTimeout)).append("\n");
    sb.append("    endorseTimeout: ").append(toIndentedString(endorseTimeout)).append("\n");
    sb.append("    strategy: ").append(toIndentedString(strategy)).append("\n");
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
    openapiFields.add("commitTimeout");
    openapiFields.add("endorseTimeout");
    openapiFields.add("strategy");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("strategy");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to GatewayEventHandlerOptions
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!GatewayEventHandlerOptions.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in GatewayEventHandlerOptions is not found in the empty JSON string", GatewayEventHandlerOptions.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!GatewayEventHandlerOptions.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `GatewayEventHandlerOptions` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : GatewayEventHandlerOptions.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!GatewayEventHandlerOptions.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'GatewayEventHandlerOptions' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<GatewayEventHandlerOptions> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(GatewayEventHandlerOptions.class));

       return (TypeAdapter<T>) new TypeAdapter<GatewayEventHandlerOptions>() {
           @Override
           public void write(JsonWriter out, GatewayEventHandlerOptions value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public GatewayEventHandlerOptions read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of GatewayEventHandlerOptions given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of GatewayEventHandlerOptions
  * @throws IOException if the JSON string is invalid with respect to GatewayEventHandlerOptions
  */
  public static GatewayEventHandlerOptions fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, GatewayEventHandlerOptions.class);
  }

 /**
  * Convert an instance of GatewayEventHandlerOptions to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

