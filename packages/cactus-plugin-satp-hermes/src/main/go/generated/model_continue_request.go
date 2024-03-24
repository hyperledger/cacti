/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
)

// checks if the ContinueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueRequest{}

// ContinueRequest Schema for a request to continue a paused transaction session.
type ContinueRequest struct {
	// A unique identifier for the transaction session to be continued.
	SessionId string `json:"sessionId"`
	// A unique identifier for the transaction context.
	ContextId string `json:"contextId"`
}

// NewContinueRequest instantiates a new ContinueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueRequest(sessionId string, contextId string) *ContinueRequest {
	this := ContinueRequest{}
	this.SessionId = sessionId
	this.ContextId = contextId
	return &this
}

// NewContinueRequestWithDefaults instantiates a new ContinueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueRequestWithDefaults() *ContinueRequest {
	this := ContinueRequest{}
	return &this
}

// GetSessionId returns the SessionId field value
func (o *ContinueRequest) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ContinueRequest) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ContinueRequest) SetSessionId(v string) {
	o.SessionId = v
}

// GetContextId returns the ContextId field value
func (o *ContinueRequest) GetContextId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContextId
}

// GetContextIdOk returns a tuple with the ContextId field value
// and a boolean to check if the value has been set.
func (o *ContinueRequest) GetContextIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContextId, true
}

// SetContextId sets field value
func (o *ContinueRequest) SetContextId(v string) {
	o.ContextId = v
}

func (o ContinueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["contextId"] = o.ContextId
	return toSerialize, nil
}

type NullableContinueRequest struct {
	value *ContinueRequest
	isSet bool
}

func (v NullableContinueRequest) Get() *ContinueRequest {
	return v.value
}

func (v *NullableContinueRequest) Set(val *ContinueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueRequest(val *ContinueRequest) *NullableContinueRequest {
	return &NullableContinueRequest{value: val, isSet: true}
}

func (v NullableContinueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


