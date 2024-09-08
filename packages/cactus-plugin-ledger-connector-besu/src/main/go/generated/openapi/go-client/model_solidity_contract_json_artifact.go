/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0-rc.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the SolidityContractJsonArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolidityContractJsonArtifact{}

// SolidityContractJsonArtifact struct for SolidityContractJsonArtifact
type SolidityContractJsonArtifact struct {
	ContractName string `json:"contractName"`
	Metadata *string `json:"metadata,omitempty"`
	Bytecode *string `json:"bytecode,omitempty"`
	DeployedBytecode *string `json:"deployedBytecode,omitempty"`
	SourceMap *string `json:"sourceMap,omitempty"`
	DeployedSourceMap *string `json:"deployedSourceMap,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty"`
	Compiler *SolidityContractJsonArtifactCompiler `json:"compiler,omitempty"`
	FunctionHashes map[string]interface{} `json:"functionHashes,omitempty"`
	GasEstimates *SolidityContractJsonArtifactGasEstimates `json:"gasEstimates,omitempty"`
}

// NewSolidityContractJsonArtifact instantiates a new SolidityContractJsonArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolidityContractJsonArtifact(contractName string) *SolidityContractJsonArtifact {
	this := SolidityContractJsonArtifact{}
	this.ContractName = contractName
	return &this
}

// NewSolidityContractJsonArtifactWithDefaults instantiates a new SolidityContractJsonArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolidityContractJsonArtifactWithDefaults() *SolidityContractJsonArtifact {
	this := SolidityContractJsonArtifact{}
	return &this
}

// GetContractName returns the ContractName field value
func (o *SolidityContractJsonArtifact) GetContractName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetContractNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractName, true
}

// SetContractName sets field value
func (o *SolidityContractJsonArtifact) SetContractName(v string) {
	o.ContractName = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *SolidityContractJsonArtifact) SetMetadata(v string) {
	o.Metadata = &v
}

// GetBytecode returns the Bytecode field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetBytecode() string {
	if o == nil || IsNil(o.Bytecode) {
		var ret string
		return ret
	}
	return *o.Bytecode
}

// GetBytecodeOk returns a tuple with the Bytecode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetBytecodeOk() (*string, bool) {
	if o == nil || IsNil(o.Bytecode) {
		return nil, false
	}
	return o.Bytecode, true
}

// HasBytecode returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasBytecode() bool {
	if o != nil && !IsNil(o.Bytecode) {
		return true
	}

	return false
}

// SetBytecode gets a reference to the given string and assigns it to the Bytecode field.
func (o *SolidityContractJsonArtifact) SetBytecode(v string) {
	o.Bytecode = &v
}

// GetDeployedBytecode returns the DeployedBytecode field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetDeployedBytecode() string {
	if o == nil || IsNil(o.DeployedBytecode) {
		var ret string
		return ret
	}
	return *o.DeployedBytecode
}

// GetDeployedBytecodeOk returns a tuple with the DeployedBytecode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetDeployedBytecodeOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedBytecode) {
		return nil, false
	}
	return o.DeployedBytecode, true
}

// HasDeployedBytecode returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasDeployedBytecode() bool {
	if o != nil && !IsNil(o.DeployedBytecode) {
		return true
	}

	return false
}

// SetDeployedBytecode gets a reference to the given string and assigns it to the DeployedBytecode field.
func (o *SolidityContractJsonArtifact) SetDeployedBytecode(v string) {
	o.DeployedBytecode = &v
}

// GetSourceMap returns the SourceMap field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetSourceMap() string {
	if o == nil || IsNil(o.SourceMap) {
		var ret string
		return ret
	}
	return *o.SourceMap
}

// GetSourceMapOk returns a tuple with the SourceMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetSourceMapOk() (*string, bool) {
	if o == nil || IsNil(o.SourceMap) {
		return nil, false
	}
	return o.SourceMap, true
}

// HasSourceMap returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasSourceMap() bool {
	if o != nil && !IsNil(o.SourceMap) {
		return true
	}

	return false
}

// SetSourceMap gets a reference to the given string and assigns it to the SourceMap field.
func (o *SolidityContractJsonArtifact) SetSourceMap(v string) {
	o.SourceMap = &v
}

// GetDeployedSourceMap returns the DeployedSourceMap field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetDeployedSourceMap() string {
	if o == nil || IsNil(o.DeployedSourceMap) {
		var ret string
		return ret
	}
	return *o.DeployedSourceMap
}

// GetDeployedSourceMapOk returns a tuple with the DeployedSourceMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetDeployedSourceMapOk() (*string, bool) {
	if o == nil || IsNil(o.DeployedSourceMap) {
		return nil, false
	}
	return o.DeployedSourceMap, true
}

// HasDeployedSourceMap returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasDeployedSourceMap() bool {
	if o != nil && !IsNil(o.DeployedSourceMap) {
		return true
	}

	return false
}

