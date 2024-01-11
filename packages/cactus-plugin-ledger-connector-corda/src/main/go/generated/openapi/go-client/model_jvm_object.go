/*
Hyperledger Cactus Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the JvmObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JvmObject{}

// JvmObject Can represent JVM primitive and reference types as well. The jvmTypeKind field indicates which one is being stored. If the jvmTypeKind field is set to REFERENCE then the jvmCtorArgs array is expected to be filled, otherwise (e.g. PRIMITIVE jvmTypeKind) it is expected that the primitiveValue property is filled with a primitive data type supported by the JSON standard such as strings, booleans, numbers, etc.
type JvmObject struct {
	JvmTypeKind JvmTypeKind `json:"jvmTypeKind"`
	PrimitiveValue map[string]interface{} `json:"primitiveValue,omitempty"`
	JvmCtorArgs []JvmObject `json:"jvmCtorArgs,omitempty"`
	JvmType JvmType `json:"jvmType"`
}

// NewJvmObject instantiates a new JvmObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJvmObject(jvmTypeKind JvmTypeKind, jvmType JvmType) *JvmObject {
	this := JvmObject{}
	this.JvmTypeKind = jvmTypeKind
	this.JvmType = jvmType
	return &this
}

// NewJvmObjectWithDefaults instantiates a new JvmObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJvmObjectWithDefaults() *JvmObject {
	this := JvmObject{}
	return &this
}

// GetJvmTypeKind returns the JvmTypeKind field value
func (o *JvmObject) GetJvmTypeKind() JvmTypeKind {
	if o == nil {
		var ret JvmTypeKind
		return ret
	}

	return o.JvmTypeKind
}

// GetJvmTypeKindOk returns a tuple with the JvmTypeKind field value
// and a boolean to check if the value has been set.
func (o *JvmObject) GetJvmTypeKindOk() (*JvmTypeKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JvmTypeKind, true
}

// SetJvmTypeKind sets field value
func (o *JvmObject) SetJvmTypeKind(v JvmTypeKind) {
	o.JvmTypeKind = v
}

// GetPrimitiveValue returns the PrimitiveValue field value if set, zero value otherwise.
func (o *JvmObject) GetPrimitiveValue() map[string]interface{} {
	if o == nil || IsNil(o.PrimitiveValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.PrimitiveValue
}

// GetPrimitiveValueOk returns a tuple with the PrimitiveValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JvmObject) GetPrimitiveValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PrimitiveValue) {
		return map[string]interface{}{}, false
	}
	return o.PrimitiveValue, true
}

// HasPrimitiveValue returns a boolean if a field has been set.
func (o *JvmObject) HasPrimitiveValue() bool {
	if o != nil && !IsNil(o.PrimitiveValue) {
		return true
	}

	return false
}

// SetPrimitiveValue gets a reference to the given map[string]interface{} and assigns it to the PrimitiveValue field.
func (o *JvmObject) SetPrimitiveValue(v map[string]interface{}) {
	o.PrimitiveValue = v
}

// GetJvmCtorArgs returns the JvmCtorArgs field value if set, zero value otherwise.
func (o *JvmObject) GetJvmCtorArgs() []JvmObject {
	if o == nil || IsNil(o.JvmCtorArgs) {
		var ret []JvmObject
		return ret
	}
	return o.JvmCtorArgs
}

// GetJvmCtorArgsOk returns a tuple with the JvmCtorArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JvmObject) GetJvmCtorArgsOk() ([]JvmObject, bool) {
	if o == nil || IsNil(o.JvmCtorArgs) {
		return nil, false
	}
	return o.JvmCtorArgs, true
}

// HasJvmCtorArgs returns a boolean if a field has been set.
func (o *JvmObject) HasJvmCtorArgs() bool {
	if o != nil && !IsNil(o.JvmCtorArgs) {
		return true
	}

	return false
}

// SetJvmCtorArgs gets a reference to the given []JvmObject and assigns it to the JvmCtorArgs field.
func (o *JvmObject) SetJvmCtorArgs(v []JvmObject) {
	o.JvmCtorArgs = v
}

// GetJvmType returns the JvmType field value
func (o *JvmObject) GetJvmType() JvmType {
	if o == nil {
		var ret JvmType
		return ret
	}

	return o.JvmType
}

// GetJvmTypeOk returns a tuple with the JvmType field value
// and a boolean to check if the value has been set.
func (o *JvmObject) GetJvmTypeOk() (*JvmType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JvmType, true
}

// SetJvmType sets field value
func (o *JvmObject) SetJvmType(v JvmType) {
	o.JvmType = v
}

func (o JvmObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JvmObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["jvmTypeKind"] = o.JvmTypeKind
	if !IsNil(o.PrimitiveValue) {
		toSerialize["primitiveValue"] = o.PrimitiveValue
	}
	if !IsNil(o.JvmCtorArgs) {
		toSerialize["jvmCtorArgs"] = o.JvmCtorArgs
	}
	toSerialize["jvmType"] = o.JvmType
	return toSerialize, nil
}

type NullableJvmObject struct {
	value *JvmObject
	isSet bool
}

func (v NullableJvmObject) Get() *JvmObject {
	return v.value
}

func (v *NullableJvmObject) Set(val *JvmObject) {
	v.value = val
	v.isSet = true
}

func (v NullableJvmObject) IsSet() bool {
	return v.isSet
}

func (v *NullableJvmObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJvmObject(val *JvmObject) *NullableJvmObject {
	return &NullableJvmObject{value: val, isSet: true}
}

func (v NullableJvmObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJvmObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


