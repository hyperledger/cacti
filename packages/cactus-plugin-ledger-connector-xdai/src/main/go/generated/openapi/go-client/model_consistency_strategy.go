/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
)

// checks if the ConsistencyStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsistencyStrategy{}

// ConsistencyStrategy struct for ConsistencyStrategy
type ConsistencyStrategy struct {
	ReceiptType ReceiptType `json:"receiptType"`
	// The amount of milliseconds to wait for the receipt to arrive to the connector. Defaults to 0 which means to wait for an unlimited amount of time. Note that this wait may be interrupted still by other parts of the infrastructure such as load balancers cutting of HTTP requests after some time even if they are the type that is supposed to be kept alive. The question of re-entrance is a broader topic not in scope to discuss here, but it is important to mention it.
	TimeoutMs *int32 `json:"timeoutMs,omitempty"`
	// The number of blocks to wait to be confirmed in addition to the block containing the transaction in question. Note that if the receipt type is set to only wait for node transaction pool ACK and this parameter is set to anything, but zero then the API will not accept the request due to conflicting parameters.
	BlockConfirmations int32 `json:"blockConfirmations"`
	// The amount of time (in milliseconds) connector will wait before making another confiramtion request to the network in case of previous confiramtion request fails
	PollIntervalMs *int32 `json:"pollIntervalMs,omitempty"`
}

// NewConsistencyStrategy instantiates a new ConsistencyStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsistencyStrategy(receiptType ReceiptType, blockConfirmations int32) *ConsistencyStrategy {
	this := ConsistencyStrategy{}
	this.ReceiptType = receiptType
	this.BlockConfirmations = blockConfirmations
	return &this
}

// NewConsistencyStrategyWithDefaults instantiates a new ConsistencyStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsistencyStrategyWithDefaults() *ConsistencyStrategy {
	this := ConsistencyStrategy{}
	return &this
}

// GetReceiptType returns the ReceiptType field value
func (o *ConsistencyStrategy) GetReceiptType() ReceiptType {
	if o == nil {
		var ret ReceiptType
		return ret
	}

	return o.ReceiptType
}

// GetReceiptTypeOk returns a tuple with the ReceiptType field value
// and a boolean to check if the value has been set.
func (o *ConsistencyStrategy) GetReceiptTypeOk() (*ReceiptType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptType, true
}

// SetReceiptType sets field value
func (o *ConsistencyStrategy) SetReceiptType(v ReceiptType) {
	o.ReceiptType = v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *ConsistencyStrategy) GetTimeoutMs() int32 {
	if o == nil || IsNil(o.TimeoutMs) {
		var ret int32
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyStrategy) GetTimeoutMsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutMs) {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *ConsistencyStrategy) HasTimeoutMs() bool {
	if o != nil && !IsNil(o.TimeoutMs) {
		return true
	}

	return false
}

// SetTimeoutMs gets a reference to the given int32 and assigns it to the TimeoutMs field.
func (o *ConsistencyStrategy) SetTimeoutMs(v int32) {
	o.TimeoutMs = &v
}

// GetBlockConfirmations returns the BlockConfirmations field value
func (o *ConsistencyStrategy) GetBlockConfirmations() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockConfirmations
}

// GetBlockConfirmationsOk returns a tuple with the BlockConfirmations field value
// and a boolean to check if the value has been set.
func (o *ConsistencyStrategy) GetBlockConfirmationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockConfirmations, true
}

// SetBlockConfirmations sets field value
func (o *ConsistencyStrategy) SetBlockConfirmations(v int32) {
	o.BlockConfirmations = v
}

// GetPollIntervalMs returns the PollIntervalMs field value if set, zero value otherwise.
func (o *ConsistencyStrategy) GetPollIntervalMs() int32 {
	if o == nil || IsNil(o.PollIntervalMs) {
		var ret int32
		return ret
	}
	return *o.PollIntervalMs
}

// GetPollIntervalMsOk returns a tuple with the PollIntervalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsistencyStrategy) GetPollIntervalMsOk() (*int32, bool) {
	if o == nil || IsNil(o.PollIntervalMs) {
		return nil, false
	}
	return o.PollIntervalMs, true
}

// HasPollIntervalMs returns a boolean if a field has been set.
func (o *ConsistencyStrategy) HasPollIntervalMs() bool {
	if o != nil && !IsNil(o.PollIntervalMs) {
		return true
	}

	return false
}

// SetPollIntervalMs gets a reference to the given int32 and assigns it to the PollIntervalMs field.
func (o *ConsistencyStrategy) SetPollIntervalMs(v int32) {
	o.PollIntervalMs = &v
}

func (o ConsistencyStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsistencyStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["receiptType"] = o.ReceiptType
	if !IsNil(o.TimeoutMs) {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	toSerialize["blockConfirmations"] = o.BlockConfirmations
	if !IsNil(o.PollIntervalMs) {
		toSerialize["pollIntervalMs"] = o.PollIntervalMs
	}
	return toSerialize, nil
}

type NullableConsistencyStrategy struct {
	value *ConsistencyStrategy
	isSet bool
}

func (v NullableConsistencyStrategy) Get() *ConsistencyStrategy {
	return v.value
}

func (v *NullableConsistencyStrategy) Set(val *ConsistencyStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableConsistencyStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableConsistencyStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsistencyStrategy(val *ConsistencyStrategy) *NullableConsistencyStrategy {
	return &NullableConsistencyStrategy{value: val, isSet: true}
}

func (v NullableConsistencyStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsistencyStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


