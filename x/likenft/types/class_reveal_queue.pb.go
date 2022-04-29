// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: likenft/class_reveal_queue.proto

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

type ClassRevealQueueEntry struct {
	RevealTime time.Time `protobuf:"bytes,1,opt,name=revealTime,proto3,stdtime" json:"revealTime"`
	ClassId    string    `protobuf:"bytes,2,opt,name=classId,proto3" json:"classId,omitempty"`
}

func (m *ClassRevealQueueEntry) Reset()         { *m = ClassRevealQueueEntry{} }
func (m *ClassRevealQueueEntry) String() string { return proto.CompactTextString(m) }
func (*ClassRevealQueueEntry) ProtoMessage()    {}
func (*ClassRevealQueueEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca4ae5e06215be6f, []int{0}
}
func (m *ClassRevealQueueEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClassRevealQueueEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassRevealQueueEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClassRevealQueueEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassRevealQueueEntry.Merge(m, src)
}
func (m *ClassRevealQueueEntry) XXX_Size() int {
	return m.Size()
}
func (m *ClassRevealQueueEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassRevealQueueEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ClassRevealQueueEntry proto.InternalMessageInfo

func (m *ClassRevealQueueEntry) GetRevealTime() time.Time {
	if m != nil {
		return m.RevealTime
	}
	return time.Time{}
}

func (m *ClassRevealQueueEntry) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClassRevealQueueEntry)(nil), "likecoin.likechain.likenft.ClassRevealQueueEntry")
}

func init() { proto.RegisterFile("likenft/class_reveal_queue.proto", fileDescriptor_ca4ae5e06215be6f) }

var fileDescriptor_ca4ae5e06215be6f = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x50, 0x3d, 0x4e, 0x84, 0x40,
	0x14, 0x66, 0x2c, 0xfc, 0x19, 0x3b, 0xa2, 0x09, 0xa1, 0x18, 0x88, 0xd5, 0x36, 0xce, 0x24, 0x7a,
	0x83, 0x55, 0x63, 0x2c, 0x25, 0x5b, 0xd9, 0x6c, 0x00, 0x1f, 0xb3, 0x13, 0x81, 0x41, 0x78, 0xa8,
	0x7b, 0x8b, 0x3d, 0xd6, 0x96, 0x5b, 0x5a, 0xa9, 0x81, 0x8b, 0x18, 0x66, 0x1c, 0xb3, 0xdd, 0xf7,
	0x92, 0xef, 0xf7, 0xd1, 0xb8, 0x54, 0x2f, 0x50, 0x17, 0x28, 0xf2, 0x32, 0xed, 0xba, 0x65, 0x0b,
	0x6f, 0x90, 0x96, 0xcb, 0xd7, 0x1e, 0x7a, 0xe0, 0x4d, 0xab, 0x51, 0xfb, 0xe1, 0xc4, 0xc8, 0xb5,
	0xaa, 0xb9, 0x01, 0xab, 0xf4, 0x0f, 0xd5, 0x05, 0x86, 0x67, 0x52, 0x4b, 0x6d, 0x68, 0x62, 0x42,
	0x56, 0x11, 0x46, 0x52, 0x6b, 0x59, 0x82, 0x30, 0x57, 0xd6, 0x17, 0x02, 0x55, 0x05, 0x1d, 0xa6,
	0x55, 0x63, 0x09, 0x17, 0xef, 0xf4, 0xfc, 0x66, 0x8a, 0x4b, 0x4c, 0xda, 0xe3, 0x14, 0x76, 0x57,
	0x63, 0xbb, 0xf6, 0x6f, 0x29, 0xb5, 0x0d, 0x16, 0xaa, 0x82, 0x80, 0xc4, 0x64, 0x76, 0x7a, 0x15,
	0x72, 0x6b, 0xc7, 0x9d, 0x1d, 0x5f, 0x38, 0xbb, 0xf9, 0xf1, 0xf6, 0x2b, 0xf2, 0x36, 0xdf, 0x11,
	0x49, 0xf6, 0x74, 0x7e, 0x40, 0x8f, 0xcc, 0x9a, 0x87, 0xe7, 0xe0, 0x20, 0x26, 0xb3, 0x93, 0xc4,
	0x9d, 0xf3, 0xfb, 0xed, 0xc0, 0xc8, 0x6e, 0x60, 0xe4, 0x67, 0x60, 0x64, 0x33, 0x32, 0x6f, 0x37,
	0x32, 0xef, 0x73, 0x64, 0xde, 0xd3, 0xa5, 0x54, 0xb8, 0xea, 0x33, 0x9e, 0xeb, 0x4a, 0xb8, 0xc1,
	0xe2, 0x7f, 0xb0, 0xf8, 0x10, 0xee, 0x4f, 0xb8, 0x6e, 0xa0, 0xcb, 0x0e, 0x4d, 0x99, 0xeb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x6e, 0x10, 0x0d, 0x3f, 0x01, 0x00, 0x00,
}

func (m *ClassRevealQueueEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassRevealQueueEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClassRevealQueueEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintClassRevealQueue(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.RevealTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.RevealTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClassRevealQueue(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClassRevealQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassRevealQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClassRevealQueueEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.RevealTime)
	n += 1 + l + sovClassRevealQueue(uint64(l))
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovClassRevealQueue(uint64(l))
	}
	return n
}

func sovClassRevealQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassRevealQueue(x uint64) (n int) {
	return sovClassRevealQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClassRevealQueueEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassRevealQueue
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
			return fmt.Errorf("proto: ClassRevealQueueEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassRevealQueueEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevealTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassRevealQueue
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
				return ErrInvalidLengthClassRevealQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassRevealQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.RevealTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassRevealQueue
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
				return ErrInvalidLengthClassRevealQueue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClassRevealQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassRevealQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassRevealQueue
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
func skipClassRevealQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassRevealQueue
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
					return 0, ErrIntOverflowClassRevealQueue
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
					return 0, ErrIntOverflowClassRevealQueue
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
				return 0, ErrInvalidLengthClassRevealQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassRevealQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassRevealQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassRevealQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassRevealQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassRevealQueue = fmt.Errorf("proto: unexpected end of group")
)
