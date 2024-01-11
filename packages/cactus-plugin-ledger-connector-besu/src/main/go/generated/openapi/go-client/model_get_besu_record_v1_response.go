/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the GetBesuRecordV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBesuRecordV1Response{}

// GetBesuRecordV1Response struct for GetBesuRecordV1Response
type GetBesuRecordV1Response struct {
	LedgerId *string `json:"ledgerId,omitempty"`
	StateContract *string `json:"stateContract,omitempty"`
	TransactionInputData interface{} `json:"transactionInputData,omitempty"`
	CallOutput interface{} `json:"callOutput,omitempty"`
}

// NewGetBesuRecordV1Response instantiates a new GetBesuRecordV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBesuRecordV1Response() *GetBesuRecordV1Response {
	this := GetBesuRecordV1Response{}
	return &this
}

// NewGetBesuRecordV1ResponseWithDefaults instantiates a new GetBesuRecordV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBesuRecordV1ResponseWithDefaults() *GetBesuRecordV1Response {
	this := GetBesuRecordV1Response{}
	return &this
}

// GetLedgerId returns the LedgerId field value if set, zero value otherwise.
func (o *GetBesuRecordV1Response) GetLedgerId() string {
	if o == nil || IsNil(o.LedgerId) {
		var ret string
		return ret
	}
	return *o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBesuRecordV1Response) GetLedgerIdOk() (*string, bool) {
	if o == nil || IsNil(o.LedgerId) {
		return nil, false
	}
	return o.LedgerId, true
}

// HasLedgerId returns a boolean if a field has been set.
func (o *GetBesuRecordV1Response) HasLedgerId() bool {
	if o != nil && !IsNil(o.LedgerId) {
		return true
	}

	return false
}

// SetLedgerId gets a reference to the given string and assigns it to the LedgerId field.
func (o *GetBesuRecordV1Response) SetLedgerId(v string) {
	o.LedgerId = &v
}

// GetStateContract returns the StateContract field value if set, zero value otherwise.
func (o *GetBesuRecordV1Response) GetStateContract() string {
	if o == nil || IsNil(o.StateContract) {
		var ret string
		return ret
	}
	return *o.StateContract
}

// GetStateContractOk returns a tuple with the StateContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBesuRecordV1Response) GetStateContractOk() (*string, bool) {
	if o == nil || IsNil(o.StateContract) {
		return nil, false
	}
	return o.StateContract, true
}

// HasStateContract returns a boolean if a field has been set.
func (o *GetBesuRecordV1Response) HasStateContract() bool {
	if o != nil && !IsNil(o.StateContract) {
		return true
	}

	return false
}

// SetStateContract gets a reference to the given string and assigns it to the StateContract field.
func (o *GetBesuRecordV1Response) SetStateContract(v string) {
	o.StateContract = &v
}

// GetTransactionInputData returns the TransactionInputData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetBesuRecordV1Response) GetTransactionInputData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TransactionInputData
}

// GetTransactionInputDataOk returns a tuple with the TransactionInputData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBesuRecordV1Response) GetTransactionInputDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TransactionInputData) {
		return nil, false
	}
	return &o.TransactionInputData, true
}

// HasTransactionInputData returns a boolean if a field has been set.
func (o *GetBesuRecordV1Response) HasTransactionInputData() bool {
	if o != nil && IsNil(o.TransactionInputData) {
		return true
	}

	return false
}

// SetTransactionInputData gets a reference to the given interface{} and assigns it to the TransactionInputData field.
func (o *GetBesuRecordV1Response) SetTransactionInputData(v interface{}) {
	o.TransactionInputData = v
}

// GetCallOutput returns the CallOutput field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetBesuRecordV1Response) GetCallOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CallOutput
}

// GetCallOutputOk returns a tuple with the CallOutput field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetBesuRecordV1Response) GetCallOutputOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CallOutput) {
		return nil, false
	}
	return &o.CallOutput, true
}

// HasCallOutput returns a boolean if a field has been set.
func (o *GetBesuRecordV1Response) HasCallOutput() bool {
	if o != nil && IsNil(o.CallOutput) {
		return true
	}

	return false
}

// SetCallOutput gets a reference to the given interface{} and assigns it to the CallOutput field.
func (o *GetBesuRecordV1Response) SetCallOutput(v interface{}) {
	o.CallOutput = v
}

func (o GetBesuRecordV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBesuRecordV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LedgerId) {
		toSerialize["ledgerId"] = o.LedgerId
	}
	if !IsNil(o.StateContract) {
		toSerialize["stateContract"] = o.StateContract
	}
	if o.TransactionInputData != nil {
		toSerialize["transactionInputData"] = o.TransactionInputData
	}
	if o.CallOutput != nil {
		toSerialize["callOutput"] = o.CallOutput
	}
	return toSerialize, nil
}

type NullableGetBesuRecordV1Response struct {
	value *GetBesuRecordV1Response
	isSet bool
}

func (v NullableGetBesuRecordV1Response) Get() *GetBesuRecordV1Response {
	return v.value
}

func (v *NullableGetBesuRecordV1Response) Set(val *GetBesuRecordV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBesuRecordV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBesuRecordV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBesuRecordV1Response(val *GetBesuRecordV1Response) *NullableGetBesuRecordV1Response {
	return &NullableGetBesuRecordV1Response{value: val, isSet: true}
}

func (v NullableGetBesuRecordV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBesuRecordV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


