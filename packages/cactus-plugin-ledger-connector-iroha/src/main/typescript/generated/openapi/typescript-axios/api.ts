/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Connector Iroha
 * Can perform basic tasks on a Iroha ledger
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface ErrorExceptionJsonResponseV1
 */
export interface ErrorExceptionJsonResponseV1 {
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionJsonResponseV1
     */
    message: string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionJsonResponseV1
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionJsonResponseV1
     */
    error?: string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionJsonResponseV1
     */
    stack?: string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionJsonResponseV1
     */
    cause?: string;
}
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
    message: string;
    /**
     * 
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    error: string;
}
/**
 * 
 * @export
 * @interface GenerateTransactionRequestV1
 */
export interface GenerateTransactionRequestV1 {
    /**
     * Iroha command name.
     * @type {IrohaCommand}
     * @memberof GenerateTransactionRequestV1
     */
    commandName: IrohaCommand;
    /**
     * Parameters for iroha command specified in commandName
     * @type {object}
     * @memberof GenerateTransactionRequestV1
     */
    commandParams: object;
    /**
     * Sender account id
     * @type {string}
     * @memberof GenerateTransactionRequestV1
     */
    creatorAccountId: string;
    /**
     * Requested transaction quorum
     * @type {number}
     * @memberof GenerateTransactionRequestV1
     */
    quorum?: number;
}
/**
 * 
 * @export
 * @interface IrohaBaseConfig
 */
export interface IrohaBaseConfig {
    [key: string]: object | any;

    /**
     * 
     * @type {string}
     * @memberof IrohaBaseConfig
     */
    irohaHost?: string;
    /**
     * 
     * @type {number}
     * @memberof IrohaBaseConfig
     */
    irohaPort?: number;
    /**
     * 
     * @type {string}
     * @memberof IrohaBaseConfig
     */
    creatorAccountId?: string;
    /**
     * 
     * @type {Array<any>}
     * @memberof IrohaBaseConfig
     */
    privKey?: Array<any>;
    /**
     * 
     * @type {number}
     * @memberof IrohaBaseConfig
     */
    quorum?: number;
    /**
     * 
     * @type {number}
     * @memberof IrohaBaseConfig
     */
    timeoutLimit?: number;
    /**
     * Can only be set to false for an insecure grpc connection.
     * @type {boolean}
     * @memberof IrohaBaseConfig
     */
    tls?: boolean;
}
/**
 * 
 * @export
 * @enum {string}
 */

export enum IrohaCommand {
    /**
    * Make entity in the system, capable of sending transactions or queries, storing signatories, personal data and identifiers.
    */
    CreateAccount = 'createAccount',
    /**
    * Set key-value information for a given account.
    */
    SetAccountDetail = 'setAccountDetail',
    /**
    * Set the number of signatories required to confirm the identity of a user, who creates the transaction.
    */
    SetAccountQuorum = 'setAccountQuorum',
    /**
    * Set key-value information for a given account if the old value matches the value passed.
    */
    CompareAndSetAccountDetail = 'compareAndSetAccountDetail',
    /**
    * Create a new type of asset, unique in a domain. An asset is a countable representation of a commodity.
    */
    CreateAsset = 'createAsset',
    /**
    * Increase the quantity of an asset on account of transaction creator.
    */
    AddAssetQuantity = 'addAssetQuantity',
    /**
    * Decrease the number of assets on account of transaction creator.
    */
    SubtractAssetQuantity = 'subtractAssetQuantity',
    /**
    * Share assets within the account in peer network: in the way that source account transfers assets to the target account.
    */
    TransferAsset = 'transferAsset',
    /**
    * Make new domain in Iroha network, which is a group of accounts.
    */
    CreateDomain = 'createDomain',
    /**
    * Create a new role in the system from the set of permissions.
    */
    CreateRole = 'createRole',
    /**
    * Detach a role from the set of roles of an account.
    */
    DetachRole = 'detachRole',
    /**
    * Promote an account to some created role in the system, where a role is a set of permissions account has to perform an action (command or query).
    */
    AppendRole = 'appendRole',
    /**
    * Add an identifier to the account. Such identifier is a public key of another device or a public key of another user.
    */
    AddSignatory = 'addSignatory',
    /**
    * Remove a public key, associated with an identity, from an account
    */
    RemoveSignatory = 'removeSignatory',
    /**
    * Give another account rights to perform actions on the account of transaction sender (give someone right to do something with my account).
    */
    GrantPermission = 'grantPermission',
    /**
    * Revoke or dismiss given granted permission from another account in the network.
    */
    RevokePermission = 'revokePermission',
    /**
    * Write into ledger the fact of peer addition into the peer network.
    */
    AddPeer = 'addPeer',
    /**
    * Write into ledger the fact of peer removal from the network.
    */
    RemovePeer = 'removePeer',
    /**
    * This command is not available for use, it was added for backward compatibility with Iroha.
    */
    SetSettingValue = 'setSettingValue',
    /**
    * This command is not availalbe for use because it is related to smart contract.
    */
    CallEngine = 'callEngine'
}

