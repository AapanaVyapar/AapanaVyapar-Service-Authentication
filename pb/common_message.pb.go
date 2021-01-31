// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: common_message.proto

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

type ProblemCode int32

const (
	ProblemCode_NoUserNameIsProvided                  ProblemCode = 0
	ProblemCode_NoPhoneNumberIsProvided               ProblemCode = 1
	ProblemCode_NoPasswordIsProvided                  ProblemCode = 2
	ProblemCode_NoPinCodeIsProvided                   ProblemCode = 4
	ProblemCode_NoEmailIsProvided                     ProblemCode = 5
	ProblemCode_UserAlreadyExist                      ProblemCode = 6
	ProblemCode_UserAlreadyExistWithSameContactNumber ProblemCode = 7
	ProblemCode_InvalidEmailAddress                   ProblemCode = 8
	ProblemCode_InvalidPasswordLength                 ProblemCode = 9
	ProblemCode_InvalidPhoneNumber                    ProblemCode = 10
	ProblemCode_InvalidPinCode                        ProblemCode = 11
	ProblemCode_UnableToSendOTP                       ProblemCode = 12
	ProblemCode_InvalidPassword                       ProblemCode = 13
	ProblemCode_InvalidUserCredentials                ProblemCode = 14
	ProblemCode_UserNotExist                          ProblemCode = 88
)

// Enum value maps for ProblemCode.
var (
	ProblemCode_name = map[int32]string{
		0:  "NoUserNameIsProvided",
		1:  "NoPhoneNumberIsProvided",
		2:  "NoPasswordIsProvided",
		4:  "NoPinCodeIsProvided",
		5:  "NoEmailIsProvided",
		6:  "UserAlreadyExist",
		7:  "UserAlreadyExistWithSameContactNumber",
		8:  "InvalidEmailAddress",
		9:  "InvalidPasswordLength",
		10: "InvalidPhoneNumber",
		11: "InvalidPinCode",
		12: "UnableToSendOTP",
		13: "InvalidPassword",
		14: "InvalidUserCredentials",
		88: "UserNotExist",
	}
	ProblemCode_value = map[string]int32{
		"NoUserNameIsProvided":                  0,
		"NoPhoneNumberIsProvided":               1,
		"NoPasswordIsProvided":                  2,
		"NoPinCodeIsProvided":                   4,
		"NoEmailIsProvided":                     5,
		"UserAlreadyExist":                      6,
		"UserAlreadyExistWithSameContactNumber": 7,
		"InvalidEmailAddress":                   8,
		"InvalidPasswordLength":                 9,
		"InvalidPhoneNumber":                    10,
		"InvalidPinCode":                        11,
		"UnableToSendOTP":                       12,
		"InvalidPassword":                       13,
		"InvalidUserCredentials":                14,
		"UserNotExist":                          88,
	}
)

func (x ProblemCode) Enum() *ProblemCode {
	p := new(ProblemCode)
	*p = x
	return p
}

func (x ProblemCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProblemCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_message_proto_enumTypes[0].Descriptor()
}

func (ProblemCode) Type() protoreflect.EnumType {
	return &file_common_message_proto_enumTypes[0]
}

func (x ProblemCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProblemCode.Descriptor instead.
func (ProblemCode) EnumDescriptor() ([]byte, []int) {
	return file_common_message_proto_rawDescGZIP(), []int{0}
}

type OTPResponse int32

const (
	OTPResponse_NotOk OTPResponse = 0
	OTPResponse_OK    OTPResponse = 1
)

// Enum value maps for OTPResponse.
var (
	OTPResponse_name = map[int32]string{
		0: "NotOk",
		1: "OK",
	}
	OTPResponse_value = map[string]int32{
		"NotOk": 0,
		"OK":    1,
	}
)

func (x OTPResponse) Enum() *OTPResponse {
	p := new(OTPResponse)
	*p = x
	return p
}

func (x OTPResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OTPResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_common_message_proto_enumTypes[1].Descriptor()
}

func (OTPResponse) Type() protoreflect.EnumType {
	return &file_common_message_proto_enumTypes[1]
}

func (x OTPResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OTPResponse.Descriptor instead.
func (OTPResponse) EnumDescriptor() ([]byte, []int) {
	return file_common_message_proto_rawDescGZIP(), []int{1}
}

type ResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,30,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken string `protobuf:"bytes,31,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *ResponseData) Reset() {
	*x = ResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseData) ProtoMessage() {}

func (x *ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_common_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseData.ProtoReflect.Descriptor instead.
func (*ResponseData) Descriptor() ([]byte, []int) {
	return file_common_message_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResponseData) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_common_message_proto protoreflect.FileDescriptor

var file_common_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x2a, 0x87, 0x03, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x14, 0x4e, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x6f,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x73, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x64, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x4e, 0x6f, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x10,
	0x02, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x50, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x6f,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x10,
	0x05, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x06, 0x12, 0x29, 0x0a, 0x25, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x53,
	0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x49,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x0a, 0x12, 0x12,
	0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x65,
	0x6e, 0x64, 0x4f, 0x54, 0x50, 0x10, 0x0c, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x58, 0x2a, 0x20, 0x0a, 0x0b, 0x4f, 0x54,
	0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x6f, 0x74,
	0x4f, 0x6b, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_message_proto_rawDescOnce sync.Once
	file_common_message_proto_rawDescData = file_common_message_proto_rawDesc
)

func file_common_message_proto_rawDescGZIP() []byte {
	file_common_message_proto_rawDescOnce.Do(func() {
		file_common_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_message_proto_rawDescData)
	})
	return file_common_message_proto_rawDescData
}

var file_common_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_message_proto_goTypes = []interface{}{
	(ProblemCode)(0),     // 0: ProblemCode
	(OTPResponse)(0),     // 1: OTPResponse
	(*ResponseData)(nil), // 2: ResponseData
}
var file_common_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_message_proto_init() }
func file_common_message_proto_init() {
	if File_common_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseData); i {
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
			RawDescriptor: file_common_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_message_proto_goTypes,
		DependencyIndexes: file_common_message_proto_depIdxs,
		EnumInfos:         file_common_message_proto_enumTypes,
		MessageInfos:      file_common_message_proto_msgTypes,
	}.Build()
	File_common_message_proto = out.File
	file_common_message_proto_rawDesc = nil
	file_common_message_proto_goTypes = nil
	file_common_message_proto_depIdxs = nil
}
