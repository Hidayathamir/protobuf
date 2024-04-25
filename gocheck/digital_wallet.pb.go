// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: gocheck/digital_wallet.proto

package gocheck

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

type ReqDigitalWalletTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientId uint64 `protobuf:"varint,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ReqDigitalWalletTransfer) Reset() {
	*x = ReqDigitalWalletTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocheck_digital_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqDigitalWalletTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqDigitalWalletTransfer) ProtoMessage() {}

func (x *ReqDigitalWalletTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_gocheck_digital_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqDigitalWalletTransfer.ProtoReflect.Descriptor instead.
func (*ReqDigitalWalletTransfer) Descriptor() ([]byte, []int) {
	return file_gocheck_digital_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *ReqDigitalWalletTransfer) GetRecipientId() uint64 {
	if x != nil {
		return x.RecipientId
	}
	return 0
}

func (x *ReqDigitalWalletTransfer) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ResDigitalWalletTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResDigitalWalletTransfer) Reset() {
	*x = ResDigitalWalletTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gocheck_digital_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResDigitalWalletTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResDigitalWalletTransfer) ProtoMessage() {}

func (x *ResDigitalWalletTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_gocheck_digital_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResDigitalWalletTransfer.ProtoReflect.Descriptor instead.
func (*ResDigitalWalletTransfer) Descriptor() ([]byte, []int) {
	return file_gocheck_digital_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *ResDigitalWalletTransfer) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_gocheck_digital_wallet_proto protoreflect.FileDescriptor

var file_gocheck_digital_wallet_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x6f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x18, 0x52, 0x65, 0x71, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x18, 0x52, 0x65, 0x73, 0x44, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x32, 0x53, 0x0a, 0x0d, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x19,
	0x2e, 0x52, 0x65, 0x71, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x52, 0x65, 0x73, 0x44,
	0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gocheck_digital_wallet_proto_rawDescOnce sync.Once
	file_gocheck_digital_wallet_proto_rawDescData = file_gocheck_digital_wallet_proto_rawDesc
)

func file_gocheck_digital_wallet_proto_rawDescGZIP() []byte {
	file_gocheck_digital_wallet_proto_rawDescOnce.Do(func() {
		file_gocheck_digital_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_gocheck_digital_wallet_proto_rawDescData)
	})
	return file_gocheck_digital_wallet_proto_rawDescData
}

var file_gocheck_digital_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gocheck_digital_wallet_proto_goTypes = []interface{}{
	(*ReqDigitalWalletTransfer)(nil), // 0: ReqDigitalWalletTransfer
	(*ResDigitalWalletTransfer)(nil), // 1: ResDigitalWalletTransfer
}
var file_gocheck_digital_wallet_proto_depIdxs = []int32{
	0, // 0: DigitalWallet.Transfer:input_type -> ReqDigitalWalletTransfer
	1, // 1: DigitalWallet.Transfer:output_type -> ResDigitalWalletTransfer
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gocheck_digital_wallet_proto_init() }
func file_gocheck_digital_wallet_proto_init() {
	if File_gocheck_digital_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gocheck_digital_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqDigitalWalletTransfer); i {
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
		file_gocheck_digital_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResDigitalWalletTransfer); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gocheck_digital_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gocheck_digital_wallet_proto_goTypes,
		DependencyIndexes: file_gocheck_digital_wallet_proto_depIdxs,
		MessageInfos:      file_gocheck_digital_wallet_proto_msgTypes,
	}.Build()
	File_gocheck_digital_wallet_proto = out.File
	file_gocheck_digital_wallet_proto_rawDesc = nil
	file_gocheck_digital_wallet_proto_goTypes = nil
	file_gocheck_digital_wallet_proto_depIdxs = nil
}
