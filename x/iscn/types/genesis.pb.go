// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iscn/genesis.proto

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

type GenesisState struct {
	Params           Params                         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ContentIdRecords []GenesisState_ContentIdRecord `protobuf:"bytes,2,rep,name=content_id_records,json=contentIdRecords,proto3" json:"content_id_records"`
	IscnRecords      []IscnInput                    `protobuf:"bytes,3,rep,name=iscn_records,json=iscnRecords,proto3,customtype=IscnInput" json:"iscn_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_540482c90f072702, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetContentIdRecords() []GenesisState_ContentIdRecord {
	if m != nil {
		return m.ContentIdRecords
	}
	return nil
}

type GenesisState_ContentIdRecord struct {
	IscnId        string `protobuf:"bytes,1,opt,name=iscn_id,json=iscnId,proto3" json:"iscn_id,omitempty"`
	Owner         string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	LatestVersion uint64 `protobuf:"varint,3,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
}

func (m *GenesisState_ContentIdRecord) Reset()         { *m = GenesisState_ContentIdRecord{} }
func (m *GenesisState_ContentIdRecord) String() string { return proto.CompactTextString(m) }
func (*GenesisState_ContentIdRecord) ProtoMessage()    {}
func (*GenesisState_ContentIdRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_540482c90f072702, []int{0, 0}
}
func (m *GenesisState_ContentIdRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState_ContentIdRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState_ContentIdRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState_ContentIdRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState_ContentIdRecord.Merge(m, src)
}
func (m *GenesisState_ContentIdRecord) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState_ContentIdRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState_ContentIdRecord.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState_ContentIdRecord proto.InternalMessageInfo

func (m *GenesisState_ContentIdRecord) GetIscnId() string {
	if m != nil {
		return m.IscnId
	}
	return ""
}

func (m *GenesisState_ContentIdRecord) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisState_ContentIdRecord) GetLatestVersion() uint64 {
	if m != nil {
		return m.LatestVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "likechain.iscn.GenesisState")
	proto.RegisterType((*GenesisState_ContentIdRecord)(nil), "likechain.iscn.GenesisState.ContentIdRecord")
}

func init() { proto.RegisterFile("iscn/genesis.proto", fileDescriptor_540482c90f072702) }

