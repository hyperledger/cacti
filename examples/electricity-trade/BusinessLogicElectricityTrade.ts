/*
 * Copyright 2020 Hyperledger Cactus Contributors
 * SPDX-License-Identifier: Apache-2.0
 *
 * BusinessLogicElectricityTrade.ts
 */

import { Request } from 'express';
import { RequestInfo } from './RequestInfo';
import { TradeInfo } from '../../packages/routing-interface/TradeInfo';
import { transactionManagement } from '../../packages/routing-interface/routes/index';
import { BusinessLogicBase } from '../../packages/business-logic-plugin/BusinessLogicBase';
import { makeRawTransaction } from './TransactionEthereum'
import { LedgerEvent } from '../../packages/ledger-plugin/LedgerPlugin';
import { json2str } from '../../packages/ledger-plugin/DriverCommon'

const fs = require('fs');
const path = require('path');
const config: any = JSON.parse(fs.readFileSync(path.resolve(__dirname, "./config/default.json"), 'utf8'));
import { getLogger } from "log4js";
const moduleName = 'BusinessLogicElectricityTrade';
const logger = getLogger(`${moduleName}`);
logger.level = config.logLevel;

export class BusinessLogicElectricityTrade extends BusinessLogicBase {
    businessLogicID: string;

    constructor(businessLogicID: string) {
        super();
        this.businessLogicID = businessLogicID;
    }

    startTransaction(req: Request, businessLogicID: string, tradeID: string) {

        logger.debug("called startTransaction()");

        // set RequestInfo
        const requestInfo: RequestInfo = new RequestInfo();
        requestInfo.businessLogicID = businessLogicID;

        // set TradeID
        requestInfo.setTradeID(tradeID);

        // Create trade information
        const tradeInfo: TradeInfo = new TradeInfo(requestInfo.businessLogicID, requestInfo.tradeID);

        this.startMonitor(tradeInfo);
    }

    startMonitor(tradeInfo: TradeInfo) {
        // Get Verifier Instance
        logger.debug(`##startMonitor(): businessLogicID: ${tradeInfo.businessLogicID}`);
        const useValidator = JSON.parse(transactionManagement.getValidatorToUse(tradeInfo.businessLogicID));
        logger.debug(`filterKey: ${config.electricityTradeInfo.sawtooth.filterKey}`);
        const options = {
            "filterKey": config.electricityTradeInfo.sawtooth.filterKey
        }
        const verifierSawtooth = transactionManagement.getVerifier(useValidator['validatorID'][0], options);
        logger.debug("getVerifierSawtooth");

    }

    remittanceTransaction(transactionSubset: object) {

        logger.debug(`called remittanceTransaction(), accountInfo = ${json2str(transactionSubset)}`);

        const accountInfo = this.getAccountInfo(transactionSubset);
        if (Object.keys(accountInfo).length === 0) {
            logger.debug(`remittanceTransaction(): skip (No target data)`);
            return;
        }

        const txParam: { fromAddress: string, fromAddressPkey: string, toAddress: string, amount: number, gas: number } = {
            fromAddress: accountInfo['fromAddress'],
            fromAddressPkey: accountInfo['fromAddressPkey'],
            toAddress: accountInfo['toAddress'],
            amount: Number(transactionSubset['Value']),
            gas: config.electricityTradeInfo.ethereum.gas
        };
        logger.debug(`####txParam = ${json2str(txParam)}`);

        // Get Verifier Instance
        logger.debug(`##remittanceTransaction(): businessLogicID: ${this.businessLogicID}`);
        const useValidator = JSON.parse(transactionManagement.getValidatorToUse(this.businessLogicID));
        const verifierEthereum = transactionManagement.getVerifier(useValidator['validatorID'][1]);
        logger.debug("getVerifierEthereum");

        // Generate parameters for// sendRawTransaction
        logger.debug(`####exec makeRawTransaction!!`);
        makeRawTransaction(txParam)
            .then(result => {
                logger.info('remittanceTransaction txId : ' + result.txId);

                // Set Parameter
                logger.debug('remittanceTransaction data : ' + json2str(result.data));
                const contract = {}; // NOTE: Since contract does not need to be specified, specify an empty object.
                const method = {type: "web3Eth", command: "sendRawTransaction"};
                const args = {"args": [result.data["serializedTx"]]};

                // Run Verifier (Ethereum)
                verifierEthereum.requestLedgerOperationNeo(contract, method, args);
            })
            .catch(err => {
                logger.error(err);
            });
    }

