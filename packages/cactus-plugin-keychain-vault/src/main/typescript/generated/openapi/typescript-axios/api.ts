/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus - Keychain API
 * Contains/describes the Keychain API types/paths for Hyperledger Cactus.
 *
 * The version of the OpenAPI document: 2.0.0-rc.3
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
 * 
 * @export
 * @interface DeleteKeychainEntryRequestV1
 */
export interface DeleteKeychainEntryRequestV1 {
    /**
     * The key of the entry to delete from the keychain.
     * @type {string}
     * @memberof DeleteKeychainEntryRequestV1
     */
    'key': string;
}
/**
 * 
 * @export
 * @interface DeleteKeychainEntryResponseV1
 */
export interface DeleteKeychainEntryResponseV1 {
    /**
     * The key of the entry that was deleted from the keychain.
     * @type {string}
     * @memberof DeleteKeychainEntryResponseV1
     */
    'key': string;
}
/**
 * 
 * @export
 * @interface GetKeychainEntryRequestV1
 */
export interface GetKeychainEntryRequestV1 {
    /**
     * The key for the entry to get from the keychain.
     * @type {string}
     * @memberof GetKeychainEntryRequestV1
     */
    'key': string;
}
/**
 * 
 * @export
 * @interface GetKeychainEntryResponseV1
 */
export interface GetKeychainEntryResponseV1 {
    /**
     * The key that was used to retrieve the value from the keychain.
     * @type {string}
     * @memberof GetKeychainEntryResponseV1
     */
    'key': string;
    /**
     * The value associated with the requested key on the keychain.
     * @type {string}
     * @memberof GetKeychainEntryResponseV1
     */
    'value': string;
}
/**
 * 
 * @export
 * @interface HasKeychainEntryRequestV1
 */
export interface HasKeychainEntryRequestV1 {
    /**
     * The key to check for presence in the keychain.
     * @type {string}
     * @memberof HasKeychainEntryRequestV1
     */
    'key': string;
}
/**
 * 
 * @export
 * @interface HasKeychainEntryResponseV1
 */
export interface HasKeychainEntryResponseV1 {
    /**
     * The key that was used to check the presence of the value in the keychain.
     * @type {string}
     * @memberof HasKeychainEntryResponseV1
     */
    'key': string;
    /**
     * Date and time encoded as JSON when the presence check was performed by the plugin backend.
     * @type {string}
     * @memberof HasKeychainEntryResponseV1
     */
    'checkedAt': string;
    /**
     * The boolean true or false indicating the presence or absence of an entry under \'key\'.
     * @type {boolean}
     * @memberof HasKeychainEntryResponseV1
     */
    'isPresent': boolean;
}
/**
 * 
 * @export
 * @interface SetKeychainEntryRequestV1
 */
export interface SetKeychainEntryRequestV1 {
    /**
     * The key for the entry to set on the keychain.
     * @type {string}
     * @memberof SetKeychainEntryRequestV1
     */
    'key': string;
    /**
     * The value that will be associated with the key on the keychain.
     * @type {string}
     * @memberof SetKeychainEntryRequestV1
     */
    'value': string;
}
/**
 * 
 * @export
 * @interface SetKeychainEntryResponseV1
 */
