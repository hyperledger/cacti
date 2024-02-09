/* tslint:disable */
/* eslint-disable */
/**
 * Hyperledger Cacti Plugin - Connector Aries
 * Can communicate with other Aries agents and Cacti Aries connectors
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
 * Request for AcceptInvitation endpoint.
 * @export
 * @interface AcceptInvitationV1Request
 */
export interface AcceptInvitationV1Request {
    /**
     * Aries label of an agent to be used to connect using URL
     * @type {string}
     * @memberof AcceptInvitationV1Request
     */
    'agentName': string;
    /**
     * Invitation URL generated by another aries agent.
     * @type {string}
     * @memberof AcceptInvitationV1Request
     */
    'invitationUrl': string;
}
/**
 * Response for AcceptInvitation endpoint.
 * @export
 * @interface AcceptInvitationV1Response
 */
export interface AcceptInvitationV1Response {
    /**
     * ID that can be used to track status of the connection
     * @type {string}
     * @memberof AcceptInvitationV1Response
     */
    'outOfBandId': string;
}
/**
 * Aries agent connection information.
 * @export
 * @interface AgentConnectionRecordV1
 */
export interface AgentConnectionRecordV1 {
    [key: string]: any;

    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'state': string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'role': string;
    /**
     * 
     * @type {boolean}
     * @memberof AgentConnectionRecordV1
     */
    'isReady': boolean;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'did'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'theirDid'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'theirLabel'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'alias'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'threadId'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'errorMessage'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'outOfBandId'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionRecordV1
     */
    'invitationDid'?: string;
}
/**
 * Fields that can be used to filter agent connection list.
 * @export
 * @interface AgentConnectionsFilterV1
 */
export interface AgentConnectionsFilterV1 {
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'did'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'invitationDid'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'outOfBandId'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'role'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'state'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'theirDid'?: string;
    /**
     * 
     * @type {string}
     * @memberof AgentConnectionsFilterV1
     */
    'threadId'?: string;
}
/**
 * Aries agent configuration to be setup and used by the connector.
 * @export
 * @interface AriesAgentConfigV1
 */
export interface AriesAgentConfigV1 {
    /**
     * Aries agent label that will also be used as wallet id.
     * @type {string}
     * @memberof AriesAgentConfigV1
     */
    'name': string;
    /**
     * Wallet private key - do not share with anyone.
     * @type {string}
     * @memberof AriesAgentConfigV1
     */
    'walletKey': string;
    /**
     * Path to wallet sqlite database to use. If not provided, the connector default path and agent name will be used.
     * @type {string}
     * @memberof AriesAgentConfigV1
     */
    'walletPath'?: string;
    /**
     * 
     * @type {Array<AriesIndyVdrPoolConfigV1>}
     * @memberof AriesAgentConfigV1
     */
    'indyNetworks': Array<AriesIndyVdrPoolConfigV1>;
    /**
     * Inbound endpoint URL for this agent. Must be unique for this connector. Must contain port.
     * @type {string}
     * @memberof AriesAgentConfigV1
     */
    'inboundUrl'?: string;
    /**
     * Flag to accept new connection by default
     * @type {boolean}
     * @memberof AriesAgentConfigV1
     */
    'autoAcceptConnections'?: boolean;
    /**
     * 
     * @type {CactiAcceptPolicyV1}
     * @memberof AriesAgentConfigV1
     */
    'autoAcceptCredentials'?: CactiAcceptPolicyV1;
    /**
     * 
     * @type {CactiAcceptPolicyV1}
     * @memberof AriesAgentConfigV1
     */
    'autoAcceptProofs'?: CactiAcceptPolicyV1;
}


/**
 * Summary of an Aries Agent configured in the connector.
 * @export
 * @interface AriesAgentSummaryV1
 */
