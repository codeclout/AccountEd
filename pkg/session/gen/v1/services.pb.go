// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/services.proto

package protov1

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

var File_v1_services_proto protoreflect.FileDescriptor

var file_v1_services_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x6d, 0x0a, 0x17, 0x41, 0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x41, 0x57, 0x53, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57,
	0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x99, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x42, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45,
	0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_v1_services_proto_goTypes = []interface{}{
	(*AWSConfigRequest)(nil),  // 0: proto.v1.AWSConfigRequest
	(*AWSConfigResponse)(nil), // 1: proto.v1.AWSConfigResponse
}
var file_v1_services_proto_depIdxs = []int32{
	0, // 0: proto.v1.AWSConfigurationService.GetAWSRoleConfiguration:input_type -> proto.v1.AWSConfigRequest
	1, // 1: proto.v1.AWSConfigurationService.GetAWSRoleConfiguration:output_type -> proto.v1.AWSConfigResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_services_proto_init() }
func file_v1_services_proto_init() {
	if File_v1_services_proto != nil {
		return
	}
	file_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_services_proto_goTypes,
		DependencyIndexes: file_v1_services_proto_depIdxs,
	}.Build()
	File_v1_services_proto = out.File
	file_v1_services_proto_rawDesc = nil
	file_v1_services_proto_goTypes = nil
	file_v1_services_proto_depIdxs = nil
}
