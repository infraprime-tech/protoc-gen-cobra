// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: types.proto

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

type Sound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapField  map[string]string `protobuf:"bytes,1,rep,name=map_field,json=mapField,proto3" json:"map_field,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ListField []string          `protobuf:"bytes,2,rep,name=list_field,json=listField,proto3" json:"list_field,omitempty"`
}

func (x *Sound) Reset() {
	*x = Sound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sound) ProtoMessage() {}

func (x *Sound) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sound.ProtoReflect.Descriptor instead.
func (*Sound) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Sound) GetMapField() map[string]string {
	if x != nil {
		return x.MapField
	}
	return nil
}

func (x *Sound) GetListField() []string {
	if x != nil {
		return x.ListField
	}
	return nil
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x05, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x39, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x4d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x2f, 0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x26, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x1a, 0x0e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(*Sound)(nil), // 0: example.Sound
	nil,           // 1: example.Sound.MapFieldEntry
}
var file_types_proto_depIdxs = []int32{
	1, // 0: example.Sound.map_field:type_name -> example.Sound.MapFieldEntry
	0, // 1: example.Types.Echo:input_type -> example.Sound
	0, // 2: example.Types.Echo:output_type -> example.Sound
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sound); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
