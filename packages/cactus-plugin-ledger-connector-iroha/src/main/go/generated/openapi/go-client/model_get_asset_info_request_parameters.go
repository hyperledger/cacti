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

// checks if the GetAssetInfoRequestParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssetInfoRequestParameters{}

// GetAssetInfoRequestParameters The list of arguments to pass in to the transaction request to Get Asset Info.
type GetAssetInfoRequestParameters struct {
	AssetId string `json:"assetId"`
}

// NewGetAssetInfoRequestParameters instantiates a new GetAssetInfoRequestParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssetInfoRequestParameters(assetId string) *GetAssetInfoRequestParameters {
	this := GetAssetInfoRequestParameters{}
	this.AssetId = assetId
	return &this
}

// NewGetAssetInfoRequestParametersWithDefaults instantiates a new GetAssetInfoRequestParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssetInfoRequestParametersWithDefaults() *GetAssetInfoRequestParameters {
	this := GetAssetInfoRequestParameters{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *GetAssetInfoRequestParameters) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *GetAssetInfoRequestParameters) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *GetAssetInfoRequestParameters) SetAssetId(v string) {
	o.AssetId = v
}

func (o GetAssetInfoRequestParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssetInfoRequestParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetId"] = o.AssetId
	return toSerialize, nil
}

type NullableGetAssetInfoRequestParameters struct {
	value *GetAssetInfoRequestParameters
	isSet bool
}

func (v NullableGetAssetInfoRequestParameters) Get() *GetAssetInfoRequestParameters {
	return v.value
}

func (v *NullableGetAssetInfoRequestParameters) Set(val *GetAssetInfoRequestParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssetInfoRequestParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssetInfoRequestParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssetInfoRequestParameters(val *GetAssetInfoRequestParameters) *NullableGetAssetInfoRequestParameters {
	return &NullableGetAssetInfoRequestParameters{value: val, isSet: true}
}

func (v NullableGetAssetInfoRequestParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssetInfoRequestParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


