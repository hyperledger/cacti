/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.apis

import java.io.IOException
import okhttp3.OkHttpClient
import okhttp3.HttpUrl

import org.openapitools.client.models.DaoTokenGetAllowanceNotFoundResponse
import org.openapitools.client.models.DaoTokenGetAllowanceRequest
import org.openapitools.client.models.DaoTokenGetAllowanceResponse
import org.openapitools.client.models.EnrollAdminV1Request
import org.openapitools.client.models.EnrollAdminV1Response

import com.squareup.moshi.Json

import org.openapitools.client.infrastructure.ApiClient
import org.openapitools.client.infrastructure.ApiResponse
import org.openapitools.client.infrastructure.ClientException
import org.openapitools.client.infrastructure.ClientError
import org.openapitools.client.infrastructure.ServerException
import org.openapitools.client.infrastructure.ServerError
import org.openapitools.client.infrastructure.MultiValueMap
import org.openapitools.client.infrastructure.PartConfig
import org.openapitools.client.infrastructure.RequestConfig
import org.openapitools.client.infrastructure.RequestMethod
import org.openapitools.client.infrastructure.ResponseType
import org.openapitools.client.infrastructure.Success
import org.openapitools.client.infrastructure.toMultiValue

class DefaultApi(basePath: kotlin.String = defaultBasePath, client: OkHttpClient = ApiClient.defaultClient) : ApiClient(basePath, client) {
    companion object {
        @JvmStatic
        val defaultBasePath: String by lazy {
            System.getProperties().getProperty(ApiClient.baseUrlKey, "http://localhost")
        }
    }

    /**
     * Get the number of tokens &#x60;spender&#x60; is approved to spend on behalf of &#x60;account&#x60;
     * 
     * @param daoTokenGetAllowanceRequest  (optional)
     * @return DaoTokenGetAllowanceResponse
     * @throws IllegalStateException If the request is not correctly configured
     * @throws IOException Rethrows the OkHttp execute method exception
     * @throws UnsupportedOperationException If the API returns an informational or redirection response
     * @throws ClientException If the API returns a client error response
     * @throws ServerException If the API returns a server error response
     */
    @Suppress("UNCHECKED_CAST")
    @Throws(IllegalStateException::class, IOException::class, UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun daoTokenGetAllowanceV1(daoTokenGetAllowanceRequest: DaoTokenGetAllowanceRequest? = null) : DaoTokenGetAllowanceResponse {
        val localVarResponse = daoTokenGetAllowanceV1WithHttpInfo(daoTokenGetAllowanceRequest = daoTokenGetAllowanceRequest)

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as DaoTokenGetAllowanceResponse
            ResponseType.Informational -> throw UnsupportedOperationException("Client does not support Informational responses.")
            ResponseType.Redirection -> throw UnsupportedOperationException("Client does not support Redirection responses.")
            ResponseType.ClientError -> {
                val localVarError = localVarResponse as ClientError<*>
                throw ClientException("Client error : ${localVarError.statusCode} ${localVarError.message.orEmpty()}", localVarError.statusCode, localVarResponse)
            }
            ResponseType.ServerError -> {
                val localVarError = localVarResponse as ServerError<*>
                throw ServerException("Server error : ${localVarError.statusCode} ${localVarError.message.orEmpty()}", localVarError.statusCode, localVarResponse)
            }
        }
    }

    /**
     * Get the number of tokens &#x60;spender&#x60; is approved to spend on behalf of &#x60;account&#x60;
     * 
     * @param daoTokenGetAllowanceRequest  (optional)
     * @return ApiResponse<DaoTokenGetAllowanceResponse?>
     * @throws IllegalStateException If the request is not correctly configured
     * @throws IOException Rethrows the OkHttp execute method exception
     */
    @Suppress("UNCHECKED_CAST")
    @Throws(IllegalStateException::class, IOException::class)
    fun daoTokenGetAllowanceV1WithHttpInfo(daoTokenGetAllowanceRequest: DaoTokenGetAllowanceRequest?) : ApiResponse<DaoTokenGetAllowanceResponse?> {
        val localVariableConfig = daoTokenGetAllowanceV1RequestConfig(daoTokenGetAllowanceRequest = daoTokenGetAllowanceRequest)

        return request<DaoTokenGetAllowanceRequest, DaoTokenGetAllowanceResponse>(
            localVariableConfig
        )
    }

