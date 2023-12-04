// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/attester.proto

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

// *
// A public key used to verify message signatures
// @param attester ECDSA uncompressed public key, hex encoded
type Attester struct {
	Attester string `protobuf:"bytes,1,opt,name=attester,proto3" json:"attester,omitempty"`
}

func (m *Attester) Reset()         { *m = Attester{} }
func (m *Attester) String() string { return proto.CompactTextString(m) }
func (*Attester) ProtoMessage()    {}
func (*Attester) Descriptor() ([]byte, []int) {
	return fileDescriptor_c84b7240ac7d2071, []int{0}
}
func (m *Attester) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attester) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attester.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attester) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attester.Merge(m, src)
}
func (m *Attester) XXX_Size() int {
	return m.Size()
}
func (m *Attester) XXX_DiscardUnknown() {
	xxx_messageInfo_Attester.DiscardUnknown(m)
}

var xxx_messageInfo_Attester proto.InternalMessageInfo

func (m *Attester) GetAttester() string {
	if m != nil {
		return m.Attester
	}
	return ""
}

func init() {
	proto.RegisterType((*Attester)(nil), "circle.cctp.v1.Attester")
}

func init() { proto.RegisterFile("circle/cctp/v1/attester.proto", fileDescriptor_c84b7240ac7d2071) }

var fileDescriptor_c84b7240ac7d2071 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2c, 0x4a,
	0xce, 0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x4f, 0x2c, 0x29, 0x49, 0x2d,
	0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x48, 0xeb, 0x81, 0xa4,
	0xf5, 0xca, 0x0c, 0x95, 0xd4, 0xb8, 0x38, 0x1c, 0xa1, 0x2a, 0x84, 0xa4, 0xb8, 0x38, 0x60, 0xaa,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x87, 0x18, 0x9e, 0x96, 0x99, 0xa7, 0x9f, 0x97, 0x9f, 0x94, 0x93, 0xaa, 0x0b, 0x76, 0x44,
	0x05, 0xc4, 0x2d, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x67, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0x02, 0x78, 0x2c, 0xa7, 0x00, 0x00, 0x00,
}

func (m *Attester) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attester) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attester) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attester) > 0 {
		i -= len(m.Attester)
		copy(dAtA[i:], m.Attester)
		i = encodeVarintAttester(dAtA, i, uint64(len(m.Attester)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttester(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttester(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Attester) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Attester)
	if l > 0 {
		n += 1 + l + sovAttester(uint64(l))
	}
	return n
}

func sovAttester(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttester(x uint64) (n int) {
	return sovAttester(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Attester) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttester
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
			return fmt.Errorf("proto: Attester: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attester: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttester
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
				return ErrInvalidLengthAttester
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttester
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttester(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttester
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
func skipAttester(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttester
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
					return 0, ErrIntOverflowAttester
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
					return 0, ErrIntOverflowAttester
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
				return 0, ErrInvalidLengthAttester
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttester
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttester
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttester        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttester          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttester = fmt.Errorf("proto: unexpected end of group")
)
