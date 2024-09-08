/*
Hyperledger Cacti Plugin - Connector CDL

Can perform basic tasks on Fujitsu CDL service.

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-cdl

import (
	"encoding/json"
)

// checks if the GetLineageRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLineageRequestV1{}

// GetLineageRequestV1 struct for GetLineageRequestV1
type GetLineageRequestV1 struct {
	AuthInfo AuthInfoV1 `json:"authInfo"`
	EventId string `json:"eventId"`
	Direction *GetLineageOptionDirectionV1 `json:"direction,omitempty"`
	Depth *string `json:"depth,omitempty"`
}

// NewGetLineageRequestV1 instantiates a new GetLineageRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLineageRequestV1(authInfo AuthInfoV1, eventId string) *GetLineageRequestV1 {
	this := GetLineageRequestV1{}
	this.AuthInfo = authInfo
	this.EventId = eventId
	var depth string = "-1"
	this.Depth = &depth
	return &this
}

// NewGetLineageRequestV1WithDefaults instantiates a new GetLineageRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLineageRequestV1WithDefaults() *GetLineageRequestV1 {
	this := GetLineageRequestV1{}
	var depth string = "-1"
	this.Depth = &depth
	return &this
}

// GetAuthInfo returns the AuthInfo field value
func (o *GetLineageRequestV1) GetAuthInfo() AuthInfoV1 {
	if o == nil {
		var ret AuthInfoV1
		return ret
	}

	return o.AuthInfo
}

// GetAuthInfoOk returns a tuple with the AuthInfo field value
// and a boolean to check if the value has been set.
func (o *GetLineageRequestV1) GetAuthInfoOk() (*AuthInfoV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthInfo, true
}

// SetAuthInfo sets field value
func (o *GetLineageRequestV1) SetAuthInfo(v AuthInfoV1) {
	o.AuthInfo = v
}

// GetEventId returns the EventId field value
func (o *GetLineageRequestV1) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *GetLineageRequestV1) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *GetLineageRequestV1) SetEventId(v string) {
	o.EventId = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *GetLineageRequestV1) GetDirection() GetLineageOptionDirectionV1 {
	if o == nil || IsNil(o.Direction) {
		var ret GetLineageOptionDirectionV1
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLineageRequestV1) GetDirectionOk() (*GetLineageOptionDirectionV1, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *GetLineageRequestV1) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given GetLineageOptionDirectionV1 and assigns it to the Direction field.
func (o *GetLineageRequestV1) SetDirection(v GetLineageOptionDirectionV1) {
	o.Direction = &v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *GetLineageRequestV1) GetDepth() string {
	if o == nil || IsNil(o.Depth) {
		var ret string
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLineageRequestV1) GetDepthOk() (*string, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *GetLineageRequestV1) HasDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given string and assigns it to the Depth field.
func (o *GetLineageRequestV1) SetDepth(v string) {
	o.Depth = &v
}

func (o GetLineageRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLineageRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authInfo"] = o.AuthInfo
	toSerialize["eventId"] = o.EventId
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}
	return toSerialize, nil
}

type NullableGetLineageRequestV1 struct {
	value *GetLineageRequestV1
	isSet bool
}

func (v NullableGetLineageRequestV1) Get() *GetLineageRequestV1 {
	return v.value
}

func (v *NullableGetLineageRequestV1) Set(val *GetLineageRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLineageRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLineageRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLineageRequestV1(val *GetLineageRequestV1) *NullableGetLineageRequestV1 {
	return &NullableGetLineageRequestV1{value: val, isSet: true}
}

func (v NullableGetLineageRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLineageRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


