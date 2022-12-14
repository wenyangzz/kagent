// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: isvc.proto

package isvc

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

type GetIsvcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelSpecGet *ModelSpec `protobuf:"bytes,1,opt,name=model_spec_get,json=modelSpecGet,proto3" json:"model_spec_get,omitempty"`
}

func (x *GetIsvcRequest) Reset() {
	*x = GetIsvcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIsvcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIsvcRequest) ProtoMessage() {}

func (x *GetIsvcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIsvcRequest.ProtoReflect.Descriptor instead.
func (*GetIsvcRequest) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{0}
}

func (x *GetIsvcRequest) GetModelSpecGet() *ModelSpec {
	if x != nil {
		return x.ModelSpecGet
	}
	return nil
}

type GetIsvcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsvcOutput *IsvcSpec `protobuf:"bytes,1,opt,name=isvc_output,json=isvcOutput,proto3" json:"isvc_output,omitempty"`
}

func (x *GetIsvcResponse) Reset() {
	*x = GetIsvcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIsvcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIsvcResponse) ProtoMessage() {}

func (x *GetIsvcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIsvcResponse.ProtoReflect.Descriptor instead.
func (*GetIsvcResponse) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{1}
}

func (x *GetIsvcResponse) GetIsvcOutput() *IsvcSpec {
	if x != nil {
		return x.IsvcOutput
	}
	return nil
}

type CreateIsvcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelSpecCreate *ModelSpec `protobuf:"bytes,1,opt,name=model_spec_create,json=modelSpecCreate,proto3" json:"model_spec_create,omitempty"`
}

func (x *CreateIsvcRequest) Reset() {
	*x = CreateIsvcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIsvcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIsvcRequest) ProtoMessage() {}

func (x *CreateIsvcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIsvcRequest.ProtoReflect.Descriptor instead.
func (*CreateIsvcRequest) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{2}
}

func (x *CreateIsvcRequest) GetModelSpecCreate() *ModelSpec {
	if x != nil {
		return x.ModelSpecCreate
	}
	return nil
}

type DeleteIsvcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelSpecDelete *ModelSpec `protobuf:"bytes,1,opt,name=model_spec_delete,json=modelSpecDelete,proto3" json:"model_spec_delete,omitempty"`
}

func (x *DeleteIsvcRequest) Reset() {
	*x = DeleteIsvcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIsvcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIsvcRequest) ProtoMessage() {}

func (x *DeleteIsvcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIsvcRequest.ProtoReflect.Descriptor instead.
func (*DeleteIsvcRequest) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteIsvcRequest) GetModelSpecDelete() *ModelSpec {
	if x != nil {
		return x.ModelSpecDelete
	}
	return nil
}

type ListIsvcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsvcList []*IsvcSpec `protobuf:"bytes,1,rep,name=isvc_list,json=isvcList,proto3" json:"isvc_list,omitempty"`
}

func (x *ListIsvcResponse) Reset() {
	*x = ListIsvcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIsvcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIsvcResponse) ProtoMessage() {}

func (x *ListIsvcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIsvcResponse.ProtoReflect.Descriptor instead.
func (*ListIsvcResponse) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{4}
}

func (x *ListIsvcResponse) GetIsvcList() []*IsvcSpec {
	if x != nil {
		return x.IsvcList
	}
	return nil
}

type ModelSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName        string `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	ModelVersion     string `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	ModelFramework   []byte `protobuf:"bytes,3,opt,name=model_framework,json=modelFramework,proto3" json:"model_framework,omitempty"`
	ModelContent     []byte `protobuf:"bytes,4,opt,name=model_content,json=modelContent,proto3" json:"model_content,omitempty"`
	ModelPreparation []byte `protobuf:"bytes,5,opt,name=model_preparation,json=modelPreparation,proto3" json:"model_preparation,omitempty"`
}

func (x *ModelSpec) Reset() {
	*x = ModelSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelSpec) ProtoMessage() {}

