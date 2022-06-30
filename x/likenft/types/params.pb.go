// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	PriceDenom             string        `protobuf:"bytes,1,opt,name=price_denom,json=priceDenom,proto3" json:"price_denom,omitempty"`
	FeePerByte             types.DecCoin `protobuf:"bytes,2,opt,name=fee_per_byte,json=feePerByte,proto3" json:"fee_per_byte"`
	MaxOfferDurationDays   uint64        `protobuf:"varint,3,opt,name=max_offer_duration_days,json=maxOfferDurationDays,proto3" json:"max_offer_duration_days,omitempty"`
	MaxListingDurationDays uint64        `protobuf:"varint,4,opt,name=max_listing_duration_days,json=maxListingDurationDays,proto3" json:"max_listing_duration_days,omitempty"`
	MaxRoyaltyBasisPoints  uint64        `protobuf:"varint,5,opt,name=max_royalty_basis_points,json=maxRoyaltyBasisPoints,proto3" json:"max_royalty_basis_points,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d35d1816b46ebf9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *Params) GetFeePerByte() types.DecCoin {
	if m != nil {
		return m.FeePerByte
	}
	return types.DecCoin{}
}

func (m *Params) GetMaxOfferDurationDays() uint64 {
	if m != nil {
		return m.MaxOfferDurationDays
	}
	return 0
}

func (m *Params) GetMaxListingDurationDays() uint64 {
	if m != nil {
		return m.MaxListingDurationDays
	}
	return 0
}

func (m *Params) GetMaxRoyaltyBasisPoints() uint64 {
	if m != nil {
		return m.MaxRoyaltyBasisPoints
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "likechain.likenft.Params")
}

func init() { proto.RegisterFile("likechain/likenft/params.proto", fileDescriptor_7d35d1816b46ebf9) }

var fileDescriptor_7d35d1816b46ebf9 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x0b, 0xd3, 0x30,
	0x18, 0x86, 0xdb, 0x59, 0x07, 0x66, 0x5e, 0x2c, 0x53, 0xeb, 0x90, 0x6c, 0x78, 0xda, 0xc5, 0x96,
	0x39, 0x86, 0xe8, 0xb1, 0xf6, 0x28, 0x6c, 0xf4, 0xe8, 0x25, 0xa4, 0xdd, 0xd7, 0x2e, 0xb8, 0x24,
	0x25, 0xc9, 0x46, 0xfb, 0x2f, 0x3c, 0x7a, 0xdc, 0xcf, 0xd9, 0x71, 0x47, 0x4f, 0x22, 0xdb, 0x1f,
	0x91, 0xa4, 0x63, 0xa0, 0xb7, 0x97, 0xef, 0xe1, 0xc9, 0x0b, 0x79, 0x11, 0xde, 0xb3, 0xef, 0x50,
	0xee, 0x28, 0x13, 0x89, 0x4d, 0xa2, 0x32, 0x49, 0x43, 0x15, 0xe5, 0x3a, 0x6e, 0x94, 0x34, 0x32,
	0x7c, 0xf1, 0xe0, 0xf1, 0x9d, 0x4f, 0x70, 0x29, 0x35, 0x97, 0x3a, 0x29, 0xa8, 0x86, 0xe4, 0xb8,
	0x28, 0xc0, 0xd0, 0x45, 0x52, 0x4a, 0x26, 0x7a, 0x65, 0x32, 0xae, 0x65, 0x2d, 0x5d, 0x4c, 0x6c,
	0xea, 0xaf, 0xef, 0x4e, 0x03, 0x34, 0xdc, 0xb8, 0x97, 0xc3, 0x29, 0x1a, 0x35, 0x8a, 0x95, 0x40,
	0xb6, 0x20, 0x24, 0x8f, 0xfc, 0x99, 0x3f, 0x7f, 0x96, 0x23, 0x77, 0xca, 0xec, 0x25, 0xcc, 0xd0,
	0xf3, 0x0a, 0x80, 0x34, 0xa0, 0x48, 0xd1, 0x19, 0x88, 0x06, 0x33, 0x7f, 0x3e, 0xfa, 0xf0, 0x36,
	0xee, 0x8b, 0x63, 0x5b, 0x1c, 0xdf, 0x8b, 0xe3, 0x0c, 0xca, 0x2f, 0x92, 0x89, 0x34, 0x38, 0xff,
	0x9e, 0x7a, 0x39, 0xaa, 0x00, 0x36, 0xa0, 0xd2, 0xce, 0x40, 0xb8, 0x42, 0xaf, 0x39, 0x6d, 0x89,
	0xac, 0x2a, 0x50, 0x64, 0x7b, 0x50, 0xd4, 0x30, 0x29, 0xc8, 0x96, 0x76, 0x3a, 0x7a, 0x32, 0xf3,
	0xe7, 0x41, 0x3e, 0xe6, 0xb4, 0x5d, 0x5b, 0x9a, 0xdd, 0x61, 0x46, 0x3b, 0x1d, 0x7e, 0x42, 0x6f,
	0xac, 0xb6, 0x67, 0xda, 0x30, 0x51, 0xff, 0x27, 0x06, 0x4e, 0x7c, 0xc5, 0x69, 0xfb, 0xb5, 0xe7,
	0xff, 0xa8, 0x1f, 0x51, 0x64, 0x55, 0x25, 0x3b, 0xba, 0x37, 0x1d, 0x29, 0xa8, 0x66, 0x9a, 0x34,
	0x92, 0x09, 0xa3, 0xa3, 0xa7, 0xce, 0x7c, 0xc9, 0x69, 0x9b, 0xf7, 0x38, 0xb5, 0x74, 0xe3, 0xe0,
	0xe7, 0xe0, 0xe7, 0x69, 0xea, 0xa5, 0xeb, 0xf3, 0x15, 0xfb, 0x97, 0x2b, 0xf6, 0xff, 0x5c, 0xb1,
	0xff, 0xe3, 0x86, 0xbd, 0xcb, 0x0d, 0x7b, 0xbf, 0x6e, 0xd8, 0xfb, 0xb6, 0xaa, 0x99, 0xd9, 0x1d,
	0x8a, 0xb8, 0x94, 0xdc, 0xcd, 0x64, 0x7f, 0xfb, 0x11, 0xde, 0xf7, 0xf3, 0x1d, 0x97, 0x49, 0xfb,
	0xd8, 0xd0, 0x74, 0x0d, 0xe8, 0x62, 0xe8, 0xbe, 0x7e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5f,
	0x98, 0xa3, 0x33, 0xe5, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxRoyaltyBasisPoints != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxRoyaltyBasisPoints))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxListingDurationDays != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxListingDurationDays))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxOfferDurationDays != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxOfferDurationDays))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.FeePerByte.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.FeePerByte.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxOfferDurationDays != 0 {
		n += 1 + sovParams(uint64(m.MaxOfferDurationDays))
	}
	if m.MaxListingDurationDays != 0 {
		n += 1 + sovParams(uint64(m.MaxListingDurationDays))
	}
	if m.MaxRoyaltyBasisPoints != 0 {
		n += 1 + sovParams(uint64(m.MaxRoyaltyBasisPoints))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePerByte", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePerByte.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOfferDurationDays", wireType)
			}
			m.MaxOfferDurationDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOfferDurationDays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxListingDurationDays", wireType)
			}
			m.MaxListingDurationDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxListingDurationDays |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRoyaltyBasisPoints", wireType)
			}
			m.MaxRoyaltyBasisPoints = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRoyaltyBasisPoints |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
