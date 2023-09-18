// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/blackbird/policy.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type Policy struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The actual policy informations. It must be one the supported policy types:
	// - BlackbirdPolicy
	Policy *types.Any `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_50403b059b43cc26, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetPolicy() *types.Any {
	if m != nil {
		return m.Policy
	}
	return nil
}

type BlackbirdPolicy struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BlackbirdPolicy) Reset()         { *m = BlackbirdPolicy{} }
func (m *BlackbirdPolicy) String() string { return proto.CompactTextString(m) }
func (*BlackbirdPolicy) ProtoMessage()    {}
func (*BlackbirdPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_50403b059b43cc26, []int{1}
}
func (m *BlackbirdPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlackbirdPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlackbirdPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlackbirdPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlackbirdPolicy.Merge(m, src)
}
func (m *BlackbirdPolicy) XXX_Size() int {
	return m.Size()
}
func (m *BlackbirdPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_BlackbirdPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_BlackbirdPolicy proto.InternalMessageInfo

func (m *BlackbirdPolicy) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlackbirdPolicyPayload struct {
	Witness []byte `protobuf:"bytes,1,opt,name=witness,proto3" json:"witness,omitempty"`
}

func (m *BlackbirdPolicyPayload) Reset()         { *m = BlackbirdPolicyPayload{} }
func (m *BlackbirdPolicyPayload) String() string { return proto.CompactTextString(m) }
func (*BlackbirdPolicyPayload) ProtoMessage()    {}
func (*BlackbirdPolicyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_50403b059b43cc26, []int{2}
}
func (m *BlackbirdPolicyPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlackbirdPolicyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlackbirdPolicyPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlackbirdPolicyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlackbirdPolicyPayload.Merge(m, src)
}
func (m *BlackbirdPolicyPayload) XXX_Size() int {
	return m.Size()
}
func (m *BlackbirdPolicyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BlackbirdPolicyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BlackbirdPolicyPayload proto.InternalMessageInfo

func (m *BlackbirdPolicyPayload) GetWitness() []byte {
	if m != nil {
		return m.Witness
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "fusionchain.blackbird.Policy")
	proto.RegisterType((*BlackbirdPolicy)(nil), "fusionchain.blackbird.BlackbirdPolicy")
	proto.RegisterType((*BlackbirdPolicyPayload)(nil), "fusionchain.blackbird.BlackbirdPolicyPayload")
}

func init() {
	proto.RegisterFile("fusionchain/blackbird/policy.proto", fileDescriptor_50403b059b43cc26)
}

var fileDescriptor_50403b059b43cc26 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbf, 0x4a, 0xc4, 0x30,
	0x1c, 0xc7, 0x9b, 0x5a, 0x2a, 0x46, 0x51, 0x08, 0x2a, 0xd5, 0x21, 0x94, 0x82, 0xd0, 0x41, 0x12,
	0x38, 0x9f, 0xc0, 0x9b, 0x1c, 0x8f, 0x8e, 0xb7, 0x25, 0x4d, 0xae, 0x17, 0xec, 0x25, 0xb5, 0x4d,
	0xd1, 0xbe, 0x85, 0x8f, 0xe5, 0x78, 0xa3, 0xa3, 0xb4, 0x2f, 0x22, 0xa4, 0xad, 0x14, 0xb7, 0xdf,
	0x0f, 0x3e, 0xdf, 0x3f, 0x7c, 0x61, 0xb2, 0x6b, 0x1b, 0x65, 0x74, 0xbe, 0x67, 0x4a, 0x53, 0x5e,
	0xb2, 0xfc, 0x95, 0xab, 0x5a, 0xd0, 0xca, 0x94, 0x2a, 0xef, 0x48, 0x55, 0x1b, 0x6b, 0xd0, 0xcd,
	0x82, 0x21, 0x7f, 0xcc, 0xfd, 0x5d, 0x61, 0x4c, 0x51, 0x4a, 0xea, 0x20, 0xde, 0xee, 0x28, 0xd3,
	0x93, 0x22, 0xd9, 0xc2, 0x70, 0xe3, 0x1c, 0xd0, 0x25, 0xf4, 0x95, 0x88, 0x40, 0x0c, 0xd2, 0x20,
	0xf3, 0x95, 0x40, 0x08, 0x06, 0x9a, 0x1d, 0x64, 0xe4, 0xc7, 0x20, 0x3d, 0xcb, 0xdc, 0x8d, 0x1e,
	0x61, 0x38, 0xe6, 0x45, 0x27, 0x31, 0x48, 0xcf, 0x57, 0xd7, 0x64, 0x74, 0x26, 0xb3, 0x33, 0x79,
	0xd6, 0x5d, 0x36, 0x31, 0xc9, 0x03, 0xbc, 0x5a, 0xcf, 0x1d, 0xa6, 0x10, 0x04, 0x03, 0xc1, 0x2c,
	0x73, 0x31, 0x17, 0x99, 0xbb, 0x93, 0x15, 0xbc, 0xfd, 0x87, 0x6d, 0x58, 0x57, 0x1a, 0x26, 0x50,
	0x04, 0x4f, 0xdf, 0x95, 0xd5, 0xb2, 0x69, 0x26, 0xc1, 0xfc, 0xae, 0x5f, 0xbe, 0x7a, 0x0c, 0x8e,
	0x3d, 0x06, 0x3f, 0x3d, 0x06, 0x9f, 0x03, 0xf6, 0x8e, 0x03, 0xf6, 0xbe, 0x07, 0xec, 0x6d, 0x49,
	0xa1, 0xec, 0xbe, 0xe5, 0x24, 0x37, 0x07, 0xfa, 0x56, 0x4b, 0x61, 0xe8, 0x72, 0xb7, 0x8f, 0xc5,
	0x72, 0xb6, 0xab, 0x64, 0xc3, 0x43, 0x57, 0xfd, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xfe,
	0xea, 0xc5, 0x5f, 0x01, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlackbirdPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlackbirdPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlackbirdPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlackbirdPolicyPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlackbirdPolicyPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlackbirdPolicyPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Witness) > 0 {
		i -= len(m.Witness)
		copy(dAtA[i:], m.Witness)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Witness)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPolicy(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func (m *BlackbirdPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func (m *BlackbirdPolicyPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Witness)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types.Any{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BlackbirdPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BlackbirdPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlackbirdPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BlackbirdPolicyPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BlackbirdPolicyPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlackbirdPolicyPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Witness", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Witness = append(m.Witness[:0], dAtA[iNdEx:postIndex]...)
			if m.Witness == nil {
				m.Witness = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
				return 0, ErrInvalidLengthPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicy = fmt.Errorf("proto: unexpected end of group")
)
