// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tableau/protobuf/meta.proto

package tableaupb

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

type WorkbookMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SheetMetaMap map[string]*SheetMeta `protobuf:"bytes,1,rep,name=sheet_meta_map,json=sheetMetaMap,proto3" json:"sheet_meta_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkbookMeta) Reset() {
	*x = WorkbookMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkbookMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkbookMeta) ProtoMessage() {}

func (x *WorkbookMeta) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkbookMeta.ProtoReflect.Descriptor instead.
func (*WorkbookMeta) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_meta_proto_rawDescGZIP(), []int{0}
}

func (x *WorkbookMeta) GetSheetMetaMap() map[string]*SheetMeta {
	if x != nil {
		return x.SheetMetaMap
	}
	return nil
}

type SheetMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sheet     string `protobuf:"bytes,1,opt,name=sheet,proto3" json:"sheet,omitempty"`
	Alias     string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	Namerow   int32  `protobuf:"varint,3,opt,name=namerow,proto3" json:"namerow,omitempty"`
	Typerow   int32  `protobuf:"varint,4,opt,name=typerow,proto3" json:"typerow,omitempty"`
	Noterow   int32  `protobuf:"varint,5,opt,name=noterow,proto3" json:"noterow,omitempty"`
	Datarow   int32  `protobuf:"varint,6,opt,name=datarow,proto3" json:"datarow,omitempty"`
	Nameline  int32  `protobuf:"varint,7,opt,name=nameline,proto3" json:"nameline,omitempty"`
	Typeline  int32  `protobuf:"varint,8,opt,name=typeline,proto3" json:"typeline,omitempty"`
	Transpose bool   `protobuf:"varint,9,opt,name=transpose,proto3" json:"transpose,omitempty"`
	// nested naming of namerow
	Nested bool   `protobuf:"varint,10,opt,name=nested,proto3" json:"nested,omitempty"`
	Sep    string `protobuf:"bytes,11,opt,name=sep,proto3" json:"sep,omitempty"`
	Subsep string `protobuf:"bytes,12,opt,name=subsep,proto3" json:"subsep,omitempty"`
	// merger of multiple sheets: [Element]...
	// each element is:
	// - a workbook name or Glob(https://pkg.go.dev/path/filepath#Glob) to merge (relative to this workbook): <Workbook>,
	//   then the sheet name is the same as this sheet.
	// - or a workbook name (relative to this workbook) with a worksheet name: <Workbook>#<Worksheet>.
	Merger []string `protobuf:"bytes,13,rep,name=merger,proto3" json:"merger,omitempty"`
	// Tableau will merge adjacent rows with the same key. If the key cell is not set,
	// it will be treated as the same as the most nearest key above the same column.
	//
	// This option is only useful for map or keyed-list.
	AdjacentKey bool `protobuf:"varint,14,opt,name=adjacent_key,json=adjacentKey,proto3" json:"adjacent_key,omitempty"`
	////////// Loader related options below //////////
	// generate ordered map accessers
	OrderedMap bool `protobuf:"varint,50,opt,name=ordered_map,json=orderedMap,proto3" json:"ordered_map,omitempty"`
	// generate index accessers.
	// The key-value pair is `<IndexColumnName>[:IndexName]`, if IndexName is not set,
	// it will be this column's parent struct type name.
	//
	// Generated APIs are:
	//
	// C++:
	// - Repeated<const STRUCT_TYPE&> Find<IndexName>(INDEX_TYPE index) const;
	// - const STRUCT_TYPE* FindFirst<IndexName>(INDEX_TYPE index);
	Index map[string]string `protobuf:"bytes,51,rep,name=index,proto3" json:"index,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SheetMeta) Reset() {
	*x = SheetMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableau_protobuf_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SheetMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SheetMeta) ProtoMessage() {}

