// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/dymensionxyz/dymension/sequencer/events.proto

package sequencer

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// EventIncreasedBond is an event emitted when a sequencer's bond is increased.
type EventIncreasedBond struct {
	// sequencer is the bech32-encoded address of the sequencer which increased its bond
	Sequencer string `protobuf:"bytes,1,opt,name=sequencer,proto3" json:"sequencer,omitempty"`
	// added_amount is the amount of coins added to the sequencer's bond
	AddedAmount types.Coin `protobuf:"bytes,2,opt,name=added_amount,json=addedAmount,proto3" json:"added_amount"`
	// bond is the new active bond amount of the sequencer
	Bond github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=bond,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"bond"`
}

func (m *EventIncreasedBond) Reset()         { *m = EventIncreasedBond{} }
func (m *EventIncreasedBond) String() string { return proto.CompactTextString(m) }
func (*EventIncreasedBond) ProtoMessage()    {}
func (*EventIncreasedBond) Descriptor() ([]byte, []int) {
	return fileDescriptor_299be36200df0dd4, []int{0}
}
func (m *EventIncreasedBond) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventIncreasedBond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventIncreasedBond.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventIncreasedBond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventIncreasedBond.Merge(m, src)
}
func (m *EventIncreasedBond) XXX_Size() int {
	return m.Size()
}
func (m *EventIncreasedBond) XXX_DiscardUnknown() {
	xxx_messageInfo_EventIncreasedBond.DiscardUnknown(m)
}

var xxx_messageInfo_EventIncreasedBond proto.InternalMessageInfo

func (m *EventIncreasedBond) GetSequencer() string {
	if m != nil {
		return m.Sequencer
	}
	return ""
}

func (m *EventIncreasedBond) GetAddedAmount() types.Coin {
	if m != nil {
		return m.AddedAmount
	}
	return types.Coin{}
}

func (m *EventIncreasedBond) GetBond() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Bond
	}
	return nil
}

type EventRotationStarted struct {
	// RollappId defines the rollapp to which the sequencer belongs.
	RollappId string `protobuf:"bytes,1,opt,name=rollapp_id,json=rollappId,proto3" json:"rollapp_id,omitempty"`
	// NextProposerAddr is the bech32-encoded address of the next proposer.
	// can be empty if no sequencer is available to be the next proposer.
	NextProposerAddr string `protobuf:"bytes,2,opt,name=next_proposer_addr,json=nextProposerAddr,proto3" json:"next_proposer_addr,omitempty"`
	// RewardAddr is a bech32-encoded address of the sequencer's reward address.
	RewardAddr string `protobuf:"bytes,3,opt,name=reward_addr,json=rewardAddr,proto3" json:"reward_addr,omitempty"`
	// WhitelistedRelayers is a list of the whitelisted relayer addresses. Addresses are bech32-encoded strings.
	WhitelistedRelayers []string `protobuf:"bytes,4,rep,name=whitelisted_relayers,json=whitelistedRelayers,proto3" json:"whitelisted_relayers,omitempty"`
}

func (m *EventRotationStarted) Reset()         { *m = EventRotationStarted{} }
func (m *EventRotationStarted) String() string { return proto.CompactTextString(m) }
func (*EventRotationStarted) ProtoMessage()    {}
func (*EventRotationStarted) Descriptor() ([]byte, []int) {
	return fileDescriptor_299be36200df0dd4, []int{1}
}
func (m *EventRotationStarted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRotationStarted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRotationStarted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRotationStarted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRotationStarted.Merge(m, src)
}
func (m *EventRotationStarted) XXX_Size() int {
	return m.Size()
}
func (m *EventRotationStarted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRotationStarted.DiscardUnknown(m)
}

var xxx_messageInfo_EventRotationStarted proto.InternalMessageInfo

func (m *EventRotationStarted) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

func (m *EventRotationStarted) GetNextProposerAddr() string {
	if m != nil {
		return m.NextProposerAddr
	}
	return ""
}

func (m *EventRotationStarted) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

func (m *EventRotationStarted) GetWhitelistedRelayers() []string {
	if m != nil {
		return m.WhitelistedRelayers
	}
	return nil
}

type EventUpdateRewardAddress struct {
	// Operator is the bech32-encoded address of the actor sending the update
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// RewardAddr is a bech32 encoded sdk acc address
	RewardAddr string `protobuf:"bytes,2,opt,name=reward_addr,json=rewardAddr,proto3" json:"reward_addr,omitempty"`
}

func (m *EventUpdateRewardAddress) Reset()         { *m = EventUpdateRewardAddress{} }
func (m *EventUpdateRewardAddress) String() string { return proto.CompactTextString(m) }
func (*EventUpdateRewardAddress) ProtoMessage()    {}
func (*EventUpdateRewardAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_299be36200df0dd4, []int{2}
}
func (m *EventUpdateRewardAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateRewardAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateRewardAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateRewardAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateRewardAddress.Merge(m, src)
}
func (m *EventUpdateRewardAddress) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateRewardAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateRewardAddress.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateRewardAddress proto.InternalMessageInfo

func (m *EventUpdateRewardAddress) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventUpdateRewardAddress) GetRewardAddr() string {
	if m != nil {
		return m.RewardAddr
	}
	return ""
}

type EventUpdateWhitelistedRelayers struct {
	// Operator is the bech32-encoded address of the actor sending the update
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Relayers is an array of the whitelisted relayer addresses. Addresses are bech32-encoded strings.
	Relayers []string `protobuf:"bytes,2,rep,name=relayers,proto3" json:"relayers,omitempty"`
}

func (m *EventUpdateWhitelistedRelayers) Reset()         { *m = EventUpdateWhitelistedRelayers{} }
func (m *EventUpdateWhitelistedRelayers) String() string { return proto.CompactTextString(m) }
func (*EventUpdateWhitelistedRelayers) ProtoMessage()    {}
func (*EventUpdateWhitelistedRelayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_299be36200df0dd4, []int{3}
}
func (m *EventUpdateWhitelistedRelayers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateWhitelistedRelayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateWhitelistedRelayers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateWhitelistedRelayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateWhitelistedRelayers.Merge(m, src)
}
func (m *EventUpdateWhitelistedRelayers) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateWhitelistedRelayers) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateWhitelistedRelayers.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateWhitelistedRelayers proto.InternalMessageInfo

func (m *EventUpdateWhitelistedRelayers) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EventUpdateWhitelistedRelayers) GetRelayers() []string {
	if m != nil {
		return m.Relayers
	}
	return nil
}

func init() {
	proto.RegisterType((*EventIncreasedBond)(nil), "dymensionxyz.dymension.sequencer.EventIncreasedBond")
	proto.RegisterType((*EventRotationStarted)(nil), "dymensionxyz.dymension.sequencer.EventRotationStarted")
	proto.RegisterType((*EventUpdateRewardAddress)(nil), "dymensionxyz.dymension.sequencer.EventUpdateRewardAddress")
	proto.RegisterType((*EventUpdateWhitelistedRelayers)(nil), "dymensionxyz.dymension.sequencer.EventUpdateWhitelistedRelayers")
}

func init() {
	proto.RegisterFile("types/dymensionxyz/dymension/sequencer/events.proto", fileDescriptor_299be36200df0dd4)
}

var fileDescriptor_299be36200df0dd4 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0xcd, 0x26, 0x11, 0x10, 0x87, 0x03, 0x32, 0x39, 0x6c, 0x23, 0xb1, 0x59, 0x45, 0x1c, 0x72,
	0xa0, 0x6b, 0x42, 0x25, 0xee, 0x5d, 0x84, 0x44, 0x6f, 0xc8, 0x55, 0x41, 0xe2, 0xb2, 0xf2, 0xae,
	0x47, 0xa9, 0x45, 0x62, 0x2f, 0xb6, 0xd3, 0x36, 0x7c, 0x05, 0xdf, 0xc1, 0x19, 0x89, 0x5f, 0xe8,
	0xb1, 0xe2, 0xc4, 0x09, 0x50, 0xf2, 0x05, 0xfc, 0x01, 0x5a, 0x7b, 0xbb, 0x8d, 0x82, 0xd4, 0x9e,
	0x76, 0x67, 0xe6, 0xbd, 0x99, 0x37, 0x7e, 0x36, 0x3a, 0xb0, 0xab, 0x12, 0x0c, 0xe1, 0xab, 0x05,
	0x48, 0x23, 0x94, 0xbc, 0x58, 0x7d, 0xbe, 0x09, 0x88, 0x81, 0x4f, 0x4b, 0x90, 0x05, 0x68, 0x02,
	0x67, 0x20, 0xad, 0x49, 0x4a, 0xad, 0xac, 0xc2, 0xf1, 0x36, 0x3c, 0x69, 0x82, 0xa4, 0x81, 0x0f,
	0xf7, 0x0a, 0x65, 0x16, 0xca, 0x64, 0x0e, 0x4f, 0x7c, 0xe0, 0xc9, 0xc3, 0xc1, 0x4c, 0xcd, 0x94,
	0xcf, 0x57, 0x7f, 0x75, 0xf6, 0xa9, 0xd7, 0xe1, 0x91, 0x24, 0x67, 0x06, 0xc8, 0xd9, 0x34, 0x07,
	0xcb, 0xa6, 0xa4, 0x50, 0x42, 0x7a, 0xd4, 0xf8, 0x6f, 0x80, 0xf0, 0xeb, 0x4a, 0xc9, 0x91, 0x2c,
	0x34, 0x30, 0x03, 0x3c, 0x55, 0x92, 0xe3, 0x97, 0xa8, 0xd7, 0x8c, 0x0e, 0x83, 0x38, 0x98, 0xf4,
	0xd2, 0xf0, 0xc7, 0xb7, 0xfd, 0x41, 0x3d, 0xf7, 0x90, 0x73, 0x0d, 0xc6, 0x1c, 0x5b, 0x2d, 0xe4,
	0x8c, 0xde, 0x40, 0x71, 0x8a, 0x1e, 0x32, 0xce, 0x81, 0x67, 0x6c, 0xa1, 0x96, 0xd2, 0x86, 0xed,
	0x38, 0x98, 0xf4, 0x5f, 0xec, 0x25, 0x35, 0xaf, 0x52, 0x91, 0xd4, 0x2a, 0x92, 0x57, 0x4a, 0xc8,
	0xb4, 0x7b, 0xf9, 0x6b, 0xd4, 0xa2, 0x7d, 0x47, 0x3a, 0x74, 0x1c, 0x9c, 0xa1, 0x6e, 0xae, 0x24,
	0x0f, 0x3b, 0x71, 0xe7, 0x76, 0xee, 0xf3, 0x8a, 0xfb, 0xf5, 0xf7, 0x68, 0x32, 0x13, 0xf6, 0x74,
	0x99, 0x27, 0x85, 0x5a, 0x5c, 0xaf, 0xeb, 0x3f, 0xfb, 0x86, 0x7f, 0x24, 0xee, 0x18, 0x1c, 0xc1,
	0x50, 0xd7, 0x78, 0xfc, 0x3d, 0x40, 0x03, 0xb7, 0x33, 0x55, 0x96, 0x59, 0xa1, 0xe4, 0xb1, 0x65,
	0xda, 0x02, 0xc7, 0x4f, 0x10, 0xd2, 0x6a, 0x3e, 0x67, 0x65, 0x99, 0x09, 0xee, 0xd7, 0xa6, 0xbd,
	0x3a, 0x73, 0xc4, 0xf1, 0x33, 0x84, 0x25, 0x5c, 0xd8, 0xca, 0x82, 0x52, 0x19, 0xd0, 0x19, 0xe3,
	0x5c, 0xbb, 0x15, 0x7b, 0xf4, 0x51, 0x55, 0x79, 0x5b, 0x17, 0xaa, 0xe3, 0xc1, 0x23, 0xd4, 0xd7,
	0x70, 0xce, 0x34, 0xf7, 0xb0, 0x8e, 0x83, 0x21, 0x9f, 0x72, 0x80, 0x29, 0x1a, 0x9c, 0x9f, 0x0a,
	0x0b, 0x73, 0x61, 0x2c, 0xf0, 0x4c, 0xc3, 0x9c, 0xad, 0x40, 0x9b, 0xb0, 0x1b, 0x77, 0x26, 0x3d,
	0xfa, 0x78, 0xab, 0x46, 0xeb, 0xd2, 0xf8, 0x04, 0x85, 0x4e, 0xf8, 0x49, 0xc9, 0x99, 0x05, 0xda,
	0xf4, 0x02, 0x63, 0x70, 0x88, 0xee, 0x57, 0x0e, 0x5a, 0x55, 0x1b, 0x46, 0xaf, 0xc3, 0x5d, 0x25,
	0xed, 0x5d, 0x25, 0xe3, 0x77, 0x28, 0xda, 0x6a, 0xfb, 0xfe, 0xff, 0xc1, 0xb7, 0x34, 0x1f, 0xa2,
	0x07, 0x8d, 0xf2, 0xb6, 0x53, 0xde, 0xc4, 0x69, 0x7e, 0xb9, 0x8e, 0x82, 0xab, 0x75, 0x14, 0xfc,
	0x59, 0x47, 0xc1, 0x97, 0x4d, 0xd4, 0xba, 0xda, 0x44, 0xad, 0x9f, 0x9b, 0xa8, 0xf5, 0xe1, 0xcd,
	0x96, 0x65, 0xbb, 0x2f, 0x45, 0x48, 0xeb, 0x4d, 0x23, 0x65, 0x7e, 0xe7, 0x33, 0xca, 0xef, 0xb9,
	0x7b, 0x7c, 0xf0, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xd2, 0xca, 0xec, 0x77, 0x03, 0x00, 0x00,
}

func (m *EventIncreasedBond) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventIncreasedBond) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventIncreasedBond) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bond) > 0 {
		for iNdEx := len(m.Bond) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bond[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.AddedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sequencer) > 0 {
		i -= len(m.Sequencer)
		copy(dAtA[i:], m.Sequencer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sequencer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRotationStarted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRotationStarted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRotationStarted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WhitelistedRelayers) > 0 {
		for iNdEx := len(m.WhitelistedRelayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.WhitelistedRelayers[iNdEx])
			copy(dAtA[i:], m.WhitelistedRelayers[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.WhitelistedRelayers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.RewardAddr) > 0 {
		i -= len(m.RewardAddr)
		copy(dAtA[i:], m.RewardAddr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RewardAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NextProposerAddr) > 0 {
		i -= len(m.NextProposerAddr)
		copy(dAtA[i:], m.NextProposerAddr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NextProposerAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateRewardAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateRewardAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateRewardAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardAddr) > 0 {
		i -= len(m.RewardAddr)
		copy(dAtA[i:], m.RewardAddr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.RewardAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateWhitelistedRelayers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateWhitelistedRelayers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateWhitelistedRelayers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Relayers) > 0 {
		for iNdEx := len(m.Relayers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Relayers[iNdEx])
			copy(dAtA[i:], m.Relayers[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Relayers[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventIncreasedBond) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sequencer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.AddedAmount.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.Bond) > 0 {
		for _, e := range m.Bond {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventRotationStarted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NextProposerAddr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RewardAddr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.WhitelistedRelayers) > 0 {
		for _, s := range m.WhitelistedRelayers {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventUpdateRewardAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.RewardAddr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateWhitelistedRelayers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Relayers) > 0 {
		for _, s := range m.Relayers {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventIncreasedBond) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventIncreasedBond: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventIncreasedBond: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequencer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sequencer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AddedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bond", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bond = append(m.Bond, types.Coin{})
			if err := m.Bond[len(m.Bond)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRotationStarted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRotationStarted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRotationStarted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextProposerAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextProposerAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhitelistedRelayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WhitelistedRelayers = append(m.WhitelistedRelayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateRewardAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateRewardAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateRewardAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateWhitelistedRelayers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateWhitelistedRelayers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateWhitelistedRelayers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayers = append(m.Relayers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
