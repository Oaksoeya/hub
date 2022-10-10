// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
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

type Params struct {
	Deposit          types.Coin                               `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	InactiveDuration time.Duration                            `protobuf:"bytes,2,opt,name=inactive_duration,json=inactiveDuration,proto3,stdduration" json:"inactive_duration"`
	MaxPrice         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=max_price,json=maxPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"max_price"`
	MinPrice         github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=min_price,json=minPrice,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_price"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a408d0240644eb, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.node.v1.Params")
}

func init() { proto.RegisterFile("sentinel/node/v1/params.proto", fileDescriptor_56a408d0240644eb) }

var fileDescriptor_56a408d0240644eb = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x93, 0xb6, 0xea, 0xdf, 0x3f, 0x2c, 0x25, 0x62, 0x28, 0x95, 0x70, 0x2b, 0xa6, 0x2e,
	0xb5, 0x29, 0x4c, 0xac, 0x85, 0x9d, 0xaa, 0x23, 0x4b, 0xe5, 0x24, 0x6e, 0x6a, 0xd1, 0xf8, 0x46,
	0xb1, 0x53, 0x95, 0xb7, 0x60, 0xe4, 0x11, 0x10, 0x4f, 0xd2, 0xb1, 0x23, 0x13, 0x85, 0x64, 0xe2,
	0x2d, 0x90, 0x9d, 0x58, 0x62, 0x62, 0x63, 0x4a, 0xac, 0x73, 0xcf, 0xf9, 0x8e, 0xad, 0xeb, 0x9d,
	0x49, 0x26, 0x14, 0x17, 0x6c, 0x4d, 0x04, 0x44, 0x8c, 0x6c, 0x26, 0x24, 0xa5, 0x19, 0x4d, 0x24,
	0x4e, 0x33, 0x50, 0xe0, 0x77, 0xad, 0x8c, 0xb5, 0x8c, 0x37, 0x93, 0xfe, 0x49, 0x0c, 0x31, 0x18,
	0x91, 0xe8, 0xbf, 0x6a, 0xae, 0x8f, 0x62, 0x80, 0x78, 0xcd, 0x88, 0x39, 0x05, 0xf9, 0x92, 0x44,
	0x79, 0x46, 0x15, 0x07, 0x61, 0xf5, 0x10, 0x64, 0x02, 0x92, 0x04, 0x54, 0x6a, 0x48, 0xc0, 0x14,
	0x9d, 0x90, 0x10, 0x78, 0xad, 0x9f, 0x7f, 0x35, 0xbc, 0xf6, 0xcc, 0x80, 0xfd, 0x6b, 0xef, 0x5f,
	0xc4, 0x52, 0x90, 0x5c, 0xf5, 0xdc, 0xa1, 0x3b, 0x3a, 0xba, 0x3c, 0xc5, 0x95, 0x19, 0x6b, 0x33,
	0xae, 0xcd, 0xf8, 0x06, 0xb8, 0x98, 0xb6, 0x76, 0xef, 0x03, 0x67, 0x6e, 0xe7, 0xfd, 0x99, 0x77,
	0xcc, 0x05, 0x0d, 0x15, 0xdf, 0xb0, 0x85, 0x2d, 0xd0, 0x6b, 0xd4, 0x21, 0x55, 0x43, 0x6c, 0x1b,
	0xe2, 0xdb, 0x7a, 0x60, 0xda, 0xd1, 0x21, 0xcf, 0x87, 0x81, 0x3b, 0xef, 0x5a, 0xb7, 0xd5, 0xfc,
	0x95, 0xf7, 0x3f, 0xa1, 0xdb, 0x45, 0x9a, 0xf1, 0x90, 0xf5, 0x9a, 0xc3, 0xe6, 0xef, 0x75, 0x2e,
	0x74, 0xd2, 0xeb, 0x61, 0x30, 0x8a, 0xb9, 0x5a, 0xe5, 0x01, 0x0e, 0x21, 0x21, 0xf5, 0xc5, 0xab,
	0xcf, 0x58, 0x46, 0x0f, 0x44, 0x3d, 0xa6, 0x4c, 0x1a, 0x83, 0x9c, 0x77, 0x12, 0xba, 0x9d, 0xe9,
	0x70, 0x43, 0xe2, 0xa2, 0x26, 0xb5, 0xfe, 0x82, 0xc4, 0x85, 0x21, 0x4d, 0xef, 0x76, 0x9f, 0xc8,
	0x79, 0x29, 0x90, 0xb3, 0x2b, 0x90, 0xbb, 0x2f, 0x90, 0xfb, 0x51, 0x20, 0xf7, 0xa9, 0x44, 0xce,
	0xbe, 0x44, 0xce, 0x5b, 0x89, 0x9c, 0xfb, 0xf1, 0x8f, 0x54, 0xbb, 0x00, 0x63, 0x58, 0x2e, 0x79,
	0xc8, 0xe9, 0x9a, 0xac, 0xf2, 0x80, 0x6c, 0xab, 0x75, 0x31, 0x80, 0xa0, 0x6d, 0xde, 0xf4, 0xea,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x70, 0x93, 0x49, 0x4c, 0x02, 0x00, 0x00,
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
	if len(m.MinPrice) > 0 {
		for iNdEx := len(m.MinPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.MaxPrice) > 0 {
		for iNdEx := len(m.MaxPrice) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MaxPrice[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.InactiveDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactiveDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.InactiveDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.MaxPrice) > 0 {
		for _, e := range m.MaxPrice {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.MinPrice) > 0 {
		for _, e := range m.MinPrice {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveDuration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.InactiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
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
			m.MaxPrice = append(m.MaxPrice, types.Coin{})
			if err := m.MaxPrice[len(m.MaxPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPrice", wireType)
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
			m.MinPrice = append(m.MinPrice, types.Coin{})
			if err := m.MinPrice[len(m.MinPrice)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
