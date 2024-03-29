// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: email/v1/messages.proto

package emailv1

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

type ValidateEmailAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoCorrect          string `protobuf:"bytes,1,opt,name=autoCorrect,proto3" json:"autoCorrect,omitempty"`
	MemberId             string `protobuf:"bytes,2,opt,name=memberId,proto3" json:"memberId,omitempty"`
	MemberIdPending      bool   `protobuf:"varint,3,opt,name=memberIdPending,proto3" json:"memberIdPending,omitempty"`
	ShouldConfirmAddress bool   `protobuf:"varint,4,opt,name=shouldConfirmAddress,proto3" json:"shouldConfirmAddress,omitempty"`
}

func (x *ValidateEmailAddressResponse) Reset() {
	*x = ValidateEmailAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateEmailAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEmailAddressResponse) ProtoMessage() {}

func (x *ValidateEmailAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEmailAddressResponse.ProtoReflect.Descriptor instead.
func (*ValidateEmailAddressResponse) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateEmailAddressResponse) GetAutoCorrect() string {
	if x != nil {
		return x.AutoCorrect
	}
	return ""
}

func (x *ValidateEmailAddressResponse) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *ValidateEmailAddressResponse) GetMemberIdPending() bool {
	if x != nil {
		return x.MemberIdPending
	}
	return false
}

func (x *ValidateEmailAddressResponse) GetShouldConfirmAddress() bool {
	if x != nil {
		return x.ShouldConfirmAddress
	}
	return false
}

type EmailVerificationPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EmailVerificationPayload) Reset() {
	*x = EmailVerificationPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailVerificationPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailVerificationPayload) ProtoMessage() {}

func (x *EmailVerificationPayload) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailVerificationPayload.ProtoReflect.Descriptor instead.
func (*EmailVerificationPayload) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *EmailVerificationPayload) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EmailVerificationPayload) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type ValidateEmailAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ValidateEmailAddressRequest) Reset() {
	*x = ValidateEmailAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateEmailAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateEmailAddressRequest) ProtoMessage() {}

func (x *ValidateEmailAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateEmailAddressRequest.ProtoReflect.Descriptor instead.
func (*ValidateEmailAddressRequest) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateEmailAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ValidFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ValidFormat) Reset() {
	*x = ValidFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidFormat) ProtoMessage() {}

func (x *ValidFormat) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidFormat.ProtoReflect.Descriptor instead.
func (*ValidFormat) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ValidFormat) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ValidFormat) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type NoReplyEmailNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwsCredentials []byte   `protobuf:"bytes,1,opt,name=awsCredentials,proto3" json:"awsCredentials,omitempty"`
	Domain         string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	FromAddress    string   `protobuf:"bytes,3,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress      []string `protobuf:"bytes,4,rep,name=toAddress,proto3" json:"toAddress,omitempty"`
	Token          string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *NoReplyEmailNotificationRequest) Reset() {
	*x = NoReplyEmailNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReplyEmailNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReplyEmailNotificationRequest) ProtoMessage() {}

func (x *NoReplyEmailNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReplyEmailNotificationRequest.ProtoReflect.Descriptor instead.
func (*NoReplyEmailNotificationRequest) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *NoReplyEmailNotificationRequest) GetAwsCredentials() []byte {
	if x != nil {
		return x.AwsCredentials
	}
	return nil
}

func (x *NoReplyEmailNotificationRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *NoReplyEmailNotificationRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *NoReplyEmailNotificationRequest) GetToAddress() []string {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *NoReplyEmailNotificationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type NoReplyEmailNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
}

func (x *NoReplyEmailNotificationResponse) Reset() {
	*x = NoReplyEmailNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoReplyEmailNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoReplyEmailNotificationResponse) ProtoMessage() {}

func (x *NoReplyEmailNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_email_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoReplyEmailNotificationResponse.ProtoReflect.Descriptor instead.
func (*NoReplyEmailNotificationResponse) Descriptor() ([]byte, []int) {
	return file_email_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *NoReplyEmailNotificationResponse) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

var File_email_v1_messages_proto protoreflect.FileDescriptor

var file_email_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0xba, 0x01, 0x0a, 0x1c, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x6f, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x64, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x44, 0x0a, 0x18, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x1b,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb7,
	0x01, 0x0a, 0x1f, 0x4e, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x77, 0x73, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x40, 0x0a, 0x20, 0x4e, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x42, 0xc0, 0x01, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x74, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x45, 0x64, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x45, 0x58, 0xaa, 0x02, 0x0e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_v1_messages_proto_rawDescOnce sync.Once
	file_email_v1_messages_proto_rawDescData = file_email_v1_messages_proto_rawDesc
)

func file_email_v1_messages_proto_rawDescGZIP() []byte {
	file_email_v1_messages_proto_rawDescOnce.Do(func() {
		file_email_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_v1_messages_proto_rawDescData)
	})
	return file_email_v1_messages_proto_rawDescData
}

var file_email_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_email_v1_messages_proto_goTypes = []interface{}{
	(*ValidateEmailAddressResponse)(nil),     // 0: proto.email.v1.ValidateEmailAddressResponse
	(*EmailVerificationPayload)(nil),         // 1: proto.email.v1.EmailVerificationPayload
	(*ValidateEmailAddressRequest)(nil),      // 2: proto.email.v1.ValidateEmailAddressRequest
	(*ValidFormat)(nil),                      // 3: proto.email.v1.ValidFormat
	(*NoReplyEmailNotificationRequest)(nil),  // 4: proto.email.v1.NoReplyEmailNotificationRequest
	(*NoReplyEmailNotificationResponse)(nil), // 5: proto.email.v1.NoReplyEmailNotificationResponse
}
var file_email_v1_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_email_v1_messages_proto_init() }
func file_email_v1_messages_proto_init() {
	if File_email_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_email_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateEmailAddressResponse); i {
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
		file_email_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailVerificationPayload); i {
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
		file_email_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateEmailAddressRequest); i {
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
		file_email_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidFormat); i {
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
		file_email_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReplyEmailNotificationRequest); i {
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
		file_email_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoReplyEmailNotificationResponse); i {
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
			RawDescriptor: file_email_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_email_v1_messages_proto_goTypes,
		DependencyIndexes: file_email_v1_messages_proto_depIdxs,
		MessageInfos:      file_email_v1_messages_proto_msgTypes,
	}.Build()
	File_email_v1_messages_proto = out.File
	file_email_v1_messages_proto_rawDesc = nil
	file_email_v1_messages_proto_goTypes = nil
	file_email_v1_messages_proto_depIdxs = nil
}
