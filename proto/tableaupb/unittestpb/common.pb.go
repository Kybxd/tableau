// Protoconf - Tableau's data interchange format
// https://tableauio.github.io/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: tableau/protobuf/unittest/common.proto

package unittestpb

import (
	_ "github.com/tableauio/tableau/proto/tableaupb"
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

type FruitType int32

const (
	FruitType_FRUIT_TYPE_UNKNOWN FruitType = 0
	FruitType_FRUIT_TYPE_APPLE   FruitType = 1
	FruitType_FRUIT_TYPE_ORANGE  FruitType = 3
	FruitType_FRUIT_TYPE_BANANA  FruitType = 4
)

// Enum value maps for FruitType.
var (
	FruitType_name = map[int32]string{
		0: "FRUIT_TYPE_UNKNOWN",
		1: "FRUIT_TYPE_APPLE",
		3: "FRUIT_TYPE_ORANGE",
		4: "FRUIT_TYPE_BANANA",
	}
	FruitType_value = map[string]int32{
		"FRUIT_TYPE_UNKNOWN": 0,
		"FRUIT_TYPE_APPLE":   1,
		"FRUIT_TYPE_ORANGE":  3,
		"FRUIT_TYPE_BANANA":  4,
	}
)

func (x FruitType) Enum() *FruitType {
	p := new(FruitType)
	*p = x
	return p
}

func (x FruitType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FruitType) Descriptor() protoreflect.EnumDescriptor {
	return file_tableau_protobuf_unittest_common_proto_enumTypes[0].Descriptor()
}

func (FruitType) Type() protoreflect.EnumType {
	return &file_tableau_protobuf_unittest_common_proto_enumTypes[0]
}

func (x FruitType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FruitType.Descriptor instead.
func (FruitType) EnumDescriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{0}
}

type FruitFlavor int32

const (
	FruitFlavor_FRUIT_FLAVOR_UNKNOWN  FruitFlavor = 0
	FruitFlavor_FRUIT_FLAVOR_FRAGRANT FruitFlavor = 1
	FruitFlavor_FRUIT_FLAVOR_SOUR     FruitFlavor = 2
	FruitFlavor_FRUIT_FLAVOR_SWEET    FruitFlavor = 3
)

// Enum value maps for FruitFlavor.
var (
	FruitFlavor_name = map[int32]string{
		0: "FRUIT_FLAVOR_UNKNOWN",
		1: "FRUIT_FLAVOR_FRAGRANT",
		2: "FRUIT_FLAVOR_SOUR",
		3: "FRUIT_FLAVOR_SWEET",
	}
	FruitFlavor_value = map[string]int32{
		"FRUIT_FLAVOR_UNKNOWN":  0,
		"FRUIT_FLAVOR_FRAGRANT": 1,
		"FRUIT_FLAVOR_SOUR":     2,
		"FRUIT_FLAVOR_SWEET":    3,
	}
)

func (x FruitFlavor) Enum() *FruitFlavor {
	p := new(FruitFlavor)
	*p = x
	return p
}

func (x FruitFlavor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FruitFlavor) Descriptor() protoreflect.EnumDescriptor {
	return file_tableau_protobuf_unittest_common_proto_enumTypes[1].Descriptor()
}

func (FruitFlavor) Type() protoreflect.EnumType {
	return &file_tableau_protobuf_unittest_common_proto_enumTypes[1]
}

func (x FruitFlavor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FruitFlavor.Descriptor instead.
func (FruitFlavor) EnumDescriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1}
}

type Target_Type int32

const (
	Target_TYPE_NIL Target_Type = 0
	Target_TYPE_PVP Target_Type = 1
	Target_TYPE_PVE Target_Type = 2
)

// Enum value maps for Target_Type.
var (
	Target_Type_name = map[int32]string{
		0: "TYPE_NIL",
		1: "TYPE_PVP",
		2: "TYPE_PVE",
	}
	Target_Type_value = map[string]int32{
		"TYPE_NIL": 0,
		"TYPE_PVP": 1,
		"TYPE_PVE": 2,
	}
)