export interface AriesAgentSummaryV1 {
    /**
     * Aries label of an agent
     * @type {string}
     * @memberof AriesAgentSummaryV1
     */
    'name': string;
    /**
     * True when Aries agent has been initialized properly.
     * @type {boolean}
     * @memberof AriesAgentSummaryV1
     */
    'isAgentInitialized': boolean;
    /**
     * True when this agents wallet has been initialized properly.
     * @type {boolean}
     * @memberof AriesAgentSummaryV1
     */
    'isWalletInitialized': boolean;
    /**
     * True when this agents wallet has been provisioned properly.
     * @type {boolean}
     * @memberof AriesAgentSummaryV1
     */
    'isWalletProvisioned': boolean;
    /**
     * 
     * @type {AriesAgentSummaryV1WalletConfig}
     * @memberof AriesAgentSummaryV1
     */
    'walletConfig': AriesAgentSummaryV1WalletConfig;
    /**
     * Aries agent endpoints configured
     * @type {Array<string>}
     * @memberof AriesAgentSummaryV1
     */
    'endpoints': Array<string>;
}
/**
 * 
 * @export
 * @interface AriesAgentSummaryV1WalletConfig
 */
export interface AriesAgentSummaryV1WalletConfig {
    /**
     * Wallet entry ID
     * @type {string}
     * @memberof AriesAgentSummaryV1WalletConfig
     */
    'id': string;
    /**
     * Wallet storage type
     * @type {string}
     * @memberof AriesAgentSummaryV1WalletConfig
     */
    'type': string;
}
/**
 * Indy VDR network configuration
 * @export
 * @interface AriesIndyVdrPoolConfigV1
 */
export interface AriesIndyVdrPoolConfigV1 {
    /**
     * Indy genesis transactions.
     * @type {string}
     * @memberof AriesIndyVdrPoolConfigV1
     */
    'genesisTransactions': string;
    /**
     * Flag to specify whether this is production or development ledger.
     * @type {boolean}
     * @memberof AriesIndyVdrPoolConfigV1
     */
    'isProduction': boolean;
    /**
     * Indy namespace
     * @type {string}
     * @memberof AriesIndyVdrPoolConfigV1
     */
    'indyNamespace': string;
    /**
     * Connect to the ledger on startup flag
     * @type {boolean}
     * @memberof AriesIndyVdrPoolConfigV1
     */
    'connectOnStartup'?: boolean;
}
/**
 * Proof exchange record from Aries framework (simplified)
 * @export
 * @interface AriesProofExchangeRecordV1
 */
export interface AriesProofExchangeRecordV1 {
    [key: string]: any;

    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'id': string;
    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'connectionId'?: string;
    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'threadId': string;
    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'state': string;
    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'protocolVersion': string;
    /**
     * 
     * @type {boolean}
     * @memberof AriesProofExchangeRecordV1
     */
    'isVerified'?: boolean;
    /**
     * 
     * @type {string}
     * @memberof AriesProofExchangeRecordV1
     */
    'errorMessage'?: string;
}
/**
 * Credential / Proof requests acceptance policies for Aries agent
 * @export
 * @enum {string}
 */

export const CactiAcceptPolicyV1 = {
    Always: 'always',
    ContentApproved: 'contentApproved',
    Never: 'never'
} as const;

export type CactiAcceptPolicyV1 = typeof CactiAcceptPolicyV1[keyof typeof CactiAcceptPolicyV1];


/**
 * Credential attribute checks to be performed by a proof request.
 * @export
 * @interface CactiProofRequestAttributeV1
 */
export interface CactiProofRequestAttributeV1 {
    /**
     * Attribute name.
     * @type {string}
     * @memberof CactiProofRequestAttributeV1
     */
    'name': string;
    /**
     * Check if attribute has specified value
     * @type {any}
     * @memberof CactiProofRequestAttributeV1
     */
    'isValueEqual'?: any;
    /**
     * Check if credentialDefinitionId has specified value
     * @type {any}
     * @memberof CactiProofRequestAttributeV1
     */
    'isCredentialDefinitionIdEqual'?: any;
}
/**
 * Request for CreateNewConnectionInvitation endpoint.
 * @export
 * @interface CreateNewConnectionInvitationV1Request
 */