func (x *ModelSpec) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelSpec.ProtoReflect.Descriptor instead.
func (*ModelSpec) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{5}
}

func (x *ModelSpec) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ModelSpec) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *ModelSpec) GetModelFramework() []byte {
	if x != nil {
		return x.ModelFramework
	}
	return nil
}

func (x *ModelSpec) GetModelContent() []byte {
	if x != nil {
		return x.ModelContent
	}
	return nil
}

func (x *ModelSpec) GetModelPreparation() []byte {
	if x != nil {
		return x.ModelPreparation
	}
	return nil
}

type IsvcSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsvcName   string `protobuf:"bytes,1,opt,name=isvc_name,json=isvcName,proto3" json:"isvc_name,omitempty"`
	IsvcStatus string `protobuf:"bytes,2,opt,name=isvc_status,json=isvcStatus,proto3" json:"isvc_status,omitempty"`
	IsvcUrl    string `protobuf:"bytes,3,opt,name=isvc_url,json=isvcUrl,proto3" json:"isvc_url,omitempty"`
}

func (x *IsvcSpec) Reset() {
	*x = IsvcSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsvcSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsvcSpec) ProtoMessage() {}

func (x *IsvcSpec) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsvcSpec.ProtoReflect.Descriptor instead.
func (*IsvcSpec) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{6}
}

func (x *IsvcSpec) GetIsvcName() string {
	if x != nil {
		return x.IsvcName
	}
	return ""
}

func (x *IsvcSpec) GetIsvcStatus() string {
	if x != nil {
		return x.IsvcStatus
	}
	return ""
}

func (x *IsvcSpec) GetIsvcUrl() string {
	if x != nil {
		return x.IsvcUrl
	}
	return ""
}

type IsvcNull struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IsvcNull) Reset() {
	*x = IsvcNull{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isvc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsvcNull) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsvcNull) ProtoMessage() {}

func (x *IsvcNull) ProtoReflect() protoreflect.Message {
	mi := &file_isvc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsvcNull.ProtoReflect.Descriptor instead.
func (*IsvcNull) Descriptor() ([]byte, []int) {
	return file_isvc_proto_rawDescGZIP(), []int{7}
}

var File_isvc_proto protoreflect.FileDescriptor

var file_isvc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x73,
	0x76, 0x63, 0x22, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69,
	0x73, 0x76, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0c, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x47, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x0b, 0x69, 0x73, 0x76, 0x63, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x49, 0x73, 0x76, 0x63, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0a, 0x69, 0x73, 0x76, 0x63, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x50, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x22, 0x50, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x73, 0x76, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x22, 0x3f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x76, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x69, 0x73, 0x76, 0x63, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x69, 0x73, 0x76,
	0x63, 0x2e, 0x49, 0x73, 0x76, 0x63, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x69, 0x73, 0x76, 0x63,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x63, 0x0a, 0x08, 0x49, 0x73, 0x76, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x76, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x73, 0x76, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x76, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x73, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x76, 0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x73, 0x76, 0x63, 0x55, 0x72, 0x6c, 0x22, 0x0a, 0x0a, 0x08, 0x49, 0x73, 0x76, 0x63, 0x4e, 0x75,
	0x6c, 0x6c, 0x32, 0xef, 0x01, 0x0a, 0x0b, 0x49, 0x73, 0x76, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x73, 0x76, 0x63, 0x12, 0x14, 0x2e,
	0x69, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73,
	0x76, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x76, 0x63, 0x12, 0x17, 0x2e, 0x69, 0x73, 0x76,
	0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x49, 0x73, 0x76, 0x63, 0x4e,
	0x75, 0x6c, 0x6c, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x73, 0x76, 0x63, 0x12, 0x17, 0x2e, 0x69, 0x73, 0x76, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x69,
	0x73, 0x76, 0x63, 0x2e, 0x49, 0x73, 0x76, 0x63, 0x4e, 0x75, 0x6c, 0x6c, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x76, 0x63, 0x12, 0x0e, 0x2e, 0x69, 0x73, 0x76,
	0x63, 0x2e, 0x49, 0x73, 0x76, 0x63, 0x4e, 0x75, 0x6c, 0x6c, 0x1a, 0x16, 0x2e, 0x69, 0x73, 0x76,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x73, 0x76, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x72, 0x64, 0x2f, 0x69, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_isvc_proto_rawDescOnce sync.Once
	file_isvc_proto_rawDescData = file_isvc_proto_rawDesc
)

