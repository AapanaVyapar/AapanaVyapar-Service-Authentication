// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: signup_message.proto

package pb

import (
	duration "github.com/golang/protobuf/ptypes/duration"
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

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,25,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,26,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNo  string `protobuf:"bytes,27,opt,name=phoneNo,proto3" json:"phoneNo,omitempty"`
	Email    string `protobuf:"bytes,28,opt,name=email,proto3" json:"email,omitempty"`
	PinCode  string `protobuf:"bytes,29,opt,name=pinCode,proto3" json:"pinCode,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*SignUpResponse_ResponseData
	//	*SignUpResponse_Code
	Data       isSignUpResponse_Data `protobuf_oneof:"data"`
	Authorized bool                  `protobuf:"varint,50,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{1}
}

func (m *SignUpResponse) GetData() isSignUpResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SignUpResponse) GetResponseData() *ResponseData {
	if x, ok := x.GetData().(*SignUpResponse_ResponseData); ok {
		return x.ResponseData
	}
	return nil
}

func (x *SignUpResponse) GetCode() ProblemCode {
	if x, ok := x.GetData().(*SignUpResponse_Code); ok {
		return x.Code
	}
	return ProblemCode_NoUserNameIsProvided
}

func (x *SignUpResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type isSignUpResponse_Data interface {
	isSignUpResponse_Data()
}

type SignUpResponse_ResponseData struct {
	ResponseData *ResponseData `protobuf:"bytes,40,opt,name=responseData,proto3,oneof"`
}

type SignUpResponse_Code struct {
	Code ProblemCode `protobuf:"varint,32,opt,name=code,proto3,enum=ProblemCode,oneof"`
}

func (*SignUpResponse_ResponseData) isSignUpResponse_Data() {}

func (*SignUpResponse_Code) isSignUpResponse_Data() {}

type ContactConformationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,33,opt,name=token,proto3" json:"token,omitempty"`
	Otp   string `protobuf:"bytes,34,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *ContactConformationRequest) Reset() {
	*x = ContactConformationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactConformationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactConformationRequest) ProtoMessage() {}

func (x *ContactConformationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactConformationRequest.ProtoReflect.Descriptor instead.
func (*ContactConformationRequest) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{2}
}

func (x *ContactConformationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ContactConformationRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type ContactConformationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,35,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken string `protobuf:"bytes,36,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *ContactConformationResponse) Reset() {
	*x = ContactConformationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactConformationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactConformationResponse) ProtoMessage() {}

func (x *ContactConformationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactConformationResponse.ProtoReflect.Descriptor instead.
func (*ContactConformationResponse) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{3}
}

func (x *ContactConformationResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ContactConformationResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ResendOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,54,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResendOTPRequest) Reset() {
	*x = ResendOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendOTPRequest) ProtoMessage() {}

func (x *ResendOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendOTPRequest.ProtoReflect.Descriptor instead.
func (*ResendOTPRequest) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{4}
}

func (x *ResendOTPRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ResendOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response                 OTPResponse        `protobuf:"varint,55,opt,name=response,proto3,enum=OTPResponse" json:"response,omitempty"`
	TimeToWaitForNextRequest *duration.Duration `protobuf:"bytes,56,opt,name=timeToWaitForNextRequest,proto3" json:"timeToWaitForNextRequest,omitempty"`
}

func (x *ResendOTPResponse) Reset() {
	*x = ResendOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_signup_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResendOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResendOTPResponse) ProtoMessage() {}

func (x *ResendOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_signup_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResendOTPResponse.ProtoReflect.Descriptor instead.
func (*ResendOTPResponse) Descriptor() ([]byte, []int) {
	return file_signup_message_proto_rawDescGZIP(), []int{5}
}

func (x *ResendOTPResponse) GetResponse() OTPResponse {
	if x != nil {
		return x.Response
	}
	return OTPResponse_NotOk
}

func (x *ResendOTPResponse) GetTimeToWaitForNextRequest() *duration.Duration {
	if x != nil {
		return x.TimeToWaitForNextRequest
	}
	return nil
}

var File_signup_message_proto protoreflect.FileDescriptor

var file_signup_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x6f, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x91, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x44, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18,
	0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x57, 0x0a, 0x1b, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x28, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x36, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x94, 0x01,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x65, 0x6e, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x37, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x18, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x4e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x38, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x6f, 0x57, 0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_signup_message_proto_rawDescOnce sync.Once
	file_signup_message_proto_rawDescData = file_signup_message_proto_rawDesc
)

func file_signup_message_proto_rawDescGZIP() []byte {
	file_signup_message_proto_rawDescOnce.Do(func() {
		file_signup_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_signup_message_proto_rawDescData)
	})
	return file_signup_message_proto_rawDescData
}

var file_signup_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_signup_message_proto_goTypes = []interface{}{
	(*SignUpRequest)(nil),               // 0: SignUpRequest
	(*SignUpResponse)(nil),              // 1: SignUpResponse
	(*ContactConformationRequest)(nil),  // 2: ContactConformationRequest
	(*ContactConformationResponse)(nil), // 3: ContactConformationResponse
	(*ResendOTPRequest)(nil),            // 4: ResendOTPRequest
	(*ResendOTPResponse)(nil),           // 5: ResendOTPResponse
	(*ResponseData)(nil),                // 6: ResponseData
	(ProblemCode)(0),                    // 7: ProblemCode
	(OTPResponse)(0),                    // 8: OTPResponse
	(*duration.Duration)(nil),           // 9: google.protobuf.Duration
}
var file_signup_message_proto_depIdxs = []int32{
	6, // 0: SignUpResponse.responseData:type_name -> ResponseData
	7, // 1: SignUpResponse.code:type_name -> ProblemCode
	8, // 2: ResendOTPResponse.response:type_name -> OTPResponse
	9, // 3: ResendOTPResponse.timeToWaitForNextRequest:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_signup_message_proto_init() }
func file_signup_message_proto_init() {
	if File_signup_message_proto != nil {
		return
	}
	file_common_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_signup_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_signup_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_signup_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactConformationRequest); i {
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
		file_signup_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactConformationResponse); i {
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
		file_signup_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendOTPRequest); i {
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
		file_signup_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResendOTPResponse); i {
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
	file_signup_message_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SignUpResponse_ResponseData)(nil),
		(*SignUpResponse_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_signup_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_signup_message_proto_goTypes,
		DependencyIndexes: file_signup_message_proto_depIdxs,
		MessageInfos:      file_signup_message_proto_msgTypes,
	}.Build()
	File_signup_message_proto = out.File
	file_signup_message_proto_rawDesc = nil
	file_signup_message_proto_goTypes = nil
	file_signup_message_proto_depIdxs = nil
}
