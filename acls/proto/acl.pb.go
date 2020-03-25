// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acl.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ApiClientACL struct {
	AllQuery             bool     `protobuf:"varint,1,opt,name=all_query,json=allQuery,proto3" json:"all_query,omitempty"`
	AnyQuery             bool     `protobuf:"varint,10,opt,name=any_query,json=anyQuery,proto3" json:"any_query,omitempty"`
	PublishQueues        []string `protobuf:"bytes,2,rep,name=publish_queues,json=publishQueues,proto3" json:"publish_queues,omitempty"`
	ReadResults          bool     `protobuf:"varint,3,opt,name=read_results,json=readResults,proto3" json:"read_results,omitempty"`
	LabelClients         bool     `protobuf:"varint,11,opt,name=label_clients,json=labelClients,proto3" json:"label_clients,omitempty"`
	CollectClient        bool     `protobuf:"varint,4,opt,name=collect_client,json=collectClient,proto3" json:"collect_client,omitempty"`
	CollectServer        bool     `protobuf:"varint,5,opt,name=collect_server,json=collectServer,proto3" json:"collect_server,omitempty"`
	ArtifactWriter       bool     `protobuf:"varint,6,opt,name=artifact_writer,json=artifactWriter,proto3" json:"artifact_writer,omitempty"`
	ServerArtifactWriter bool     `protobuf:"varint,15,opt,name=server_artifact_writer,json=serverArtifactWriter,proto3" json:"server_artifact_writer,omitempty"`
	Execve               bool     `protobuf:"varint,7,opt,name=execve,proto3" json:"execve,omitempty"`
	NotebookEditor       bool     `protobuf:"varint,8,opt,name=notebook_editor,json=notebookEditor,proto3" json:"notebook_editor,omitempty"`
	ServerAdmin          bool     `protobuf:"varint,12,opt,name=server_admin,json=serverAdmin,proto3" json:"server_admin,omitempty"`
	FilesystemRead       bool     `protobuf:"varint,13,opt,name=filesystem_read,json=filesystemRead,proto3" json:"filesystem_read,omitempty"`
	FilesystemWrite      bool     `protobuf:"varint,14,opt,name=filesystem_write,json=filesystemWrite,proto3" json:"filesystem_write,omitempty"`
	MachineState         bool     `protobuf:"varint,16,opt,name=machine_state,json=machineState,proto3" json:"machine_state,omitempty"`
	// A list of roles in lieu of the permissions above. These will be
	// interpolated into this ACL object.
	Roles                []string `protobuf:"bytes,9,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiClientACL) Reset()         { *m = ApiClientACL{} }
func (m *ApiClientACL) String() string { return proto.CompactTextString(m) }
func (*ApiClientACL) ProtoMessage()    {}
func (*ApiClientACL) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0}
}

func (m *ApiClientACL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClientACL.Unmarshal(m, b)
}
func (m *ApiClientACL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClientACL.Marshal(b, m, deterministic)
}
func (m *ApiClientACL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClientACL.Merge(m, src)
}
func (m *ApiClientACL) XXX_Size() int {
	return xxx_messageInfo_ApiClientACL.Size(m)
}
func (m *ApiClientACL) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClientACL.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClientACL proto.InternalMessageInfo

func (m *ApiClientACL) GetAllQuery() bool {
	if m != nil {
		return m.AllQuery
	}
	return false
}

func (m *ApiClientACL) GetAnyQuery() bool {
	if m != nil {
		return m.AnyQuery
	}
	return false
}

func (m *ApiClientACL) GetPublishQueues() []string {
	if m != nil {
		return m.PublishQueues
	}
	return nil
}

func (m *ApiClientACL) GetReadResults() bool {
	if m != nil {
		return m.ReadResults
	}
	return false
}

func (m *ApiClientACL) GetLabelClients() bool {
	if m != nil {
		return m.LabelClients
	}
	return false
}

func (m *ApiClientACL) GetCollectClient() bool {
	if m != nil {
		return m.CollectClient
	}
	return false
}

func (m *ApiClientACL) GetCollectServer() bool {
	if m != nil {
		return m.CollectServer
	}
	return false
}

func (m *ApiClientACL) GetArtifactWriter() bool {
	if m != nil {
		return m.ArtifactWriter
	}
	return false
}

func (m *ApiClientACL) GetServerArtifactWriter() bool {
	if m != nil {
		return m.ServerArtifactWriter
	}
	return false
}

func (m *ApiClientACL) GetExecve() bool {
	if m != nil {
		return m.Execve
	}
	return false
}

func (m *ApiClientACL) GetNotebookEditor() bool {
	if m != nil {
		return m.NotebookEditor
	}
	return false
}

func (m *ApiClientACL) GetServerAdmin() bool {
	if m != nil {
		return m.ServerAdmin
	}
	return false
}

func (m *ApiClientACL) GetFilesystemRead() bool {
	if m != nil {
		return m.FilesystemRead
	}
	return false
}

func (m *ApiClientACL) GetFilesystemWrite() bool {
	if m != nil {
		return m.FilesystemWrite
	}
	return false
}

func (m *ApiClientACL) GetMachineState() bool {
	if m != nil {
		return m.MachineState
	}
	return false
}

func (m *ApiClientACL) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

// A role is a named sets of ACL permissions. A user may possess
// multiple roles.
type Role struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions          *ApiClientACL `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPermissions() *ApiClientACL {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiClientACL)(nil), "proto.ApiClientACL")
	proto.RegisterType((*Role)(nil), "proto.Role")
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_a452f070aeef01eb) }

