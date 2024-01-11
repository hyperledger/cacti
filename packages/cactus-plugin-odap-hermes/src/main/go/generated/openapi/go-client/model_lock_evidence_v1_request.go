/*
Hyperledger Cactus Plugin - Odap Hermes

Implementation for Odap and Hermes

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-odap-hermes

import (
	"encoding/json"
)

// checks if the LockEvidenceV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockEvidenceV1Request{}

// LockEvidenceV1Request struct for LockEvidenceV1Request
type LockEvidenceV1Request struct {
	SessionID string `json:"sessionID"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	LockEvidenceClaim string `json:"lockEvidenceClaim"`
	LockEvidenceFormat map[string]interface{} `json:"lockEvidenceFormat,omitempty"`
	LockEvidenceExpiration string `json:"lockEvidenceExpiration"`
	HashCommenceAckRequest string `json:"hashCommenceAckRequest"`
	ClientTransferNumber NullableInt32 `json:"clientTransferNumber,omitempty"`
	Signature string `json:"signature"`
	MessageType string `json:"messageType"`
	MessageHash *string `json:"messageHash,omitempty"`
	SequenceNumber float32 `json:"sequenceNumber"`
}

// NewLockEvidenceV1Request instantiates a new LockEvidenceV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockEvidenceV1Request(sessionID string, clientIdentityPubkey string, serverIdentityPubkey string, lockEvidenceClaim string, lockEvidenceExpiration string, hashCommenceAckRequest string, signature string, messageType string, sequenceNumber float32) *LockEvidenceV1Request {
	this := LockEvidenceV1Request{}
	this.SessionID = sessionID
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.LockEvidenceClaim = lockEvidenceClaim
	this.LockEvidenceExpiration = lockEvidenceExpiration
	this.HashCommenceAckRequest = hashCommenceAckRequest
	this.Signature = signature
	this.MessageType = messageType
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewLockEvidenceV1RequestWithDefaults instantiates a new LockEvidenceV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockEvidenceV1RequestWithDefaults() *LockEvidenceV1Request {
	this := LockEvidenceV1Request{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *LockEvidenceV1Request) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *LockEvidenceV1Request) SetSessionID(v string) {
	o.SessionID = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *LockEvidenceV1Request) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *LockEvidenceV1Request) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *LockEvidenceV1Request) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *LockEvidenceV1Request) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetLockEvidenceClaim returns the LockEvidenceClaim field value
func (o *LockEvidenceV1Request) GetLockEvidenceClaim() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockEvidenceClaim
}

// GetLockEvidenceClaimOk returns a tuple with the LockEvidenceClaim field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetLockEvidenceClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockEvidenceClaim, true
}

// SetLockEvidenceClaim sets field value
func (o *LockEvidenceV1Request) SetLockEvidenceClaim(v string) {
	o.LockEvidenceClaim = v
}

// GetLockEvidenceFormat returns the LockEvidenceFormat field value if set, zero value otherwise.
func (o *LockEvidenceV1Request) GetLockEvidenceFormat() map[string]interface{} {
	if o == nil || IsNil(o.LockEvidenceFormat) {
		var ret map[string]interface{}
		return ret
	}
	return o.LockEvidenceFormat
}

// GetLockEvidenceFormatOk returns a tuple with the LockEvidenceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetLockEvidenceFormatOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LockEvidenceFormat) {
		return map[string]interface{}{}, false
	}
	return o.LockEvidenceFormat, true
}

// HasLockEvidenceFormat returns a boolean if a field has been set.
func (o *LockEvidenceV1Request) HasLockEvidenceFormat() bool {
	if o != nil && !IsNil(o.LockEvidenceFormat) {
		return true
	}

	return false
}

// SetLockEvidenceFormat gets a reference to the given map[string]interface{} and assigns it to the LockEvidenceFormat field.
func (o *LockEvidenceV1Request) SetLockEvidenceFormat(v map[string]interface{}) {
	o.LockEvidenceFormat = v
}

// GetLockEvidenceExpiration returns the LockEvidenceExpiration field value
func (o *LockEvidenceV1Request) GetLockEvidenceExpiration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LockEvidenceExpiration
}

// GetLockEvidenceExpirationOk returns a tuple with the LockEvidenceExpiration field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetLockEvidenceExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockEvidenceExpiration, true
}

// SetLockEvidenceExpiration sets field value
func (o *LockEvidenceV1Request) SetLockEvidenceExpiration(v string) {
	o.LockEvidenceExpiration = v
}

// GetHashCommenceAckRequest returns the HashCommenceAckRequest field value
func (o *LockEvidenceV1Request) GetHashCommenceAckRequest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashCommenceAckRequest
}

// GetHashCommenceAckRequestOk returns a tuple with the HashCommenceAckRequest field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetHashCommenceAckRequestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashCommenceAckRequest, true
}

// SetHashCommenceAckRequest sets field value
func (o *LockEvidenceV1Request) SetHashCommenceAckRequest(v string) {
	o.HashCommenceAckRequest = v
}

// GetClientTransferNumber returns the ClientTransferNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LockEvidenceV1Request) GetClientTransferNumber() int32 {
	if o == nil || IsNil(o.ClientTransferNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ClientTransferNumber.Get()
}

// GetClientTransferNumberOk returns a tuple with the ClientTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LockEvidenceV1Request) GetClientTransferNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientTransferNumber.Get(), o.ClientTransferNumber.IsSet()
}

// HasClientTransferNumber returns a boolean if a field has been set.
func (o *LockEvidenceV1Request) HasClientTransferNumber() bool {
	if o != nil && o.ClientTransferNumber.IsSet() {
		return true
	}

	return false
}

// SetClientTransferNumber gets a reference to the given NullableInt32 and assigns it to the ClientTransferNumber field.
func (o *LockEvidenceV1Request) SetClientTransferNumber(v int32) {
	o.ClientTransferNumber.Set(&v)
}
// SetClientTransferNumberNil sets the value for ClientTransferNumber to be an explicit nil
func (o *LockEvidenceV1Request) SetClientTransferNumberNil() {
	o.ClientTransferNumber.Set(nil)
}

// UnsetClientTransferNumber ensures that no value is present for ClientTransferNumber, not even an explicit nil
func (o *LockEvidenceV1Request) UnsetClientTransferNumber() {
	o.ClientTransferNumber.Unset()
}

// GetSignature returns the Signature field value
func (o *LockEvidenceV1Request) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *LockEvidenceV1Request) SetSignature(v string) {
	o.Signature = v
}

// GetMessageType returns the MessageType field value
func (o *LockEvidenceV1Request) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *LockEvidenceV1Request) SetMessageType(v string) {
	o.MessageType = v
}

// GetMessageHash returns the MessageHash field value if set, zero value otherwise.
func (o *LockEvidenceV1Request) GetMessageHash() string {
	if o == nil || IsNil(o.MessageHash) {
		var ret string
		return ret
	}
	return *o.MessageHash
}

// GetMessageHashOk returns a tuple with the MessageHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetMessageHashOk() (*string, bool) {
	if o == nil || IsNil(o.MessageHash) {
		return nil, false
	}
	return o.MessageHash, true
}

// HasMessageHash returns a boolean if a field has been set.
func (o *LockEvidenceV1Request) HasMessageHash() bool {
	if o != nil && !IsNil(o.MessageHash) {
		return true
	}

	return false
}

// SetMessageHash gets a reference to the given string and assigns it to the MessageHash field.
func (o *LockEvidenceV1Request) SetMessageHash(v string) {
	o.MessageHash = &v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *LockEvidenceV1Request) GetSequenceNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *LockEvidenceV1Request) GetSequenceNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *LockEvidenceV1Request) SetSequenceNumber(v float32) {
	o.SequenceNumber = v
}

func (o LockEvidenceV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockEvidenceV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["lockEvidenceClaim"] = o.LockEvidenceClaim
	if !IsNil(o.LockEvidenceFormat) {
		toSerialize["lockEvidenceFormat"] = o.LockEvidenceFormat
	}
	toSerialize["lockEvidenceExpiration"] = o.LockEvidenceExpiration
	toSerialize["hashCommenceAckRequest"] = o.HashCommenceAckRequest
	if o.ClientTransferNumber.IsSet() {
		toSerialize["clientTransferNumber"] = o.ClientTransferNumber.Get()
	}
	toSerialize["signature"] = o.Signature
	toSerialize["messageType"] = o.MessageType
	if !IsNil(o.MessageHash) {
		toSerialize["messageHash"] = o.MessageHash
	}
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableLockEvidenceV1Request struct {
	value *LockEvidenceV1Request
	isSet bool
}

func (v NullableLockEvidenceV1Request) Get() *LockEvidenceV1Request {
	return v.value
}

func (v *NullableLockEvidenceV1Request) Set(val *LockEvidenceV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableLockEvidenceV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableLockEvidenceV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockEvidenceV1Request(val *LockEvidenceV1Request) *NullableLockEvidenceV1Request {
	return &NullableLockEvidenceV1Request{value: val, isSet: true}
}

func (v NullableLockEvidenceV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockEvidenceV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


