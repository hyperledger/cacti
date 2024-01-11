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

// checks if the RunTransactionSignedRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunTransactionSignedRequestV1{}

// RunTransactionSignedRequestV1 struct for RunTransactionSignedRequestV1
type RunTransactionSignedRequestV1 struct {
	// Signed transaction binary data received from generate-transaction endpoint.
	SignedTransaction string `json:"signedTransaction"`
	BaseConfig *IrohaBaseConfig `json:"baseConfig,omitempty"`
}

// NewRunTransactionSignedRequestV1 instantiates a new RunTransactionSignedRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunTransactionSignedRequestV1(signedTransaction string) *RunTransactionSignedRequestV1 {
	this := RunTransactionSignedRequestV1{}
	this.SignedTransaction = signedTransaction
	return &this
}

// NewRunTransactionSignedRequestV1WithDefaults instantiates a new RunTransactionSignedRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunTransactionSignedRequestV1WithDefaults() *RunTransactionSignedRequestV1 {
	this := RunTransactionSignedRequestV1{}
	return &this
}

// GetSignedTransaction returns the SignedTransaction field value
func (o *RunTransactionSignedRequestV1) GetSignedTransaction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedTransaction
}

// GetSignedTransactionOk returns a tuple with the SignedTransaction field value
// and a boolean to check if the value has been set.
func (o *RunTransactionSignedRequestV1) GetSignedTransactionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedTransaction, true
}

// SetSignedTransaction sets field value
func (o *RunTransactionSignedRequestV1) SetSignedTransaction(v string) {
	o.SignedTransaction = v
}

// GetBaseConfig returns the BaseConfig field value if set, zero value otherwise.
func (o *RunTransactionSignedRequestV1) GetBaseConfig() IrohaBaseConfig {
	if o == nil || IsNil(o.BaseConfig) {
		var ret IrohaBaseConfig
		return ret
	}
	return *o.BaseConfig
}

// GetBaseConfigOk returns a tuple with the BaseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunTransactionSignedRequestV1) GetBaseConfigOk() (*IrohaBaseConfig, bool) {
	if o == nil || IsNil(o.BaseConfig) {
		return nil, false
	}
	return o.BaseConfig, true
}

// HasBaseConfig returns a boolean if a field has been set.
func (o *RunTransactionSignedRequestV1) HasBaseConfig() bool {
	if o != nil && !IsNil(o.BaseConfig) {
		return true
	}

	return false
}

// SetBaseConfig gets a reference to the given IrohaBaseConfig and assigns it to the BaseConfig field.
func (o *RunTransactionSignedRequestV1) SetBaseConfig(v IrohaBaseConfig) {
	o.BaseConfig = &v
}

func (o RunTransactionSignedRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunTransactionSignedRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signedTransaction"] = o.SignedTransaction
	if !IsNil(o.BaseConfig) {
		toSerialize["baseConfig"] = o.BaseConfig
	}
	return toSerialize, nil
}

type NullableRunTransactionSignedRequestV1 struct {
	value *RunTransactionSignedRequestV1
	isSet bool
}

func (v NullableRunTransactionSignedRequestV1) Get() *RunTransactionSignedRequestV1 {
	return v.value
}

func (v *NullableRunTransactionSignedRequestV1) Set(val *RunTransactionSignedRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableRunTransactionSignedRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableRunTransactionSignedRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunTransactionSignedRequestV1(val *RunTransactionSignedRequestV1) *NullableRunTransactionSignedRequestV1 {
	return &NullableRunTransactionSignedRequestV1{value: val, isSet: true}
}

func (v NullableRunTransactionSignedRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunTransactionSignedRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


