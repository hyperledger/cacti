// Copyright IBM Corp. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: common/access_control.proto

package common

import (
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

// AccessControlPolicy specifies a set of data that can be accessed by some
// SecurityGroup
type AccessControlPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityDomain string  `protobuf:"bytes,1,opt,name=securityDomain,proto3" json:"securityDomain,omitempty"`
	Rules          []*Rule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *AccessControlPolicy) Reset() {
	*x = AccessControlPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_access_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessControlPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessControlPolicy) ProtoMessage() {}

func (x *AccessControlPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_common_access_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessControlPolicy.ProtoReflect.Descriptor instead.
func (*AccessControlPolicy) Descriptor() ([]byte, []int) {
	return file_common_access_control_proto_rawDescGZIP(), []int{0}
}

func (x *AccessControlPolicy) GetSecurityDomain() string {
	if x != nil {
		return x.SecurityDomain
	}
	return ""
}

func (x *AccessControlPolicy) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// Rule represents a single data access rule for the AccessControlPolicy
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principal     string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	PrincipalType string `protobuf:"bytes,2,opt,name=principalType,proto3" json:"principalType,omitempty"`
	Resource      string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Read          bool   `protobuf:"varint,4,opt,name=read,proto3" json:"read,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_access_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_common_access_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_common_access_control_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *Rule) GetPrincipalType() string {
	if x != nil {
		return x.PrincipalType
	}
	return ""
}

func (x *Rule) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Rule) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

var File_common_access_control_proto protoreflect.FileDescriptor

var file_common_access_control_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x22, 0x70, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x7a, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65,
	0x61, 0x64, 0x42, 0x7b, 0x0a, 0x39, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2e, 0x77, 0x65, 0x61, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2f, 0x77, 0x65,
	0x61, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_access_control_proto_rawDescOnce sync.Once
	file_common_access_control_proto_rawDescData = file_common_access_control_proto_rawDesc
)

func file_common_access_control_proto_rawDescGZIP() []byte {
	file_common_access_control_proto_rawDescOnce.Do(func() {
		file_common_access_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_access_control_proto_rawDescData)
	})
	return file_common_access_control_proto_rawDescData
}

var file_common_access_control_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_access_control_proto_goTypes = []interface{}{
	(*AccessControlPolicy)(nil), // 0: common.access_control.AccessControlPolicy
	(*Rule)(nil),                // 1: common.access_control.Rule
}
var file_common_access_control_proto_depIdxs = []int32{
	1, // 0: common.access_control.AccessControlPolicy.rules:type_name -> common.access_control.Rule
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_access_control_proto_init() }
func file_common_access_control_proto_init() {
	if File_common_access_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_access_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessControlPolicy); i {
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
		file_common_access_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
			RawDescriptor: file_common_access_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_access_control_proto_goTypes,
		DependencyIndexes: file_common_access_control_proto_depIdxs,
		MessageInfos:      file_common_access_control_proto_msgTypes,
	}.Build()
	File_common_access_control_proto = out.File
	file_common_access_control_proto_rawDesc = nil
	file_common_access_control_proto_goTypes = nil
	file_common_access_control_proto_depIdxs = nil
}
