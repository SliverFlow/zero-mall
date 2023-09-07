// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: system.proto

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

type NilReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NilReply) Reset() {
	*x = NilReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NilReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NilReply) ProtoMessage() {}

func (x *NilReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NilReply.ProtoReflect.Descriptor instead.
func (*NilReply) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

type NilReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NilReq) Reset() {
	*x = NilReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NilReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NilReq) ProtoMessage() {}

func (x *NilReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NilReq.ProtoReflect.Descriptor instead.
func (*NilReq) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{1}
}

// 登录入参
type SystemLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SystemLoginReq) Reset() {
	*x = SystemLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemLoginReq) ProtoMessage() {}

func (x *SystemLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemLoginReq.ProtoReflect.Descriptor instead.
func (*SystemLoginReq) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{2}
}

func (x *SystemLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SystemLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// 登录出参
type SystemLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SystemLoginReply) Reset() {
	*x = SystemLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemLoginReply) ProtoMessage() {}

func (x *SystemLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemLoginReply.ProtoReflect.Descriptor instead.
func (*SystemLoginReply) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{3}
}

func (x *SystemLoginReply) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// sys_user
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UUID      string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Nickname  string `protobuf:"bytes,5,opt,name=Nickname,proto3" json:"Nickname,omitempty"`
	Avatar    string `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Type      int64  `protobuf:"varint,7,opt,name=Type,proto3" json:"Type,omitempty"`
	Status    int64  `protobuf:"varint,8,opt,name=Status,proto3" json:"Status,omitempty"`
	ID        uint64 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt int64  `protobuf:"varint,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *User) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *User) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *User) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *CreateRole) Reset() {
	*x = CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRole) ProtoMessage() {}

func (x *CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRole.ProtoReflect.Descriptor instead.
func (*CreateRole) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ParentId  int64  `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
	Component string `protobuf:"bytes,5,opt,name=Component,proto3" json:"Component,omitempty"`
	Sorted    int64  `protobuf:"varint,6,opt,name=Sorted,proto3" json:"Sorted,omitempty"`
	Title     string `protobuf:"bytes,7,opt,name=Title,proto3" json:"Title,omitempty"`
	Icon      string `protobuf:"bytes,8,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Status    int64  `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
	Role      int64  `protobuf:"varint,10,opt,name=Role,proto3" json:"Role,omitempty"`
	CreatedAt int64  `protobuf:"varint,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,12,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{6}
}

func (x *Menu) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Menu) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Menu) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Menu) GetSorted() int64 {
	if x != nil {
		return x.Sorted
	}
	return 0
}

func (x *Menu) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Menu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Menu) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Menu) GetRole() int64 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *Menu) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Menu) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type RoleIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RoleIDReq) Reset() {
	*x = RoleIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleIDReq) ProtoMessage() {}

func (x *RoleIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleIDReq.ProtoReflect.Descriptor instead.
func (*RoleIDReq) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{7}
}

func (x *RoleIDReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type MenuListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Menu `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *MenuListReply) Reset() {
	*x = MenuListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuListReply) ProtoMessage() {}

func (x *MenuListReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuListReply.ProtoReflect.Descriptor instead.
func (*MenuListReply) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{8}
}

func (x *MenuListReply) GetList() []*Menu {
	if x != nil {
		return x.List
	}
	return nil
}

type MenuChangeStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PID    int64 `protobuf:"varint,2,opt,name=PID,proto3" json:"PID,omitempty"`
	Status int64 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *MenuChangeStatusReq) Reset() {
	*x = MenuChangeStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuChangeStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuChangeStatusReq) ProtoMessage() {}

func (x *MenuChangeStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuChangeStatusReq.ProtoReflect.Descriptor instead.
func (*MenuChangeStatusReq) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{9}
}

func (x *MenuChangeStatusReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MenuChangeStatusReq) GetPID() int64 {
	if x != nil {
		return x.PID
	}
	return 0
}

