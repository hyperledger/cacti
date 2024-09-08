/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the FabricX509CertificateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricX509CertificateV1{}

// FabricX509CertificateV1 Transaction endorser certificate object
type FabricX509CertificateV1 struct {
	Issuer string `json:"issuer"`
	SerialNumber string `json:"serialNumber"`
	Subject string `json:"subject"`
	SubjectAltName string `json:"subjectAltName"`
	ValidFrom string `json:"validFrom"`
	ValidTo string `json:"validTo"`
	Pem string `json:"pem"`
}

// NewFabricX509CertificateV1 instantiates a new FabricX509CertificateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricX509CertificateV1(issuer string, serialNumber string, subject string, subjectAltName string, validFrom string, validTo string, pem string) *FabricX509CertificateV1 {
	this := FabricX509CertificateV1{}
	this.Issuer = issuer
	this.SerialNumber = serialNumber
	this.Subject = subject
	this.SubjectAltName = subjectAltName
	this.ValidFrom = validFrom
	this.ValidTo = validTo
	this.Pem = pem
	return &this
}

// NewFabricX509CertificateV1WithDefaults instantiates a new FabricX509CertificateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricX509CertificateV1WithDefaults() *FabricX509CertificateV1 {
	this := FabricX509CertificateV1{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *FabricX509CertificateV1) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *FabricX509CertificateV1) SetIssuer(v string) {
	o.Issuer = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *FabricX509CertificateV1) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *FabricX509CertificateV1) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetSubject returns the Subject field value
func (o *FabricX509CertificateV1) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *FabricX509CertificateV1) SetSubject(v string) {
	o.Subject = v
}

// GetSubjectAltName returns the SubjectAltName field value
func (o *FabricX509CertificateV1) GetSubjectAltName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectAltName
}

// GetSubjectAltNameOk returns a tuple with the SubjectAltName field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetSubjectAltNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectAltName, true
}

// SetSubjectAltName sets field value
func (o *FabricX509CertificateV1) SetSubjectAltName(v string) {
	o.SubjectAltName = v
}

// GetValidFrom returns the ValidFrom field value
func (o *FabricX509CertificateV1) GetValidFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetValidFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *FabricX509CertificateV1) SetValidFrom(v string) {
	o.ValidFrom = v
}

// GetValidTo returns the ValidTo field value
func (o *FabricX509CertificateV1) GetValidTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetValidToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidTo, true
}

// SetValidTo sets field value
func (o *FabricX509CertificateV1) SetValidTo(v string) {
	o.ValidTo = v
}

// GetPem returns the Pem field value
func (o *FabricX509CertificateV1) GetPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pem
}

// GetPemOk returns a tuple with the Pem field value
// and a boolean to check if the value has been set.
func (o *FabricX509CertificateV1) GetPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pem, true
}

// SetPem sets field value
func (o *FabricX509CertificateV1) SetPem(v string) {
	o.Pem = v
}

func (o FabricX509CertificateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricX509CertificateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuer"] = o.Issuer
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["subject"] = o.Subject
	toSerialize["subjectAltName"] = o.SubjectAltName
	toSerialize["validFrom"] = o.ValidFrom
	toSerialize["validTo"] = o.ValidTo
	toSerialize["pem"] = o.Pem
	return toSerialize, nil
}

type NullableFabricX509CertificateV1 struct {
	value *FabricX509CertificateV1
	isSet bool
}

func (v NullableFabricX509CertificateV1) Get() *FabricX509CertificateV1 {
	return v.value
}

func (v *NullableFabricX509CertificateV1) Set(val *FabricX509CertificateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricX509CertificateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricX509CertificateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricX509CertificateV1(val *FabricX509CertificateV1) *NullableFabricX509CertificateV1 {
	return &NullableFabricX509CertificateV1{value: val, isSet: true}
}

func (v NullableFabricX509CertificateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricX509CertificateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


