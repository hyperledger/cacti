/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
	"fmt"
)

// SearchLineageTypeV1 the model 'SearchLineageTypeV1'
type SearchLineageTypeV1 string

// List of SearchLineageTypeV1
const (
	ExactMatch SearchLineageTypeV1 = "exactmatch"
	PartialMatch SearchLineageTypeV1 = "partialmatch"
	RegexMatch SearchLineageTypeV1 = "regexpmatch"
)

// All allowed values of SearchLineageTypeV1 enum
var AllowedSearchLineageTypeV1EnumValues = []SearchLineageTypeV1{
	"exactmatch",
	"partialmatch",
	"regexpmatch",
}

func (v *SearchLineageTypeV1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SearchLineageTypeV1(value)
	for _, existing := range AllowedSearchLineageTypeV1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SearchLineageTypeV1", value)
}

// NewSearchLineageTypeV1FromValue returns a pointer to a valid SearchLineageTypeV1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSearchLineageTypeV1FromValue(v string) (*SearchLineageTypeV1, error) {
	ev := SearchLineageTypeV1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SearchLineageTypeV1: valid values are %v", v, AllowedSearchLineageTypeV1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SearchLineageTypeV1) IsValid() bool {
	for _, existing := range AllowedSearchLineageTypeV1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SearchLineageTypeV1 value
func (v SearchLineageTypeV1) Ptr() *SearchLineageTypeV1 {
	return &v
}

type NullableSearchLineageTypeV1 struct {
	value *SearchLineageTypeV1
	isSet bool
}

func (v NullableSearchLineageTypeV1) Get() *SearchLineageTypeV1 {
	return v.value
}

func (v *NullableSearchLineageTypeV1) Set(val *SearchLineageTypeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchLineageTypeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchLineageTypeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchLineageTypeV1(val *SearchLineageTypeV1) *NullableSearchLineageTypeV1 {
	return &NullableSearchLineageTypeV1{value: val, isSet: true}
}

func (v NullableSearchLineageTypeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchLineageTypeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

