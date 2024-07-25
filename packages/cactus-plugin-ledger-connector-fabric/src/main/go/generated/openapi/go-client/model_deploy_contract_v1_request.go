/*
Hyperledger Cactus Plugin - Connector Fabric

Can perform basic tasks on a fabric ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-fabric

import (
	"encoding/json"
)

// checks if the DeployContractV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployContractV1Request{}

// DeployContractV1Request struct for DeployContractV1Request
type DeployContractV1Request struct {
	CcLang ChainCodeProgrammingLanguage `json:"ccLang"`
	// File-system path pointing at the CA file.
	CaFile string `json:"caFile"`
	// Ordering service endpoint specified as <hostname or IP address>:<port>
	Orderer string `json:"orderer"`
	// The hostname override to use when validating the TLS connection to the orderer
	OrdererTLSHostnameOverride string `json:"ordererTLSHostnameOverride"`
	// Timeout for client to connect (default 3s)
	ConnTimeout *int32 `json:"connTimeout,omitempty"`
	// Passed in to the peer via the --signature-policy argument on the command line. See also: https://hyperledger-fabric.readthedocs.io/en/release-2.2/endorsement-policies.html#setting-chaincode-level-endorsement-policies
	SignaturePolicy *string `json:"signaturePolicy,omitempty"`
	// Name of the collections config file as present in the sourceFiles array of the request.
	CollectionsConfigFile *string `json:"collectionsConfigFile,omitempty"`
	// The name of the Fabric channel where the contract will get instantiated.
	ChannelId string `json:"channelId"`
	TargetOrganizations []DeploymentTargetOrganization `json:"targetOrganizations"`
	ConstructorArgs *DeployContractGoSourceV1RequestConstructorArgs `json:"constructorArgs,omitempty"`
	CcSequence float32 `json:"ccSequence"`
	CcVersion string `json:"ccVersion"`
	CcName string `json:"ccName"`
	// Human readable label to uniquely identify the contract. Recommended to include in this at least the contract name and the exact version in order to make it easily distinguishable from other deployments of the same contract.
	CcLabel string `json:"ccLabel"`
	// The your-smart-contract.go file where the functionality of your contract is implemented.
	SourceFiles []FileBase64 `json:"sourceFiles"`
}

// NewDeployContractV1Request instantiates a new DeployContractV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployContractV1Request(ccLang ChainCodeProgrammingLanguage, caFile string, orderer string, ordererTLSHostnameOverride string, channelId string, targetOrganizations []DeploymentTargetOrganization, ccSequence float32, ccVersion string, ccName string, ccLabel string, sourceFiles []FileBase64) *DeployContractV1Request {
	this := DeployContractV1Request{}
	this.CcLang = ccLang
	this.CaFile = caFile
	this.Orderer = orderer
	this.OrdererTLSHostnameOverride = ordererTLSHostnameOverride
	this.ChannelId = channelId
	this.TargetOrganizations = targetOrganizations
	this.CcSequence = ccSequence
	this.CcVersion = ccVersion
	this.CcName = ccName
	this.CcLabel = ccLabel
	this.SourceFiles = sourceFiles
	return &this
}

// NewDeployContractV1RequestWithDefaults instantiates a new DeployContractV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployContractV1RequestWithDefaults() *DeployContractV1Request {
	this := DeployContractV1Request{}
	return &this
}

// GetCcLang returns the CcLang field value
func (o *DeployContractV1Request) GetCcLang() ChainCodeProgrammingLanguage {
	if o == nil {
		var ret ChainCodeProgrammingLanguage
		return ret
	}

	return o.CcLang
}

// GetCcLangOk returns a tuple with the CcLang field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCcLangOk() (*ChainCodeProgrammingLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcLang, true
}

// SetCcLang sets field value
func (o *DeployContractV1Request) SetCcLang(v ChainCodeProgrammingLanguage) {
	o.CcLang = v
}

// GetCaFile returns the CaFile field value
func (o *DeployContractV1Request) GetCaFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaFile
}

// GetCaFileOk returns a tuple with the CaFile field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCaFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaFile, true
}

// SetCaFile sets field value
func (o *DeployContractV1Request) SetCaFile(v string) {
	o.CaFile = v
}

// GetOrderer returns the Orderer field value
func (o *DeployContractV1Request) GetOrderer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orderer
}

// GetOrdererOk returns a tuple with the Orderer field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetOrdererOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orderer, true
}

// SetOrderer sets field value
func (o *DeployContractV1Request) SetOrderer(v string) {
	o.Orderer = v
}

// GetOrdererTLSHostnameOverride returns the OrdererTLSHostnameOverride field value
func (o *DeployContractV1Request) GetOrdererTLSHostnameOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrdererTLSHostnameOverride
}

// GetOrdererTLSHostnameOverrideOk returns a tuple with the OrdererTLSHostnameOverride field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetOrdererTLSHostnameOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrdererTLSHostnameOverride, true
}

// SetOrdererTLSHostnameOverride sets field value
func (o *DeployContractV1Request) SetOrdererTLSHostnameOverride(v string) {
	o.OrdererTLSHostnameOverride = v
}

// GetConnTimeout returns the ConnTimeout field value if set, zero value otherwise.
func (o *DeployContractV1Request) GetConnTimeout() int32 {
	if o == nil || IsNil(o.ConnTimeout) {
		var ret int32
		return ret
	}
	return *o.ConnTimeout
}

// GetConnTimeoutOk returns a tuple with the ConnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetConnTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnTimeout) {
		return nil, false
	}
	return o.ConnTimeout, true
}

// HasConnTimeout returns a boolean if a field has been set.
func (o *DeployContractV1Request) HasConnTimeout() bool {
	if o != nil && !IsNil(o.ConnTimeout) {
		return true
	}

	return false
}

// SetConnTimeout gets a reference to the given int32 and assigns it to the ConnTimeout field.
func (o *DeployContractV1Request) SetConnTimeout(v int32) {
	o.ConnTimeout = &v
}

// GetSignaturePolicy returns the SignaturePolicy field value if set, zero value otherwise.
func (o *DeployContractV1Request) GetSignaturePolicy() string {
	if o == nil || IsNil(o.SignaturePolicy) {
		var ret string
		return ret
	}
	return *o.SignaturePolicy
}

// GetSignaturePolicyOk returns a tuple with the SignaturePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetSignaturePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.SignaturePolicy) {
		return nil, false
	}
	return o.SignaturePolicy, true
}

// HasSignaturePolicy returns a boolean if a field has been set.
func (o *DeployContractV1Request) HasSignaturePolicy() bool {
	if o != nil && !IsNil(o.SignaturePolicy) {
		return true
	}

	return false
}

// SetSignaturePolicy gets a reference to the given string and assigns it to the SignaturePolicy field.
func (o *DeployContractV1Request) SetSignaturePolicy(v string) {
	o.SignaturePolicy = &v
}

// GetCollectionsConfigFile returns the CollectionsConfigFile field value if set, zero value otherwise.
func (o *DeployContractV1Request) GetCollectionsConfigFile() string {
	if o == nil || IsNil(o.CollectionsConfigFile) {
		var ret string
		return ret
	}
	return *o.CollectionsConfigFile
}

// GetCollectionsConfigFileOk returns a tuple with the CollectionsConfigFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCollectionsConfigFileOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionsConfigFile) {
		return nil, false
	}
	return o.CollectionsConfigFile, true
}

// HasCollectionsConfigFile returns a boolean if a field has been set.
func (o *DeployContractV1Request) HasCollectionsConfigFile() bool {
	if o != nil && !IsNil(o.CollectionsConfigFile) {
		return true
	}

	return false
}

// SetCollectionsConfigFile gets a reference to the given string and assigns it to the CollectionsConfigFile field.
func (o *DeployContractV1Request) SetCollectionsConfigFile(v string) {
	o.CollectionsConfigFile = &v
}

// GetChannelId returns the ChannelId field value
func (o *DeployContractV1Request) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *DeployContractV1Request) SetChannelId(v string) {
	o.ChannelId = v
}

// GetTargetOrganizations returns the TargetOrganizations field value
func (o *DeployContractV1Request) GetTargetOrganizations() []DeploymentTargetOrganization {
	if o == nil {
		var ret []DeploymentTargetOrganization
		return ret
	}

	return o.TargetOrganizations
}

// GetTargetOrganizationsOk returns a tuple with the TargetOrganizations field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetTargetOrganizationsOk() ([]DeploymentTargetOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetOrganizations, true
}

// SetTargetOrganizations sets field value
func (o *DeployContractV1Request) SetTargetOrganizations(v []DeploymentTargetOrganization) {
	o.TargetOrganizations = v
}

// GetConstructorArgs returns the ConstructorArgs field value if set, zero value otherwise.
func (o *DeployContractV1Request) GetConstructorArgs() DeployContractGoSourceV1RequestConstructorArgs {
	if o == nil || IsNil(o.ConstructorArgs) {
		var ret DeployContractGoSourceV1RequestConstructorArgs
		return ret
	}
	return *o.ConstructorArgs
}

// GetConstructorArgsOk returns a tuple with the ConstructorArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetConstructorArgsOk() (*DeployContractGoSourceV1RequestConstructorArgs, bool) {
	if o == nil || IsNil(o.ConstructorArgs) {
		return nil, false
	}
	return o.ConstructorArgs, true
}

// HasConstructorArgs returns a boolean if a field has been set.
func (o *DeployContractV1Request) HasConstructorArgs() bool {
	if o != nil && !IsNil(o.ConstructorArgs) {
		return true
	}

	return false
}

// SetConstructorArgs gets a reference to the given DeployContractGoSourceV1RequestConstructorArgs and assigns it to the ConstructorArgs field.
func (o *DeployContractV1Request) SetConstructorArgs(v DeployContractGoSourceV1RequestConstructorArgs) {
	o.ConstructorArgs = &v
}

// GetCcSequence returns the CcSequence field value
func (o *DeployContractV1Request) GetCcSequence() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CcSequence
}

// GetCcSequenceOk returns a tuple with the CcSequence field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCcSequenceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcSequence, true
}

// SetCcSequence sets field value
func (o *DeployContractV1Request) SetCcSequence(v float32) {
	o.CcSequence = v
}

// GetCcVersion returns the CcVersion field value
func (o *DeployContractV1Request) GetCcVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CcVersion
}

// GetCcVersionOk returns a tuple with the CcVersion field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCcVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcVersion, true
}

// SetCcVersion sets field value
func (o *DeployContractV1Request) SetCcVersion(v string) {
	o.CcVersion = v
}

// GetCcName returns the CcName field value
func (o *DeployContractV1Request) GetCcName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CcName
}

// GetCcNameOk returns a tuple with the CcName field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCcNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcName, true
}

// SetCcName sets field value
func (o *DeployContractV1Request) SetCcName(v string) {
	o.CcName = v
}

// GetCcLabel returns the CcLabel field value
func (o *DeployContractV1Request) GetCcLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CcLabel
}

// GetCcLabelOk returns a tuple with the CcLabel field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetCcLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcLabel, true
}

// SetCcLabel sets field value
func (o *DeployContractV1Request) SetCcLabel(v string) {
	o.CcLabel = v
}

// GetSourceFiles returns the SourceFiles field value
func (o *DeployContractV1Request) GetSourceFiles() []FileBase64 {
	if o == nil {
		var ret []FileBase64
		return ret
	}

	return o.SourceFiles
}

// GetSourceFilesOk returns a tuple with the SourceFiles field value
// and a boolean to check if the value has been set.
func (o *DeployContractV1Request) GetSourceFilesOk() ([]FileBase64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceFiles, true
}

// SetSourceFiles sets field value
func (o *DeployContractV1Request) SetSourceFiles(v []FileBase64) {
	o.SourceFiles = v
}

func (o DeployContractV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployContractV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ccLang"] = o.CcLang
	toSerialize["caFile"] = o.CaFile
	toSerialize["orderer"] = o.Orderer
	toSerialize["ordererTLSHostnameOverride"] = o.OrdererTLSHostnameOverride
	if !IsNil(o.ConnTimeout) {
		toSerialize["connTimeout"] = o.ConnTimeout
	}
	if !IsNil(o.SignaturePolicy) {
		toSerialize["signaturePolicy"] = o.SignaturePolicy
	}
	if !IsNil(o.CollectionsConfigFile) {
		toSerialize["collectionsConfigFile"] = o.CollectionsConfigFile
	}
	toSerialize["channelId"] = o.ChannelId
	toSerialize["targetOrganizations"] = o.TargetOrganizations
	if !IsNil(o.ConstructorArgs) {
		toSerialize["constructorArgs"] = o.ConstructorArgs
	}
	toSerialize["ccSequence"] = o.CcSequence
	toSerialize["ccVersion"] = o.CcVersion
	toSerialize["ccName"] = o.CcName
	toSerialize["ccLabel"] = o.CcLabel
	toSerialize["sourceFiles"] = o.SourceFiles
	return toSerialize, nil
}

type NullableDeployContractV1Request struct {
	value *DeployContractV1Request
	isSet bool
}

func (v NullableDeployContractV1Request) Get() *DeployContractV1Request {
	return v.value
}

func (v *NullableDeployContractV1Request) Set(val *DeployContractV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployContractV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployContractV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployContractV1Request(val *DeployContractV1Request) *NullableDeployContractV1Request {
	return &NullableDeployContractV1Request{value: val, isSet: true}
}

func (v NullableDeployContractV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployContractV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


