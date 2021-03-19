// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: markers/internal/unknown.proto

package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyPayload struct {
	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *MyPayload) Reset()         { *m = MyPayload{} }
func (m *MyPayload) String() string { return proto.CompactTextString(m) }
func (*MyPayload) ProtoMessage()    {}
func (*MyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_unknown_e3b6b20424084bae, []int{0}
}
func (m *MyPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *MyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyPayload.Merge(dst, src)
}
func (m *MyPayload) XXX_Size() int {
	return m.Size()
}
func (m *MyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MyPayload proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MyPayload)(nil), "cockroach.errors.markers.internal.MyPayload")
}
func (m *MyPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MyPayload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Val != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUnknown(dAtA, i, uint64(m.Val))
	}
	return i, nil
}

func encodeVarintUnknown(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MyPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Val != 0 {
		n += 1 + sovUnknown(uint64(m.Val))
	}
	return n
}

func sovUnknown(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUnknown(x uint64) (n int) {
	return sovUnknown(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MyPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnknown
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MyPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MyPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			m.Val = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnknown
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Val |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnknown(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnknown
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
func skipUnknown(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnknown
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
					return 0, ErrIntOverflowUnknown
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowUnknown
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthUnknown
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUnknown
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipUnknown(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthUnknown = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnknown   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("markers/internal/unknown.proto", fileDescriptor_unknown_e3b6b20424084bae)
}

var fileDescriptor_unknown_e3b6b20424084bae = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4d, 0x2c, 0xca,
	0x4e, 0x2d, 0x2a, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0xcd, 0xcb,
	0xce, 0xcb, 0x2f, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4c, 0xce, 0x4f, 0xce,
	0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x4b, 0x2d, 0x2a, 0xca, 0x2f, 0x2a, 0xd6, 0x83, 0x6a, 0xd0,
	0x83, 0x69, 0x50, 0x92, 0xe5, 0xe2, 0xf4, 0xad, 0x0c, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x11,
	0x12, 0xe0, 0x62, 0x2e, 0x4b, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x9d,
	0xb4, 0x4e, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x1b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0x38, 0x60, 0x46, 0x25, 0xb1, 0x81, 0x2d, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x85, 0x26, 0xa4, 0xd9, 0x96, 0x00, 0x00, 0x00,
}
