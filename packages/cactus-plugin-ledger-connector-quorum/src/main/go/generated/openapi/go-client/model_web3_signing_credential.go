/*
Hyperledger Cactus Plugin - Connector Quorum

Can perform basic tasks on a Quorum ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-quorum

import (
	"encoding/json"
	"fmt"
)

// Web3SigningCredential - struct for Web3SigningCredential
type Web3SigningCredential struct {
	Web3SigningCredentialCactusKeychainRef *Web3SigningCredentialCactusKeychainRef
	Web3SigningCredentialGethKeychainPassword *Web3SigningCredentialGethKeychainPassword
	Web3SigningCredentialNone *Web3SigningCredentialNone
	Web3SigningCredentialPrivateKeyHex *Web3SigningCredentialPrivateKeyHex
}

// Web3SigningCredentialCactusKeychainRefAsWeb3SigningCredential is a convenience function that returns Web3SigningCredentialCactusKeychainRef wrapped in Web3SigningCredential
func Web3SigningCredentialCactusKeychainRefAsWeb3SigningCredential(v *Web3SigningCredentialCactusKeychainRef) Web3SigningCredential {
	return Web3SigningCredential{
		Web3SigningCredentialCactusKeychainRef: v,
	}
}

// Web3SigningCredentialGethKeychainPasswordAsWeb3SigningCredential is a convenience function that returns Web3SigningCredentialGethKeychainPassword wrapped in Web3SigningCredential
func Web3SigningCredentialGethKeychainPasswordAsWeb3SigningCredential(v *Web3SigningCredentialGethKeychainPassword) Web3SigningCredential {
	return Web3SigningCredential{
		Web3SigningCredentialGethKeychainPassword: v,
	}
}

// Web3SigningCredentialNoneAsWeb3SigningCredential is a convenience function that returns Web3SigningCredentialNone wrapped in Web3SigningCredential
func Web3SigningCredentialNoneAsWeb3SigningCredential(v *Web3SigningCredentialNone) Web3SigningCredential {
	return Web3SigningCredential{
		Web3SigningCredentialNone: v,
	}
}

// Web3SigningCredentialPrivateKeyHexAsWeb3SigningCredential is a convenience function that returns Web3SigningCredentialPrivateKeyHex wrapped in Web3SigningCredential
func Web3SigningCredentialPrivateKeyHexAsWeb3SigningCredential(v *Web3SigningCredentialPrivateKeyHex) Web3SigningCredential {
	return Web3SigningCredential{
		Web3SigningCredentialPrivateKeyHex: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Web3SigningCredential) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Web3SigningCredentialCactusKeychainRef
	err = newStrictDecoder(data).Decode(&dst.Web3SigningCredentialCactusKeychainRef)
	if err == nil {
		jsonWeb3SigningCredentialCactusKeychainRef, _ := json.Marshal(dst.Web3SigningCredentialCactusKeychainRef)
		if string(jsonWeb3SigningCredentialCactusKeychainRef) == "{}" { // empty struct
			dst.Web3SigningCredentialCactusKeychainRef = nil
		} else {
			match++
		}
	} else {
		dst.Web3SigningCredentialCactusKeychainRef = nil
	}

	// try to unmarshal data into Web3SigningCredentialGethKeychainPassword
	err = newStrictDecoder(data).Decode(&dst.Web3SigningCredentialGethKeychainPassword)
	if err == nil {
		jsonWeb3SigningCredentialGethKeychainPassword, _ := json.Marshal(dst.Web3SigningCredentialGethKeychainPassword)
		if string(jsonWeb3SigningCredentialGethKeychainPassword) == "{}" { // empty struct
			dst.Web3SigningCredentialGethKeychainPassword = nil
		} else {
			match++
		}
	} else {
		dst.Web3SigningCredentialGethKeychainPassword = nil
	}

	// try to unmarshal data into Web3SigningCredentialNone
	err = newStrictDecoder(data).Decode(&dst.Web3SigningCredentialNone)
	if err == nil {
		jsonWeb3SigningCredentialNone, _ := json.Marshal(dst.Web3SigningCredentialNone)
		if string(jsonWeb3SigningCredentialNone) == "{}" { // empty struct
			dst.Web3SigningCredentialNone = nil
		} else {
			match++
		}
	} else {
		dst.Web3SigningCredentialNone = nil
	}

	// try to unmarshal data into Web3SigningCredentialPrivateKeyHex
	err = newStrictDecoder(data).Decode(&dst.Web3SigningCredentialPrivateKeyHex)
	if err == nil {
		jsonWeb3SigningCredentialPrivateKeyHex, _ := json.Marshal(dst.Web3SigningCredentialPrivateKeyHex)
		if string(jsonWeb3SigningCredentialPrivateKeyHex) == "{}" { // empty struct
			dst.Web3SigningCredentialPrivateKeyHex = nil
		} else {
			match++
		}
	} else {
		dst.Web3SigningCredentialPrivateKeyHex = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Web3SigningCredentialCactusKeychainRef = nil
		dst.Web3SigningCredentialGethKeychainPassword = nil
		dst.Web3SigningCredentialNone = nil
		dst.Web3SigningCredentialPrivateKeyHex = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Web3SigningCredential)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Web3SigningCredential)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Web3SigningCredential) MarshalJSON() ([]byte, error) {
	if src.Web3SigningCredentialCactusKeychainRef != nil {
		return json.Marshal(&src.Web3SigningCredentialCactusKeychainRef)
	}

	if src.Web3SigningCredentialGethKeychainPassword != nil {
		return json.Marshal(&src.Web3SigningCredentialGethKeychainPassword)
	}

	if src.Web3SigningCredentialNone != nil {
		return json.Marshal(&src.Web3SigningCredentialNone)
	}

	if src.Web3SigningCredentialPrivateKeyHex != nil {
		return json.Marshal(&src.Web3SigningCredentialPrivateKeyHex)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Web3SigningCredential) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Web3SigningCredentialCactusKeychainRef != nil {
		return obj.Web3SigningCredentialCactusKeychainRef
	}

	if obj.Web3SigningCredentialGethKeychainPassword != nil {
		return obj.Web3SigningCredentialGethKeychainPassword
	}

	if obj.Web3SigningCredentialNone != nil {
		return obj.Web3SigningCredentialNone
	}

	if obj.Web3SigningCredentialPrivateKeyHex != nil {
		return obj.Web3SigningCredentialPrivateKeyHex
	}

	// all schemas are nil
	return nil
}

type NullableWeb3SigningCredential struct {
	value *Web3SigningCredential
	isSet bool
}

func (v NullableWeb3SigningCredential) Get() *Web3SigningCredential {
	return v.value
}

func (v *NullableWeb3SigningCredential) Set(val *Web3SigningCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3SigningCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3SigningCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3SigningCredential(val *Web3SigningCredential) *NullableWeb3SigningCredential {
	return &NullableWeb3SigningCredential{value: val, isSet: true}
}

func (v NullableWeb3SigningCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3SigningCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


