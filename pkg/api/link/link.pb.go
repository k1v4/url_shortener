// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.2
// source: link/link.proto

package linkv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SaveUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullUrl string `protobuf:"bytes,1,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"`
}

func (x *SaveUrlRequest) Reset() {
	*x = SaveUrlRequest{}
	mi := &file_link_link_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUrlRequest) ProtoMessage() {}

func (x *SaveUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_link_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUrlRequest.ProtoReflect.Descriptor instead.
func (*SaveUrlRequest) Descriptor() ([]byte, []int) {
	return file_link_link_proto_rawDescGZIP(), []int{0}
}

func (x *SaveUrlRequest) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

type SaveUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *SaveUrlResponse) Reset() {
	*x = SaveUrlResponse{}
	mi := &file_link_link_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUrlResponse) ProtoMessage() {}

func (x *SaveUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_link_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUrlResponse.ProtoReflect.Descriptor instead.
func (*SaveUrlResponse) Descriptor() ([]byte, []int) {
	return file_link_link_proto_rawDescGZIP(), []int{1}
}

func (x *SaveUrlResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetOriginRequest) Reset() {
	*x = GetOriginRequest{}
	mi := &file_link_link_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginRequest) ProtoMessage() {}

func (x *GetOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_link_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginRequest.ProtoReflect.Descriptor instead.
func (*GetOriginRequest) Descriptor() ([]byte, []int) {
	return file_link_link_proto_rawDescGZIP(), []int{2}
}

func (x *GetOriginRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullUrl string `protobuf:"bytes,1,opt,name=full_url,json=fullUrl,proto3" json:"full_url,omitempty"`
}

func (x *GetOriginResponse) Reset() {
	*x = GetOriginResponse{}
	mi := &file_link_link_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginResponse) ProtoMessage() {}

func (x *GetOriginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_link_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginResponse.ProtoReflect.Descriptor instead.
func (*GetOriginResponse) Descriptor() ([]byte, []int) {
	return file_link_link_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginResponse) GetFullUrl() string {
	if x != nil {
		return x.FullUrl
	}
	return ""
}

var File_link_link_proto protoreflect.FileDescriptor

var file_link_link_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55, 0x72,
	0x6c, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x22, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x55,
	0x72, 0x6c, 0x32, 0xc5, 0x01, 0x0a, 0x0c, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x72,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x2f, 0x7b, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x7d, 0x12, 0x5d, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12,
	0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2f, 0x7b,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x7d, 0x42, 0x16, 0x5a, 0x14, 0x6b, 0x31,
	0x76, 0x63, 0x68, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x3b, 0x6c, 0x69, 0x6e, 0x6b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_link_proto_rawDescOnce sync.Once
	file_link_link_proto_rawDescData = file_link_link_proto_rawDesc
)

func file_link_link_proto_rawDescGZIP() []byte {
	file_link_link_proto_rawDescOnce.Do(func() {
		file_link_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_link_proto_rawDescData)
	})
	return file_link_link_proto_rawDescData
}

var file_link_link_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_link_link_proto_goTypes = []any{
	(*SaveUrlRequest)(nil),    // 0: api.SaveUrlRequest
	(*SaveUrlResponse)(nil),   // 1: api.SaveUrlResponse
	(*GetOriginRequest)(nil),  // 2: api.GetOriginRequest
	(*GetOriginResponse)(nil), // 3: api.GetOriginResponse
}
var file_link_link_proto_depIdxs = []int32{
	0, // 0: api.UrlShortener.SaveUrl:input_type -> api.SaveUrlRequest
	2, // 1: api.UrlShortener.GetOrigin:input_type -> api.GetOriginRequest
	1, // 2: api.UrlShortener.SaveUrl:output_type -> api.SaveUrlResponse
	3, // 3: api.UrlShortener.GetOrigin:output_type -> api.GetOriginResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_link_link_proto_init() }
func file_link_link_proto_init() {
	if File_link_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_link_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_link_proto_goTypes,
		DependencyIndexes: file_link_link_proto_depIdxs,
		MessageInfos:      file_link_link_proto_msgTypes,
	}.Build()
	File_link_link_proto = out.File
	file_link_link_proto_rawDesc = nil
	file_link_link_proto_goTypes = nil
	file_link_link_proto_depIdxs = nil
}
