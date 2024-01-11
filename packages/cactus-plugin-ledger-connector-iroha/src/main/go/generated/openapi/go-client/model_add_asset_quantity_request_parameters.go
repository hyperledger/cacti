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

// checks if the AddAssetQuantityRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAssetQuantityRequestParameters{}

// AddAssetQuantityRequestParameters The list of arguments to pass in to the transaction request to Add Asset Quantity.
type AddAssetQuantityRequestParameters struct {
	AssetId string `json:"assetId"`
	Amount float32 `json:"amount"`
}

// NewAddAssetQuantityRequestParameters instantiates a new AddAssetQuantityRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAssetQuantityRequestParameters(assetId string, amount float32) *AddAssetQuantityRequestParameters {
	this := AddAssetQuantityRequestParameters{}
	this.AssetId = assetId
	this.Amount = amount
	return &this
}

// NewAddAssetQuantityRequestParametersWithDefaults instantiates a new AddAssetQuantityRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAssetQuantityRequestParametersWithDefaults() *AddAssetQuantityRequestParameters {
	this := AddAssetQuantityRequestParameters{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *AddAssetQuantityRequestParameters) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *AddAssetQuantityRequestParameters) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *AddAssetQuantityRequestParameters) SetAssetId(v string) {
	o.AssetId = v
}

// GetAmount returns the Amount field value
func (o *AddAssetQuantityRequestParameters) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AddAssetQuantityRequestParameters) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AddAssetQuantityRequestParameters) SetAmount(v float32) {
	o.Amount = v
}

func (o AddAssetQuantityRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAssetQuantityRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetId"] = o.AssetId
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableAddAssetQuantityRequestParameters struct {
	value *AddAssetQuantityRequestParameters
	isSet bool
}

func (v NullableAddAssetQuantityRequestParameters) Get() *AddAssetQuantityRequestParameters {
	return v.value
}

func (v *NullableAddAssetQuantityRequestParameters) Set(val *AddAssetQuantityRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAssetQuantityRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAssetQuantityRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAssetQuantityRequestParameters(val *AddAssetQuantityRequestParameters) *NullableAddAssetQuantityRequestParameters {
	return &NullableAddAssetQuantityRequestParameters{value: val, isSet: true}
}

func (v NullableAddAssetQuantityRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAssetQuantityRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


