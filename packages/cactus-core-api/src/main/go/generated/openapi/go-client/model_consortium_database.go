/*
Hyperledger Core API

Contains/describes the core API types for Cactus. Does not describe actual endpoints on its own as this is left to the implementing plugins who can import and re-use commonly needed type definitions from this specification. One example of said commonly used type definitions would be the types related to consortium management, cactus nodes, ledgers, etc..

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-core-api

import (
	"encoding/json"
)

// checks if the ConsortiumDatabase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsortiumDatabase{}

// ConsortiumDatabase struct for ConsortiumDatabase
type ConsortiumDatabase struct {
	// A collection of Consortium entities. In practice this should only ever contain a single consortium, but we defined it as an array to keep the convention up with the rest of the collections defined in the Consortium data in general. Also, if we ever decide to somehow have some sort of consortium to consortium integration (which does not make much sense in the current frame of mind of the author in the year 2020) then having this as an array will have proven itself to be an excellent long term compatibility/extensibility decision indeed.
	Consortium []Consortium `json:"consortium"`
	// The complete collection of all ledger entities in existence within the consortium.
	Ledger []Ledger `json:"ledger"`
	// The complete collection of all consortium member entities in existence within the consortium.
	ConsortiumMember []ConsortiumMember `json:"consortiumMember"`
	// The complete collection of all cactus nodes entities in existence within the consortium.
	CactusNode []CactusNode `json:"cactusNode"`
	// The complete collection of all plugin instance entities in existence within the consortium.
	PluginInstance []PluginInstance `json:"pluginInstance"`
}

// NewConsortiumDatabase instantiates a new ConsortiumDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsortiumDatabase(consortium []Consortium, ledger []Ledger, consortiumMember []ConsortiumMember, cactusNode []CactusNode, pluginInstance []PluginInstance) *ConsortiumDatabase {
	this := ConsortiumDatabase{}
	this.Consortium = consortium
	this.Ledger = ledger
	this.ConsortiumMember = consortiumMember
	this.CactusNode = cactusNode
	this.PluginInstance = pluginInstance
	return &this
}

// NewConsortiumDatabaseWithDefaults instantiates a new ConsortiumDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsortiumDatabaseWithDefaults() *ConsortiumDatabase {
	this := ConsortiumDatabase{}
	return &this
}

// GetConsortium returns the Consortium field value
func (o *ConsortiumDatabase) GetConsortium() []Consortium {
	if o == nil {
		var ret []Consortium
		return ret
	}

	return o.Consortium
}

// GetConsortiumOk returns a tuple with the Consortium field value
// and a boolean to check if the value has been set.
func (o *ConsortiumDatabase) GetConsortiumOk() ([]Consortium, bool) {
	if o == nil {
		return nil, false
	}
	return o.Consortium, true
}

// SetConsortium sets field value
func (o *ConsortiumDatabase) SetConsortium(v []Consortium) {
	o.Consortium = v
}

// GetLedger returns the Ledger field value
func (o *ConsortiumDatabase) GetLedger() []Ledger {
	if o == nil {
		var ret []Ledger
		return ret
	}

	return o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value
// and a boolean to check if the value has been set.
func (o *ConsortiumDatabase) GetLedgerOk() ([]Ledger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ledger, true
}

// SetLedger sets field value
func (o *ConsortiumDatabase) SetLedger(v []Ledger) {
	o.Ledger = v
}

// GetConsortiumMember returns the ConsortiumMember field value
func (o *ConsortiumDatabase) GetConsortiumMember() []ConsortiumMember {
	if o == nil {
		var ret []ConsortiumMember
		return ret
	}

	return o.ConsortiumMember
}

// GetConsortiumMemberOk returns a tuple with the ConsortiumMember field value
// and a boolean to check if the value has been set.
func (o *ConsortiumDatabase) GetConsortiumMemberOk() ([]ConsortiumMember, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsortiumMember, true
}

// SetConsortiumMember sets field value
func (o *ConsortiumDatabase) SetConsortiumMember(v []ConsortiumMember) {
	o.ConsortiumMember = v
}

// GetCactusNode returns the CactusNode field value
func (o *ConsortiumDatabase) GetCactusNode() []CactusNode {
	if o == nil {
		var ret []CactusNode
		return ret
	}

	return o.CactusNode
}

// GetCactusNodeOk returns a tuple with the CactusNode field value
// and a boolean to check if the value has been set.
func (o *ConsortiumDatabase) GetCactusNodeOk() ([]CactusNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.CactusNode, true
}

// SetCactusNode sets field value
func (o *ConsortiumDatabase) SetCactusNode(v []CactusNode) {
	o.CactusNode = v
}

// GetPluginInstance returns the PluginInstance field value
func (o *ConsortiumDatabase) GetPluginInstance() []PluginInstance {
	if o == nil {
		var ret []PluginInstance
		return ret
	}

	return o.PluginInstance
}

// GetPluginInstanceOk returns a tuple with the PluginInstance field value
// and a boolean to check if the value has been set.
func (o *ConsortiumDatabase) GetPluginInstanceOk() ([]PluginInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginInstance, true
}

// SetPluginInstance sets field value
func (o *ConsortiumDatabase) SetPluginInstance(v []PluginInstance) {
	o.PluginInstance = v
}

func (o ConsortiumDatabase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsortiumDatabase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consortium"] = o.Consortium
	toSerialize["ledger"] = o.Ledger
	toSerialize["consortiumMember"] = o.ConsortiumMember
	toSerialize["cactusNode"] = o.CactusNode
	toSerialize["pluginInstance"] = o.PluginInstance
	return toSerialize, nil
}

type NullableConsortiumDatabase struct {
	value *ConsortiumDatabase
	isSet bool
}

func (v NullableConsortiumDatabase) Get() *ConsortiumDatabase {
	return v.value
}

func (v *NullableConsortiumDatabase) Set(val *ConsortiumDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableConsortiumDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableConsortiumDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsortiumDatabase(val *ConsortiumDatabase) *NullableConsortiumDatabase {
	return &NullableConsortiumDatabase{value: val, isSet: true}
}

func (v NullableConsortiumDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsortiumDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


