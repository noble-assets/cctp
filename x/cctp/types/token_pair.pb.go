// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/token_pair.proto

package types

import (
	fmt "fmt"
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

//
// TokenPair is used to look up the Noble token (i.e. "uusdc") from a remote domain token address
// Multiple remote_domain + remote_token pairs can map to the same local_token
//
// @param remote_domain the remote domain_id corresponding to the token
// @param remote_token the remote token address
// @param local_token the corresponding Noble token denom in uunits
type TokenPair struct {
	RemoteDomain uint32 `protobuf:"varint,1,opt,name=remote_domain,json=remoteDomain,proto3" json:"remote_domain,omitempty"`
	RemoteToken  []byte `protobuf:"bytes,2,opt,name=remote_token,json=remoteToken,proto3" json:"remote_token,omitempty"`
	LocalToken   string `protobuf:"bytes,3,opt,name=local_token,json=localToken,proto3" json:"local_token,omitempty"`
}

func (m *TokenPair) Reset()         { *m = TokenPair{} }
func (m *TokenPair) String() string { return proto.CompactTextString(m) }
func (*TokenPair) ProtoMessage()    {}
func (*TokenPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a764822a752ee9, []int{0}
}
func (m *TokenPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPair.Merge(m, src)
}
func (m *TokenPair) XXX_Size() int {
	return m.Size()
}
func (m *TokenPair) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPair.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPair proto.InternalMessageInfo

func (m *TokenPair) GetRemoteDomain() uint32 {
	if m != nil {
		return m.RemoteDomain
	}
	return 0
}

func (m *TokenPair) GetRemoteToken() []byte {
	if m != nil {
		return m.RemoteToken
	}
	return nil
}

func (m *TokenPair) GetLocalToken() string {
	if m != nil {
		return m.LocalToken
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenPair)(nil), "noble.cctp.v1.TokenPair")
}

func init() { proto.RegisterFile("noble/cctp/v1/token_pair.proto", fileDescriptor_73a764822a752ee9) }

var fileDescriptor_73a764822a752ee9 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0x59, 0x4d, 0x4c, 0x6e, 0xef, 0x68, 0xa8, 0xae, 0x5a, 0x51, 0x1b, 0x1a, 0xd8, 0x5c,
	0x7c, 0x03, 0x63, 0xac, 0x0d, 0xb1, 0xb2, 0xc1, 0x05, 0x27, 0xe7, 0x46, 0xd8, 0xd9, 0x2c, 0xc3,
	0x46, 0xdf, 0xc2, 0xc7, 0xb2, 0xbc, 0xd2, 0xd2, 0xc0, 0x8b, 0x5c, 0x58, 0x28, 0xe7, 0x9b, 0x2f,
	0x7f, 0xfe, 0x9f, 0x0b, 0x83, 0x75, 0x0b, 0xb2, 0x69, 0xc8, 0x4a, 0x7f, 0x90, 0x84, 0x9f, 0x60,
	0x2a, 0xab, 0xb4, 0x2b, 0xac, 0x43, 0xc2, 0x24, 0x0e, 0xff, 0x62, 0xfe, 0x17, 0xfe, 0x70, 0x4b,
	0x7c, 0xf3, 0x32, 0x2b, 0xcf, 0x4a, 0xbb, 0xe4, 0x8e, 0xc7, 0x0e, 0x3a, 0x24, 0xa8, 0xde, 0xb1,
	0x53, 0xda, 0xec, 0x59, 0xca, 0xb2, 0xb8, 0xdc, 0x2d, 0xf0, 0x31, 0xb0, 0xe4, 0x86, 0xaf, 0x77,
	0x15, 0xb2, 0xf7, 0x17, 0x29, 0xcb, 0x76, 0xe5, 0x76, 0x61, 0x21, 0x2b, 0xb9, 0xe6, 0xdb, 0x16,
	0x1b, 0xd5, 0xae, 0xc6, 0x65, 0xca, 0xb2, 0x4d, 0xc9, 0x03, 0x0a, 0xc2, 0xc3, 0xdb, 0xef, 0x28,
	0xd8, 0x69, 0x14, 0xec, 0x7f, 0x14, 0xec, 0x67, 0x12, 0xd1, 0x69, 0x12, 0xd1, 0xdf, 0x24, 0xa2,
	0xd7, 0xa7, 0xa3, 0xa6, 0x8f, 0xa1, 0x2e, 0x1a, 0xec, 0x64, 0x4f, 0x4e, 0x99, 0x23, 0xb4, 0xe8,
	0x21, 0xf7, 0x60, 0x68, 0x70, 0xd0, 0xcb, 0x50, 0x3f, 0x9f, 0xeb, 0xe7, 0x0e, 0x07, 0x02, 0x97,
	0x5b, 0xa7, 0xbd, 0x22, 0x90, 0x5f, 0xcb, 0x68, 0xfa, 0xb6, 0xd0, 0xd7, 0x57, 0x61, 0xed, 0xfd,
	0x39, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xd6, 0x63, 0x1a, 0x0f, 0x01, 0x00, 0x00,
}

func (m *TokenPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LocalToken) > 0 {
		i -= len(m.LocalToken)
		copy(dAtA[i:], m.LocalToken)
		i = encodeVarintTokenPair(dAtA, i, uint64(len(m.LocalToken)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RemoteToken) > 0 {
		i -= len(m.RemoteToken)
		copy(dAtA[i:], m.RemoteToken)
		i = encodeVarintTokenPair(dAtA, i, uint64(len(m.RemoteToken)))
		i--
		dAtA[i] = 0x12
	}
	if m.RemoteDomain != 0 {
		i = encodeVarintTokenPair(dAtA, i, uint64(m.RemoteDomain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RemoteDomain != 0 {
		n += 1 + sovTokenPair(uint64(m.RemoteDomain))
	}
	l = len(m.RemoteToken)
	if l > 0 {
		n += 1 + l + sovTokenPair(uint64(l))
	}
	l = len(m.LocalToken)
	if l > 0 {
		n += 1 + l + sovTokenPair(uint64(l))
	}
	return n
}

func sovTokenPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenPair(x uint64) (n int) {
	return sovTokenPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenPair
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
			return fmt.Errorf("proto: TokenPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteDomain", wireType)
			}
			m.RemoteDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemoteDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteToken", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
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
				return ErrInvalidLengthTokenPair
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteToken = append(m.RemoteToken[:0], dAtA[iNdEx:postIndex]...)
			if m.RemoteToken == nil {
				m.RemoteToken = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenPair
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
				return ErrInvalidLengthTokenPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LocalToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenPair
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
func skipTokenPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenPair
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
					return 0, ErrIntOverflowTokenPair
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
					return 0, ErrIntOverflowTokenPair
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
				return 0, ErrInvalidLengthTokenPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenPair = fmt.Errorf("proto: unexpected end of group")
)
