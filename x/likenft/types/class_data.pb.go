// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/class_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClassParentType int32

const (
	ClassParentType_UNKNOWN ClassParentType = 0
	ClassParentType_ISCN    ClassParentType = 1
	ClassParentType_ACCOUNT ClassParentType = 2
)

var ClassParentType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ISCN",
	2: "ACCOUNT",
}

var ClassParentType_value = map[string]int32{
	"UNKNOWN": 0,
	"ISCN":    1,
	"ACCOUNT": 2,
}

func (x ClassParentType) String() string {
	return proto.EnumName(ClassParentType_name, int32(x))
}

func (ClassParentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{0}
}

type ClassData struct {
	Metadata      JsonInput   `protobuf:"bytes,1,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
	Parent        ClassParent `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent"`
	Config        ClassConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config"`
	MintableCount uint64      `protobuf:"varint,4,opt,name=mintableCount,proto3" json:"mintableCount,omitempty"`
}

func (m *ClassData) Reset()         { *m = ClassData{} }
func (m *ClassData) String() string { return proto.CompactTextString(m) }
func (*ClassData) ProtoMessage()    {}
func (*ClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{0}
}
func (m *ClassData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassData.Merge(m, src)
}
func (m *ClassData) XXX_Size() int {
	return m.Size()
}
func (m *ClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassData.DiscardUnknown(m)
}

var xxx_messageInfo_ClassData proto.InternalMessageInfo

func (m *ClassData) GetParent() ClassParent {
	if m != nil {
		return m.Parent
	}
	return ClassParent{}
}

func (m *ClassData) GetConfig() ClassConfig {
	if m != nil {
		return m.Config
	}
	return ClassConfig{}
}

func (m *ClassData) GetMintableCount() uint64 {
	if m != nil {
		return m.MintableCount
	}
	return 0
}

type ClassParent struct {
	Type              ClassParentType `protobuf:"varint,1,opt,name=type,proto3,enum=likecoin.likechain.likenft.ClassParentType" json:"type,omitempty"`
	IscnIdPrefix      string          `protobuf:"bytes,2,opt,name=iscnIdPrefix,proto3" json:"iscnIdPrefix,omitempty"`
	IscnVersionAtMint uint64          `protobuf:"varint,3,opt,name=iscnVersionAtMint,proto3" json:"iscnVersionAtMint,omitempty"`
	Account           string          `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *ClassParent) Reset()         { *m = ClassParent{} }
func (m *ClassParent) String() string { return proto.CompactTextString(m) }
func (*ClassParent) ProtoMessage()    {}
func (*ClassParent) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{1}
}
func (m *ClassParent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassParent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassParent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassParent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassParent.Merge(m, src)
}
func (m *ClassParent) XXX_Size() int {
	return m.Size()
}
func (m *ClassParent) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassParent.DiscardUnknown(m)
}

var xxx_messageInfo_ClassParent proto.InternalMessageInfo

func (m *ClassParent) GetType() ClassParentType {
	if m != nil {
		return m.Type
	}
	return ClassParentType_UNKNOWN
}

func (m *ClassParent) GetIscnIdPrefix() string {
	if m != nil {
		return m.IscnIdPrefix
	}
	return ""
}

func (m *ClassParent) GetIscnVersionAtMint() uint64 {
	if m != nil {
		return m.IscnVersionAtMint
	}
	return 0
}

func (m *ClassParent) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type MintPeriod struct {
	StartTime        *time.Time `protobuf:"bytes,1,opt,name=startTime,proto3,stdtime" json:"startTime,omitempty"`
	AllowedAddresses []string   `protobuf:"bytes,2,rep,name=allowedAddresses,proto3" json:"allowedAddresses,omitempty"`
	MintPrice        uint64     `protobuf:"varint,3,opt,name=mintPrice,proto3" json:"mintPrice,omitempty"`
}

func (m *MintPeriod) Reset()         { *m = MintPeriod{} }
func (m *MintPeriod) String() string { return proto.CompactTextString(m) }
func (*MintPeriod) ProtoMessage()    {}
func (*MintPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{2}
}
func (m *MintPeriod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintPeriod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintPeriod.Merge(m, src)
}
func (m *MintPeriod) XXX_Size() int {
	return m.Size()
}
func (m *MintPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_MintPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_MintPeriod proto.InternalMessageInfo

func (m *MintPeriod) GetStartTime() *time.Time {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *MintPeriod) GetAllowedAddresses() []string {
	if m != nil {
		return m.AllowedAddresses
	}
	return nil
}

func (m *MintPeriod) GetMintPrice() uint64 {
	if m != nil {
		return m.MintPrice
	}
	return 0
}

type ClassConfig struct {
	Burnable       bool         `protobuf:"varint,1,opt,name=burnable,proto3" json:"burnable,omitempty"`
	MaxSupply      uint64       `protobuf:"varint,2,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	EnableBlindBox bool         `protobuf:"varint,3,opt,name=enableBlindBox,proto3" json:"enableBlindBox,omitempty"`
	MintPeriods    []MintPeriod `protobuf:"bytes,4,rep,name=mintPeriods,proto3" json:"mintPeriods"`
	RevealTime     *time.Time   `protobuf:"bytes,5,opt,name=revealTime,proto3,stdtime" json:"revealTime,omitempty"`
}

