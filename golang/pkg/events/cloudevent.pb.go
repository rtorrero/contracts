// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: cloudevent.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// CloudEvent is copied from
// https://github.com/cloudevents/spec/blob/master/protobuf-format.md.
type CloudEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique event identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// URI of the event source.
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Version of the spec in use.
	SpecVersion string `protobuf:"bytes,3,opt,name=spec_version,json=specVersion,proto3" json:"spec_version,omitempty"`
	// Event type identifier.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Optional & Extension Attributes
	Attributes map[string]*CloudEventAttributeValue `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// CloudEvent Data (Bytes, Text, or Proto)
	//
	// Types that are assignable to Data:
	//
	//	*CloudEvent_BinaryData
	//	*CloudEvent_TextData
	//	*CloudEvent_ProtoData
	Data isCloudEvent_Data `protobuf_oneof:"data"`
}

func (x *CloudEvent) Reset() {
	*x = CloudEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudevent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEvent) ProtoMessage() {}

func (x *CloudEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cloudevent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEvent.ProtoReflect.Descriptor instead.
func (*CloudEvent) Descriptor() ([]byte, []int) {
	return file_cloudevent_proto_rawDescGZIP(), []int{0}
}

func (x *CloudEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloudEvent) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CloudEvent) GetSpecVersion() string {
	if x != nil {
		return x.SpecVersion
	}
	return ""
}

func (x *CloudEvent) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CloudEvent) GetAttributes() map[string]*CloudEventAttributeValue {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (m *CloudEvent) GetData() isCloudEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *CloudEvent) GetBinaryData() []byte {
	if x, ok := x.GetData().(*CloudEvent_BinaryData); ok {
		return x.BinaryData
	}
	return nil
}

func (x *CloudEvent) GetTextData() string {
	if x, ok := x.GetData().(*CloudEvent_TextData); ok {
		return x.TextData
	}
	return ""
}

func (x *CloudEvent) GetProtoData() *anypb.Any {
	if x, ok := x.GetData().(*CloudEvent_ProtoData); ok {
		return x.ProtoData
	}
	return nil
}

type isCloudEvent_Data interface {
	isCloudEvent_Data()
}

type CloudEvent_BinaryData struct {
	// If the event is binary data then the datacontenttype attribute
	// should be set to an appropriate media-type.
	BinaryData []byte `protobuf:"bytes,6,opt,name=binary_data,json=binaryData,proto3,oneof"`
}

type CloudEvent_TextData struct {
	// If the event is string data then the datacontenttype attribute
	// should be set to an appropriate media-type such as application/json.
	TextData string `protobuf:"bytes,7,opt,name=text_data,json=textData,proto3,oneof"`
}

type CloudEvent_ProtoData struct {
	// If the event is a protobuf then it must be encoded using this Any
	// type. The datacontenttype attribute should be set to
	// application/protobuf and the dataschema attribute set to the message
	// type.
	ProtoData *anypb.Any `protobuf:"bytes,8,opt,name=proto_data,json=protoData,proto3,oneof"`
}

func (*CloudEvent_BinaryData) isCloudEvent_Data() {}

func (*CloudEvent_TextData) isCloudEvent_Data() {}

func (*CloudEvent_ProtoData) isCloudEvent_Data() {}

// CloudEventAttribute enables extensions to use any of the seven allowed
// data types as the value of an envelope key.
type CloudEventAttributeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value can be any one of these types.
	//
	// Types that are assignable to Attr:
	//
	//	*CloudEventAttributeValue_CeBoolean
	//	*CloudEventAttributeValue_CeInteger
	//	*CloudEventAttributeValue_CeString
	//	*CloudEventAttributeValue_CeBytes
	//	*CloudEventAttributeValue_CeUri
	//	*CloudEventAttributeValue_CeUriRef
	//	*CloudEventAttributeValue_CeTimestamp
	Attr isCloudEventAttributeValue_Attr `protobuf_oneof:"attr"`
}

func (x *CloudEventAttributeValue) Reset() {
	*x = CloudEventAttributeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudevent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudEventAttributeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudEventAttributeValue) ProtoMessage() {}

func (x *CloudEventAttributeValue) ProtoReflect() protoreflect.Message {
	mi := &file_cloudevent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudEventAttributeValue.ProtoReflect.Descriptor instead.
func (*CloudEventAttributeValue) Descriptor() ([]byte, []int) {
	return file_cloudevent_proto_rawDescGZIP(), []int{1}
}

func (m *CloudEventAttributeValue) GetAttr() isCloudEventAttributeValue_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (x *CloudEventAttributeValue) GetCeBoolean() bool {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeBoolean); ok {
		return x.CeBoolean
	}
	return false
}

func (x *CloudEventAttributeValue) GetCeInteger() int32 {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeInteger); ok {
		return x.CeInteger
	}
	return 0
}

func (x *CloudEventAttributeValue) GetCeString() string {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeString); ok {
		return x.CeString
	}
	return ""
}

func (x *CloudEventAttributeValue) GetCeBytes() []byte {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeBytes); ok {
		return x.CeBytes
	}
	return nil
}

func (x *CloudEventAttributeValue) GetCeUri() string {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeUri); ok {
		return x.CeUri
	}
	return ""
}

func (x *CloudEventAttributeValue) GetCeUriRef() string {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeUriRef); ok {
		return x.CeUriRef
	}
	return ""
}

func (x *CloudEventAttributeValue) GetCeTimestamp() *timestamppb.Timestamp {
	if x, ok := x.GetAttr().(*CloudEventAttributeValue_CeTimestamp); ok {
		return x.CeTimestamp
	}
	return nil
}

type isCloudEventAttributeValue_Attr interface {
	isCloudEventAttributeValue_Attr()
}

type CloudEventAttributeValue_CeBoolean struct {
	// Boolean value.
	CeBoolean bool `protobuf:"varint,1,opt,name=ce_boolean,json=ceBoolean,proto3,oneof"`
}

type CloudEventAttributeValue_CeInteger struct {
	// Integer value.
	CeInteger int32 `protobuf:"varint,2,opt,name=ce_integer,json=ceInteger,proto3,oneof"`
}

type CloudEventAttributeValue_CeString struct {
	// String value.
	CeString string `protobuf:"bytes,3,opt,name=ce_string,json=ceString,proto3,oneof"`
}

type CloudEventAttributeValue_CeBytes struct {
	// Byte string value.
	CeBytes []byte `protobuf:"bytes,4,opt,name=ce_bytes,json=ceBytes,proto3,oneof"`
}

type CloudEventAttributeValue_CeUri struct {
	// URI value.
	CeUri string `protobuf:"bytes,5,opt,name=ce_uri,json=ceUri,proto3,oneof"`
}

type CloudEventAttributeValue_CeUriRef struct {
	// URI reference value.
	CeUriRef string `protobuf:"bytes,6,opt,name=ce_uri_ref,json=ceUriRef,proto3,oneof"`
}

type CloudEventAttributeValue_CeTimestamp struct {
	// Timestamp value.
	CeTimestamp *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=ce_timestamp,json=ceTimestamp,proto3,oneof"`
}

