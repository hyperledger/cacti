/*
 * Hyperledger Cactus Plugin - Odap Hermes
 * Implementation for Odap and Hermes
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
 * LockEvidenceV1Response
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class LockEvidenceV1Response {
  public static final String SERIALIZED_NAME_SESSION_I_D = "sessionID";
  @SerializedName(SERIALIZED_NAME_SESSION_I_D)
  private String sessionID;

  public static final String SERIALIZED_NAME_CLIENT_IDENTITY_PUBKEY = "clientIdentityPubkey";
  @SerializedName(SERIALIZED_NAME_CLIENT_IDENTITY_PUBKEY)
  private String clientIdentityPubkey;

  public static final String SERIALIZED_NAME_SERVER_IDENTITY_PUBKEY = "serverIdentityPubkey";
  @SerializedName(SERIALIZED_NAME_SERVER_IDENTITY_PUBKEY)
  private String serverIdentityPubkey;

  public static final String SERIALIZED_NAME_HASH_LOCK_EVIDENCE_REQUEST = "hashLockEvidenceRequest";
  @SerializedName(SERIALIZED_NAME_HASH_LOCK_EVIDENCE_REQUEST)
  private String hashLockEvidenceRequest;

  public static final String SERIALIZED_NAME_SERVER_TRANSFER_NUMBER = "serverTransferNumber";
  @SerializedName(SERIALIZED_NAME_SERVER_TRANSFER_NUMBER)
  private Integer serverTransferNumber;

  public static final String SERIALIZED_NAME_SIGNATURE = "signature";
  @SerializedName(SERIALIZED_NAME_SIGNATURE)
  private String signature;

  public static final String SERIALIZED_NAME_MESSAGE_TYPE = "messageType";
  @SerializedName(SERIALIZED_NAME_MESSAGE_TYPE)
  private String messageType;

  public static final String SERIALIZED_NAME_SEQUENCE_NUMBER = "sequenceNumber";
  @SerializedName(SERIALIZED_NAME_SEQUENCE_NUMBER)
  private BigDecimal sequenceNumber;

  public LockEvidenceV1Response() {
  }

  public LockEvidenceV1Response sessionID(String sessionID) {
    
    this.sessionID = sessionID;
    return this;
  }

   /**
   * Get sessionID
   * @return sessionID
  **/
  @javax.annotation.Nonnull
  public String getSessionID() {
    return sessionID;
  }


  public void setSessionID(String sessionID) {
    this.sessionID = sessionID;
  }


  public LockEvidenceV1Response clientIdentityPubkey(String clientIdentityPubkey) {
    
    this.clientIdentityPubkey = clientIdentityPubkey;
    return this;
  }

   /**
   * Get clientIdentityPubkey
   * @return clientIdentityPubkey
  **/
  @javax.annotation.Nonnull
  public String getClientIdentityPubkey() {
    return clientIdentityPubkey;
  }


  public void setClientIdentityPubkey(String clientIdentityPubkey) {
    this.clientIdentityPubkey = clientIdentityPubkey;
  }


  public LockEvidenceV1Response serverIdentityPubkey(String serverIdentityPubkey) {
    
    this.serverIdentityPubkey = serverIdentityPubkey;
    return this;
  }

   /**
   * Get serverIdentityPubkey
   * @return serverIdentityPubkey
  **/
  @javax.annotation.Nonnull
  public String getServerIdentityPubkey() {
    return serverIdentityPubkey;
  }


  public void setServerIdentityPubkey(String serverIdentityPubkey) {
    this.serverIdentityPubkey = serverIdentityPubkey;
  }


  public LockEvidenceV1Response hashLockEvidenceRequest(String hashLockEvidenceRequest) {
    
    this.hashLockEvidenceRequest = hashLockEvidenceRequest;
    return this;
  }

   /**
   * Get hashLockEvidenceRequest
   * @return hashLockEvidenceRequest
  **/
  @javax.annotation.Nonnull
  public String getHashLockEvidenceRequest() {
    return hashLockEvidenceRequest;
  }


  public void setHashLockEvidenceRequest(String hashLockEvidenceRequest) {
    this.hashLockEvidenceRequest = hashLockEvidenceRequest;
  }


  public LockEvidenceV1Response serverTransferNumber(Integer serverTransferNumber) {
    
    this.serverTransferNumber = serverTransferNumber;
    return this;
  }

   /**
   * Get serverTransferNumber
   * @return serverTransferNumber
  **/
  @javax.annotation.Nullable
  public Integer getServerTransferNumber() {
    return serverTransferNumber;
  }


  public void setServerTransferNumber(Integer serverTransferNumber) {
    this.serverTransferNumber = serverTransferNumber;
  }


  public LockEvidenceV1Response signature(String signature) {
    
    this.signature = signature;
    return this;
  }

   /**
   * Get signature
   * @return signature
  **/
  @javax.annotation.Nonnull
  public String getSignature() {
    return signature;
  }


  public void setSignature(String signature) {
    this.signature = signature;
  }


  public LockEvidenceV1Response messageType(String messageType) {
    
    this.messageType = messageType;
    return this;
  }

   /**
   * Get messageType
   * @return messageType
  **/
  @javax.annotation.Nonnull
  public String getMessageType() {
    return messageType;
  }


  public void setMessageType(String messageType) {
    this.messageType = messageType;
  }


  public LockEvidenceV1Response sequenceNumber(BigDecimal sequenceNumber) {
    
    this.sequenceNumber = sequenceNumber;
    return this;
  }

   /**
   * Get sequenceNumber
   * @return sequenceNumber
  **/
  @javax.annotation.Nonnull
  public BigDecimal getSequenceNumber() {
    return sequenceNumber;
  }


  public void setSequenceNumber(BigDecimal sequenceNumber) {
    this.sequenceNumber = sequenceNumber;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    LockEvidenceV1Response lockEvidenceV1Response = (LockEvidenceV1Response) o;
    return Objects.equals(this.sessionID, lockEvidenceV1Response.sessionID) &&
        Objects.equals(this.clientIdentityPubkey, lockEvidenceV1Response.clientIdentityPubkey) &&
        Objects.equals(this.serverIdentityPubkey, lockEvidenceV1Response.serverIdentityPubkey) &&
        Objects.equals(this.hashLockEvidenceRequest, lockEvidenceV1Response.hashLockEvidenceRequest) &&
        Objects.equals(this.serverTransferNumber, lockEvidenceV1Response.serverTransferNumber) &&
        Objects.equals(this.signature, lockEvidenceV1Response.signature) &&
        Objects.equals(this.messageType, lockEvidenceV1Response.messageType) &&
        Objects.equals(this.sequenceNumber, lockEvidenceV1Response.sequenceNumber);
  }

  private static <T> boolean equalsNullable(JsonNullable<T> a, JsonNullable<T> b) {
    return a == b || (a != null && b != null && a.isPresent() && b.isPresent() && Objects.deepEquals(a.get(), b.get()));
  }

  @Override
  public int hashCode() {
    return Objects.hash(sessionID, clientIdentityPubkey, serverIdentityPubkey, hashLockEvidenceRequest, serverTransferNumber, signature, messageType, sequenceNumber);
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
    sb.append("class LockEvidenceV1Response {\n");
    sb.append("    sessionID: ").append(toIndentedString(sessionID)).append("\n");
    sb.append("    clientIdentityPubkey: ").append(toIndentedString(clientIdentityPubkey)).append("\n");
    sb.append("    serverIdentityPubkey: ").append(toIndentedString(serverIdentityPubkey)).append("\n");
    sb.append("    hashLockEvidenceRequest: ").append(toIndentedString(hashLockEvidenceRequest)).append("\n");
    sb.append("    serverTransferNumber: ").append(toIndentedString(serverTransferNumber)).append("\n");
    sb.append("    signature: ").append(toIndentedString(signature)).append("\n");
    sb.append("    messageType: ").append(toIndentedString(messageType)).append("\n");
    sb.append("    sequenceNumber: ").append(toIndentedString(sequenceNumber)).append("\n");
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
    openapiFields.add("sessionID");
    openapiFields.add("clientIdentityPubkey");
    openapiFields.add("serverIdentityPubkey");
    openapiFields.add("hashLockEvidenceRequest");
    openapiFields.add("serverTransferNumber");
    openapiFields.add("signature");
    openapiFields.add("messageType");
    openapiFields.add("sequenceNumber");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("sessionID");
    openapiRequiredFields.add("clientIdentityPubkey");
    openapiRequiredFields.add("serverIdentityPubkey");
    openapiRequiredFields.add("hashLockEvidenceRequest");
    openapiRequiredFields.add("signature");
    openapiRequiredFields.add("messageType");
    openapiRequiredFields.add("sequenceNumber");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to LockEvidenceV1Response
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!LockEvidenceV1Response.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in LockEvidenceV1Response is not found in the empty JSON string", LockEvidenceV1Response.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!LockEvidenceV1Response.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `LockEvidenceV1Response` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : LockEvidenceV1Response.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if (!jsonObj.get("sessionID").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sessionID` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sessionID").toString()));
      }
      if (!jsonObj.get("clientIdentityPubkey").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `clientIdentityPubkey` to be a primitive type in the JSON string but got `%s`", jsonObj.get("clientIdentityPubkey").toString()));
      }
      if (!jsonObj.get("serverIdentityPubkey").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `serverIdentityPubkey` to be a primitive type in the JSON string but got `%s`", jsonObj.get("serverIdentityPubkey").toString()));
      }
      if (!jsonObj.get("hashLockEvidenceRequest").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `hashLockEvidenceRequest` to be a primitive type in the JSON string but got `%s`", jsonObj.get("hashLockEvidenceRequest").toString()));
      }
      if (!jsonObj.get("signature").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `signature` to be a primitive type in the JSON string but got `%s`", jsonObj.get("signature").toString()));
      }
      if (!jsonObj.get("messageType").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `messageType` to be a primitive type in the JSON string but got `%s`", jsonObj.get("messageType").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!LockEvidenceV1Response.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'LockEvidenceV1Response' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<LockEvidenceV1Response> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(LockEvidenceV1Response.class));

       return (TypeAdapter<T>) new TypeAdapter<LockEvidenceV1Response>() {
           @Override
           public void write(JsonWriter out, LockEvidenceV1Response value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public LockEvidenceV1Response read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of LockEvidenceV1Response given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of LockEvidenceV1Response
  * @throws IOException if the JSON string is invalid with respect to LockEvidenceV1Response
  */
  public static LockEvidenceV1Response fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, LockEvidenceV1Response.class);
  }

 /**
  * Convert an instance of LockEvidenceV1Response to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

