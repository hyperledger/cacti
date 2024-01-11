/*
 * Hyperledger Cacti Plugin - Connector Ethereum
 * Can perform basic tasks on a Ethereum ledger
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
import org.openapitools.client.model.Web3SigningCredentialType;

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
 * Web3SigningCredentialCactiKeychainRef
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class Web3SigningCredentialCactiKeychainRef {
  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private Web3SigningCredentialType type;

  public static final String SERIALIZED_NAME_ETH_ACCOUNT = "ethAccount";
  @SerializedName(SERIALIZED_NAME_ETH_ACCOUNT)
  private String ethAccount;

  public static final String SERIALIZED_NAME_KEYCHAIN_ENTRY_KEY = "keychainEntryKey";
  @SerializedName(SERIALIZED_NAME_KEYCHAIN_ENTRY_KEY)
  private String keychainEntryKey;

  public static final String SERIALIZED_NAME_KEYCHAIN_ID = "keychainId";
  @SerializedName(SERIALIZED_NAME_KEYCHAIN_ID)
  private String keychainId;

  public Web3SigningCredentialCactiKeychainRef() {
  }

  public Web3SigningCredentialCactiKeychainRef type(Web3SigningCredentialType type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nonnull

  public Web3SigningCredentialType getType() {
    return type;
  }


  public void setType(Web3SigningCredentialType type) {
    this.type = type;
  }


  public Web3SigningCredentialCactiKeychainRef ethAccount(String ethAccount) {
    
    this.ethAccount = ethAccount;
    return this;
  }

   /**
   * The ethereum account (public key) that the credential  belongs to. Basically the username in the traditional  terminology of authentication.
   * @return ethAccount
  **/
  @javax.annotation.Nonnull

  public String getEthAccount() {
    return ethAccount;
  }


  public void setEthAccount(String ethAccount) {
    this.ethAccount = ethAccount;
  }


  public Web3SigningCredentialCactiKeychainRef keychainEntryKey(String keychainEntryKey) {
    
    this.keychainEntryKey = keychainEntryKey;
    return this;
  }

   /**
   * The key to use when looking up the the keychain entry holding the secret pointed to by the  keychainEntryKey parameter.
   * @return keychainEntryKey
  **/
  @javax.annotation.Nonnull

  public String getKeychainEntryKey() {
    return keychainEntryKey;
  }


  public void setKeychainEntryKey(String keychainEntryKey) {
    this.keychainEntryKey = keychainEntryKey;
  }


  public Web3SigningCredentialCactiKeychainRef keychainId(String keychainId) {
    
    this.keychainId = keychainId;
    return this;
  }

   /**
   * The keychain ID to use when looking up the the keychain plugin instance that will be used to retrieve the secret pointed to by the keychainEntryKey parameter.
   * @return keychainId
  **/
  @javax.annotation.Nullable

  public String getKeychainId() {
    return keychainId;
  }


  public void setKeychainId(String keychainId) {
    this.keychainId = keychainId;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Web3SigningCredentialCactiKeychainRef web3SigningCredentialCactiKeychainRef = (Web3SigningCredentialCactiKeychainRef) o;
    return Objects.equals(this.type, web3SigningCredentialCactiKeychainRef.type) &&
        Objects.equals(this.ethAccount, web3SigningCredentialCactiKeychainRef.ethAccount) &&
        Objects.equals(this.keychainEntryKey, web3SigningCredentialCactiKeychainRef.keychainEntryKey) &&
        Objects.equals(this.keychainId, web3SigningCredentialCactiKeychainRef.keychainId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, ethAccount, keychainEntryKey, keychainId);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Web3SigningCredentialCactiKeychainRef {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    ethAccount: ").append(toIndentedString(ethAccount)).append("\n");
    sb.append("    keychainEntryKey: ").append(toIndentedString(keychainEntryKey)).append("\n");
    sb.append("    keychainId: ").append(toIndentedString(keychainId)).append("\n");
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
    openapiFields.add("ethAccount");
    openapiFields.add("keychainEntryKey");
    openapiFields.add("keychainId");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("type");
    openapiRequiredFields.add("ethAccount");
    openapiRequiredFields.add("keychainEntryKey");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Web3SigningCredentialCactiKeychainRef
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!Web3SigningCredentialCactiKeychainRef.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in Web3SigningCredentialCactiKeychainRef is not found in the empty JSON string", Web3SigningCredentialCactiKeychainRef.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!Web3SigningCredentialCactiKeychainRef.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `Web3SigningCredentialCactiKeychainRef` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : Web3SigningCredentialCactiKeychainRef.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("ethAccount").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `ethAccount` to be a primitive type in the JSON string but got `%s`", jsonObj.get("ethAccount").toString()));
      }
      if (!jsonObj.get("keychainEntryKey").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `keychainEntryKey` to be a primitive type in the JSON string but got `%s`", jsonObj.get("keychainEntryKey").toString()));
      }
      if ((jsonObj.get("keychainId") != null && !jsonObj.get("keychainId").isJsonNull()) && !jsonObj.get("keychainId").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `keychainId` to be a primitive type in the JSON string but got `%s`", jsonObj.get("keychainId").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!Web3SigningCredentialCactiKeychainRef.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'Web3SigningCredentialCactiKeychainRef' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<Web3SigningCredentialCactiKeychainRef> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(Web3SigningCredentialCactiKeychainRef.class));

       return (TypeAdapter<T>) new TypeAdapter<Web3SigningCredentialCactiKeychainRef>() {
           @Override
           public void write(JsonWriter out, Web3SigningCredentialCactiKeychainRef value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public Web3SigningCredentialCactiKeychainRef read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of Web3SigningCredentialCactiKeychainRef given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Web3SigningCredentialCactiKeychainRef
  * @throws IOException if the JSON string is invalid with respect to Web3SigningCredentialCactiKeychainRef
  */
  public static Web3SigningCredentialCactiKeychainRef fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Web3SigningCredentialCactiKeychainRef.class);
  }

 /**
  * Convert an instance of Web3SigningCredentialCactiKeychainRef to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

