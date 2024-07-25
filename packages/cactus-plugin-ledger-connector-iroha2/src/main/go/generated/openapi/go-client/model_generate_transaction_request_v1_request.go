/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
	"fmt"
)

// GenerateTransactionRequestV1Request - struct for GenerateTransactionRequestV1Request
type GenerateTransactionRequestV1Request struct {
	IrohaQueryDefinitionV1 *IrohaQueryDefinitionV1
	IrohaTransactionDefinitionV1 *IrohaTransactionDefinitionV1
}

// IrohaQueryDefinitionV1AsGenerateTransactionRequestV1Request is a convenience function that returns IrohaQueryDefinitionV1 wrapped in GenerateTransactionRequestV1Request
func IrohaQueryDefinitionV1AsGenerateTransactionRequestV1Request(v *IrohaQueryDefinitionV1) GenerateTransactionRequestV1Request {
	return GenerateTransactionRequestV1Request{
		IrohaQueryDefinitionV1: v,
	}
}

// IrohaTransactionDefinitionV1AsGenerateTransactionRequestV1Request is a convenience function that returns IrohaTransactionDefinitionV1 wrapped in GenerateTransactionRequestV1Request
func IrohaTransactionDefinitionV1AsGenerateTransactionRequestV1Request(v *IrohaTransactionDefinitionV1) GenerateTransactionRequestV1Request {
	return GenerateTransactionRequestV1Request{
		IrohaTransactionDefinitionV1: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GenerateTransactionRequestV1Request) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into IrohaQueryDefinitionV1
	err = newStrictDecoder(data).Decode(&dst.IrohaQueryDefinitionV1)
	if err == nil {
		jsonIrohaQueryDefinitionV1, _ := json.Marshal(dst.IrohaQueryDefinitionV1)
		if string(jsonIrohaQueryDefinitionV1) == "{}" { // empty struct
			dst.IrohaQueryDefinitionV1 = nil
		} else {
			match++
		}
	} else {
		dst.IrohaQueryDefinitionV1 = nil
	}

	// try to unmarshal data into IrohaTransactionDefinitionV1
	err = newStrictDecoder(data).Decode(&dst.IrohaTransactionDefinitionV1)
	if err == nil {
		jsonIrohaTransactionDefinitionV1, _ := json.Marshal(dst.IrohaTransactionDefinitionV1)
		if string(jsonIrohaTransactionDefinitionV1) == "{}" { // empty struct
			dst.IrohaTransactionDefinitionV1 = nil
		} else {
			match++
		}
	} else {
		dst.IrohaTransactionDefinitionV1 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.IrohaQueryDefinitionV1 = nil
		dst.IrohaTransactionDefinitionV1 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GenerateTransactionRequestV1Request)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GenerateTransactionRequestV1Request)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GenerateTransactionRequestV1Request) MarshalJSON() ([]byte, error) {
	if src.IrohaQueryDefinitionV1 != nil {
		return json.Marshal(&src.IrohaQueryDefinitionV1)
	}

	if src.IrohaTransactionDefinitionV1 != nil {
		return json.Marshal(&src.IrohaTransactionDefinitionV1)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GenerateTransactionRequestV1Request) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.IrohaQueryDefinitionV1 != nil {
		return obj.IrohaQueryDefinitionV1
	}

	if obj.IrohaTransactionDefinitionV1 != nil {
		return obj.IrohaTransactionDefinitionV1
	}

	// all schemas are nil
	return nil
}

type NullableGenerateTransactionRequestV1Request struct {
	value *GenerateTransactionRequestV1Request
	isSet bool
}

func (v NullableGenerateTransactionRequestV1Request) Get() *GenerateTransactionRequestV1Request {
	return v.value
}

func (v *NullableGenerateTransactionRequestV1Request) Set(val *GenerateTransactionRequestV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTransactionRequestV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTransactionRequestV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTransactionRequestV1Request(val *GenerateTransactionRequestV1Request) *NullableGenerateTransactionRequestV1Request {
	return &NullableGenerateTransactionRequestV1Request{value: val, isSet: true}
}

func (v NullableGenerateTransactionRequestV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateTransactionRequestV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


