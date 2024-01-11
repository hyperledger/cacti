/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the WatchBlocksFullResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchBlocksFullResponseV1{}

// WatchBlocksFullResponseV1 Response that corresponds to Fabric SDK 'full' EventType.
type WatchBlocksFullResponseV1 struct {
	// Full commited block.
	FullBlock interface{} `json:"fullBlock"`
}

// NewWatchBlocksFullResponseV1 instantiates a new WatchBlocksFullResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchBlocksFullResponseV1(fullBlock interface{}) *WatchBlocksFullResponseV1 {
	this := WatchBlocksFullResponseV1{}
	this.FullBlock = fullBlock
	return &this
}

// NewWatchBlocksFullResponseV1WithDefaults instantiates a new WatchBlocksFullResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchBlocksFullResponseV1WithDefaults() *WatchBlocksFullResponseV1 {
	this := WatchBlocksFullResponseV1{}
	return &this
}

// GetFullBlock returns the FullBlock field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *WatchBlocksFullResponseV1) GetFullBlock() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.FullBlock
}

// GetFullBlockOk returns a tuple with the FullBlock field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WatchBlocksFullResponseV1) GetFullBlockOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FullBlock) {
		return nil, false
	}
	return &o.FullBlock, true
}

// SetFullBlock sets field value
func (o *WatchBlocksFullResponseV1) SetFullBlock(v interface{}) {
	o.FullBlock = v
}

func (o WatchBlocksFullResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchBlocksFullResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FullBlock != nil {
		toSerialize["fullBlock"] = o.FullBlock
	}
	return toSerialize, nil
}

type NullableWatchBlocksFullResponseV1 struct {
	value *WatchBlocksFullResponseV1
	isSet bool
}

func (v NullableWatchBlocksFullResponseV1) Get() *WatchBlocksFullResponseV1 {
	return v.value
}

func (v *NullableWatchBlocksFullResponseV1) Set(val *WatchBlocksFullResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchBlocksFullResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchBlocksFullResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchBlocksFullResponseV1(val *WatchBlocksFullResponseV1) *NullableWatchBlocksFullResponseV1 {
	return &NullableWatchBlocksFullResponseV1{value: val, isSet: true}
}

func (v NullableWatchBlocksFullResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchBlocksFullResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


