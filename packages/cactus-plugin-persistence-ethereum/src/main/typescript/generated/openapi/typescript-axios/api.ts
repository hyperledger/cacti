/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Persistence Ethereum
 * Synchronizes state of an ethereum ledger into a DB that can later be viewed in GUI
 *
 * The version of the OpenAPI document: v2.0.0-alpha.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from "./configuration.js";
import type { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from "./common.js";
import type { RequestArgs } from "./base.js";
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError } from "./base.js";

/**
 * 
 * @export
 * @interface ErrorExceptionResponseV1
 */
export interface ErrorExceptionResponseV1 {
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    'message': string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    'error': string;
}
/**
 * Ethereum tokens that are being monitored by the persistence plugin.
 * @export
 * @interface MonitoredToken
 */
export interface MonitoredToken {
    /**
     * 
     * @type {TokenTypeV1}
     * @memberof MonitoredToken
     */
    'type': TokenTypeV1;
    /**
     * Token name
     * @type {string}
     * @memberof MonitoredToken
     */
    'name': string;
    /**
     * Token symbol
     * @type {string}
     * @memberof MonitoredToken
     */
    'symbol': string;
}


/**
 * Response with plugin status report.
 * @export
 * @interface StatusResponseV1
 */
export interface StatusResponseV1 {
    /**
     * Plugin instance id.
     * @type {string}
     * @memberof StatusResponseV1
     */
    'instanceId': string;
    /**
     * True if successfully connected to the database, false otherwise.
     * @type {boolean}
     * @memberof StatusResponseV1
     */
    'connected': boolean;
    /**
     * True if web services were correctly exported.
     * @type {boolean}
     * @memberof StatusResponseV1
     */
    'webServicesRegistered': boolean;
    /**
     * Total number of tokens being monitored by the plugin.
     * @type {number}
     * @memberof StatusResponseV1
     */
    'monitoredTokensCount': number;
    /**
     * 
     * @type {Array<TrackedOperationV1>}
     * @memberof StatusResponseV1
     */
    'operationsRunning': Array<TrackedOperationV1>;
    /**
     * True if block monitoring is running, false otherwise.
     * @type {boolean}
     * @memberof StatusResponseV1
     */
    'monitorRunning': boolean;
    /**
     * Number of the last block seen by the block monitor.
     * @type {number}
     * @memberof StatusResponseV1
     */
    'lastSeenBlock': number;
}
/**
 * 
 * @export
 * @enum {string}
 */

export const TokenTypeV1 = {
    /**
    * EIP-20: Token Standard
    */
    ERC20: 'erc20',
    /**
    * EIP-721: Non-Fungible Token Standard
    */
    ERC721: 'erc721'
} as const;

export type TokenTypeV1 = typeof TokenTypeV1[keyof typeof TokenTypeV1];


/**
 * Persistence plugin operation that is tracked and returned in status report.
 * @export
 * @interface TrackedOperationV1
 */
export interface TrackedOperationV1 {
    /**
     * Start time of the operation.
     * @type {string}
     * @memberof TrackedOperationV1
     */
    'startAt': string;
    /**
     * Operation name.
     * @type {string}
     * @memberof TrackedOperationV1
     */
    'operation': string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get the status of persistence plugin for ethereum
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getStatusV1: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-persistence-ethereum/status`;
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
         * @summary Get the status of persistence plugin for ethereum
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getStatusV1(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<StatusResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getStatusV1(options);
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
         * @summary Get the status of persistence plugin for ethereum
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getStatusV1(options?: any): AxiosPromise<StatusResponseV1> {
            return localVarFp.getStatusV1(options).then((request) => request(axios, basePath));
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
     * @summary Get the status of persistence plugin for ethereum
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getStatusV1(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getStatusV1(options).then((request) => request(this.axios, this.basePath));
    }
}


