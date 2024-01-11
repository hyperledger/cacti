/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
	"fmt"
)

// EthContractInvocationType the model 'EthContractInvocationType'
type EthContractInvocationType string

// List of EthContractInvocationType
const (
	SEND EthContractInvocationType = "SEND"
	CALL EthContractInvocationType = "CALL"
)

// All allowed values of EthContractInvocationType enum
var AllowedEthContractInvocationTypeEnumValues = []EthContractInvocationType{
	"SEND",
	"CALL",
}

func (v *EthContractInvocationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EthContractInvocationType(value)
	for _, existing := range AllowedEthContractInvocationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EthContractInvocationType", value)
}

// NewEthContractInvocationTypeFromValue returns a pointer to a valid EthContractInvocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEthContractInvocationTypeFromValue(v string) (*EthContractInvocationType, error) {
	ev := EthContractInvocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EthContractInvocationType: valid values are %v", v, AllowedEthContractInvocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EthContractInvocationType) IsValid() bool {
	for _, existing := range AllowedEthContractInvocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EthContractInvocationType value
func (v EthContractInvocationType) Ptr() *EthContractInvocationType {
	return &v
}

type NullableEthContractInvocationType struct {
	value *EthContractInvocationType
	isSet bool
}

func (v NullableEthContractInvocationType) Get() *EthContractInvocationType {
	return v.value
}

func (v *NullableEthContractInvocationType) Set(val *EthContractInvocationType) {
	v.value = val
	v.isSet = true
}

func (v NullableEthContractInvocationType) IsSet() bool {
	return v.isSet
}

func (v *NullableEthContractInvocationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEthContractInvocationType(val *EthContractInvocationType) *NullableEthContractInvocationType {
	return &NullableEthContractInvocationType{value: val, isSet: true}
}

func (v NullableEthContractInvocationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEthContractInvocationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