func (x *SheetMeta) ProtoReflect() protoreflect.Message {
	mi := &file_tableau_protobuf_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SheetMeta.ProtoReflect.Descriptor instead.
func (*SheetMeta) Descriptor() ([]byte, []int) {
	return file_tableau_protobuf_meta_proto_rawDescGZIP(), []int{1}
}

func (x *SheetMeta) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *SheetMeta) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SheetMeta) GetNamerow() int32 {
	if x != nil {
		return x.Namerow
	}
	return 0
}

func (x *SheetMeta) GetTyperow() int32 {
	if x != nil {
		return x.Typerow
	}
	return 0
}

func (x *SheetMeta) GetNoterow() int32 {
	if x != nil {
		return x.Noterow
	}
	return 0
}

func (x *SheetMeta) GetDatarow() int32 {
	if x != nil {
		return x.Datarow
	}
	return 0
}

func (x *SheetMeta) GetNameline() int32 {
	if x != nil {
		return x.Nameline
	}
	return 0
}

func (x *SheetMeta) GetTypeline() int32 {
	if x != nil {
		return x.Typeline
	}
	return 0
}

func (x *SheetMeta) GetTranspose() bool {
	if x != nil {
		return x.Transpose
	}
	return false
}

func (x *SheetMeta) GetNested() bool {
	if x != nil {
		return x.Nested
	}
	return false
}

func (x *SheetMeta) GetSep() string {
	if x != nil {
		return x.Sep
	}
	return ""
}

func (x *SheetMeta) GetSubsep() string {
	if x != nil {
		return x.Subsep
	}
	return ""
}

func (x *SheetMeta) GetMerger() []string {
	if x != nil {
		return x.Merger
	}
	return nil
}

func (x *SheetMeta) GetAdjacentKey() bool {
	if x != nil {
		return x.AdjacentKey
	}
	return false
}

func (x *SheetMeta) GetOrderedMap() bool {
	if x != nil {
		return x.OrderedMap
	}
	return false
}

func (x *SheetMeta) GetIndex() map[string]string {
	if x != nil {
		return x.Index
	}
	return nil
}

var File_tableau_protobuf_meta_proto protoreflect.FileDescriptor

var file_tableau_protobuf_meta_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x1a, 0x1e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x62,
	0x6f, 0x6f, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x5a, 0x0a, 0x0e, 0x73, 0x68, 0x65, 0x65, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x6f,
	0x6f, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0b, 0x82, 0xb5, 0x18, 0x07, 0x1a, 0x05,
	0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x0c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x4d, 0x61, 0x70, 0x1a, 0x53, 0x0a, 0x11, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x61, 0x75, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x0a, 0x08,
	0x40, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x41, 0x55, 0x10, 0x01, 0x28, 0x02, 0x22, 0x92, 0x06, 0x0a,
	0x09, 0x53, 0x68, 0x65, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x82, 0xb5, 0x18, 0x07, 0x0a,
	0x05, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x23, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x82, 0xb5,
	0x18, 0x09, 0x0a, 0x05, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x40, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x6f, 0x77, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xb5, 0x18, 0x0b, 0x0a, 0x07, 0x4e, 0x61, 0x6d, 0x65, 0x72,
	0x6f, 0x77, 0x40, 0x01, 0x52, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x72, 0x6f, 0x77, 0x12, 0x29, 0x0a,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x72, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f,
	0x82, 0xb5, 0x18, 0x0b, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x72, 0x6f, 0x77, 0x40, 0x01, 0x52,
	0x07, 0x74, 0x79, 0x70, 0x65, 0x72, 0x6f, 0x77, 0x12, 0x29, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65,
	0x72, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xb5, 0x18, 0x0b, 0x0a,
	0x07, 0x4e, 0x6f, 0x74, 0x65, 0x72, 0x6f, 0x77, 0x40, 0x01, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x65,
	0x72, 0x6f, 0x77, 0x12, 0x29, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x72, 0x6f, 0x77, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xb5, 0x18, 0x0b, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61,
	0x72, 0x6f, 0x77, 0x40, 0x01, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x72, 0x6f, 0x77, 0x12, 0x2c,
	0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x10, 0x82, 0xb5, 0x18, 0x0c, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x40, 0x01, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x74, 0x79, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x10,
	0x82, 0xb5, 0x18, 0x0c, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x40, 0x01,
	0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x11, 0x82,
	0xb5, 0x18, 0x0d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x40, 0x01,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0e, 0x82, 0xb5, 0x18,
	0x0a, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x40, 0x01, 0x52, 0x06, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x73, 0x65, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0x82, 0xb5, 0x18, 0x07, 0x0a, 0x03, 0x53, 0x65, 0x70, 0x40, 0x01, 0x52, 0x03, 0x73,
	0x65, 0x70, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x73, 0x65, 0x70, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0x82, 0xb5, 0x18, 0x0a, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x73, 0x65, 0x70,
	0x40, 0x01, 0x52, 0x06, 0x73, 0x75, 0x62, 0x73, 0x65, 0x70, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x42, 0x10, 0x82, 0xb5, 0x18, 0x0c,
	0x0a, 0x06, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x72, 0x20, 0x03, 0x40, 0x01, 0x52, 0x06, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x42, 0x13, 0x82, 0xb5, 0x18, 0x0f,
	0x0a, 0x0b, 0x41, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x40, 0x01, 0x52,
	0x0b, 0x61, 0x64, 0x6a, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x12, 0x82, 0xb5, 0x18, 0x0e, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64,
	0x4d, 0x61, 0x70, 0x40, 0x01, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x61,
	0x70, 0x12, 0x44, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x33, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x2e, 0x53, 0x68, 0x65, 0x65, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x0f, 0x82, 0xb5, 0x18, 0x0b, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x20, 0x03, 0x40, 0x01,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x38, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x69, 0x6f, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61,
	0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x75, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableau_protobuf_meta_proto_rawDescOnce sync.Once
	file_tableau_protobuf_meta_proto_rawDescData = file_tableau_protobuf_meta_proto_rawDesc
)

