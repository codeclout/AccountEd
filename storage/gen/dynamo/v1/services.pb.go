// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: dynamo/v1/services.proto

package dynamov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_dynamo_v1_services_proto protoreflect.FileDescriptor

var file_dynamo_v1_services_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd4, 0x02, 0x0a, 0x16, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f,
	0x44, 0x42, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x55, 0x0a, 0x0a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc1, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x45, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x44, 0x58, 0xaa, 0x02, 0x0f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_dynamo_v1_services_proto_goTypes = []interface{}{
	(*FetchTokenRequest)(nil),                  // 0: proto.dynamo.v1.FetchTokenRequest
	(*TokenStoreRequest)(nil),                  // 1: proto.dynamo.v1.TokenStoreRequest
	(*StoreConfirmedRegistrationRequest)(nil),  // 2: proto.dynamo.v1.StoreConfirmedRegistrationRequest
	(*FetchTokenResponse)(nil),                 // 3: proto.dynamo.v1.FetchTokenResponse
	(*TokenStoreResponse)(nil),                 // 4: proto.dynamo.v1.TokenStoreResponse
	(*StoreConfirmedRegistrationResponse)(nil), // 5: proto.dynamo.v1.StoreConfirmedRegistrationResponse
}
var file_dynamo_v1_services_proto_depIdxs = []int32{
	0, // 0: proto.dynamo.v1.DynamoDBStorageService.FetchToken:input_type -> proto.dynamo.v1.FetchTokenRequest
	1, // 1: proto.dynamo.v1.DynamoDBStorageService.StorePublicToken:input_type -> proto.dynamo.v1.TokenStoreRequest
	2, // 2: proto.dynamo.v1.DynamoDBStorageService.StoreConfirmedRegistration:input_type -> proto.dynamo.v1.StoreConfirmedRegistrationRequest
	3, // 3: proto.dynamo.v1.DynamoDBStorageService.FetchToken:output_type -> proto.dynamo.v1.FetchTokenResponse
	4, // 4: proto.dynamo.v1.DynamoDBStorageService.StorePublicToken:output_type -> proto.dynamo.v1.TokenStoreResponse
	5, // 5: proto.dynamo.v1.DynamoDBStorageService.StoreConfirmedRegistration:output_type -> proto.dynamo.v1.StoreConfirmedRegistrationResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dynamo_v1_services_proto_init() }
func file_dynamo_v1_services_proto_init() {
	if File_dynamo_v1_services_proto != nil {
		return
	}
	file_dynamo_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dynamo_v1_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dynamo_v1_services_proto_goTypes,
		DependencyIndexes: file_dynamo_v1_services_proto_depIdxs,
	}.Build()
	File_dynamo_v1_services_proto = out.File
	file_dynamo_v1_services_proto_rawDesc = nil
	file_dynamo_v1_services_proto_goTypes = nil
	file_dynamo_v1_services_proto_depIdxs = nil
}
