// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/services.proto

package protov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/codeclout/AccountEd/pkg/notifications/gen/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EmailNotificationServiceName is the fully-qualified name of the EmailNotificationService service.
	EmailNotificationServiceName = "proto.v1.EmailNotificationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EmailNotificationServiceValidateEmailAddressProcedure is the fully-qualified name of the
	// EmailNotificationService's ValidateEmailAddress RPC.
	EmailNotificationServiceValidateEmailAddressProcedure = "/proto.v1.EmailNotificationService/ValidateEmailAddress"
)

// EmailNotificationServiceClient is a client for the proto.v1.EmailNotificationService service.
type EmailNotificationServiceClient interface {
	ValidateEmailAddress(context.Context, *connect_go.Request[v1.ValidateEmailAddressRequest]) (*connect_go.Response[v1.ValidateEmailAddressResponse], error)
}

// NewEmailNotificationServiceClient constructs a client for the proto.v1.EmailNotificationService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEmailNotificationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EmailNotificationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &emailNotificationServiceClient{
		validateEmailAddress: connect_go.NewClient[v1.ValidateEmailAddressRequest, v1.ValidateEmailAddressResponse](
			httpClient,
			baseURL+EmailNotificationServiceValidateEmailAddressProcedure,
			opts...,
		),
	}
}

// emailNotificationServiceClient implements EmailNotificationServiceClient.
type emailNotificationServiceClient struct {
	validateEmailAddress *connect_go.Client[v1.ValidateEmailAddressRequest, v1.ValidateEmailAddressResponse]
}

// ValidateEmailAddress calls proto.v1.EmailNotificationService.ValidateEmailAddress.
func (c *emailNotificationServiceClient) ValidateEmailAddress(ctx context.Context, req *connect_go.Request[v1.ValidateEmailAddressRequest]) (*connect_go.Response[v1.ValidateEmailAddressResponse], error) {
	return c.validateEmailAddress.CallUnary(ctx, req)
}

// EmailNotificationServiceHandler is an implementation of the proto.v1.EmailNotificationService
// service.
type EmailNotificationServiceHandler interface {
	ValidateEmailAddress(context.Context, *connect_go.Request[v1.ValidateEmailAddressRequest]) (*connect_go.Response[v1.ValidateEmailAddressResponse], error)
}

// NewEmailNotificationServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEmailNotificationServiceHandler(svc EmailNotificationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(EmailNotificationServiceValidateEmailAddressProcedure, connect_go.NewUnaryHandler(
		EmailNotificationServiceValidateEmailAddressProcedure,
		svc.ValidateEmailAddress,
		opts...,
	))
	return "/proto.v1.EmailNotificationService/", mux
}

// UnimplementedEmailNotificationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEmailNotificationServiceHandler struct{}

func (UnimplementedEmailNotificationServiceHandler) ValidateEmailAddress(context.Context, *connect_go.Request[v1.ValidateEmailAddressRequest]) (*connect_go.Response[v1.ValidateEmailAddressResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("proto.v1.EmailNotificationService.ValidateEmailAddress is not implemented"))
}
