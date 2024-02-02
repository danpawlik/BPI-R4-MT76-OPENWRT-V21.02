// Copyright 2015 Google LLC
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
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/api/usage.proto

package serviceconfig

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration controlling usage of a service.
type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	//
	// For Google APIs, a Terms of Service requirement must be included here.
	// Google Cloud APIs must include "serviceusage.googleapis.com/tos/cloud".
	// Other Google APIs should include
	// "serviceusage.googleapis.com/tos/universal". Additional ToS can be
	// included based on the business needs.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	// A list of usage rules that apply to individual API methods.
	//
	// **NOTE:** All service configuration rules follow "last one wins" order.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules,proto3" json:"rules,omitempty"`
	// The full resource name of a channel used for sending notifications to the
	// service producer.
	//
	// Google Service Management currently only supports
	// [Google Cloud Pub/Sub](https://cloud.google.com/pubsub) as a notification
	// channel. To use Google Cloud Pub/Sub as the channel, this must be the name
	// of a Cloud Pub/Sub topic that uses the Cloud Pub/Sub topic name format
	// documented in https://cloud.google.com/pubsub/docs/overview.
	ProducerNotificationChannel string `protobuf:"bytes,7,opt,name=producer_notification_channel,json=producerNotificationChannel,proto3" json:"producer_notification_channel,omitempty"`
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_google_api_usage_proto_rawDescGZIP(), []int{0}
}

func (x *Usage) GetRequirements() []string {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *Usage) GetRules() []*UsageRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Usage) GetProducerNotificationChannel() string {
	if x != nil {
		return x.ProducerNotificationChannel
	}
	return ""
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//	usage:
//	  rules:
//	  - selector: "*"
//	    allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//	usage:
//	  rules:
//	  - selector: "google.example.library.v1.LibraryService.CreateBook"
//	    allow_unregistered_calls: true
type UsageRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// If true, the selected method allows unregistered calls, e.g. calls
	// that don't identify any user or application.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls,proto3" json:"allow_unregistered_calls,omitempty"`
	// If true, the selected method should skip service control and the control
	// plane features, such as quota and billing, will not be available.
	// This flag is used by Google Cloud Endpoints to bypass checks for internal
	// methods, such as service health check methods.
	SkipServiceControl bool `protobuf:"varint,3,opt,name=skip_service_control,json=skipServiceControl,proto3" json:"skip_service_control,omitempty"`
}

func (x *UsageRule) Reset() {
	*x = UsageRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageRule) ProtoMessage() {}

func (x *UsageRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageRule.ProtoReflect.Descriptor instead.
func (*UsageRule) Descriptor() ([]byte, []int) {
	return file_google_api_usage_proto_rawDescGZIP(), []int{1}
}

func (x *UsageRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *UsageRule) GetAllowUnregisteredCalls() bool {
	if x != nil {
		return x.AllowUnregisteredCalls
	}
	return false
}

func (x *UsageRule) GetSkipServiceControl() bool {
	if x != nil {
		return x.SkipServiceControl
	}
	return false
}

var File_google_api_usage_proto protoreflect.FileDescriptor

var file_google_api_usage_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0x9c, 0x01, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12,
	0x42, 0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a,
	0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x6b, 0x69, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x73, 0x6b, 0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0a, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0xa2, 0x02, 0x04, 0x47, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_usage_proto_rawDescOnce sync.Once
	file_google_api_usage_proto_rawDescData = file_google_api_usage_proto_rawDesc
)

func file_google_api_usage_proto_rawDescGZIP() []byte {
	file_google_api_usage_proto_rawDescOnce.Do(func() {
		file_google_api_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_usage_proto_rawDescData)
	})
	return file_google_api_usage_proto_rawDescData
}

var file_google_api_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_api_usage_proto_goTypes = []interface{}{
	(*Usage)(nil),     // 0: google.api.Usage
	(*UsageRule)(nil), // 1: google.api.UsageRule
}
var file_google_api_usage_proto_depIdxs = []int32{
	1, // 0: google.api.Usage.rules:type_name -> google.api.UsageRule
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_api_usage_proto_init() }
func file_google_api_usage_proto_init() {
	if File_google_api_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_usage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_google_api_usage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageRule); i {
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
			RawDescriptor: file_google_api_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_usage_proto_goTypes,
		DependencyIndexes: file_google_api_usage_proto_depIdxs,
		MessageInfos:      file_google_api_usage_proto_msgTypes,
	}.Build()
	File_google_api_usage_proto = out.File
	file_google_api_usage_proto_rawDesc = nil
	file_google_api_usage_proto_goTypes = nil
	file_google_api_usage_proto_depIdxs = nil
}
