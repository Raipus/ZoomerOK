// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: authorization.proto

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

type AuthorizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthorizationRequest) Reset() {
	*x = AuthorizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationRequest) ProtoMessage() {}

func (x *AuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthorizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AuthorizationResponse) Reset() {
	*x = AuthorizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationResponse) ProtoMessage() {}

func (x *AuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizationResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

var File_authorization_proto protoreflect.FileDescriptor

var file_authorization_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x2c, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x31, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x55, 0x0a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authorization_proto_rawDescOnce sync.Once
	file_authorization_proto_rawDescData = file_authorization_proto_rawDesc
)

func file_authorization_proto_rawDescGZIP() []byte {
	file_authorization_proto_rawDescOnce.Do(func() {
		file_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_authorization_proto_rawDescData)
	})
	return file_authorization_proto_rawDescData
}

var file_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_authorization_proto_goTypes = []interface{}{
	(*AuthorizationRequest)(nil),  // 0: pb.AuthorizationRequest
	(*AuthorizationResponse)(nil), // 1: pb.AuthorizationResponse
}
var file_authorization_proto_depIdxs = []int32{
	0, // 0: pb.Authorization.Authorization:input_type -> pb.AuthorizationRequest
	1, // 1: pb.Authorization.Authorization:output_type -> pb.AuthorizationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authorization_proto_init() }
func file_authorization_proto_init() {
	if File_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationRequest); i {
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
		file_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationResponse); i {
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
			RawDescriptor: file_authorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authorization_proto_goTypes,
		DependencyIndexes: file_authorization_proto_depIdxs,
		MessageInfos:      file_authorization_proto_msgTypes,
	}.Build()
	File_authorization_proto = out.File
	file_authorization_proto_rawDesc = nil
	file_authorization_proto_goTypes = nil
	file_authorization_proto_depIdxs = nil
}
