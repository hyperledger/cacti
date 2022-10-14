/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cactus Plugin - Connector Iroha V2
 * Can perform basic tasks on a Iroha V2 ledger
 *
 * The version of the OpenAPI document: 1.0.0
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
 * Iroha V2 block response type.
 * @export
 * @enum {string}
 */

export enum BlockTypeV1 {
    /**
    * Default JSON-encoded string full block data.
    */
    Raw = 'raw',
    /**
    * Encoded format that must be decoded with Iroha SDK on client side before use
    */
    Binary = 'binary'
}

/**
 * Error response from the connector.
 * @export
 * @interface ErrorExceptionResponseV1
 */
export interface ErrorExceptionResponseV1 {
    /**
     * Short error description message.
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    message: string;
    /**
     * Detailed error information.
     * @type {string}
     * @memberof ErrorExceptionResponseV1
     */
    error: string;
}
/**
 * Request for generating transaction payload that can be signed on the client side.
 * @export
 * @interface GenerateTransactionRequestV1
 */
export interface GenerateTransactionRequestV1 {
    /**
     * 
     * @type {IrohaTransactionDefinitionV1}
     * @memberof GenerateTransactionRequestV1
     */
    transaction: IrohaTransactionDefinitionV1;
    /**
     * 
     * @type {Iroha2BaseConfig}
     * @memberof GenerateTransactionRequestV1
     */
    baseConfig?: Iroha2BaseConfig;
}
/**
 * Iroha V2 account ID.
 * @export
 * @interface Iroha2AccountId
 */
export interface Iroha2AccountId {
    /**
     * 
     * @type {string}
     * @memberof Iroha2AccountId
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof Iroha2AccountId
     */
    domainId: string;
}
/**
 * Iroha V2 connection configuration.
 * @export
 * @interface Iroha2BaseConfig
 */
export interface Iroha2BaseConfig {
    /**
     * 
     * @type {Iroha2BaseConfigTorii}
     * @memberof Iroha2BaseConfig
     */
    torii: Iroha2BaseConfigTorii;
    /**
     * 
     * @type {Iroha2AccountId}
     * @memberof Iroha2BaseConfig
     */
    accountId?: Iroha2AccountId;
    /**
     * 
     * @type {Iroha2KeyPair | KeychainReference}
     * @memberof Iroha2BaseConfig
     */
    signingCredential?: Iroha2KeyPair | KeychainReference;
}
/**
 * Iroha V2 peer connection information.
 * @export
 * @interface Iroha2BaseConfigTorii
 */
export interface Iroha2BaseConfigTorii {
    /**
     * 
     * @type {string}
     * @memberof Iroha2BaseConfigTorii
     */
    apiURL?: string;
    /**
     * 
     * @type {string}
     * @memberof Iroha2BaseConfigTorii
     */
    telemetryURL?: string;
}
/**
 * Private/Public key JSON containing payload and digest function.
 * @export
 * @interface Iroha2KeyJson
 */
export interface Iroha2KeyJson {
    /**
     * 
     * @type {string}
     * @memberof Iroha2KeyJson
     */
    digestFunction: string;
    /**
     * 
     * @type {string}
     * @memberof Iroha2KeyJson
     */
    payload: string;
}
/**
 * Pair of Iroha account private and public keys.
 * @export
 * @interface Iroha2KeyPair
 */
export interface Iroha2KeyPair {
    /**
     * 
     * @type {Iroha2KeyJson}
     * @memberof Iroha2KeyPair
     */
    privateKey: Iroha2KeyJson;
    /**
     * 
     * @type {string}
     * @memberof Iroha2KeyPair
     */
    publicKey: string;
}
/**
 * Command names that correspond to Iroha Special Instructions (https://hyperledger.github.io/iroha-2-docs/guide/advanced/isi.html)
 * @export
 * @enum {string}
 */

export enum IrohaInstruction {
    /**
    * Register new domain
    */
    RegisterDomain = 'registerDomain',
    /**
    * Register new asset definition
    */
    RegisterAssetDefinition = 'registerAssetDefinition',
    /**
    * Register new asset
    */
    RegisterAsset = 'registerAsset',
    /**
    * Mint asset value
    */
    MintAsset = 'mintAsset',
    /**
    * Burn asset value
    */
    BurnAsset = 'burnAsset',
    /**
    * Transfer asset between accounts
    */
    TransferAsset = 'transferAsset',
    /**
    * Register new account
    */
    RegisterAccount = 'registerAccount'
}

/**
 * Single Iroha V2 instruction to be executed request.
 * @export
 * @interface IrohaInstructionRequestV1
 */
export interface IrohaInstructionRequestV1 {
    /**
     * Iroha V2 instruction name.
     * @type {IrohaInstruction}
     * @memberof IrohaInstructionRequestV1
     */
    name: IrohaInstruction;
    /**
     * The list of arguments to pass with specified instruction.
     * @type {Array<any>}
     * @memberof IrohaInstructionRequestV1
     */
    params: Array<any>;
}
/**
 * Command names that correspond to Iroha queries (https://hyperledger.github.io/iroha-2-docs/guide/advanced/queries.html)
 * @export
 * @enum {string}
 */

