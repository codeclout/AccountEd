// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aws/v1/services.proto

package awsv1

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

var File_aws_v1_services_proto protoreflect.FileDescriptor

var file_aws_v1_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x61, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x77, 0x0a, 0x18,
	0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41,
	0x57, 0x53, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x77, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb0, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75,
	0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x77, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x41, 0x58, 0xaa, 0x02,
	0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_aws_v1_services_proto_goTypes = []interface{}{
	(*AWSConfigRequest)(nil),  // 0: proto.aws.v1.AWSConfigRequest
	(*AWSConfigResponse)(nil), // 1: proto.aws.v1.AWSConfigResponse
}
var file_aws_v1_services_proto_depIdxs = []int32{
	0, // 0: proto.aws.v1.AWSResourceClientService.GetAWSSessionCredentials:input_type -> proto.aws.v1.AWSConfigRequest
	1, // 1: proto.aws.v1.AWSResourceClientService.GetAWSSessionCredentials:output_type -> proto.aws.v1.AWSConfigResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aws_v1_services_proto_init() }
func file_aws_v1_services_proto_init() {
	if File_aws_v1_services_proto != nil {
		return
	}
	file_aws_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aws_v1_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aws_v1_services_proto_goTypes,
		DependencyIndexes: file_aws_v1_services_proto_depIdxs,
	}.Build()
	File_aws_v1_services_proto = out.File
	file_aws_v1_services_proto_rawDesc = nil
	file_aws_v1_services_proto_goTypes = nil
	file_aws_v1_services_proto_depIdxs = nil
}
