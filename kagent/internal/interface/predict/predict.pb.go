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
// source: predict.proto

package predict

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

type KPredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelName    string   `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	ModelVersion string   `protobuf:"bytes,2,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	ModelType    string   `protobuf:"bytes,3,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Inputs       []byte   `protobuf:"bytes,4,opt,name=inputs,proto3" json:"inputs,omitempty"`
	OutputFilter []string `protobuf:"bytes,5,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
}

func (x *KPredictRequest) Reset() {
	*x = KPredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KPredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KPredictRequest) ProtoMessage() {}

func (x *KPredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KPredictRequest.ProtoReflect.Descriptor instead.
func (*KPredictRequest) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{0}
}

func (x *KPredictRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *KPredictRequest) GetModelVersion() string {
	if x != nil {
		return x.ModelVersion
	}
	return ""
}

func (x *KPredictRequest) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *KPredictRequest) GetInputs() []byte {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *KPredictRequest) GetOutputFilter() []string {
	if x != nil {
		return x.OutputFilter
	}
	return nil
}

type KPredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outputs []byte `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *KPredictResponse) Reset() {
	*x = KPredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_predict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KPredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KPredictResponse) ProtoMessage() {}

func (x *KPredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_predict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KPredictResponse.ProtoReflect.Descriptor instead.
func (*KPredictResponse) Descriptor() ([]byte, []int) {
	return file_predict_proto_rawDescGZIP(), []int{1}
}

func (x *KPredictResponse) GetOutputs() []byte {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_predict_proto protoreflect.FileDescriptor

var file_predict_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x0f, 0x4b, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x10,
	0x4b, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x32, 0x57, 0x0a, 0x12, 0x4b, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x08, 0x4b, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x4b, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x2e, 0x4b, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_predict_proto_rawDescOnce sync.Once
	file_predict_proto_rawDescData = file_predict_proto_rawDesc
)

func file_predict_proto_rawDescGZIP() []byte {
	file_predict_proto_rawDescOnce.Do(func() {
		file_predict_proto_rawDescData = protoimpl.X.CompressGZIP(file_predict_proto_rawDescData)
	})
	return file_predict_proto_rawDescData
}

var file_predict_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_predict_proto_goTypes = []interface{}{
	(*KPredictRequest)(nil),  // 0: predict.KPredictRequest
	(*KPredictResponse)(nil), // 1: predict.KPredictResponse
}
var file_predict_proto_depIdxs = []int32{
	0, // 0: predict.KPredictionService.KPredict:input_type -> predict.KPredictRequest
	1, // 1: predict.KPredictionService.KPredict:output_type -> predict.KPredictResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_predict_proto_init() }
func file_predict_proto_init() {
	if File_predict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_predict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KPredictRequest); i {
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
		file_predict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KPredictResponse); i {
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
			RawDescriptor: file_predict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_predict_proto_goTypes,
		DependencyIndexes: file_predict_proto_depIdxs,
		MessageInfos:      file_predict_proto_msgTypes,
	}.Build()
	File_predict_proto = out.File
	file_predict_proto_rawDesc = nil
	file_predict_proto_goTypes = nil
	file_predict_proto_depIdxs = nil
}