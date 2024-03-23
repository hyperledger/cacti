/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: 0.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
)

// checks if the CancelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelRequest{}

// CancelRequest Request to cancel an ongoing transaction session, identified by the session ID.
type CancelRequest struct {
	// Unique identifier (UUID) for the session.
	SessionID string `json:"sessionID"`
}

// NewCancelRequest instantiates a new CancelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelRequest(sessionID string) *CancelRequest {
	this := CancelRequest{}
	this.SessionID = sessionID
	return &this
}

// NewCancelRequestWithDefaults instantiates a new CancelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelRequestWithDefaults() *CancelRequest {
	this := CancelRequest{}
	return &this
}

// GetSessionID returns the SessionID field value
func (o *CancelRequest) GetSessionID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionID
}

// GetSessionIDOk returns a tuple with the SessionID field value
// and a boolean to check if the value has been set.
func (o *CancelRequest) GetSessionIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionID, true
}

// SetSessionID sets field value
func (o *CancelRequest) SetSessionID(v string) {
	o.SessionID = v
}

func (o CancelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionID"] = o.SessionID
	return toSerialize, nil
}

type NullableCancelRequest struct {
	value *CancelRequest
	isSet bool
}

func (v NullableCancelRequest) Get() *CancelRequest {
	return v.value
}

func (v *NullableCancelRequest) Set(val *CancelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelRequest(val *CancelRequest) *NullableCancelRequest {
	return &NullableCancelRequest{value: val, isSet: true}
}

func (v NullableCancelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


