/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - BUNGEE-Hermes
 * Can create blockchain views of different networks
 *
 * The version of the OpenAPI document: 2.0.0-alpha.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from './base';

/**
 * Request object for createViewV1 endpoint
 * @export
 * @interface CreateViewRequest
 */
export interface CreateViewRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof CreateViewRequest
     */
    'stateIds'?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequest
     */
    'tI'?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequest
     */
    'tF'?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequest
     */
    'viewID'?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequest
     */
    'strategyId': string;
    /**
     * 
     * @type {CreateViewRequestNetworkDetails}
     * @memberof CreateViewRequest
     */
    'networkDetails': CreateViewRequestNetworkDetails;
}
/**
 * 
 * @export
 * @interface CreateViewRequestNetworkDetails
 */
export interface CreateViewRequestNetworkDetails {
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequestNetworkDetails
     */
    'connectorApiPath': string;
    /**
     * 
     * @type {string}
     * @memberof CreateViewRequestNetworkDetails
     */
    'participant': string;
}
/**
 * This is the response for a viewRequests
 * @export
 * @interface CreateViewResponse
 */
export interface CreateViewResponse {
    /**
     * 
     * @type {string}
     * @memberof CreateViewResponse
     */
    'view'?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateViewResponse
     */
    'signature'?: string;
}
/**
 * public key from bungee-hermes plugin instance
 * @export
 * @interface GetPublicKeyResponse
 */
export interface GetPublicKeyResponse {
    /**
     * 
     * @type {string}
     * @memberof GetPublicKeyResponse
     */
    'pubKey'?: string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Creates a Blockchain View.
         * @param {CreateViewRequest} createViewRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createViewV1: async (createViewRequest: CreateViewRequest, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'createViewRequest' is not null or undefined
            assertParamExists('createViewV1', 'createViewRequest', createViewRequest)
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/create-view`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(createViewRequest, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Queries plugin\'s available strategies for ledger capture
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAvailableStrategies: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/get-available-strategies`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Queries plugin\'s public key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPublicKey: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-bungee-hermes/get-public-key`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * DefaultApi - functional programming interface
 * @export
 */
export const DefaultApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = DefaultApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Creates a Blockchain View.
         * @param {CreateViewRequest} createViewRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createViewV1(createViewRequest: CreateViewRequest, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<CreateViewResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.createViewV1(createViewRequest, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Queries plugin\'s available strategies for ledger capture
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAvailableStrategies(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<string>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getAvailableStrategies(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Queries plugin\'s public key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPublicKey(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetPublicKeyResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPublicKey(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * DefaultApi - factory interface
 * @export
 */
export const DefaultApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = DefaultApiFp(configuration)
    return {
        /**
         * 
         * @summary Creates a Blockchain View.
         * @param {CreateViewRequest} createViewRequest 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createViewV1(createViewRequest: CreateViewRequest, options?: any): AxiosPromise<CreateViewResponse> {
            return localVarFp.createViewV1(createViewRequest, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Queries plugin\'s available strategies for ledger capture
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAvailableStrategies(options?: any): AxiosPromise<Array<string>> {
            return localVarFp.getAvailableStrategies(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Queries plugin\'s public key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPublicKey(options?: any): AxiosPromise<GetPublicKeyResponse> {
            return localVarFp.getPublicKey(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * DefaultApi - object-oriented interface
 * @export
 * @class DefaultApi
 * @extends {BaseAPI}
 */
export class DefaultApi extends BaseAPI {
    /**
     * 
     * @summary Creates a Blockchain View.
     * @param {CreateViewRequest} createViewRequest 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public createViewV1(createViewRequest: CreateViewRequest, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).createViewV1(createViewRequest, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Queries plugin\'s available strategies for ledger capture
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getAvailableStrategies(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getAvailableStrategies(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Queries plugin\'s public key
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPublicKey(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getPublicKey(options).then((request) => request(this.axios, this.basePath));
    }
}


