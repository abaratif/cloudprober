// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: github.com/cloudprober/cloudprober/targets/testdata/testdata.proto

package testdata

import (
	proto "github.com/cloudprober/cloudprober/targets/proto"
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

type FancyTargets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *FancyTargets) Reset() {
	*x = FancyTargets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FancyTargets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FancyTargets) ProtoMessage() {}

func (x *FancyTargets) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FancyTargets.ProtoReflect.Descriptor instead.
func (*FancyTargets) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescGZIP(), []int{0}
}

func (x *FancyTargets) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type AnotherFancyTargets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *AnotherFancyTargets) Reset() {
	*x = AnotherFancyTargets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnotherFancyTargets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnotherFancyTargets) ProtoMessage() {}

func (x *AnotherFancyTargets) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnotherFancyTargets.ProtoReflect.Descriptor instead.
func (*AnotherFancyTargets) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *AnotherFancyTargets) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*proto.TargetsDef)(nil),
		ExtensionType: (*FancyTargets)(nil),
		Field:         200,
		Name:          "cloudprober.targets.testdata.fancy_targets",
		Tag:           "bytes,200,opt,name=fancy_targets",
		Filename:      "github.com/cloudprober/cloudprober/targets/testdata/testdata.proto",
	},
	{
		ExtendedType:  (*proto.TargetsDef)(nil),
		ExtensionType: (*AnotherFancyTargets)(nil),
		Field:         201,
		Name:          "cloudprober.targets.testdata.another_fancy_targets",
		Tag:           "bytes,201,opt,name=another_fancy_targets",
		Filename:      "github.com/cloudprober/cloudprober/targets/testdata/testdata.proto",
	},
}

// Extension fields to proto.TargetsDef.
var (
	// optional cloudprober.targets.testdata.FancyTargets fancy_targets = 200;
	E_FancyTargets = &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_extTypes[0]
	// optional cloudprober.targets.testdata.AnotherFancyTargets another_fancy_targets = 201;
	E_AnotherFancyTargets = &file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_extTypes[1]
)

var File_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x3a, 0x71, 0x0a, 0x0d, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x44, 0x65, 0x66, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x0c, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x3a, 0x87, 0x01, 0x0a, 0x15, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x66, 0x61, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x1f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x44, 0x65, 0x66, 0x18,
	0xc9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x61, 0x6e,
	0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x13, 0x61, 0x6e, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescData = file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_goTypes = []interface{}{
	(*FancyTargets)(nil),        // 0: cloudprober.targets.testdata.FancyTargets
	(*AnotherFancyTargets)(nil), // 1: cloudprober.targets.testdata.AnotherFancyTargets
	(*proto.TargetsDef)(nil),    // 2: cloudprober.targets.TargetsDef
}
var file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_depIdxs = []int32{
	2, // 0: cloudprober.targets.testdata.fancy_targets:extendee -> cloudprober.targets.TargetsDef
	2, // 1: cloudprober.targets.testdata.another_fancy_targets:extendee -> cloudprober.targets.TargetsDef
	0, // 2: cloudprober.targets.testdata.fancy_targets:type_name -> cloudprober.targets.testdata.FancyTargets
	1, // 3: cloudprober.targets.testdata.another_fancy_targets:type_name -> cloudprober.targets.testdata.AnotherFancyTargets
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_init() }
func file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_init() {
	if File_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FancyTargets); i {
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
		file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnotherFancyTargets); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_msgTypes,
		ExtensionInfos:    file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_extTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto = out.File
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_targets_testdata_testdata_proto_depIdxs = nil
}
