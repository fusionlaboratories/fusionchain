// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/treasury/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the treasury module's genesis state.
type GenesisState struct {
	Params       Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Keys         []Key         `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys"`
	KeyRequests  []KeyRequest  `protobuf:"bytes,3,rep,name=key_requests,json=keyRequests,proto3" json:"key_requests"`
	SignRequests []SignRequest `protobuf:"bytes,4,rep,name=sign_requests,json=signRequests,proto3" json:"sign_requests"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_acebfe64ea42de8c, []int{0}
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

func (m *GenesisState) GetKeys() []Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *GenesisState) GetKeyRequests() []KeyRequest {
	if m != nil {
		return m.KeyRequests
	}
	return nil
}

func (m *GenesisState) GetSignRequests() []SignRequest {
	if m != nil {
		return m.SignRequests
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "fusionchain.treasury.GenesisState")
}

func init() {
	proto.RegisterFile("fusionchain/treasury/genesis.proto", fileDescriptor_acebfe64ea42de8c)
}

var fileDescriptor_acebfe64ea42de8c = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x13, 0x40, 0x0c, 0x86, 0xbb, 0x44, 0x0c, 0x5c, 0x54, 0xb9, 0xc0, 0xc4, 0xd2, 0x44,
	0x82, 0xad, 0x23, 0x0b, 0xaa, 0xda, 0xa1, 0x82, 0xad, 0x4b, 0x15, 0xd2, 0x53, 0x63, 0x45, 0x89,
	0x83, 0x8f, 0x23, 0xd5, 0x6f, 0xd1, 0xb1, 0x8f, 0xc4, 0xc8, 0xd8, 0xa9, 0xaa, 0x92, 0x17, 0xa9,
	0x70, 0xd2, 0x86, 0x4a, 0xee, 0x66, 0x1d, 0x7f, 0xff, 0xe7, 0xe3, 0x9f, 0x4c, 0x9f, 0x73, 0xe4,
	0x22, 0x8d, 0x76, 0x21, 0x4f, 0x03, 0x25, 0x21, 0xc4, 0x5c, 0xea, 0x80, 0x41, 0x0a, 0xc8, 0xd1,
	0xcf, 0xa4, 0x50, 0xc2, 0x1b, 0x9c, 0x31, 0xfe, 0x37, 0x33, 0x1a, 0x30, 0xc1, 0x84, 0x01, 0x82,
	0xd3, 0xa9, 0x62, 0x47, 0x13, 0xab, 0x2f, 0x0b, 0x65, 0x98, 0xd4, 0xba, 0x11, 0xb5, 0x22, 0x31,
	0xe8, 0xfa, 0xde, 0xbe, 0x52, 0x92, 0x45, 0xc8, 0x59, 0x5a, 0x31, 0xd3, 0xb7, 0x16, 0xe9, 0xaf,
	0xaa, 0x25, 0x37, 0x2a, 0x54, 0xe0, 0x5d, 0x93, 0x6e, 0xf5, 0xc8, 0xd0, 0x1d, 0xbb, 0xb3, 0xde,
	0xfc, 0xc2, 0xb7, 0x2d, 0xed, 0xdf, 0x1b, 0x66, 0xd9, 0x39, 0x7c, 0x5c, 0x3a, 0xeb, 0x3a, 0xe1,
	0x2d, 0x48, 0x27, 0x06, 0x8d, 0xc3, 0xd6, 0xb8, 0x3d, 0xeb, 0xcd, 0xff, 0xdb, 0x93, 0xb7, 0xa0,
	0xeb, 0x98, 0x81, 0xbd, 0x1b, 0xd2, 0x8f, 0x41, 0x3f, 0x4a, 0xd8, 0xe7, 0x80, 0x0a, 0x87, 0x6d,
	0x13, 0x1e, 0xff, 0x19, 0x5e, 0x57, 0x60, 0xed, 0xe8, 0xc5, 0x3f, 0x13, 0xf4, 0xee, 0xc8, 0xbf,
	0xd3, 0xd7, 0x1a, 0x57, 0xc7, 0xb8, 0x26, 0x76, 0xd7, 0x86, 0xb3, 0xf4, 0xb7, 0xac, 0x8f, 0xcd,
	0x08, 0x97, 0xab, 0x43, 0x41, 0xdd, 0x63, 0x41, 0xdd, 0xcf, 0x82, 0xba, 0xaf, 0x25, 0x75, 0x8e,
	0x25, 0x75, 0xde, 0x4b, 0xea, 0x3c, 0x5c, 0x31, 0xae, 0x76, 0xf9, 0xd6, 0x8f, 0x44, 0x12, 0xec,
	0x25, 0x3c, 0x89, 0xe0, 0xbc, 0xe9, 0x97, 0xa6, 0x6b, 0xa5, 0x33, 0xc0, 0x6d, 0xd7, 0x54, 0xbd,
	0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x60, 0x94, 0x5d, 0xaa, 0x23, 0x02, 0x00, 0x00,
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
	if len(m.SignRequests) > 0 {
		for iNdEx := len(m.SignRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.KeyRequests) > 0 {
		for iNdEx := len(m.KeyRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeyRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Keys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Keys) > 0 {
		for _, e := range m.Keys {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.KeyRequests) > 0 {
		for _, e := range m.KeyRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SignRequests) > 0 {
		for _, e := range m.SignRequests {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
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
			m.Keys = append(m.Keys, Key{})
			if err := m.Keys[len(m.Keys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyRequests", wireType)
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
			m.KeyRequests = append(m.KeyRequests, KeyRequest{})
			if err := m.KeyRequests[len(m.KeyRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignRequests", wireType)
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
			m.SignRequests = append(m.SignRequests, SignRequest{})
			if err := m.SignRequests[len(m.SignRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
