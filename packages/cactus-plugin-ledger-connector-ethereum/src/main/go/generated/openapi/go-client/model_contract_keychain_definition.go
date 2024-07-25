/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the ContractKeychainDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractKeychainDefinition{}

// ContractKeychainDefinition struct for ContractKeychainDefinition
type ContractKeychainDefinition struct {
	// The contract name for retrieve the contracts json on the keychain.
	ContractName string `json:"contractName"`
	// The keychainId for retrieve the contracts json.
	KeychainId string `json:"keychainId"`
}

// NewContractKeychainDefinition instantiates a new ContractKeychainDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractKeychainDefinition(contractName string, keychainId string) *ContractKeychainDefinition {
	this := ContractKeychainDefinition{}
	this.ContractName = contractName
	this.KeychainId = keychainId
	return &this
}

// NewContractKeychainDefinitionWithDefaults instantiates a new ContractKeychainDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractKeychainDefinitionWithDefaults() *ContractKeychainDefinition {
	this := ContractKeychainDefinition{}
	return &this
}

// GetContractName returns the ContractName field value
func (o *ContractKeychainDefinition) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *ContractKeychainDefinition) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *ContractKeychainDefinition) SetContractName(v string) {
	o.ContractName = v
}

// GetKeychainId returns the KeychainId field value
func (o *ContractKeychainDefinition) GetKeychainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value
// and a boolean to check if the value has been set.
func (o *ContractKeychainDefinition) GetKeychainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainId, true
}

// SetKeychainId sets field value
func (o *ContractKeychainDefinition) SetKeychainId(v string) {
	o.KeychainId = v
}

func (o ContractKeychainDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractKeychainDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractName"] = o.ContractName
	toSerialize["keychainId"] = o.KeychainId
	return toSerialize, nil
}

type NullableContractKeychainDefinition struct {
	value *ContractKeychainDefinition
	isSet bool
}

func (v NullableContractKeychainDefinition) Get() *ContractKeychainDefinition {
	return v.value
}

func (v *NullableContractKeychainDefinition) Set(val *ContractKeychainDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableContractKeychainDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableContractKeychainDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractKeychainDefinition(val *ContractKeychainDefinition) *NullableContractKeychainDefinition {
	return &NullableContractKeychainDefinition{value: val, isSet: true}
}

func (v NullableContractKeychainDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractKeychainDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