func file_tableau_protobuf_meta_proto_rawDescGZIP() []byte {
	file_tableau_protobuf_meta_proto_rawDescOnce.Do(func() {
		file_tableau_protobuf_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableau_protobuf_meta_proto_rawDescData)
	})
	return file_tableau_protobuf_meta_proto_rawDescData
}

var file_tableau_protobuf_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tableau_protobuf_meta_proto_goTypes = []interface{}{
	(*WorkbookMeta)(nil), // 0: tableau.WorkbookMeta
	(*SheetMeta)(nil),    // 1: tableau.SheetMeta
	nil,                  // 2: tableau.WorkbookMeta.SheetMetaMapEntry
	nil,                  // 3: tableau.SheetMeta.IndexEntry
}
var file_tableau_protobuf_meta_proto_depIdxs = []int32{
	2, // 0: tableau.WorkbookMeta.sheet_meta_map:type_name -> tableau.WorkbookMeta.SheetMetaMapEntry
	3, // 1: tableau.SheetMeta.index:type_name -> tableau.SheetMeta.IndexEntry
	1, // 2: tableau.WorkbookMeta.SheetMetaMapEntry.value:type_name -> tableau.SheetMeta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tableau_protobuf_meta_proto_init() }
func file_tableau_protobuf_meta_proto_init() {
	if File_tableau_protobuf_meta_proto != nil {
		return
	}
	file_tableau_protobuf_tableau_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tableau_protobuf_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkbookMeta); i {
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
		file_tableau_protobuf_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SheetMeta); i {
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
			RawDescriptor: file_tableau_protobuf_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableau_protobuf_meta_proto_goTypes,
		DependencyIndexes: file_tableau_protobuf_meta_proto_depIdxs,
		MessageInfos:      file_tableau_protobuf_meta_proto_msgTypes,
	}.Build()
	File_tableau_protobuf_meta_proto = out.File
	file_tableau_protobuf_meta_proto_rawDesc = nil
	file_tableau_protobuf_meta_proto_goTypes = nil
	file_tableau_protobuf_meta_proto_depIdxs = nil
}
