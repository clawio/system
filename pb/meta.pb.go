// Code generated by protoc-gen-go.
// source: meta.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MetaData_FileSystemElement int32

const (
	MetaData_CONTAINER MetaData_FileSystemElement = 0
	MetaData_FILE      MetaData_FileSystemElement = 1
)

var MetaData_FileSystemElement_name = map[int32]string{
	0: "CONTAINER",
	1: "FILE",
}
var MetaData_FileSystemElement_value = map[string]int32{
	"CONTAINER": 0,
	"FILE":      1,
}

func (x MetaData_FileSystemElement) String() string {
	return proto.EnumName(MetaData_FileSystemElement_name, int32(x))
}
func (MetaData_FileSystemElement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

type MetaData struct {
	Path           string                     `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Type           MetaData_FileSystemElement `protobuf:"varint,2,opt,name=type,enum=pb.MetaData_FileSystemElement" json:"type,omitempty"`
	Size           int64                      `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	ServerModified int64                      `protobuf:"varint,4,opt,name=server_modified,json=serverModified" json:"server_modified,omitempty"`
	Checksum       string                     `protobuf:"bytes,5,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *MetaData) Reset()                    { *m = MetaData{} }
func (m *MetaData) String() string            { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()               {}
func (*MetaData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MetaData) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MetaData) GetType() MetaData_FileSystemElement {
	if m != nil {
		return m.Type
	}
	return MetaData_CONTAINER
}

func (m *MetaData) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MetaData) GetServerModified() int64 {
	if m != nil {
		return m.ServerModified
	}
	return 0
}

func (m *MetaData) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func init() {
	proto.RegisterType((*MetaData)(nil), "pb.MetaData")
	proto.RegisterEnum("pb.MetaData_FileSystemElement", MetaData_FileSystemElement_name, MetaData_FileSystemElement_value)
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2d, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xba, 0xc9, 0xc8, 0xc5, 0xe1,
	0x9b, 0x5a, 0x92, 0xe8, 0x92, 0x58, 0x92, 0x28, 0x24, 0xc4, 0xc5, 0x52, 0x90, 0x58, 0x92, 0x21,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x19, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4,
	0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe9, 0x15, 0x24, 0xe9, 0xc1, 0xd4, 0xeb, 0xb9,
	0x65, 0xe6, 0xa4, 0x06, 0x57, 0x16, 0x97, 0xa4, 0xe6, 0xba, 0xe6, 0xa4, 0xe6, 0xa6, 0xe6, 0x95,
	0x04, 0x81, 0xd5, 0x82, 0xcc, 0x29, 0xce, 0xac, 0x4a, 0x95, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x0e,
	0x02, 0xb3, 0x85, 0xd4, 0xb9, 0xf8, 0x8b, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0xe2, 0x73, 0xf3, 0x53,
	0x32, 0xd3, 0x32, 0x53, 0x53, 0x24, 0x58, 0xc0, 0xd2, 0x7c, 0x10, 0x61, 0x5f, 0xa8, 0xa8, 0x90,
	0x14, 0x17, 0x47, 0x72, 0x46, 0x6a, 0x72, 0x76, 0x71, 0x69, 0xae, 0x04, 0x2b, 0xd8, 0x21, 0x70,
	0xbe, 0x92, 0x0e, 0x97, 0x20, 0x86, 0x9d, 0x42, 0xbc, 0x5c, 0x9c, 0xce, 0xfe, 0x7e, 0x21, 0x8e,
	0x9e, 0x7e, 0xae, 0x41, 0x02, 0x0c, 0x42, 0x1c, 0x5c, 0x2c, 0x6e, 0x9e, 0x3e, 0xae, 0x02, 0x8c,
	0x49, 0x6c, 0x60, 0x6f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x1b, 0x44, 0xe6, 0xf4,
	0x00, 0x00, 0x00,
}
