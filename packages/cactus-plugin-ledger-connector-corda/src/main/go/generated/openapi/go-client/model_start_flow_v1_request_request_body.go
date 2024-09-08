/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the StartFlowV1RequestRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartFlowV1RequestRequestBody{}

// StartFlowV1RequestRequestBody struct for StartFlowV1RequestRequestBody
type StartFlowV1RequestRequestBody struct {
	ChatName *string `json:"chatName,omitempty"`
	OtherMember *string `json:"otherMember,omitempty"`
	Message *string `json:"message,omitempty"`
	NumberOfRecords *string `json:"numberOfRecords,omitempty"`
}

// NewStartFlowV1RequestRequestBody instantiates a new StartFlowV1RequestRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartFlowV1RequestRequestBody() *StartFlowV1RequestRequestBody {
	this := StartFlowV1RequestRequestBody{}
	return &this
}

// NewStartFlowV1RequestRequestBodyWithDefaults instantiates a new StartFlowV1RequestRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartFlowV1RequestRequestBodyWithDefaults() *StartFlowV1RequestRequestBody {
	this := StartFlowV1RequestRequestBody{}
	return &this
}

// GetChatName returns the ChatName field value if set, zero value otherwise.
func (o *StartFlowV1RequestRequestBody) GetChatName() string {
	if o == nil || IsNil(o.ChatName) {
		var ret string
		return ret
	}
	return *o.ChatName
}

// GetChatNameOk returns a tuple with the ChatName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartFlowV1RequestRequestBody) GetChatNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChatName) {
		return nil, false
	}
	return o.ChatName, true
}

// HasChatName returns a boolean if a field has been set.
func (o *StartFlowV1RequestRequestBody) HasChatName() bool {
	if o != nil && !IsNil(o.ChatName) {
		return true
	}

	return false
}

// SetChatName gets a reference to the given string and assigns it to the ChatName field.
func (o *StartFlowV1RequestRequestBody) SetChatName(v string) {
	o.ChatName = &v
}

// GetOtherMember returns the OtherMember field value if set, zero value otherwise.
func (o *StartFlowV1RequestRequestBody) GetOtherMember() string {
	if o == nil || IsNil(o.OtherMember) {
		var ret string
		return ret
	}
	return *o.OtherMember
}

// GetOtherMemberOk returns a tuple with the OtherMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartFlowV1RequestRequestBody) GetOtherMemberOk() (*string, bool) {
	if o == nil || IsNil(o.OtherMember) {
		return nil, false
	}
	return o.OtherMember, true
}

// HasOtherMember returns a boolean if a field has been set.
func (o *StartFlowV1RequestRequestBody) HasOtherMember() bool {
	if o != nil && !IsNil(o.OtherMember) {
		return true
	}

	return false
}

// SetOtherMember gets a reference to the given string and assigns it to the OtherMember field.
func (o *StartFlowV1RequestRequestBody) SetOtherMember(v string) {
	o.OtherMember = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *StartFlowV1RequestRequestBody) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartFlowV1RequestRequestBody) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *StartFlowV1RequestRequestBody) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *StartFlowV1RequestRequestBody) SetMessage(v string) {
	o.Message = &v
}

// GetNumberOfRecords returns the NumberOfRecords field value if set, zero value otherwise.
func (o *StartFlowV1RequestRequestBody) GetNumberOfRecords() string {
	if o == nil || IsNil(o.NumberOfRecords) {
		var ret string
		return ret
	}
	return *o.NumberOfRecords
}

// GetNumberOfRecordsOk returns a tuple with the NumberOfRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartFlowV1RequestRequestBody) GetNumberOfRecordsOk() (*string, bool) {
	if o == nil || IsNil(o.NumberOfRecords) {
		return nil, false
	}
	return o.NumberOfRecords, true
}

// HasNumberOfRecords returns a boolean if a field has been set.
func (o *StartFlowV1RequestRequestBody) HasNumberOfRecords() bool {
	if o != nil && !IsNil(o.NumberOfRecords) {
		return true
	}

	return false
}

// SetNumberOfRecords gets a reference to the given string and assigns it to the NumberOfRecords field.
func (o *StartFlowV1RequestRequestBody) SetNumberOfRecords(v string) {
	o.NumberOfRecords = &v
}

func (o StartFlowV1RequestRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartFlowV1RequestRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChatName) {
		toSerialize["chatName"] = o.ChatName
	}
	if !IsNil(o.OtherMember) {
		toSerialize["otherMember"] = o.OtherMember
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.NumberOfRecords) {
		toSerialize["numberOfRecords"] = o.NumberOfRecords
	}
	return toSerialize, nil
}

type NullableStartFlowV1RequestRequestBody struct {
	value *StartFlowV1RequestRequestBody
	isSet bool
}

func (v NullableStartFlowV1RequestRequestBody) Get() *StartFlowV1RequestRequestBody {
	return v.value
}

func (v *NullableStartFlowV1RequestRequestBody) Set(val *StartFlowV1RequestRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStartFlowV1RequestRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStartFlowV1RequestRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartFlowV1RequestRequestBody(val *StartFlowV1RequestRequestBody) *NullableStartFlowV1RequestRequestBody {
	return &NullableStartFlowV1RequestRequestBody{value: val, isSet: true}
}

func (v NullableStartFlowV1RequestRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartFlowV1RequestRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


