// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/rerunworkflowrequest.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RerunWorkflowRequest struct {
	ReRunFromWorkflowId  string                    `protobuf:"bytes,1,opt,name=re_run_from_workflow_id,json=reRunFromWorkflowId" json:"re_run_from_workflow_id,omitempty"`
	WorkflowInput        map[string]*_struct.Value `protobuf:"bytes,2,rep,name=workflow_input,json=workflowInput" json:"workflow_input,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReRunFromTaskId      string                    `protobuf:"bytes,3,opt,name=re_run_from_task_id,json=reRunFromTaskId" json:"re_run_from_task_id,omitempty"`
	TaskInput            map[string]*_struct.Value `protobuf:"bytes,4,rep,name=task_input,json=taskInput" json:"task_input,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	CorrelationId        string                    `protobuf:"bytes,5,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RerunWorkflowRequest) Reset()         { *m = RerunWorkflowRequest{} }
func (m *RerunWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*RerunWorkflowRequest) ProtoMessage()    {}
func (*RerunWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rerunworkflowrequest_54d9ae665213e0b8, []int{0}
}
func (m *RerunWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RerunWorkflowRequest.Unmarshal(m, b)
}
func (m *RerunWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RerunWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *RerunWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RerunWorkflowRequest.Merge(dst, src)
}
func (m *RerunWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_RerunWorkflowRequest.Size(m)
}
func (m *RerunWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RerunWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RerunWorkflowRequest proto.InternalMessageInfo

func (m *RerunWorkflowRequest) GetReRunFromWorkflowId() string {
	if m != nil {
		return m.ReRunFromWorkflowId
	}
	return ""
}

func (m *RerunWorkflowRequest) GetWorkflowInput() map[string]*_struct.Value {
	if m != nil {
		return m.WorkflowInput
	}
	return nil
}

func (m *RerunWorkflowRequest) GetReRunFromTaskId() string {
	if m != nil {
		return m.ReRunFromTaskId
	}
	return ""
}

func (m *RerunWorkflowRequest) GetTaskInput() map[string]*_struct.Value {
	if m != nil {
		return m.TaskInput
	}
	return nil
}

func (m *RerunWorkflowRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func init() {
	proto.RegisterType((*RerunWorkflowRequest)(nil), "conductor.proto.RerunWorkflowRequest")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.RerunWorkflowRequest.TaskInputEntry")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.RerunWorkflowRequest.WorkflowInputEntry")
}

func init() {
	proto.RegisterFile("model/rerunworkflowrequest.proto", fileDescriptor_rerunworkflowrequest_54d9ae665213e0b8)
}

var fileDescriptor_rerunworkflowrequest_54d9ae665213e0b8 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0xef, 0xd2, 0x30,
	0x18, 0xc7, 0x33, 0xf8, 0x61, 0x42, 0x09, 0x60, 0x8a, 0x41, 0x82, 0x1e, 0x16, 0x13, 0x13, 0x0e,
	0xa4, 0x4b, 0x90, 0x03, 0xe1, 0x48, 0xa2, 0x09, 0x37, 0x33, 0x89, 0x1a, 0x2f, 0xcb, 0xd6, 0x75,
	0x73, 0xd9, 0xd6, 0x07, 0x9e, 0xb5, 0x22, 0xaf, 0xc0, 0xb7, 0x6d, 0xd6, 0x6d, 0x30, 0x91, 0x83,
	0x07, 0x6f, 0xeb, 0xb3, 0xef, 0x9f, 0xcf, 0x9e, 0x8e, 0xd8, 0x39, 0x84, 0x22, 0x73, 0x50, 0xa0,
	0x96, 0x67, 0xc0, 0x34, 0xca, 0xe0, 0x8c, 0xe2, 0xa4, 0x45, 0xa1, 0xd8, 0x11, 0x41, 0x01, 0x1d,
	0x73, 0x90, 0xa1, 0xe6, 0x0a, 0xb0, 0x1a, 0xcc, 0x5f, 0xc7, 0x00, 0x71, 0x26, 0x1c, 0x73, 0x0a,
	0x74, 0xe4, 0x14, 0x0a, 0x35, 0xaf, 0xe5, 0x6f, 0x7e, 0x3d, 0x91, 0x17, 0x6e, 0x99, 0xf6, 0xa5,
	0x4e, 0x73, 0xab, 0x34, 0xba, 0x26, 0x2f, 0x51, 0x78, 0xa8, 0xa5, 0x17, 0x21, 0xe4, 0x5e, 0x53,
	0xe6, 0x25, 0xe1, 0xcc, 0xb2, 0xad, 0x45, 0xdf, 0x9d, 0xa0, 0x70, 0xb5, 0xfc, 0x80, 0x90, 0x37,
	0xd6, 0x7d, 0x48, 0x3d, 0x32, 0xba, 0x29, 0xe5, 0x51, 0xab, 0x59, 0xc7, 0xee, 0x2e, 0x06, 0xab,
	0x0d, 0xbb, 0xc3, 0x62, 0x8f, 0x4a, 0xd9, 0x35, 0xa9, 0xb4, 0xbe, 0x97, 0x0a, 0x2f, 0xee, 0xf0,
	0xdc, 0x9e, 0xd1, 0x25, 0x99, 0xb4, 0xb1, 0x94, 0x5f, 0xa4, 0x25, 0x52, 0xd7, 0x20, 0x8d, 0xaf,
	0x48, 0x07, 0xbf, 0x48, 0xf7, 0x21, 0xfd, 0x44, 0x48, 0xa5, 0x30, 0x28, 0x4f, 0x06, 0x65, 0xfd,
	0x6f, 0x28, 0x26, 0xe1, 0x86, 0xd1, 0x57, 0xcd, 0x99, 0xbe, 0x25, 0x23, 0x0e, 0x88, 0x22, 0xf3,
	0x55, 0x02, 0xb2, 0x6c, 0xef, 0x99, 0xf6, 0x61, 0x6b, 0xba, 0x0f, 0xe7, 0x5f, 0x09, 0xfd, 0xfb,
	0x73, 0xe8, 0x73, 0xd2, 0x4d, 0xc5, 0xa5, 0x5e, 0x61, 0xf9, 0x48, 0x97, 0xa4, 0xf7, 0xc3, 0xcf,
	0xb4, 0x98, 0x75, 0x6c, 0x6b, 0x31, 0x58, 0x4d, 0x59, 0x75, 0x5f, 0xac, 0xb9, 0x2f, 0xf6, 0xb9,
	0x7c, 0xeb, 0x56, 0xa2, 0x6d, 0x67, 0x63, 0xcd, 0x0f, 0x64, 0xf4, 0x27, 0xdd, 0xff, 0x48, 0xdd,
	0x9d, 0xc8, 0x2b, 0x0e, 0x39, 0x93, 0x42, 0x45, 0x59, 0xf2, 0xf3, 0x7e, 0x49, 0xbb, 0xe9, 0xa3,
	0x2d, 0x7d, 0x0c, 0xbe, 0x6d, 0xe3, 0x44, 0x7d, 0xd7, 0x01, 0xe3, 0x90, 0x3b, 0xb5, 0xd7, 0xb9,
	0x7a, 0x1d, 0x9e, 0x25, 0x42, 0x2a, 0x27, 0x86, 0x18, 0x8f, 0xbc, 0x35, 0x37, 0xbf, 0x71, 0xf0,
	0xcc, 0x44, 0xbf, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xce, 0x3a, 0x9b, 0x51, 0xd6, 0x02, 0x00,
	0x00,
}
