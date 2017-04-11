// Code generated by protoc-gen-go.
// source: google/iam/v1/iam_policy.proto
// DO NOT EDIT!

package google_iam_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// REQUIRED: The resource for which policy is being specified.
	// Resource is usually specified as a path, such as,
	// projects/{project}/zones/{zone}/disks/{disk}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the 'resource'. The size of
	// the policy is limited to a few 10s of KB. An empty policy is in general a
	// valid policy but certain services (like Projects) might reject them.
	Policy *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
}

func (m *SetIamPolicyRequest) Reset()                    { *m = SetIamPolicyRequest{} }
func (m *SetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*SetIamPolicyRequest) ProtoMessage()               {}
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetIamPolicyRequest) GetPolicy() *Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	// REQUIRED: The resource for which policy is being requested. Resource
	// is usually specified as a path, such as, projects/{project}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
}

func (m *GetIamPolicyRequest) Reset()                    { *m = GetIamPolicyRequest{} }
func (m *GetIamPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIamPolicyRequest) ProtoMessage()               {}
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	// REQUIRED: The resource for which policy detail is being requested.
	// Resource is usually specified as a path, such as, projects/{project}.
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	// The set of permissions to check for the 'resource'. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed.
	Permissions []string `protobuf:"bytes,2,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsRequest) Reset()                    { *m = TestIamPermissionsRequest{} }
func (m *TestIamPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*TestIamPermissionsRequest) ProtoMessage()               {}
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *TestIamPermissionsResponse) Reset()                    { *m = TestIamPermissionsResponse{} }
func (m *TestIamPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*TestIamPermissionsResponse) ProtoMessage()               {}
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*SetIamPolicyRequest)(nil), "google.iam.v1.SetIamPolicyRequest")
	proto.RegisterType((*GetIamPolicyRequest)(nil), "google.iam.v1.GetIamPolicyRequest")
	proto.RegisterType((*TestIamPermissionsRequest)(nil), "google.iam.v1.TestIamPermissionsRequest")
	proto.RegisterType((*TestIamPermissionsResponse)(nil), "google.iam.v1.TestIamPermissionsResponse")
}

func init() { proto.RegisterFile("google/iam/v1/iam_policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0x04, 0x51, 0xf1, 0x05, 0xf9, 0x39, 0x99,
	0xc9, 0x95, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x79, 0x3d, 0xa0, 0x84, 0x5e,
	0x99, 0xa1, 0x94, 0x14, 0xaa, 0x72, 0x64, 0xa5, 0x4a, 0x09, 0x5c, 0xc2, 0xc1, 0xa9, 0x25, 0x9e,
	0x89, 0xb9, 0x01, 0x60, 0xd1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x8e,
	0xa2, 0xd4, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38,
	0x5f, 0x48, 0x97, 0x8b, 0x0d, 0x62, 0x84, 0x04, 0x13, 0x50, 0x86, 0xdb, 0x48, 0x54, 0x0f, 0xc5,
	0x3a, 0x3d, 0xa8, 0x49, 0x50, 0x45, 0x4a, 0x86, 0x5c, 0xc2, 0xee, 0xa4, 0xd9, 0xa0, 0x14, 0xc9,
	0x25, 0x19, 0x02, 0x54, 0x03, 0xd2, 0x93, 0x5a, 0x94, 0x9b, 0x59, 0x5c, 0x9c, 0x99, 0x9f, 0x57,
	0x4c, 0x8c, 0xd3, 0x14, 0xb8, 0xb8, 0x0b, 0x10, 0x3a, 0x80, 0xee, 0x63, 0x06, 0x4a, 0x23, 0x0b,
	0x29, 0xd9, 0x71, 0x49, 0x61, 0x33, 0xba, 0xb8, 0x00, 0x48, 0x61, 0xe8, 0x67, 0xc4, 0xd0, 0x6f,
	0xd4, 0xc3, 0xc4, 0xc5, 0xe9, 0xe9, 0xe8, 0x0b, 0xf1, 0x8b, 0x90, 0x27, 0x17, 0x0f, 0x72, 0xe8,
	0x09, 0x29, 0xa1, 0x05, 0x05, 0x96, 0xa0, 0x95, 0xc2, 0x1e, 0x5c, 0x20, 0xa3, 0xdc, 0xf1, 0x19,
	0xe5, 0x4e, 0xbc, 0x51, 0x99, 0x5c, 0x42, 0x98, 0x7e, 0x14, 0xd2, 0x40, 0x53, 0x8c, 0x33, 0x84,
	0xa5, 0x34, 0x89, 0x50, 0x09, 0x09, 0x30, 0x27, 0x55, 0x2e, 0xc1, 0xe4, 0xfc, 0x5c, 0x54, 0xf5,
	0x4e, 0x7c, 0x70, 0x87, 0x06, 0x80, 0xd2, 0x58, 0x00, 0x63, 0x12, 0x1b, 0x38, 0xb1, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xc8, 0x3c, 0x1b, 0xb9, 0x02, 0x00, 0x00,
}