    onEvent(ledgerEvent: LedgerEvent, targetIndex: number): void {
        logger.debug(`##in BLP:onEvent()`);
        logger.debug(`##onEvent(): ${json2str(ledgerEvent['data']['blockData'][targetIndex])}`);

        switch (ledgerEvent.verifierId) {
              case config.electricityTradeInfo.sawtooth.validatorID:
                this.onEventSawtooth(ledgerEvent.data, targetIndex);
                break;
              case config.electricityTradeInfo.ethereum.validatorID:
                this.onEventEthereum(ledgerEvent.data, targetIndex);
                break;
            default:
                logger.error(`##onEvent(), invalid verifierId: ${ledgerEvent.verifierId}`);
                return;
        }
    }

    onEventSawtooth(event: object, targetIndex: number): void {
        logger.debug(`##in onEventSawtooth()`);
        const tx = this.getTransactionFromSawtoothEvent(event, targetIndex);
        if (tx == null) {
            logger.error(`##onEventSawtooth(): invalid event: ${json2str(event)}`);
            return;
        }

        try {
            const txId = tx['header_signature'];
            logger.debug(`##txId = ${txId}`);

            if (tx['payload_decoded'][0].Verb !== "set") {
                const transactionSubset = {
                    "Name": tx['payload_decoded'][0].Name,
                    "Value": tx['payload_decoded'][0].Value,
                    "Verb": tx['payload_decoded'][0].Verb
                };
                this.remittanceTransaction(transactionSubset);
            }
        }
        catch (err) {
            logger.error(`##onEventSawtooth(): err: ${err}, event: ${json2str(event)}`);
        }
    }

    getTransactionFromSawtoothEvent(event: object, targetIndex: number): object | null {
        try {
            const retTransaction = event['blockData'][targetIndex];
            logger.debug(`##getTransactionFromSawtoothEvent(), retTransaction: ${retTransaction}`);
            return retTransaction;
        }
        catch (err) {
            logger.error(`##getTransactionFromSawtoothEvent(): invalid even, err:${err}, event:${event}`);
        }
    }

    onEventEthereum(event: object, targetIndex: number): void {
        logger.debug(`##in onEventEthereum()`);
        const tx = this.getTransactionFromEthereumEvent(event, targetIndex);
        if (tx == null) {
            logger.error(`##onEventEthereum(): invalid event: ${json2str(event)}`);
            return;
        }

        try {
            const txId = tx['hash'];
            const status = event['status']
            logger.debug(`##txId = ${txId}`);
            logger.debug(`##status =${status}`);

            if (status !== 200) {
                logger.error(`##onEventEthereum(): error event, status: ${status}, txId: ${txId}`);
                return;
            }
        }
        catch (err) {
            logger.error(`##onEventEthereum(): err: ${err}, event: ${json2str(event)}`);
        }
    }

    getTransactionFromEthereumEvent(event: object, targetIndex: number): object | null {
        try {
            const retTransaction = event['blockData']['transactions'][targetIndex];
            logger.debug(`##getTransactionFromEthereumEvent(), retTransaction: ${retTransaction}`);
            return retTransaction;
        }
        catch (err) {
            logger.error(`##getTransactionFromEthereumEvent(): invalid even, err:${err}, event:${event}`);
        }
    }

    getOperationStatus(tradeID: string): object {
        logger.debug(`##in getOperationStatus()`);
        return {};
    }

    getTxIDFromEvent(ledgerEvent: LedgerEvent, targetIndex: number): string | null {
        logger.debug(`##in getTxIDFromEvent`);
//        logger.debug(`##event: ${json2str(ledgerEvent)}`);

        switch (ledgerEvent.verifierId) {
            case config.electricityTradeInfo.sawtooth.validatorID:
                return this.getTxIDFromEventSawtooth(ledgerEvent.data, targetIndex);
            case config.electricityTradeInfo.ethereum.validatorID:
                return this.getTxIDFromEventEtherem(ledgerEvent.data, targetIndex);
            default:
                logger.error(`##getTxIDFromEvent(): invalid verifierId: ${ledgerEvent.verifierId}`);
        }
        return null;
    }

