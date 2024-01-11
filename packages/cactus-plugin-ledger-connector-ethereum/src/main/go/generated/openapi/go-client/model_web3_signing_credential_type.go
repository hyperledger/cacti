/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
	"fmt"
)

// Web3SigningCredentialType the model 'Web3SigningCredentialType'
type Web3SigningCredentialType string

// List of Web3SigningCredentialType
const (
	CACTI_KEYCHAIN_REF Web3SigningCredentialType = "CACTI_KEYCHAIN_REF"
	GETH_KEYCHAIN_PASSWORD Web3SigningCredentialType = "GETH_KEYCHAIN_PASSWORD"
	PRIVATE_KEY_HEX Web3SigningCredentialType = "PRIVATE_KEY_HEX"
	NONE Web3SigningCredentialType = "NONE"
)

// All allowed values of Web3SigningCredentialType enum
var AllowedWeb3SigningCredentialTypeEnumValues = []Web3SigningCredentialType{
	"CACTI_KEYCHAIN_REF",
	"GETH_KEYCHAIN_PASSWORD",
	"PRIVATE_KEY_HEX",
	"NONE",
}

func (v *Web3SigningCredentialType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Web3SigningCredentialType(value)
	for _, existing := range AllowedWeb3SigningCredentialTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Web3SigningCredentialType", value)
}

// NewWeb3SigningCredentialTypeFromValue returns a pointer to a valid Web3SigningCredentialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWeb3SigningCredentialTypeFromValue(v string) (*Web3SigningCredentialType, error) {
	ev := Web3SigningCredentialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Web3SigningCredentialType: valid values are %v", v, AllowedWeb3SigningCredentialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Web3SigningCredentialType) IsValid() bool {
	for _, existing := range AllowedWeb3SigningCredentialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Web3SigningCredentialType value
func (v Web3SigningCredentialType) Ptr() *Web3SigningCredentialType {
	return &v
}

type NullableWeb3SigningCredentialType struct {
	value *Web3SigningCredentialType
	isSet bool
}

func (v NullableWeb3SigningCredentialType) Get() *Web3SigningCredentialType {
	return v.value
}

func (v *NullableWeb3SigningCredentialType) Set(val *Web3SigningCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3SigningCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3SigningCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3SigningCredentialType(val *Web3SigningCredentialType) *NullableWeb3SigningCredentialType {
	return &NullableWeb3SigningCredentialType{value: val, isSet: true}
}

func (v NullableWeb3SigningCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3SigningCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