export interface SetKeychainEntryResponseV1 {
    /**
     * The key that was used to set the value on the keychain.
     * @type {string}
     * @memberof SetKeychainEntryResponseV1
     */
    'key': string;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Deletes an entry from the keychain stored under the provided key.
         * @param {DeleteKeychainEntryRequestV1} [deleteKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteKeychainEntryV1: async (deleteKeychainEntryRequestV1?: DeleteKeychainEntryRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-keychain-vault/delete-keychain-entry`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(deleteKeychainEntryRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Retrieves the contents of a keychain entry from the backend.
         * @param {GetKeychainEntryRequestV1} getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getKeychainEntryV1: async (getKeychainEntryRequestV1: GetKeychainEntryRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'getKeychainEntryRequestV1' is not null or undefined
            assertParamExists('getKeychainEntryV1', 'getKeychainEntryRequestV1', getKeychainEntryRequestV1)
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-keychain-vault/get-keychain-entry`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(getKeychainEntryRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusMetricsV1: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-keychain-vault/get-prometheus-exporter-metrics`;
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
         * @summary Retrieves the information regarding a key being present on the keychain or not.
         * @param {HasKeychainEntryRequestV1} [hasKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        hasKeychainEntryV1: async (hasKeychainEntryRequestV1?: HasKeychainEntryRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-keychain-vault/has-keychain-entry`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(hasKeychainEntryRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Sets a value under a key on the keychain backend.
         * @param {SetKeychainEntryRequestV1} setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        setKeychainEntryV1: async (setKeychainEntryRequestV1: SetKeychainEntryRequestV1, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'setKeychainEntryRequestV1' is not null or undefined
            assertParamExists('setKeychainEntryV1', 'setKeychainEntryRequestV1', setKeychainEntryRequestV1)
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-keychain-vault/set-keychain-entry`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(setKeychainEntryRequestV1, localVarRequestOptions, configuration)

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
         * @summary Deletes an entry from the keychain stored under the provided key.
         * @param {DeleteKeychainEntryRequestV1} [deleteKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async deleteKeychainEntryV1(deleteKeychainEntryRequestV1?: DeleteKeychainEntryRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<DeleteKeychainEntryResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.deleteKeychainEntryV1(deleteKeychainEntryRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Retrieves the contents of a keychain entry from the backend.
         * @param {GetKeychainEntryRequestV1} getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getKeychainEntryV1(getKeychainEntryRequestV1: GetKeychainEntryRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetKeychainEntryResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getKeychainEntryV1(getKeychainEntryRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPrometheusMetricsV1(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPrometheusMetricsV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Retrieves the information regarding a key being present on the keychain or not.
         * @param {HasKeychainEntryRequestV1} [hasKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async hasKeychainEntryV1(hasKeychainEntryRequestV1?: HasKeychainEntryRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<HasKeychainEntryResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.hasKeychainEntryV1(hasKeychainEntryRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Sets a value under a key on the keychain backend.
         * @param {SetKeychainEntryRequestV1} setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async setKeychainEntryV1(setKeychainEntryRequestV1: SetKeychainEntryRequestV1, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<SetKeychainEntryResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.setKeychainEntryV1(setKeychainEntryRequestV1, options);
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
         * @summary Deletes an entry from the keychain stored under the provided key.
         * @param {DeleteKeychainEntryRequestV1} [deleteKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        deleteKeychainEntryV1(deleteKeychainEntryRequestV1?: DeleteKeychainEntryRequestV1, options?: any): AxiosPromise<DeleteKeychainEntryResponseV1> {
            return localVarFp.deleteKeychainEntryV1(deleteKeychainEntryRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Retrieves the contents of a keychain entry from the backend.
         * @param {GetKeychainEntryRequestV1} getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getKeychainEntryV1(getKeychainEntryRequestV1: GetKeychainEntryRequestV1, options?: any): AxiosPromise<GetKeychainEntryResponseV1> {
            return localVarFp.getKeychainEntryV1(getKeychainEntryRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getPrometheusMetricsV1(options?: any): AxiosPromise<string> {
            return localVarFp.getPrometheusMetricsV1(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Retrieves the information regarding a key being present on the keychain or not.
         * @param {HasKeychainEntryRequestV1} [hasKeychainEntryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        hasKeychainEntryV1(hasKeychainEntryRequestV1?: HasKeychainEntryRequestV1, options?: any): AxiosPromise<HasKeychainEntryResponseV1> {
            return localVarFp.hasKeychainEntryV1(hasKeychainEntryRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Sets a value under a key on the keychain backend.
         * @param {SetKeychainEntryRequestV1} setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        setKeychainEntryV1(setKeychainEntryRequestV1: SetKeychainEntryRequestV1, options?: any): AxiosPromise<SetKeychainEntryResponseV1> {
            return localVarFp.setKeychainEntryV1(setKeychainEntryRequestV1, options).then((request) => request(axios, basePath));
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
     * @summary Deletes an entry from the keychain stored under the provided key.
     * @param {DeleteKeychainEntryRequestV1} [deleteKeychainEntryRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public deleteKeychainEntryV1(deleteKeychainEntryRequestV1?: DeleteKeychainEntryRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).deleteKeychainEntryV1(deleteKeychainEntryRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Retrieves the contents of a keychain entry from the backend.
     * @param {GetKeychainEntryRequestV1} getKeychainEntryRequestV1 Request body to obtain a keychain entry via its key
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getKeychainEntryV1(getKeychainEntryRequestV1: GetKeychainEntryRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getKeychainEntryV1(getKeychainEntryRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the Prometheus Metrics
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPrometheusMetricsV1(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getPrometheusMetricsV1(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Retrieves the information regarding a key being present on the keychain or not.
     * @param {HasKeychainEntryRequestV1} [hasKeychainEntryRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public hasKeychainEntryV1(hasKeychainEntryRequestV1?: HasKeychainEntryRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).hasKeychainEntryV1(hasKeychainEntryRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Sets a value under a key on the keychain backend.
     * @param {SetKeychainEntryRequestV1} setKeychainEntryRequestV1 Request body to write/update a keychain entry via its key
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public setKeychainEntryV1(setKeychainEntryRequestV1: SetKeychainEntryRequestV1, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).setKeychainEntryV1(setKeychainEntryRequestV1, options).then((request) => request(this.axios, this.basePath));
    }
}


