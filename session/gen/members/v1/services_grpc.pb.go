// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: members/v1/services.proto

package membersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MemberSession_GetEncryptedSessionId_FullMethodName = "/proto.members.v1.MemberSession/GetEncryptedSessionId"
)

// MemberSessionClient is the client API for MemberSession service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberSessionClient interface {
	GetEncryptedSessionId(ctx context.Context, in *EncryptedStringRequest, opts ...grpc.CallOption) (*EncryptedStringResponse, error)
}

type memberSessionClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberSessionClient(cc grpc.ClientConnInterface) MemberSessionClient {
	return &memberSessionClient{cc}
}

func (c *memberSessionClient) GetEncryptedSessionId(ctx context.Context, in *EncryptedStringRequest, opts ...grpc.CallOption) (*EncryptedStringResponse, error) {
	out := new(EncryptedStringResponse)
	err := c.cc.Invoke(ctx, MemberSession_GetEncryptedSessionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberSessionServer is the server API for MemberSession service.
// All implementations should embed UnimplementedMemberSessionServer
// for forward compatibility
type MemberSessionServer interface {
	GetEncryptedSessionId(context.Context, *EncryptedStringRequest) (*EncryptedStringResponse, error)
}

// UnimplementedMemberSessionServer should be embedded to have forward compatible implementations.
type UnimplementedMemberSessionServer struct {
}

func (UnimplementedMemberSessionServer) GetEncryptedSessionId(context.Context, *EncryptedStringRequest) (*EncryptedStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncryptedSessionId not implemented")
}

// UnsafeMemberSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberSessionServer will
// result in compilation errors.
type UnsafeMemberSessionServer interface {
	mustEmbedUnimplementedMemberSessionServer()
}

func RegisterMemberSessionServer(s grpc.ServiceRegistrar, srv MemberSessionServer) {
	s.RegisterService(&MemberSession_ServiceDesc, srv)
}

func _MemberSession_GetEncryptedSessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptedStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberSessionServer).GetEncryptedSessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemberSession_GetEncryptedSessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberSessionServer).GetEncryptedSessionId(ctx, req.(*EncryptedStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberSession_ServiceDesc is the grpc.ServiceDesc for MemberSession service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberSession_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.members.v1.MemberSession",
	HandlerType: (*MemberSessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEncryptedSessionId",
			Handler:    _MemberSession_GetEncryptedSessionId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "members/v1/services.proto",
}
