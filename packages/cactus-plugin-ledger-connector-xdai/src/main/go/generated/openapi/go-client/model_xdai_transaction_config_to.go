/*
Hyperledger Cactus Plugin - Connector Xdai

Can perform basic tasks on a Xdai ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-xdai

import (
	"encoding/json"
	"fmt"
)

// XdaiTransactionConfigTo - struct for XdaiTransactionConfigTo
type XdaiTransactionConfigTo struct {
	String *string
}

// stringAsXdaiTransactionConfigTo is a convenience function that returns string wrapped in XdaiTransactionConfigTo
func StringAsXdaiTransactionConfigTo(v *string) XdaiTransactionConfigTo {
	return XdaiTransactionConfigTo{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *XdaiTransactionConfigTo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(XdaiTransactionConfigTo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(XdaiTransactionConfigTo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src XdaiTransactionConfigTo) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *XdaiTransactionConfigTo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableXdaiTransactionConfigTo struct {
	value *XdaiTransactionConfigTo
	isSet bool
}

func (v NullableXdaiTransactionConfigTo) Get() *XdaiTransactionConfigTo {
	return v.value
}

func (v *NullableXdaiTransactionConfigTo) Set(val *XdaiTransactionConfigTo) {
	v.value = val
	v.isSet = true
}

func (v NullableXdaiTransactionConfigTo) IsSet() bool {
	return v.isSet
}

func (v *NullableXdaiTransactionConfigTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXdaiTransactionConfigTo(val *XdaiTransactionConfigTo) *NullableXdaiTransactionConfigTo {
	return &NullableXdaiTransactionConfigTo{value: val, isSet: true}
}

func (v NullableXdaiTransactionConfigTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXdaiTransactionConfigTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