// SetDeployedSourceMap gets a reference to the given string and assigns it to the DeployedSourceMap field.
func (o *SolidityContractJsonArtifact) SetDeployedSourceMap(v string) {
	o.DeployedSourceMap = &v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetSourcePath() string {
	if o == nil || IsNil(o.SourcePath) {
		var ret string
		return ret
	}
	return *o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetSourcePathOk() (*string, bool) {
	if o == nil || IsNil(o.SourcePath) {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasSourcePath() bool {
	if o != nil && !IsNil(o.SourcePath) {
		return true
	}

	return false
}

// SetSourcePath gets a reference to the given string and assigns it to the SourcePath field.
func (o *SolidityContractJsonArtifact) SetSourcePath(v string) {
	o.SourcePath = &v
}

// GetCompiler returns the Compiler field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetCompiler() SolidityContractJsonArtifactCompiler {
	if o == nil || IsNil(o.Compiler) {
		var ret SolidityContractJsonArtifactCompiler
		return ret
	}
	return *o.Compiler
}

// GetCompilerOk returns a tuple with the Compiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetCompilerOk() (*SolidityContractJsonArtifactCompiler, bool) {
	if o == nil || IsNil(o.Compiler) {
		return nil, false
	}
	return o.Compiler, true
}

// HasCompiler returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasCompiler() bool {
	if o != nil && !IsNil(o.Compiler) {
		return true
	}

	return false
}

// SetCompiler gets a reference to the given SolidityContractJsonArtifactCompiler and assigns it to the Compiler field.
func (o *SolidityContractJsonArtifact) SetCompiler(v SolidityContractJsonArtifactCompiler) {
	o.Compiler = &v
}

// GetFunctionHashes returns the FunctionHashes field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetFunctionHashes() map[string]interface{} {
	if o == nil || IsNil(o.FunctionHashes) {
		var ret map[string]interface{}
		return ret
	}
	return o.FunctionHashes
}

// GetFunctionHashesOk returns a tuple with the FunctionHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetFunctionHashesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FunctionHashes) {
		return map[string]interface{}{}, false
	}
	return o.FunctionHashes, true
}

// HasFunctionHashes returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasFunctionHashes() bool {
	if o != nil && !IsNil(o.FunctionHashes) {
		return true
	}

	return false
}

// SetFunctionHashes gets a reference to the given map[string]interface{} and assigns it to the FunctionHashes field.
func (o *SolidityContractJsonArtifact) SetFunctionHashes(v map[string]interface{}) {
	o.FunctionHashes = v
}

// GetGasEstimates returns the GasEstimates field value if set, zero value otherwise.
func (o *SolidityContractJsonArtifact) GetGasEstimates() SolidityContractJsonArtifactGasEstimates {
	if o == nil || IsNil(o.GasEstimates) {
		var ret SolidityContractJsonArtifactGasEstimates
		return ret
	}
	return *o.GasEstimates
}

// GetGasEstimatesOk returns a tuple with the GasEstimates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolidityContractJsonArtifact) GetGasEstimatesOk() (*SolidityContractJsonArtifactGasEstimates, bool) {
	if o == nil || IsNil(o.GasEstimates) {
		return nil, false
	}
	return o.GasEstimates, true
}

// HasGasEstimates returns a boolean if a field has been set.
func (o *SolidityContractJsonArtifact) HasGasEstimates() bool {
	if o != nil && !IsNil(o.GasEstimates) {
		return true
	}

	return false
}

// SetGasEstimates gets a reference to the given SolidityContractJsonArtifactGasEstimates and assigns it to the GasEstimates field.
func (o *SolidityContractJsonArtifact) SetGasEstimates(v SolidityContractJsonArtifactGasEstimates) {
	o.GasEstimates = &v
}

func (o SolidityContractJsonArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolidityContractJsonArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractName"] = o.ContractName
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Bytecode) {
		toSerialize["bytecode"] = o.Bytecode
	}
	if !IsNil(o.DeployedBytecode) {
		toSerialize["deployedBytecode"] = o.DeployedBytecode
	}
	if !IsNil(o.SourceMap) {
		toSerialize["sourceMap"] = o.SourceMap
	}
	if !IsNil(o.DeployedSourceMap) {
		toSerialize["deployedSourceMap"] = o.DeployedSourceMap
	}
	if !IsNil(o.SourcePath) {
		toSerialize["sourcePath"] = o.SourcePath
	}
	if !IsNil(o.Compiler) {
		toSerialize["compiler"] = o.Compiler
	}
	if !IsNil(o.FunctionHashes) {
		toSerialize["functionHashes"] = o.FunctionHashes
	}
	if !IsNil(o.GasEstimates) {
		toSerialize["gasEstimates"] = o.GasEstimates
	}
	return toSerialize, nil
}

type NullableSolidityContractJsonArtifact struct {
	value *SolidityContractJsonArtifact
	isSet bool
}

func (v NullableSolidityContractJsonArtifact) Get() *SolidityContractJsonArtifact {
	return v.value
}

func (v *NullableSolidityContractJsonArtifact) Set(val *SolidityContractJsonArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableSolidityContractJsonArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableSolidityContractJsonArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolidityContractJsonArtifact(val *SolidityContractJsonArtifact) *NullableSolidityContractJsonArtifact {
	return &NullableSolidityContractJsonArtifact{value: val, isSet: true}
}

func (v NullableSolidityContractJsonArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolidityContractJsonArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


