/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type TransactionApi interface {

	/*
	Cancel Cancel a transaction session

	Attempts to cancel a previously submitted transaction intent using its session ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCancelRequest
	*/
	Cancel(ctx context.Context) ApiCancelRequest

	// CancelExecute executes the request
	//  @return Cancel200Response
	CancelExecute(r ApiCancelRequest) (*Cancel200Response, *http.Response, error)

	/*
	GetRoutes Get a list of routes for a gateway-to-gateway asset transfer

	Get a list of possible routes for swapping one asset for another across multiple exchanges

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRoutesRequest
	*/
	GetRoutes(ctx context.Context) ApiGetRoutesRequest

	// GetRoutesExecute executes the request
	//  @return GetRoutes200Response
	GetRoutesExecute(r ApiGetRoutesRequest) (*GetRoutes200Response, *http.Response, error)

	/*
	Transact Submit a transaction intent

	Allows users to queue intents for transactions based on specified parameters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiTransactRequest
	*/
	Transact(ctx context.Context) ApiTransactRequest

	// TransactExecute executes the request
	//  @return Transact200Response
	TransactExecute(r ApiTransactRequest) (*Transact200Response, *http.Response, error)
}

// TransactionApiService TransactionApi service
type TransactionApiService service

type ApiCancelRequest struct {
	ctx context.Context
	ApiService TransactionApi
	cancelRequest *CancelRequest
}

func (r ApiCancelRequest) CancelRequest(cancelRequest CancelRequest) ApiCancelRequest {
	r.cancelRequest = &cancelRequest
	return r
}

func (r ApiCancelRequest) Execute() (*Cancel200Response, *http.Response, error) {
	return r.ApiService.CancelExecute(r)
}

/*
Cancel Cancel a transaction session

Attempts to cancel a previously submitted transaction intent using its session ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelRequest
*/
func (a *TransactionApiService) Cancel(ctx context.Context) ApiCancelRequest {
	return ApiCancelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Cancel200Response
func (a *TransactionApiService) CancelExecute(r ApiCancelRequest) (*Cancel200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Cancel200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionApiService.Cancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cancelRequest == nil {
		return localVarReturnValue, nil, reportError("cancelRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.cancelRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v TransactDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRoutesRequest struct {
	ctx context.Context
	ApiService TransactionApi
	fromNetworkID *string
	fromAmount *string
	fromToken *string
	toDLTNetwork *string
	toToken *string
	fromAddress *string
	toAddress *string
}

// The sending DLT Network.
func (r ApiGetRoutesRequest) FromNetworkID(fromNetworkID string) ApiGetRoutesRequest {
	r.fromNetworkID = &fromNetworkID
	return r
}

// The amount that should be sent including all decimals.
func (r ApiGetRoutesRequest) FromAmount(fromAmount string) ApiGetRoutesRequest {
	r.fromAmount = &fromAmount
	return r
}

// The token that should be transferred. Can be the address or the symbol.
func (r ApiGetRoutesRequest) FromToken(fromToken string) ApiGetRoutesRequest {
	r.fromToken = &fromToken
	return r
}

// The receiving DLT Network.
func (r ApiGetRoutesRequest) ToDLTNetwork(toDLTNetwork string) ApiGetRoutesRequest {
	r.toDLTNetwork = &toDLTNetwork
	return r
}

// The token that should be transferred to. Can be the address or the symbol.
func (r ApiGetRoutesRequest) ToToken(toToken string) ApiGetRoutesRequest {
	r.toToken = &toToken
	return r
}

// The sending wallet address.
func (r ApiGetRoutesRequest) FromAddress(fromAddress string) ApiGetRoutesRequest {
	r.fromAddress = &fromAddress
	return r
}

// The receiving wallet address. If none is provided, the fromAddress will be used.
func (r ApiGetRoutesRequest) ToAddress(toAddress string) ApiGetRoutesRequest {
	r.toAddress = &toAddress
	return r
}

func (r ApiGetRoutesRequest) Execute() (*GetRoutes200Response, *http.Response, error) {
	return r.ApiService.GetRoutesExecute(r)
}

/*
GetRoutes Get a list of routes for a gateway-to-gateway asset transfer

Get a list of possible routes for swapping one asset for another across multiple exchanges

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRoutesRequest
*/
func (a *TransactionApiService) GetRoutes(ctx context.Context) ApiGetRoutesRequest {
	return ApiGetRoutesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetRoutes200Response
func (a *TransactionApiService) GetRoutesExecute(r ApiGetRoutesRequest) (*GetRoutes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetRoutes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionApiService.GetRoutes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/routes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fromNetworkID == nil {
		return localVarReturnValue, nil, reportError("fromNetworkID is required and must be specified")
	}
	if r.fromAmount == nil {
		return localVarReturnValue, nil, reportError("fromAmount is required and must be specified")
	}
	if r.fromToken == nil {
		return localVarReturnValue, nil, reportError("fromToken is required and must be specified")
	}
	if r.toDLTNetwork == nil {
		return localVarReturnValue, nil, reportError("toDLTNetwork is required and must be specified")
	}
	if r.toToken == nil {
		return localVarReturnValue, nil, reportError("toToken is required and must be specified")
	}
	if r.fromAddress == nil {
		return localVarReturnValue, nil, reportError("fromAddress is required and must be specified")
	}
	if r.toAddress == nil {
		return localVarReturnValue, nil, reportError("toAddress is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fromNetworkID", r.fromNetworkID, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "fromAmount", r.fromAmount, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "fromToken", r.fromToken, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "toDLTNetwork", r.toDLTNetwork, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "toToken", r.toToken, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "fromAddress", r.fromAddress, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "toAddress", r.toAddress, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v TransactDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiTransactRequest struct {
	ctx context.Context
	ApiService TransactionApi
	transactRequest *TransactRequest
}

func (r ApiTransactRequest) TransactRequest(transactRequest TransactRequest) ApiTransactRequest {
	r.transactRequest = &transactRequest
	return r
}

func (r ApiTransactRequest) Execute() (*Transact200Response, *http.Response, error) {
	return r.ApiService.TransactExecute(r)
}

/*
Transact Submit a transaction intent

Allows users to queue intents for transactions based on specified parameters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTransactRequest
*/
func (a *TransactionApiService) Transact(ctx context.Context) ApiTransactRequest {
	return ApiTransactRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Transact200Response
func (a *TransactionApiService) TransactExecute(r ApiTransactRequest) (*Transact200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Transact200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionApiService.Transact")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/transact"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.transactRequest == nil {
		return localVarReturnValue, nil, reportError("transactRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.transactRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v TransactDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
