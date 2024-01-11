/*
Hyperledger Cactus Plugin - Connector Iroha

Can perform basic tasks on a Iroha ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha

import (
	"encoding/json"
)

// checks if the AddSignatoryRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSignatoryRequestParameters{}

// AddSignatoryRequestParameters The list of arguments to pass in to the transaction request to Add Signatory.
type AddSignatoryRequestParameters struct {
	AccountId string `json:"accountId"`
	PublicKey string `json:"publicKey"`
}

// NewAddSignatoryRequestParameters instantiates a new AddSignatoryRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSignatoryRequestParameters(accountId string, publicKey string) *AddSignatoryRequestParameters {
	this := AddSignatoryRequestParameters{}
	this.AccountId = accountId
	this.PublicKey = publicKey
	return &this
}

// NewAddSignatoryRequestParametersWithDefaults instantiates a new AddSignatoryRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSignatoryRequestParametersWithDefaults() *AddSignatoryRequestParameters {
	this := AddSignatoryRequestParameters{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AddSignatoryRequestParameters) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AddSignatoryRequestParameters) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AddSignatoryRequestParameters) SetAccountId(v string) {
	o.AccountId = v
}

// GetPublicKey returns the PublicKey field value
func (o *AddSignatoryRequestParameters) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *AddSignatoryRequestParameters) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *AddSignatoryRequestParameters) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o AddSignatoryRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSignatoryRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

type NullableAddSignatoryRequestParameters struct {
	value *AddSignatoryRequestParameters
	isSet bool
}

func (v NullableAddSignatoryRequestParameters) Get() *AddSignatoryRequestParameters {
	return v.value
}

func (v *NullableAddSignatoryRequestParameters) Set(val *AddSignatoryRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSignatoryRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSignatoryRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSignatoryRequestParameters(val *AddSignatoryRequestParameters) *NullableAddSignatoryRequestParameters {
	return &NullableAddSignatoryRequestParameters{value: val, isSet: true}
}

func (v NullableAddSignatoryRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSignatoryRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


