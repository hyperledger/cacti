/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the TransactReceiptTransactionEndorsement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactReceiptTransactionEndorsement{}

// TransactReceiptTransactionEndorsement struct for TransactReceiptTransactionEndorsement
type TransactReceiptTransactionEndorsement struct {
	Mspid *string `json:"mspid,omitempty"`
	EndorserID *string `json:"endorserID,omitempty"`
	Signature *string `json:"signature,omitempty"`
}

// NewTransactReceiptTransactionEndorsement instantiates a new TransactReceiptTransactionEndorsement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactReceiptTransactionEndorsement() *TransactReceiptTransactionEndorsement {
	this := TransactReceiptTransactionEndorsement{}
	return &this
}

// NewTransactReceiptTransactionEndorsementWithDefaults instantiates a new TransactReceiptTransactionEndorsement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactReceiptTransactionEndorsementWithDefaults() *TransactReceiptTransactionEndorsement {
	this := TransactReceiptTransactionEndorsement{}
	return &this
}

// GetMspid returns the Mspid field value if set, zero value otherwise.
func (o *TransactReceiptTransactionEndorsement) GetMspid() string {
	if o == nil || IsNil(o.Mspid) {
		var ret string
		return ret
	}
	return *o.Mspid
}

// GetMspidOk returns a tuple with the Mspid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptTransactionEndorsement) GetMspidOk() (*string, bool) {
	if o == nil || IsNil(o.Mspid) {
		return nil, false
	}
	return o.Mspid, true
}

// HasMspid returns a boolean if a field has been set.
func (o *TransactReceiptTransactionEndorsement) HasMspid() bool {
	if o != nil && !IsNil(o.Mspid) {
		return true
	}

	return false
}

// SetMspid gets a reference to the given string and assigns it to the Mspid field.
func (o *TransactReceiptTransactionEndorsement) SetMspid(v string) {
	o.Mspid = &v
}

// GetEndorserID returns the EndorserID field value if set, zero value otherwise.
func (o *TransactReceiptTransactionEndorsement) GetEndorserID() string {
	if o == nil || IsNil(o.EndorserID) {
		var ret string
		return ret
	}
	return *o.EndorserID
}

// GetEndorserIDOk returns a tuple with the EndorserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptTransactionEndorsement) GetEndorserIDOk() (*string, bool) {
	if o == nil || IsNil(o.EndorserID) {
		return nil, false
	}
	return o.EndorserID, true
}

// HasEndorserID returns a boolean if a field has been set.
func (o *TransactReceiptTransactionEndorsement) HasEndorserID() bool {
	if o != nil && !IsNil(o.EndorserID) {
		return true
	}

	return false
}

// SetEndorserID gets a reference to the given string and assigns it to the EndorserID field.
func (o *TransactReceiptTransactionEndorsement) SetEndorserID(v string) {
	o.EndorserID = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *TransactReceiptTransactionEndorsement) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactReceiptTransactionEndorsement) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *TransactReceiptTransactionEndorsement) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *TransactReceiptTransactionEndorsement) SetSignature(v string) {
	o.Signature = &v
}

func (o TransactReceiptTransactionEndorsement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactReceiptTransactionEndorsement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mspid) {
		toSerialize["mspid"] = o.Mspid
	}
	if !IsNil(o.EndorserID) {
		toSerialize["endorserID"] = o.EndorserID
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	return toSerialize, nil
}

type NullableTransactReceiptTransactionEndorsement struct {
	value *TransactReceiptTransactionEndorsement
	isSet bool
}

func (v NullableTransactReceiptTransactionEndorsement) Get() *TransactReceiptTransactionEndorsement {
	return v.value
}

func (v *NullableTransactReceiptTransactionEndorsement) Set(val *TransactReceiptTransactionEndorsement) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactReceiptTransactionEndorsement) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactReceiptTransactionEndorsement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactReceiptTransactionEndorsement(val *TransactReceiptTransactionEndorsement) *NullableTransactReceiptTransactionEndorsement {
	return &NullableTransactReceiptTransactionEndorsement{value: val, isSet: true}
}

func (v NullableTransactReceiptTransactionEndorsement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactReceiptTransactionEndorsement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


