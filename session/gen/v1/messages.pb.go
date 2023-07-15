// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/messages.proto

package protov1

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

type AWSConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arn    string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *AWSConfigRequest) Reset() {
	*x = AWSConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSConfigRequest) ProtoMessage() {}

func (x *AWSConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSConfigRequest.ProtoReflect.Descriptor instead.
func (*AWSConfigRequest) Descriptor() ([]byte, []int) {
	return file_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AWSConfigRequest) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *AWSConfigRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type AWSConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwsSession []byte `protobuf:"bytes,1,opt,name=awsSession,proto3" json:"awsSession,omitempty"`
}

func (x *AWSConfigResponse) Reset() {
	*x = AWSConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSConfigResponse) ProtoMessage() {}

func (x *AWSConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSConfigResponse.ProtoReflect.Descriptor instead.
func (*AWSConfigResponse) Descriptor() ([]byte, []int) {
	return file_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *AWSConfigResponse) GetAwsSession() []byte {
	if x != nil {
		return x.AwsSession
	}
	return nil
}

type EncryptedStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
}

func (x *EncryptedStringRequest) Reset() {
	*x = EncryptedStringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedStringRequest) ProtoMessage() {}

func (x *EncryptedStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedStringRequest.ProtoReflect.Descriptor instead.
func (*EncryptedStringRequest) Descriptor() ([]byte, []int) {
	return file_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptedStringRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type EncryptedStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedSessionId string `protobuf:"bytes,1,opt,name=encryptedSessionId,proto3" json:"encryptedSessionId,omitempty"`
}

func (x *EncryptedStringResponse) Reset() {
	*x = EncryptedStringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedStringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedStringResponse) ProtoMessage() {}

func (x *EncryptedStringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedStringResponse.ProtoReflect.Descriptor instead.
func (*EncryptedStringResponse) Descriptor() ([]byte, []int) {
	return file_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptedStringResponse) GetEncryptedSessionId() string {
	if x != nil {
		return x.EncryptedSessionId
	}
	return ""
}

var File_v1_messages_proto protoreflect.FileDescriptor

var file_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x3c, 0x0a,
	0x10, 0x41, 0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x72, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x11, 0x41,
	0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x77, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x36, 0x0a, 0x16, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x17, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x42, 0xa2, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa,
	0x02, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_messages_proto_rawDescOnce sync.Once
	file_v1_messages_proto_rawDescData = file_v1_messages_proto_rawDesc
)

func file_v1_messages_proto_rawDescGZIP() []byte {
	file_v1_messages_proto_rawDescOnce.Do(func() {
		file_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_messages_proto_rawDescData)
	})
	return file_v1_messages_proto_rawDescData
}

var file_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_messages_proto_goTypes = []interface{}{
	(*AWSConfigRequest)(nil),        // 0: proto.v1.AWSConfigRequest
	(*AWSConfigResponse)(nil),       // 1: proto.v1.AWSConfigResponse
	(*EncryptedStringRequest)(nil),  // 2: proto.v1.EncryptedStringRequest
	(*EncryptedStringResponse)(nil), // 3: proto.v1.EncryptedStringResponse
}
var file_v1_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_messages_proto_init() }
func file_v1_messages_proto_init() {
	if File_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSConfigRequest); i {
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
		file_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSConfigResponse); i {
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
		file_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedStringRequest); i {
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
		file_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedStringResponse); i {
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
			RawDescriptor: file_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_messages_proto_goTypes,
		DependencyIndexes: file_v1_messages_proto_depIdxs,
		MessageInfos:      file_v1_messages_proto_msgTypes,
	}.Build()
	File_v1_messages_proto = out.File
	file_v1_messages_proto_rawDesc = nil
	file_v1_messages_proto_goTypes = nil
	file_v1_messages_proto_depIdxs = nil
}
