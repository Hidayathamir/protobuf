// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: goout/erajol_bike.proto

package pbgoout

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

type ReqErajolBikeOrderDriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId uint64 `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	DriverId   uint64 `protobuf:"varint,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Price      int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ReqErajolBikeOrderDriver) Reset() {
	*x = ReqErajolBikeOrderDriver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goout_erajol_bike_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqErajolBikeOrderDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqErajolBikeOrderDriver) ProtoMessage() {}

func (x *ReqErajolBikeOrderDriver) ProtoReflect() protoreflect.Message {
	mi := &file_goout_erajol_bike_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqErajolBikeOrderDriver.ProtoReflect.Descriptor instead.
func (*ReqErajolBikeOrderDriver) Descriptor() ([]byte, []int) {
	return file_goout_erajol_bike_proto_rawDescGZIP(), []int{0}
}

func (x *ReqErajolBikeOrderDriver) GetCustomerId() uint64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *ReqErajolBikeOrderDriver) GetDriverId() uint64 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *ReqErajolBikeOrderDriver) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ResErajolBikeOrderDriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *ResErajolBikeOrderDriver) Reset() {
	*x = ResErajolBikeOrderDriver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goout_erajol_bike_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResErajolBikeOrderDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResErajolBikeOrderDriver) ProtoMessage() {}

func (x *ResErajolBikeOrderDriver) ProtoReflect() protoreflect.Message {
	mi := &file_goout_erajol_bike_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResErajolBikeOrderDriver.ProtoReflect.Descriptor instead.
func (*ResErajolBikeOrderDriver) Descriptor() ([]byte, []int) {
	return file_goout_erajol_bike_proto_rawDescGZIP(), []int{1}
}

func (x *ResErajolBikeOrderDriver) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_goout_erajol_bike_proto protoreflect.FileDescriptor

var file_goout_erajol_bike_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x6f, 0x75, 0x74, 0x2f, 0x65, 0x72, 0x61, 0x6a, 0x6f, 0x6c, 0x5f, 0x62,
	0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x18, 0x52, 0x65, 0x71,
	0x45, 0x72, 0x61, 0x6a, 0x6f, 0x6c, 0x42, 0x69, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x35, 0x0a, 0x18, 0x52, 0x65, 0x73,
	0x45, 0x72, 0x61, 0x6a, 0x6f, 0x6c, 0x42, 0x69, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x32, 0x53, 0x0a, 0x0a, 0x45, 0x72, 0x61, 0x6a, 0x6f, 0x6c, 0x42, 0x69, 0x6b, 0x65, 0x12, 0x45,
	0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x52, 0x65, 0x71, 0x45, 0x72, 0x61, 0x6a, 0x6f, 0x6c, 0x42, 0x69, 0x6b, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x52, 0x65, 0x73, 0x45, 0x72,
	0x61, 0x6a, 0x6f, 0x6c, 0x42, 0x69, 0x6b, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x69, 0x64, 0x61, 0x79, 0x61, 0x74, 0x68, 0x61, 0x6d, 0x69, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x6f, 0x75,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goout_erajol_bike_proto_rawDescOnce sync.Once
	file_goout_erajol_bike_proto_rawDescData = file_goout_erajol_bike_proto_rawDesc
)

func file_goout_erajol_bike_proto_rawDescGZIP() []byte {
	file_goout_erajol_bike_proto_rawDescOnce.Do(func() {
		file_goout_erajol_bike_proto_rawDescData = protoimpl.X.CompressGZIP(file_goout_erajol_bike_proto_rawDescData)
	})
	return file_goout_erajol_bike_proto_rawDescData
}

var file_goout_erajol_bike_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_goout_erajol_bike_proto_goTypes = []interface{}{
	(*ReqErajolBikeOrderDriver)(nil), // 0: ReqErajolBikeOrderDriver
	(*ResErajolBikeOrderDriver)(nil), // 1: ResErajolBikeOrderDriver
}
var file_goout_erajol_bike_proto_depIdxs = []int32{
	0, // 0: ErajolBike.OrderDriver:input_type -> ReqErajolBikeOrderDriver
	1, // 1: ErajolBike.OrderDriver:output_type -> ResErajolBikeOrderDriver
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goout_erajol_bike_proto_init() }
func file_goout_erajol_bike_proto_init() {
	if File_goout_erajol_bike_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goout_erajol_bike_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqErajolBikeOrderDriver); i {
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
		file_goout_erajol_bike_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResErajolBikeOrderDriver); i {
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
			RawDescriptor: file_goout_erajol_bike_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goout_erajol_bike_proto_goTypes,
		DependencyIndexes: file_goout_erajol_bike_proto_depIdxs,
		MessageInfos:      file_goout_erajol_bike_proto_msgTypes,
	}.Build()
	File_goout_erajol_bike_proto = out.File
	file_goout_erajol_bike_proto_rawDesc = nil
	file_goout_erajol_bike_proto_goTypes = nil
	file_goout_erajol_bike_proto_depIdxs = nil
}
