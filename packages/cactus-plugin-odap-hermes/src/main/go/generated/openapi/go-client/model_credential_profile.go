/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-odap-hermes

import (
	"encoding/json"
	"fmt"
)

// CredentialProfile the model 'CredentialProfile'
type CredentialProfile string

// List of CredentialProfile
const (
	SAML CredentialProfile = "SAML"
	OAUTH CredentialProfile = "OAUTH"
	X509 CredentialProfile = "X509"
)

// All allowed values of CredentialProfile enum
var AllowedCredentialProfileEnumValues = []CredentialProfile{
	"SAML",
	"OAUTH",
	"X509",
}

func (v *CredentialProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CredentialProfile(value)
	for _, existing := range AllowedCredentialProfileEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CredentialProfile", value)
}

// NewCredentialProfileFromValue returns a pointer to a valid CredentialProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCredentialProfileFromValue(v string) (*CredentialProfile, error) {
	ev := CredentialProfile(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CredentialProfile: valid values are %v", v, AllowedCredentialProfileEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CredentialProfile) IsValid() bool {
	for _, existing := range AllowedCredentialProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CredentialProfile value
func (v CredentialProfile) Ptr() *CredentialProfile {
	return &v
}

type NullableCredentialProfile struct {
	value *CredentialProfile
	isSet bool
}

func (v NullableCredentialProfile) Get() *CredentialProfile {
	return v.value
}

func (v *NullableCredentialProfile) Set(val *CredentialProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialProfile(val *CredentialProfile) *NullableCredentialProfile {
	return &NullableCredentialProfile{value: val, isSet: true}
}

func (v NullableCredentialProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