func (x Target_Type) Enum() *Target_Type {
	p := new(Target_Type)
	*p = x
	return p
}

func (x Target_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Target_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_tableau_protobuf_unittest_common_proto_enumTypes[2].Descriptor()
}

func (Target_Type) Type() protoreflect.EnumType {
	return &file_tableau_protobuf_unittest_common_proto_enumTypes[2]
}

func (x Target_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Target_Type.Descriptor instead.
func (Target_Type) EnumDescriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1, 0}
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int32  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Target_Type `protobuf:"varint,9999,opt,name=type,proto3,enum=unittest.Target_Type" json:"type,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Target_Pvp_
	//	*Target_Pve_
	Value isTarget_Value `protobuf_oneof:"value"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1}
}

func (x *Target) GetType() Target_Type {
	if x != nil {
		return x.Type
	}
	return Target_TYPE_NIL
}

func (m *Target) GetValue() isTarget_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Target) GetPvp() *Target_Pvp {
	if x, ok := x.GetValue().(*Target_Pvp_); ok {
		return x.Pvp
	}
	return nil
}

func (x *Target) GetPve() *Target_Pve {
	if x, ok := x.GetValue().(*Target_Pve_); ok {
		return x.Pve
	}
	return nil
}

type isTarget_Value interface {
	isTarget_Value()
}

type Target_Pvp_ struct {
	Pvp *Target_Pvp `protobuf:"bytes,1,opt,name=pvp,proto3,oneof"` // Bound to enum value: TYPE_PVP.
}

type Target_Pve_ struct {
	Pve *Target_Pve `protobuf:"bytes,2,opt,name=pve,proto3,oneof"` // Bound to enum value: TYPE_PVP.
}

func (*Target_Pvp_) isTarget_Value() {}

func (*Target_Pve_) isTarget_Value() {}

type Target_Pvp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Health uint32 `protobuf:"varint,2,opt,name=health,proto3" json:"health,omitempty"`
	Damage int64  `protobuf:"varint,3,opt,name=damage,proto3" json:"damage,omitempty"`
}

func (x *Target_Pvp) Reset() {
	*x = Target_Pvp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target_Pvp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target_Pvp) ProtoMessage() {}

func (x *Target_Pvp) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target_Pvp.ProtoReflect.Descriptor instead.
func (*Target_Pvp) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Target_Pvp) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Target_Pvp) GetHealth() uint32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Target_Pvp) GetDamage() int64 {
	if x != nil {
		return x.Damage
	}
	return 0
}

