/*
Hyperledger Cacti Plugin - Connector Ethereum

Can perform basic tasks on a Ethereum ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-ethereum

import (
	"encoding/json"
)

// checks if the Web3SigningCredentialCactiKeychainRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3SigningCredentialCactiKeychainRef{}

// Web3SigningCredentialCactiKeychainRef struct for Web3SigningCredentialCactiKeychainRef
type Web3SigningCredentialCactiKeychainRef struct {
	Type Web3SigningCredentialType `json:"type"`
	// The ethereum account (public key) that the credential  belongs to. Basically the username in the traditional  terminology of authentication.
	EthAccount string `json:"ethAccount"`
	// The key to use when looking up the the keychain entry holding the secret pointed to by the  keychainEntryKey parameter.
	KeychainEntryKey string `json:"keychainEntryKey"`
	// The keychain ID to use when looking up the the keychain plugin instance that will be used to retrieve the secret pointed to by the keychainEntryKey parameter.
	KeychainId *string `json:"keychainId,omitempty"`
}

// NewWeb3SigningCredentialCactiKeychainRef instantiates a new Web3SigningCredentialCactiKeychainRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3SigningCredentialCactiKeychainRef(type_ Web3SigningCredentialType, ethAccount string, keychainEntryKey string) *Web3SigningCredentialCactiKeychainRef {
	this := Web3SigningCredentialCactiKeychainRef{}
	this.Type = type_
	this.EthAccount = ethAccount
	this.KeychainEntryKey = keychainEntryKey
	return &this
}

// NewWeb3SigningCredentialCactiKeychainRefWithDefaults instantiates a new Web3SigningCredentialCactiKeychainRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3SigningCredentialCactiKeychainRefWithDefaults() *Web3SigningCredentialCactiKeychainRef {
	this := Web3SigningCredentialCactiKeychainRef{}
	return &this
}

// GetType returns the Type field value
func (o *Web3SigningCredentialCactiKeychainRef) GetType() Web3SigningCredentialType {
	if o == nil {
		var ret Web3SigningCredentialType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Web3SigningCredentialCactiKeychainRef) GetTypeOk() (*Web3SigningCredentialType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Web3SigningCredentialCactiKeychainRef) SetType(v Web3SigningCredentialType) {
	o.Type = v
}

// GetEthAccount returns the EthAccount field value
func (o *Web3SigningCredentialCactiKeychainRef) GetEthAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EthAccount
}

// GetEthAccountOk returns a tuple with the EthAccount field value
// and a boolean to check if the value has been set.
func (o *Web3SigningCredentialCactiKeychainRef) GetEthAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EthAccount, true
}

// SetEthAccount sets field value
func (o *Web3SigningCredentialCactiKeychainRef) SetEthAccount(v string) {
	o.EthAccount = v
}

// GetKeychainEntryKey returns the KeychainEntryKey field value
func (o *Web3SigningCredentialCactiKeychainRef) GetKeychainEntryKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeychainEntryKey
}

// GetKeychainEntryKeyOk returns a tuple with the KeychainEntryKey field value
// and a boolean to check if the value has been set.
func (o *Web3SigningCredentialCactiKeychainRef) GetKeychainEntryKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeychainEntryKey, true
}

// SetKeychainEntryKey sets field value
func (o *Web3SigningCredentialCactiKeychainRef) SetKeychainEntryKey(v string) {
	o.KeychainEntryKey = v
}

// GetKeychainId returns the KeychainId field value if set, zero value otherwise.
func (o *Web3SigningCredentialCactiKeychainRef) GetKeychainId() string {
	if o == nil || isNil(o.KeychainId) {
		var ret string
		return ret
	}
	return *o.KeychainId
}

// GetKeychainIdOk returns a tuple with the KeychainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Web3SigningCredentialCactiKeychainRef) GetKeychainIdOk() (*string, bool) {
	if o == nil || isNil(o.KeychainId) {
		return nil, false
	}
	return o.KeychainId, true
}

// HasKeychainId returns a boolean if a field has been set.
func (o *Web3SigningCredentialCactiKeychainRef) HasKeychainId() bool {
	if o != nil && !isNil(o.KeychainId) {
		return true
	}

	return false
}

// SetKeychainId gets a reference to the given string and assigns it to the KeychainId field.
func (o *Web3SigningCredentialCactiKeychainRef) SetKeychainId(v string) {
	o.KeychainId = &v
}

func (o Web3SigningCredentialCactiKeychainRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3SigningCredentialCactiKeychainRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["ethAccount"] = o.EthAccount
	toSerialize["keychainEntryKey"] = o.KeychainEntryKey
	if !isNil(o.KeychainId) {
		toSerialize["keychainId"] = o.KeychainId
	}
	return toSerialize, nil
}

type NullableWeb3SigningCredentialCactiKeychainRef struct {
	value *Web3SigningCredentialCactiKeychainRef
	isSet bool
}

func (v NullableWeb3SigningCredentialCactiKeychainRef) Get() *Web3SigningCredentialCactiKeychainRef {
	return v.value
}

func (v *NullableWeb3SigningCredentialCactiKeychainRef) Set(val *Web3SigningCredentialCactiKeychainRef) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3SigningCredentialCactiKeychainRef) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3SigningCredentialCactiKeychainRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3SigningCredentialCactiKeychainRef(val *Web3SigningCredentialCactiKeychainRef) *NullableWeb3SigningCredentialCactiKeychainRef {
	return &NullableWeb3SigningCredentialCactiKeychainRef{value: val, isSet: true}
}

func (v NullableWeb3SigningCredentialCactiKeychainRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3SigningCredentialCactiKeychainRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


