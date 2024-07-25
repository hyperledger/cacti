/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-satp-hermes

import (
	"encoding/json"
)

// checks if the CommitPreparationV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitPreparationV1Response{}

// CommitPreparationV1Response struct for CommitPreparationV1Response
type CommitPreparationV1Response struct {
	SessionID string `json:"sessionID"`
	MessageType string `json:"messageType"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	HashCommitPrep string `json:"hashCommitPrep"`
	ServerTransferNumber *string `json:"serverTransferNumber,omitempty"`
	Signature string `json:"signature"`
	SequenceNumber float32 `json:"sequenceNumber"`
}

// NewCommitPreparationV1Response instantiates a new CommitPreparationV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitPreparationV1Response(sessionID string, messageType string, clientIdentityPubkey string, serverIdentityPubkey string, hashCommitPrep string, signature string, sequenceNumber float32) *CommitPreparationV1Response {
	this := CommitPreparationV1Response{}
	this.SessionID = sessionID
	this.MessageType = messageType
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.HashCommitPrep = hashCommitPrep
	this.Signature = signature
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewCommitPreparationV1ResponseWithDefaults instantiates a new CommitPreparationV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitPreparationV1ResponseWithDefaults() *CommitPreparationV1Response {
	this := CommitPreparationV1Response{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *CommitPreparationV1Response) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *CommitPreparationV1Response) SetSessionID(v string) {
	o.SessionID = v
}

// GetMessageType returns the MessageType field value
func (o *CommitPreparationV1Response) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *CommitPreparationV1Response) SetMessageType(v string) {
	o.MessageType = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *CommitPreparationV1Response) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *CommitPreparationV1Response) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *CommitPreparationV1Response) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *CommitPreparationV1Response) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetHashCommitPrep returns the HashCommitPrep field value
func (o *CommitPreparationV1Response) GetHashCommitPrep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashCommitPrep
}

// GetHashCommitPrepOk returns a tuple with the HashCommitPrep field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetHashCommitPrepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashCommitPrep, true
}

// SetHashCommitPrep sets field value
func (o *CommitPreparationV1Response) SetHashCommitPrep(v string) {
	o.HashCommitPrep = v
}

// GetServerTransferNumber returns the ServerTransferNumber field value if set, zero value otherwise.
func (o *CommitPreparationV1Response) GetServerTransferNumber() string {
	if o == nil || IsNil(o.ServerTransferNumber) {
		var ret string
		return ret
	}
	return *o.ServerTransferNumber
}

// GetServerTransferNumberOk returns a tuple with the ServerTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetServerTransferNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ServerTransferNumber) {
		return nil, false
	}
	return o.ServerTransferNumber, true
}

// HasServerTransferNumber returns a boolean if a field has been set.
func (o *CommitPreparationV1Response) HasServerTransferNumber() bool {
	if o != nil && !IsNil(o.ServerTransferNumber) {
		return true
	}

	return false
}

// SetServerTransferNumber gets a reference to the given string and assigns it to the ServerTransferNumber field.
func (o *CommitPreparationV1Response) SetServerTransferNumber(v string) {
	o.ServerTransferNumber = &v
}

// GetSignature returns the Signature field value
func (o *CommitPreparationV1Response) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *CommitPreparationV1Response) SetSignature(v string) {
	o.Signature = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *CommitPreparationV1Response) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *CommitPreparationV1Response) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *CommitPreparationV1Response) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

func (o CommitPreparationV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitPreparationV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["messageType"] = o.MessageType
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["hashCommitPrep"] = o.HashCommitPrep
	if !IsNil(o.ServerTransferNumber) {
		toSerialize["serverTransferNumber"] = o.ServerTransferNumber
	}
	toSerialize["signature"] = o.Signature
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableCommitPreparationV1Response struct {
	value *CommitPreparationV1Response
	isSet bool
}

func (v NullableCommitPreparationV1Response) Get() *CommitPreparationV1Response {
	return v.value
}

func (v *NullableCommitPreparationV1Response) Set(val *CommitPreparationV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitPreparationV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitPreparationV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitPreparationV1Response(val *CommitPreparationV1Response) *NullableCommitPreparationV1Response {
	return &NullableCommitPreparationV1Response{value: val, isSet: true}
}

func (v NullableCommitPreparationV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitPreparationV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


