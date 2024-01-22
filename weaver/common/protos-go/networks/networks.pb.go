// Copyright IBM Corp. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: networks/networks.proto

package networks

import (
	common "github.com/hyperledger/cacti/weaver/common/protos-go/v2/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DbName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DbName) Reset() {
	*x = DbName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbName) ProtoMessage() {}

func (x *DbName) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbName.ProtoReflect.Descriptor instead.
func (*DbName) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{0}
}

func (x *DbName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RelayDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pairs map[string]string `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RelayDatabase) Reset() {
	*x = RelayDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayDatabase) ProtoMessage() {}

func (x *RelayDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayDatabase.ProtoReflect.Descriptor instead.
func (*RelayDatabase) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{1}
}

func (x *RelayDatabase) GetPairs() map[string]string {
	if x != nil {
		return x.Pairs
	}
	return nil
}

type GetStateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *GetStateMessage) Reset() {
	*x = GetStateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateMessage) ProtoMessage() {}

func (x *GetStateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateMessage.ProtoReflect.Descriptor instead.
func (*GetStateMessage) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{2}
}

func (x *GetStateMessage) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type NetworkQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy             []string `protobuf:"bytes,1,rep,name=policy,proto3" json:"policy,omitempty"`
	Address            string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	RequestingRelay    string   `protobuf:"bytes,3,opt,name=requesting_relay,json=requestingRelay,proto3" json:"requesting_relay,omitempty"`
	RequestingNetwork  string   `protobuf:"bytes,4,opt,name=requesting_network,json=requestingNetwork,proto3" json:"requesting_network,omitempty"`
	Certificate        string   `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	RequestorSignature string   `protobuf:"bytes,6,opt,name=requestor_signature,json=requestorSignature,proto3" json:"requestor_signature,omitempty"`
	Nonce              string   `protobuf:"bytes,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RequestingOrg      string   `protobuf:"bytes,8,opt,name=requesting_org,json=requestingOrg,proto3" json:"requesting_org,omitempty"`
	Confidential       bool     `protobuf:"varint,9,opt,name=confidential,proto3" json:"confidential,omitempty"`
}

func (x *NetworkQuery) Reset() {
	*x = NetworkQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkQuery) ProtoMessage() {}

func (x *NetworkQuery) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkQuery.ProtoReflect.Descriptor instead.
func (*NetworkQuery) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkQuery) GetPolicy() []string {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *NetworkQuery) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NetworkQuery) GetRequestingRelay() string {
	if x != nil {
		return x.RequestingRelay
	}
	return ""
}

func (x *NetworkQuery) GetRequestingNetwork() string {
	if x != nil {
		return x.RequestingNetwork
	}
	return ""
}

func (x *NetworkQuery) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *NetworkQuery) GetRequestorSignature() string {
	if x != nil {
		return x.RequestorSignature
	}
	return ""
}

func (x *NetworkQuery) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *NetworkQuery) GetRequestingOrg() string {
	if x != nil {
		return x.RequestingOrg
	}
	return ""
}

func (x *NetworkQuery) GetConfidential() bool {
	if x != nil {
		return x.Confidential
	}
	return false
}

// Below message is used for network/client to dest-relay communication
type NetworkEventSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventMatcher         *common.EventMatcher     `protobuf:"bytes,1,opt,name=event_matcher,json=eventMatcher,proto3" json:"event_matcher,omitempty"`
	Query                *NetworkQuery            `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	EventPublicationSpec *common.EventPublication `protobuf:"bytes,3,opt,name=event_publication_spec,json=eventPublicationSpec,proto3" json:"event_publication_spec,omitempty"`
}

func (x *NetworkEventSubscription) Reset() {
	*x = NetworkEventSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkEventSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkEventSubscription) ProtoMessage() {}

func (x *NetworkEventSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkEventSubscription.ProtoReflect.Descriptor instead.
func (*NetworkEventSubscription) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkEventSubscription) GetEventMatcher() *common.EventMatcher {
	if x != nil {
		return x.EventMatcher
	}
	return nil
}

func (x *NetworkEventSubscription) GetQuery() *NetworkQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *NetworkEventSubscription) GetEventPublicationSpec() *common.EventPublication {
	if x != nil {
		return x.EventPublicationSpec
	}
	return nil
}

type NetworkEventUnsubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request   *NetworkEventSubscription `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	RequestId string                    `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *NetworkEventUnsubscription) Reset() {
	*x = NetworkEventUnsubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkEventUnsubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkEventUnsubscription) ProtoMessage() {}

func (x *NetworkEventUnsubscription) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkEventUnsubscription.ProtoReflect.Descriptor instead.
func (*NetworkEventUnsubscription) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkEventUnsubscription) GetRequest() *NetworkEventSubscription {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *NetworkEventUnsubscription) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type NetworkAssetTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policy             []string `protobuf:"bytes,1,rep,name=policy,proto3" json:"policy,omitempty"`
	Address            string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	RequestingRelay    string   `protobuf:"bytes,3,opt,name=requesting_relay,json=requestingRelay,proto3" json:"requesting_relay,omitempty"`
	RequestingNetwork  string   `protobuf:"bytes,4,opt,name=requesting_network,json=requestingNetwork,proto3" json:"requesting_network,omitempty"`
	Certificate        string   `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	RequestorSignature string   `protobuf:"bytes,6,opt,name=requestor_signature,json=requestorSignature,proto3" json:"requestor_signature,omitempty"`
	Nonce              string   `protobuf:"bytes,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RequestingOrg      string   `protobuf:"bytes,8,opt,name=requesting_org,json=requestingOrg,proto3" json:"requesting_org,omitempty"`
	Confidential       bool     `protobuf:"varint,9,opt,name=confidential,proto3" json:"confidential,omitempty"`
}

func (x *NetworkAssetTransfer) Reset() {
	*x = NetworkAssetTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networks_networks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkAssetTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkAssetTransfer) ProtoMessage() {}

func (x *NetworkAssetTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_networks_networks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkAssetTransfer.ProtoReflect.Descriptor instead.
func (*NetworkAssetTransfer) Descriptor() ([]byte, []int) {
	return file_networks_networks_proto_rawDescGZIP(), []int{6}
}

func (x *NetworkAssetTransfer) GetPolicy() []string {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *NetworkAssetTransfer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NetworkAssetTransfer) GetRequestingRelay() string {
	if x != nil {
		return x.RequestingRelay
	}
	return ""
}

func (x *NetworkAssetTransfer) GetRequestingNetwork() string {
	if x != nil {
		return x.RequestingNetwork
	}
	return ""
}

func (x *NetworkAssetTransfer) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *NetworkAssetTransfer) GetRequestorSignature() string {
	if x != nil {
		return x.RequestorSignature
	}
	return ""
}

func (x *NetworkAssetTransfer) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *NetworkAssetTransfer) GetRequestingOrg() string {
	if x != nil {
		return x.RequestingOrg
	}
	return ""
}

func (x *NetworkAssetTransfer) GetConfidential() bool {
	if x != nil {
		return x.Confidential
	}
	return false
}

var File_networks_networks_proto protoreflect.FileDescriptor

var file_networks_networks_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x1a, 0x10, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x06, 0x44, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x69, 0x72, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xce, 0x02, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x72, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x4f, 0x72, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0xea, 0x01, 0x0a, 0x18, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x55, 0x0a,
	0x16, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x82, 0x01, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xd6, 0x02, 0x0a, 0x14, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x72, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x32, 0xa7, 0x05, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x42,
	0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00,
	0x12, 0x50, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x19, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x44, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x20,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b,
	0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61,
	0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x00, 0x12, 0x54, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61,
	0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0x00, 0x42, 0x79, 0x0a, 0x35,
	0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x63, 0x61, 0x63, 0x74, 0x69, 0x2e, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x61,
	0x63, 0x74, 0x69, 0x2f, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_networks_networks_proto_rawDescOnce sync.Once
	file_networks_networks_proto_rawDescData = file_networks_networks_proto_rawDesc
)

func file_networks_networks_proto_rawDescGZIP() []byte {
	file_networks_networks_proto_rawDescOnce.Do(func() {
		file_networks_networks_proto_rawDescData = protoimpl.X.CompressGZIP(file_networks_networks_proto_rawDescData)
	})
	return file_networks_networks_proto_rawDescData
}

