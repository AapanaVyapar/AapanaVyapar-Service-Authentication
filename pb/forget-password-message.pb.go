// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: forget-password-message.proto

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

type ForgetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,41,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	PhoNo  string `protobuf:"bytes,42,opt,name=phoNo,proto3" json:"phoNo,omitempty"`
}

func (x *ForgetPasswordRequest) Reset() {
	*x = ForgetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordRequest) ProtoMessage() {}

func (x *ForgetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{0}
}

func (x *ForgetPasswordRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *ForgetPasswordRequest) GetPhoNo() string {
	if x != nil {
		return x.PhoNo
	}
	return ""
}

type ForgetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseData *ResponseData `protobuf:"bytes,43,opt,name=responseData,proto3" json:"responseData,omitempty"`
}

func (x *ForgetPasswordResponse) Reset() {
	*x = ForgetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgetPasswordResponse) ProtoMessage() {}

func (x *ForgetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgetPasswordResponse.ProtoReflect.Descriptor instead.
func (*ForgetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{1}
}

func (x *ForgetPasswordResponse) GetResponseData() *ResponseData {
	if x != nil {
		return x.ResponseData
	}
	return nil
}

type ConformForgetPasswordOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,44,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	Otp    string `protobuf:"bytes,45,opt,name=otp,proto3" json:"otp,omitempty"`
	Token  string `protobuf:"bytes,46,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ConformForgetPasswordOTPRequest) Reset() {
	*x = ConformForgetPasswordOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConformForgetPasswordOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformForgetPasswordOTPRequest) ProtoMessage() {}

func (x *ConformForgetPasswordOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConformForgetPasswordOTPRequest.ProtoReflect.Descriptor instead.
func (*ConformForgetPasswordOTPRequest) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{2}
}

func (x *ConformForgetPasswordOTPRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *ConformForgetPasswordOTPRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *ConformForgetPasswordOTPRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ConformForgetPasswordOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPassToken string `protobuf:"bytes,47,opt,name=newPassToken,proto3" json:"newPassToken,omitempty"`
}

func (x *ConformForgetPasswordOTPResponse) Reset() {
	*x = ConformForgetPasswordOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConformForgetPasswordOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformForgetPasswordOTPResponse) ProtoMessage() {}

func (x *ConformForgetPasswordOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConformForgetPasswordOTPResponse.ProtoReflect.Descriptor instead.
func (*ConformForgetPasswordOTPResponse) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{3}
}

func (x *ConformForgetPasswordOTPResponse) GetNewPassToken() string {
	if x != nil {
		return x.NewPassToken
	}
	return ""
}

type SetNewPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey       string `protobuf:"bytes,48,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	NewPassToken string `protobuf:"bytes,49,opt,name=newPassToken,proto3" json:"newPassToken,omitempty"`
	NewPassword  string `protobuf:"bytes,50,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *SetNewPasswordRequest) Reset() {
	*x = SetNewPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNewPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNewPasswordRequest) ProtoMessage() {}

func (x *SetNewPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNewPasswordRequest.ProtoReflect.Descriptor instead.
func (*SetNewPasswordRequest) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{4}
}

func (x *SetNewPasswordRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *SetNewPasswordRequest) GetNewPassToken() string {
	if x != nil {
		return x.NewPassToken
	}
	return ""
}

func (x *SetNewPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type SetNewPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,51,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SetNewPasswordResponse) Reset() {
	*x = SetNewPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forget_password_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNewPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNewPasswordResponse) ProtoMessage() {}

func (x *SetNewPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forget_password_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNewPasswordResponse.ProtoReflect.Descriptor instead.
func (*SetNewPasswordResponse) Descriptor() ([]byte, []int) {
	return file_forget_password_message_proto_rawDescGZIP(), []int{5}
}

func (x *SetNewPasswordResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_forget_password_message_proto protoreflect.FileDescriptor

var file_forget_password_message_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x4e, 0x6f, 0x18,
	0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x4e, 0x6f, 0x22, 0x4b, 0x0a, 0x16,
	0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x61, 0x0a, 0x1f, 0x43, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x2d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x46, 0x0a, 0x20,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x75, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x18, 0x30, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x53,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x30, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61,
	0x72, 0x2e, 0x61, 0x61, 0x70, 0x61, 0x6e, 0x61, 0x76, 0x79, 0x61, 0x70, 0x61, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forget_password_message_proto_rawDescOnce sync.Once
	file_forget_password_message_proto_rawDescData = file_forget_password_message_proto_rawDesc
)

func file_forget_password_message_proto_rawDescGZIP() []byte {
	file_forget_password_message_proto_rawDescOnce.Do(func() {
		file_forget_password_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_forget_password_message_proto_rawDescData)
	})
	return file_forget_password_message_proto_rawDescData
}

var file_forget_password_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_forget_password_message_proto_goTypes = []interface{}{
	(*ForgetPasswordRequest)(nil),            // 0: ForgetPasswordRequest
	(*ForgetPasswordResponse)(nil),           // 1: ForgetPasswordResponse
	(*ConformForgetPasswordOTPRequest)(nil),  // 2: ConformForgetPasswordOTPRequest
	(*ConformForgetPasswordOTPResponse)(nil), // 3: ConformForgetPasswordOTPResponse
	(*SetNewPasswordRequest)(nil),            // 4: SetNewPasswordRequest
	(*SetNewPasswordResponse)(nil),           // 5: SetNewPasswordResponse
	(*ResponseData)(nil),                     // 6: ResponseData
}
var file_forget_password_message_proto_depIdxs = []int32{
	6, // 0: ForgetPasswordResponse.responseData:type_name -> ResponseData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_forget_password_message_proto_init() }
func file_forget_password_message_proto_init() {
	if File_forget_password_message_proto != nil {
		return
	}
	file_common_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_forget_password_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetPasswordRequest); i {
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
		file_forget_password_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgetPasswordResponse); i {
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
		file_forget_password_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConformForgetPasswordOTPRequest); i {
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
		file_forget_password_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConformForgetPasswordOTPResponse); i {
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
		file_forget_password_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNewPasswordRequest); i {
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
		file_forget_password_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNewPasswordResponse); i {
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
			RawDescriptor: file_forget_password_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_forget_password_message_proto_goTypes,
		DependencyIndexes: file_forget_password_message_proto_depIdxs,
		MessageInfos:      file_forget_password_message_proto_msgTypes,
	}.Build()
	File_forget_password_message_proto = out.File
	file_forget_password_message_proto_rawDesc = nil
	file_forget_password_message_proto_goTypes = nil
	file_forget_password_message_proto_depIdxs = nil
}