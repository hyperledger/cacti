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

// checks if the TransferCommenceV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferCommenceV1Request{}

// TransferCommenceV1Request struct for TransferCommenceV1Request
type TransferCommenceV1Request struct {
	SessionID string `json:"sessionID"`
	MessageType string `json:"messageType"`
	OriginatorPubkey string `json:"originatorPubkey"`
	BeneficiaryPubkey string `json:"beneficiaryPubkey"`
	SenderDltSystem string `json:"senderDltSystem"`
	RecipientDltSystem string `json:"recipientDltSystem"`
	ClientIdentityPubkey string `json:"clientIdentityPubkey"`
	ServerIdentityPubkey string `json:"serverIdentityPubkey"`
	HashAssetProfile string `json:"hashAssetProfile"`
	AssetUnit *int32 `json:"assetUnit,omitempty"`
	HashPrevMessage string `json:"hashPrevMessage"`
	ClientTransferNumber NullableInt32 `json:"clientTransferNumber,omitempty"`
	Signature string `json:"signature"`
	SequenceNumber int32 `json:"sequenceNumber"`
}

// NewTransferCommenceV1Request instantiates a new TransferCommenceV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferCommenceV1Request(sessionID string, messageType string, originatorPubkey string, beneficiaryPubkey string, senderDltSystem string, recipientDltSystem string, clientIdentityPubkey string, serverIdentityPubkey string, hashAssetProfile string, hashPrevMessage string, signature string, sequenceNumber int32) *TransferCommenceV1Request {
	this := TransferCommenceV1Request{}
	this.SessionID = sessionID
	this.MessageType = messageType
	this.OriginatorPubkey = originatorPubkey
	this.BeneficiaryPubkey = beneficiaryPubkey
	this.SenderDltSystem = senderDltSystem
	this.RecipientDltSystem = recipientDltSystem
	this.ClientIdentityPubkey = clientIdentityPubkey
	this.ServerIdentityPubkey = serverIdentityPubkey
	this.HashAssetProfile = hashAssetProfile
	this.HashPrevMessage = hashPrevMessage
	this.Signature = signature
	this.SequenceNumber = sequenceNumber
	return &this
}

// NewTransferCommenceV1RequestWithDefaults instantiates a new TransferCommenceV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferCommenceV1RequestWithDefaults() *TransferCommenceV1Request {
	this := TransferCommenceV1Request{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *TransferCommenceV1Request) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *TransferCommenceV1Request) SetSessionID(v string) {
	o.SessionID = v
}

// GetMessageType returns the MessageType field value
func (o *TransferCommenceV1Request) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *TransferCommenceV1Request) SetMessageType(v string) {
	o.MessageType = v
}

// GetOriginatorPubkey returns the OriginatorPubkey field value
func (o *TransferCommenceV1Request) GetOriginatorPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginatorPubkey
}

// GetOriginatorPubkeyOk returns a tuple with the OriginatorPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetOriginatorPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginatorPubkey, true
}

// SetOriginatorPubkey sets field value
func (o *TransferCommenceV1Request) SetOriginatorPubkey(v string) {
	o.OriginatorPubkey = v
}

// GetBeneficiaryPubkey returns the BeneficiaryPubkey field value
func (o *TransferCommenceV1Request) GetBeneficiaryPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeneficiaryPubkey
}

// GetBeneficiaryPubkeyOk returns a tuple with the BeneficiaryPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetBeneficiaryPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeneficiaryPubkey, true
}

// SetBeneficiaryPubkey sets field value
func (o *TransferCommenceV1Request) SetBeneficiaryPubkey(v string) {
	o.BeneficiaryPubkey = v
}

// GetSenderDltSystem returns the SenderDltSystem field value
func (o *TransferCommenceV1Request) GetSenderDltSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderDltSystem
}

// GetSenderDltSystemOk returns a tuple with the SenderDltSystem field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetSenderDltSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderDltSystem, true
}

// SetSenderDltSystem sets field value
func (o *TransferCommenceV1Request) SetSenderDltSystem(v string) {
	o.SenderDltSystem = v
}

// GetRecipientDltSystem returns the RecipientDltSystem field value
func (o *TransferCommenceV1Request) GetRecipientDltSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientDltSystem
}

// GetRecipientDltSystemOk returns a tuple with the RecipientDltSystem field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetRecipientDltSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientDltSystem, true
}

// SetRecipientDltSystem sets field value
func (o *TransferCommenceV1Request) SetRecipientDltSystem(v string) {
	o.RecipientDltSystem = v
}

// GetClientIdentityPubkey returns the ClientIdentityPubkey field value
func (o *TransferCommenceV1Request) GetClientIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIdentityPubkey
}

// GetClientIdentityPubkeyOk returns a tuple with the ClientIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetClientIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIdentityPubkey, true
}

// SetClientIdentityPubkey sets field value
func (o *TransferCommenceV1Request) SetClientIdentityPubkey(v string) {
	o.ClientIdentityPubkey = v
}

// GetServerIdentityPubkey returns the ServerIdentityPubkey field value
func (o *TransferCommenceV1Request) GetServerIdentityPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerIdentityPubkey
}

// GetServerIdentityPubkeyOk returns a tuple with the ServerIdentityPubkey field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetServerIdentityPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerIdentityPubkey, true
}

// SetServerIdentityPubkey sets field value
func (o *TransferCommenceV1Request) SetServerIdentityPubkey(v string) {
	o.ServerIdentityPubkey = v
}