func file_isvc_proto_rawDescGZIP() []byte {
	file_isvc_proto_rawDescOnce.Do(func() {
		file_isvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_isvc_proto_rawDescData)
	})
	return file_isvc_proto_rawDescData
}

var file_isvc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_isvc_proto_goTypes = []interface{}{
	(*GetIsvcRequest)(nil),    // 0: isvc.GetIsvcRequest
	(*GetIsvcResponse)(nil),   // 1: isvc.GetIsvcResponse
	(*CreateIsvcRequest)(nil), // 2: isvc.CreateIsvcRequest
	(*DeleteIsvcRequest)(nil), // 3: isvc.DeleteIsvcRequest
	(*ListIsvcResponse)(nil),  // 4: isvc.ListIsvcResponse
	(*ModelSpec)(nil),         // 5: isvc.ModelSpec
	(*IsvcSpec)(nil),          // 6: isvc.IsvcSpec
	(*IsvcNull)(nil),          // 7: isvc.IsvcNull
}
var file_isvc_proto_depIdxs = []int32{
	5, // 0: isvc.GetIsvcRequest.model_spec_get:type_name -> isvc.ModelSpec
	6, // 1: isvc.GetIsvcResponse.isvc_output:type_name -> isvc.IsvcSpec
	5, // 2: isvc.CreateIsvcRequest.model_spec_create:type_name -> isvc.ModelSpec
	5, // 3: isvc.DeleteIsvcRequest.model_spec_delete:type_name -> isvc.ModelSpec
	6, // 4: isvc.ListIsvcResponse.isvc_list:type_name -> isvc.IsvcSpec
	0, // 5: isvc.IsvcService.GetIsvc:input_type -> isvc.GetIsvcRequest
	2, // 6: isvc.IsvcService.CreateIsvc:input_type -> isvc.CreateIsvcRequest
	3, // 7: isvc.IsvcService.DeleteIsvc:input_type -> isvc.DeleteIsvcRequest
	7, // 8: isvc.IsvcService.ListIsvc:input_type -> isvc.IsvcNull
	1, // 9: isvc.IsvcService.GetIsvc:output_type -> isvc.GetIsvcResponse
	7, // 10: isvc.IsvcService.CreateIsvc:output_type -> isvc.IsvcNull
	7, // 11: isvc.IsvcService.DeleteIsvc:output_type -> isvc.IsvcNull
	4, // 12: isvc.IsvcService.ListIsvc:output_type -> isvc.ListIsvcResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_isvc_proto_init() }
func file_isvc_proto_init() {
	if File_isvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIsvcRequest); i {
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
		file_isvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIsvcResponse); i {
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
		file_isvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIsvcRequest); i {
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
		file_isvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIsvcRequest); i {
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
		file_isvc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIsvcResponse); i {
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
		file_isvc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelSpec); i {
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
		file_isvc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsvcSpec); i {
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
		file_isvc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsvcNull); i {
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
			RawDescriptor: file_isvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_isvc_proto_goTypes,
		DependencyIndexes: file_isvc_proto_depIdxs,
		MessageInfos:      file_isvc_proto_msgTypes,
	}.Build()
	File_isvc_proto = out.File
	file_isvc_proto_rawDesc = nil
	file_isvc_proto_goTypes = nil
	file_isvc_proto_depIdxs = nil
}