var fileDescriptor_540482c90f072702 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x37, 0x86, 0x18, 0x0a, 0xa2, 0x34, 0x44, 0x17, 0x0e, 0x63, 0x31, 0x31, 0xd9, 0x41,
	0xb7, 0x04, 0xf8, 0x04, 0x78, 0x30, 0x4b, 0x3c, 0x98, 0x99, 0x78, 0xf0, 0x82, 0xa3, 0x6b, 0x46,
	0x23, 0xb4, 0xcb, 0x5a, 0x50, 0xbf, 0x85, 0x5f, 0xca, 0x84, 0x23, 0x47, 0xe3, 0x81, 0x18, 0xf6,
	0x45, 0x4c, 0xdb, 0x85, 0x08, 0xb7, 0xbe, 0xff, 0xfb, 0xff, 0x7f, 0x2f, 0x7d, 0x0f, 0x40, 0xc2,
	0x11, 0x0d, 0x52, 0x4c, 0x31, 0x27, 0xdc, 0xcf, 0x72, 0x26, 0x18, 0x6c, 0xcd, 0xc8, 0x2b, 0x46,
	0xd3, 0x98, 0x50, 0x5f, 0x76, 0xbb, 0x9d, 0x94, 0xa5, 0x4c, 0xb5, 0x02, 0xf9, 0xd2, 0xae, 0x6e,
	0x5b, 0x25, 0xb3, 0x38, 0x8f, 0xe7, 0x65, 0xf0, 0xf2, 0xab, 0x02, 0x9a, 0x77, 0x1a, 0xf5, 0x28,
	0x62, 0x81, 0xe1, 0x10, 0xd4, 0xb4, 0xc1, 0x36, 0x5d, 0xd3, 0x6b, 0xf4, 0xcf, 0xfd, 0x7d, 0xb4,
	0xff, 0xa0, 0xba, 0xa3, 0xea, 0x6a, 0xd3, 0x33, 0xa2, 0xd2, 0x0b, 0x5f, 0x00, 0x44, 0x8c, 0x0a,
	0x4c, 0xc5, 0x98, 0x24, 0xe3, 0x1c, 0x23, 0x96, 0x27, 0xdc, 0xae, 0xb8, 0x96, 0xd7, 0xe8, 0x5f,
	0x1f, 0x12, 0xfe, 0xcf, 0xf3, 0x6f, 0x75, 0x2c, 0x4c, 0x22, 0x15, 0x2a, 0xb9, 0x67, 0x68, 0x5f,
	0xe6, 0x70, 0x08, 0x9a, 0x32, 0xbc, 0x63, 0x5b, 0xae, 0xe5, 0x35, 0x47, 0x6d, 0xe9, 0xfe, 0xd9,
	0xf4, 0xea, 0x21, 0x47, 0x34, 0xa4, 0xd9, 0x42, 0x44, 0x0d, 0x69, 0x2b, 0x53, 0xdd, 0x14, 0x9c,
	0x1e, 0x0c, 0x80, 0x17, 0xe0, 0x58, 0x81, 0x48, 0xa2, 0x7e, 0x58, 0x8f, 0x6a, 0xb2, 0x0c, 0x13,
	0xd8, 0x01, 0x47, 0xec, 0x8d, 0xe2, 0xdc, 0xae, 0x28, 0x59, 0x17, 0xf0, 0x0a, 0xb4, 0x66, 0xb1,
	0xc0, 0x5c, 0x8c, 0x97, 0x38, 0xe7, 0x84, 0x51, 0xdb, 0x72, 0x4d, 0xaf, 0x1a, 0x9d, 0x68, 0xf5,
	0x49, 0x8b, 0xa3, 0xfb, 0xd5, 0xd6, 0x31, 0xd7, 0x5b, 0xc7, 0xfc, 0xdd, 0x3a, 0xe6, 0x67, 0xe1,
	0x18, 0xeb, 0xc2, 0x31, 0xbe, 0x0b, 0xc7, 0x78, 0xee, 0xa7, 0x44, 0x4c, 0x17, 0x13, 0x1f, 0xb1,
	0x79, 0xa0, 0x16, 0xc1, 0x08, 0xdd, 0x3d, 0x6e, 0xd4, 0x5a, 0x82, 0xe5, 0x20, 0x78, 0x0f, 0xd4,
	0x71, 0xc4, 0x47, 0x86, 0xf9, 0xa4, 0xa6, 0x8e, 0x33, 0xf8, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0x48, 0xcf, 0xe1, 0xeb, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IscnRecords) > 0 {
		for iNdEx := len(m.IscnRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.IscnRecords[iNdEx].Size()
				i -= size
				if _, err := m.IscnRecords[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ContentIdRecords) > 0 {
		for iNdEx := len(m.ContentIdRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContentIdRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisState_ContentIdRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState_ContentIdRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState_ContentIdRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestVersion != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LatestVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IscnId) > 0 {
		i -= len(m.IscnId)
		copy(dAtA[i:], m.IscnId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IscnId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ContentIdRecords) > 0 {
		for _, e := range m.ContentIdRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IscnRecords) > 0 {
		for _, e := range m.IscnRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState_ContentIdRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IscnId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LatestVersion != 0 {
		n += 1 + sovGenesis(uint64(m.LatestVersion))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentIdRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentIdRecords = append(m.ContentIdRecords, GenesisState_ContentIdRecord{})
			if err := m.ContentIdRecords[len(m.ContentIdRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnRecords", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v IscnInput
			m.IscnRecords = append(m.IscnRecords, v)
			if err := m.IscnRecords[len(m.IscnRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState_ContentIdRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ContentIdRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentIdRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IscnId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IscnId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestVersion", wireType)
			}
			m.LatestVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestVersion |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