export enum IrohaQuery {
    /**
    * Get list of all registered domains
    */
    FindAllDomains = 'findAllDomains',
    /**
    * Get domain with specified ID
    */
    FindDomainById = 'findDomainById',
    /**
    * Get asset definition with specified ID
    */
    FindAssetDefinitionById = 'findAssetDefinitionById',
    /**
    * Get list of all registered asset definition
    */
    FindAllAssetsDefinitions = 'findAllAssetsDefinitions',
    /**
    * Get asset with specified ID
    */
    FindAssetById = 'findAssetById',
    /**
    * Get list of all registered assets
    */
    FindAllAssets = 'findAllAssets',
    /**
    * Get list of all ledger peers
    */
    FindAllPeers = 'findAllPeers',
    /**
    * Get list of all ledger blocks
    */
    FindAllBlocks = 'findAllBlocks',
    /**
    * Get account with specified ID
    */
    FindAccountById = 'findAccountById',
    /**
    * Get list of all registered accounts
    */
    FindAllAccounts = 'findAllAccounts',
    /**
    * Get list of all transactions
    */
    FindAllTransactions = 'findAllTransactions',
    /**
    * Get transaction with specified hash
    */
    FindTransactionByHash = 'findTransactionByHash'
}

/**
 * Iroha V2 transaction definition
 * @export
 * @interface IrohaTransactionDefinitionV1
 */
export interface IrohaTransactionDefinitionV1 {
    /**
     * 
     * @type {IrohaInstructionRequestV1 | Array<IrohaInstructionRequestV1>}
     * @memberof IrohaTransactionDefinitionV1
     */
    instruction: IrohaInstructionRequestV1 | Array<IrohaInstructionRequestV1>;
    /**
     * 
     * @type {IrohaTransactionParametersV1}
     * @memberof IrohaTransactionDefinitionV1
     */
    params?: IrohaTransactionParametersV1;
}
/**
 * Iroha V2 transaction payload parameters
 * @export
 * @interface IrohaTransactionParametersV1
 */
export interface IrohaTransactionParametersV1 {
    [key: string]: object | any;

    /**
     * BigInt time to live.
     * @type {string}
     * @memberof IrohaTransactionParametersV1
     */
    ttl?: string;
    /**
     * BigInt creation time
     * @type {string}
     * @memberof IrohaTransactionParametersV1
     */
    creationTime?: string;
    /**
     * Transaction nonce
     * @type {number}
     * @memberof IrohaTransactionParametersV1
     */
    nonce?: number;
}
/**
 * Reference to entry stored in Cactus keychain plugin.
 * @export
 * @interface KeychainReference
 */
export interface KeychainReference {
    /**
     * Keychain plugin ID.
     * @type {string}
     * @memberof KeychainReference
     */
    keychainId: string;
    /**
     * Key reference name.
     * @type {string}
     * @memberof KeychainReference
     */
    keychainRef: string;
}
/**
 * Request to query endpoint.
 * @export
 * @interface QueryRequestV1
 */
export interface QueryRequestV1 {
    /**
     * Name of the query to be executed.
     * @type {IrohaQuery}
     * @memberof QueryRequestV1
     */
    queryName: IrohaQuery;
    /**
     * 
     * @type {Iroha2BaseConfig}
     * @memberof QueryRequestV1
     */
    baseConfig?: Iroha2BaseConfig;
    /**
     * The list of arguments to pass with the query.
     * @type {Array<any>}
     * @memberof QueryRequestV1
     */
    params?: Array<any>;
}
/**
 * Response with query results.
 * @export
 * @interface QueryResponseV1
 */
export interface QueryResponseV1 {
    /**
     * Query response data that varies between different queries.
     * @type {any}
     * @memberof QueryResponseV1
     */
    response: any;
}
/**
 * Request to transact endpoint.
 * @export
 * @interface TransactRequestV1
 */
export interface TransactRequestV1 {
    /**
     * Signed transaction binary data received from generate-transaction endpoint.
     * @type {any}
     * @memberof TransactRequestV1
     */
    signedTransaction?: any;
    /**
     * 
     * @type {IrohaTransactionDefinitionV1}
     * @memberof TransactRequestV1
     */
    transaction?: IrohaTransactionDefinitionV1;
    /**
     * Wait unitl transaction is sent and return the final status (committed / rejected)
     * @type {boolean}
     * @memberof TransactRequestV1
     */
    waitForCommit?: boolean;
    /**
     * 
     * @type {Iroha2BaseConfig}
     * @memberof TransactRequestV1
     */
    baseConfig?: Iroha2BaseConfig;
}
/**
 * Response from transaction endpoint with operation status.
 * @export
 * @interface TransactResponseV1
 */
export interface TransactResponseV1 {
    /**
     * Hexadecimal hash of the transaction sent to the ledger.
     * @type {string}
     * @memberof TransactResponseV1
     */
    hash: string;
    /**
     * 
     * @type {TransactionStatus}
     * @memberof TransactResponseV1
     */
    status: TransactionStatus;
    /**
     * When waitForCommit was suplied and the transaction was rejected, contains the reason of the rejection.
     * @type {string}
     * @memberof TransactResponseV1
     */
    rejectReason?: string;
}
/**
 * Status of Iroha V2 transaction.
 * @export
 * @enum {string}
 */

