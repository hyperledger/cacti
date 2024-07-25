/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the Iroha2AccountId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Iroha2AccountId{}

// Iroha2AccountId Iroha V2 account ID.
type Iroha2AccountId struct {
	Name string `json:"name"`
	DomainId string `json:"domainId"`
}

// NewIroha2AccountId instantiates a new Iroha2AccountId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIroha2AccountId(name string, domainId string) *Iroha2AccountId {
	this := Iroha2AccountId{}
	this.Name = name
	this.DomainId = domainId
	return &this
}

// NewIroha2AccountIdWithDefaults instantiates a new Iroha2AccountId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIroha2AccountIdWithDefaults() *Iroha2AccountId {
	this := Iroha2AccountId{}
	return &this
}

// GetName returns the Name field value
func (o *Iroha2AccountId) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Iroha2AccountId) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Iroha2AccountId) SetName(v string) {
	o.Name = v
}

// GetDomainId returns the DomainId field value
func (o *Iroha2AccountId) GetDomainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *Iroha2AccountId) GetDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value
func (o *Iroha2AccountId) SetDomainId(v string) {
	o.DomainId = v
}

func (o Iroha2AccountId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Iroha2AccountId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["domainId"] = o.DomainId
	return toSerialize, nil
}

type NullableIroha2AccountId struct {
	value *Iroha2AccountId
	isSet bool
}

func (v NullableIroha2AccountId) Get() *Iroha2AccountId {
	return v.value
}

func (v *NullableIroha2AccountId) Set(val *Iroha2AccountId) {
	v.value = val
	v.isSet = true
}

func (v NullableIroha2AccountId) IsSet() bool {
	return v.isSet
}

func (v *NullableIroha2AccountId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIroha2AccountId(val *Iroha2AccountId) *NullableIroha2AccountId {
	return &NullableIroha2AccountId{value: val, isSet: true}
}

func (v NullableIroha2AccountId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIroha2AccountId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


