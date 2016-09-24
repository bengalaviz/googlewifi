// Code generated by protoc-gen-go.
// source: diagnosticreport.proto
// DO NOT EDIT!

/*
Package diagnosticreport is a generated protocol buffer package.

It is generated from these files:
	diagnosticreport.proto

It has these top-level messages:
	File
	FileLength
	CommandOutput
	UnknownPair
	DiagnosticReport
*/
package diagnosticreport

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

type File struct {
	Path    string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FileLength struct {
	Path   string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Length int32  `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
}

func (m *FileLength) Reset()                    { *m = FileLength{} }
func (m *FileLength) String() string            { return proto.CompactTextString(m) }
func (*FileLength) ProtoMessage()               {}
func (*FileLength) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CommandOutput struct {
	Command    string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
	Output     string `protobuf:"bytes,2,opt,name=output" json:"output,omitempty"`
	StatusCode int32  `protobuf:"varint,3,opt,name=statusCode" json:"statusCode,omitempty"`
}

func (m *CommandOutput) Reset()                    { *m = CommandOutput{} }
func (m *CommandOutput) String() string            { return proto.CompactTextString(m) }
func (*CommandOutput) ProtoMessage()               {}
func (*CommandOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type UnknownPair struct {
	Unknown1 string `protobuf:"bytes,1,opt,name=unknown1" json:"unknown1,omitempty"`
	Unknown2 string `protobuf:"bytes,2,opt,name=unknown2" json:"unknown2,omitempty"`
}

func (m *UnknownPair) Reset()                    { *m = UnknownPair{} }
func (m *UnknownPair) String() string            { return proto.CompactTextString(m) }
func (*UnknownPair) ProtoMessage()               {}
func (*UnknownPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DiagnosticReport struct {
	Version          string           `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Files            []*File          `protobuf:"bytes,2,rep,name=files" json:"files,omitempty"`
	StormVersion     string           `protobuf:"bytes,3,opt,name=stormVersion" json:"stormVersion,omitempty"`
	WhirlwindVersion string           `protobuf:"bytes,4,opt,name=whirlwindVersion" json:"whirlwindVersion,omitempty"`
	NetworkConfig    string           `protobuf:"bytes,5,opt,name=networkConfig" json:"networkConfig,omitempty"`
	FileLengths      []*FileLength    `protobuf:"bytes,7,rep,name=fileLengths" json:"fileLengths,omitempty"`
	WanInfo          string           `protobuf:"bytes,8,opt,name=wanInfo" json:"wanInfo,omitempty"`
	CommandOutputs   []*CommandOutput `protobuf:"bytes,9,rep,name=commandOutputs" json:"commandOutputs,omitempty"`
	InfoJSON         string           `protobuf:"bytes,10,opt,name=infoJSON" json:"infoJSON,omitempty"`
	Unknown1         int32            `protobuf:"varint,11,opt,name=unknown1" json:"unknown1,omitempty"`
	UnknownPairs     []*UnknownPair   `protobuf:"bytes,12,rep,name=unknownPairs" json:"unknownPairs,omitempty"`
	UnixTime         int32            `protobuf:"varint,13,opt,name=unixTime" json:"unixTime,omitempty"`
}

func (m *DiagnosticReport) Reset()                    { *m = DiagnosticReport{} }
func (m *DiagnosticReport) String() string            { return proto.CompactTextString(m) }
func (*DiagnosticReport) ProtoMessage()               {}
func (*DiagnosticReport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DiagnosticReport) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func (m *DiagnosticReport) GetFileLengths() []*FileLength {
	if m != nil {
		return m.FileLengths
	}
	return nil
}

func (m *DiagnosticReport) GetCommandOutputs() []*CommandOutput {
	if m != nil {
		return m.CommandOutputs
	}
	return nil
}

func (m *DiagnosticReport) GetUnknownPairs() []*UnknownPair {
	if m != nil {
		return m.UnknownPairs
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "diagnosticreport.File")
	proto.RegisterType((*FileLength)(nil), "diagnosticreport.FileLength")
	proto.RegisterType((*CommandOutput)(nil), "diagnosticreport.CommandOutput")
	proto.RegisterType((*UnknownPair)(nil), "diagnosticreport.UnknownPair")
	proto.RegisterType((*DiagnosticReport)(nil), "diagnosticreport.DiagnosticReport")
}

func init() { proto.RegisterFile("diagnosticreport.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0xcb, 0xae, 0xda, 0x30,
	0x14, 0xd4, 0xbd, 0x84, 0x0b, 0x9c, 0x40, 0x85, 0xbc, 0x40, 0x56, 0xd5, 0x07, 0x8a, 0xba, 0x40,
	0x55, 0x85, 0x54, 0xda, 0x45, 0x57, 0x95, 0x2a, 0xfa, 0x50, 0xab, 0xaa, 0x54, 0xe9, 0x63, 0x9f,
	0x12, 0x07, 0x2c, 0x12, 0x3b, 0xb2, 0x9d, 0xa6, 0xdf, 0xd0, 0xaf, 0xbe, 0xca, 0xb1, 0x13, 0x12,
	0xc2, 0x8e, 0x19, 0x9f, 0x99, 0x39, 0xf6, 0x10, 0x58, 0xc4, 0x3c, 0x3a, 0x08, 0xa9, 0x0d, 0xdf,
	0x2b, 0x96, 0x4b, 0x65, 0xd6, 0xb9, 0x92, 0x46, 0x92, 0xf9, 0x25, 0x1f, 0xbc, 0x06, 0xef, 0x23,
	0x4f, 0x19, 0x21, 0xe0, 0xe5, 0x91, 0x39, 0xd2, 0x9b, 0xe5, 0xcd, 0x6a, 0x12, 0xe2, 0x6f, 0x42,
	0x61, 0xb4, 0x97, 0xc2, 0x30, 0x61, 0xe8, 0x2d, 0xd2, 0x35, 0x0c, 0xde, 0x00, 0x54, 0xaa, 0xaf,
	0x4c, 0x1c, 0xcc, 0xf1, 0xaa, 0x76, 0x01, 0x77, 0x29, 0x9e, 0xa2, 0x74, 0x18, 0x3a, 0x14, 0x44,
	0x30, 0xdb, 0xca, 0x2c, 0x8b, 0x44, 0xbc, 0x2b, 0x4c, 0x5e, 0x18, 0x1b, 0x82, 0x84, 0xd3, 0xd7,
	0xb0, 0xb2, 0x90, 0x38, 0xe3, 0xd2, 0x1d, 0x22, 0x4f, 0x00, 0xb4, 0x89, 0x4c, 0xa1, 0xb7, 0x32,
	0x66, 0x74, 0x80, 0xf6, 0x2d, 0x26, 0xf8, 0x00, 0xfe, 0x2f, 0x71, 0x12, 0xb2, 0x14, 0xdf, 0x23,
	0xae, 0xc8, 0x43, 0x18, 0x17, 0x16, 0xbe, 0x74, 0x09, 0x0d, 0x6e, 0x9d, 0x6d, 0x5c, 0x48, 0x83,
	0x83, 0xff, 0x1e, 0xcc, 0xdf, 0x37, 0xcf, 0x15, 0xe2, 0x73, 0x55, 0xdb, 0xfe, 0x65, 0x4a, 0x73,
	0x29, 0xea, 0x6d, 0x1d, 0x24, 0x2f, 0x60, 0x98, 0xf0, 0x94, 0x69, 0x7a, 0xbb, 0x1c, 0xac, 0xfc,
	0xcd, 0x62, 0xdd, 0xab, 0xa0, 0x7a, 0xb1, 0xd0, 0x0e, 0x91, 0x00, 0xa6, 0xda, 0x48, 0x95, 0xfd,
	0x76, 0x66, 0x03, 0x34, 0xeb, 0x70, 0xe4, 0x39, 0xcc, 0xcb, 0x23, 0x57, 0x69, 0xc9, 0x45, 0x5c,
	0xcf, 0x79, 0x38, 0xd7, 0xe3, 0xc9, 0x33, 0x98, 0x09, 0x66, 0x4a, 0xa9, 0x4e, 0x5b, 0x29, 0x12,
	0x7e, 0xa0, 0x43, 0x1c, 0xec, 0x92, 0xe4, 0x2d, 0xf8, 0x49, 0x53, 0x9b, 0xa6, 0x23, 0xdc, 0xf4,
	0xd1, 0xf5, 0x4d, 0xed, 0x50, 0xd8, 0x16, 0x54, 0xb7, 0x2f, 0x23, 0xf1, 0x59, 0x24, 0x92, 0x8e,
	0xed, 0xed, 0x1d, 0x24, 0x9f, 0xe0, 0xc1, 0xbe, 0x5d, 0xab, 0xa6, 0x13, 0x34, 0x7f, 0xda, 0x37,
	0xef, 0xd4, 0x1f, 0x5e, 0xc8, 0xaa, 0x46, 0xb8, 0x48, 0xe4, 0x97, 0x1f, 0xbb, 0x6f, 0x14, 0x6c,
	0x23, 0x35, 0xee, 0x34, 0xe9, 0x63, 0xed, 0xe7, 0x26, 0xdf, 0xc1, 0xb4, 0x38, 0x97, 0xae, 0xe9,
	0x14, 0xe3, 0x1f, 0xf7, 0xe3, 0x5b, 0x7f, 0x8d, 0xb0, 0x23, 0xb1, 0xf6, 0xfc, 0xdf, 0x4f, 0x9e,
	0x31, 0x3a, 0xab, 0xed, 0x2d, 0xfe, 0x73, 0x87, 0xdf, 0xcf, 0xab, 0xfb, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3b, 0x8f, 0xae, 0x05, 0x59, 0x03, 0x00, 0x00,
}