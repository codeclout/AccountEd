// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: members/v1/messages.proto

package membersv1

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

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_members_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_members_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValidToken bool `protobuf:"varint,1,opt,name=isValidToken,proto3" json:"isValidToken,omitempty"`
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_members_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_members_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_members_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateTokenResponse) GetIsValidToken() bool {
	if x != nil {
		return x.IsValidToken
	}
	return false
}

var File_members_v1_messages_proto protoreflect.FileDescriptor

var file_members_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a,
	0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b, 0x0a, 0x15, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0xc8, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x64, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_members_v1_messages_proto_rawDescOnce sync.Once
	file_members_v1_messages_proto_rawDescData = file_members_v1_messages_proto_rawDesc
)

func file_members_v1_messages_proto_rawDescGZIP() []byte {
	file_members_v1_messages_proto_rawDescOnce.Do(func() {
		file_members_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_members_v1_messages_proto_rawDescData)
	})
	return file_members_v1_messages_proto_rawDescData
}

var file_members_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_members_v1_messages_proto_goTypes = []interface{}{
	(*ValidateTokenRequest)(nil),  // 0: proto.members.v1.ValidateTokenRequest
	(*ValidateTokenResponse)(nil), // 1: proto.members.v1.ValidateTokenResponse
}
var file_members_v1_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_members_v1_messages_proto_init() }
func file_members_v1_messages_proto_init() {
	if File_members_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_members_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
		file_members_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenResponse); i {
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
			RawDescriptor: file_members_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_members_v1_messages_proto_goTypes,
		DependencyIndexes: file_members_v1_messages_proto_depIdxs,
		MessageInfos:      file_members_v1_messages_proto_msgTypes,
	}.Build()
	File_members_v1_messages_proto = out.File
	file_members_v1_messages_proto_rawDesc = nil
	file_members_v1_messages_proto_goTypes = nil
	file_members_v1_messages_proto_depIdxs = nil
}
