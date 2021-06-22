// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: sorting.proto

package gen

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

type LoadItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *LoadItemsRequest) Reset() {
	*x = LoadItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadItemsRequest) ProtoMessage() {}

func (x *LoadItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadItemsRequest.ProtoReflect.Descriptor instead.
func (*LoadItemsRequest) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{0}
}

func (x *LoadItemsRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type LoadItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoadItemsResponse) Reset() {
	*x = LoadItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadItemsResponse) ProtoMessage() {}

func (x *LoadItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadItemsResponse.ProtoReflect.Descriptor instead.
func (*LoadItemsResponse) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{1}
}

type MoveItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cubby *Cubby `protobuf:"bytes,1,opt,name=cubby,proto3" json:"cubby,omitempty"`
}

func (x *MoveItemRequest) Reset() {
	*x = MoveItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveItemRequest) ProtoMessage() {}

func (x *MoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveItemRequest.ProtoReflect.Descriptor instead.
func (*MoveItemRequest) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{2}
}

func (x *MoveItemRequest) GetCubby() *Cubby {
	if x != nil {
		return x.Cubby
	}
	return nil
}

type RemoveItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemCodes []string `protobuf:"bytes,1,rep,name=itemCodes,proto3" json:"itemCodes,omitempty"`
}

func (x *RemoveItemsRequest) Reset() {
	*x = RemoveItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemsRequest) ProtoMessage() {}

func (x *RemoveItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemsRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemsRequest) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveItemsRequest) GetItemCodes() []string {
	if x != nil {
		return x.ItemCodes
	}
	return nil
}

type MoveItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MoveItemResponse) Reset() {
	*x = MoveItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveItemResponse) ProtoMessage() {}

func (x *MoveItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveItemResponse.ProtoReflect.Descriptor instead.
func (*MoveItemResponse) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{4}
}

type SelectItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SelectItemRequest) Reset() {
	*x = SelectItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectItemRequest) ProtoMessage() {}

func (x *SelectItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectItemRequest.ProtoReflect.Descriptor instead.
func (*SelectItemRequest) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{5}
}

type SelectItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *SelectItemResponse) Reset() {
	*x = SelectItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectItemResponse) ProtoMessage() {}

func (x *SelectItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectItemResponse.ProtoReflect.Descriptor instead.
func (*SelectItemResponse) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{6}
}

func (x *SelectItemResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type RemoveItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveItemsResponse) Reset() {
	*x = RemoveItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sorting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemsResponse) ProtoMessage() {}

func (x *RemoveItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sorting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemsResponse.ProtoReflect.Descriptor instead.
func (*RemoveItemsResponse) Descriptor() ([]byte, []int) {
	return file_sorting_proto_rawDescGZIP(), []int{7}
}

var File_sorting_proto protoreflect.FileDescriptor

var file_sorting_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x10,
	0x4c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a, 0x0f, 0x4d, 0x6f, 0x76, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x63,
	0x75, 0x62, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x62, 0x62, 0x79, 0x52, 0x05, 0x63, 0x75, 0x62, 0x62, 0x79, 0x22,
	0x32, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x12,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf0, 0x01, 0x0a, 0x0c, 0x53,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x4c,
	0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x11, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e,
	0x4d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6d, 0x69,
	0x74, 0x61, 0x72, 0x6b, 0x6f, 0x76, 0x61, 0x63, 0x68, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2d, 0x61, 0x74, 0x2d, 0x6f, 0x63, 0x61, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sorting_proto_rawDescOnce sync.Once
	file_sorting_proto_rawDescData = file_sorting_proto_rawDesc
)

func file_sorting_proto_rawDescGZIP() []byte {
	file_sorting_proto_rawDescOnce.Do(func() {
		file_sorting_proto_rawDescData = protoimpl.X.CompressGZIP(file_sorting_proto_rawDescData)
	})
	return file_sorting_proto_rawDescData
}

var file_sorting_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sorting_proto_goTypes = []interface{}{
	(*LoadItemsRequest)(nil),    // 0: LoadItemsRequest
	(*LoadItemsResponse)(nil),   // 1: LoadItemsResponse
	(*MoveItemRequest)(nil),     // 2: MoveItemRequest
	(*RemoveItemsRequest)(nil),  // 3: RemoveItemsRequest
	(*MoveItemResponse)(nil),    // 4: MoveItemResponse
	(*SelectItemRequest)(nil),   // 5: SelectItemRequest
	(*SelectItemResponse)(nil),  // 6: SelectItemResponse
	(*RemoveItemsResponse)(nil), // 7: RemoveItemsResponse
	(*Item)(nil),                // 8: types.Item
	(*Cubby)(nil),               // 9: types.Cubby
}
var file_sorting_proto_depIdxs = []int32{
	8, // 0: LoadItemsRequest.items:type_name -> types.Item
	9, // 1: MoveItemRequest.cubby:type_name -> types.Cubby
	8, // 2: SelectItemResponse.item:type_name -> types.Item
	0, // 3: SortingRobot.LoadItems:input_type -> LoadItemsRequest
	2, // 4: SortingRobot.MoveItem:input_type -> MoveItemRequest
	5, // 5: SortingRobot.SelectItem:input_type -> SelectItemRequest
	3, // 6: SortingRobot.RemoveItemsByCode:input_type -> RemoveItemsRequest
	1, // 7: SortingRobot.LoadItems:output_type -> LoadItemsResponse
	4, // 8: SortingRobot.MoveItem:output_type -> MoveItemResponse
	6, // 9: SortingRobot.SelectItem:output_type -> SelectItemResponse
	7, // 10: SortingRobot.RemoveItemsByCode:output_type -> RemoveItemsResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sorting_proto_init() }
func file_sorting_proto_init() {
	if File_sorting_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sorting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadItemsRequest); i {
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
		file_sorting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadItemsResponse); i {
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
		file_sorting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveItemRequest); i {
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
		file_sorting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemsRequest); i {
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
		file_sorting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveItemResponse); i {
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
		file_sorting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectItemRequest); i {
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
		file_sorting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectItemResponse); i {
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
		file_sorting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveItemsResponse); i {
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
			RawDescriptor: file_sorting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sorting_proto_goTypes,
		DependencyIndexes: file_sorting_proto_depIdxs,
		MessageInfos:      file_sorting_proto_msgTypes,
	}.Build()
	File_sorting_proto = out.File
	file_sorting_proto_rawDesc = nil
	file_sorting_proto_goTypes = nil
	file_sorting_proto_depIdxs = nil
}
