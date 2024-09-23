/*
Hyperledger Cactus Plugin - Connector Polkadot

Can perform basic tasks on a Polkadot parachain

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-polkadot

import (
	"encoding/json"
)

// checks if the RunTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionResponse{}

// RunTransactionResponse struct for RunTransactionResponse
type RunTransactionResponse struct {
	Success bool `json:"success"`
	TxHash *string `json:"txHash,omitempty"`
	BlockHash *string `json:"blockHash,omitempty"`
}

// NewRunTransactionResponse instantiates a new RunTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionResponse(success bool) *RunTransactionResponse {
	this := RunTransactionResponse{}
	this.Success = success
	return &this
}

// NewRunTransactionResponseWithDefaults instantiates a new RunTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionResponseWithDefaults() *RunTransactionResponse {
	this := RunTransactionResponse{}
	return &this
}

// GetSuccess returns the Success field value
func (o *RunTransactionResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *RunTransactionResponse) SetSuccess(v bool) {
	o.Success = v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *RunTransactionResponse) GetTxHash() string {
	if o == nil || IsNil(o.TxHash) {
		var ret string
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetTxHashOk() (*string, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *RunTransactionResponse) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given string and assigns it to the TxHash field.
func (o *RunTransactionResponse) SetTxHash(v string) {
	o.TxHash = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *RunTransactionResponse) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionResponse) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *RunTransactionResponse) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *RunTransactionResponse) SetBlockHash(v string) {
	o.BlockHash = &v
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
	toSerialize["success"] = o.Success
	if !IsNil(o.TxHash) {
		toSerialize["txHash"] = o.TxHash
	}
	if !IsNil(o.BlockHash) {
		toSerialize["blockHash"] = o.BlockHash
	}
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

