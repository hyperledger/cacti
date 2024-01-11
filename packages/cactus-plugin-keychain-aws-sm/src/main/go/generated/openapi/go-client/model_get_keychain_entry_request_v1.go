/*
Hyperledger Cactus - Keychain API

Contains/describes the Keychain API types/paths for Hyperledger Cactus.

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-keychain-aws-sm

import (
	"encoding/json"
)

// checks if the GetKeychainEntryRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetKeychainEntryRequestV1{}

// GetKeychainEntryRequestV1 struct for GetKeychainEntryRequestV1
type GetKeychainEntryRequestV1 struct {
	// The key for the entry to get from the keychain.
	Key string `json:"key"`
}

// NewGetKeychainEntryRequestV1 instantiates a new GetKeychainEntryRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetKeychainEntryRequestV1(key string) *GetKeychainEntryRequestV1 {
	this := GetKeychainEntryRequestV1{}
	this.Key = key
	return &this
}

// NewGetKeychainEntryRequestV1WithDefaults instantiates a new GetKeychainEntryRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetKeychainEntryRequestV1WithDefaults() *GetKeychainEntryRequestV1 {
	this := GetKeychainEntryRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *GetKeychainEntryRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetKeychainEntryRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetKeychainEntryRequestV1) SetKey(v string) {
	o.Key = v
}

func (o GetKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetKeychainEntryRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableGetKeychainEntryRequestV1 struct {
	value *GetKeychainEntryRequestV1
	isSet bool
}

func (v NullableGetKeychainEntryRequestV1) Get() *GetKeychainEntryRequestV1 {
	return v.value
}

func (v *NullableGetKeychainEntryRequestV1) Set(val *GetKeychainEntryRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetKeychainEntryRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetKeychainEntryRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetKeychainEntryRequestV1(val *GetKeychainEntryRequestV1) *NullableGetKeychainEntryRequestV1 {
	return &NullableGetKeychainEntryRequestV1{value: val, isSet: true}
}

func (v NullableGetKeychainEntryRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetKeychainEntryRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


