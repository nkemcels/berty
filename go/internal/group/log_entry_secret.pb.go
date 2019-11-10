// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-internal/log_entry_secret.proto

package group

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
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

type SecretEntryEnvelope struct {
	EncryptedSecretPayload []byte   `protobuf:"bytes,1,opt,name=encrypted_secret_payload,json=encryptedSecretPayload,proto3" json:"encrypted_secret_payload,omitempty"`
	SecretPayloadSignature []byte   `protobuf:"bytes,2,opt,name=secret_payload_signature,json=secretPayloadSignature,proto3" json:"secret_payload_signature,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SecretEntryEnvelope) Reset()         { *m = SecretEntryEnvelope{} }
func (m *SecretEntryEnvelope) String() string { return proto.CompactTextString(m) }
func (*SecretEntryEnvelope) ProtoMessage()    {}
func (*SecretEntryEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1711af9e257970, []int{0}
}
func (m *SecretEntryEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretEntryEnvelope.Unmarshal(m, b)
}
func (m *SecretEntryEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretEntryEnvelope.Marshal(b, m, deterministic)
}
func (m *SecretEntryEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretEntryEnvelope.Merge(m, src)
}
func (m *SecretEntryEnvelope) XXX_Size() int {
	return xxx_messageInfo_SecretEntryEnvelope.Size(m)
}
func (m *SecretEntryEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretEntryEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_SecretEntryEnvelope proto.InternalMessageInfo

func (m *SecretEntryEnvelope) GetEncryptedSecretPayload() []byte {
	if m != nil {
		return m.EncryptedSecretPayload
	}
	return nil
}

func (m *SecretEntryEnvelope) GetSecretPayloadSignature() []byte {
	if m != nil {
		return m.SecretPayloadSignature
	}
	return nil
}

type SecretEntryPayload struct {
	DestMemberPubKey      []byte   `protobuf:"bytes,1,opt,name=dest_member_pub_key,json=destMemberPubKey,proto3" json:"dest_member_pub_key,omitempty"`
	SenderDevicePubKey    []byte   `protobuf:"bytes,2,opt,name=sender_device_pub_key,json=senderDevicePubKey,proto3" json:"sender_device_pub_key,omitempty"`
	EncryptedDeviceSecret []byte   `protobuf:"bytes,3,opt,name=encrypted_device_secret,json=encryptedDeviceSecret,proto3" json:"encrypted_device_secret,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SecretEntryPayload) Reset()         { *m = SecretEntryPayload{} }
func (m *SecretEntryPayload) String() string { return proto.CompactTextString(m) }
func (*SecretEntryPayload) ProtoMessage()    {}
func (*SecretEntryPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1711af9e257970, []int{1}
}
func (m *SecretEntryPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretEntryPayload.Unmarshal(m, b)
}
func (m *SecretEntryPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretEntryPayload.Marshal(b, m, deterministic)
}
func (m *SecretEntryPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretEntryPayload.Merge(m, src)
}
func (m *SecretEntryPayload) XXX_Size() int {
	return xxx_messageInfo_SecretEntryPayload.Size(m)
}
func (m *SecretEntryPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretEntryPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SecretEntryPayload proto.InternalMessageInfo

func (m *SecretEntryPayload) GetDestMemberPubKey() []byte {
	if m != nil {
		return m.DestMemberPubKey
	}
	return nil
}

func (m *SecretEntryPayload) GetSenderDevicePubKey() []byte {
	if m != nil {
		return m.SenderDevicePubKey
	}
	return nil
}

func (m *SecretEntryPayload) GetEncryptedDeviceSecret() []byte {
	if m != nil {
		return m.EncryptedDeviceSecret
	}
	return nil
}

type DeviceSecret struct {
	DerivationState      []byte   `protobuf:"bytes,1,opt,name=derivation_state,json=derivationState,proto3" json:"derivation_state,omitempty"`
	Counter              uint64   `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSecret) Reset()         { *m = DeviceSecret{} }
func (m *DeviceSecret) String() string { return proto.CompactTextString(m) }
func (*DeviceSecret) ProtoMessage()    {}
func (*DeviceSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e1711af9e257970, []int{2}
}
func (m *DeviceSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSecret.Unmarshal(m, b)
}
func (m *DeviceSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSecret.Marshal(b, m, deterministic)
}
func (m *DeviceSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSecret.Merge(m, src)
}
func (m *DeviceSecret) XXX_Size() int {
	return xxx_messageInfo_DeviceSecret.Size(m)
}
func (m *DeviceSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSecret.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSecret proto.InternalMessageInfo

func (m *DeviceSecret) GetDerivationState() []byte {
	if m != nil {
		return m.DerivationState
	}
	return nil
}

func (m *DeviceSecret) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*SecretEntryEnvelope)(nil), "SecretEntryEnvelope")
	proto.RegisterType((*SecretEntryPayload)(nil), "SecretEntryPayload")
	proto.RegisterType((*DeviceSecret)(nil), "DeviceSecret")
}

func init() { proto.RegisterFile("go-internal/log_entry_secret.proto", fileDescriptor_5e1711af9e257970) }

var fileDescriptor_5e1711af9e257970 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0x8a, 0x42, 0x18, 0x38, 0x32, 0xa6, 0x3d, 0x88, 0x48, 0x4f, 0x7a, 0xd8, 0x86,
	0x08, 0xe2, 0x59, 0xdc, 0x49, 0x84, 0xb1, 0xde, 0xbc, 0x84, 0xb4, 0x79, 0xd4, 0x62, 0x97, 0x84,
	0x97, 0xd7, 0x41, 0x3e, 0x82, 0x5f, 0xc6, 0xcf, 0x28, 0x4d, 0xba, 0xae, 0x1e, 0xdf, 0xfb, 0xe5,
	0x97, 0xfc, 0xc3, 0x9f, 0x65, 0x95, 0x59, 0xd4, 0x9a, 0x00, 0xb5, 0x6c, 0x56, 0x8d, 0xa9, 0x04,
	0x68, 0x42, 0x2f, 0x1c, 0x94, 0x08, 0xb4, 0xb4, 0x68, 0xc8, 0x64, 0x3f, 0x09, 0x9b, 0xe5, 0x61,
	0xb1, 0xee, 0xe0, 0x5a, 0xef, 0xa1, 0x31, 0x16, 0xf8, 0x0b, 0x4b, 0x41, 0x97, 0xe8, 0x2d, 0x81,
	0xea, 0x0d, 0x61, 0xa5, 0x6f, 0x8c, 0x54, 0x69, 0x72, 0x97, 0xdc, 0x4f, 0xb6, 0x57, 0x03, 0x8f,
	0xfe, 0x26, 0xd2, 0xce, 0xfc, 0x7f, 0x5e, 0xb8, 0xba, 0xd2, 0x92, 0x5a, 0x84, 0xf4, 0x24, 0x9a,
	0x6e, 0x2c, 0xe4, 0x07, 0x9a, 0xfd, 0x26, 0x8c, 0x8f, 0xb2, 0x1c, 0x2e, 0x5c, 0xb0, 0x99, 0x02,
	0x47, 0x62, 0x07, 0xbb, 0x02, 0x50, 0xd8, 0xb6, 0x10, 0xdf, 0xe0, 0xfb, 0x14, 0xd3, 0x0e, 0x7d,
	0x04, 0xb2, 0x69, 0x8b, 0x77, 0xf0, 0xfc, 0x91, 0xcd, 0x1d, 0x68, 0x05, 0x28, 0x14, 0xec, 0xeb,
	0x12, 0x06, 0x21, 0x3e, 0xce, 0x23, 0x7c, 0x0b, 0xac, 0x57, 0x9e, 0xd9, 0xf5, 0xf1, 0xb3, 0xbd,
	0x15, 0x33, 0xa6, 0xa7, 0x41, 0x9a, 0x0f, 0x38, 0x7a, 0x31, 0x65, 0x96, 0xb3, 0xc9, 0x78, 0xe6,
	0x0f, 0x6c, 0xaa, 0x00, 0xeb, 0xbd, 0xa4, 0xda, 0x68, 0xe1, 0x48, 0x12, 0xf4, 0x31, 0x2f, 0x8f,
	0xfb, 0xbc, 0x5b, 0xf3, 0x94, 0x5d, 0x94, 0xa6, 0xed, 0xca, 0x09, 0xb9, 0xce, 0xb6, 0x87, 0xf1,
	0xf5, 0xf6, 0xf3, 0xa6, 0x00, 0x24, 0xbf, 0x24, 0x28, 0xbf, 0x56, 0x95, 0x59, 0x0d, 0x15, 0x56,
	0x68, 0x5a, 0x5b, 0x9c, 0x87, 0xe2, 0x9e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x35, 0xc1,
	0x02, 0xde, 0x01, 0x00, 0x00,
}
