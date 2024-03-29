/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an BondAsset and TokenAsset
type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface, ccType string) error {
	if ccType == "Bond" {
		return s.InitBondAssetLedger(ctx)
	} else if ccType == "Token" {
		return s.InitTokenAssetLedger(ctx)
	}
	return fmt.Errorf("only Bond and Token are accepted as ccType.")
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	_, ok := os.LookupEnv("EXTERNAL_SERVICE")
	if ok {
		server := &shim.ChaincodeServer{
			CCID:    os.Getenv("CHAINCODE_CCID"),
			Address: os.Getenv("CHAINCODE_ADDRESS"),
			CC:      chaincode,
			TLSProps: shim.TLSProperties{
				Disabled: true,
			},
		}
		// Start the chaincode external server
		err = server.Start()
	} else {
		err = chaincode.Start()
	}
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}

}
