// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: third_party/tink/proto/jwt_ecdsa.proto

package jwt_ecdsa_go_proto

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

type JwtEcdsaAlgorithm int32

const (
	JwtEcdsaAlgorithm_ES_UNKNOWN JwtEcdsaAlgorithm = 0
	JwtEcdsaAlgorithm_ES256      JwtEcdsaAlgorithm = 1
	JwtEcdsaAlgorithm_ES384      JwtEcdsaAlgorithm = 2
	JwtEcdsaAlgorithm_ES512      JwtEcdsaAlgorithm = 3
)

// Enum value maps for JwtEcdsaAlgorithm.
var (
	JwtEcdsaAlgorithm_name = map[int32]string{
		0: "ES_UNKNOWN",
		1: "ES256",
		2: "ES384",
		3: "ES512",
	}
	JwtEcdsaAlgorithm_value = map[string]int32{
		"ES_UNKNOWN": 0,
		"ES256":      1,
		"ES384":      2,
		"ES512":      3,
	}
)

func (x JwtEcdsaAlgorithm) Enum() *JwtEcdsaAlgorithm {
	p := new(JwtEcdsaAlgorithm)
	*p = x
	return p
}

func (x JwtEcdsaAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JwtEcdsaAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_jwt_ecdsa_proto_enumTypes[0].Descriptor()
}

func (JwtEcdsaAlgorithm) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_jwt_ecdsa_proto_enumTypes[0]
}

func (x JwtEcdsaAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JwtEcdsaAlgorithm.Descriptor instead.
func (JwtEcdsaAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP(), []int{0}
}

// key_type: type.googleapis.com/google.crypto.tink.JwtEcdsaPublicKey
type JwtEcdsaPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32                       `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Algorithm JwtEcdsaAlgorithm            `protobuf:"varint,2,opt,name=algorithm,proto3,enum=google.crypto.tink.JwtEcdsaAlgorithm" json:"algorithm,omitempty"`
	X         []byte                       `protobuf:"bytes,3,opt,name=x,proto3" json:"x,omitempty"`
	Y         []byte                       `protobuf:"bytes,4,opt,name=y,proto3" json:"y,omitempty"`
	CustomKid *JwtEcdsaPublicKey_CustomKid `protobuf:"bytes,5,opt,name=custom_kid,json=customKid,proto3" json:"custom_kid,omitempty"`
}

func (x *JwtEcdsaPublicKey) Reset() {
	*x = JwtEcdsaPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtEcdsaPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtEcdsaPublicKey) ProtoMessage() {}

func (x *JwtEcdsaPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtEcdsaPublicKey.ProtoReflect.Descriptor instead.
func (*JwtEcdsaPublicKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP(), []int{0}
}

func (x *JwtEcdsaPublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *JwtEcdsaPublicKey) GetAlgorithm() JwtEcdsaAlgorithm {
	if x != nil {
		return x.Algorithm
	}
	return JwtEcdsaAlgorithm_ES_UNKNOWN
}

func (x *JwtEcdsaPublicKey) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *JwtEcdsaPublicKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *JwtEcdsaPublicKey) GetCustomKid() *JwtEcdsaPublicKey_CustomKid {
	if x != nil {
		return x.CustomKid
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.JwtEcdsaPrivateKey
type JwtEcdsaPrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32             `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PublicKey *JwtEcdsaPublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Unsigned big integer in bigendian representation.
	KeyValue []byte `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *JwtEcdsaPrivateKey) Reset() {
	*x = JwtEcdsaPrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtEcdsaPrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtEcdsaPrivateKey) ProtoMessage() {}

func (x *JwtEcdsaPrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtEcdsaPrivateKey.ProtoReflect.Descriptor instead.
func (*JwtEcdsaPrivateKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP(), []int{1}
}

func (x *JwtEcdsaPrivateKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *JwtEcdsaPrivateKey) GetPublicKey() *JwtEcdsaPublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *JwtEcdsaPrivateKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

type JwtEcdsaKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32            `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Algorithm JwtEcdsaAlgorithm `protobuf:"varint,2,opt,name=algorithm,proto3,enum=google.crypto.tink.JwtEcdsaAlgorithm" json:"algorithm,omitempty"`
}

func (x *JwtEcdsaKeyFormat) Reset() {
	*x = JwtEcdsaKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtEcdsaKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtEcdsaKeyFormat) ProtoMessage() {}

func (x *JwtEcdsaKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtEcdsaKeyFormat.ProtoReflect.Descriptor instead.
func (*JwtEcdsaKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP(), []int{2}
}

func (x *JwtEcdsaKeyFormat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *JwtEcdsaKeyFormat) GetAlgorithm() JwtEcdsaAlgorithm {
	if x != nil {
		return x.Algorithm
	}
	return JwtEcdsaAlgorithm_ES_UNKNOWN
}

// Optional, custom kid header value to be used with "RAW" keys.
// "TINK" keys with this value set will be rejected.
type JwtEcdsaPublicKey_CustomKid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JwtEcdsaPublicKey_CustomKid) Reset() {
	*x = JwtEcdsaPublicKey_CustomKid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtEcdsaPublicKey_CustomKid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtEcdsaPublicKey_CustomKid) ProtoMessage() {}

