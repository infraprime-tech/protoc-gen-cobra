// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: nested.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TopLevelNestedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TopLevelNestedType) Reset() {
	*x = TopLevelNestedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopLevelNestedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopLevelNestedType) ProtoMessage() {}

func (x *TopLevelNestedType) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopLevelNestedType.ProtoReflect.Descriptor instead.
func (*TopLevelNestedType) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{0}
}

func (x *TopLevelNestedType) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type NestedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner    *NestedRequest_InnerNestedType `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	TopLevel *TopLevelNestedType            `protobuf:"bytes,2,opt,name=top_level,json=topLevel,proto3" json:"top_level,omitempty"`
}

func (x *NestedRequest) Reset() {
	*x = NestedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRequest) ProtoMessage() {}

func (x *NestedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRequest.ProtoReflect.Descriptor instead.
func (*NestedRequest) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{1}
}

func (x *NestedRequest) GetInner() *NestedRequest_InnerNestedType {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *NestedRequest) GetTopLevel() *TopLevelNestedType {
	if x != nil {
		return x.TopLevel
	}
	return nil
}

type NestedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Return string `protobuf:"bytes,1,opt,name=return,proto3" json:"return,omitempty"`
}

func (x *NestedResponse) Reset() {
	*x = NestedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedResponse) ProtoMessage() {}

func (x *NestedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedResponse.ProtoReflect.Descriptor instead.
func (*NestedResponse) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{2}
}

func (x *NestedResponse) GetReturn() string {
	if x != nil {
		return x.Return
	}
	return ""
}

type DeeplyNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L0 *DeeplyNested_DeeplyNestedOuter `protobuf:"bytes,1,opt,name=l0,proto3" json:"l0,omitempty"`
}

func (x *DeeplyNested) Reset() {
	*x = DeeplyNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeeplyNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeeplyNested) ProtoMessage() {}

func (x *DeeplyNested) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeeplyNested.ProtoReflect.Descriptor instead.
func (*DeeplyNested) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{3}
}

func (x *DeeplyNested) GetL0() *DeeplyNested_DeeplyNestedOuter {
	if x != nil {
		return x.L0
	}
	return nil
}

type NestedRequest_InnerNestedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NestedRequest_InnerNestedType) Reset() {
	*x = NestedRequest_InnerNestedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NestedRequest_InnerNestedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NestedRequest_InnerNestedType) ProtoMessage() {}

func (x *NestedRequest_InnerNestedType) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NestedRequest_InnerNestedType.ProtoReflect.Descriptor instead.
func (*NestedRequest_InnerNestedType) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{1, 0}
}

func (x *NestedRequest_InnerNestedType) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeeplyNested_DeeplyNestedOuter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L1 *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner `protobuf:"bytes,1,opt,name=l1,proto3" json:"l1,omitempty"`
}

func (x *DeeplyNested_DeeplyNestedOuter) Reset() {
	*x = DeeplyNested_DeeplyNestedOuter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeeplyNested_DeeplyNestedOuter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeeplyNested_DeeplyNestedOuter) ProtoMessage() {}

func (x *DeeplyNested_DeeplyNestedOuter) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeeplyNested_DeeplyNestedOuter.ProtoReflect.Descriptor instead.
func (*DeeplyNested_DeeplyNestedOuter) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DeeplyNested_DeeplyNestedOuter) GetL1() *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner {
	if x != nil {
		return x.L1
	}
	return nil
}

type DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L2 *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost `protobuf:"bytes,1,opt,name=l2,proto3" json:"l2,omitempty"`
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) Reset() {
	*x = DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) ProtoMessage() {}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner.ProtoReflect.Descriptor instead.
func (*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{3, 0, 0}
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner) GetL2() *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost {
	if x != nil {
		return x.L2
	}
	return nil
}

type DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	L3 string `protobuf:"bytes,1,opt,name=l3,proto3" json:"l3,omitempty"`
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) Reset() {
	*x = DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nested_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) ProtoMessage() {}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) ProtoReflect() protoreflect.Message {
	mi := &file_nested_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost.ProtoReflect.Descriptor instead.
func (*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) Descriptor() ([]byte, []int) {
	return file_nested_proto_rawDescGZIP(), []int{3, 0, 0, 0}
}

func (x *DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost) GetL3() string {
	if x != nil {
		return x.L3
	}
	return ""
}

var File_nested_proto protoreflect.FileDescriptor

var file_nested_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x54, 0x6f, 0x70, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x27, 0x0a,
	0x0f, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x22, 0xc8, 0x02, 0x0a, 0x0c, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x37, 0x0a, 0x02, 0x6c, 0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x52, 0x02, 0x6c, 0x30, 0x1a, 0xfe, 0x01, 0x0a, 0x11, 0x44,
	0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x65, 0x72,
	0x12, 0x49, 0x0a, 0x02, 0x6c, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x02, 0x6c, 0x31, 0x1a, 0x9d, 0x01, 0x0a, 0x11,
	0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x5f, 0x0a, 0x02, 0x6c, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x52, 0x02,
	0x6c, 0x32, 0x1a, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x6c,
	0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6c, 0x33, 0x32, 0x83, 0x01, 0x0a, 0x06,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x65, 0x70, 0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x12, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x65, 0x70,
	0x6c, 0x79, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_nested_proto_rawDescOnce sync.Once
	file_nested_proto_rawDescData = file_nested_proto_rawDesc
)

func file_nested_proto_rawDescGZIP() []byte {
	file_nested_proto_rawDescOnce.Do(func() {
		file_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_nested_proto_rawDescData)
	})
	return file_nested_proto_rawDescData
}

var file_nested_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_nested_proto_goTypes = []interface{}{
	(*TopLevelNestedType)(nil),                               // 0: example.TopLevelNestedType
	(*NestedRequest)(nil),                                    // 1: example.NestedRequest
	(*NestedResponse)(nil),                                   // 2: example.NestedResponse
	(*DeeplyNested)(nil),                                     // 3: example.DeeplyNested
	(*NestedRequest_InnerNestedType)(nil),                    // 4: example.NestedRequest.InnerNestedType
	(*DeeplyNested_DeeplyNestedOuter)(nil),                   // 5: example.DeeplyNested.DeeplyNestedOuter
	(*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner)(nil), // 6: example.DeeplyNested.DeeplyNestedOuter.DeeplyNestedInner
	(*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost)(nil), // 7: example.DeeplyNested.DeeplyNestedOuter.DeeplyNestedInner.DeeplyNestedInnermost
}
var file_nested_proto_depIdxs = []int32{
	4, // 0: example.NestedRequest.inner:type_name -> example.NestedRequest.InnerNestedType
	0, // 1: example.NestedRequest.top_level:type_name -> example.TopLevelNestedType
	5, // 2: example.DeeplyNested.l0:type_name -> example.DeeplyNested.DeeplyNestedOuter
	6, // 3: example.DeeplyNested.DeeplyNestedOuter.l1:type_name -> example.DeeplyNested.DeeplyNestedOuter.DeeplyNestedInner
	7, // 4: example.DeeplyNested.DeeplyNestedOuter.DeeplyNestedInner.l2:type_name -> example.DeeplyNested.DeeplyNestedOuter.DeeplyNestedInner.DeeplyNestedInnermost
	1, // 5: example.Nested.Get:input_type -> example.NestedRequest
	3, // 6: example.Nested.GetDeeplyNested:input_type -> example.DeeplyNested
	2, // 7: example.Nested.Get:output_type -> example.NestedResponse
	2, // 8: example.Nested.GetDeeplyNested:output_type -> example.NestedResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nested_proto_init() }
func file_nested_proto_init() {
	if File_nested_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopLevelNestedType); i {
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
		file_nested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRequest); i {
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
		file_nested_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedResponse); i {
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
		file_nested_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeeplyNested); i {
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
		file_nested_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NestedRequest_InnerNestedType); i {
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
		file_nested_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeeplyNested_DeeplyNestedOuter); i {
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
		file_nested_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner); i {
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
		file_nested_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeeplyNested_DeeplyNestedOuter_DeeplyNestedInner_DeeplyNestedInnermost); i {
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
			RawDescriptor: file_nested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nested_proto_goTypes,
		DependencyIndexes: file_nested_proto_depIdxs,
		MessageInfos:      file_nested_proto_msgTypes,
	}.Build()
	File_nested_proto = out.File
	file_nested_proto_rawDesc = nil
	file_nested_proto_goTypes = nil
	file_nested_proto_depIdxs = nil
}