export enum TransactionStatus {
    /**
    * Transaction was submitted to the ledger - use other tools to check if it was accepted and committed.
    */
    Submitted = 'submitted',
    /**
    * Transaction was committed to the ledger.
    */
    Committed = 'committed',
    /**
    * Transaction was rejected.
    */
    Rejected = 'rejected'
}

/**
 * Binary encoded response of block data.
 * @export
 * @interface WatchBlocksBinaryResponseV1
 */
export interface WatchBlocksBinaryResponseV1 {
    /**
     * 
     * @type {any}
     * @memberof WatchBlocksBinaryResponseV1
     */
    binaryBlock: any;
}
/**
 * Options passed when subscribing to block monitoring.
 * @export
 * @interface WatchBlocksOptionsV1
 */
export interface WatchBlocksOptionsV1 {
    /**
     * 
     * @type {BlockTypeV1}
     * @memberof WatchBlocksOptionsV1
     */
    type?: BlockTypeV1;
    /**
     * Number of block to start monitoring from.
     * @type {string}
     * @memberof WatchBlocksOptionsV1
     */
    startBlock?: string;
    /**
     * 
     * @type {Iroha2BaseConfig}
     * @memberof WatchBlocksOptionsV1
     */
    baseConfig?: Iroha2BaseConfig;
}
/**
 * Default JSON-encoded string full block data.
 * @export
 * @interface WatchBlocksRawResponseV1
 */
export interface WatchBlocksRawResponseV1 {
    /**
     * 
     * @type {string}
     * @memberof WatchBlocksRawResponseV1
     */
    blockData: string;
}
/**
 * @type WatchBlocksResponseV1
 * @export
 */
export type WatchBlocksResponseV1 = ErrorExceptionResponseV1 | WatchBlocksBinaryResponseV1 | WatchBlocksRawResponseV1;

/**
 * Websocket requests for monitoring new blocks.
 * @export
 * @enum {string}
 */

export enum WatchBlocksV1 {
    Subscribe = 'org.hyperledger.cactus.api.async.hliroha2.WatchBlocksV1.Subscribe',
    Next = 'org.hyperledger.cactus.api.async.hliroha2.WatchBlocksV1.Next',
    Unsubscribe = 'org.hyperledger.cactus.api.async.hliroha2.WatchBlocksV1.Unsubscribe',
    Error = 'org.hyperledger.cactus.api.async.hliroha2.WatchBlocksV1.Error',
    Complete = 'org.hyperledger.cactus.api.async.hliroha2.WatchBlocksV1.Complete'
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
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha2/generate-transaction`;
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
         * @summary Executes a query on a Iroha V2 ledger and returns it\'s results.
         * @param {QueryRequestV1} [queryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        queryV1: async (queryRequestV1?: QueryRequestV1, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha2/query`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(queryRequestV1, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Executes a transaction on a Iroha V2 ledger (by sending some instructions)
         * @param {TransactRequestV1} [transactRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        transactV1: async (transactRequestV1?: TransactRequestV1, options: any = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-iroha2/transact`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(transactRequestV1, localVarRequestOptions, configuration)

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
         * @summary Executes a query on a Iroha V2 ledger and returns it\'s results.
         * @param {QueryRequestV1} [queryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async queryV1(queryRequestV1?: QueryRequestV1, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<QueryResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.queryV1(queryRequestV1, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Executes a transaction on a Iroha V2 ledger (by sending some instructions)
         * @param {TransactRequestV1} [transactRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async transactV1(transactRequestV1?: TransactRequestV1, options?: any): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<TransactResponseV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.transactV1(transactRequestV1, options);
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
         * @summary Executes a query on a Iroha V2 ledger and returns it\'s results.
         * @param {QueryRequestV1} [queryRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        queryV1(queryRequestV1?: QueryRequestV1, options?: any): AxiosPromise<QueryResponseV1> {
            return localVarFp.queryV1(queryRequestV1, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Executes a transaction on a Iroha V2 ledger (by sending some instructions)
         * @param {TransactRequestV1} [transactRequestV1] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        transactV1(transactRequestV1?: TransactRequestV1, options?: any): AxiosPromise<TransactResponseV1> {
            return localVarFp.transactV1(transactRequestV1, options).then((request) => request(axios, basePath));
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
     * @summary Executes a query on a Iroha V2 ledger and returns it\'s results.
     * @param {QueryRequestV1} [queryRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public queryV1(queryRequestV1?: QueryRequestV1, options?: any) {
        return DefaultApiFp(this.configuration).queryV1(queryRequestV1, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Executes a transaction on a Iroha V2 ledger (by sending some instructions)
     * @param {TransactRequestV1} [transactRequestV1] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public transactV1(transactRequestV1?: TransactRequestV1, options?: any) {
        return DefaultApiFp(this.configuration).transactV1(transactRequestV1, options).then((request) => request(this.axios, this.basePath));
    }
}


