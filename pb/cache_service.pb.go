// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: cache_service.proto

package pb

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

type RegCacheRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseURL string `protobuf:"bytes,1,opt,name=baseURL,proto3" json:"baseURL,omitempty"`
}

func (x *RegCacheRequest) Reset() {
	*x = RegCacheRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegCacheRequest) ProtoMessage() {}

func (x *RegCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegCacheRequest.ProtoReflect.Descriptor instead.
func (*RegCacheRequest) Descriptor() ([]byte, []int) {
	return file_cache_service_proto_rawDescGZIP(), []int{0}
}

func (x *RegCacheRequest) GetBaseURL() string {
	if x != nil {
		return x.BaseURL
	}
	return ""
}

type RegCacheResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegCacheResponse) Reset() {
	*x = RegCacheResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegCacheResponse) ProtoMessage() {}

func (x *RegCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegCacheResponse.ProtoReflect.Descriptor instead.
func (*RegCacheResponse) Descriptor() ([]byte, []int) {
	return file_cache_service_proto_rawDescGZIP(), []int{1}
}

func (x *RegCacheResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type NodesAdders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodesAdders []string `protobuf:"bytes,1,rep,name=NodesAdders,proto3" json:"NodesAdders,omitempty"`
}

func (x *NodesAdders) Reset() {
	*x = NodesAdders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodesAdders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodesAdders) ProtoMessage() {}

func (x *NodesAdders) ProtoReflect() protoreflect.Message {
	mi := &file_cache_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodesAdders.ProtoReflect.Descriptor instead.
func (*NodesAdders) Descriptor() ([]byte, []int) {
	return file_cache_service_proto_rawDescGZIP(), []int{2}
}

func (x *NodesAdders) GetNodesAdders() []string {
	if x != nil {
		return x.NodesAdders
	}
	return nil
}

var File_cache_service_proto protoreflect.FileDescriptor

var file_cache_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x73, 0x65, 0x55, 0x52, 0x4c, 0x22, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x2f, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x41, 0x64, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x41, 0x64, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x41, 0x64,
	0x64, 0x65, 0x72, 0x73, 0x32, 0xcb, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x41, 0x64, 0x64, 0x65, 0x72, 0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x72,
	0x65, 0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x2e, 0x52, 0x65,
	0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x52, 0x65, 0x67, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x04,
	0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x06, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x22, 0x00, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x06, 0x2e, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x1a, 0x04, 0x2e, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x12, 0x04, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x04, 0x2e, 0x4b, 0x65, 0x79,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x62, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cache_service_proto_rawDescOnce sync.Once
	file_cache_service_proto_rawDescData = file_cache_service_proto_rawDesc
)

func file_cache_service_proto_rawDescGZIP() []byte {
	file_cache_service_proto_rawDescOnce.Do(func() {
		file_cache_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_service_proto_rawDescData)
	})
	return file_cache_service_proto_rawDescData
}

var file_cache_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cache_service_proto_goTypes = []interface{}{
	(*RegCacheRequest)(nil),  // 0: RegCacheRequest
	(*RegCacheResponse)(nil), // 1: RegCacheResponse
	(*NodesAdders)(nil),      // 2: NodesAdders
	(*Key)(nil),              // 3: Key
	(*Cache)(nil),            // 4: Cache
}
var file_cache_service_proto_depIdxs = []int32{
	0, // 0: CacheService.findAllCacheNode:input_type -> RegCacheRequest
	0, // 1: CacheService.regCacheNode:input_type -> RegCacheRequest
	3, // 2: CacheService.getCache:input_type -> Key
	4, // 3: CacheService.setCache:input_type -> Cache
	3, // 4: CacheService.remove:input_type -> Key
	2, // 5: CacheService.findAllCacheNode:output_type -> NodesAdders
	1, // 6: CacheService.regCacheNode:output_type -> RegCacheResponse
	4, // 7: CacheService.getCache:output_type -> Cache
	3, // 8: CacheService.setCache:output_type -> Key
	3, // 9: CacheService.remove:output_type -> Key
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cache_service_proto_init() }
func file_cache_service_proto_init() {
	if File_cache_service_proto != nil {
		return
	}
	file_cache_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cache_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegCacheRequest); i {
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
		file_cache_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegCacheResponse); i {
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
		file_cache_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodesAdders); i {
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
			RawDescriptor: file_cache_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_service_proto_goTypes,
		DependencyIndexes: file_cache_service_proto_depIdxs,
		MessageInfos:      file_cache_service_proto_msgTypes,
	}.Build()
	File_cache_service_proto = out.File
	file_cache_service_proto_rawDesc = nil
	file_cache_service_proto_goTypes = nil
	file_cache_service_proto_depIdxs = nil
}