func (m *ClassConfig) Reset()         { *m = ClassConfig{} }
func (m *ClassConfig) String() string { return proto.CompactTextString(m) }
func (*ClassConfig) ProtoMessage()    {}
func (*ClassConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{3}
}
func (m *ClassConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassConfig.Merge(m, src)
}
func (m *ClassConfig) XXX_Size() int {
	return m.Size()
}
func (m *ClassConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClassConfig proto.InternalMessageInfo

func (m *ClassConfig) GetBurnable() bool {
	if m != nil {
		return m.Burnable
	}
	return false
}

func (m *ClassConfig) GetMaxSupply() uint64 {
	if m != nil {
		return m.MaxSupply
	}
	return 0
}

func (m *ClassConfig) GetEnableBlindBox() bool {
	if m != nil {
		return m.EnableBlindBox
	}
	return false
}

func (m *ClassConfig) GetMintPeriods() []MintPeriod {
	if m != nil {
		return m.MintPeriods
	}
	return nil
}

func (m *ClassConfig) GetRevealTime() *time.Time {
	if m != nil {
		return m.RevealTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("likecoin.likechain.likenft.ClassParentType", ClassParentType_name, ClassParentType_value)
	proto.RegisterType((*ClassData)(nil), "likecoin.likechain.likenft.ClassData")
	proto.RegisterType((*ClassParent)(nil), "likecoin.likechain.likenft.ClassParent")
	proto.RegisterType((*MintPeriod)(nil), "likecoin.likechain.likenft.MintPeriod")
	proto.RegisterType((*ClassConfig)(nil), "likecoin.likechain.likenft.ClassConfig")
}

func init() { proto.RegisterFile("likenft/class_data.proto", fileDescriptor_084dbaa322495271) }

var fileDescriptor_084dbaa322495271 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xae, 0xbb, 0xfc, 0xb6, 0xd6, 0xdd, 0x6f, 0x74, 0x16, 0x87, 0xa8, 0x42, 0x69, 0x55, 0xa1,
	0x51, 0x0d, 0x2d, 0x91, 0x8a, 0xb8, 0x02, 0x4b, 0x40, 0x68, 0x20, 0xb2, 0x2a, 0xeb, 0x40, 0xe2,
	0x82, 0xdc, 0xc4, 0xcd, 0x2c, 0x12, 0x3b, 0x8a, 0x1d, 0xe8, 0x8e, 0x7c, 0x83, 0x9d, 0xf8, 0x4c,
	0x3b, 0x70, 0xd8, 0x11, 0x71, 0x18, 0x68, 0xbb, 0xf3, 0x19, 0x90, 0x9d, 0xa6, 0xfb, 0x27, 0xd0,
	0x6e, 0xf1, 0xe3, 0xe7, 0x7d, 0xde, 0xf7, 0x79, 0xfc, 0x2a, 0xd0, 0x4c, 0xe8, 0x47, 0xc2, 0xa6,
	0xd2, 0x09, 0x13, 0x2c, 0xc4, 0x87, 0x08, 0x4b, 0x6c, 0x67, 0x39, 0x97, 0x1c, 0x75, 0xd4, 0x4d,
	0xc8, 0x29, 0xb3, 0xf5, 0xc7, 0x01, 0x9e, 0x7f, 0xb1, 0xa9, 0xec, 0xdc, 0x8d, 0x79, 0xcc, 0x35,
	0xcd, 0x51, 0x5f, 0x65, 0x45, 0xa7, 0x1b, 0x73, 0x1e, 0x27, 0xc4, 0xd1, 0xa7, 0x49, 0x31, 0x75,
	0x24, 0x4d, 0x89, 0x90, 0x38, 0xcd, 0x4a, 0x42, 0xff, 0x37, 0x80, 0x4d, 0x4f, 0xf5, 0x79, 0x8e,
	0x25, 0x46, 0x5b, 0xb0, 0x91, 0x12, 0x89, 0x55, 0x4b, 0x13, 0xf4, 0xc0, 0x60, 0xd5, 0x5d, 0x3f,
	0x3e, 0xed, 0xd6, 0x7e, 0x9c, 0x76, 0x9b, 0xaf, 0x04, 0x67, 0x3b, 0x2c, 0x2b, 0x64, 0xb0, 0xa0,
	0xa0, 0x17, 0x70, 0x39, 0xc3, 0x39, 0x61, 0xd2, 0xac, 0xf7, 0xc0, 0xa0, 0x35, 0x7c, 0x60, 0xff,
	0x7d, 0x40, 0x5b, 0x77, 0x19, 0x69, 0xba, 0x6b, 0x28, 0xd5, 0x60, 0x5e, 0xac, 0x64, 0x42, 0xce,
	0xa6, 0x34, 0x36, 0x97, 0x6e, 0x29, 0xe3, 0x69, 0x7a, 0x25, 0x53, 0x16, 0xa3, 0xfb, 0xf0, 0xff,
	0x94, 0x32, 0x89, 0x27, 0x09, 0xf1, 0x78, 0xc1, 0xa4, 0x69, 0xf4, 0xc0, 0xc0, 0x08, 0xae, 0x82,
	0xfd, 0x6f, 0x00, 0xb6, 0x2e, 0x8d, 0x82, 0x9e, 0x42, 0x43, 0x1e, 0x66, 0x44, 0xdb, 0x5d, 0x1b,
	0x3e, 0xbc, 0xa5, 0x83, 0xf1, 0x61, 0x46, 0x02, 0x5d, 0x88, 0x06, 0x70, 0x95, 0x8a, 0x90, 0xed,
	0x44, 0xa3, 0x9c, 0x4c, 0xe9, 0x4c, 0x47, 0xd1, 0xd4, 0xa3, 0x81, 0xe0, 0xca, 0x0d, 0x1a, 0xc2,
	0x75, 0x75, 0x7e, 0x4b, 0x72, 0x41, 0x39, 0xdb, 0x96, 0x6f, 0x28, 0x93, 0xda, 0xb2, 0x31, 0xa7,
	0xdf, 0xbc, 0x46, 0x16, 0x5c, 0xc1, 0x61, 0xb8, 0xb0, 0x53, 0x09, 0x57, 0x60, 0xff, 0x2b, 0x80,
	0x50, 0x11, 0x47, 0x24, 0xa7, 0x3c, 0x42, 0x4f, 0x60, 0x53, 0x48, 0x9c, 0xcb, 0x31, 0x4d, 0x4b,
	0x4b, 0xad, 0x61, 0xc7, 0x2e, 0x77, 0xc0, 0xae, 0x76, 0xc0, 0x1e, 0x57, 0x3b, 0xe0, 0x1a, 0x47,
	0x3f, 0xbb, 0x20, 0xb8, 0x28, 0x41, 0x9b, 0xb0, 0x8d, 0x93, 0x84, 0x7f, 0x26, 0xd1, 0x76, 0x14,
	0xe5, 0x44, 0x08, 0x22, 0xcc, 0x7a, 0x6f, 0x69, 0xd0, 0x0c, 0x6e, 0xe0, 0xe8, 0x1e, 0x6c, 0xaa,
	0x68, 0x47, 0x39, 0x0d, 0x49, 0x69, 0x23, 0xb8, 0x00, 0xfa, 0x5f, 0xea, 0xf3, 0x9c, 0xcb, 0xb7,
	0x42, 0x1d, 0xd8, 0x98, 0x14, 0x39, 0x53, 0x0f, 0xa1, 0x07, 0x6b, 0x04, 0x8b, 0xb3, 0x56, 0xc2,
	0xb3, 0xbd, 0x22, 0xcb, 0x92, 0x43, 0x9d, 0x9f, 0x52, 0xaa, 0x00, 0xb4, 0x01, 0xd7, 0x88, 0xe6,
	0xb9, 0x09, 0x65, 0x91, 0xcb, 0x67, 0xba, 0x59, 0x23, 0xb8, 0x86, 0x22, 0x1f, 0xb6, 0xd2, 0x45,
	0x12, 0xc2, 0x34, 0x7a, 0x4b, 0x83, 0xd6, 0x70, 0xe3, 0x5f, 0x0f, 0x7a, 0x11, 0xdc, 0x7c, 0x95,
	0x2e, 0x0b, 0xa0, 0x67, 0x10, 0xe6, 0xe4, 0x13, 0xc1, 0x89, 0x0e, 0xf3, 0xbf, 0x5b, 0x86, 0x79,
	0xa9, 0x66, 0xf3, 0x31, 0xbc, 0x73, 0x6d, 0x67, 0x50, 0x0b, 0xae, 0xec, 0xfb, 0xaf, 0xfd, 0xdd,
	0x77, 0x7e, 0xbb, 0x86, 0x1a, 0xd0, 0xd8, 0xd9, 0xf3, 0xfc, 0x36, 0x50, 0xf0, 0xb6, 0xe7, 0xed,
	0xee, 0xfb, 0xe3, 0x76, 0xdd, 0x7d, 0x79, 0x7c, 0x66, 0x81, 0x93, 0x33, 0x0b, 0xfc, 0x3a, 0xb3,
	0xc0, 0xd1, 0xb9, 0x55, 0x3b, 0x39, 0xb7, 0x6a, 0xdf, 0xcf, 0xad, 0xda, 0xfb, 0xad, 0x98, 0xca,
	0x83, 0x62, 0x62, 0x87, 0x3c, 0x75, 0x2a, 0x5f, 0xce, 0xc2, 0x97, 0x33, 0x73, 0xaa, 0x5f, 0x87,
	0xda, 0x4c, 0x31, 0x59, 0xd6, 0x53, 0x3e, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x26, 0x83, 0xdc,
	0xb7, 0x52, 0x04, 0x00, 0x00,
}

func (m *ClassData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintableCount != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.MintableCount))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClassData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Parent.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClassData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Metadata.Size()
		i -= size
		if _, err := m.Metadata.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClassData(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ClassParent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassParent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassParent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintClassData(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x22
	}
	if m.IscnVersionAtMint != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.IscnVersionAtMint))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IscnIdPrefix) > 0 {
		i -= len(m.IscnIdPrefix)
		copy(dAtA[i:], m.IscnIdPrefix)
		i = encodeVarintClassData(dAtA, i, uint64(len(m.IscnIdPrefix)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MintPeriod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintPeriod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintPeriod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintPrice != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.MintPrice))
		i--
		dAtA[i] = 0x18
	}
	if len(m.AllowedAddresses) > 0 {
		for iNdEx := len(m.AllowedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedAddresses[iNdEx])
			copy(dAtA[i:], m.AllowedAddresses[iNdEx])
			i = encodeVarintClassData(dAtA, i, uint64(len(m.AllowedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartTime != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintClassData(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClassConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevealTime != nil {
		n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.RevealTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.RevealTime):])
		if err4 != nil {
			return 0, err4
		}
		i -= n4
		i = encodeVarintClassData(dAtA, i, uint64(n4))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MintPeriods) > 0 {
		for iNdEx := len(m.MintPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClassData(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EnableBlindBox {
		i--
		if m.EnableBlindBox {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.MaxSupply != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.MaxSupply))
		i--
		dAtA[i] = 0x10
	}
	if m.Burnable {
		i--
		if m.Burnable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintClassData(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Metadata.Size()
	n += 1 + l + sovClassData(uint64(l))
	l = m.Parent.Size()
	n += 1 + l + sovClassData(uint64(l))
	l = m.Config.Size()
	n += 1 + l + sovClassData(uint64(l))
	if m.MintableCount != 0 {
		n += 1 + sovClassData(uint64(m.MintableCount))
	}
	return n
}

func (m *ClassParent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovClassData(uint64(m.Type))
	}
	l = len(m.IscnIdPrefix)
	if l > 0 {
		n += 1 + l + sovClassData(uint64(l))
	}
	if m.IscnVersionAtMint != 0 {
		n += 1 + sovClassData(uint64(m.IscnVersionAtMint))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovClassData(uint64(l))
	}
	return n
}

func (m *MintPeriod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime)
		n += 1 + l + sovClassData(uint64(l))
	}
	if len(m.AllowedAddresses) > 0 {
		for _, s := range m.AllowedAddresses {
			l = len(s)
			n += 1 + l + sovClassData(uint64(l))
		}
	}
	if m.MintPrice != 0 {
		n += 1 + sovClassData(uint64(m.MintPrice))
	}
	return n
}

func (m *ClassConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Burnable {
		n += 2
	}
	if m.MaxSupply != 0 {
		n += 1 + sovClassData(uint64(m.MaxSupply))
	}
	if m.EnableBlindBox {
		n += 2
	}
	if len(m.MintPeriods) > 0 {
		for _, e := range m.MintPeriods {
			l = e.Size()
			n += 1 + l + sovClassData(uint64(l))
		}
	}
	if m.RevealTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.RevealTime)
		n += 1 + l + sovClassData(uint64(l))
	}
	return n
}

func sovClassData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassData(x uint64) (n int) {
	return sovClassData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Parent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintableCount", wireType)
			}
			m.MintableCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintableCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClassParent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassParent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassParent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ClassParentType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnIdPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IscnIdPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnVersionAtMint", wireType)
			}
			m.IscnVersionAtMint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IscnVersionAtMint |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MintPeriod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MintPeriod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintPeriod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedAddresses = append(m.AllowedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintPrice", wireType)
			}
			m.MintPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MintPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClassConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassData
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClassConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burnable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Burnable = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSupply", wireType)
			}
			m.MaxSupply = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSupply |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableBlindBox", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableBlindBox = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintPeriods = append(m.MintPeriods, MintPeriod{})
			if err := m.MintPeriods[len(m.MintPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClassData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RevealTime == nil {
				m.RevealTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.RevealTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassData
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClassData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassData
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClassData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassData = fmt.Errorf("proto: unexpected end of group")
)
