// Code generated by protoc-gen-gogo.
// source: google/rpc/error_details.proto
// DO NOT EDIT!

package google_rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retires have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay *google_protobuf.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
}

func (m *RetryInfo) Reset()                    { *m = RetryInfo{} }
func (m *RetryInfo) String() string            { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()               {}
func (*RetryInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{0} }

func (m *RetryInfo) GetRetryDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (m *DebugInfo) Reset()                    { *m = DebugInfo{} }
func (m *DebugInfo) String() string            { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()               {}
func (*DebugInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{1} }

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryDetail and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
}

func (m *QuotaFailure) Reset()                    { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()               {}
func (*QuotaFailure) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{2} }

func (m *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "project:<Google
	// developer project id>".
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *QuotaFailure_Violation) Reset()         { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()    {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptorErrorDetails, []int{2, 0}
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	// Describes all violations in a client request.
	FieldViolations []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations" json:"field_violations,omitempty"`
}

func (m *BadRequest) Reset()                    { *m = BadRequest{} }
func (m *BadRequest) String() string            { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()               {}
func (*BadRequest) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{3} }

func (m *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if m != nil {
		return m.FieldViolations
	}
	return nil
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	// A path leading to a field in the request body. The value will be a
	// sequence of dot-separated identifiers that identify a protocol buffer
	// field. E.g., "violations.field" would identify this field.
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// A description of why the request element is bad.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *BadRequest_FieldViolation) Reset()         { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()    {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return fileDescriptorErrorDetails, []int{3, 0}
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service generating
	// it. For example, it can be used to identify requests in the service's logs.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string `protobuf:"bytes,2,opt,name=serving_data,json=servingData,proto3" json:"serving_data,omitempty"`
}

func (m *RequestInfo) Reset()                    { *m = RequestInfo{} }
func (m *RequestInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()               {}
func (*RequestInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{4} }

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	// "cloud storage bucket", "file", "Google calendar"; or the type URL
	// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The name of the resource being accessed.  For example, a shared calendar
	// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The owner of the resource (optional).
	// For example, "user:<owner email>" or "project:<Google developer project
	// id>".
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Describes what error is encountered when accessing this resource.
	// For example, updating a cloud project may require the `writer` permission
	// on the developer console project.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ResourceInfo) Reset()                    { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string            { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()               {}
func (*ResourceInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{5} }

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links []*Help_Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{6} }

func (m *Help) GetLinks() []*Help_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// Describes a URL link.
type Help_Link struct {
	// Describes what the link offers.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The URL of the link.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Help_Link) Reset()                    { *m = Help_Link{} }
func (m *Help_Link) String() string            { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()               {}
func (*Help_Link) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{6, 0} }

func init() {
	proto.RegisterType((*RetryInfo)(nil), "google.rpc.RetryInfo")
	proto.RegisterType((*DebugInfo)(nil), "google.rpc.DebugInfo")
	proto.RegisterType((*QuotaFailure)(nil), "google.rpc.QuotaFailure")
	proto.RegisterType((*QuotaFailure_Violation)(nil), "google.rpc.QuotaFailure.Violation")
	proto.RegisterType((*BadRequest)(nil), "google.rpc.BadRequest")
	proto.RegisterType((*BadRequest_FieldViolation)(nil), "google.rpc.BadRequest.FieldViolation")
	proto.RegisterType((*RequestInfo)(nil), "google.rpc.RequestInfo")
	proto.RegisterType((*ResourceInfo)(nil), "google.rpc.ResourceInfo")
	proto.RegisterType((*Help)(nil), "google.rpc.Help")
	proto.RegisterType((*Help_Link)(nil), "google.rpc.Help.Link")
}

var fileDescriptorErrorDetails = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x49, 0x5a, 0xe4, 0x71, 0x28, 0x65, 0x05, 0x28, 0x44, 0x02, 0x15, 0xa3, 0x4a, 0x95,
	0x90, 0x1c, 0xa9, 0xdc, 0x7a, 0x8c, 0xdc, 0x92, 0x4a, 0x08, 0x82, 0x85, 0xb8, 0x5a, 0x8e, 0x3d,
	0x89, 0x4c, 0x5d, 0xaf, 0xd9, 0x8f, 0xa2, 0xfc, 0x0b, 0xee, 0xfc, 0x01, 0x7e, 0x26, 0xfb, 0xd9,
	0x38, 0xcd, 0x85, 0x9b, 0xdf, 0x9b, 0xb7, 0x6f, 0xe7, 0xed, 0x78, 0xe0, 0xcd, 0x9a, 0xd2, 0x75,
	0x83, 0x53, 0xd6, 0x95, 0x53, 0x64, 0x8c, 0xb2, 0xbc, 0x42, 0x51, 0xd4, 0x0d, 0x4f, 0x3a, 0x46,
	0x05, 0x25, 0x60, 0xeb, 0x89, 0xaa, 0x4f, 0xbc, 0xd6, 0x54, 0x96, 0x72, 0x35, 0xad, 0x24, 0x2b,
	0x44, 0x4d, 0x5b, 0xab, 0x8d, 0x3f, 0x42, 0x98, 0xa1, 0x60, 0x9b, 0xeb, 0x76, 0x45, 0xc9, 0x05,
	0x44, 0x4c, 0x03, 0xe5, 0xd7, 0x14, 0x9b, 0x71, 0x70, 0x12, 0x9c, 0x45, 0xe7, 0xaf, 0x12, 0x67,
	0xe7, 0x2d, 0x92, 0xd4, 0x59, 0x64, 0x60, 0xd4, 0xa9, 0x16, 0xc7, 0x73, 0x08, 0x53, 0x5c, 0xca,
	0xb5, 0x31, 0x7a, 0x07, 0x4f, 0xb8, 0x28, 0xca, 0x9b, 0x1c, 0x5b, 0xc1, 0x6a, 0xe4, 0xca, 0x6a,
	0x70, 0x16, 0x66, 0x23, 0x43, 0x5e, 0x5a, 0x8e, 0xbc, 0x84, 0x43, 0xdb, 0xf7, 0xf8, 0x91, 0xba,
	0x28, 0xcc, 0x1c, 0x8a, 0xff, 0x04, 0x30, 0xfa, 0x2a, 0xa9, 0x28, 0xae, 0x14, 0x92, 0x0c, 0xc9,
	0x0c, 0xe0, 0xae, 0xa6, 0x8d, 0xb9, 0xd3, 0x5a, 0x45, 0xe7, 0x71, 0xb2, 0x0d, 0x99, 0xf4, 0xd5,
	0xc9, 0x77, 0x2f, 0xcd, 0x7a, 0xa7, 0x26, 0x2a, 0xe7, 0x7d, 0x81, 0x8c, 0xe1, 0x31, 0x97, 0xcb,
	0x1f, 0x58, 0x0a, 0x93, 0x31, 0xcc, 0x3c, 0x24, 0x27, 0x10, 0x55, 0xc8, 0x4b, 0x56, 0x77, 0x5a,
	0xe8, 0x1a, 0xeb, 0x53, 0xf1, 0xdf, 0x00, 0x60, 0x56, 0x54, 0x19, 0xfe, 0x94, 0xc8, 0x05, 0x59,
	0xc0, 0xf1, 0xaa, 0xc6, 0xa6, 0xca, 0xf7, 0x3a, 0x3c, 0xed, 0x77, 0xb8, 0x3d, 0x91, 0x5c, 0x69,
	0xf9, 0xb6, 0xc9, 0xa7, 0xab, 0x1d, 0xcc, 0x27, 0x73, 0x38, 0xda, 0x95, 0x90, 0xe7, 0x70, 0x60,
	0x44, 0xae, 0x59, 0x0b, 0xfe, 0xa3, 0xd5, 0x2f, 0x10, 0xb9, 0x4b, 0xcd, 0x50, 0x5e, 0x83, 0x9a,
	0x97, 0x81, 0x79, 0xed, 0xbd, 0x42, 0xc7, 0x5c, 0x57, 0xe4, 0x2d, 0x8c, 0x38, 0xb2, 0xbb, 0xba,
	0x5d, 0xe7, 0x55, 0x21, 0x0a, 0x6f, 0xe8, 0xb8, 0x54, 0x51, 0xf1, 0x6f, 0x35, 0x99, 0x0c, 0x39,
	0x95, 0xac, 0x44, 0x3f, 0x67, 0xe6, 0x70, 0x2e, 0x36, 0x1d, 0x3a, 0xd7, 0x91, 0x27, 0xbf, 0x29,
	0x6e, 0x47, 0xd4, 0x16, 0xb7, 0xe8, 0x9c, 0xef, 0x45, 0x9f, 0x15, 0xa7, 0x33, 0xd2, 0x5f, 0x2d,
	0xb2, 0xf1, 0xc0, 0x66, 0x34, 0xe0, 0x61, 0xc6, 0xe1, 0x7e, 0x46, 0x0a, 0xc3, 0x39, 0x36, 0x1d,
	0x79, 0x0f, 0x07, 0x4d, 0xdd, 0xde, 0xf8, 0xc7, 0x7f, 0xd1, 0x7f, 0x7c, 0x2d, 0x48, 0x3e, 0xa9,
	0x6a, 0x66, 0x35, 0x93, 0x0b, 0x18, 0x6a, 0xf8, 0xd0, 0x3e, 0xd8, 0xb3, 0x27, 0xc7, 0x30, 0x90,
	0xcc, 0xff, 0xa0, 0xfa, 0x73, 0x76, 0x0a, 0x47, 0x25, 0xbd, 0xed, 0xd9, 0xcf, 0x9e, 0x5d, 0xea,
	0x1d, 0x4c, 0xed, 0x0a, 0x2e, 0xf4, 0x92, 0x2c, 0x82, 0xe5, 0xa1, 0xd9, 0x96, 0x0f, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0x02, 0xe9, 0x90, 0xac, 0x03, 0x00, 0x00,
}