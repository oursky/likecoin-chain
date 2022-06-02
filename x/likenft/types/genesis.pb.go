// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likechain/likenft/genesis.proto

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

// GenesisState defines the likenft module's genesis state.
type GenesisState struct {
	Params               Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ClassesByIscnList    []ClassesByISCN         `protobuf:"bytes,2,rep,name=classes_by_iscn_list,json=classesByIscnList,proto3" json:"classes_by_iscn_list"`
	ClassesByAccountList []ClassesByAccount      `protobuf:"bytes,3,rep,name=classes_by_account_list,json=classesByAccountList,proto3" json:"classes_by_account_list"`
	MintableNftList      []MintableNFT           `protobuf:"bytes,4,rep,name=mintable_nft_list,json=mintableNftList,proto3" json:"mintable_nft_list"`
	ClassRevealQueue     []ClassRevealQueueEntry `protobuf:"bytes,5,rep,name=class_reveal_queue,json=classRevealQueue,proto3" json:"class_reveal_queue"`
	OfferList            []Offer                 `protobuf:"bytes,6,rep,name=offer_list,json=offerList,proto3" json:"offer_list"`
	ListingList          []Listing               `protobuf:"bytes,7,rep,name=listing_list,json=listingList,proto3" json:"listing_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae2e28366593bfc, []int{0}
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

func (m *GenesisState) GetClassesByIscnList() []ClassesByISCN {
	if m != nil {
		return m.ClassesByIscnList
	}
	return nil
}

func (m *GenesisState) GetClassesByAccountList() []ClassesByAccount {
	if m != nil {
		return m.ClassesByAccountList
	}
	return nil
}

func (m *GenesisState) GetMintableNftList() []MintableNFT {
	if m != nil {
		return m.MintableNftList
	}
	return nil
}

func (m *GenesisState) GetClassRevealQueue() []ClassRevealQueueEntry {
	if m != nil {
		return m.ClassRevealQueue
	}
	return nil
}

func (m *GenesisState) GetOfferList() []Offer {
	if m != nil {
		return m.OfferList
	}
	return nil
}

func (m *GenesisState) GetListingList() []Listing {
	if m != nil {
		return m.ListingList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "likechain.likenft.GenesisState")
}

func init() { proto.RegisterFile("likechain/likenft/genesis.proto", fileDescriptor_fae2e28366593bfc) }

var fileDescriptor_fae2e28366593bfc = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0x5b, 0xef, 0x9c, 0x98, 0x5d, 0xd0, 0x95, 0x81, 0x75, 0x60, 0xee, 0x50, 0xc1, 0x21,
	0xd8, 0xc1, 0xf5, 0xc1, 0x27, 0x1f, 0xdc, 0xd0, 0x8b, 0x70, 0x9d, 0x73, 0x13, 0x04, 0x11, 0x6a,
	0x1a, 0xd2, 0x2e, 0xd8, 0x25, 0xb3, 0x49, 0xc5, 0x7e, 0x0b, 0x3f, 0x8d, 0x9f, 0x61, 0x8f, 0x7b,
	0xf4, 0x49, 0x64, 0xfb, 0x22, 0xd2, 0xe4, 0x6c, 0xc8, 0x5a, 0x77, 0xdf, 0x4a, 0xce, 0x2f, 0xbf,
	0x9c, 0x73, 0xfa, 0x47, 0x67, 0x29, 0xff, 0xc2, 0xe8, 0x9c, 0x70, 0x31, 0x28, 0xbf, 0x44, 0xac,
	0x07, 0x09, 0x13, 0x4c, 0x71, 0x15, 0x2c, 0x33, 0xa9, 0xa5, 0xd7, 0xde, 0x03, 0x01, 0x00, 0xdd,
	0x4e, 0x22, 0x13, 0x69, 0xaa, 0x83, 0xf2, 0xcb, 0x82, 0xdd, 0xc7, 0x55, 0x13, 0x4d, 0x89, 0x52,
	0x61, 0xc6, 0xbe, 0x31, 0x92, 0x86, 0x5f, 0x73, 0x96, 0xb3, 0x2b, 0x58, 0xa6, 0xc2, 0xa8, 0x08,
	0x09, 0xa5, 0x32, 0x17, 0x1a, 0xd8, 0x47, 0x47, 0x59, 0xae, 0xa8, 0x00, 0xb0, 0x66, 0x94, 0x94,
	0x2b, 0xcd, 0x45, 0x02, 0xc0, 0xc3, 0x2a, 0xb0, 0xe0, 0x42, 0x93, 0x28, 0x65, 0xa1, 0x88, 0x77,
	0xef, 0xdd, 0xab, 0x52, 0x32, 0x8e, 0x59, 0x06, 0x65, 0x5c, 0x2d, 0x2f, 0x49, 0x46, 0x16, 0xb0,
	0xaf, 0xfb, 0x3f, 0x1b, 0xe8, 0xf4, 0xc2, 0x6e, 0x70, 0xa6, 0x89, 0x66, 0xde, 0x33, 0xd4, 0xb4,
	0x80, 0xef, 0xf6, 0xdc, 0x7e, 0xeb, 0xfc, 0x6e, 0x50, 0xd9, 0x68, 0x30, 0x31, 0xc0, 0xb0, 0xb1,
	0xfa, 0x7d, 0xe6, 0x4c, 0x01, 0xf7, 0x3e, 0xa0, 0xce, 0xc1, 0xa0, 0x61, 0x39, 0x8f, 0x7f, 0xad,
	0x77, 0xd2, 0x6f, 0x9d, 0xf7, 0x6a, 0x34, 0x23, 0x8b, 0x0f, 0x8b, 0xd7, 0xb3, 0xd1, 0x18, 0x6c,
	0x6d, 0xba, 0x3f, 0x54, 0x54, 0x5c, 0x72, 0xa5, 0xbd, 0xcf, 0xe8, 0x4e, 0x75, 0xdb, 0xd6, 0x7d,
	0x62, 0xdc, 0x0f, 0x8e, 0xb9, 0x5f, 0x58, 0x1e, 0xf4, 0x1d, 0x7a, 0x70, 0x6e, 0x5e, 0x98, 0xa0,
	0xf6, 0xbf, 0x9b, 0xb5, 0xee, 0x86, 0x71, 0xe3, 0x1a, 0xf7, 0x1b, 0x60, 0xc7, 0xaf, 0xde, 0x83,
	0xf6, 0xd6, 0xee, 0xfa, 0x38, 0xb6, 0xc6, 0x4f, 0xc8, 0xab, 0xa6, 0xc9, 0xbf, 0x6e, 0x94, 0xfd,
	0xff, 0xb5, 0x3b, 0x35, 0xec, 0xbb, 0x12, 0x7d, 0x29, 0x74, 0x56, 0x80, 0xfc, 0x36, 0x3d, 0x28,
	0x7a, 0xcf, 0x11, 0x32, 0xff, 0xd8, 0x36, 0xda, 0x34, 0x56, 0xbf, 0xc6, 0xfa, 0xb6, 0x84, 0xc0,
	0x72, 0xd3, 0xdc, 0x30, 0xcd, 0x8d, 0xd0, 0x29, 0x24, 0xcd, 0x0a, 0x6e, 0x18, 0x41, 0xb7, 0x46,
	0x70, 0x69, 0x31, 0x50, 0xb4, 0xe0, 0x56, 0x79, 0x3a, 0xbc, 0x58, 0x6d, 0xb0, 0xbb, 0xde, 0x60,
	0xf7, 0xcf, 0x06, 0xbb, 0x3f, 0xb6, 0xd8, 0x59, 0x6f, 0xb1, 0xf3, 0x6b, 0x8b, 0x9d, 0x8f, 0x4f,
	0x12, 0xae, 0xe7, 0x79, 0x14, 0x50, 0xb9, 0x30, 0x99, 0xa3, 0x12, 0xc2, 0x67, 0x63, 0xf8, 0x7d,
	0x1f, 0x44, 0x5d, 0x2c, 0x99, 0x8a, 0x9a, 0x26, 0x88, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xca, 0x62, 0x37, 0xdb, 0x03, 0x00, 0x00,
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
	if len(m.ListingList) > 0 {
		for iNdEx := len(m.ListingList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ListingList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.OfferList) > 0 {
		for iNdEx := len(m.OfferList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OfferList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ClassRevealQueue) > 0 {
		for iNdEx := len(m.ClassRevealQueue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassRevealQueue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.MintableNftList) > 0 {
		for iNdEx := len(m.MintableNftList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MintableNftList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ClassesByAccountList) > 0 {
		for iNdEx := len(m.ClassesByAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassesByAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClassesByIscnList) > 0 {
		for iNdEx := len(m.ClassesByIscnList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClassesByIscnList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ClassesByIscnList) > 0 {
		for _, e := range m.ClassesByIscnList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassesByAccountList) > 0 {
		for _, e := range m.ClassesByAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MintableNftList) > 0 {
		for _, e := range m.MintableNftList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClassRevealQueue) > 0 {
		for _, e := range m.ClassRevealQueue {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.OfferList) > 0 {
		for _, e := range m.OfferList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ListingList) > 0 {
		for _, e := range m.ListingList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClassesByIscnList", wireType)
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
			m.ClassesByIscnList = append(m.ClassesByIscnList, ClassesByISCN{})
			if err := m.ClassesByIscnList[len(m.ClassesByIscnList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassesByAccountList", wireType)
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
			m.ClassesByAccountList = append(m.ClassesByAccountList, ClassesByAccount{})
			if err := m.ClassesByAccountList[len(m.ClassesByAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintableNftList", wireType)
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
			m.MintableNftList = append(m.MintableNftList, MintableNFT{})
			if err := m.MintableNftList[len(m.MintableNftList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassRevealQueue", wireType)
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
			m.ClassRevealQueue = append(m.ClassRevealQueue, ClassRevealQueueEntry{})
			if err := m.ClassRevealQueue[len(m.ClassRevealQueue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OfferList", wireType)
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
			m.OfferList = append(m.OfferList, Offer{})
			if err := m.OfferList[len(m.OfferList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListingList", wireType)
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
			m.ListingList = append(m.ListingList, Listing{})
			if err := m.ListingList[len(m.ListingList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
