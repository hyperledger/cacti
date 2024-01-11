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

// checks if the IrohaBlockProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IrohaBlockProgress{}

// IrohaBlockProgress struct for IrohaBlockProgress
type IrohaBlockProgress struct {
	TransactionReceipt IrohaBlockResponse `json:"transactionReceipt"`
}

// NewIrohaBlockProgress instantiates a new IrohaBlockProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIrohaBlockProgress(transactionReceipt IrohaBlockResponse) *IrohaBlockProgress {
	this := IrohaBlockProgress{}
	this.TransactionReceipt = transactionReceipt
	return &this
}

// NewIrohaBlockProgressWithDefaults instantiates a new IrohaBlockProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIrohaBlockProgressWithDefaults() *IrohaBlockProgress {
	this := IrohaBlockProgress{}
	return &this
}

// GetTransactionReceipt returns the TransactionReceipt field value
func (o *IrohaBlockProgress) GetTransactionReceipt() IrohaBlockResponse {
	if o == nil {
		var ret IrohaBlockResponse
		return ret
	}

	return o.TransactionReceipt
}

// GetTransactionReceiptOk returns a tuple with the TransactionReceipt field value
// and a boolean to check if the value has been set.
func (o *IrohaBlockProgress) GetTransactionReceiptOk() (*IrohaBlockResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionReceipt, true
}

// SetTransactionReceipt sets field value
func (o *IrohaBlockProgress) SetTransactionReceipt(v IrohaBlockResponse) {
	o.TransactionReceipt = v
}

func (o IrohaBlockProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IrohaBlockProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionReceipt"] = o.TransactionReceipt
	return toSerialize, nil
}

type NullableIrohaBlockProgress struct {
	value *IrohaBlockProgress
	isSet bool
}

func (v NullableIrohaBlockProgress) Get() *IrohaBlockProgress {
	return v.value
}

func (v *NullableIrohaBlockProgress) Set(val *IrohaBlockProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaBlockProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaBlockProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaBlockProgress(val *IrohaBlockProgress) *NullableIrohaBlockProgress {
	return &NullableIrohaBlockProgress{value: val, isSet: true}
}

func (v NullableIrohaBlockProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaBlockProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