var fileDescriptor_a452f070aeef01eb = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0x15, 0x9a, 0x84, 0xac, 0xf3, 0xd1, 0xca, 0x54, 0x95, 0x05, 0x17, 0x53, 0x04, 0xa4,
	0x42, 0xda, 0x88, 0xaf, 0x0b, 0xb7, 0xb4, 0xe2, 0x44, 0x0f, 0x64, 0x7b, 0x80, 0xdb, 0xca, 0xf1,
	0x4e, 0x1a, 0x0b, 0xaf, 0x1d, 0x6c, 0x27, 0x69, 0x5e, 0x8d, 0x33, 0x4f, 0x02, 0xaf, 0xc1, 0x01,
	0x79, 0xbc, 0x11, 0x29, 0xa7, 0x5d, 0xff, 0xfc, 0xf3, 0x7f, 0xac, 0xf1, 0x90, 0x4c, 0x48, 0x9d,
	0xaf, 0x9c, 0x0d, 0x96, 0x76, 0xf0, 0xf3, 0xf8, 0xc3, 0x76, 0xbb, 0xcd, 0x37, 0xa0, 0xad, 0x54,
	0x15, 0xdc, 0xe5, 0xd2, 0xd6, 0x93, 0x5b, 0xab, 0x85, 0xb9, 0x9d, 0x24, 0xe8, 0xc4, 0x2a, 0x58,
	0x37, 0x41, 0x79, 0xe2, 0xa1, 0x16, 0x26, 0x28, 0x99, 0x22, 0xce, 0x7f, 0x74, 0xc8, 0x60, 0xba,
	0x52, 0x57, 0x5a, 0x81, 0x09, 0xd3, 0xab, 0x6b, 0xfa, 0x89, 0x64, 0x42, 0xeb, 0xf2, 0xfb, 0x1a,
	0xdc, 0x8e, 0xb5, 0x78, 0x6b, 0xdc, 0xbb, 0xcc, 0x7f, 0xfd, 0xf9, 0xfd, 0xb3, 0x35, 0xa6, 0x2f,
	0x3e, 0x3b, 0xbb, 0x51, 0x15, 0x78, 0x2e, 0xdc, 0x5c, 0x05, 0x27, 0xdc, 0x8e, 0xa3, 0xc7, 0x35,
	0x6c, 0x40, 0x73, 0x21, 0x25, 0x78, 0x9f, 0x17, 0x3d, 0xa1, 0xf5, 0x2c, 0x72, 0xfa, 0x84, 0x64,
	0xc2, 0xec, 0x9a, 0x30, 0x12, 0xc3, 0x8a, 0x9e, 0x30, 0xbb, 0xb4, 0xf9, 0x95, 0x8c, 0x56, 0xeb,
	0xb9, 0x56, 0x7e, 0x19, 0x85, 0x35, 0x78, 0xf6, 0x80, 0x1f, 0x8d, 0xb3, 0xcb, 0xd7, 0x58, 0xee,
	0x15, 0xbd, 0xb8, 0x56, 0x3e, 0x70, 0xbb, 0xe0, 0x69, 0x97, 0x87, 0x25, 0x70, 0x89, 0xf7, 0xe4,
	0x52, 0x18, 0xde, 0x9c, 0xe5, 0xc1, 0xe6, 0xc5, 0xb0, 0x59, 0xcc, 0xd0, 0xa4, 0x4f, 0xc9, 0xc0,
	0x81, 0xa8, 0x4a, 0x07, 0x7e, 0xad, 0x83, 0x67, 0x47, 0x58, 0xb9, 0x1f, 0x59, 0x91, 0x10, 0x7d,
	0x46, 0x86, 0x5a, 0xcc, 0x41, 0x97, 0x29, 0xd1, 0xb3, 0x3e, 0x3a, 0x03, 0x84, 0xa9, 0x1b, 0x9e,
	0x3e, 0x27, 0x23, 0x69, 0xb5, 0x06, 0x19, 0x1a, 0x8d, 0xb5, 0xd1, 0x1a, 0x36, 0x34, 0x79, 0x87,
	0x9a, 0x07, 0xb7, 0x01, 0xc7, 0x3a, 0xf7, 0xb4, 0x1b, 0x84, 0xf4, 0x25, 0x39, 0x16, 0x2e, 0xa8,
	0x85, 0x90, 0xa1, 0xdc, 0x3a, 0x15, 0xc0, 0xb1, 0x2e, 0x7a, 0xa3, 0x3d, 0xfe, 0x82, 0x94, 0xbe,
	0x23, 0x67, 0x29, 0xa7, 0xfc, 0xdf, 0x3f, 0x46, 0xff, 0x34, 0xed, 0x4e, 0xef, 0x9f, 0x3a, 0x23,
	0x5d, 0xb8, 0x03, 0xb9, 0x01, 0xf6, 0x10, 0xad, 0x66, 0x15, 0xcb, 0x1a, 0x1b, 0x60, 0x6e, 0xed,
	0xb7, 0x12, 0x2a, 0x15, 0xac, 0x63, 0xbd, 0x54, 0x76, 0x8f, 0x3f, 0x22, 0x8d, 0x5d, 0xdb, 0x97,
	0xad, 0x6a, 0x65, 0xd8, 0x20, 0x75, 0xad, 0x29, 0x16, 0x51, 0xcc, 0x5a, 0x28, 0x0d, 0x7e, 0xe7,
	0x03, 0xd4, 0x65, 0xec, 0x27, 0x1b, 0xa6, 0xac, 0x7f, 0xb8, 0x00, 0x51, 0xd1, 0x0b, 0x72, 0x72,
	0x20, 0xe2, 0xed, 0xd9, 0x08, 0xcd, 0x83, 0x00, 0xbc, 0x78, 0x7c, 0x89, 0x5a, 0xc8, 0xa5, 0x32,
	0x50, 0xfa, 0x20, 0x02, 0xb0, 0x93, 0xf4, 0x12, 0x0d, 0xbc, 0x89, 0x8c, 0x9e, 0x92, 0x8e, 0xb3,
	0x1a, 0x3c, 0xcb, 0xe2, 0x88, 0x14, 0x69, 0x71, 0x3e, 0x23, 0xed, 0xc2, 0x6a, 0xa0, 0x94, 0xb4,
	0x8d, 0xa8, 0x01, 0xc7, 0x35, 0x2b, 0xf0, 0x9f, 0xbe, 0x27, 0xfd, 0x15, 0xb8, 0x5a, 0x79, 0xaf,
	0xac, 0x89, 0xa3, 0xd5, 0x1a, 0xf7, 0xdf, 0x3c, 0x4a, 0x53, 0x9f, 0x1f, 0x4e, 0x7c, 0x71, 0xe8,
	0xcd, 0xbb, 0x28, 0xbc, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x01, 0x5d, 0x2a, 0x66, 0x03,
	0x00, 0x00,
}