export interface CreateNewConnectionInvitationV1Request {
    /**
     * Aries label of an agent to use to generate an invitation
     * @type {string}
     * @memberof CreateNewConnectionInvitationV1Request
     */
    'agentName': string;
    /**
     * Invitation URL domain to use. If not specified, then connector default domain will be used
     * @type {string}
     * @memberof CreateNewConnectionInvitationV1Request
     */
    'invitationDomain'?: string;
}
/**
 * Response for CreateNewConnectionInvitation endpoint.
 * @export
 * @interface CreateNewConnectionInvitationV1Response
 */
export interface CreateNewConnectionInvitationV1Response {
    /**
     * Invitation URL that can be used by another aries agent to connect to us.
     * @type {string}
     * @memberof CreateNewConnectionInvitationV1Response
     */
    'invitationUrl': string;
    /**
     * ID that can be used to track status of the connection
     * @type {string}
     * @memberof CreateNewConnectionInvitationV1Response
     */
    'outOfBandId': string;
}
/**
 * Error response from the connector.
 * @export
 * @interface ErrorExceptionV1Response
 */
export interface ErrorExceptionV1Response {
    /**
     * Short error description message.
     * @type {string}
     * @memberof ErrorExceptionV1Response
     */
    'message': string;
    /**
     * Detailed error information.
     * @type {string}
     * @memberof ErrorExceptionV1Response
     */
    'error': string;
}
/**
 * Request for GetConnections endpoint.
 * @export
 * @interface GetConnectionsV1Request
 */
export interface GetConnectionsV1Request {
    /**
     * 
     * @type {string}
     * @memberof GetConnectionsV1Request
     */
    'agentName': string;
    /**
     * 
     * @type {AgentConnectionsFilterV1}
     * @memberof GetConnectionsV1Request
     */
    'filter'?: AgentConnectionsFilterV1;
}
/**
 * Request for RequestProof endpoint.
 * @export
 * @interface RequestProofV1Request
 */
export interface RequestProofV1Request {
    /**
     * Aries label of an agent to be used to connect using URL
     * @type {string}
     * @memberof RequestProofV1Request
     */
    'agentName': string;
    /**
     * Peer connection ID from which we want to request a proof.
     * @type {string}
     * @memberof RequestProofV1Request
     */
    'connectionId': string;
    /**
     * 
     * @type {Array<CactiProofRequestAttributeV1>}
     * @memberof RequestProofV1Request
     */
    'proofAttributes': Array<CactiProofRequestAttributeV1>;
}
/**
 * Options passed when monitoring connection change events.
 * @export
 * @interface WatchConnectionStateOptionsV1
 */
export interface WatchConnectionStateOptionsV1 {
    /**
     * Aries agent label that will also be used as wallet id.
     * @type {string}
     * @memberof WatchConnectionStateOptionsV1
     */
    'agentName': string;
}
/**
 * Values pushed on each connection state change.
 * @export
 * @interface WatchConnectionStateProgressV1
 */
export interface WatchConnectionStateProgressV1 {
    /**
     * 
     * @type {AgentConnectionRecordV1}
     * @memberof WatchConnectionStateProgressV1
     */
    'connectionRecord': AgentConnectionRecordV1;
    /**
     * 
     * @type {string}
     * @memberof WatchConnectionStateProgressV1
     */
    'previousState': string | null;
}
/**
 * Websocket requests for monitoring connection change events.
 * @export
 * @enum {string}
 */

