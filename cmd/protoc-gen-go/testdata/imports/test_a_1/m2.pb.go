// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_1/m2.proto

package test_a_1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type M2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *M2) Reset() {
	*x = M2{}
}

func (x *M2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*M2) ProtoMessage() {}

func (x *M2) ProtoReflect() protoreflect.Message {
	mi := &file_imports_test_a_1_m2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *M2) XXX_Methods() *protoiface.Methods {
	return file_imports_test_a_1_m2_proto_msgTypes[0].Methods()
}

// Deprecated: Use M2.ProtoReflect.Type instead.
func (*M2) Descriptor() ([]byte, []int) {
	return file_imports_test_a_1_m2_proto_rawDescGZIP(), []int{0}
}

var File_imports_test_a_1_m2_proto protoreflect.FileDescriptor

var file_imports_test_a_1_m2_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x31, 0x2f, 0x6d, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x32, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x61, 0x5f, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imports_test_a_1_m2_proto_rawDescOnce sync.Once
	file_imports_test_a_1_m2_proto_rawDescData = file_imports_test_a_1_m2_proto_rawDesc
)

func file_imports_test_a_1_m2_proto_rawDescGZIP() []byte {
	file_imports_test_a_1_m2_proto_rawDescOnce.Do(func() {
		file_imports_test_a_1_m2_proto_rawDescData = protoimpl.X.CompressGZIP(file_imports_test_a_1_m2_proto_rawDescData)
	})
	return file_imports_test_a_1_m2_proto_rawDescData
}

var file_imports_test_a_1_m2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_imports_test_a_1_m2_proto_goTypes = []interface{}{
	(*M2)(nil), // 0: test.a.M2
}
var file_imports_test_a_1_m2_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_imports_test_a_1_m2_proto_init() }
func file_imports_test_a_1_m2_proto_init() {
	if File_imports_test_a_1_m2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imports_test_a_1_m2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*M2); i {
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
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_imports_test_a_1_m2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_imports_test_a_1_m2_proto_goTypes,
		DependencyIndexes: file_imports_test_a_1_m2_proto_depIdxs,
		MessageInfos:      file_imports_test_a_1_m2_proto_msgTypes,
	}.Build()
	File_imports_test_a_1_m2_proto = out.File
	file_imports_test_a_1_m2_proto_rawDesc = nil
	file_imports_test_a_1_m2_proto_goTypes = nil
	file_imports_test_a_1_m2_proto_depIdxs = nil
}
