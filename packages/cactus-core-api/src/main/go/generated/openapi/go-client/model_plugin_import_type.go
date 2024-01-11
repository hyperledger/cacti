/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
	"fmt"
)

// PluginImportType the model 'PluginImportType'
type PluginImportType string

// List of PluginImportType
const (
	LOCAL PluginImportType = "org.hyperledger.cactus.plugin_import_type.LOCAL"
	REMOTE PluginImportType = "org.hyperledger.cactus.plugin_import_type.REMOTE"
)

// All allowed values of PluginImportType enum
var AllowedPluginImportTypeEnumValues = []PluginImportType{
	"org.hyperledger.cactus.plugin_import_type.LOCAL",
	"org.hyperledger.cactus.plugin_import_type.REMOTE",
}

func (v *PluginImportType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PluginImportType(value)
	for _, existing := range AllowedPluginImportTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PluginImportType", value)
}

// NewPluginImportTypeFromValue returns a pointer to a valid PluginImportType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPluginImportTypeFromValue(v string) (*PluginImportType, error) {
	ev := PluginImportType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PluginImportType: valid values are %v", v, AllowedPluginImportTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PluginImportType) IsValid() bool {
	for _, existing := range AllowedPluginImportTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PluginImportType value
func (v PluginImportType) Ptr() *PluginImportType {
	return &v
}

type NullablePluginImportType struct {
	value *PluginImportType
	isSet bool
}

func (v NullablePluginImportType) Get() *PluginImportType {
	return v.value
}

func (v *NullablePluginImportType) Set(val *PluginImportType) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginImportType) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginImportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginImportType(val *PluginImportType) *NullablePluginImportType {
	return &NullablePluginImportType{value: val, isSet: true}
}

func (v NullablePluginImportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginImportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