var file_networks_networks_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_networks_networks_proto_goTypes = []interface{}{
	(*DbName)(nil),                        // 0: networks.networks.DbName
	(*RelayDatabase)(nil),                 // 1: networks.networks.RelayDatabase
	(*GetStateMessage)(nil),               // 2: networks.networks.GetStateMessage
	(*NetworkQuery)(nil),                  // 3: networks.networks.NetworkQuery
	(*NetworkEventSubscription)(nil),      // 4: networks.networks.NetworkEventSubscription
	(*NetworkEventUnsubscription)(nil),    // 5: networks.networks.NetworkEventUnsubscription
	(*NetworkAssetTransfer)(nil),          // 6: networks.networks.NetworkAssetTransfer
	nil,                                   // 7: networks.networks.RelayDatabase.PairsEntry
	(*common.EventMatcher)(nil),           // 8: common.events.EventMatcher
	(*common.EventPublication)(nil),       // 9: common.events.EventPublication
	(*common.Ack)(nil),                    // 10: common.ack.Ack
	(*common.RequestState)(nil),           // 11: common.state.RequestState
	(*common.EventSubscriptionState)(nil), // 12: common.events.EventSubscriptionState
	(*common.EventStates)(nil),            // 13: common.events.EventStates
}
var file_networks_networks_proto_depIdxs = []int32{
	7,  // 0: networks.networks.RelayDatabase.pairs:type_name -> networks.networks.RelayDatabase.PairsEntry
	8,  // 1: networks.networks.NetworkEventSubscription.event_matcher:type_name -> common.events.EventMatcher
	3,  // 2: networks.networks.NetworkEventSubscription.query:type_name -> networks.networks.NetworkQuery
	9,  // 3: networks.networks.NetworkEventSubscription.event_publication_spec:type_name -> common.events.EventPublication
	4,  // 4: networks.networks.NetworkEventUnsubscription.request:type_name -> networks.networks.NetworkEventSubscription
	3,  // 5: networks.networks.Network.RequestState:input_type -> networks.networks.NetworkQuery
	2,  // 6: networks.networks.Network.GetState:input_type -> networks.networks.GetStateMessage
	0,  // 7: networks.networks.Network.RequestDatabase:input_type -> networks.networks.DbName
	6,  // 8: networks.networks.Network.RequestAssetTransfer:input_type -> networks.networks.NetworkAssetTransfer
	4,  // 9: networks.networks.Network.SubscribeEvent:input_type -> networks.networks.NetworkEventSubscription
	2,  // 10: networks.networks.Network.GetEventSubscriptionState:input_type -> networks.networks.GetStateMessage
	5,  // 11: networks.networks.Network.UnsubscribeEvent:input_type -> networks.networks.NetworkEventUnsubscription
	2,  // 12: networks.networks.Network.GetEventStates:input_type -> networks.networks.GetStateMessage
	10, // 13: networks.networks.Network.RequestState:output_type -> common.ack.Ack
	11, // 14: networks.networks.Network.GetState:output_type -> common.state.RequestState
	1,  // 15: networks.networks.Network.RequestDatabase:output_type -> networks.networks.RelayDatabase
	10, // 16: networks.networks.Network.RequestAssetTransfer:output_type -> common.ack.Ack
	10, // 17: networks.networks.Network.SubscribeEvent:output_type -> common.ack.Ack
	12, // 18: networks.networks.Network.GetEventSubscriptionState:output_type -> common.events.EventSubscriptionState
	10, // 19: networks.networks.Network.UnsubscribeEvent:output_type -> common.ack.Ack
	13, // 20: networks.networks.Network.GetEventStates:output_type -> common.events.EventStates
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_networks_networks_proto_init() }
func file_networks_networks_proto_init() {
	if File_networks_networks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networks_networks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbName); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayDatabase); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkEventSubscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkEventUnsubscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_networks_networks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkAssetTransfer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_networks_networks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_networks_networks_proto_goTypes,
		DependencyIndexes: file_networks_networks_proto_depIdxs,
		MessageInfos:      file_networks_networks_proto_msgTypes,
	}.Build()
	File_networks_networks_proto = out.File
	file_networks_networks_proto_rawDesc = nil
	file_networks_networks_proto_goTypes = nil
	file_networks_networks_proto_depIdxs = nil
}
