// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: pb/auth/auth_service.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x12, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x62, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70,
	0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x62,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x62,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xa8, 0x0c, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x48, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x2e, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x62,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_pb_auth_auth_service_proto_goTypes = []interface{}{
	(*GetUserInfoRequest)(nil),           // 0: pb.auth.GetUserInfoRequest
	(*HasAccessRequest)(nil),             // 1: pb.auth.HasAccessRequest
	(*RefreshTokenRequest)(nil),          // 2: pb.auth.RefreshTokenRequest
	(*LoginRequest)(nil),                 // 3: pb.auth.LoginRequest
	(*RegisterRequest)(nil),              // 4: pb.auth.RegisterRequest
	(*LogoutRequest)(nil),                // 5: pb.auth.LogoutRequest
	(*FindPermissionByIDRequest)(nil),    // 6: pb.auth.FindPermissionByIDRequest
	(*FindPermissionByNameRequest)(nil),  // 7: pb.auth.FindPermissionByNameRequest
	(*CreatePermissionRequest)(nil),      // 8: pb.auth.CreatePermissionRequest
	(*DeletePermissionRequest)(nil),      // 9: pb.auth.DeletePermissionRequest
	(*FindGroupByIDRequest)(nil),         // 10: pb.auth.FindGroupByIDRequest
	(*FindGroupByNameRequest)(nil),       // 11: pb.auth.FindGroupByNameRequest
	(*CreateGroupRequest)(nil),           // 12: pb.auth.CreateGroupRequest
	(*DeleteGroupRequest)(nil),           // 13: pb.auth.DeleteGroupRequest
	(*FindGroupPermissionRequest)(nil),   // 14: pb.auth.FindGroupPermissionRequest
	(*CreateGroupPermissionRequest)(nil), // 15: pb.auth.CreateGroupPermissionRequest
	(*DeleteGroupPermissionRequest)(nil), // 16: pb.auth.DeleteGroupPermissionRequest
	(*FindAllUserGroupsRequest)(nil),     // 17: pb.auth.FindAllUserGroupsRequest
	(*FindUserGroupRequest)(nil),         // 18: pb.auth.FindUserGroupRequest
	(*CreateUserGroupRequest)(nil),       // 19: pb.auth.CreateUserGroupRequest
	(*DeleteUserGroupRequest)(nil),       // 20: pb.auth.DeleteUserGroupRequest
	(*User)(nil),                         // 21: pb.auth.User
	(*wrapperspb.BoolValue)(nil),         // 22: google.protobuf.BoolValue
	(*AuthResponse)(nil),                 // 23: pb.auth.AuthResponse
	(*emptypb.Empty)(nil),                // 24: google.protobuf.Empty
	(*Permission)(nil),                   // 25: pb.auth.Permission
	(*Group)(nil),                        // 26: pb.auth.Group
	(*GroupPermission)(nil),              // 27: pb.auth.GroupPermission
	(*FindAllUserGroupsResponse)(nil),    // 28: pb.auth.FindAllUserGroupsResponse
	(*UserGroup)(nil),                    // 29: pb.auth.UserGroup
}
var file_pb_auth_auth_service_proto_depIdxs = []int32{
	0,  // 0: pb.auth.AuthService.GetUserInfo:input_type -> pb.auth.GetUserInfoRequest
	1,  // 1: pb.auth.AuthService.HasAccess:input_type -> pb.auth.HasAccessRequest
	2,  // 2: pb.auth.AuthService.RefreshToken:input_type -> pb.auth.RefreshTokenRequest
	3,  // 3: pb.auth.AuthService.Login:input_type -> pb.auth.LoginRequest
	4,  // 4: pb.auth.AuthService.Register:input_type -> pb.auth.RegisterRequest
	5,  // 5: pb.auth.AuthService.Logout:input_type -> pb.auth.LogoutRequest
	6,  // 6: pb.auth.AuthService.FindPermissionByID:input_type -> pb.auth.FindPermissionByIDRequest
	7,  // 7: pb.auth.AuthService.FindPermissionByName:input_type -> pb.auth.FindPermissionByNameRequest
	8,  // 8: pb.auth.AuthService.CreatePermission:input_type -> pb.auth.CreatePermissionRequest
	9,  // 9: pb.auth.AuthService.DeletePermission:input_type -> pb.auth.DeletePermissionRequest
	10, // 10: pb.auth.AuthService.FindGroupByID:input_type -> pb.auth.FindGroupByIDRequest
	11, // 11: pb.auth.AuthService.FindGroupByName:input_type -> pb.auth.FindGroupByNameRequest
	12, // 12: pb.auth.AuthService.CreateGroup:input_type -> pb.auth.CreateGroupRequest
	13, // 13: pb.auth.AuthService.DeleteGroupByID:input_type -> pb.auth.DeleteGroupRequest
	14, // 14: pb.auth.AuthService.FindGroupPermission:input_type -> pb.auth.FindGroupPermissionRequest
	15, // 15: pb.auth.AuthService.CreateGroupPermission:input_type -> pb.auth.CreateGroupPermissionRequest
	16, // 16: pb.auth.AuthService.DeleteGroupPermission:input_type -> pb.auth.DeleteGroupPermissionRequest
	17, // 17: pb.auth.AuthService.FindAllUserGroups:input_type -> pb.auth.FindAllUserGroupsRequest
	18, // 18: pb.auth.AuthService.FindUserGroup:input_type -> pb.auth.FindUserGroupRequest
	19, // 19: pb.auth.AuthService.CreateUserGroup:input_type -> pb.auth.CreateUserGroupRequest
	20, // 20: pb.auth.AuthService.DeleteUserGroup:input_type -> pb.auth.DeleteUserGroupRequest
	21, // 21: pb.auth.AuthService.GetUserInfo:output_type -> pb.auth.User
	22, // 22: pb.auth.AuthService.HasAccess:output_type -> google.protobuf.BoolValue
	23, // 23: pb.auth.AuthService.RefreshToken:output_type -> pb.auth.AuthResponse
	23, // 24: pb.auth.AuthService.Login:output_type -> pb.auth.AuthResponse
	23, // 25: pb.auth.AuthService.Register:output_type -> pb.auth.AuthResponse
	24, // 26: pb.auth.AuthService.Logout:output_type -> google.protobuf.Empty
	25, // 27: pb.auth.AuthService.FindPermissionByID:output_type -> pb.auth.Permission
	25, // 28: pb.auth.AuthService.FindPermissionByName:output_type -> pb.auth.Permission
	25, // 29: pb.auth.AuthService.CreatePermission:output_type -> pb.auth.Permission
	24, // 30: pb.auth.AuthService.DeletePermission:output_type -> google.protobuf.Empty
	26, // 31: pb.auth.AuthService.FindGroupByID:output_type -> pb.auth.Group
	26, // 32: pb.auth.AuthService.FindGroupByName:output_type -> pb.auth.Group
	26, // 33: pb.auth.AuthService.CreateGroup:output_type -> pb.auth.Group
	24, // 34: pb.auth.AuthService.DeleteGroupByID:output_type -> google.protobuf.Empty
	27, // 35: pb.auth.AuthService.FindGroupPermission:output_type -> pb.auth.GroupPermission
	27, // 36: pb.auth.AuthService.CreateGroupPermission:output_type -> pb.auth.GroupPermission
	24, // 37: pb.auth.AuthService.DeleteGroupPermission:output_type -> google.protobuf.Empty
	28, // 38: pb.auth.AuthService.FindAllUserGroups:output_type -> pb.auth.FindAllUserGroupsResponse
	29, // 39: pb.auth.AuthService.FindUserGroup:output_type -> pb.auth.UserGroup
	29, // 40: pb.auth.AuthService.CreateUserGroup:output_type -> pb.auth.UserGroup
	24, // 41: pb.auth.AuthService.DeleteUserGroup:output_type -> google.protobuf.Empty
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pb_auth_auth_service_proto_init() }
func file_pb_auth_auth_service_proto_init() {
	if File_pb_auth_auth_service_proto != nil {
		return
	}
	file_pb_auth_user_proto_init()
	file_pb_auth_auth_proto_init()
	file_pb_auth_permission_proto_init()
	file_pb_auth_group_proto_init()
	file_pb_auth_group_permission_proto_init()
	file_pb_auth_user_group_proto_init()
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
