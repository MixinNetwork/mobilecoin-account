// Copyright (c) 2018-2022 The MobileCoin Foundation

/// Protos to be used for displaying encoded strings to users

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: printable.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

/// Message for a payment request, which combines a public address
/// with an a requested payment amount and memo field
type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// The public address of the user requesting a payment
	PublicAddress *PublicAddress `protobuf:"bytes,1,opt,name=public_address,json=publicAddress,proto3" json:"public_address,omitempty"`
	/// The requested value of the payment
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	/// Any additional text explaining the request
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
	/// Token id to transact in.
	TokenId uint64 `protobuf:"varint,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_printable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_printable_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetPublicAddress() *PublicAddress {
	if x != nil {
		return x.PublicAddress
	}
	return nil
}

func (x *PaymentRequest) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *PaymentRequest) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *PaymentRequest) GetTokenId() uint64 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

/// Message encoding a private key and a UTXO, for the purpose of
/// giving someone access to an output. This would most likely be
/// used for gift cards.
type TransferPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// [Deprecated] The root entropy, allowing the recipient to spend the money.
	/// This has been replaced by a BIP39 entropy.
	//
	// Deprecated: Do not use.
	RootEntropy []byte `protobuf:"bytes,1,opt,name=root_entropy,json=rootEntropy,proto3" json:"root_entropy,omitempty"`
	/// The public key of the UTXO to spend. This is an optimization, meaning
	/// the recipient does not need to scan the entire ledger.
	TxOutPublicKey *CompressedRistretto `protobuf:"bytes,2,opt,name=tx_out_public_key,json=txOutPublicKey,proto3" json:"tx_out_public_key,omitempty"`
	/// Any additional text explaining the gift
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
	/// BIP39 entropy, allowing the recipient to spend the money.
	/// When deriving an AccountKey from this entropy, account_index is always 0.
	Bip39Entropy []byte `protobuf:"bytes,4,opt,name=bip39_entropy,json=bip39Entropy,proto3" json:"bip39_entropy,omitempty"`
}

func (x *TransferPayload) Reset() {
	*x = TransferPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferPayload) ProtoMessage() {}

func (x *TransferPayload) ProtoReflect() protoreflect.Message {
	mi := &file_printable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferPayload.ProtoReflect.Descriptor instead.
func (*TransferPayload) Descriptor() ([]byte, []int) {
	return file_printable_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *TransferPayload) GetRootEntropy() []byte {
	if x != nil {
		return x.RootEntropy
	}
	return nil
}

func (x *TransferPayload) GetTxOutPublicKey() *CompressedRistretto {
	if x != nil {
		return x.TxOutPublicKey
	}
	return nil
}

func (x *TransferPayload) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TransferPayload) GetBip39Entropy() []byte {
	if x != nil {
		return x.Bip39Entropy
	}
	return nil
}

/// Message encoding information required to locate a TxOut,
/// un-blind the amount, and spend the TxOut. This can be used to give
/// MobileCoin to both FOG & non-FOG users who may not yet have
// a MobileCoin account enabled
type TxOutGiftCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The global index of the TxOut that has been gifted. This allows
	// the receiver to find & uniquely identify the TxOut
	GlobalIndex uint64 `protobuf:"varint,1,opt,name=global_index,json=globalIndex,proto3" json:"global_index,omitempty"`
	// The one-time private key which can be used to spend the TxOut
	OnetimePrivateKey *RistrettoPrivate `protobuf:"bytes,2,opt,name=onetime_private_key,json=onetimePrivateKey,proto3" json:"onetime_private_key,omitempty"`
	// The shared secret used to un-blind the amount of the TxOut
	SharedSecret *CompressedRistretto `protobuf:"bytes,3,opt,name=shared_secret,json=sharedSecret,proto3" json:"shared_secret,omitempty"`
}

func (x *TxOutGiftCode) Reset() {
	*x = TxOutGiftCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxOutGiftCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxOutGiftCode) ProtoMessage() {}

func (x *TxOutGiftCode) ProtoReflect() protoreflect.Message {
	mi := &file_printable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxOutGiftCode.ProtoReflect.Descriptor instead.
func (*TxOutGiftCode) Descriptor() ([]byte, []int) {
	return file_printable_proto_rawDescGZIP(), []int{2}
}

func (x *TxOutGiftCode) GetGlobalIndex() uint64 {
	if x != nil {
		return x.GlobalIndex
	}
	return 0
}

func (x *TxOutGiftCode) GetOnetimePrivateKey() *RistrettoPrivate {
	if x != nil {
		return x.OnetimePrivateKey
	}
	return nil
}

func (x *TxOutGiftCode) GetSharedSecret() *CompressedRistretto {
	if x != nil {
		return x.SharedSecret
	}
	return nil
}

/// This wraps all of the above messages using "oneof", allowing us to
/// have a single encoding scheme and extend as necessary simply by adding
/// new messages without breaking backwards compatibility
type PrintableWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Wrapper:
	//	*PrintableWrapper_PublicAddress
	//	*PrintableWrapper_PaymentRequest
	//	*PrintableWrapper_TransferPayload
	//	*PrintableWrapper_TxOutGiftCode
	Wrapper isPrintableWrapper_Wrapper `protobuf_oneof:"wrapper"`
}

func (x *PrintableWrapper) Reset() {
	*x = PrintableWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_printable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrintableWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrintableWrapper) ProtoMessage() {}

func (x *PrintableWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_printable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrintableWrapper.ProtoReflect.Descriptor instead.
func (*PrintableWrapper) Descriptor() ([]byte, []int) {
	return file_printable_proto_rawDescGZIP(), []int{3}
}

func (m *PrintableWrapper) GetWrapper() isPrintableWrapper_Wrapper {
	if m != nil {
		return m.Wrapper
	}
	return nil
}

func (x *PrintableWrapper) GetPublicAddress() *PublicAddress {
	if x, ok := x.GetWrapper().(*PrintableWrapper_PublicAddress); ok {
		return x.PublicAddress
	}
	return nil
}

func (x *PrintableWrapper) GetPaymentRequest() *PaymentRequest {
	if x, ok := x.GetWrapper().(*PrintableWrapper_PaymentRequest); ok {
		return x.PaymentRequest
	}
	return nil
}

func (x *PrintableWrapper) GetTransferPayload() *TransferPayload {
	if x, ok := x.GetWrapper().(*PrintableWrapper_TransferPayload); ok {
		return x.TransferPayload
	}
	return nil
}

func (x *PrintableWrapper) GetTxOutGiftCode() *TxOutGiftCode {
	if x, ok := x.GetWrapper().(*PrintableWrapper_TxOutGiftCode); ok {
		return x.TxOutGiftCode
	}
	return nil
}

type isPrintableWrapper_Wrapper interface {
	isPrintableWrapper_Wrapper()
}

type PrintableWrapper_PublicAddress struct {
	PublicAddress *PublicAddress `protobuf:"bytes,1,opt,name=public_address,json=publicAddress,proto3,oneof"`
}

type PrintableWrapper_PaymentRequest struct {
	PaymentRequest *PaymentRequest `protobuf:"bytes,2,opt,name=payment_request,json=paymentRequest,proto3,oneof"`
}

type PrintableWrapper_TransferPayload struct {
	TransferPayload *TransferPayload `protobuf:"bytes,3,opt,name=transfer_payload,json=transferPayload,proto3,oneof"`
}

type PrintableWrapper_TxOutGiftCode struct {
	TxOutGiftCode *TxOutGiftCode `protobuf:"bytes,4,opt,name=tx_out_gift_code,json=txOutGiftCode,proto3,oneof"`
}

func (*PrintableWrapper_PublicAddress) isPrintableWrapper_Wrapper() {}

func (*PrintableWrapper_PaymentRequest) isPrintableWrapper_Wrapper() {}

func (*PrintableWrapper_TransferPayload) isPrintableWrapper_Wrapper() {}

func (*PrintableWrapper_TxOutGiftCode) isPrintableWrapper_Wrapper() {}

var File_printable_proto protoreflect.FileDescriptor

var file_printable_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x0e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3e, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x25, 0x0a, 0x0c, 0x72, 0x6f, 0x6f, 0x74,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x12,
	0x48, 0x0a, 0x11, 0x74, 0x78, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x52, 0x69, 0x73, 0x74, 0x72, 0x65, 0x74, 0x74, 0x6f, 0x52, 0x0e, 0x74, 0x78, 0x4f, 0x75, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x23, 0x0a,
	0x0d, 0x62, 0x69, 0x70, 0x33, 0x39, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x62, 0x69, 0x70, 0x33, 0x39, 0x45, 0x6e, 0x74, 0x72, 0x6f,
	0x70, 0x79, 0x22, 0xc2, 0x01, 0x0a, 0x0d, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x47, 0x69, 0x66, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x4a, 0x0a, 0x13, 0x6f, 0x6e, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x52, 0x69, 0x73, 0x74, 0x72, 0x65, 0x74, 0x74, 0x6f, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x11, 0x6f, 0x6e, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x52, 0x69, 0x73, 0x74, 0x72, 0x65, 0x74, 0x74, 0x6f, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x69, 0x6e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x44,
	0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x43, 0x0a,
	0x10, 0x74, 0x78, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x54, 0x78, 0x4f, 0x75, 0x74, 0x47, 0x69, 0x66, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x78, 0x4f, 0x75, 0x74, 0x47, 0x69, 0x66, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x1c, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_printable_proto_rawDescOnce sync.Once
	file_printable_proto_rawDescData = file_printable_proto_rawDesc
)

func file_printable_proto_rawDescGZIP() []byte {
	file_printable_proto_rawDescOnce.Do(func() {
		file_printable_proto_rawDescData = protoimpl.X.CompressGZIP(file_printable_proto_rawDescData)
	})
	return file_printable_proto_rawDescData
}

var file_printable_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_printable_proto_goTypes = []interface{}{
	(*PaymentRequest)(nil),      // 0: printable.PaymentRequest
	(*TransferPayload)(nil),     // 1: printable.TransferPayload
	(*TxOutGiftCode)(nil),       // 2: printable.TxOutGiftCode
	(*PrintableWrapper)(nil),    // 3: printable.PrintableWrapper
	(*PublicAddress)(nil),       // 4: external.PublicAddress
	(*CompressedRistretto)(nil), // 5: external.CompressedRistretto
	(*RistrettoPrivate)(nil),    // 6: external.RistrettoPrivate
}
var file_printable_proto_depIdxs = []int32{
	4, // 0: printable.PaymentRequest.public_address:type_name -> external.PublicAddress
	5, // 1: printable.TransferPayload.tx_out_public_key:type_name -> external.CompressedRistretto
	6, // 2: printable.TxOutGiftCode.onetime_private_key:type_name -> external.RistrettoPrivate
	5, // 3: printable.TxOutGiftCode.shared_secret:type_name -> external.CompressedRistretto
	4, // 4: printable.PrintableWrapper.public_address:type_name -> external.PublicAddress
	0, // 5: printable.PrintableWrapper.payment_request:type_name -> printable.PaymentRequest
	1, // 6: printable.PrintableWrapper.transfer_payload:type_name -> printable.TransferPayload
	2, // 7: printable.PrintableWrapper.tx_out_gift_code:type_name -> printable.TxOutGiftCode
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_printable_proto_init() }
func file_printable_proto_init() {
	if File_printable_proto != nil {
		return
	}
	file_external_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_printable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_printable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_printable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxOutGiftCode); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_printable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrintableWrapper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_printable_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*PrintableWrapper_PublicAddress)(nil),
		(*PrintableWrapper_PaymentRequest)(nil),
		(*PrintableWrapper_TransferPayload)(nil),
		(*PrintableWrapper_TxOutGiftCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_printable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_printable_proto_goTypes,
		DependencyIndexes: file_printable_proto_depIdxs,
		MessageInfos:      file_printable_proto_msgTypes,
	}.Build()
	File_printable_proto = out.File
	file_printable_proto_rawDesc = nil
	file_printable_proto_goTypes = nil
	file_printable_proto_depIdxs = nil
}
