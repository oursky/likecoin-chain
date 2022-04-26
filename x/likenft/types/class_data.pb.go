// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/class_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
	Metadata       JsonInput   `protobuf:"bytes,1,opt,name=metadata,proto3,customtype=JsonInput" json:"metadata"`
	Parent         ClassParent `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent"`
	Config         ClassConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config"`
	ClaimableCount uint64      `protobuf:"varint,4,opt,name=claimableCount,proto3" json:"claimableCount,omitempty"`
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

func (m *ClassData) GetClaimableCount() uint64 {
	if m != nil {
		return m.ClaimableCount
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

type ClassConfig struct {
	Burnable        bool   `protobuf:"varint,1,opt,name=burnable,proto3" json:"burnable,omitempty"`
	MaxSupply       uint64 `protobuf:"varint,2,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	EnablePayToMint bool   `protobuf:"varint,3,opt,name=enablePayToMint,proto3" json:"enablePayToMint,omitempty"`
	MintPrice       uint64 `protobuf:"varint,4,opt,name=mintPrice,proto3" json:"mintPrice,omitempty"`
}

func (m *ClassConfig) Reset()         { *m = ClassConfig{} }
func (m *ClassConfig) String() string { return proto.CompactTextString(m) }
func (*ClassConfig) ProtoMessage()    {}
func (*ClassConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_084dbaa322495271, []int{2}
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

func (m *ClassConfig) GetEnablePayToMint() bool {
	if m != nil {
		return m.EnablePayToMint
	}
	return false
}

func (m *ClassConfig) GetMintPrice() uint64 {
	if m != nil {
		return m.MintPrice
	}
	return 0
}

func init() {
	proto.RegisterEnum("likecoin.likechain.likenft.ClassParentType", ClassParentType_name, ClassParentType_value)
	proto.RegisterType((*ClassData)(nil), "likecoin.likechain.likenft.ClassData")
	proto.RegisterType((*ClassParent)(nil), "likecoin.likechain.likenft.ClassParent")
	proto.RegisterType((*ClassConfig)(nil), "likecoin.likechain.likenft.ClassConfig")
}

func init() { proto.RegisterFile("likenft/class_data.proto", fileDescriptor_084dbaa322495271) }

var fileDescriptor_084dbaa322495271 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0x16, 0xab, 0x75, 0x36, 0x55, 0x9b, 0xae, 0x38, 0x58, 0x11, 0x72, 0xa3, 0x1c, 0xc0,
	0x02, 0xd5, 0x96, 0x82, 0x38, 0xa3, 0xc6, 0x20, 0x14, 0x10, 0x6e, 0xb4, 0x4d, 0x41, 0xe2, 0x82,
	0x36, 0xdb, 0x4d, 0xba, 0xc2, 0xd9, 0xb5, 0xec, 0xb5, 0x94, 0x7c, 0x05, 0x7c, 0x56, 0x0f, 0x1c,
	0x7a, 0x44, 0x1c, 0x2a, 0x94, 0xfc, 0x00, 0x9f, 0x80, 0x76, 0x63, 0xbb, 0x10, 0x54, 0x29, 0xb7,
	0xd9, 0x37, 0x6f, 0xde, 0xcc, 0x9b, 0xb1, 0xa1, 0x9b, 0xf0, 0x2f, 0x4c, 0x4c, 0x54, 0x48, 0x13,
	0x92, 0xe7, 0x9f, 0x2f, 0x89, 0x22, 0x41, 0x9a, 0x49, 0x25, 0x51, 0x5b, 0x67, 0xa8, 0xe4, 0x22,
	0x30, 0xc1, 0x15, 0x29, 0x23, 0x31, 0x51, 0xed, 0x87, 0x53, 0x39, 0x95, 0x86, 0x16, 0xea, 0x68,
	0x5d, 0xd1, 0xfd, 0x0d, 0x60, 0x23, 0xd2, 0x32, 0xaf, 0x88, 0x22, 0xe8, 0x04, 0x3a, 0x33, 0xa6,
	0x88, 0x56, 0x74, 0x41, 0x07, 0xf8, 0xfb, 0xfd, 0xa3, 0xeb, 0xdb, 0x63, 0xeb, 0xe7, 0xed, 0x71,
	0xe3, 0x6d, 0x2e, 0xc5, 0x40, 0xa4, 0x85, 0xc2, 0x35, 0x05, 0xbd, 0x86, 0xbb, 0x29, 0xc9, 0x98,
	0x50, 0xee, 0x4e, 0x07, 0xf8, 0xcd, 0xde, 0x93, 0xe0, 0xfe, 0xfe, 0x81, 0xe9, 0x32, 0x34, 0xf4,
	0xbe, 0xad, 0x55, 0x71, 0x59, 0xac, 0x65, 0xa8, 0x14, 0x13, 0x3e, 0x75, 0x1f, 0x6c, 0x29, 0x13,
	0x19, 0x7a, 0x25, 0xb3, 0x2e, 0x46, 0x8f, 0xe1, 0x01, 0x4d, 0x08, 0x9f, 0x91, 0x71, 0xc2, 0x22,
	0x59, 0x08, 0xe5, 0xda, 0x1d, 0xe0, 0xdb, 0x78, 0x03, 0xed, 0x7e, 0x07, 0xb0, 0xf9, 0xd7, 0x30,
	0xe8, 0x25, 0xb4, 0xd5, 0x22, 0x65, 0xc6, 0xf0, 0x41, 0xef, 0xd9, 0x96, 0x1e, 0x46, 0x8b, 0x94,
	0x61, 0x53, 0x88, 0x7c, 0xb8, 0xcf, 0x73, 0x2a, 0x06, 0x97, 0xc3, 0x8c, 0x4d, 0xf8, 0xdc, 0x2c,
	0xa3, 0x61, 0x86, 0x03, 0xf8, 0x9f, 0x0c, 0xea, 0xc1, 0x23, 0xfd, 0xfe, 0xc0, 0xb2, 0x9c, 0x4b,
	0x71, 0xaa, 0xde, 0x73, 0xa1, 0x8c, 0x69, 0xbb, 0xa4, 0xff, 0x9f, 0x46, 0x1e, 0xdc, 0x23, 0x94,
	0xd6, 0x7e, 0x2a, 0xe1, 0x0a, 0xec, 0x7e, 0xad, 0xec, 0xac, 0x97, 0x82, 0xda, 0xd0, 0x19, 0x17,
	0x99, 0xd0, 0x7e, 0x8d, 0x25, 0x07, 0xd7, 0x6f, 0xf4, 0x08, 0x36, 0x66, 0x64, 0x7e, 0x5e, 0xa4,
	0x69, 0xb2, 0x30, 0x63, 0xda, 0xf8, 0x0e, 0x40, 0x3e, 0x3c, 0x64, 0x86, 0x37, 0x24, 0x8b, 0x91,
	0xac, 0x67, 0x73, 0xf0, 0x26, 0x6c, 0x74, 0xb8, 0x50, 0xc3, 0x8c, 0x53, 0x56, 0x6e, 0xf9, 0x0e,
	0x78, 0xfa, 0x02, 0x1e, 0x6e, 0x2c, 0x0a, 0x35, 0xe1, 0xde, 0x45, 0xfc, 0x2e, 0x3e, 0xfb, 0x18,
	0xb7, 0x2c, 0xe4, 0x40, 0x7b, 0x70, 0x1e, 0xc5, 0x2d, 0xa0, 0xe1, 0xd3, 0x28, 0x3a, 0xbb, 0x88,
	0x47, 0xad, 0x9d, 0xfe, 0x9b, 0xeb, 0xa5, 0x07, 0x6e, 0x96, 0x1e, 0xf8, 0xb5, 0xf4, 0xc0, 0xb7,
	0x95, 0x67, 0xdd, 0xac, 0x3c, 0xeb, 0xc7, 0xca, 0xb3, 0x3e, 0x9d, 0x4c, 0xb9, 0xba, 0x2a, 0xc6,
	0x01, 0x95, 0xb3, 0xb0, 0xba, 0x4e, 0x58, 0x5f, 0x27, 0x9c, 0x87, 0xd5, 0x0f, 0xa1, 0xcf, 0x91,
	0x8f, 0x77, 0xcd, 0xa7, 0xfd, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0x18, 0x66, 0x94,
	0x28, 0x03, 0x00, 0x00,
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
	if m.ClaimableCount != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.ClaimableCount))
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
	if m.MintPrice != 0 {
		i = encodeVarintClassData(dAtA, i, uint64(m.MintPrice))
		i--
		dAtA[i] = 0x20
	}
	if m.EnablePayToMint {
		i--
		if m.EnablePayToMint {
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
	if m.ClaimableCount != 0 {
		n += 1 + sovClassData(uint64(m.ClaimableCount))
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
	if m.EnablePayToMint {
		n += 2
	}
	if m.MintPrice != 0 {
		n += 1 + sovClassData(uint64(m.MintPrice))
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableCount", wireType)
			}
			m.ClaimableCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClaimableCount |= uint64(b&0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field EnablePayToMint", wireType)
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
			m.EnablePayToMint = bool(v != 0)
		case 4:
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