    getTxIDFromEventSawtooth(event: object, targetIndex: number): string | null {
        logger.debug(`##in getTxIDFromEventSawtooth()`);
        const tx = this.getTransactionFromSawtoothEvent(event, targetIndex);
        if (tx == null) {
            logger.warn(`#getTxIDFromEventSawtooth(): skip(not found tx)`);
            return null;
        }

        try {
            const txId = tx['header_signature'];

            if (typeof txId !== 'string') {
                logger.warn(`#getTxIDFromEventSawtooth(): skip(invalid block, not found txId.), event: ${json2str(event)}`);
                return null;
            }

            logger.debug(`###getTxIDFromEventSawtooth(): txId: ${txId}`);
            return txId;

        }
        catch (err) {
            logger.error(`##getTxIDFromEventSawtooth(): err: ${err}, event: ${json2str(event)}`);
            return null;
        }
    }

    getTxIDFromEventEtherem(event: object, targetIndex: number): string | null {
        logger.debug(`##in getTxIDFromEventEtherem()`);
        const tx = this.getTransactionFromEthereumEvent(event, targetIndex);
        if (tx == null) {
            logger.warn(`#getTxIDFromEventEtherem(): skip(not found tx)`);
            return null;
        }

        try {
            const txId = tx['hash'];

            if (typeof txId !== 'string') {
                logger.warn(`#getTxIDFromEventEtherem(): skip(invalid block, not found txId.), event: ${json2str(event)}`);
                return null;
            }

            logger.debug(`###getTxIDFromEventEtherem(): txId: ${txId}`);
            return txId;

        }
        catch (err) {
            logger.error(`##getTxIDFromEventEtherem(): err: ${err}, event: ${json2str(event)}`);
            return null;
        }
    }

    getEventDataNum(ledgerEvent: LedgerEvent): number {
        logger.debug(`##in BLP:getEventDataNum(), ledgerEvent.verifierId: ${ledgerEvent.verifierId}`);
        const event = ledgerEvent.data;
        let retEventNum = 0;

        try {
            switch (ledgerEvent.verifierId) {
                case config.electricityTradeInfo.sawtooth.validatorID:
                    retEventNum = event['blockData'].length;
                    break;
                case config.electricityTradeInfo.ethereum.validatorID:
                    retEventNum = event['blockData']['transactions'].length;
                    break;
                default:
                    logger.error(`##getEventDataNum(): invalid verifierId: ${ledgerEvent.verifierId}`);
                    break;
            }
            logger.debug(`##getEventDataNum(): retEventNum: ${retEventNum}, verifierId: ${ledgerEvent.verifierId}`);
            return retEventNum;
        }
        catch (err) {
            logger.error(`##getEventDataNum(): invalid even, err: ${err}, event: ${event}`);
            return 0;
        }
    }

    getAccountInfo(transactionSubset: object): object {
        const transactionInfo = {};
        for (const customer of config.electricityTradeInfo.ethereum.account.customer) {
            logger.debug(`customer = ${json2str(customer)}`);
            if (transactionSubset['Name'] === customer.number) {
                if (transactionSubset['Verb'] === "inc") {
                    logger.debug('getAccountInfo(): A to B');
                    transactionInfo['fromAddress'] = customer.accountA;
                    transactionInfo['fromAddressPkey'] = config.electricityTradeInfo.ethereum.account.pKey[customer.accountA];
                    transactionInfo['toAddress'] = customer.accountB;
                }
                else if (transactionSubset['Verb'] === "dec"){
                    logger.debug('getAccountInfo(): B to A');
                    transactionInfo['fromAddress'] = customer.accountB;
                    transactionInfo['fromAddressPkey'] = config.electricityTradeInfo.ethereum.account.pKey[customer.accountB];
                    transactionInfo['toAddress'] = customer.accountA;
                }
            }
        }
        return transactionInfo;
    }
}