export const WatchConnectionStateV1 = {
    Subscribe: 'org.hyperledger.cactus.api.async.hlaries.WatchConnectionStateV1.Subscribe',
    Next: 'org.hyperledger.cactus.api.async.hlaries.WatchConnectionStateV1.Next',
    Unsubscribe: 'org.hyperledger.cactus.api.async.hlaries.WatchConnectionStateV1.Unsubscribe',
    Error: 'org.hyperledger.cactus.api.async.hlaries.WatchConnectionStateV1.Error',
    Complete: 'org.hyperledger.cactus.api.async.hlaries.WatchConnectionStateV1.Complete'
} as const;

export type WatchConnectionStateV1 = typeof WatchConnectionStateV1[keyof typeof WatchConnectionStateV1];


/**
 * Options passed when monitoring proof state change events.
 * @export
 * @interface WatchProofStateOptionsV1
 */
export interface WatchProofStateOptionsV1 {
    /**
     * Aries agent label that will also be used as wallet id.
     * @type {string}
     * @memberof WatchProofStateOptionsV1
     */
    'agentName': string;
}
/**
 * Values pushed on each proof state change.
 * @export
 * @interface WatchProofStateProgressV1
 */
export interface WatchProofStateProgressV1 {
    /**
     * 
     * @type {AriesProofExchangeRecordV1}
     * @memberof WatchProofStateProgressV1
     */
    'proofRecord': AriesProofExchangeRecordV1;
    /**
     * 
     * @type {string}
     * @memberof WatchProofStateProgressV1
     */
    'previousState': string | null;
}
/**
 * Websocket requests for monitoring proof state change events.
 * @export
 * @enum {string}
 */

export const WatchProofStateV1 = {
    Subscribe: 'org.hyperledger.cactus.api.async.hlaries.WatchProofStateV1.Subscribe',
    Next: 'org.hyperledger.cactus.api.async.hlaries.WatchProofStateV1.Next',
    Unsubscribe: 'org.hyperledger.cactus.api.async.hlaries.WatchProofStateV1.Unsubscribe',
    Error: 'org.hyperledger.cactus.api.async.hlaries.WatchProofStateV1.Error',
    Complete: 'org.hyperledger.cactus.api.async.hlaries.WatchProofStateV1.Complete'
} as const;

export type WatchProofStateV1 = typeof WatchProofStateV1[keyof typeof WatchProofStateV1];



/**
 * DefaultApi - axios parameter creator
 * @export
 */