    /**
     * To obtain the request config of the operation daoTokenGetAllowanceV1
     *
     * @param daoTokenGetAllowanceRequest  (optional)
     * @return RequestConfig
     */
    fun daoTokenGetAllowanceV1RequestConfig(daoTokenGetAllowanceRequest: DaoTokenGetAllowanceRequest?) : RequestConfig<DaoTokenGetAllowanceRequest> {
        val localVariableBody = daoTokenGetAllowanceRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()
        localVariableHeaders["Content-Type"] = "application/json"
        localVariableHeaders["Accept"] = "application/json"

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-example-carbon-accounting-backend/dao-token/get-allowance",
            query = localVariableQuery,
            headers = localVariableHeaders,
            requiresAuthentication = false,
            body = localVariableBody
        )
    }

    /**
     * Registers an admin account within the Fabric organization specified.
     * 
     * @param enrollAdminV1Request  (optional)
     * @return EnrollAdminV1Response
     * @throws IllegalStateException If the request is not correctly configured
     * @throws IOException Rethrows the OkHttp execute method exception
     * @throws UnsupportedOperationException If the API returns an informational or redirection response
     * @throws ClientException If the API returns a client error response
     * @throws ServerException If the API returns a server error response
     */
    @Suppress("UNCHECKED_CAST")
    @Throws(IllegalStateException::class, IOException::class, UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun enrollAdminV1(enrollAdminV1Request: EnrollAdminV1Request? = null) : EnrollAdminV1Response {
        val localVarResponse = enrollAdminV1WithHttpInfo(enrollAdminV1Request = enrollAdminV1Request)

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as EnrollAdminV1Response
            ResponseType.Informational -> throw UnsupportedOperationException("Client does not support Informational responses.")
            ResponseType.Redirection -> throw UnsupportedOperationException("Client does not support Redirection responses.")
            ResponseType.ClientError -> {
                val localVarError = localVarResponse as ClientError<*>
                throw ClientException("Client error : ${localVarError.statusCode} ${localVarError.message.orEmpty()}", localVarError.statusCode, localVarResponse)
            }
            ResponseType.ServerError -> {
                val localVarError = localVarResponse as ServerError<*>
                throw ServerException("Server error : ${localVarError.statusCode} ${localVarError.message.orEmpty()}", localVarError.statusCode, localVarResponse)
            }
        }
    }

    /**
     * Registers an admin account within the Fabric organization specified.
     * 
     * @param enrollAdminV1Request  (optional)
     * @return ApiResponse<EnrollAdminV1Response?>
     * @throws IllegalStateException If the request is not correctly configured
     * @throws IOException Rethrows the OkHttp execute method exception
     */
    @Suppress("UNCHECKED_CAST")
    @Throws(IllegalStateException::class, IOException::class)
    fun enrollAdminV1WithHttpInfo(enrollAdminV1Request: EnrollAdminV1Request?) : ApiResponse<EnrollAdminV1Response?> {
        val localVariableConfig = enrollAdminV1RequestConfig(enrollAdminV1Request = enrollAdminV1Request)

        return request<EnrollAdminV1Request, EnrollAdminV1Response>(
            localVariableConfig
        )
    }

    /**
     * To obtain the request config of the operation enrollAdminV1
     *
     * @param enrollAdminV1Request  (optional)
     * @return RequestConfig
     */
    fun enrollAdminV1RequestConfig(enrollAdminV1Request: EnrollAdminV1Request?) : RequestConfig<EnrollAdminV1Request> {
        val localVariableBody = enrollAdminV1Request
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()
        localVariableHeaders["Content-Type"] = "application/json"
        localVariableHeaders["Accept"] = "application/json"

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/utilityemissionchannel/registerEnroll/admin",
            query = localVariableQuery,
            headers = localVariableHeaders,
            requiresAuthentication = false,
            body = localVariableBody
        )
    }


    private fun encodeURIComponent(uriComponent: kotlin.String): kotlin.String =
        HttpUrl.Builder().scheme("http").host("localhost").addPathSegment(uriComponent).build().encodedPathSegments[0]
}