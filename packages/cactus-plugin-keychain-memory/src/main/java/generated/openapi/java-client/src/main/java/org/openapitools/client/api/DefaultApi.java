/*
 * Hyperledger Cactus Plugin - Keychain Memory 
 * Contains/describes the Hyperledger Cacti Keychain Memory plugin.
 *
 * The version of the OpenAPI document: v2.0.0-alpha.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiCallback;
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.ApiResponse;
import org.openapitools.client.Configuration;
import org.openapitools.client.Pair;
import org.openapitools.client.ProgressRequestBody;
import org.openapitools.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import org.openapitools.client.model.DeleteKeychainEntryRequestV1;
import org.openapitools.client.model.DeleteKeychainEntryResponseV1;
import org.openapitools.client.model.GetKeychainEntryRequestV1;
import org.openapitools.client.model.GetKeychainEntryResponseV1;
import org.openapitools.client.model.HasKeychainEntryRequestV1;
import org.openapitools.client.model.HasKeychainEntryResponseV1;
import org.openapitools.client.model.SetKeychainEntryRequestV1;
import org.openapitools.client.model.SetKeychainEntryResponseV1;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import javax.ws.rs.core.GenericType;

public class DefaultApi {
    private ApiClient localVarApiClient;
    private int localHostIndex;
    private String localCustomBaseUrl;

    public DefaultApi() {
        this(Configuration.getDefaultApiClient());
    }

    public DefaultApi(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public int getHostIndex() {
        return localHostIndex;
    }

    public void setHostIndex(int hostIndex) {
        this.localHostIndex = hostIndex;
    }

    public String getCustomBaseUrl() {
        return localCustomBaseUrl;
    }

    public void setCustomBaseUrl(String customBaseUrl) {
        this.localCustomBaseUrl = customBaseUrl;
    }

    /**
     * Build call for deleteKeychainEntryV1
     * @param deleteKeychainEntryRequestV1 Request body to delete a keychain entry via its key (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call deleteKeychainEntryV1Call(DeleteKeychainEntryRequestV1 deleteKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = deleteKeychainEntryRequestV1;

        // create path and map variables
        String localVarPath = "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/delete-keychain-entry";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call deleteKeychainEntryV1ValidateBeforeCall(DeleteKeychainEntryRequestV1 deleteKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'deleteKeychainEntryRequestV1' is set
        if (deleteKeychainEntryRequestV1 == null) {
            throw new ApiException("Missing the required parameter 'deleteKeychainEntryRequestV1' when calling deleteKeychainEntryV1(Async)");
        }

        return deleteKeychainEntryV1Call(deleteKeychainEntryRequestV1, _callback);

    }

    /**
     * Deletes an entry under a key on the keychain backend.
     * 
     * @param deleteKeychainEntryRequestV1 Request body to delete a keychain entry via its key (required)
     * @return DeleteKeychainEntryResponseV1
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public DeleteKeychainEntryResponseV1 deleteKeychainEntryV1(DeleteKeychainEntryRequestV1 deleteKeychainEntryRequestV1) throws ApiException {
        ApiResponse<DeleteKeychainEntryResponseV1> localVarResp = deleteKeychainEntryV1WithHttpInfo(deleteKeychainEntryRequestV1);
        return localVarResp.getData();
    }

    /**
     * Deletes an entry under a key on the keychain backend.
     * 
     * @param deleteKeychainEntryRequestV1 Request body to delete a keychain entry via its key (required)
     * @return ApiResponse&lt;DeleteKeychainEntryResponseV1&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<DeleteKeychainEntryResponseV1> deleteKeychainEntryV1WithHttpInfo(DeleteKeychainEntryRequestV1 deleteKeychainEntryRequestV1) throws ApiException {
        okhttp3.Call localVarCall = deleteKeychainEntryV1ValidateBeforeCall(deleteKeychainEntryRequestV1, null);
        Type localVarReturnType = new TypeToken<DeleteKeychainEntryResponseV1>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Deletes an entry under a key on the keychain backend. (asynchronously)
     * 
     * @param deleteKeychainEntryRequestV1 Request body to delete a keychain entry via its key (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call deleteKeychainEntryV1Async(DeleteKeychainEntryRequestV1 deleteKeychainEntryRequestV1, final ApiCallback<DeleteKeychainEntryResponseV1> _callback) throws ApiException {

        okhttp3.Call localVarCall = deleteKeychainEntryV1ValidateBeforeCall(deleteKeychainEntryRequestV1, _callback);
        Type localVarReturnType = new TypeToken<DeleteKeychainEntryResponseV1>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getKeychainEntryV1
     * @param getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> A keychain item with the specified key was not found. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getKeychainEntryV1Call(GetKeychainEntryRequestV1 getKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = getKeychainEntryRequestV1;

        // create path and map variables
        String localVarPath = "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/get-keychain-entry";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getKeychainEntryV1ValidateBeforeCall(GetKeychainEntryRequestV1 getKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'getKeychainEntryRequestV1' is set
        if (getKeychainEntryRequestV1 == null) {
            throw new ApiException("Missing the required parameter 'getKeychainEntryRequestV1' when calling getKeychainEntryV1(Async)");
        }

        return getKeychainEntryV1Call(getKeychainEntryRequestV1, _callback);

    }

    /**
     * Retrieves the contents of a keychain entry from the backend.
     * 
     * @param getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key (required)
     * @return GetKeychainEntryResponseV1
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> A keychain item with the specified key was not found. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public GetKeychainEntryResponseV1 getKeychainEntryV1(GetKeychainEntryRequestV1 getKeychainEntryRequestV1) throws ApiException {
        ApiResponse<GetKeychainEntryResponseV1> localVarResp = getKeychainEntryV1WithHttpInfo(getKeychainEntryRequestV1);
        return localVarResp.getData();
    }

    /**
     * Retrieves the contents of a keychain entry from the backend.
     * 
     * @param getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key (required)
     * @return ApiResponse&lt;GetKeychainEntryResponseV1&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> A keychain item with the specified key was not found. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<GetKeychainEntryResponseV1> getKeychainEntryV1WithHttpInfo(GetKeychainEntryRequestV1 getKeychainEntryRequestV1) throws ApiException {
        okhttp3.Call localVarCall = getKeychainEntryV1ValidateBeforeCall(getKeychainEntryRequestV1, null);
        Type localVarReturnType = new TypeToken<GetKeychainEntryResponseV1>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Retrieves the contents of a keychain entry from the backend. (asynchronously)
     * 
     * @param getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> A keychain item with the specified key was not found. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getKeychainEntryV1Async(GetKeychainEntryRequestV1 getKeychainEntryRequestV1, final ApiCallback<GetKeychainEntryResponseV1> _callback) throws ApiException {

        okhttp3.Call localVarCall = getKeychainEntryV1ValidateBeforeCall(getKeychainEntryRequestV1, _callback);
        Type localVarReturnType = new TypeToken<GetKeychainEntryResponseV1>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for getPrometheusMetricsV1
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getPrometheusMetricsV1Call(final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/get-prometheus-exporter-metrics";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "text/plain"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(basePath, localVarPath, "GET", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call getPrometheusMetricsV1ValidateBeforeCall(final ApiCallback _callback) throws ApiException {
        return getPrometheusMetricsV1Call(_callback);

    }

    /**
     * Get the Prometheus Metrics
     * 
     * @return String
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
     </table>
     */
    public String getPrometheusMetricsV1() throws ApiException {
        ApiResponse<String> localVarResp = getPrometheusMetricsV1WithHttpInfo();
        return localVarResp.getData();
    }

    /**
     * Get the Prometheus Metrics
     * 
     * @return ApiResponse&lt;String&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<String> getPrometheusMetricsV1WithHttpInfo() throws ApiException {
        okhttp3.Call localVarCall = getPrometheusMetricsV1ValidateBeforeCall(null);
        Type localVarReturnType = new TypeToken<String>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Get the Prometheus Metrics (asynchronously)
     * 
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call getPrometheusMetricsV1Async(final ApiCallback<String> _callback) throws ApiException {

        okhttp3.Call localVarCall = getPrometheusMetricsV1ValidateBeforeCall(_callback);
        Type localVarReturnType = new TypeToken<String>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for hasKeychainEntryV1
     * @param hasKeychainEntryRequestV1 Request body for checking a keychain entry via its key (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call hasKeychainEntryV1Call(HasKeychainEntryRequestV1 hasKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = hasKeychainEntryRequestV1;

        // create path and map variables
        String localVarPath = "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/has-keychain-entry";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call hasKeychainEntryV1ValidateBeforeCall(HasKeychainEntryRequestV1 hasKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'hasKeychainEntryRequestV1' is set
        if (hasKeychainEntryRequestV1 == null) {
            throw new ApiException("Missing the required parameter 'hasKeychainEntryRequestV1' when calling hasKeychainEntryV1(Async)");
        }

        return hasKeychainEntryV1Call(hasKeychainEntryRequestV1, _callback);

    }

    /**
     * Checks that an entry exists under a key on the keychain backend
     * 
     * @param hasKeychainEntryRequestV1 Request body for checking a keychain entry via its key (required)
     * @return HasKeychainEntryResponseV1
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public HasKeychainEntryResponseV1 hasKeychainEntryV1(HasKeychainEntryRequestV1 hasKeychainEntryRequestV1) throws ApiException {
        ApiResponse<HasKeychainEntryResponseV1> localVarResp = hasKeychainEntryV1WithHttpInfo(hasKeychainEntryRequestV1);
        return localVarResp.getData();
    }

    /**
     * Checks that an entry exists under a key on the keychain backend
     * 
     * @param hasKeychainEntryRequestV1 Request body for checking a keychain entry via its key (required)
     * @return ApiResponse&lt;HasKeychainEntryResponseV1&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<HasKeychainEntryResponseV1> hasKeychainEntryV1WithHttpInfo(HasKeychainEntryRequestV1 hasKeychainEntryRequestV1) throws ApiException {
        okhttp3.Call localVarCall = hasKeychainEntryV1ValidateBeforeCall(hasKeychainEntryRequestV1, null);
        Type localVarReturnType = new TypeToken<HasKeychainEntryResponseV1>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Checks that an entry exists under a key on the keychain backend (asynchronously)
     * 
     * @param hasKeychainEntryRequestV1 Request body for checking a keychain entry via its key (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call hasKeychainEntryV1Async(HasKeychainEntryRequestV1 hasKeychainEntryRequestV1, final ApiCallback<HasKeychainEntryResponseV1> _callback) throws ApiException {

        okhttp3.Call localVarCall = hasKeychainEntryV1ValidateBeforeCall(hasKeychainEntryRequestV1, _callback);
        Type localVarReturnType = new TypeToken<HasKeychainEntryResponseV1>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
    /**
     * Build call for setKeychainEntryV1
     * @param setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key (required)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call setKeychainEntryV1Call(SetKeychainEntryRequestV1 setKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        String basePath = null;
        // Operation Servers
        String[] localBasePaths = new String[] {  };

        // Determine Base Path to Use
        if (localCustomBaseUrl != null){
            basePath = localCustomBaseUrl;
        } else if ( localBasePaths.length > 0 ) {
            basePath = localBasePaths[localHostIndex];
        } else {
            basePath = null;
        }

        Object localVarPostBody = setKeychainEntryRequestV1;

        // create path and map variables
        String localVarPath = "/api/v1/plugins/@hyperledger/cactus-plugin-keychain-memory/set-keychain-entry";

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "application/json"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        if (localVarContentType != null) {
            localVarHeaderParams.put("Content-Type", localVarContentType);
        }

        String[] localVarAuthNames = new String[] {  };
        return localVarApiClient.buildCall(basePath, localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call setKeychainEntryV1ValidateBeforeCall(SetKeychainEntryRequestV1 setKeychainEntryRequestV1, final ApiCallback _callback) throws ApiException {
        // verify the required parameter 'setKeychainEntryRequestV1' is set
        if (setKeychainEntryRequestV1 == null) {
            throw new ApiException("Missing the required parameter 'setKeychainEntryRequestV1' when calling setKeychainEntryV1(Async)");
        }

        return setKeychainEntryV1Call(setKeychainEntryRequestV1, _callback);

    }

    /**
     * Sets a value under a key on the keychain backend.
     * 
     * @param setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key (required)
     * @return SetKeychainEntryResponseV1
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public SetKeychainEntryResponseV1 setKeychainEntryV1(SetKeychainEntryRequestV1 setKeychainEntryRequestV1) throws ApiException {
        ApiResponse<SetKeychainEntryResponseV1> localVarResp = setKeychainEntryV1WithHttpInfo(setKeychainEntryRequestV1);
        return localVarResp.getData();
    }

    /**
     * Sets a value under a key on the keychain backend.
     * 
     * @param setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key (required)
     * @return ApiResponse&lt;SetKeychainEntryResponseV1&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<SetKeychainEntryResponseV1> setKeychainEntryV1WithHttpInfo(SetKeychainEntryRequestV1 setKeychainEntryRequestV1) throws ApiException {
        okhttp3.Call localVarCall = setKeychainEntryV1ValidateBeforeCall(setKeychainEntryRequestV1, null);
        Type localVarReturnType = new TypeToken<SetKeychainEntryResponseV1>(){}.getType();
        return localVarApiClient.execute(localVarCall, localVarReturnType);
    }

    /**
     * Sets a value under a key on the keychain backend. (asynchronously)
     * 
     * @param setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key (required)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> OK </td><td>  -  </td></tr>
        <tr><td> 400 </td><td> Bad request. Key must be a string and longer than 0, shorter than 1024 characters. </td><td>  -  </td></tr>
        <tr><td> 401 </td><td> Authorization information is missing or invalid. </td><td>  -  </td></tr>
        <tr><td> 500 </td><td> Unexpected error. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call setKeychainEntryV1Async(SetKeychainEntryRequestV1 setKeychainEntryRequestV1, final ApiCallback<SetKeychainEntryResponseV1> _callback) throws ApiException {

        okhttp3.Call localVarCall = setKeychainEntryV1ValidateBeforeCall(setKeychainEntryRequestV1, _callback);
        Type localVarReturnType = new TypeToken<SetKeychainEntryResponseV1>(){}.getType();
        localVarApiClient.executeAsync(localVarCall, localVarReturnType, _callback);
        return localVarCall;
    }
}
