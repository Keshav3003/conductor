// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowsummary.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

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

type WorkflowSummary struct {
	WorkflowType             string                  `protobuf:"bytes,1,opt,name=workflow_type,json=workflowType" json:"workflow_type,omitempty"`
	Version                  int32                   `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	WorkflowId               string                  `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId" json:"workflow_id,omitempty"`
	CorrelationId            string                  `protobuf:"bytes,4,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	StartTime                string                  `protobuf:"bytes,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	UpdateTime               string                  `protobuf:"bytes,6,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	EndTime                  string                  `protobuf:"bytes,7,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	Status                   Workflow_WorkflowStatus `protobuf:"varint,8,opt,name=status,enum=conductor.proto.Workflow_WorkflowStatus" json:"status,omitempty"`
	Input                    string                  `protobuf:"bytes,9,opt,name=input" json:"input,omitempty"`
	Output                   string                  `protobuf:"bytes,10,opt,name=output" json:"output,omitempty"`
	ReasonForIncompletion    string                  `protobuf:"bytes,11,opt,name=reason_for_incompletion,json=reasonForIncompletion" json:"reason_for_incompletion,omitempty"`
	ExecutionTime            int64                   `protobuf:"varint,12,opt,name=execution_time,json=executionTime" json:"execution_time,omitempty"`
	Event                    string                  `protobuf:"bytes,13,opt,name=event" json:"event,omitempty"`
	FailedReferenceTaskNames string                  `protobuf:"bytes,14,opt,name=failed_reference_task_names,json=failedReferenceTaskNames" json:"failed_reference_task_names,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *WorkflowSummary) Reset()         { *m = WorkflowSummary{} }
func (m *WorkflowSummary) String() string { return proto.CompactTextString(m) }
func (*WorkflowSummary) ProtoMessage()    {}
func (*WorkflowSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowsummary_3f8ed40c0bd9261f, []int{0}
}
func (m *WorkflowSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowSummary.Unmarshal(m, b)
}
func (m *WorkflowSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowSummary.Marshal(b, m, deterministic)
}
func (dst *WorkflowSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowSummary.Merge(dst, src)
}
func (m *WorkflowSummary) XXX_Size() int {
	return xxx_messageInfo_WorkflowSummary.Size(m)
}
func (m *WorkflowSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowSummary.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowSummary proto.InternalMessageInfo

func (m *WorkflowSummary) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *WorkflowSummary) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowSummary) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowSummary) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *WorkflowSummary) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *WorkflowSummary) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *WorkflowSummary) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *WorkflowSummary) GetStatus() Workflow_WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return Workflow_RUNNING
}

func (m *WorkflowSummary) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *WorkflowSummary) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *WorkflowSummary) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *WorkflowSummary) GetExecutionTime() int64 {
	if m != nil {
		return m.ExecutionTime
	}
	return 0
}

func (m *WorkflowSummary) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *WorkflowSummary) GetFailedReferenceTaskNames() string {
	if m != nil {
		return m.FailedReferenceTaskNames
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowSummary)(nil), "conductor.proto.WorkflowSummary")
}

func init() {
	proto.RegisterFile("model/workflowsummary.proto", fileDescriptor_workflowsummary_3f8ed40c0bd9261f)
}

