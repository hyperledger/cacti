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

// checks if the WebSocketKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebSocketKey{}

// WebSocketKey web-socket key details for signing fabric message with private key stored with external client
type WebSocketKey struct {
	// session Id to access client
	SessionId string `json:"sessionId"`
	// signature of the session ID
	Signature string `json:"signature"`
}

// NewWebSocketKey instantiates a new WebSocketKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebSocketKey(sessionId string, signature string) *WebSocketKey {
	this := WebSocketKey{}
	this.SessionId = sessionId
	this.Signature = signature
	return &this
}

// NewWebSocketKeyWithDefaults instantiates a new WebSocketKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebSocketKeyWithDefaults() *WebSocketKey {
	this := WebSocketKey{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *WebSocketKey) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *WebSocketKey) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *WebSocketKey) SetSessionId(v string) {
	o.SessionId = v
}

// GetSignature returns the Signature field value
func (o *WebSocketKey) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *WebSocketKey) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *WebSocketKey) SetSignature(v string) {
	o.Signature = v
}

func (o WebSocketKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebSocketKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableWebSocketKey struct {
	value *WebSocketKey
	isSet bool
}

func (v NullableWebSocketKey) Get() *WebSocketKey {
	return v.value
}

func (v *NullableWebSocketKey) Set(val *WebSocketKey) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSocketKey) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSocketKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSocketKey(val *WebSocketKey) *NullableWebSocketKey {
	return &NullableWebSocketKey{value: val, isSet: true}
}

func (v NullableWebSocketKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSocketKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


