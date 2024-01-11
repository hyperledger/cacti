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

// checks if the CompareAndSetAccountDetailRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompareAndSetAccountDetailRequestParameters{}

// CompareAndSetAccountDetailRequestParameters The list of arguments to pass in to the transaction request to Compare And Set Account Detail.
type CompareAndSetAccountDetailRequestParameters struct {
	AccountId string `json:"accountId"`
	Key string `json:"key"`
	Value string `json:"value"`
	OldValue *string `json:"oldValue,omitempty"`
	CheckEmpty bool `json:"check_empty"`
}

// NewCompareAndSetAccountDetailRequestParameters instantiates a new CompareAndSetAccountDetailRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompareAndSetAccountDetailRequestParameters(accountId string, key string, value string, checkEmpty bool) *CompareAndSetAccountDetailRequestParameters {
	this := CompareAndSetAccountDetailRequestParameters{}
	this.AccountId = accountId
	this.Key = key
	this.Value = value
	this.CheckEmpty = checkEmpty
	return &this
}

// NewCompareAndSetAccountDetailRequestParametersWithDefaults instantiates a new CompareAndSetAccountDetailRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompareAndSetAccountDetailRequestParametersWithDefaults() *CompareAndSetAccountDetailRequestParameters {
	this := CompareAndSetAccountDetailRequestParameters{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *CompareAndSetAccountDetailRequestParameters) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CompareAndSetAccountDetailRequestParameters) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *CompareAndSetAccountDetailRequestParameters) SetAccountId(v string) {
	o.AccountId = v
}

// GetKey returns the Key field value
func (o *CompareAndSetAccountDetailRequestParameters) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CompareAndSetAccountDetailRequestParameters) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CompareAndSetAccountDetailRequestParameters) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *CompareAndSetAccountDetailRequestParameters) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CompareAndSetAccountDetailRequestParameters) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CompareAndSetAccountDetailRequestParameters) SetValue(v string) {
	o.Value = v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise.
func (o *CompareAndSetAccountDetailRequestParameters) GetOldValue() string {
	if o == nil || IsNil(o.OldValue) {
		var ret string
		return ret
	}
	return *o.OldValue
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompareAndSetAccountDetailRequestParameters) GetOldValueOk() (*string, bool) {
	if o == nil || IsNil(o.OldValue) {
		return nil, false
	}
	return o.OldValue, true
}

// HasOldValue returns a boolean if a field has been set.
func (o *CompareAndSetAccountDetailRequestParameters) HasOldValue() bool {
	if o != nil && !IsNil(o.OldValue) {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given string and assigns it to the OldValue field.
func (o *CompareAndSetAccountDetailRequestParameters) SetOldValue(v string) {
	o.OldValue = &v
}

// GetCheckEmpty returns the CheckEmpty field value
func (o *CompareAndSetAccountDetailRequestParameters) GetCheckEmpty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CheckEmpty
}

// GetCheckEmptyOk returns a tuple with the CheckEmpty field value
// and a boolean to check if the value has been set.
func (o *CompareAndSetAccountDetailRequestParameters) GetCheckEmptyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckEmpty, true
}

// SetCheckEmpty sets field value
func (o *CompareAndSetAccountDetailRequestParameters) SetCheckEmpty(v bool) {
	o.CheckEmpty = v
}

func (o CompareAndSetAccountDetailRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompareAndSetAccountDetailRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	if !IsNil(o.OldValue) {
		toSerialize["oldValue"] = o.OldValue
	}
	toSerialize["check_empty"] = o.CheckEmpty
	return toSerialize, nil
}

type NullableCompareAndSetAccountDetailRequestParameters struct {
	value *CompareAndSetAccountDetailRequestParameters
	isSet bool
}

func (v NullableCompareAndSetAccountDetailRequestParameters) Get() *CompareAndSetAccountDetailRequestParameters {
	return v.value
}

func (v *NullableCompareAndSetAccountDetailRequestParameters) Set(val *CompareAndSetAccountDetailRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCompareAndSetAccountDetailRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCompareAndSetAccountDetailRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompareAndSetAccountDetailRequestParameters(val *CompareAndSetAccountDetailRequestParameters) *NullableCompareAndSetAccountDetailRequestParameters {
	return &NullableCompareAndSetAccountDetailRequestParameters{value: val, isSet: true}
}

func (v NullableCompareAndSetAccountDetailRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompareAndSetAccountDetailRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


