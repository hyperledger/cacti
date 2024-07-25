/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the CactusNodeMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CactusNodeMeta{}

// CactusNodeMeta A Cactus node meta information
type CactusNodeMeta struct {
	NodeApiHost string `json:"nodeApiHost"`
	// The PEM encoded public key that was used to generate the JWS included in the response (the jws property)
	PublicKeyPem string `json:"publicKeyPem"`
}

// NewCactusNodeMeta instantiates a new CactusNodeMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCactusNodeMeta(nodeApiHost string, publicKeyPem string) *CactusNodeMeta {
	this := CactusNodeMeta{}
	this.NodeApiHost = nodeApiHost
	this.PublicKeyPem = publicKeyPem
	return &this
}

// NewCactusNodeMetaWithDefaults instantiates a new CactusNodeMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCactusNodeMetaWithDefaults() *CactusNodeMeta {
	this := CactusNodeMeta{}
	return &this
}

// GetNodeApiHost returns the NodeApiHost field value
func (o *CactusNodeMeta) GetNodeApiHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeApiHost
}

// GetNodeApiHostOk returns a tuple with the NodeApiHost field value
// and a boolean to check if the value has been set.
func (o *CactusNodeMeta) GetNodeApiHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeApiHost, true
}

// SetNodeApiHost sets field value
func (o *CactusNodeMeta) SetNodeApiHost(v string) {
	o.NodeApiHost = v
}

// GetPublicKeyPem returns the PublicKeyPem field value
func (o *CactusNodeMeta) GetPublicKeyPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyPem
}

// GetPublicKeyPemOk returns a tuple with the PublicKeyPem field value
// and a boolean to check if the value has been set.
func (o *CactusNodeMeta) GetPublicKeyPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyPem, true
}

// SetPublicKeyPem sets field value
func (o *CactusNodeMeta) SetPublicKeyPem(v string) {
	o.PublicKeyPem = v
}

func (o CactusNodeMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CactusNodeMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeApiHost"] = o.NodeApiHost
	toSerialize["publicKeyPem"] = o.PublicKeyPem
	return toSerialize, nil
}

type NullableCactusNodeMeta struct {
	value *CactusNodeMeta
	isSet bool
}

func (v NullableCactusNodeMeta) Get() *CactusNodeMeta {
	return v.value
}

func (v *NullableCactusNodeMeta) Set(val *CactusNodeMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCactusNodeMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCactusNodeMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCactusNodeMeta(val *CactusNodeMeta) *NullableCactusNodeMeta {
	return &NullableCactusNodeMeta{value: val, isSet: true}
}

func (v NullableCactusNodeMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCactusNodeMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


