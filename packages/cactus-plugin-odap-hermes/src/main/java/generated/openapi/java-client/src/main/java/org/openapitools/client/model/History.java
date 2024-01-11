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
import java.util.ArrayList;
import java.util.List;

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
 * History
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class History {
  public static final String SERIALIZED_NAME_TRANSACTIONS = "Transactions";
  @SerializedName(SERIALIZED_NAME_TRANSACTIONS)
  private List<Object> transactions;

  public static final String SERIALIZED_NAME_ACTIONS = "Actions";
  @SerializedName(SERIALIZED_NAME_ACTIONS)
  private List<Object> actions;

  public static final String SERIALIZED_NAME_ORIGIN = "Origin";
  @SerializedName(SERIALIZED_NAME_ORIGIN)
  private String origin;

  public static final String SERIALIZED_NAME_DESTINATION = "Destination";
  @SerializedName(SERIALIZED_NAME_DESTINATION)
  private String destination;

  public static final String SERIALIZED_NAME_BALANCE = "Balance";
  @SerializedName(SERIALIZED_NAME_BALANCE)
  private String balance;

  public static final String SERIALIZED_NAME_CURRENT_STATUS = "CurrentStatus";
  @SerializedName(SERIALIZED_NAME_CURRENT_STATUS)
  private Object currentStatus;

  public static final String SERIALIZED_NAME_APPLICATION_SPECIFIC_PARAMETERS = "ApplicationSpecificParameters";
  @SerializedName(SERIALIZED_NAME_APPLICATION_SPECIFIC_PARAMETERS)
  private Object applicationSpecificParameters;

  public History() {
  }

  public History transactions(List<Object> transactions) {
    
    this.transactions = transactions;
    return this;
  }

  public History addTransactionsItem(Object transactionsItem) {
    if (this.transactions == null) {
      this.transactions = new ArrayList<>();
    }
    this.transactions.add(transactionsItem);
    return this;
  }

   /**
   * Get transactions
   * @return transactions
  **/
  @javax.annotation.Nullable
  public List<Object> getTransactions() {
    return transactions;
  }


  public void setTransactions(List<Object> transactions) {
    this.transactions = transactions;
  }


  public History actions(List<Object> actions) {
    
    this.actions = actions;
    return this;
  }

  public History addActionsItem(Object actionsItem) {
    if (this.actions == null) {
      this.actions = new ArrayList<>();
    }
    this.actions.add(actionsItem);
    return this;
  }

   /**
   * Get actions
   * @return actions
  **/
  @javax.annotation.Nullable
  public List<Object> getActions() {
    return actions;
  }


  public void setActions(List<Object> actions) {
    this.actions = actions;
  }


  public History origin(String origin) {
    
    this.origin = origin;
    return this;
  }

   /**
   * Get origin
   * @return origin
  **/
  @javax.annotation.Nullable
  public String getOrigin() {
    return origin;
  }


  public void setOrigin(String origin) {
    this.origin = origin;
  }


  public History destination(String destination) {
    
    this.destination = destination;
    return this;
  }

   /**
   * Get destination
   * @return destination
  **/
  @javax.annotation.Nullable
  public String getDestination() {
    return destination;
  }


  public void setDestination(String destination) {
    this.destination = destination;
  }


  public History balance(String balance) {
    
    this.balance = balance;
    return this;
  }

   /**
   * Get balance
   * @return balance
  **/
  @javax.annotation.Nullable
  public String getBalance() {
    return balance;
  }


  public void setBalance(String balance) {
    this.balance = balance;
  }


  public History currentStatus(Object currentStatus) {
    
    this.currentStatus = currentStatus;
    return this;
  }

   /**
   * Get currentStatus
   * @return currentStatus
  **/
  @javax.annotation.Nullable
  public Object getCurrentStatus() {
    return currentStatus;
  }


  public void setCurrentStatus(Object currentStatus) {
    this.currentStatus = currentStatus;
  }


  public History applicationSpecificParameters(Object applicationSpecificParameters) {
    
    this.applicationSpecificParameters = applicationSpecificParameters;
    return this;
  }

   /**
   * Get applicationSpecificParameters
   * @return applicationSpecificParameters
  **/
  @javax.annotation.Nullable
  public Object getApplicationSpecificParameters() {
    return applicationSpecificParameters;
  }


  public void setApplicationSpecificParameters(Object applicationSpecificParameters) {
    this.applicationSpecificParameters = applicationSpecificParameters;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    History history = (History) o;
    return Objects.equals(this.transactions, history.transactions) &&
        Objects.equals(this.actions, history.actions) &&
        Objects.equals(this.origin, history.origin) &&
        Objects.equals(this.destination, history.destination) &&
        Objects.equals(this.balance, history.balance) &&
        Objects.equals(this.currentStatus, history.currentStatus) &&
        Objects.equals(this.applicationSpecificParameters, history.applicationSpecificParameters);
  }

  @Override
  public int hashCode() {
    return Objects.hash(transactions, actions, origin, destination, balance, currentStatus, applicationSpecificParameters);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class History {\n");
    sb.append("    transactions: ").append(toIndentedString(transactions)).append("\n");
    sb.append("    actions: ").append(toIndentedString(actions)).append("\n");
    sb.append("    origin: ").append(toIndentedString(origin)).append("\n");
    sb.append("    destination: ").append(toIndentedString(destination)).append("\n");
    sb.append("    balance: ").append(toIndentedString(balance)).append("\n");
    sb.append("    currentStatus: ").append(toIndentedString(currentStatus)).append("\n");
    sb.append("    applicationSpecificParameters: ").append(toIndentedString(applicationSpecificParameters)).append("\n");
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
    openapiFields.add("Transactions");
    openapiFields.add("Actions");
    openapiFields.add("Origin");
    openapiFields.add("Destination");
    openapiFields.add("Balance");
    openapiFields.add("CurrentStatus");
    openapiFields.add("ApplicationSpecificParameters");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to History
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!History.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in History is not found in the empty JSON string", History.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!History.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `History` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("Transactions") != null && !jsonObj.get("Transactions").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `Transactions` to be an array in the JSON string but got `%s`", jsonObj.get("Transactions").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("Actions") != null && !jsonObj.get("Actions").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `Actions` to be an array in the JSON string but got `%s`", jsonObj.get("Actions").toString()));
      }
      if ((jsonObj.get("Origin") != null && !jsonObj.get("Origin").isJsonNull()) && !jsonObj.get("Origin").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `Origin` to be a primitive type in the JSON string but got `%s`", jsonObj.get("Origin").toString()));
      }
      if ((jsonObj.get("Destination") != null && !jsonObj.get("Destination").isJsonNull()) && !jsonObj.get("Destination").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `Destination` to be a primitive type in the JSON string but got `%s`", jsonObj.get("Destination").toString()));
      }
      if ((jsonObj.get("Balance") != null && !jsonObj.get("Balance").isJsonNull()) && !jsonObj.get("Balance").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `Balance` to be a primitive type in the JSON string but got `%s`", jsonObj.get("Balance").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!History.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'History' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<History> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(History.class));

       return (TypeAdapter<T>) new TypeAdapter<History>() {
           @Override
           public void write(JsonWriter out, History value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public History read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of History given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of History
  * @throws IOException if the JSON string is invalid with respect to History
  */
  public static History fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, History.class);
  }

 /**
  * Convert an instance of History to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

