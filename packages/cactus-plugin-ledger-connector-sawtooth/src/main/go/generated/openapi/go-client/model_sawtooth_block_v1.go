/*
Hyperledger Cacti Plugin - Connector Sawtooth

Can perform basic tasks on a Sawtooth ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-sawtooth

import (
	"encoding/json"
)

// checks if the SawtoothBlockV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SawtoothBlockV1{}

// SawtoothBlockV1 struct for SawtoothBlockV1
type SawtoothBlockV1 struct {
	Header SawtoothBlockHeaderV1 `json:"header"`
	HeaderSignature string `json:"header_signature"`
	Batches []SawtoothBatchV1 `json:"batches"`
}

// NewSawtoothBlockV1 instantiates a new SawtoothBlockV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSawtoothBlockV1(header SawtoothBlockHeaderV1, headerSignature string, batches []SawtoothBatchV1) *SawtoothBlockV1 {
	this := SawtoothBlockV1{}
	this.Header = header
	this.HeaderSignature = headerSignature
	this.Batches = batches
	return &this
}

// NewSawtoothBlockV1WithDefaults instantiates a new SawtoothBlockV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSawtoothBlockV1WithDefaults() *SawtoothBlockV1 {
	this := SawtoothBlockV1{}
	return &this
}

// GetHeader returns the Header field value
func (o *SawtoothBlockV1) GetHeader() SawtoothBlockHeaderV1 {
	if o == nil {
		var ret SawtoothBlockHeaderV1
		return ret
	}

	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value
// and a boolean to check if the value has been set.
func (o *SawtoothBlockV1) GetHeaderOk() (*SawtoothBlockHeaderV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Header, true
}

// SetHeader sets field value
func (o *SawtoothBlockV1) SetHeader(v SawtoothBlockHeaderV1) {
	o.Header = v
}

// GetHeaderSignature returns the HeaderSignature field value
func (o *SawtoothBlockV1) GetHeaderSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderSignature
}

// GetHeaderSignatureOk returns a tuple with the HeaderSignature field value
// and a boolean to check if the value has been set.
func (o *SawtoothBlockV1) GetHeaderSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderSignature, true
}

// SetHeaderSignature sets field value
func (o *SawtoothBlockV1) SetHeaderSignature(v string) {
	o.HeaderSignature = v
}

// GetBatches returns the Batches field value
func (o *SawtoothBlockV1) GetBatches() []SawtoothBatchV1 {
	if o == nil {
		var ret []SawtoothBatchV1
		return ret
	}

	return o.Batches
}

// GetBatchesOk returns a tuple with the Batches field value
// and a boolean to check if the value has been set.
func (o *SawtoothBlockV1) GetBatchesOk() ([]SawtoothBatchV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Batches, true
}

// SetBatches sets field value
func (o *SawtoothBlockV1) SetBatches(v []SawtoothBatchV1) {
	o.Batches = v
}

func (o SawtoothBlockV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SawtoothBlockV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["header"] = o.Header
	toSerialize["header_signature"] = o.HeaderSignature
	toSerialize["batches"] = o.Batches
	return toSerialize, nil
}

type NullableSawtoothBlockV1 struct {
	value *SawtoothBlockV1
	isSet bool
}

func (v NullableSawtoothBlockV1) Get() *SawtoothBlockV1 {
	return v.value
}

func (v *NullableSawtoothBlockV1) Set(val *SawtoothBlockV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSawtoothBlockV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSawtoothBlockV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSawtoothBlockV1(val *SawtoothBlockV1) *NullableSawtoothBlockV1 {
	return &NullableSawtoothBlockV1{value: val, isSet: true}
}

func (v NullableSawtoothBlockV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSawtoothBlockV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


