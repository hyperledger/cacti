/*
 * Hyperledger Cactus Plugin - Connector Iroha V2
 * Can perform basic tasks on a Iroha V2 ledger
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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.IrohaQuery;
import org.openapitools.client.model.IrohaQueryDefinitionV1;
import org.openapitools.client.model.IrohaTransactionDefinitionV1;
import org.openapitools.client.model.IrohaTransactionDefinitionV1Instruction;

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
public class GenerateTransactionRequestV1Request extends AbstractOpenApiSchema {
    private static final Logger log = Logger.getLogger(GenerateTransactionRequestV1Request.class.getName());

    public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
        @SuppressWarnings("unchecked")
        @Override
        public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
            if (!GenerateTransactionRequestV1Request.class.isAssignableFrom(type.getRawType())) {
                return null; // this class only serializes 'GenerateTransactionRequestV1Request' and its subtypes
            }
            final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
            final TypeAdapter<IrohaQueryDefinitionV1> adapterIrohaQueryDefinitionV1 = gson.getDelegateAdapter(this, TypeToken.get(IrohaQueryDefinitionV1.class));
            final TypeAdapter<IrohaTransactionDefinitionV1> adapterIrohaTransactionDefinitionV1 = gson.getDelegateAdapter(this, TypeToken.get(IrohaTransactionDefinitionV1.class));

            return (TypeAdapter<T>) new TypeAdapter<GenerateTransactionRequestV1Request>() {
                @Override
                public void write(JsonWriter out, GenerateTransactionRequestV1Request value) throws IOException {
                    if (value == null || value.getActualInstance() == null) {
                        elementAdapter.write(out, null);
                        return;
                    }

                    // check if the actual instance is of the type `IrohaQueryDefinitionV1`
                    if (value.getActualInstance() instanceof IrohaQueryDefinitionV1) {
                        JsonObject obj = adapterIrohaQueryDefinitionV1.toJsonTree((IrohaQueryDefinitionV1)value.getActualInstance()).getAsJsonObject();
                        elementAdapter.write(out, obj);
                        return;
                    }

                    // check if the actual instance is of the type `IrohaTransactionDefinitionV1`
                    if (value.getActualInstance() instanceof IrohaTransactionDefinitionV1) {
                        JsonObject obj = adapterIrohaTransactionDefinitionV1.toJsonTree((IrohaTransactionDefinitionV1)value.getActualInstance()).getAsJsonObject();
                        elementAdapter.write(out, obj);
                        return;
                    }

                    throw new IOException("Failed to serialize as the type doesn't match oneOf schemas: IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1");
                }

                @Override
                public GenerateTransactionRequestV1Request read(JsonReader in) throws IOException {
                    Object deserialized = null;
                    JsonObject jsonObject = elementAdapter.read(in).getAsJsonObject();

                    int match = 0;
                    ArrayList<String> errorMessages = new ArrayList<>();
                    TypeAdapter actualAdapter = elementAdapter;

                    // deserialize IrohaQueryDefinitionV1
                    try {
                        // validate the JSON object to see if any exception is thrown
                        IrohaQueryDefinitionV1.validateJsonObject(jsonObject);
                        actualAdapter = adapterIrohaQueryDefinitionV1;
                        match++;
                        log.log(Level.FINER, "Input data matches schema 'IrohaQueryDefinitionV1'");
                    } catch (Exception e) {
                        // deserialization failed, continue
                        errorMessages.add(String.format("Deserialization for IrohaQueryDefinitionV1 failed with `%s`.", e.getMessage()));
                        log.log(Level.FINER, "Input data does not match schema 'IrohaQueryDefinitionV1'", e);
                    }

                    // deserialize IrohaTransactionDefinitionV1
                    try {
                        // validate the JSON object to see if any exception is thrown
                        IrohaTransactionDefinitionV1.validateJsonObject(jsonObject);
                        actualAdapter = adapterIrohaTransactionDefinitionV1;
                        match++;
                        log.log(Level.FINER, "Input data matches schema 'IrohaTransactionDefinitionV1'");
                    } catch (Exception e) {
                        // deserialization failed, continue
                        errorMessages.add(String.format("Deserialization for IrohaTransactionDefinitionV1 failed with `%s`.", e.getMessage()));
                        log.log(Level.FINER, "Input data does not match schema 'IrohaTransactionDefinitionV1'", e);
                    }

                    if (match == 1) {
                        GenerateTransactionRequestV1Request ret = new GenerateTransactionRequestV1Request();
                        ret.setActualInstance(actualAdapter.fromJsonTree(jsonObject));
                        return ret;
                    }

                    throw new IOException(String.format("Failed deserialization for GenerateTransactionRequestV1Request: %d classes match result, expected 1. Detailed failure message for oneOf schemas: %s. JSON: %s", match, errorMessages, jsonObject.toString()));
                }
            }.nullSafe();
        }
    }

    // store a list of schema names defined in oneOf
    public static final Map<String, GenericType> schemas = new HashMap<String, GenericType>();

    public GenerateTransactionRequestV1Request() {
        super("oneOf", Boolean.FALSE);
    }

    public GenerateTransactionRequestV1Request(IrohaQueryDefinitionV1 o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    public GenerateTransactionRequestV1Request(IrohaTransactionDefinitionV1 o) {
        super("oneOf", Boolean.FALSE);
        setActualInstance(o);
    }

    static {
        schemas.put("IrohaQueryDefinitionV1", new GenericType<IrohaQueryDefinitionV1>() {
        });
        schemas.put("IrohaTransactionDefinitionV1", new GenericType<IrohaTransactionDefinitionV1>() {
        });
    }

    @Override
    public Map<String, GenericType> getSchemas() {
        return GenerateTransactionRequestV1Request.schemas;
    }

    /**
     * Set the instance that matches the oneOf child schema, check
     * the instance parameter is valid against the oneOf child schemas:
     * IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1
     *
     * It could be an instance of the 'oneOf' schemas.
     * The oneOf child schemas may themselves be a composed schema (allOf, anyOf, oneOf).
     */
    @Override
    public void setActualInstance(Object instance) {
        if (instance instanceof IrohaQueryDefinitionV1) {
            super.setActualInstance(instance);
            return;
        }

        if (instance instanceof IrohaTransactionDefinitionV1) {
            super.setActualInstance(instance);
            return;
        }

        throw new RuntimeException("Invalid instance type. Must be IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1");
    }

    /**
     * Get the actual instance, which can be the following:
     * IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1
     *
     * @return The actual instance (IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1)
     */
    @Override
    public Object getActualInstance() {
        return super.getActualInstance();
    }

    /**
     * Get the actual instance of `IrohaQueryDefinitionV1`. If the actual instance is not `IrohaQueryDefinitionV1`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `IrohaQueryDefinitionV1`
     * @throws ClassCastException if the instance is not `IrohaQueryDefinitionV1`
     */
    public IrohaQueryDefinitionV1 getIrohaQueryDefinitionV1() throws ClassCastException {
        return (IrohaQueryDefinitionV1)super.getActualInstance();
    }

    /**
     * Get the actual instance of `IrohaTransactionDefinitionV1`. If the actual instance is not `IrohaTransactionDefinitionV1`,
     * the ClassCastException will be thrown.
     *
     * @return The actual instance of `IrohaTransactionDefinitionV1`
     * @throws ClassCastException if the instance is not `IrohaTransactionDefinitionV1`
     */
    public IrohaTransactionDefinitionV1 getIrohaTransactionDefinitionV1() throws ClassCastException {
        return (IrohaTransactionDefinitionV1)super.getActualInstance();
    }


 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to GenerateTransactionRequestV1Request
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
    // validate oneOf schemas one by one
    int validCount = 0;
    ArrayList<String> errorMessages = new ArrayList<>();
    // validate the json string with IrohaQueryDefinitionV1
    try {
      IrohaQueryDefinitionV1.validateJsonObject(jsonObj);
      validCount++;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for IrohaQueryDefinitionV1 failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    // validate the json string with IrohaTransactionDefinitionV1
    try {
      IrohaTransactionDefinitionV1.validateJsonObject(jsonObj);
      validCount++;
    } catch (Exception e) {
      errorMessages.add(String.format("Deserialization for IrohaTransactionDefinitionV1 failed with `%s`.", e.getMessage()));
      // continue to the next one
    }
    if (validCount != 1) {
      throw new IOException(String.format("The JSON string is invalid for GenerateTransactionRequestV1Request with oneOf schemas: IrohaQueryDefinitionV1, IrohaTransactionDefinitionV1. %d class(es) match the result, expected 1. Detailed failure message for oneOf schemas: %s. JSON: %s", validCount, errorMessages, jsonObj.toString()));
    }
  }

 /**
  * Create an instance of GenerateTransactionRequestV1Request given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of GenerateTransactionRequestV1Request
  * @throws IOException if the JSON string is invalid with respect to GenerateTransactionRequestV1Request
  */
  public static GenerateTransactionRequestV1Request fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, GenerateTransactionRequestV1Request.class);
  }

 /**
  * Convert an instance of GenerateTransactionRequestV1Request to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

