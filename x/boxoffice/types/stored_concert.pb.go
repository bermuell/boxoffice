// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: boxoffice/boxoffice/stored_concert.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type StoredConcert struct {
	Index  string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Volume uint64 `protobuf:"varint,3,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (m *StoredConcert) Reset()         { *m = StoredConcert{} }
func (m *StoredConcert) String() string { return proto.CompactTextString(m) }
func (*StoredConcert) ProtoMessage()    {}
func (*StoredConcert) Descriptor() ([]byte, []int) {
	return fileDescriptor_6bccf74095168321, []int{0}
}
func (m *StoredConcert) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredConcert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredConcert.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredConcert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredConcert.Merge(m, src)
}
func (m *StoredConcert) XXX_Size() int {
	return m.Size()
}
func (m *StoredConcert) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredConcert.DiscardUnknown(m)
}

var xxx_messageInfo_StoredConcert proto.InternalMessageInfo

func (m *StoredConcert) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredConcert) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StoredConcert) GetVolume() uint64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func init() {
	proto.RegisterType((*StoredConcert)(nil), "boxoffice.boxoffice.StoredConcert")
}

func init() {
	proto.RegisterFile("boxoffice/boxoffice/stored_concert.proto", fileDescriptor_6bccf74095168321)
}

var fileDescriptor_6bccf74095168321 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xca, 0xaf, 0xc8,
	0x4f, 0x4b, 0xcb, 0x4c, 0x4e, 0xd5, 0x47, 0xb0, 0x8a, 0x4b, 0xf2, 0x8b, 0x52, 0x53, 0xe2, 0x93,
	0xf3, 0xf3, 0x92, 0x53, 0x8b, 0x4a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0xf2,
	0x7a, 0x70, 0x96, 0x52, 0x20, 0x17, 0x6f, 0x30, 0x58, 0xb1, 0x33, 0x44, 0xad, 0x90, 0x08, 0x17,
	0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0x24,
	0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x04, 0xb3, 0x85, 0xc4, 0xb8, 0xd8,
	0xca, 0xf2, 0x73, 0x4a, 0x73, 0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0x58, 0x82, 0xa0, 0x3c, 0x27,
	0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4a, 0x2d, 0xca, 0x2d, 0x4d, 0xcd, 0xc9, 0x41,
	0x72, 0x75, 0x05, 0x12, 0xbb, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x72, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xac, 0x91, 0xfd, 0xe5, 0x00, 0x00, 0x00,
}

func (m *StoredConcert) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredConcert) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredConcert) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Volume != 0 {
		i = encodeVarintStoredConcert(dAtA, i, uint64(m.Volume))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStoredConcert(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredConcert(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredConcert(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredConcert(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredConcert) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredConcert(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStoredConcert(uint64(l))
	}
	if m.Volume != 0 {
		n += 1 + sovStoredConcert(uint64(m.Volume))
	}
	return n
}

func sovStoredConcert(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredConcert(x uint64) (n int) {
	return sovStoredConcert(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredConcert) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredConcert
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
			return fmt.Errorf("proto: StoredConcert: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredConcert: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredConcert
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
				return ErrInvalidLengthStoredConcert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredConcert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredConcert
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
				return ErrInvalidLengthStoredConcert
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredConcert
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			m.Volume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredConcert
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Volume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStoredConcert(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredConcert
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
func skipStoredConcert(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredConcert
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
					return 0, ErrIntOverflowStoredConcert
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
					return 0, ErrIntOverflowStoredConcert
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
				return 0, ErrInvalidLengthStoredConcert
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredConcert
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredConcert
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredConcert        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredConcert          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredConcert = fmt.Errorf("proto: unexpected end of group")
)