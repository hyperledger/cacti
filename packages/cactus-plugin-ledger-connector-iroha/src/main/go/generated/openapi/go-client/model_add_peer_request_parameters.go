/*
Hyperledger Cactus Plugin - Connector Iroha

Can perform basic tasks on a Iroha ledger

API version: v2.0.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha

import (
	"encoding/json"
)

// checks if the AddPeerRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPeerRequestParameters{}

// AddPeerRequestParameters The list of arguments to pass in to the transaction request to Add Peer.
type AddPeerRequestParameters struct {
	Address string `json:"address"`
	PeerKey string `json:"peerKey"`
	TlsCertificate *string `json:"tlsCertificate,omitempty"`
	SyncingPeer *bool `json:"syncingPeer,omitempty"`
}

// NewAddPeerRequestParameters instantiates a new AddPeerRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPeerRequestParameters(address string, peerKey string) *AddPeerRequestParameters {
	this := AddPeerRequestParameters{}
	this.Address = address
	this.PeerKey = peerKey
	return &this
}

// NewAddPeerRequestParametersWithDefaults instantiates a new AddPeerRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPeerRequestParametersWithDefaults() *AddPeerRequestParameters {
	this := AddPeerRequestParameters{}
	return &this
}

// GetAddress returns the Address field value
func (o *AddPeerRequestParameters) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AddPeerRequestParameters) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AddPeerRequestParameters) SetAddress(v string) {
	o.Address = v
}

// GetPeerKey returns the PeerKey field value
func (o *AddPeerRequestParameters) GetPeerKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeerKey
}

// GetPeerKeyOk returns a tuple with the PeerKey field value
// and a boolean to check if the value has been set.
func (o *AddPeerRequestParameters) GetPeerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerKey, true
}

// SetPeerKey sets field value
func (o *AddPeerRequestParameters) SetPeerKey(v string) {
	o.PeerKey = v
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *AddPeerRequestParameters) GetTlsCertificate() string {
	if o == nil || IsNil(o.TlsCertificate) {
		var ret string
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeerRequestParameters) GetTlsCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.TlsCertificate) {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *AddPeerRequestParameters) HasTlsCertificate() bool {
	if o != nil && !IsNil(o.TlsCertificate) {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given string and assigns it to the TlsCertificate field.
func (o *AddPeerRequestParameters) SetTlsCertificate(v string) {
	o.TlsCertificate = &v
}

// GetSyncingPeer returns the SyncingPeer field value if set, zero value otherwise.
func (o *AddPeerRequestParameters) GetSyncingPeer() bool {
	if o == nil || IsNil(o.SyncingPeer) {
		var ret bool
		return ret
	}
	return *o.SyncingPeer
}

// GetSyncingPeerOk returns a tuple with the SyncingPeer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeerRequestParameters) GetSyncingPeerOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncingPeer) {
		return nil, false
	}
	return o.SyncingPeer, true
}

// HasSyncingPeer returns a boolean if a field has been set.
func (o *AddPeerRequestParameters) HasSyncingPeer() bool {
	if o != nil && !IsNil(o.SyncingPeer) {
		return true
	}

	return false
}

// SetSyncingPeer gets a reference to the given bool and assigns it to the SyncingPeer field.
func (o *AddPeerRequestParameters) SetSyncingPeer(v bool) {
	o.SyncingPeer = &v
}

func (o AddPeerRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPeerRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["peerKey"] = o.PeerKey
	if !IsNil(o.TlsCertificate) {
		toSerialize["tlsCertificate"] = o.TlsCertificate
	}
	if !IsNil(o.SyncingPeer) {
		toSerialize["syncingPeer"] = o.SyncingPeer
	}
	return toSerialize, nil
}

type NullableAddPeerRequestParameters struct {
	value *AddPeerRequestParameters
	isSet bool
}

func (v NullableAddPeerRequestParameters) Get() *AddPeerRequestParameters {
	return v.value
}

func (v *NullableAddPeerRequestParameters) Set(val *AddPeerRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPeerRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPeerRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPeerRequestParameters(val *AddPeerRequestParameters) *NullableAddPeerRequestParameters {
	return &NullableAddPeerRequestParameters{value: val, isSet: true}
}

func (v NullableAddPeerRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPeerRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


