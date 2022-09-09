// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: execution_requested.proto

package events

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

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId string   `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Checks  []string `protobuf:"bytes,2,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_requested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_execution_requested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_execution_requested_proto_rawDescGZIP(), []int{0}
}

func (x *Target) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *Target) GetChecks() []string {
	if x != nil {
		return x.Checks
	}
	return nil
}

type ExecutionRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionId string    `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	GroupId     string    `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Targets     []*Target `protobuf:"bytes,3,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *ExecutionRequested) Reset() {
	*x = ExecutionRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execution_requested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionRequested) ProtoMessage() {}

func (x *ExecutionRequested) ProtoReflect() protoreflect.Message {
	mi := &file_execution_requested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionRequested.ProtoReflect.Descriptor instead.
func (*ExecutionRequested) Descriptor() ([]byte, []int) {
	return file_execution_requested_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionRequested) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *ExecutionRequested) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ExecutionRequested) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

var File_execution_requested_proto protoreflect.FileDescriptor

var file_execution_requested_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x72, 0x65,
	0x6e, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x3b, 0x0a,
	0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_execution_requested_proto_rawDescOnce sync.Once
	file_execution_requested_proto_rawDescData = file_execution_requested_proto_rawDesc
)

func file_execution_requested_proto_rawDescGZIP() []byte {
	file_execution_requested_proto_rawDescOnce.Do(func() {
		file_execution_requested_proto_rawDescData = protoimpl.X.CompressGZIP(file_execution_requested_proto_rawDescData)
	})
	return file_execution_requested_proto_rawDescData
}

var file_execution_requested_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_execution_requested_proto_goTypes = []interface{}{
	(*Target)(nil),             // 0: trento.checks.v1.Target
	(*ExecutionRequested)(nil), // 1: trento.checks.v1.ExecutionRequested
}
var file_execution_requested_proto_depIdxs = []int32{
	0, // 0: trento.checks.v1.ExecutionRequested.targets:type_name -> trento.checks.v1.Target
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_execution_requested_proto_init() }
func file_execution_requested_proto_init() {
	if File_execution_requested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execution_requested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_execution_requested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionRequested); i {
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
			RawDescriptor: file_execution_requested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_execution_requested_proto_goTypes,
		DependencyIndexes: file_execution_requested_proto_depIdxs,
		MessageInfos:      file_execution_requested_proto_msgTypes,
	}.Build()
	File_execution_requested_proto = out.File
	file_execution_requested_proto_rawDesc = nil
	file_execution_requested_proto_goTypes = nil
	file_execution_requested_proto_depIdxs = nil
}
