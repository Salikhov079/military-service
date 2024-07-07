// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: soldiers.proto

package soldiers

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

type UseB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuantityWeapon    int32  `protobuf:"varint,1,opt,name=quantity_weapon,json=quantityWeapon,proto3" json:"quantity_weapon,omitempty"`
	QuantityBigWeapon int32  `protobuf:"varint,2,opt,name=quantity_big_weapon,json=quantityBigWeapon,proto3" json:"quantity_big_weapon,omitempty"`
	SoldierId         string `protobuf:"bytes,3,opt,name=soldier_id,json=soldierId,proto3" json:"soldier_id,omitempty"`
	Date              string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *UseB) Reset() {
	*x = UseB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseB) ProtoMessage() {}

func (x *UseB) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseB.ProtoReflect.Descriptor instead.
func (*UseB) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{0}
}

func (x *UseB) GetQuantityWeapon() int32 {
	if x != nil {
		return x.QuantityWeapon
	}
	return 0
}

func (x *UseB) GetQuantityBigWeapon() int32 {
	if x != nil {
		return x.QuantityBigWeapon
	}
	return 0
}

func (x *UseB) GetSoldierId() string {
	if x != nil {
		return x.SoldierId
	}
	return ""
}

func (x *UseB) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type UseF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diesel    int32  `protobuf:"varint,1,opt,name=diesel,proto3" json:"diesel,omitempty"`
	Petrol    int32  `protobuf:"varint,2,opt,name=petrol,proto3" json:"petrol,omitempty"`
	SoldierId string `protobuf:"bytes,3,opt,name=soldier_id,json=soldierId,proto3" json:"soldier_id,omitempty"`
	Date      string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *UseF) Reset() {
	*x = UseF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseF) ProtoMessage() {}

func (x *UseF) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseF.ProtoReflect.Descriptor instead.
func (*UseF) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{1}
}

func (x *UseF) GetDiesel() int32 {
	if x != nil {
		return x.Diesel
	}
	return 0
}

func (x *UseF) GetPetrol() int32 {
	if x != nil {
		return x.Petrol
	}
	return 0
}

func (x *UseF) GetSoldierId() string {
	if x != nil {
		return x.SoldierId
	}
	return ""
}

func (x *UseF) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ById) Reset() {
	*x = ById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ById) ProtoMessage() {}

func (x *ById) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ById.ProtoReflect.Descriptor instead.
func (*ById) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{2}
}

func (x *ById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{3}
}

type SoldierReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	DateOfBirth string `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	GroupId     string `protobuf:"bytes,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	JoinDate    string `protobuf:"bytes,6,opt,name=join_date,json=joinDate,proto3" json:"join_date,omitempty"`
}

func (x *SoldierReq) Reset() {
	*x = SoldierReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoldierReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoldierReq) ProtoMessage() {}

func (x *SoldierReq) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoldierReq.ProtoReflect.Descriptor instead.
func (*SoldierReq) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{4}
}

func (x *SoldierReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SoldierReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SoldierReq) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *SoldierReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SoldierReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *SoldierReq) GetJoinDate() string {
	if x != nil {
		return x.JoinDate
	}
	return ""
}

type Soldier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	DateOfBirth string `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	GroupId     string `protobuf:"bytes,6,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	JoinDate    string `protobuf:"bytes,7,opt,name=join_date,json=joinDate,proto3" json:"join_date,omitempty"`
}

func (x *Soldier) Reset() {
	*x = Soldier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Soldier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Soldier) ProtoMessage() {}

func (x *Soldier) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Soldier.ProtoReflect.Descriptor instead.
func (*Soldier) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{5}
}

func (x *Soldier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Soldier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Soldier) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Soldier) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Soldier) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Soldier) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Soldier) GetJoinDate() string {
	if x != nil {
		return x.JoinDate
	}
	return ""
}

type AllSoldiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Soldiers []*Soldier `protobuf:"bytes,1,rep,name=soldiers,proto3" json:"soldiers,omitempty"`
}

func (x *AllSoldiers) Reset() {
	*x = AllSoldiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_soldiers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSoldiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSoldiers) ProtoMessage() {}

