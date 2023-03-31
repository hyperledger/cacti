/**
 * Hyperledger Cactus Plugin - HTLC ETH BESU ERC20
 *
 * Allows Cactus nodes to interact with HTLC contracts with ERC-20 Tokens
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.apis

import org.openapitools.client.models.GetSingleStatusRequest
import org.openapitools.client.models.GetStatusRequest
import org.openapitools.client.models.InitializeRequest
import org.openapitools.client.models.InvokeContractV1Response
import org.openapitools.client.models.NewContractRequest
import org.openapitools.client.models.RefundRequest
import org.openapitools.client.models.RunTransactionResponse
import org.openapitools.client.models.WithdrawRequest

import org.openapitools.client.infrastructure.ApiClient
import org.openapitools.client.infrastructure.ClientException
import org.openapitools.client.infrastructure.ClientError
import org.openapitools.client.infrastructure.ServerException
import org.openapitools.client.infrastructure.ServerError
import org.openapitools.client.infrastructure.MultiValueMap
import org.openapitools.client.infrastructure.RequestConfig
import org.openapitools.client.infrastructure.RequestMethod
import org.openapitools.client.infrastructure.ResponseType
import org.openapitools.client.infrastructure.Success
import org.openapitools.client.infrastructure.toMultiValue

class DefaultApi(basePath: kotlin.String = defaultBasePath) : ApiClient(basePath) {
    companion object {
        @JvmStatic
        val defaultBasePath: String by lazy {
            System.getProperties().getProperty("org.openapitools.client.baseUrl", "http://localhost")
        }
    }

    /**
    * 
    * 
    * @param getSingleStatusRequest  (optional)
    * @return kotlin.Int
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun getSingleStatusV1(getSingleStatusRequest: GetSingleStatusRequest?) : kotlin.Int {
        val localVariableConfig = getSingleStatusV1RequestConfig(getSingleStatusRequest = getSingleStatusRequest)

        val localVarResponse = request<GetSingleStatusRequest, kotlin.Int>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as kotlin.Int
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
    * To obtain the request config of the operation getSingleStatusV1
    *
    * @param getSingleStatusRequest  (optional)
    * @return RequestConfig
    */
    fun getSingleStatusV1RequestConfig(getSingleStatusRequest: GetSingleStatusRequest?) : RequestConfig<GetSingleStatusRequest> {
        val localVariableBody = getSingleStatusRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/get-single-status",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

    /**
    * 
    * 
    * @param getStatusRequest  (optional)
    * @return kotlin.collections.List<kotlin.Int>
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun getStatusV1(getStatusRequest: GetStatusRequest?) : kotlin.collections.List<kotlin.Int> {
        val localVariableConfig = getStatusV1RequestConfig(getStatusRequest = getStatusRequest)

        val localVarResponse = request<GetStatusRequest, kotlin.collections.List<kotlin.Int>>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as kotlin.collections.List<kotlin.Int>
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
    * To obtain the request config of the operation getStatusV1
    *
    * @param getStatusRequest  (optional)
    * @return RequestConfig
    */
    fun getStatusV1RequestConfig(getStatusRequest: GetStatusRequest?) : RequestConfig<GetStatusRequest> {
        val localVariableBody = getStatusRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/get-status",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

    /**
    * Initialize contract
    * 
    * @param initializeRequest  (optional)
    * @return RunTransactionResponse
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun initializeV1(initializeRequest: InitializeRequest?) : RunTransactionResponse {
        val localVariableConfig = initializeV1RequestConfig(initializeRequest = initializeRequest)

        val localVarResponse = request<InitializeRequest, RunTransactionResponse>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as RunTransactionResponse
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
    * To obtain the request config of the operation initializeV1
    *
    * @param initializeRequest  (optional)
    * @return RequestConfig
    */
    fun initializeV1RequestConfig(initializeRequest: InitializeRequest?) : RequestConfig<InitializeRequest> {
        val localVariableBody = initializeRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/initialize",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

    /**
    * Create a new hashtimelock contract
    * 
    * @param newContractRequest  (optional)
    * @return InvokeContractV1Response
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun newContractV1(newContractRequest: NewContractRequest?) : InvokeContractV1Response {
        val localVariableConfig = newContractV1RequestConfig(newContractRequest = newContractRequest)

        val localVarResponse = request<NewContractRequest, InvokeContractV1Response>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as InvokeContractV1Response
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
    * To obtain the request config of the operation newContractV1
    *
    * @param newContractRequest  (optional)
    * @return RequestConfig
    */
    fun newContractV1RequestConfig(newContractRequest: NewContractRequest?) : RequestConfig<NewContractRequest> {
        val localVariableBody = newContractRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/new-contract",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

    /**
    * Refund a hashtimelock contract
    * 
    * @param refundRequest  (optional)
    * @return InvokeContractV1Response
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun refundV1(refundRequest: RefundRequest?) : InvokeContractV1Response {
        val localVariableConfig = refundV1RequestConfig(refundRequest = refundRequest)

        val localVarResponse = request<RefundRequest, InvokeContractV1Response>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as InvokeContractV1Response
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
    * To obtain the request config of the operation refundV1
    *
    * @param refundRequest  (optional)
    * @return RequestConfig
    */
    fun refundV1RequestConfig(refundRequest: RefundRequest?) : RequestConfig<RefundRequest> {
        val localVariableBody = refundRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/refund",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

    /**
    * Withdraw a hashtimelock contract
    * 
    * @param withdrawRequest  (optional)
    * @return InvokeContractV1Response
    * @throws UnsupportedOperationException If the API returns an informational or redirection response
    * @throws ClientException If the API returns a client error response
    * @throws ServerException If the API returns a server error response
    */
    @Suppress("UNCHECKED_CAST")
    @Throws(UnsupportedOperationException::class, ClientException::class, ServerException::class)
    fun withdrawV1(withdrawRequest: WithdrawRequest?) : InvokeContractV1Response {
        val localVariableConfig = withdrawV1RequestConfig(withdrawRequest = withdrawRequest)

        val localVarResponse = request<WithdrawRequest, InvokeContractV1Response>(
            localVariableConfig
        )

        return when (localVarResponse.responseType) {
            ResponseType.Success -> (localVarResponse as Success<*>).data as InvokeContractV1Response
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
    * To obtain the request config of the operation withdrawV1
    *
    * @param withdrawRequest  (optional)
    * @return RequestConfig
    */
    fun withdrawV1RequestConfig(withdrawRequest: WithdrawRequest?) : RequestConfig<WithdrawRequest> {
        val localVariableBody = withdrawRequest
        val localVariableQuery: MultiValueMap = mutableMapOf()
        val localVariableHeaders: MutableMap<String, String> = mutableMapOf()

        return RequestConfig(
            method = RequestMethod.POST,
            path = "/api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu-erc20/withdraw",
            query = localVariableQuery,
            headers = localVariableHeaders,
            body = localVariableBody
        )
    }

}
