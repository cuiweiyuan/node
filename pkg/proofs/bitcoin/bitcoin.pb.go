// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/pkg/proofs/bitcoin/bitcoin.proto

package bitcoin

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

type Proof struct {
	TxBytes []byte `protobuf:"bytes,1,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	Path    []byte `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Index   uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_761b957bca33df27, []int{0}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetTxBytes() []byte {
	if m != nil {
		return m.TxBytes
	}
	return nil
}

func (m *Proof) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Proof) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*Proof)(nil), "zetachain.zetacore.pkg.proofs.bitcoin.Proof")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/pkg/proofs/bitcoin/bitcoin.proto", fileDescriptor_761b957bca33df27)
}

var fileDescriptor_761b957bca33df27 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xae, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x0b, 0xb2, 0xd3, 0xf5,
	0x0b, 0x8a, 0xf2, 0xf3, 0xd3, 0x8a, 0xf5, 0x93, 0x32, 0x4b, 0x92, 0xf3, 0x33, 0xf3, 0x60, 0xb4,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x2a, 0x5c, 0x93, 0x1e, 0x4c, 0x93, 0x5e, 0x41, 0x76,
	0xba, 0x1e, 0x44, 0x93, 0x1e, 0x54, 0xb1, 0x92, 0x0f, 0x17, 0x6b, 0x00, 0x48, 0x44, 0x48, 0x92,
	0x8b, 0xa3, 0xa4, 0x22, 0x3e, 0xa9, 0xb2, 0x24, 0xb5, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27,
	0x88, 0xbd, 0xa4, 0xc2, 0x09, 0xc4, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0xc9, 0x90, 0x60,
	0x02, 0x0b, 0x83, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0xcc, 0x0a,
	0x8c, 0x1a, 0xbc, 0x41, 0x10, 0x8e, 0x93, 0xf7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x82, 0x3d, 0xa1,
	0x4b, 0xd0, 0x3f, 0x49, 0x6c, 0x60, 0x8f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xf6,
	0xf8, 0x28, 0xff, 0x00, 0x00, 0x00,
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBitcoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovBitcoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovBitcoin(uint64(m.Index))
	}
	return n
}

func sovBitcoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBitcoin(x uint64) (n int) {
	return sovBitcoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoin
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = append(m.TxBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.TxBytes == nil {
				m.TxBytes = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path[:0], dAtA[iNdEx:postIndex]...)
			if m.Path == nil {
				m.Path = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcoin
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
func skipBitcoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBitcoin
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
					return 0, ErrIntOverflowBitcoin
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
					return 0, ErrIntOverflowBitcoin
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
				return 0, ErrInvalidLengthBitcoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBitcoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBitcoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBitcoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBitcoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBitcoin = fmt.Errorf("proto: unexpected end of group")
)