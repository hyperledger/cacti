/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the Consortium type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Consortium{}

// Consortium struct for Consortium
type Consortium struct {
	Id string `json:"id"`
	Name string `json:"name"`
	MainApiHost string `json:"mainApiHost"`
	// The collection (array) of primary keys of consortium member entities that belong to this Consortium.
	MemberIds []string `json:"memberIds"`
}

// NewConsortium instantiates a new Consortium object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsortium(id string, name string, mainApiHost string, memberIds []string) *Consortium {
	this := Consortium{}
	this.Id = id
	this.Name = name
	this.MainApiHost = mainApiHost
	this.MemberIds = memberIds
	return &this
}

// NewConsortiumWithDefaults instantiates a new Consortium object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsortiumWithDefaults() *Consortium {
	this := Consortium{}
	return &this
}

// GetId returns the Id field value
func (o *Consortium) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Consortium) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Consortium) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Consortium) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Consortium) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Consortium) SetName(v string) {
	o.Name = v
}

// GetMainApiHost returns the MainApiHost field value
func (o *Consortium) GetMainApiHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MainApiHost
}

// GetMainApiHostOk returns a tuple with the MainApiHost field value
// and a boolean to check if the value has been set.
func (o *Consortium) GetMainApiHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MainApiHost, true
}

// SetMainApiHost sets field value
func (o *Consortium) SetMainApiHost(v string) {
	o.MainApiHost = v
}

// GetMemberIds returns the MemberIds field value
func (o *Consortium) GetMemberIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MemberIds
}

// GetMemberIdsOk returns a tuple with the MemberIds field value
// and a boolean to check if the value has been set.
func (o *Consortium) GetMemberIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemberIds, true
}

// SetMemberIds sets field value
func (o *Consortium) SetMemberIds(v []string) {
	o.MemberIds = v
}

func (o Consortium) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Consortium) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["mainApiHost"] = o.MainApiHost
	toSerialize["memberIds"] = o.MemberIds
	return toSerialize, nil
}

type NullableConsortium struct {
	value *Consortium
	isSet bool
}

func (v NullableConsortium) Get() *Consortium {
	return v.value
}

func (v *NullableConsortium) Set(val *Consortium) {
	v.value = val
	v.isSet = true
}

func (v NullableConsortium) IsSet() bool {
	return v.isSet
}

func (v *NullableConsortium) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsortium(val *Consortium) *NullableConsortium {
	return &NullableConsortium{value: val, isSet: true}
}

func (v NullableConsortium) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsortium) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


