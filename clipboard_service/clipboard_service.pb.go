// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: clipboard_service/clipboard_service.proto

package clipboard_service

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

type CreateClipboardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values   []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	DeviceId string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *CreateClipboardsRequest) Reset() {
	*x = CreateClipboardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClipboardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClipboardsRequest) ProtoMessage() {}

func (x *CreateClipboardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClipboardsRequest.ProtoReflect.Descriptor instead.
func (*CreateClipboardsRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClipboardsRequest) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *CreateClipboardsRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type CreateClipboardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *CreateClipboardsResponse) Reset() {
	*x = CreateClipboardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClipboardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClipboardsResponse) ProtoMessage() {}

func (x *CreateClipboardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClipboardsResponse.ProtoReflect.Descriptor instead.
func (*CreateClipboardsResponse) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClipboardsResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetClipboardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetClipboardsRequest) Reset() {
	*x = GetClipboardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardsRequest) ProtoMessage() {}

func (x *GetClipboardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardsRequest.ProtoReflect.Descriptor instead.
func (*GetClipboardsRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetClipboardsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetClipboardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clipboards []*ClipboardItem `protobuf:"bytes,1,rep,name=clipboards,proto3" json:"clipboards,omitempty"`
}

func (x *GetClipboardsResponse) Reset() {
	*x = GetClipboardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardsResponse) ProtoMessage() {}

func (x *GetClipboardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardsResponse.ProtoReflect.Descriptor instead.
func (*GetClipboardsResponse) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetClipboardsResponse) GetClipboards() []*ClipboardItem {
	if x != nil {
		return x.Clipboards
	}
	return nil
}

type DeleteClipboardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteClipboardsRequest) Reset() {
	*x = DeleteClipboardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClipboardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClipboardsRequest) ProtoMessage() {}

func (x *DeleteClipboardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClipboardsRequest.ProtoReflect.Descriptor instead.
func (*DeleteClipboardsRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteClipboardsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteClipboardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteClipboardsResponse) Reset() {
	*x = DeleteClipboardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClipboardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClipboardsResponse) ProtoMessage() {}

func (x *DeleteClipboardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClipboardsResponse.ProtoReflect.Descriptor instead.
func (*DeleteClipboardsResponse) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteClipboardsResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SubscribeClipboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeClipboardRequest) Reset() {
	*x = SubscribeClipboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeClipboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeClipboardRequest) ProtoMessage() {}

func (x *SubscribeClipboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeClipboardRequest.ProtoReflect.Descriptor instead.
func (*SubscribeClipboardRequest) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{6}
}

type ClipboardMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items     []*ClipboardItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Operation string           `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *ClipboardMessage) Reset() {
	*x = ClipboardMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardMessage) ProtoMessage() {}

func (x *ClipboardMessage) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardMessage.ProtoReflect.Descriptor instead.
func (*ClipboardMessage) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{7}
}

func (x *ClipboardMessage) GetItems() []*ClipboardItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ClipboardMessage) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

type ClipboardItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content  string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	DeviceId string `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *ClipboardItem) Reset() {
	*x = ClipboardItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clipboard_service_clipboard_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClipboardItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClipboardItem) ProtoMessage() {}

func (x *ClipboardItem) ProtoReflect() protoreflect.Message {
	mi := &file_clipboard_service_clipboard_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClipboardItem.ProtoReflect.Descriptor instead.
func (*ClipboardItem) Descriptor() ([]byte, []int) {
	return file_clipboard_service_clipboard_service_proto_rawDescGZIP(), []int{8}
}

func (x *ClipboardItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClipboardItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ClipboardItem) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

var File_clipboard_service_clipboard_service_proto protoreflect.FileDescriptor

var file_clipboard_service_clipboard_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4e,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2c,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x28, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x59, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x22, 0x2b, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x34,
	0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x68, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x0d, 0x43,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x32, 0xc3, 0x03, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2a, 0x2e, 0x63,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a,
	0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x2c, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6d, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2a,
	0x2e, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x39, 0x33, 0x39, 0x33, 0x32, 0x33, 0x37,
	0x34, 0x39, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clipboard_service_clipboard_service_proto_rawDescOnce sync.Once
	file_clipboard_service_clipboard_service_proto_rawDescData = file_clipboard_service_clipboard_service_proto_rawDesc
)