func (x *JwtEcdsaPublicKey_CustomKid) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtEcdsaPublicKey_CustomKid.ProtoReflect.Descriptor instead.
func (*JwtEcdsaPublicKey_CustomKid) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP(), []int{0, 0}
}

func (x *JwtEcdsaPublicKey_CustomKid) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_third_party_tink_proto_jwt_ecdsa_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_jwt_ecdsa_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x65, 0x63, 0x64,
	0x73, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x22, 0x81, 0x02, 0x0a,
	0x11, 0x4a, 0x77, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4a, 0x77, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x78, 0x12,
	0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79, 0x12, 0x4e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4a, 0x77, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b,
	0x69, 0x64, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b, 0x69, 0x64, 0x1a, 0x21, 0x0a,
	0x09, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4b, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x91, 0x01, 0x0a, 0x12, 0x4a, 0x77, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4a, 0x77, 0x74, 0x45, 0x63,
	0x64, 0x73, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x72, 0x0a, 0x11, 0x4a, 0x77, 0x74, 0x45, 0x63, 0x64, 0x73, 0x61,
	0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x4a, 0x77, 0x74, 0x45,
	0x63, 0x64, 0x73, 0x61, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2a, 0x44, 0x0a, 0x11, 0x4a, 0x77, 0x74, 0x45,
	0x63, 0x64, 0x73, 0x61, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0e, 0x0a,
	0x0a, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x53, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x53, 0x33, 0x38,
	0x34, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x53, 0x35, 0x31, 0x32, 0x10, 0x03, 0x42, 0x51,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a,
	0x77, 0x74, 0x5f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_jwt_ecdsa_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_jwt_ecdsa_proto_rawDescData = file_third_party_tink_proto_jwt_ecdsa_proto_rawDesc
)

func file_third_party_tink_proto_jwt_ecdsa_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_jwt_ecdsa_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_jwt_ecdsa_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_jwt_ecdsa_proto_rawDescData)
	})
	return file_third_party_tink_proto_jwt_ecdsa_proto_rawDescData
}

var file_third_party_tink_proto_jwt_ecdsa_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_third_party_tink_proto_jwt_ecdsa_proto_goTypes = []interface{}{
	(JwtEcdsaAlgorithm)(0),              // 0: google.crypto.tink.JwtEcdsaAlgorithm
	(*JwtEcdsaPublicKey)(nil),           // 1: google.crypto.tink.JwtEcdsaPublicKey
	(*JwtEcdsaPrivateKey)(nil),          // 2: google.crypto.tink.JwtEcdsaPrivateKey
	(*JwtEcdsaKeyFormat)(nil),           // 3: google.crypto.tink.JwtEcdsaKeyFormat
	(*JwtEcdsaPublicKey_CustomKid)(nil), // 4: google.crypto.tink.JwtEcdsaPublicKey.CustomKid
}
var file_third_party_tink_proto_jwt_ecdsa_proto_depIdxs = []int32{
	0, // 0: google.crypto.tink.JwtEcdsaPublicKey.algorithm:type_name -> google.crypto.tink.JwtEcdsaAlgorithm
	4, // 1: google.crypto.tink.JwtEcdsaPublicKey.custom_kid:type_name -> google.crypto.tink.JwtEcdsaPublicKey.CustomKid
	1, // 2: google.crypto.tink.JwtEcdsaPrivateKey.public_key:type_name -> google.crypto.tink.JwtEcdsaPublicKey
	0, // 3: google.crypto.tink.JwtEcdsaKeyFormat.algorithm:type_name -> google.crypto.tink.JwtEcdsaAlgorithm
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_jwt_ecdsa_proto_init() }
func file_third_party_tink_proto_jwt_ecdsa_proto_init() {
	if File_third_party_tink_proto_jwt_ecdsa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtEcdsaPublicKey); i {
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
		file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtEcdsaPrivateKey); i {
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
		file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtEcdsaKeyFormat); i {
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
		file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtEcdsaPublicKey_CustomKid); i {
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
			RawDescriptor: file_third_party_tink_proto_jwt_ecdsa_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_jwt_ecdsa_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_jwt_ecdsa_proto_depIdxs,
		EnumInfos:         file_third_party_tink_proto_jwt_ecdsa_proto_enumTypes,
		MessageInfos:      file_third_party_tink_proto_jwt_ecdsa_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_jwt_ecdsa_proto = out.File
	file_third_party_tink_proto_jwt_ecdsa_proto_rawDesc = nil
	file_third_party_tink_proto_jwt_ecdsa_proto_goTypes = nil
	file_third_party_tink_proto_jwt_ecdsa_proto_depIdxs = nil
}
