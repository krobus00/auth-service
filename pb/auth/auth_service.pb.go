// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/auth/auth_service.proto

package auth

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

var File_pb_auth_auth_service_proto protoreflect.FileDescriptor

var file_pb_auth_auth_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x02, 0x0a, 0x0b, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x48, 0x61,
	0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70,
	0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pb_auth_auth_service_proto_goTypes = []interface{}{
	(*GetUserInfoRequest)(nil),  // 0: pb.auth.GetUserInfoRequest
	(*HasAccessRequest)(nil),    // 1: pb.auth.HasAccessRequest
	(*LoginRequest)(nil),        // 2: pb.auth.LoginRequest
	(*RegisterRequest)(nil),     // 3: pb.auth.RegisterRequest
	(*RefreshTokenRequest)(nil), // 4: pb.auth.RefreshTokenRequest
	(*User)(nil),                // 5: pb.auth.User
	(*wrappers.BoolValue)(nil),  // 6: google.protobuf.BoolValue
	(*AuthResponse)(nil),        // 7: pb.auth.AuthResponse
}
var file_pb_auth_auth_service_proto_depIdxs = []int32{
	0, // 0: pb.auth.AuthService.GetUserInfo:input_type -> pb.auth.GetUserInfoRequest
	1, // 1: pb.auth.AuthService.HasAccess:input_type -> pb.auth.HasAccessRequest
	2, // 2: pb.auth.AuthService.Login:input_type -> pb.auth.LoginRequest
	3, // 3: pb.auth.AuthService.Register:input_type -> pb.auth.RegisterRequest
	4, // 4: pb.auth.AuthService.RefreshToken:input_type -> pb.auth.RefreshTokenRequest
	5, // 5: pb.auth.AuthService.GetUserInfo:output_type -> pb.auth.User
	6, // 6: pb.auth.AuthService.HasAccess:output_type -> google.protobuf.BoolValue
	7, // 7: pb.auth.AuthService.Login:output_type -> pb.auth.AuthResponse
	7, // 8: pb.auth.AuthService.Register:output_type -> pb.auth.AuthResponse
	7, // 9: pb.auth.AuthService.RefreshToken:output_type -> pb.auth.AuthResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_auth_auth_service_proto_init() }
func file_pb_auth_auth_service_proto_init() {
	if File_pb_auth_auth_service_proto != nil {
		return
	}
	file_pb_auth_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_auth_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_auth_service_proto_goTypes,
		DependencyIndexes: file_pb_auth_auth_service_proto_depIdxs,
	}.Build()
	File_pb_auth_auth_service_proto = out.File
	file_pb_auth_auth_service_proto_rawDesc = nil
	file_pb_auth_auth_service_proto_goTypes = nil
	file_pb_auth_auth_service_proto_depIdxs = nil
}