var fileDescriptor_workflowsummary_3f8ed40c0bd9261f = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x89, 0x77, 0x6d, 0xaf, 0x7b, 0xd7, 0x1e, 0x2e, 0x55, 0x57, 0x8b, 0x58, 0x14, 0x21,
	0x4f, 0x09, 0x28, 0xf8, 0x20, 0x08, 0x72, 0x0f, 0x42, 0x5f, 0x44, 0x72, 0x05, 0xc1, 0x97, 0xb0,
	0xdd, 0x4c, 0xea, 0xd2, 0x64, 0x27, 0xec, 0x6e, 0xee, 0xda, 0x4f, 0xe2, 0xd7, 0x95, 0xcc, 0x36,
	0xa5, 0xd7, 0xb7, 0xcc, 0xef, 0xff, 0x9b, 0xdd, 0xcc, 0xb0, 0x6c, 0x5e, 0x63, 0x01, 0x55, 0xfa,
	0x88, 0x76, 0x5b, 0x56, 0xf8, 0xe8, 0xda, 0xba, 0x96, 0x76, 0x9f, 0x34, 0x16, 0x3d, 0xf2, 0x5b,
	0x85, 0xa6, 0x68, 0x95, 0x47, 0x1b, 0xc0, 0x9b, 0xd9, 0x53, 0x3b, 0xd0, 0xf7, 0xff, 0x2e, 0xd9,
	0xed, 0xef, 0x03, 0xba, 0x0f, 0x07, 0xf0, 0x0f, 0x6c, 0xd2, 0x5b, 0xb9, 0xdf, 0x37, 0x20, 0xa2,
	0x45, 0x14, 0x8f, 0xb3, 0x9b, 0x1e, 0xae, 0xf6, 0x0d, 0x70, 0xc1, 0x46, 0x0f, 0x60, 0x9d, 0x46,
	0x23, 0x9e, 0x2d, 0xa2, 0x78, 0x90, 0xf5, 0x25, 0x7f, 0xc7, 0xae, 0x8f, 0xed, 0xba, 0x10, 0x17,
	0xd4, 0xcc, 0x7a, 0xb4, 0x2c, 0xf8, 0x47, 0x36, 0x55, 0x68, 0x2d, 0x54, 0xd2, 0x6b, 0x34, 0x9d,
	0x73, 0x49, 0xce, 0xe4, 0x84, 0x2e, 0x0b, 0xfe, 0x96, 0x31, 0xe7, 0xa5, 0xf5, 0xb9, 0xd7, 0x35,
	0x88, 0x01, 0x29, 0x63, 0x22, 0x2b, 0x5d, 0x43, 0x77, 0x4d, 0xdb, 0x14, 0xd2, 0x43, 0xc8, 0x87,
	0xe1, 0x9a, 0x80, 0x48, 0x78, 0xcd, 0xae, 0xc0, 0x14, 0x21, 0x1d, 0x51, 0x3a, 0x02, 0x53, 0x50,
	0xf4, 0x9d, 0x0d, 0x9d, 0x97, 0xbe, 0x75, 0xe2, 0x6a, 0x11, 0xc5, 0xd3, 0x4f, 0x71, 0x72, 0xb6,
	0xad, 0xa4, 0xdf, 0xc9, 0xf1, 0xe3, 0x9e, 0xfc, 0xec, 0xd0, 0xc7, 0x67, 0x6c, 0xa0, 0x4d, 0xd3,
	0x7a, 0x31, 0xa6, 0x93, 0x43, 0xc1, 0x5f, 0xb2, 0x21, 0xb6, 0xbe, 0xc3, 0x8c, 0xf0, 0xa1, 0xe2,
	0x5f, 0xd8, 0x2b, 0x0b, 0xd2, 0xa1, 0xc9, 0x4b, 0xb4, 0xb9, 0x36, 0x0a, 0xeb, 0xa6, 0x82, 0x6e,
	0x4e, 0x71, 0x4d, 0xe2, 0x8b, 0x10, 0xff, 0x40, 0xbb, 0x3c, 0x09, 0xbb, 0x4d, 0xc1, 0x0e, 0x54,
	0x4b, 0x7b, 0xa2, 0x41, 0x6e, 0x16, 0x51, 0x7c, 0x91, 0x4d, 0x8e, 0x94, 0xc6, 0x99, 0xb1, 0x01,
	0x3c, 0x80, 0xf1, 0x62, 0x12, 0x7e, 0x86, 0x0a, 0xfe, 0x8d, 0xcd, 0x4b, 0xa9, 0x2b, 0x28, 0x72,
	0x0b, 0x25, 0x58, 0x30, 0x0a, 0x72, 0x2f, 0xdd, 0x36, 0x37, 0xb2, 0x06, 0x27, 0xa6, 0xe4, 0x8a,
	0xa0, 0x64, 0xbd, 0xb1, 0x92, 0x6e, 0xfb, 0xb3, 0xcb, 0xef, 0x2a, 0x36, 0x57, 0x58, 0x27, 0x06,
	0x7c, 0x59, 0xe9, 0xdd, 0xf9, 0x82, 0xee, 0x9e, 0x9f, 0xbd, 0x9a, 0x5f, 0xeb, 0x3f, 0x5f, 0x37,
	0xda, 0xff, 0x6d, 0xd7, 0x89, 0xc2, 0x3a, 0x3d, 0xb4, 0xa5, 0xc7, 0xb6, 0x54, 0x55, 0x1a, 0x8c,
	0x4f, 0x37, 0xb8, 0xb1, 0x8d, 0x3a, 0xe1, 0xf4, 0x2c, 0xd7, 0x43, 0x3a, 0xf5, 0xf3, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0xee, 0x86, 0xf0, 0xd4, 0x02, 0x00, 0x00,
}
