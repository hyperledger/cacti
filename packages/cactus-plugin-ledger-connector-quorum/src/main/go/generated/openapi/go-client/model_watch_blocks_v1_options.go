/*
Hyperledger Cactus Plugin - Connector Quorum

Can perform basic tasks on a Quorum ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-quorum

import (
	"encoding/json"
)

// checks if the WatchBlocksV1Options type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksV1Options{}

// WatchBlocksV1Options struct for WatchBlocksV1Options
type WatchBlocksV1Options struct {
	GetBlockData *bool `json:"getBlockData,omitempty"`
}

// NewWatchBlocksV1Options instantiates a new WatchBlocksV1Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksV1Options() *WatchBlocksV1Options {
	this := WatchBlocksV1Options{}
	return &this
}

// NewWatchBlocksV1OptionsWithDefaults instantiates a new WatchBlocksV1Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksV1OptionsWithDefaults() *WatchBlocksV1Options {
	this := WatchBlocksV1Options{}
	return &this
}

// GetGetBlockData returns the GetBlockData field value if set, zero value otherwise.
func (o *WatchBlocksV1Options) GetGetBlockData() bool {
	if o == nil || IsNil(o.GetBlockData) {
		var ret bool
		return ret
	}
	return *o.GetBlockData
}

// GetGetBlockDataOk returns a tuple with the GetBlockData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchBlocksV1Options) GetGetBlockDataOk() (*bool, bool) {
	if o == nil || IsNil(o.GetBlockData) {
		return nil, false
	}
	return o.GetBlockData, true
}

// HasGetBlockData returns a boolean if a field has been set.
func (o *WatchBlocksV1Options) HasGetBlockData() bool {
	if o != nil && !IsNil(o.GetBlockData) {
		return true
	}

	return false
}

// SetGetBlockData gets a reference to the given bool and assigns it to the GetBlockData field.
func (o *WatchBlocksV1Options) SetGetBlockData(v bool) {
	o.GetBlockData = &v
}

func (o WatchBlocksV1Options) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksV1Options) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GetBlockData) {
		toSerialize["getBlockData"] = o.GetBlockData
	}
	return toSerialize, nil
}

type NullableWatchBlocksV1Options struct {
	value *WatchBlocksV1Options
	isSet bool
}

func (v NullableWatchBlocksV1Options) Get() *WatchBlocksV1Options {
	return v.value
}

func (v *NullableWatchBlocksV1Options) Set(val *WatchBlocksV1Options) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksV1Options) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksV1Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksV1Options(val *WatchBlocksV1Options) *NullableWatchBlocksV1Options {
	return &NullableWatchBlocksV1Options{value: val, isSet: true}
}

func (v NullableWatchBlocksV1Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksV1Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