func file_clipboard_service_clipboard_service_proto_rawDescGZIP() []byte {
	file_clipboard_service_clipboard_service_proto_rawDescOnce.Do(func() {
		file_clipboard_service_clipboard_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_clipboard_service_clipboard_service_proto_rawDescData)
	})
	return file_clipboard_service_clipboard_service_proto_rawDescData
}

var file_clipboard_service_clipboard_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_clipboard_service_clipboard_service_proto_goTypes = []interface{}{
	(*CreateClipboardsRequest)(nil),   // 0: clipboard_service.CreateClipboardsRequest
	(*CreateClipboardsResponse)(nil),  // 1: clipboard_service.CreateClipboardsResponse
	(*GetClipboardsRequest)(nil),      // 2: clipboard_service.GetClipboardsRequest
	(*GetClipboardsResponse)(nil),     // 3: clipboard_service.GetClipboardsResponse
	(*DeleteClipboardsRequest)(nil),   // 4: clipboard_service.DeleteClipboardsRequest
	(*DeleteClipboardsResponse)(nil),  // 5: clipboard_service.DeleteClipboardsResponse
	(*SubscribeClipboardRequest)(nil), // 6: clipboard_service.SubscribeClipboardRequest
	(*ClipboardMessage)(nil),          // 7: clipboard_service.ClipboardMessage
	(*ClipboardItem)(nil),             // 8: clipboard_service.ClipboardItem
}
var file_clipboard_service_clipboard_service_proto_depIdxs = []int32{
	8, // 0: clipboard_service.GetClipboardsResponse.clipboards:type_name -> clipboard_service.ClipboardItem
	8, // 1: clipboard_service.ClipboardMessage.items:type_name -> clipboard_service.ClipboardItem
	0, // 2: clipboard_service.ClipboardService.CreateClipboards:input_type -> clipboard_service.CreateClipboardsRequest
	2, // 3: clipboard_service.ClipboardService.GetClipboards:input_type -> clipboard_service.GetClipboardsRequest
	6, // 4: clipboard_service.ClipboardService.SubscribeClipboard:input_type -> clipboard_service.SubscribeClipboardRequest
	4, // 5: clipboard_service.ClipboardService.DeleteClipboards:input_type -> clipboard_service.DeleteClipboardsRequest
	1, // 6: clipboard_service.ClipboardService.CreateClipboards:output_type -> clipboard_service.CreateClipboardsResponse
	3, // 7: clipboard_service.ClipboardService.GetClipboards:output_type -> clipboard_service.GetClipboardsResponse
	7, // 8: clipboard_service.ClipboardService.SubscribeClipboard:output_type -> clipboard_service.ClipboardMessage
	5, // 9: clipboard_service.ClipboardService.DeleteClipboards:output_type -> clipboard_service.DeleteClipboardsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_clipboard_service_clipboard_service_proto_init() }
func file_clipboard_service_clipboard_service_proto_init() {
	if File_clipboard_service_clipboard_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clipboard_service_clipboard_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClipboardsRequest); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClipboardsResponse); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardsRequest); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardsResponse); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClipboardsRequest); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClipboardsResponse); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeClipboardRequest); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardMessage); i {
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
		file_clipboard_service_clipboard_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClipboardItem); i {
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
			RawDescriptor: file_clipboard_service_clipboard_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clipboard_service_clipboard_service_proto_goTypes,
		DependencyIndexes: file_clipboard_service_clipboard_service_proto_depIdxs,
		MessageInfos:      file_clipboard_service_clipboard_service_proto_msgTypes,
	}.Build()
	File_clipboard_service_clipboard_service_proto = out.File
	file_clipboard_service_clipboard_service_proto_rawDesc = nil
	file_clipboard_service_clipboard_service_proto_goTypes = nil
	file_clipboard_service_clipboard_service_proto_depIdxs = nil
}
