// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: cart.proto

package pb

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

type CartNil struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartNil) Reset() {
	*x = CartNil{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartNil) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartNil) ProtoMessage() {}

func (x *CartNil) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartNil.ProtoReflect.Descriptor instead.
func (*CartNil) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{0}
}

type CartInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ProductID   string `protobuf:"bytes,2,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Quantity    int64  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Checked     int64  `protobuf:"varint,4,opt,name=Checked,proto3" json:"Checked,omitempty"`
	CartID      string `protobuf:"bytes,5,opt,name=CartID,proto3" json:"CartID,omitempty"`
	ProductName string `protobuf:"bytes,6,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
}

func (x *CartInfo) Reset() {
	*x = CartInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartInfo) ProtoMessage() {}

func (x *CartInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartInfo.ProtoReflect.Descriptor instead.
func (*CartInfo) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CartInfo) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CartInfo) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartInfo) GetChecked() int64 {
	if x != nil {
		return x.Checked
	}
	return 0
}

func (x *CartInfo) GetCartID() string {
	if x != nil {
		return x.CartID
	}
	return ""
}

func (x *CartInfo) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type CartCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID      string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ProductID   string `protobuf:"bytes,2,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Quantity    int64  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Checked     int64  `protobuf:"varint,4,opt,name=Checked,proto3" json:"Checked,omitempty"`
	ProductName string `protobuf:"bytes,5,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
}

func (x *CartCreateReq) Reset() {
	*x = CartCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartCreateReq) ProtoMessage() {}

func (x *CartCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartCreateReq.ProtoReflect.Descriptor instead.
func (*CartCreateReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CartCreateReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CartCreateReq) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *CartCreateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartCreateReq) GetChecked() int64 {
	if x != nil {
		return x.Checked
	}
	return 0
}

func (x *CartCreateReq) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

type CartListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	KeyWord  string `protobuf:"bytes,3,opt,name=KeyWord,proto3" json:"KeyWord,omitempty"`
	UserID   string `protobuf:"bytes,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CartListReq) Reset() {
	*x = CartListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartListReq) ProtoMessage() {}

func (x *CartListReq) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartListReq.ProtoReflect.Descriptor instead.
func (*CartListReq) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{3}
}

func (x *CartListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CartListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CartListReq) GetKeyWord() string {
	if x != nil {
		return x.KeyWord
	}
	return ""
}

func (x *CartListReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type CartListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64       `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	List  []*CartInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *CartListReply) Reset() {
	*x = CartListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartListReply) ProtoMessage() {}

func (x *CartListReply) ProtoReflect() protoreflect.Message {
	mi := &file_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartListReply.ProtoReflect.Descriptor instead.
func (*CartListReply) Descriptor() ([]byte, []int) {
	return file_cart_proto_rawDescGZIP(), []int{4}
}

func (x *CartListReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CartListReply) GetList() []*CartInfo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_cart_proto protoreflect.FileDescriptor

var file_cart_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x09, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x74, 0x4e, 0x69, 0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x08,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x61, 0x72, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x9d,
	0x01, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x6f,
	0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4b, 0x65, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x47, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x64, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74,
	0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x4e, 0x69, 0x6c, 0x12, 0x2e,
	0x0a, 0x08, 0x63, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_proto_rawDescOnce sync.Once
	file_cart_proto_rawDescData = file_cart_proto_rawDesc
)

func file_cart_proto_rawDescGZIP() []byte {
	file_cart_proto_rawDescOnce.Do(func() {
		file_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_proto_rawDescData)
	})
	return file_cart_proto_rawDescData
}

var file_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cart_proto_goTypes = []interface{}{
	(*CartNil)(nil),       // 0: pb.CartNil
	(*CartInfo)(nil),      // 1: pb.CartInfo
	(*CartCreateReq)(nil), // 2: pb.CartCreateReq
	(*CartListReq)(nil),   // 3: pb.CartListReq
	(*CartListReply)(nil), // 4: pb.CartListReply
}
var file_cart_proto_depIdxs = []int32{
	1, // 0: pb.CartListReply.List:type_name -> pb.CartInfo
	2, // 1: pb.cart.cartCreate:input_type -> pb.CartCreateReq
	3, // 2: pb.cart.cartList:input_type -> pb.CartListReq
	0, // 3: pb.cart.cartCreate:output_type -> pb.CartNil
	4, // 4: pb.cart.cartList:output_type -> pb.CartListReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cart_proto_init() }
func file_cart_proto_init() {
	if File_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartNil); i {
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
		file_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartInfo); i {
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
		file_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartCreateReq); i {
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
		file_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartListReq); i {
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
		file_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartListReply); i {
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
			RawDescriptor: file_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_proto_goTypes,
		DependencyIndexes: file_cart_proto_depIdxs,
		MessageInfos:      file_cart_proto_msgTypes,
	}.Build()
	File_cart_proto = out.File
	file_cart_proto_rawDesc = nil
	file_cart_proto_goTypes = nil
	file_cart_proto_depIdxs = nil
}