func (*CloudEventAttributeValue_CeBoolean) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeInteger) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeString) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeBytes) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeUri) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeUriRef) isCloudEventAttributeValue_Attr() {}

func (*CloudEventAttributeValue_CeTimestamp) isCloudEventAttributeValue_Attr() {}

var File_cloudevent_proto protoreflect.FileDescriptor

var file_cloudevent_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x03, 0x0a, 0x0a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x63, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x64, 0x0a, 0x0f, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9a, 0x02, 0x0a, 0x18, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x6f,
	0x6c, 0x65, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x63, 0x65,
	0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x65, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x63, 0x65, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x63, 0x65, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x65, 0x55, 0x72, 0x69, 0x52, 0x65, 0x66, 0x12, 0x3f, 0x0a,
	0x0c, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06,
	0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudevent_proto_rawDescOnce sync.Once
	file_cloudevent_proto_rawDescData = file_cloudevent_proto_rawDesc
)

func file_cloudevent_proto_rawDescGZIP() []byte {
	file_cloudevent_proto_rawDescOnce.Do(func() {
		file_cloudevent_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudevent_proto_rawDescData)
	})
	return file_cloudevent_proto_rawDescData
}

var file_cloudevent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloudevent_proto_goTypes = []interface{}{
	(*CloudEvent)(nil),               // 0: cloudevents.CloudEvent
	(*CloudEventAttributeValue)(nil), // 1: cloudevents.CloudEventAttributeValue
	nil,                              // 2: cloudevents.CloudEvent.AttributesEntry
	(*anypb.Any)(nil),                // 3: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_cloudevent_proto_depIdxs = []int32{
	2, // 0: cloudevents.CloudEvent.attributes:type_name -> cloudevents.CloudEvent.AttributesEntry
	3, // 1: cloudevents.CloudEvent.proto_data:type_name -> google.protobuf.Any
	4, // 2: cloudevents.CloudEventAttributeValue.ce_timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: cloudevents.CloudEvent.AttributesEntry.value:type_name -> cloudevents.CloudEventAttributeValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cloudevent_proto_init() }
func file_cloudevent_proto_init() {
	if File_cloudevent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudevent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEvent); i {
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
		file_cloudevent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudEventAttributeValue); i {
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
	file_cloudevent_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CloudEvent_BinaryData)(nil),
		(*CloudEvent_TextData)(nil),
		(*CloudEvent_ProtoData)(nil),
	}
	file_cloudevent_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CloudEventAttributeValue_CeBoolean)(nil),
		(*CloudEventAttributeValue_CeInteger)(nil),
		(*CloudEventAttributeValue_CeString)(nil),
		(*CloudEventAttributeValue_CeBytes)(nil),
		(*CloudEventAttributeValue_CeUri)(nil),
		(*CloudEventAttributeValue_CeUriRef)(nil),
		(*CloudEventAttributeValue_CeTimestamp)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudevent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudevent_proto_goTypes,
		DependencyIndexes: file_cloudevent_proto_depIdxs,
		MessageInfos:      file_cloudevent_proto_msgTypes,
	}.Build()
	File_cloudevent_proto = out.File
	file_cloudevent_proto_rawDesc = nil
	file_cloudevent_proto_goTypes = nil
	file_cloudevent_proto_depIdxs = nil
}
