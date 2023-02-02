// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: github.com/cloudprober/cloudprober/surfacers/probestatus/proto/config.proto

package proto

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

type SurfacerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// default 60s
	ResolutionSec *int32 `protobuf:"varint,1,opt,name=resolution_sec,json=resolutionSec,def=60" json:"resolution_sec,omitempty"`
	// Number of points in each timeseries. This field dictates how far back
	// can you go up to (resolution_sec * timeseries_size). Note that higher
	// this number, more memory you'll use.
	TimeseriesSize *int32 `protobuf:"varint,2,opt,name=timeseries_size,json=timeseriesSize,def=4320" json:"timeseries_size,omitempty"`
	// Max targets per probe.
	MaxTargetsPerProbe *int32 `protobuf:"varint,3,opt,name=max_targets_per_probe,json=maxTargetsPerProbe,def=20" json:"max_targets_per_probe,omitempty"`
	// ProbeStatus URL
	// Note that older default URL /probestatus forwards to this URL to avoid
	// breaking older default setups.
	Url *string `protobuf:"bytes,4,opt,name=url,def=/status" json:"url,omitempty"`
	// Page cache time
	CacheTimeSec *int32 `protobuf:"varint,5,opt,name=cache_time_sec,json=cacheTimeSec,def=2" json:"cache_time_sec,omitempty"`
	// Probestatus surfacer is enabled by default. To disable it, set this
	// option.
	Disable *bool `protobuf:"varint,6,opt,name=disable" json:"disable,omitempty"`
}

// Default values for SurfacerConf fields.
const (
	Default_SurfacerConf_ResolutionSec      = int32(60)
	Default_SurfacerConf_TimeseriesSize     = int32(4320)
	Default_SurfacerConf_MaxTargetsPerProbe = int32(20)
	Default_SurfacerConf_Url                = string("/status")
	Default_SurfacerConf_CacheTimeSec       = int32(2)
)

func (x *SurfacerConf) Reset() {
	*x = SurfacerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurfacerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurfacerConf) ProtoMessage() {}

func (x *SurfacerConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurfacerConf.ProtoReflect.Descriptor instead.
func (*SurfacerConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *SurfacerConf) GetResolutionSec() int32 {
	if x != nil && x.ResolutionSec != nil {
		return *x.ResolutionSec
	}
	return Default_SurfacerConf_ResolutionSec
}

func (x *SurfacerConf) GetTimeseriesSize() int32 {
	if x != nil && x.TimeseriesSize != nil {
		return *x.TimeseriesSize
	}
	return Default_SurfacerConf_TimeseriesSize
}

func (x *SurfacerConf) GetMaxTargetsPerProbe() int32 {
	if x != nil && x.MaxTargetsPerProbe != nil {
		return *x.MaxTargetsPerProbe
	}
	return Default_SurfacerConf_MaxTargetsPerProbe
}

func (x *SurfacerConf) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return Default_SurfacerConf_Url
}

func (x *SurfacerConf) GetCacheTimeSec() int32 {
	if x != nil && x.CacheTimeSec != nil {
		return *x.CacheTimeSec
	}
	return Default_SurfacerConf_CacheTimeSec
}

func (x *SurfacerConf) GetDisable() bool {
	if x != nil && x.Disable != nil {
		return *x.Disable
	}
	return false
}

var File_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x73, 0x75, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xfd, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x12, 0x29, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x36, 0x30, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x2d, 0x0a, 0x0f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x04, 0x34, 0x33, 0x32, 0x30, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x35, 0x0a, 0x15, 0x6d, 0x61,
	0x78, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x32, 0x30, 0x52, 0x12, 0x6d,
	0x61, 0x78, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x07,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0e,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_goTypes = []interface{}{
	(*SurfacerConf)(nil), // 0: cloudprober.surfacer.probestatus.SurfacerConf
}
var file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurfacerConf); i {
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
			RawDescriptor: file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_surfacers_probestatus_proto_config_proto_depIdxs = nil
}
