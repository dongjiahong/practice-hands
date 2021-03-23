// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: helloworld.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/api/kratos/api"
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

type Helloworld int32

const (
	Helloworld_MissingName Helloworld = 0
)

// Enum value maps for Helloworld.
var (
	Helloworld_name = map[int32]string{
		0: "MissingName",
	}
	Helloworld_value = map[string]int32{
		"MissingName": 0,
	}
)

func (x Helloworld) Enum() *Helloworld {
	p := new(Helloworld)
	*p = x
	return p
}

func (x Helloworld) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Helloworld) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_proto_enumTypes[0].Descriptor()
}

func (Helloworld) Type() protoreflect.EnumType {
	return &file_helloworld_proto_enumTypes[0]
}

func (x Helloworld) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Helloworld.Descriptor instead.
func (Helloworld) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_proto_rawDescGZIP(), []int{0}
}

var File_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x22, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x00, 0x1a, 0x03, 0xa0, 0x45, 0x01, 0x42, 0x4e, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x18, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x10, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x44, 0x65,
	0x6d, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_proto_rawDescData = file_helloworld_proto_rawDesc
)

func file_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_proto_rawDescData)
	})
	return file_helloworld_proto_rawDescData
}

var file_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_proto_goTypes = []interface{}{
	(Helloworld)(0), // 0: kratos.demo.errors.Helloworld
}
var file_helloworld_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_proto_init() }
func file_helloworld_proto_init() {
	if File_helloworld_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_helloworld_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_proto_depIdxs,
		EnumInfos:         file_helloworld_proto_enumTypes,
	}.Build()
	File_helloworld_proto = out.File
	file_helloworld_proto_rawDesc = nil
	file_helloworld_proto_goTypes = nil
	file_helloworld_proto_depIdxs = nil
}
