/*
Hyperledger Cacti Plugin - Connector Corda

Can perform basic tasks on a Corda ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-corda

import (
	"encoding/json"
)

// checks if the DeployContractJarsBadRequestV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractJarsBadRequestV1Response{}

// DeployContractJarsBadRequestV1Response struct for DeployContractJarsBadRequestV1Response
type DeployContractJarsBadRequestV1Response struct {
	Errors []string `json:"errors"`
}

// NewDeployContractJarsBadRequestV1Response instantiates a new DeployContractJarsBadRequestV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractJarsBadRequestV1Response(errors []string) *DeployContractJarsBadRequestV1Response {
	this := DeployContractJarsBadRequestV1Response{}
	this.Errors = errors
	return &this
}

// NewDeployContractJarsBadRequestV1ResponseWithDefaults instantiates a new DeployContractJarsBadRequestV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractJarsBadRequestV1ResponseWithDefaults() *DeployContractJarsBadRequestV1Response {
	this := DeployContractJarsBadRequestV1Response{}
	return &this
}

// GetErrors returns the Errors field value
func (o *DeployContractJarsBadRequestV1Response) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *DeployContractJarsBadRequestV1Response) GetErrorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *DeployContractJarsBadRequestV1Response) SetErrors(v []string) {
	o.Errors = v
}

func (o DeployContractJarsBadRequestV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployContractJarsBadRequestV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableDeployContractJarsBadRequestV1Response struct {
	value *DeployContractJarsBadRequestV1Response
	isSet bool
}

func (v NullableDeployContractJarsBadRequestV1Response) Get() *DeployContractJarsBadRequestV1Response {
	return v.value
}

func (v *NullableDeployContractJarsBadRequestV1Response) Set(val *DeployContractJarsBadRequestV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractJarsBadRequestV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractJarsBadRequestV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractJarsBadRequestV1Response(val *DeployContractJarsBadRequestV1Response) *NullableDeployContractJarsBadRequestV1Response {
	return &NullableDeployContractJarsBadRequestV1Response{value: val, isSet: true}
}

func (v NullableDeployContractJarsBadRequestV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractJarsBadRequestV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


