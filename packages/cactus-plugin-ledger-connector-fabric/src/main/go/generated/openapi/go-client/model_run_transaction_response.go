/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the RunTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionResponse{}

// RunTransactionResponse struct for RunTransactionResponse
type RunTransactionResponse struct {
	FunctionOutput string `json:"functionOutput"`
	TransactionId string `json:"transactionId"`
}

// NewRunTransactionResponse instantiates a new RunTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionResponse(functionOutput string, transactionId string) *RunTransactionResponse {
	this := RunTransactionResponse{}
	this.FunctionOutput = functionOutput
	this.TransactionId = transactionId
	return &this
}

// NewRunTransactionResponseWithDefaults instantiates a new RunTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionResponseWithDefaults() *RunTransactionResponse {
	this := RunTransactionResponse{}
	return &this
}

// GetFunctionOutput returns the FunctionOutput field value
func (o *RunTransactionResponse) GetFunctionOutput() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionOutput
}

// GetFunctionOutputOk returns a tuple with the FunctionOutput field value
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetFunctionOutputOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionOutput, true
}

// SetFunctionOutput sets field value
func (o *RunTransactionResponse) SetFunctionOutput(v string) {
	o.FunctionOutput = v
}

// GetTransactionId returns the TransactionId field value
func (o *RunTransactionResponse) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *RunTransactionResponse) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o RunTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["functionOutput"] = o.FunctionOutput
	toSerialize["transactionId"] = o.TransactionId
	return toSerialize, nil
}

type NullableRunTransactionResponse struct {
	value *RunTransactionResponse
	isSet bool
}

func (v NullableRunTransactionResponse) Get() *RunTransactionResponse {
	return v.value
}

func (v *NullableRunTransactionResponse) Set(val *RunTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionResponse(val *RunTransactionResponse) *NullableRunTransactionResponse {
	return &NullableRunTransactionResponse{value: val, isSet: true}
}

func (v NullableRunTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