/**
 * 
 * @export
 * @enum {string}
 */

export enum IrohaQuery {
    /**
    * To get the state of an account
    */
    GetAccount = 'getAccount',
    /**
    * To get details of the account.
    */
    GetAccountDetail = 'getAccountDetail',
    /**
    * To get information on the given asset (as for now - its precision).
    */
    GetAssetInfo = 'getAssetInfo',
    /**
    * To get the state of all assets in an account (a balance).
    */
    GetAccountAssets = 'getAccountAssets',
    /**
    * To retrieve information about transactions, based on their hashes.
    */
    GetTransactions = 'getTransactions',
    /**
    * To retrieve a list of pending (not fully signed) multisignature transactions or batches of transactions issued by account of query creator.
    */
    GetPendingTransactions = 'getPendingTransactions',
    /**
    * To retrieve a list of transactions per account.
    */
    GetAccountTransactions = 'getAccountTransactions',
    /**
    * To retrieve all transactions associated with given account and asset.
    */
    GetAccountAssetTransactions = 'getAccountAssetTransactions',
    /**
    * To get existing roles in the system.
    */
    GetRoles = 'getRoles',
    /**
    * To get signatories, which act as an identity of the account.
    */
    GetSignatories = 'getSignatories',
    /**
    * To get available permissions per role in the system.
    */
    GetRolePermissions = 'getRolePermissions',
    /**
    * To get a specific block, using its height as an identifier.
    */
    GetBlock = 'getBlock',
    /**
    * To retrieve a receipt of a CallEngine command. Allows to access the event log created during computations inside the EVM.
    */
    GetEngineReceipts = 'getEngineReceipts',
    /**
    * To get new blocks as soon as they are committed, a user can invoke FetchCommits RPC call to Iroha network.
    */
    FetchCommits = 'fetchCommits',
    /**
    * A query that returns a list of peers in Iroha network.
    */
    GetPeers = 'getPeers'
}

/**
 * 
 * @export
 * @interface KeyPair
 */
export interface KeyPair {
    /**
     * SHA-3 ed25519 public keys of length 64 are recommended.
     * @type {string}
     * @memberof KeyPair
     */
    publicKey: string;
    /**
     * SHA-3 ed25519 private keys of length 64 are recommended.
     * @type {string}
     * @memberof KeyPair
     */
    privateKey: string;
}
/**
 * 
 * @export
 * @interface RunTransactionRequestV1
 */
export interface RunTransactionRequestV1 {
    /**
     * 
     * @type {string}
     * @memberof RunTransactionRequestV1
     */
    commandName: string;
    /**
     * 
     * @type {IrohaBaseConfig}
     * @memberof RunTransactionRequestV1
     */
    baseConfig: IrohaBaseConfig;
    /**
     * The list of arguments to pass in to the transaction request.
     * @type {Array<any>}
     * @memberof RunTransactionRequestV1
     */
    params: Array<any>;
}
/**
 * 
 * @export
 * @interface RunTransactionResponse
 */
