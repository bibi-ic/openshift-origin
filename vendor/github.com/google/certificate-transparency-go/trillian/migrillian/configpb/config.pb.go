// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package configpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import configpb "github.com/google/certificate-transparency-go/trillian/ctfe/configpb"
import keyspb "github.com/google/trillian/crypto/keyspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MigrationConfig describes the configuration options for a single CT log
// migration instance.
type MigrationConfig struct {
	// The URI of the source CT log, e.g. "https://ct.googleapis.com/pilot".
	SourceUri string `protobuf:"bytes,1,opt,name=source_uri,json=sourceUri,proto3" json:"source_uri,omitempty"`
	// The public key of the source log.
	PublicKey *keyspb.PublicKey `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// The name of the backend which this log migrates to. The name must be one of
	// those defined in the LogBackendSet.
	LogBackendName string `protobuf:"bytes,3,opt,name=log_backend_name,json=logBackendName,proto3" json:"log_backend_name,omitempty"`
	// The ID of a Trillian PREORDERED_LOG tree that stores the log data.
	LogId int64 `protobuf:"varint,4,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	// Max number of entries per get-entries request from the source log.
	BatchSize int32 `protobuf:"varint,5,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// Determines whether the migration should run continuously, i.e. watch and
	// follow the updates of the source log's STH. For example, this mode can be
	// used to support a mirror CT log.
	IsContinuous bool `protobuf:"varint,6,opt,name=is_continuous,json=isContinuous,proto3" json:"is_continuous,omitempty"`
	// The log entry index to start fetching at. Ignored in continuous mode which
	// starts at the point where it stopped (e.g. the current Trillian tree size
	// in a simple case).
	StartIndex int64 `protobuf:"varint,7,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	// The log index to end fetching at, non-inclusive. If zero, fetch up to the
	// source log's current STH. Ignored in continuous mode which keeps updating
	// STH and fetching up to that.
	EndIndex             int64    `protobuf:"varint,8,opt,name=end_index,json=endIndex,proto3" json:"end_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MigrationConfig) Reset()         { *m = MigrationConfig{} }
func (m *MigrationConfig) String() string { return proto.CompactTextString(m) }
func (*MigrationConfig) ProtoMessage()    {}
func (*MigrationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_bf213aaca32ca340, []int{0}
}
func (m *MigrationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationConfig.Unmarshal(m, b)
}
func (m *MigrationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationConfig.Marshal(b, m, deterministic)
}
func (dst *MigrationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationConfig.Merge(dst, src)
}
func (m *MigrationConfig) XXX_Size() int {
	return xxx_messageInfo_MigrationConfig.Size(m)
}
func (m *MigrationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationConfig proto.InternalMessageInfo

func (m *MigrationConfig) GetSourceUri() string {
	if m != nil {
		return m.SourceUri
	}
	return ""
}

func (m *MigrationConfig) GetPublicKey() *keyspb.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *MigrationConfig) GetLogBackendName() string {
	if m != nil {
		return m.LogBackendName
	}
	return ""
}

func (m *MigrationConfig) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

func (m *MigrationConfig) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *MigrationConfig) GetIsContinuous() bool {
	if m != nil {
		return m.IsContinuous
	}
	return false
}

func (m *MigrationConfig) GetStartIndex() int64 {
	if m != nil {
		return m.StartIndex
	}
	return 0
}

func (m *MigrationConfig) GetEndIndex() int64 {
	if m != nil {
		return m.EndIndex
	}
	return 0
}

// MigrationConfigSet is a set of MigrationConfig messages.
type MigrationConfigSet struct {
	Config               []*MigrationConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MigrationConfigSet) Reset()         { *m = MigrationConfigSet{} }
func (m *MigrationConfigSet) String() string { return proto.CompactTextString(m) }
func (*MigrationConfigSet) ProtoMessage()    {}
func (*MigrationConfigSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_bf213aaca32ca340, []int{1}
}
func (m *MigrationConfigSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationConfigSet.Unmarshal(m, b)
}
func (m *MigrationConfigSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationConfigSet.Marshal(b, m, deterministic)
}
func (dst *MigrationConfigSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationConfigSet.Merge(dst, src)
}
func (m *MigrationConfigSet) XXX_Size() int {
	return xxx_messageInfo_MigrationConfigSet.Size(m)
}
func (m *MigrationConfigSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationConfigSet.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationConfigSet proto.InternalMessageInfo

func (m *MigrationConfigSet) GetConfig() []*MigrationConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// MigrillianConfig holds configuration for multiple migration / mirroring jobs.
type MigrillianConfig struct {
	// The set of backends that this configuration will use to send requests to.
	// The names of the backends in the LogBackendSet must all be distinct.
	Backends *configpb.LogBackendSet `protobuf:"bytes,1,opt,name=backends,proto3" json:"backends,omitempty"`
	// The set of migrations that will use the above backends. All the protos in
	// it must set a valid log_backend_name for the config to be usable.
	MigrationConfigs     *MigrationConfigSet `protobuf:"bytes,2,opt,name=migration_configs,json=migrationConfigs,proto3" json:"migration_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MigrillianConfig) Reset()         { *m = MigrillianConfig{} }
func (m *MigrillianConfig) String() string { return proto.CompactTextString(m) }
func (*MigrillianConfig) ProtoMessage()    {}
func (*MigrillianConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_bf213aaca32ca340, []int{2}
}
func (m *MigrillianConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrillianConfig.Unmarshal(m, b)
}
func (m *MigrillianConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrillianConfig.Marshal(b, m, deterministic)
}
func (dst *MigrillianConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrillianConfig.Merge(dst, src)
}
func (m *MigrillianConfig) XXX_Size() int {
	return xxx_messageInfo_MigrillianConfig.Size(m)
}
func (m *MigrillianConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrillianConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MigrillianConfig proto.InternalMessageInfo

func (m *MigrillianConfig) GetBackends() *configpb.LogBackendSet {
	if m != nil {
		return m.Backends
	}
	return nil
}

func (m *MigrillianConfig) GetMigrationConfigs() *MigrationConfigSet {
	if m != nil {
		return m.MigrationConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*MigrationConfig)(nil), "configpb.MigrationConfig")
	proto.RegisterType((*MigrationConfigSet)(nil), "configpb.MigrationConfigSet")
	proto.RegisterType((*MigrillianConfig)(nil), "configpb.MigrillianConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_bf213aaca32ca340) }

var fileDescriptor_config_bf213aaca32ca340 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0x8e, 0x94, 0x30,
	0x14, 0xc6, 0xc3, 0x8e, 0x33, 0x42, 0x67, 0xd5, 0xd9, 0x26, 0x46, 0x5c, 0x35, 0x92, 0xf1, 0x86,
	0x9b, 0x05, 0x9d, 0x8d, 0x2f, 0xe0, 0x5e, 0x98, 0x89, 0x7f, 0xa2, 0x4c, 0xbc, 0x26, 0xa5, 0x9c,
	0xe9, 0x9e, 0x0c, 0xb4, 0xa4, 0x2d, 0x89, 0xec, 0x63, 0xf8, 0x0e, 0xbe, 0xa7, 0xa1, 0xc0, 0xac,
	0xae, 0xd9, 0x2b, 0xe0, 0xfb, 0xbe, 0xfe, 0x38, 0xf9, 0x4e, 0xc9, 0x29, 0x57, 0x72, 0x8f, 0x22,
	0x69, 0xb4, 0xb2, 0x8a, 0xfa, 0xc3, 0x57, 0x53, 0x9c, 0x7f, 0x17, 0x68, 0xaf, 0xdb, 0x22, 0xe1,
	0xaa, 0x4e, 0x85, 0x52, 0xa2, 0x82, 0x94, 0x83, 0xb6, 0xb8, 0x47, 0xce, 0x2c, 0x5c, 0x58, 0xcd,
	0xa4, 0x69, 0x98, 0x06, 0xc9, 0xbb, 0x0b, 0xa1, 0x52, 0xab, 0xb1, 0xaa, 0x90, 0xc9, 0x94, 0xdb,
	0x3d, 0xa4, 0x13, 0x25, 0xfd, 0x1b, 0x7e, 0xfe, 0xfe, 0x7f, 0xe4, 0xed, 0x31, 0xdd, 0x35, 0x56,
	0xa5, 0x07, 0xe8, 0x4c, 0x53, 0x8c, 0x8f, 0xe1, 0xd8, 0xfa, 0xf7, 0x09, 0x79, 0xf2, 0x05, 0x85,
	0x66, 0x16, 0x95, 0xbc, 0x72, 0x40, 0xfa, 0x8a, 0x10, 0xa3, 0x5a, 0xcd, 0x21, 0x6f, 0x35, 0x86,
	0x5e, 0xe4, 0xc5, 0x41, 0x16, 0x0c, 0xca, 0x0f, 0x8d, 0xf4, 0x2d, 0x21, 0x4d, 0x5b, 0x54, 0xc8,
	0xf3, 0x03, 0x74, 0xe1, 0x49, 0xe4, 0xc5, 0xcb, 0xcd, 0x59, 0x32, 0x52, 0xbf, 0x39, 0xe7, 0x13,
	0x74, 0x59, 0xd0, 0x4c, 0xaf, 0x34, 0x26, 0xab, 0x4a, 0x89, 0xbc, 0x60, 0xfc, 0x00, 0xb2, 0xcc,
	0x25, 0xab, 0x21, 0x9c, 0x39, 0xec, 0xe3, 0x4a, 0x89, 0x0f, 0x83, 0xfc, 0x95, 0xd5, 0x40, 0x9f,
	0x92, 0x45, 0x9f, 0xc4, 0x32, 0x7c, 0x10, 0x79, 0xf1, 0x2c, 0x9b, 0x57, 0x4a, 0x6c, 0xcb, 0x7e,
	0xa2, 0x82, 0x59, 0x7e, 0x9d, 0x1b, 0xbc, 0x81, 0x70, 0x1e, 0x79, 0xf1, 0x3c, 0x0b, 0x9c, 0xb2,
	0xc3, 0x1b, 0xa0, 0x6f, 0xc8, 0x23, 0x34, 0x39, 0x57, 0xd2, 0xa2, 0x6c, 0x55, 0x6b, 0xc2, 0x45,
	0xe4, 0xc5, 0x7e, 0x76, 0x8a, 0xe6, 0xea, 0xa8, 0xd1, 0xd7, 0x64, 0x69, 0x2c, 0xd3, 0x36, 0x47,
	0x59, 0xc2, 0xcf, 0xf0, 0xa1, 0xe3, 0x13, 0x27, 0x6d, 0x7b, 0x85, 0xbe, 0x20, 0x41, 0x3f, 0xdd,
	0x60, 0xfb, 0xce, 0xf6, 0x41, 0x96, 0xce, 0x5c, 0x7f, 0x24, 0xf4, 0x4e, 0x4d, 0x3b, 0xb0, 0xf4,
	0x1d, 0x59, 0x0c, 0x4b, 0x08, 0xbd, 0x68, 0x16, 0x2f, 0x37, 0xcf, 0x93, 0x69, 0x39, 0xc9, 0x9d,
	0x74, 0x36, 0x06, 0xd7, 0xbf, 0x3c, 0xb2, 0xea, 0xbd, 0x61, 0x35, 0x63, 0xe3, 0x97, 0xc4, 0x1f,
	0xcb, 0x31, 0xae, 0xef, 0xe5, 0xe6, 0xd9, 0x2d, 0xe9, 0xf3, 0xb1, 0xa2, 0x1d, 0xd8, 0xec, 0x18,
	0xa4, 0x5b, 0x72, 0x56, 0x4f, 0x3f, 0xc9, 0x87, 0xb4, 0x19, 0xd7, 0xf1, 0xf2, 0xde, 0x39, 0x7a,
	0xc4, 0xaa, 0xfe, 0x57, 0x33, 0xc5, 0xc2, 0x5d, 0x86, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x51, 0xcc, 0x28, 0x13, 0xb0, 0x02, 0x00, 0x00,
}