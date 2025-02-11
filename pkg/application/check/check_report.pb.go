// Copyright 2023 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: google/cloud/apigeeregistry/v1/check/check_report.proto

// (-- api-linter: core::0215::versioned-packages=disabled
//     aip.dev/not-precedent: Support protos for the apigeeregistry.v1 API. --)

package check

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible severities for the violation of a rule.
type Problem_Severity int32

const (
	// The default value, unused.
	Problem_SEVERITY_UNSPECIFIED Problem_Severity = 0
	// Violation of the rule is an error that must be fixed.
	Problem_ERROR Problem_Severity = 1
	// Violation of the rule is a pattern that is wrong,
	// and should be warned about.
	Problem_WARNING Problem_Severity = 2
	// Violation of the rule is not necessarily a bad pattern
	// or error, but information the user should be aware of.
	Problem_INFO Problem_Severity = 3
	// Violation of the rule is a hint that is provided to
	// the user to fix their spec's design.
	Problem_HINT Problem_Severity = 4
)

// Enum value maps for Problem_Severity.
var (
	Problem_Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "ERROR",
		2: "WARNING",
		3: "INFO",
		4: "HINT",
	}
	Problem_Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"ERROR":                1,
		"WARNING":              2,
		"INFO":                 3,
		"HINT":                 4,
	}
)

func (x Problem_Severity) Enum() *Problem_Severity {
	p := new(Problem_Severity)
	*p = x
	return p
}

func (x Problem_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Problem_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apigeeregistry_v1_check_check_report_proto_enumTypes[0].Descriptor()
}

func (Problem_Severity) Type() protoreflect.EnumType {
	return &file_google_cloud_apigeeregistry_v1_check_check_report_proto_enumTypes[0]
}

func (x Problem_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Problem_Severity.Descriptor instead.
func (Problem_Severity) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescGZIP(), []int{1, 0}
}

// CheckReport is the results of running the check command.
type CheckReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier of the response.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Artifact kind. May be used in YAML representations to identify the type of
	// this artifact.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Creation timestamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// A list of Problems found.
	Problems []*Problem `protobuf:"bytes,4,rep,name=problems,proto3" json:"problems,omitempty"`
	// Populated if check wasn't able to complete due to an error.
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CheckReport) Reset() {
	*x = CheckReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReport) ProtoMessage() {}

func (x *CheckReport) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReport.ProtoReflect.Descriptor instead.
func (*CheckReport) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescGZIP(), []int{0}
}

func (x *CheckReport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CheckReport) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CheckReport) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CheckReport) GetProblems() []*Problem {
	if x != nil {
		return x.Problems
	}
	return nil
}

func (x *CheckReport) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Problem is a result of a rule check.
type Problem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message provides a short description of the problem.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Suggestion provides a suggested fix, if applicable.
	Suggestion string `protobuf:"bytes,3,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// Location provides the location of the problem.
	// If for a Resource, it is the Resource name.
	// If for a field, this is the Resource name + "::" + field name.
	Location string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	// RuleId provides the ID of the rule that this problem belongs to.
	RuleId string `protobuf:"bytes,5,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// RuleDocUri provides a uri to the documented explaination of this rule.
	RuleDocUri string `protobuf:"bytes,6,opt,name=rule_doc_uri,json=ruleDocUri,proto3" json:"rule_doc_uri,omitempty"`
	// Severity provides information on the criticality of the Problem.
	Severity Problem_Severity `protobuf:"varint,7,opt,name=severity,proto3,enum=google.cloud.apigeeregistry.v1.style.Problem_Severity" json:"severity,omitempty"`
}

func (x *Problem) Reset() {
	*x = Problem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Problem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Problem) ProtoMessage() {}

func (x *Problem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Problem.ProtoReflect.Descriptor instead.
func (*Problem) Descriptor() ([]byte, []int) {
	return file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescGZIP(), []int{1}
}

func (x *Problem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Problem) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *Problem) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Problem) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *Problem) GetRuleDocUri() string {
	if x != nil {
		return x.RuleDocUri
	}
	return ""
}

func (x *Problem) GetSeverity() Problem_Severity {
	if x != nil {
		return x.Severity
	}
	return Problem_SEVERITY_UNSPECIFIED
}

var File_google_cloud_apigeeregistry_v1_check_check_report_proto protoreflect.FileDescriptor

var file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc5, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x55, 0x72, 0x69, 0x12,
	0x52, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x49, 0x4e, 0x54, 0x10, 0x04, 0x42, 0x76, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x42, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x65, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescOnce sync.Once
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescData = file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDesc
)

func file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescGZIP() []byte {
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescOnce.Do(func() {
		file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescData)
	})
	return file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDescData
}

var file_google_cloud_apigeeregistry_v1_check_check_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_apigeeregistry_v1_check_check_report_proto_goTypes = []interface{}{
	(Problem_Severity)(0),         // 0: google.cloud.apigeeregistry.v1.style.Problem.Severity
	(*CheckReport)(nil),           // 1: google.cloud.apigeeregistry.v1.style.CheckReport
	(*Problem)(nil),               // 2: google.cloud.apigeeregistry.v1.style.Problem
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_apigeeregistry_v1_check_check_report_proto_depIdxs = []int32{
	3, // 0: google.cloud.apigeeregistry.v1.style.CheckReport.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.apigeeregistry.v1.style.CheckReport.problems:type_name -> google.cloud.apigeeregistry.v1.style.Problem
	0, // 2: google.cloud.apigeeregistry.v1.style.Problem.severity:type_name -> google.cloud.apigeeregistry.v1.style.Problem.Severity
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_apigeeregistry_v1_check_check_report_proto_init() }
func file_google_cloud_apigeeregistry_v1_check_check_report_proto_init() {
	if File_google_cloud_apigeeregistry_v1_check_check_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckReport); i {
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
		file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Problem); i {
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
			RawDescriptor: file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apigeeregistry_v1_check_check_report_proto_goTypes,
		DependencyIndexes: file_google_cloud_apigeeregistry_v1_check_check_report_proto_depIdxs,
		EnumInfos:         file_google_cloud_apigeeregistry_v1_check_check_report_proto_enumTypes,
		MessageInfos:      file_google_cloud_apigeeregistry_v1_check_check_report_proto_msgTypes,
	}.Build()
	File_google_cloud_apigeeregistry_v1_check_check_report_proto = out.File
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_rawDesc = nil
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_goTypes = nil
	file_google_cloud_apigeeregistry_v1_check_check_report_proto_depIdxs = nil
}