func (x *AllSoldiers) ProtoReflect() protoreflect.Message {
	mi := &file_soldiers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSoldiers.ProtoReflect.Descriptor instead.
func (*AllSoldiers) Descriptor() ([]byte, []int) {
	return file_soldiers_proto_rawDescGZIP(), []int{6}
}

func (x *AllSoldiers) GetSoldiers() []*Soldier {
	if x != nil {
		return x.Soldiers
	}
	return nil
}

var File_soldiers_proto protoreflect.FileDescriptor

var file_soldiers_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x42, 0x12, 0x27, 0x0a, 0x0f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x69, 0x67, 0x5f, 0x77, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x42, 0x69, 0x67, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x69, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x46, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x65, 0x73, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x69, 0x65, 0x73, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x65, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x70, 0x65, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x6c,
	0x64, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x16, 0x0a, 0x04, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x0a, 0x53,
	0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f,
	0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a,
	0x6f, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3c, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x53, 0x6f,
	0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x6f, 0x6c,
	0x64, 0x69, 0x65, 0x72, 0x73, 0x32, 0xd0, 0x02, 0x0a, 0x0e, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x6f,
	0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x6f,
	0x6c, 0x64, 0x69, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a,
	0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x28, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72,
	0x73, 0x2e, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x53,
	0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x73, 0x6f, 0x6c, 0x64,
	0x69, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73,
	0x12, 0x2b, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x42, 0x1a, 0x0e, 0x2e,
	0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x46, 0x75, 0x65, 0x6c, 0x12, 0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x46, 0x1a, 0x0e, 0x2e, 0x73, 0x6f, 0x6c, 0x64, 0x69,
	0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x64, 0x69, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_soldiers_proto_rawDescOnce sync.Once
	file_soldiers_proto_rawDescData = file_soldiers_proto_rawDesc
)

func file_soldiers_proto_rawDescGZIP() []byte {
	file_soldiers_proto_rawDescOnce.Do(func() {
		file_soldiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_soldiers_proto_rawDescData)
	})
	return file_soldiers_proto_rawDescData
}

var file_soldiers_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_soldiers_proto_goTypes = []interface{}{
	(*UseB)(nil),        // 0: soldiers.UseB
	(*UseF)(nil),        // 1: soldiers.UseF
	(*ById)(nil),        // 2: soldiers.ById
	(*Void)(nil),        // 3: soldiers.Void
	(*SoldierReq)(nil),  // 4: soldiers.SoldierReq
	(*Soldier)(nil),     // 5: soldiers.Soldier
	(*AllSoldiers)(nil), // 6: soldiers.AllSoldiers
}
var file_soldiers_proto_depIdxs = []int32{
	5, // 0: soldiers.AllSoldiers.soldiers:type_name -> soldiers.Soldier
	4, // 1: soldiers.SoldierService.Create:input_type -> soldiers.SoldierReq
	5, // 2: soldiers.SoldierService.Update:input_type -> soldiers.Soldier
	2, // 3: soldiers.SoldierService.Delete:input_type -> soldiers.ById
	2, // 4: soldiers.SoldierService.Get:input_type -> soldiers.ById
	4, // 5: soldiers.SoldierService.GetAll:input_type -> soldiers.SoldierReq
	0, // 6: soldiers.SoldierService.UseBullet:input_type -> soldiers.UseB
	1, // 7: soldiers.SoldierService.UseFuel:input_type -> soldiers.UseF
	3, // 8: soldiers.SoldierService.Create:output_type -> soldiers.Void
	3, // 9: soldiers.SoldierService.Update:output_type -> soldiers.Void
	3, // 10: soldiers.SoldierService.Delete:output_type -> soldiers.Void
	5, // 11: soldiers.SoldierService.Get:output_type -> soldiers.Soldier
	6, // 12: soldiers.SoldierService.GetAll:output_type -> soldiers.AllSoldiers
	3, // 13: soldiers.SoldierService.UseBullet:output_type -> soldiers.Void
	3, // 14: soldiers.SoldierService.UseFuel:output_type -> soldiers.Void
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_soldiers_proto_init() }
func file_soldiers_proto_init() {
	if File_soldiers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_soldiers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseB); i {
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
		file_soldiers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseF); i {
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
		file_soldiers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ById); i {
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
		file_soldiers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_soldiers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoldierReq); i {
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
		file_soldiers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Soldier); i {
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
		file_soldiers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSoldiers); i {
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
			RawDescriptor: file_soldiers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_soldiers_proto_goTypes,
		DependencyIndexes: file_soldiers_proto_depIdxs,
		MessageInfos:      file_soldiers_proto_msgTypes,
	}.Build()
	File_soldiers_proto = out.File
	file_soldiers_proto_rawDesc = nil
	file_soldiers_proto_goTypes = nil
	file_soldiers_proto_depIdxs = nil
}