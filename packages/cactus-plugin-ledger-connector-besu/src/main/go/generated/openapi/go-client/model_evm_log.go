/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the EvmLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EvmLog{}

// EvmLog struct for EvmLog
type EvmLog struct {
	Address string `json:"address"`
	Data string `json:"data"`
	BlockHash string `json:"blockHash"`
	TransactionHash string `json:"transactionHash"`
	Topics []string `json:"topics"`
	LogIndex float32 `json:"logIndex"`
	TransactionIndex float32 `json:"transactionIndex"`
	BlockNumber float32 `json:"blockNumber"`
}

// NewEvmLog instantiates a new EvmLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvmLog(address string, data string, blockHash string, transactionHash string, topics []string, logIndex float32, transactionIndex float32, blockNumber float32) *EvmLog {
	this := EvmLog{}
	this.Address = address
	this.Data = data
	this.BlockHash = blockHash
	this.TransactionHash = transactionHash
	this.Topics = topics
	this.LogIndex = logIndex
	this.TransactionIndex = transactionIndex
	this.BlockNumber = blockNumber
	return &this
}

// NewEvmLogWithDefaults instantiates a new EvmLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvmLogWithDefaults() *EvmLog {
	this := EvmLog{}
	return &this
}

// GetAddress returns the Address field value
func (o *EvmLog) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *EvmLog) SetAddress(v string) {
	o.Address = v
}

// GetData returns the Data field value
func (o *EvmLog) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EvmLog) SetData(v string) {
	o.Data = v
}

// GetBlockHash returns the BlockHash field value
func (o *EvmLog) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *EvmLog) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *EvmLog) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *EvmLog) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetTopics returns the Topics field value
func (o *EvmLog) GetTopics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetTopicsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topics, true
}

// SetTopics sets field value
func (o *EvmLog) SetTopics(v []string) {
	o.Topics = v
}

// GetLogIndex returns the LogIndex field value
func (o *EvmLog) GetLogIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetLogIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogIndex, true
}

// SetLogIndex sets field value
func (o *EvmLog) SetLogIndex(v float32) {
	o.LogIndex = v
}

// GetTransactionIndex returns the TransactionIndex field value
func (o *EvmLog) GetTransactionIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransactionIndex
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetTransactionIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionIndex, true
}

// SetTransactionIndex sets field value
func (o *EvmLog) SetTransactionIndex(v float32) {
	o.TransactionIndex = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *EvmLog) GetBlockNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *EvmLog) GetBlockNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *EvmLog) SetBlockNumber(v float32) {
	o.BlockNumber = v
}

func (o EvmLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EvmLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["data"] = o.Data
	toSerialize["blockHash"] = o.BlockHash
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["topics"] = o.Topics
	toSerialize["logIndex"] = o.LogIndex
	toSerialize["transactionIndex"] = o.TransactionIndex
	toSerialize["blockNumber"] = o.BlockNumber
	return toSerialize, nil
}

type NullableEvmLog struct {
	value *EvmLog
	isSet bool
}

func (v NullableEvmLog) Get() *EvmLog {
	return v.value
}

func (v *NullableEvmLog) Set(val *EvmLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEvmLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEvmLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvmLog(val *EvmLog) *NullableEvmLog {
	return &NullableEvmLog{value: val, isSet: true}
}

func (v NullableEvmLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvmLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