func (x *MenuChangeStatusReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type MenuCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId  int64  `protobuf:"varint,2,opt,name=ParentId,proto3" json:"ParentId,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Path      string `protobuf:"bytes,4,opt,name=Path,proto3" json:"Path,omitempty"`
	Component string `protobuf:"bytes,5,opt,name=Component,proto3" json:"Component,omitempty"`
	Sorted    int64  `protobuf:"varint,6,opt,name=Sorted,proto3" json:"Sorted,omitempty"`
	Title     string `protobuf:"bytes,7,opt,name=Title,proto3" json:"Title,omitempty"`
	Icon      string `protobuf:"bytes,8,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Status    int64  `protobuf:"varint,9,opt,name=Status,proto3" json:"Status,omitempty"`
	Role      int64  `protobuf:"varint,10,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *MenuCreateReq) Reset() {
	*x = MenuCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuCreateReq) ProtoMessage() {}

func (x *MenuCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuCreateReq.ProtoReflect.Descriptor instead.
func (*MenuCreateReq) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{10}
}

func (x *MenuCreateReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuCreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuCreateReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuCreateReq) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *MenuCreateReq) GetSorted() int64 {
	if x != nil {
		return x.Sorted
	}
	return 0
}

func (x *MenuCreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuCreateReq) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *MenuCreateReq) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MenuCreateReq) GetRole() int64 {
	if x != nil {
		return x.Role
	}
	return 0
}

var File_system_proto protoreflect.FileDescriptor

var file_system_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x0a, 0x0a, 0x08, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x08,
	0x0a, 0x06, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x90, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x20, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa2, 0x02, 0x0a, 0x04, 0x4d, 0x65,
	0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1b,
	0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x2d, 0x0a, 0x0d, 0x4d,
	0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x4d, 0x65,
	0x6e, 0x75, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x50, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0d,
	0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x32, 0xbc, 0x02,
	0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x0a, 0x52,
	0x6f, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e,
	0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x0e, 0x4d, 0x65, 0x6e, 0x75, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x11, 0x4d,
	0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x39, 0x0a, 0x10, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x6e, 0x75,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a,
	0x0a, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_proto_rawDescOnce sync.Once
	file_system_proto_rawDescData = file_system_proto_rawDesc
)

func file_system_proto_rawDescGZIP() []byte {
	file_system_proto_rawDescOnce.Do(func() {
		file_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_proto_rawDescData)
	})
	return file_system_proto_rawDescData
}

var file_system_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_system_proto_goTypes = []interface{}{
	(*NilReply)(nil),            // 0: pb.NilReply
	(*NilReq)(nil),              // 1: pb.NilReq
	(*SystemLoginReq)(nil),      // 2: pb.SystemLoginReq
	(*SystemLoginReply)(nil),    // 3: pb.SystemLoginReply
	(*User)(nil),                // 4: pb.User
	(*CreateRole)(nil),          // 5: pb.CreateRole
	(*Menu)(nil),                // 6: pb.Menu
	(*RoleIDReq)(nil),           // 7: pb.RoleIDReq
	(*MenuListReply)(nil),       // 8: pb.MenuListReply
	(*MenuChangeStatusReq)(nil), // 9: pb.MenuChangeStatusReq
	(*MenuCreateReq)(nil),       // 10: pb.MenuCreateReq
}
var file_system_proto_depIdxs = []int32{
	4,  // 0: pb.SystemLoginReply.user:type_name -> pb.User
	6,  // 1: pb.MenuListReply.List:type_name -> pb.Menu
	2,  // 2: pb.system.Login:input_type -> pb.SystemLoginReq
	5,  // 3: pb.system.RoleCreate:input_type -> pb.CreateRole
	7,  // 4: pb.system.MenuListByRole:input_type -> pb.RoleIDReq
	7,  // 5: pb.system.MenuListAllByRole:input_type -> pb.RoleIDReq
	9,  // 6: pb.system.MenuChangeStatus:input_type -> pb.MenuChangeStatusReq
	10, // 7: pb.system.MenuCreate:input_type -> pb.MenuCreateReq
	3,  // 8: pb.system.Login:output_type -> pb.SystemLoginReply
	0,  // 9: pb.system.RoleCreate:output_type -> pb.NilReply
	8,  // 10: pb.system.MenuListByRole:output_type -> pb.MenuListReply
	8,  // 11: pb.system.MenuListAllByRole:output_type -> pb.MenuListReply
	0,  // 12: pb.system.MenuChangeStatus:output_type -> pb.NilReply
	0,  // 13: pb.system.MenuCreate:output_type -> pb.NilReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_system_proto_init() }
func file_system_proto_init() {
	if File_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NilReply); i {
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
		file_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NilReq); i {
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
		file_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemLoginReq); i {
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
		file_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemLoginReply); i {
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
		file_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRole); i {
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
		file_system_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_system_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleIDReq); i {
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
		file_system_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuListReply); i {
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
		file_system_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuChangeStatusReq); i {
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
		file_system_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuCreateReq); i {
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
			RawDescriptor: file_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_proto_goTypes,
		DependencyIndexes: file_system_proto_depIdxs,
		MessageInfos:      file_system_proto_msgTypes,
	}.Build()
	File_system_proto = out.File
	file_system_proto_rawDesc = nil
	file_system_proto_goTypes = nil
	file_system_proto_depIdxs = nil
}