type Target_Pve struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mission  *Target_Pve_Mission `protobuf:"bytes,1,opt,name=mission,proto3" json:"mission,omitempty"`                                                                                             // incell struct
	Heros    []int32             `protobuf:"varint,2,rep,packed,name=heros,proto3" json:"heros,omitempty"`                                                                                         // incell list
	Dungeons map[int32]int64     `protobuf:"bytes,3,rep,name=dungeons,proto3" json:"dungeons,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // incell map
}

func (x *Target_Pve) Reset() {
	*x = Target_Pve{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target_Pve) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target_Pve) ProtoMessage() {}

func (x *Target_Pve) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target_Pve.ProtoReflect.Descriptor instead.
func (*Target_Pve) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Target_Pve) GetMission() *Target_Pve_Mission {
	if x != nil {
		return x.Mission
	}
	return nil
}

func (x *Target_Pve) GetHeros() []int32 {
	if x != nil {
		return x.Heros
	}
	return nil
}

func (x *Target_Pve) GetDungeons() map[int32]int64 {
	if x != nil {
		return x.Dungeons
	}
	return nil
}

type Target_Pve_Mission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Level  uint32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Damage int64  `protobuf:"varint,3,opt,name=damage,proto3" json:"damage,omitempty"`
}

func (x *Target_Pve_Mission) Reset() {
	*x = Target_Pve_Mission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target_Pve_Mission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target_Pve_Mission) ProtoMessage() {}

func (x *Target_Pve_Mission) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_unittest_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target_Pve_Mission.ProtoReflect.Descriptor instead.
func (*Target_Pve_Mission) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_unittest_common_proto_rawDescGZIP(), []int{1, 1, 1}
}

func (x *Target_Pve_Mission) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Target_Pve_Mission) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Target_Pve_Mission) GetDamage() int64 {
	if x != nil {
		return x.Damage
	}
	return 0
}

var File_tableau_protobuf_unittest_common_proto protoreflect.FileDescriptor

var file_tableau_protobuf_unittest_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x0a, 0x02, 0x49, 0x44,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x09, 0x82, 0xb5, 0x18, 0x05, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x52, 0x03, 0x6e, 0x75,
	0x6d, 0x22, 0xee, 0x04, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x8f, 0x4e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x0a, 0x82, 0xb5, 0x18, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x70, 0x76, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x75, 0x6e, 0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x50, 0x76, 0x70, 0x48, 0x00, 0x52, 0x03, 0x70, 0x76, 0x70, 0x12, 0x28,
	0x0a, 0x03, 0x70, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x50, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x03, 0x70, 0x76, 0x65, 0x1a, 0x49, 0x0a, 0x03, 0x50, 0x76, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x1a, 0xa8, 0x02, 0x0a, 0x03, 0x50, 0x76, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75,
	0x6e, 0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x50,
	0x76, 0x65, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x82, 0xb5, 0x18, 0x09,
	0x0a, 0x07, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x65, 0x72, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x68, 0x65, 0x72, 0x6f, 0x73, 0x12, 0x3e, 0x0a, 0x08, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x75, 0x6e, 0x69,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x50, 0x76, 0x65,
	0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47, 0x0a, 0x07, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x46,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e,
	0x49, 0x4c, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x56, 0x50,
	0x10, 0x01, 0x1a, 0x09, 0x82, 0xb5, 0x18, 0x05, 0x0a, 0x03, 0x50, 0x56, 0x50, 0x12, 0x17, 0x0a,
	0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x56, 0x45, 0x10, 0x02, 0x1a, 0x09, 0x82, 0xb5, 0x18,
	0x05, 0x0a, 0x03, 0x50, 0x56, 0x45, 0x3a, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0b, 0x82, 0xb5, 0x18, 0x07, 0x12, 0x05, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x2a, 0x9f, 0x01, 0x0a, 0x09, 0x46, 0x72, 0x75, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x12, 0x46, 0x52, 0x55, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0d, 0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x21, 0x0a, 0x10, 0x46, 0x52, 0x55, 0x49, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45, 0x10, 0x01, 0x1a, 0x0b, 0x82,
	0xb5, 0x18, 0x07, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x11, 0x46, 0x52,
	0x55, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x03, 0x1a, 0x0c, 0x82, 0xb5, 0x18, 0x08, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x23, 0x0a, 0x11, 0x46, 0x52, 0x55, 0x49, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41,
	0x4e, 0x41, 0x4e, 0x41, 0x10, 0x04, 0x1a, 0x0c, 0x82, 0xb5, 0x18, 0x08, 0x0a, 0x06, 0x42, 0x61,
	0x6e, 0x61, 0x6e, 0x61, 0x2a, 0xa9, 0x01, 0x0a, 0x0b, 0x46, 0x72, 0x75, 0x69, 0x74, 0x46, 0x6c,
	0x61, 0x76, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x14, 0x46, 0x52, 0x55, 0x49, 0x54, 0x5f, 0x46, 0x4c,
	0x41, 0x56, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0d,
	0x82, 0xb5, 0x18, 0x09, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x12, 0x29, 0x0a,
	0x15, 0x46, 0x52, 0x55, 0x49, 0x54, 0x5f, 0x46, 0x4c, 0x41, 0x56, 0x4f, 0x52, 0x5f, 0x46, 0x52,
	0x41, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x1a, 0x0e, 0x82, 0xb5, 0x18, 0x0a, 0x0a, 0x08,
	0x46, 0x72, 0x61, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x11, 0x46, 0x52, 0x55, 0x49,
	0x54, 0x5f, 0x46, 0x4c, 0x41, 0x56, 0x4f, 0x52, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x10, 0x02, 0x1a,
	0x0a, 0x82, 0xb5, 0x18, 0x06, 0x0a, 0x04, 0x53, 0x6f, 0x75, 0x72, 0x12, 0x23, 0x0a, 0x12, 0x46,
	0x52, 0x55, 0x49, 0x54, 0x5f, 0x46, 0x4c, 0x41, 0x56, 0x4f, 0x52, 0x5f, 0x53, 0x57, 0x45, 0x45,
	0x54, 0x10, 0x03, 0x1a, 0x0b, 0x82, 0xb5, 0x18, 0x07, 0x0a, 0x05, 0x53, 0x77, 0x65, 0x65, 0x74,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x69, 0x6f, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x70, 0x62,
	0x2f, 0x75, 0x6e, 0x69, 0x74, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tableau_protobuf_unittest_common_proto_rawDescOnce sync.Once
	file_tableau_protobuf_unittest_common_proto_rawDescData = file_tableau_protobuf_unittest_common_proto_rawDesc
)

func file_tableau_protobuf_unittest_common_proto_rawDescGZIP() []byte {
	file_tableau_protobuf_unittest_common_proto_rawDescOnce.Do(func() {
		file_tableau_protobuf_unittest_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableau_protobuf_unittest_common_proto_rawDescData)
	})
	return file_tableau_protobuf_unittest_common_proto_rawDescData
}

var file_tableau_protobuf_unittest_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tableau_protobuf_unittest_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tableau_protobuf_unittest_common_proto_goTypes = []interface{}{
	(FruitType)(0),             // 0: unittest.FruitType
	(FruitFlavor)(0),           // 1: unittest.FruitFlavor
	(Target_Type)(0),           // 2: unittest.Target.Type
	(*Item)(nil),               // 3: unittest.Item
	(*Target)(nil),             // 4: unittest.Target
	(*Target_Pvp)(nil),         // 5: unittest.Target.Pvp
	(*Target_Pve)(nil),         // 6: unittest.Target.Pve
	nil,                        // 7: unittest.Target.Pve.DungeonsEntry
	(*Target_Pve_Mission)(nil), // 8: unittest.Target.Pve.Mission
}
var file_tableau_protobuf_unittest_common_proto_depIdxs = []int32{
	2, // 0: unittest.Target.type:type_name -> unittest.Target.Type
	5, // 1: unittest.Target.pvp:type_name -> unittest.Target.Pvp
	6, // 2: unittest.Target.pve:type_name -> unittest.Target.Pve
	8, // 3: unittest.Target.Pve.mission:type_name -> unittest.Target.Pve.Mission
	7, // 4: unittest.Target.Pve.dungeons:type_name -> unittest.Target.Pve.DungeonsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tableau_protobuf_unittest_common_proto_init() }
func file_tableau_protobuf_unittest_common_proto_init() {
	if File_tableau_protobuf_unittest_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tableau_protobuf_unittest_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_tableau_protobuf_unittest_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_tableau_protobuf_unittest_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target_Pvp); i {
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
		file_tableau_protobuf_unittest_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target_Pve); i {
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
		file_tableau_protobuf_unittest_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target_Pve_Mission); i {
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
	file_tableau_protobuf_unittest_common_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Target_Pvp_)(nil),
		(*Target_Pve_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tableau_protobuf_unittest_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableau_protobuf_unittest_common_proto_goTypes,
		DependencyIndexes: file_tableau_protobuf_unittest_common_proto_depIdxs,
		EnumInfos:         file_tableau_protobuf_unittest_common_proto_enumTypes,
		MessageInfos:      file_tableau_protobuf_unittest_common_proto_msgTypes,
	}.Build()
	File_tableau_protobuf_unittest_common_proto = out.File
	file_tableau_protobuf_unittest_common_proto_rawDesc = nil
	file_tableau_protobuf_unittest_common_proto_goTypes = nil
	file_tableau_protobuf_unittest_common_proto_depIdxs = nil
}