// GetHashAssetProfile returns the HashAssetProfile field value
func (o *TransferCommenceV1Request) GetHashAssetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashAssetProfile
}

// GetHashAssetProfileOk returns a tuple with the HashAssetProfile field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetHashAssetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashAssetProfile, true
}

// SetHashAssetProfile sets field value
func (o *TransferCommenceV1Request) SetHashAssetProfile(v string) {
	o.HashAssetProfile = v
}

// GetAssetUnit returns the AssetUnit field value if set, zero value otherwise.
func (o *TransferCommenceV1Request) GetAssetUnit() int32 {
	if o == nil || IsNil(o.AssetUnit) {
		var ret int32
		return ret
	}
	return *o.AssetUnit
}

// GetAssetUnitOk returns a tuple with the AssetUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetAssetUnitOk() (*int32, bool) {
	if o == nil || IsNil(o.AssetUnit) {
		return nil, false
	}
	return o.AssetUnit, true
}

// HasAssetUnit returns a boolean if a field has been set.
func (o *TransferCommenceV1Request) HasAssetUnit() bool {
	if o != nil && !IsNil(o.AssetUnit) {
		return true
	}

	return false
}

// SetAssetUnit gets a reference to the given int32 and assigns it to the AssetUnit field.
func (o *TransferCommenceV1Request) SetAssetUnit(v int32) {
	o.AssetUnit = &v
}

// GetHashPrevMessage returns the HashPrevMessage field value
func (o *TransferCommenceV1Request) GetHashPrevMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HashPrevMessage
}

// GetHashPrevMessageOk returns a tuple with the HashPrevMessage field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetHashPrevMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashPrevMessage, true
}

// SetHashPrevMessage sets field value
func (o *TransferCommenceV1Request) SetHashPrevMessage(v string) {
	o.HashPrevMessage = v
}

// GetClientTransferNumber returns the ClientTransferNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransferCommenceV1Request) GetClientTransferNumber() int32 {
	if o == nil || IsNil(o.ClientTransferNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ClientTransferNumber.Get()
}

// GetClientTransferNumberOk returns a tuple with the ClientTransferNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferCommenceV1Request) GetClientTransferNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientTransferNumber.Get(), o.ClientTransferNumber.IsSet()
}

// HasClientTransferNumber returns a boolean if a field has been set.
func (o *TransferCommenceV1Request) HasClientTransferNumber() bool {
	if o != nil && o.ClientTransferNumber.IsSet() {
		return true
	}

	return false
}

// SetClientTransferNumber gets a reference to the given NullableInt32 and assigns it to the ClientTransferNumber field.
func (o *TransferCommenceV1Request) SetClientTransferNumber(v int32) {
	o.ClientTransferNumber.Set(&v)
}
// SetClientTransferNumberNil sets the value for ClientTransferNumber to be an explicit nil
func (o *TransferCommenceV1Request) SetClientTransferNumberNil() {
	o.ClientTransferNumber.Set(nil)
}

// UnsetClientTransferNumber ensures that no value is present for ClientTransferNumber, not even an explicit nil
func (o *TransferCommenceV1Request) UnsetClientTransferNumber() {
	o.ClientTransferNumber.Unset()
}

// GetSignature returns the Signature field value
func (o *TransferCommenceV1Request) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *TransferCommenceV1Request) SetSignature(v string) {
	o.Signature = v
}

// GetSequenceNumber returns the SequenceNumber field value
func (o *TransferCommenceV1Request) GetSequenceNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SequenceNumber
}

// GetSequenceNumberOk returns a tuple with the SequenceNumber field value
// and a boolean to check if the value has been set.
func (o *TransferCommenceV1Request) GetSequenceNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SequenceNumber, true
}

// SetSequenceNumber sets field value
func (o *TransferCommenceV1Request) SetSequenceNumber(v int32) {
	o.SequenceNumber = v
}

func (o TransferCommenceV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferCommenceV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	toSerialize["messageType"] = o.MessageType
	toSerialize["originatorPubkey"] = o.OriginatorPubkey
	toSerialize["beneficiaryPubkey"] = o.BeneficiaryPubkey
	toSerialize["senderDltSystem"] = o.SenderDltSystem
	toSerialize["recipientDltSystem"] = o.RecipientDltSystem
	toSerialize["clientIdentityPubkey"] = o.ClientIdentityPubkey
	toSerialize["serverIdentityPubkey"] = o.ServerIdentityPubkey
	toSerialize["hashAssetProfile"] = o.HashAssetProfile
	if !IsNil(o.AssetUnit) {
		toSerialize["assetUnit"] = o.AssetUnit
	}
	toSerialize["hashPrevMessage"] = o.HashPrevMessage
	if o.ClientTransferNumber.IsSet() {
		toSerialize["clientTransferNumber"] = o.ClientTransferNumber.Get()
	}
	toSerialize["signature"] = o.Signature
	toSerialize["sequenceNumber"] = o.SequenceNumber
	return toSerialize, nil
}

type NullableTransferCommenceV1Request struct {
	value *TransferCommenceV1Request
	isSet bool
}

func (v NullableTransferCommenceV1Request) Get() *TransferCommenceV1Request {
	return v.value
}

func (v *NullableTransferCommenceV1Request) Set(val *TransferCommenceV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferCommenceV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferCommenceV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferCommenceV1Request(val *TransferCommenceV1Request) *NullableTransferCommenceV1Request {
	return &NullableTransferCommenceV1Request{value: val, isSet: true}
}

func (v NullableTransferCommenceV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferCommenceV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