export interface RunTransactionResponse {
    /**
     * 
     * @type {any}
     * @memberof RunTransactionResponse
     */
    transactionReceipt: any | null;
}
/**
 * 
 * @export
 * @interface RunTransactionSignedRequestV1
 */
export interface RunTransactionSignedRequestV1 {
    /**
     * Signed transaction binary data received from generate-transaction endpoint.
     * @type {any}
     * @memberof RunTransactionSignedRequestV1
     */
    signedTransaction: any;
    /**
     * 
     * @type {IrohaBaseConfig}
     * @memberof RunTransactionSignedRequestV1
     */
    baseConfig?: IrohaBaseConfig;
}

/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Generate transaction that can be signed locally.
         * @param {GenerateTransactionRequestV1} [generateTransactionRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        generateTransactionV1: async (generateTransactionRequestV1?: GenerateTransactionRequestV1, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha/generate-transaction`;
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

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(generateTransactionRequestV1, localVarRequestOptions, configuration)

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
        getPrometheusMetricsV1: async (options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha/get-prometheus-exporter-metrics`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Executes a transaction on a Iroha ledger
         * @param {RunTransactionRequestV1 | RunTransactionSignedRequestV1} [runTransactionRequestV1RunTransactionSignedRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runTransactionV1: async (runTransactionRequestV1RunTransactionSignedRequestV1?: RunTransactionRequestV1 | RunTransactionSignedRequestV1, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha/run-transaction`;
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

            setSearchParams(localVarUrlObj, localVarQueryParameter, options.query);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(runTransactionRequestV1RunTransactionSignedRequestV1, localVarRequestOptions, configuration)

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
         * @summary Generate transaction that can be signed locally.
         * @param {GenerateTransactionRequestV1} [generateTransactionRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async generateTransactionV1(generateTransactionRequestV1?: GenerateTransactionRequestV1, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<any>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.generateTransactionV1(generateTransactionRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get the Prometheus Metrics
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getPrometheusMetricsV1(options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<string>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getPrometheusMetricsV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Executes a transaction on a Iroha ledger
         * @param {RunTransactionRequestV1 | RunTransactionSignedRequestV1} [runTransactionRequestV1RunTransactionSignedRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1?: RunTransactionRequestV1 | RunTransactionSignedRequestV1, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<RunTransactionResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1, options);
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
         * @summary Generate transaction that can be signed locally.
         * @param {GenerateTransactionRequestV1} [generateTransactionRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        generateTransactionV1(generateTransactionRequestV1?: GenerateTransactionRequestV1, options?: any): AxiosPromise<any> {
            return localVarFp.generateTransactionV1(generateTransactionRequestV1, options).then((request) => request(axios, basePath));
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
         * @summary Executes a transaction on a Iroha ledger
         * @param {RunTransactionRequestV1 | RunTransactionSignedRequestV1} [runTransactionRequestV1RunTransactionSignedRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1?: RunTransactionRequestV1 | RunTransactionSignedRequestV1, options?: any): AxiosPromise<RunTransactionResponse> {
            return localVarFp.runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1, options).then((request) => request(axios, basePath));
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
     * @summary Generate transaction that can be signed locally.
     * @param {GenerateTransactionRequestV1} [generateTransactionRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public generateTransactionV1(generateTransactionRequestV1?: GenerateTransactionRequestV1, options?: any) {
        return DefaultApiFp(this.configuration).generateTransactionV1(generateTransactionRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get the Prometheus Metrics
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getPrometheusMetricsV1(options?: any) {
        return DefaultApiFp(this.configuration).getPrometheusMetricsV1(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Executes a transaction on a Iroha ledger
     * @param {RunTransactionRequestV1 | RunTransactionSignedRequestV1} [runTransactionRequestV1RunTransactionSignedRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1?: RunTransactionRequestV1 | RunTransactionSignedRequestV1, options?: any) {
        return DefaultApiFp(this.configuration).runTransactionV1(runTransactionRequestV1RunTransactionSignedRequestV1, options).then((request) => request(this.axios, this.basePath));
    }
}