export const DefaultApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Connect to another agent using it\'s invitation URL
         * @param {AcceptInvitationV1Request} [acceptInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        acceptInvitationV1: async (acceptInvitationV1Request?: AcceptInvitationV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-aries/accept-invitation`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(acceptInvitationV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create new aries agent invitation that other agents can use to connect.
         * @param {CreateNewConnectionInvitationV1Request} [createNewConnectionInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createNewConnectionInvitationV1: async (createNewConnectionInvitationV1Request?: CreateNewConnectionInvitationV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-aries/create-new-connection-invitation`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(createNewConnectionInvitationV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get all Aries agents configured in this connector plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAgentsV1: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-aries/get-agents`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
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
         * @summary Get all connections of given aries agent.
         * @param {GetConnectionsV1Request} [getConnectionsV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getConnectionsV1: async (getConnectionsV1Request?: GetConnectionsV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-aries/get-connections`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(getConnectionsV1Request, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Request proof matching provided requriements from connected peer agent.
         * @param {RequestProofV1Request} [requestProofV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestProofV1: async (requestProofV1Request?: RequestProofV1Request, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-aries/request-proof`;
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
            localVarRequestOptions.data = serializeDataIfNeeded(requestProofV1Request, localVarRequestOptions, configuration)

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
         * @summary Connect to another agent using it\'s invitation URL
         * @param {AcceptInvitationV1Request} [acceptInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async acceptInvitationV1(acceptInvitationV1Request?: AcceptInvitationV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<AcceptInvitationV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.acceptInvitationV1(acceptInvitationV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Create new aries agent invitation that other agents can use to connect.
         * @param {CreateNewConnectionInvitationV1Request} [createNewConnectionInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createNewConnectionInvitationV1(createNewConnectionInvitationV1Request?: CreateNewConnectionInvitationV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<CreateNewConnectionInvitationV1Response>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.createNewConnectionInvitationV1(createNewConnectionInvitationV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get all Aries agents configured in this connector plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getAgentsV1(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<AriesAgentSummaryV1>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getAgentsV1(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get all connections of given aries agent.
         * @param {GetConnectionsV1Request} [getConnectionsV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getConnectionsV1(getConnectionsV1Request?: GetConnectionsV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<AgentConnectionRecordV1>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getConnectionsV1(getConnectionsV1Request, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Request proof matching provided requriements from connected peer agent.
         * @param {RequestProofV1Request} [requestProofV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async requestProofV1(requestProofV1Request?: RequestProofV1Request, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<AriesProofExchangeRecordV1>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.requestProofV1(requestProofV1Request, options);
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
         * @summary Connect to another agent using it\'s invitation URL
         * @param {AcceptInvitationV1Request} [acceptInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        acceptInvitationV1(acceptInvitationV1Request?: AcceptInvitationV1Request, options?: any): AxiosPromise<AcceptInvitationV1Response> {
            return localVarFp.acceptInvitationV1(acceptInvitationV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create new aries agent invitation that other agents can use to connect.
         * @param {CreateNewConnectionInvitationV1Request} [createNewConnectionInvitationV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createNewConnectionInvitationV1(createNewConnectionInvitationV1Request?: CreateNewConnectionInvitationV1Request, options?: any): AxiosPromise<CreateNewConnectionInvitationV1Response> {
            return localVarFp.createNewConnectionInvitationV1(createNewConnectionInvitationV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get all Aries agents configured in this connector plugin.
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getAgentsV1(options?: any): AxiosPromise<Array<AriesAgentSummaryV1>> {
            return localVarFp.getAgentsV1(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get all connections of given aries agent.
         * @param {GetConnectionsV1Request} [getConnectionsV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getConnectionsV1(getConnectionsV1Request?: GetConnectionsV1Request, options?: any): AxiosPromise<Array<AgentConnectionRecordV1>> {
            return localVarFp.getConnectionsV1(getConnectionsV1Request, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Request proof matching provided requriements from connected peer agent.
         * @param {RequestProofV1Request} [requestProofV1Request] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        requestProofV1(requestProofV1Request?: RequestProofV1Request, options?: any): AxiosPromise<AriesProofExchangeRecordV1> {
            return localVarFp.requestProofV1(requestProofV1Request, options).then((request) => request(axios, basePath));
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
     * @summary Connect to another agent using it\'s invitation URL
     * @param {AcceptInvitationV1Request} [acceptInvitationV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public acceptInvitationV1(acceptInvitationV1Request?: AcceptInvitationV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).acceptInvitationV1(acceptInvitationV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create new aries agent invitation that other agents can use to connect.
     * @param {CreateNewConnectionInvitationV1Request} [createNewConnectionInvitationV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public createNewConnectionInvitationV1(createNewConnectionInvitationV1Request?: CreateNewConnectionInvitationV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).createNewConnectionInvitationV1(createNewConnectionInvitationV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get all Aries agents configured in this connector plugin.
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getAgentsV1(options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getAgentsV1(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get all connections of given aries agent.
     * @param {GetConnectionsV1Request} [getConnectionsV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public getConnectionsV1(getConnectionsV1Request?: GetConnectionsV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).getConnectionsV1(getConnectionsV1Request, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Request proof matching provided requriements from connected peer agent.
     * @param {RequestProofV1Request} [requestProofV1Request] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DefaultApi
     */
    public requestProofV1(requestProofV1Request?: RequestProofV1Request, options?: AxiosRequestConfig) {
        return DefaultApiFp(this.configuration).requestProofV1(requestProofV1Request, options).then((request) => request(this.axios, this.basePath));
    }
}


