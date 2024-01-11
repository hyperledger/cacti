/*
 * Hyperledger Cactus Plugin - Connector Xdai
 * Can perform basic tasks on a Xdai ledger
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
import org.openapitools.client.model.Web3SigningCredentialCactusKeychainRef;
import org.openapitools.client.model.Web3SigningCredentialNone;
import org.openapitools.client.model.Web3SigningCredentialPrivateKeyHex;
import org.openapitools.client.model.Web3SigningCredentialType;

import javax.ws.rs.core.GenericType;

import java.io.IOException;
import java.lang.reflect.Type;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashSet;
import java.util.HashMap;
import java.util.Map;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapter;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;
import com.google.gson.JsonPrimitive;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;

import org.openapitools.client.JSON;

@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class Web3SigningCredential extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(Web3SigningCredential.class.getName());

    public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
        @SuppressWarnings("unchecked")
        @Override
        public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
            if (!Web3SigningCredential.class.isAssignableFrom(type.getRawType())) {
                return null; // this class only serializes 'Web3SigningCredential' and its subtypes
            }
            final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
            final TypeAdapter<Web3SigningCredentialCactusKeychainRef> adapterWeb3SigningCredentialCactusKeychainRef = gson.getDelegateAdapter(this, TypeToken.get(Web3SigningCredentialCactusKeychainRef.class));
            final TypeAdapter<Web3SigningCredentialNone> adapterWeb3SigningCredentialNone = gson.getDelegateAdapter(this, TypeToken.get(Web3SigningCredentialNone.class));
            final TypeAdapter<Web3SigningCredentialPrivateKeyHex> adapterWeb3SigningCredentialPrivateKeyHex = gson.getDelegateAdapter(this, TypeToken.get(Web3SigningCredentialPrivateKeyHex.class));

            return (TypeAdapter<T>) new TypeAdapter<Web3SigningCredential>() {
                @Override
                public void write(JsonWriter out, Web3SigningCredential value) throws IOException {
                    if (value == null || value.getActualInstance() == null) {
                        elementAdapter.write(out, null);
                        return;
                    }

                    // check if the actual instance is of the type `Web3SigningCredentialCactusKeychainRef`
                    if (value.getActualInstance() instanceof Web3SigningCredentialCactusKeychainRef) {
                        JsonObject obj = adapterWeb3SigningCredentialCactusKeychainRef.toJsonTree((Web3SigningCredentialCactusKeychainRef)value.getActualInstance()).getAsJsonObject();
                        elementAdapter.write(out, obj);
                        return;
                    }

                    // check if the actual instance is of the type `Web3SigningCredentialNone`
                    if (value.getActualInstance() instanceof Web3SigningCredentialNone) {
                        JsonObject obj = adapterWeb3SigningCredentialNone.toJsonTree((Web3SigningCredentialNone)value.getActualInstance()).getAsJsonObject();
                        elementAdapter.write(out, obj);
                        return;
                    }

                    // check if the actual instance is of the type `Web3SigningCredentialPrivateKeyHex`
                    if (value.getActualInstance() instanceof Web3SigningCredentialPrivateKeyHex) {
                        JsonObject obj = adapterWeb3SigningCredentialPrivateKeyHex.toJsonTree((Web3SigningCredentialPrivateKeyHex)value.getActualInstance()).getAsJsonObject();
                        elementAdapter.write(out, obj);
                        return;
                    }

                    throw new IOException("Failed to serialize as the type doesn't match oneOf schemas: Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex");
                }

                @Override
                public Web3SigningCredential read(JsonReader in) throws IOException {
                    Object deserialized = null;
                    JsonObject jsonObject = elementAdapter.read(in).getAsJsonObject();

                    int match = 0;
                    ArrayList<String> errorMessages = new ArrayList<>();
                    TypeAdapter actualAdapter = elementAdapter;

                    // deserialize Web3SigningCredentialCactusKeychainRef
                    try {
                        // validate the JSON object to see if any exception is thrown
                        Web3SigningCredentialCactusKeychainRef.validateJsonObject(jsonObject);
                        actualAdapter = adapterWeb3SigningCredentialCactusKeychainRef;
                        match++;
                        log.log(Level.FINER, "Input data matches schema 'Web3SigningCredentialCactusKeychainRef'");
                    } catch (Exception e) {
                        // deserialization failed, continue
                        errorMessages.add(String.format("Deserialization for Web3SigningCredentialCactusKeychainRef failed with `%s`.", e.getMessage()));
                        log.log(Level.FINER, "Input data does not match schema 'Web3SigningCredentialCactusKeychainRef'", e);
                    }

                    // deserialize Web3SigningCredentialNone
                    try {
                        // validate the JSON object to see if any exception is thrown
                        Web3SigningCredentialNone.validateJsonObject(jsonObject);
                        actualAdapter = adapterWeb3SigningCredentialNone;
                        match++;
                        log.log(Level.FINER, "Input data matches schema 'Web3SigningCredentialNone'");
                    } catch (Exception e) {
                        // deserialization failed, continue
                        errorMessages.add(String.format("Deserialization for Web3SigningCredentialNone failed with `%s`.", e.getMessage()));
                        log.log(Level.FINER, "Input data does not match schema 'Web3SigningCredentialNone'", e);
                    }

                    // deserialize Web3SigningCredentialPrivateKeyHex
                    try {
                        // validate the JSON object to see if any exception is thrown
                        Web3SigningCredentialPrivateKeyHex.validateJsonObject(jsonObject);
                        actualAdapter = adapterWeb3SigningCredentialPrivateKeyHex;
                        match++;
                        log.log(Level.FINER, "Input data matches schema 'Web3SigningCredentialPrivateKeyHex'");
                    } catch (Exception e) {
                        // deserialization failed, continue
                        errorMessages.add(String.format("Deserialization for Web3SigningCredentialPrivateKeyHex failed with `%s`.", e.getMessage()));
                        log.log(Level.FINER, "Input data does not match schema 'Web3SigningCredentialPrivateKeyHex'", e);
                    }

                    if (match == 1) {
                        Web3SigningCredential ret = new Web3SigningCredential();
                        ret.setActualInstance(actualAdapter.fromJsonTree(jsonObject));
                        return ret;
                    }

                    throw new IOException(String.format("Failed deserialization for Web3SigningCredential: %d classes match result, expected 1. Detailed failure message for oneOf schemas: %s. JSON: %s", match, errorMessages, jsonObject.toString()));
                }
            }.nullSafe();
        }
    }

    // store a list of schema names defined in oneOf
    public static final Map<String, GenericType> schemas = new HashMap<String, GenericType>();

    public Web3SigningCredential() {
        super("oneOf", Boolean.FALSE);
    }

    public Web3SigningCredential(Web3SigningCredentialCactusKeychainRef o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public Web3SigningCredential(Web3SigningCredentialNone o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public Web3SigningCredential(Web3SigningCredentialPrivateKeyHex o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    static {
        schemas.put("Web3SigningCredentialCactusKeychainRef", new GenericType<Web3SigningCredentialCactusKeychainRef>() {
        });
        schemas.put("Web3SigningCredentialNone", new GenericType<Web3SigningCredentialNone>() {
        });
        schemas.put("Web3SigningCredentialPrivateKeyHex", new GenericType<Web3SigningCredentialPrivateKeyHex>() {
        });
    }

    @Override
    public Map<String, GenericType> getSchemas() {
        return Web3SigningCredential.schemas;
    }

    /**
     * Set the instance that matches the oneOf child schema, check
     * the instance parameter is valid against the oneOf child schemas:
     * Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex
     *
     * It could be an instance of the 'oneOf' schemas.
     * The oneOf child schemas may themselves be a composed schema (allOf, anyOf, oneOf).
     */
    @Override
    public void setActualInstance(Object instance) {
        if (instance instanceof Web3SigningCredentialCactusKeychainRef) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Web3SigningCredentialNone) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof Web3SigningCredentialPrivateKeyHex) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex");
    }

    /**
     * Get the actual instance, which can be the following:
     * Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex
     *
     * @return The actual instance (Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex)
     */
    @Override
    public Object getActualInstance() {
        return super.getActualInstance();
    }

    /**
     * Get the actual instance of `Web3SigningCredentialCactusKeychainRef`. If the actual instance is not `Web3SigningCredentialCactusKeychainRef`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Web3SigningCredentialCactusKeychainRef`
     * @throws ClassCastException if the instance is not `Web3SigningCredentialCactusKeychainRef`
     */
    public Web3SigningCredentialCactusKeychainRef getWeb3SigningCredentialCactusKeychainRef() throws ClassCastException {
        return (Web3SigningCredentialCactusKeychainRef)super.getActualInstance();
    }

    /**
     * Get the actual instance of `Web3SigningCredentialNone`. If the actual instance is not `Web3SigningCredentialNone`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Web3SigningCredentialNone`
     * @throws ClassCastException if the instance is not `Web3SigningCredentialNone`
     */
    public Web3SigningCredentialNone getWeb3SigningCredentialNone() throws ClassCastException {
        return (Web3SigningCredentialNone)super.getActualInstance();
    }

    /**
     * Get the actual instance of `Web3SigningCredentialPrivateKeyHex`. If the actual instance is not `Web3SigningCredentialPrivateKeyHex`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `Web3SigningCredentialPrivateKeyHex`
     * @throws ClassCastException if the instance is not `Web3SigningCredentialPrivateKeyHex`
     */
    public Web3SigningCredentialPrivateKeyHex getWeb3SigningCredentialPrivateKeyHex() throws ClassCastException {
        return (Web3SigningCredentialPrivateKeyHex)super.getActualInstance();
    }


 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to Web3SigningCredential
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
    // validate oneOf schemas one by one
    int validCount = 0;
    ArrayList<String> errorMessages = new ArrayList<>();
    // validate the json string with Web3SigningCredentialCactusKeychainRef
    try {
      Web3SigningCredentialCactusKeychainRef.validateJsonObject(jsonObj);
      validCount++;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Web3SigningCredentialCactusKeychainRef failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Web3SigningCredentialNone
    try {
      Web3SigningCredentialNone.validateJsonObject(jsonObj);
      validCount++;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Web3SigningCredentialNone failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with Web3SigningCredentialPrivateKeyHex
    try {
      Web3SigningCredentialPrivateKeyHex.validateJsonObject(jsonObj);
      validCount++;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for Web3SigningCredentialPrivateKeyHex failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    if (validCount != 1) {
      throw new IOException(String.format("The JSON string is invalid for Web3SigningCredential with oneOf schemas: Web3SigningCredentialCactusKeychainRef, Web3SigningCredentialNone, Web3SigningCredentialPrivateKeyHex. %d class(es) match the result, expected 1. Detailed failure message for oneOf schemas: %s. JSON: %s", validCount, errorMessages, jsonObj.toString()));
    }
  }

 /**
  * Create an instance of Web3SigningCredential given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of Web3SigningCredential
  * @throws IOException if the JSON string is invalid with respect to Web3SigningCredential
  */
  public static Web3SigningCredential fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, Web3SigningCredential.class);
  }

 /**
  * Convert an instance of Web3SigningCredential to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

