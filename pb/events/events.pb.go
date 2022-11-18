// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: pb/events/events.proto

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

// Event
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	City     string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Postcode string `protobuf:"bytes,3,opt,name=postcode,proto3" json:"postcode,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Location) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Location) GetPostcode() string {
	if x != nil {
		return x.Postcode
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RangeStart float64 `protobuf:"fixed64,1,opt,name=rangeStart,proto3" json:"rangeStart,omitempty"`
	RangeEnd   float64 `protobuf:"fixed64,2,opt,name=rangeEnd,proto3" json:"rangeEnd,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetRangeStart() float64 {
	if x != nil {
		return x.RangeStart
	}
	return 0
}

func (x *Price) GetRangeEnd() float64 {
	if x != nil {
		return x.RangeEnd
	}
	return 0
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date        string    `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Location    *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Description string    `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Tags        []string  `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Category    string    `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty"`
	Price       *Price    `protobuf:"bytes,8,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Event) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Event) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Event) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

// Find One
type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Event  *Event `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// Find Many
type FindManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *string `protobuf:"bytes,1,opt,name=category,proto3,oneof" json:"category,omitempty"`
	City     *string `protobuf:"bytes,2,opt,name=city,proto3,oneof" json:"city,omitempty"`
	Price    *Price  `protobuf:"bytes,3,opt,name=price,proto3,oneof" json:"price,omitempty"`
}

func (x *FindManyRequest) Reset() {
	*x = FindManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyRequest) ProtoMessage() {}

func (x *FindManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyRequest.ProtoReflect.Descriptor instead.
func (*FindManyRequest) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{5}
}

func (x *FindManyRequest) GetCategory() string {
	if x != nil && x.Category != nil {
		return *x.Category
	}
	return ""
}

func (x *FindManyRequest) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *FindManyRequest) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

type FindManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Events []*Event `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *FindManyResponse) Reset() {
	*x = FindManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_events_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindManyResponse) ProtoMessage() {}

func (x *FindManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_events_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindManyResponse.ProtoReflect.Descriptor instead.
func (*FindManyResponse) Descriptor() ([]byte, []int) {
	return file_pb_events_events_proto_rawDescGZIP(), []int{6}
}

func (x *FindManyResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindManyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindManyResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_pb_events_events_proto protoreflect.FileDescriptor

var file_pb_events_events_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x54, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x43, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x95, 0x01, 0x0a,
	0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x48, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x87, 0x01,
	0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61,
	0x6e, 0x79, 0x12, 0x17, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_events_events_proto_rawDescOnce sync.Once
	file_pb_events_events_proto_rawDescData = file_pb_events_events_proto_rawDesc
)

func file_pb_events_events_proto_rawDescGZIP() []byte {
	file_pb_events_events_proto_rawDescOnce.Do(func() {
		file_pb_events_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_events_events_proto_rawDescData)
	})
	return file_pb_events_events_proto_rawDescData
}

var file_pb_events_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_events_events_proto_goTypes = []interface{}{
	(*Location)(nil),         // 0: events.Location
	(*Price)(nil),            // 1: events.Price
	(*Event)(nil),            // 2: events.Event
	(*FindOneRequest)(nil),   // 3: events.FindOneRequest
	(*FindOneResponse)(nil),  // 4: events.FindOneResponse
	(*FindManyRequest)(nil),  // 5: events.FindManyRequest
	(*FindManyResponse)(nil), // 6: events.FindManyResponse
}
var file_pb_events_events_proto_depIdxs = []int32{
	0, // 0: events.Event.location:type_name -> events.Location
	1, // 1: events.Event.price:type_name -> events.Price
	2, // 2: events.FindOneResponse.event:type_name -> events.Event
	1, // 3: events.FindManyRequest.price:type_name -> events.Price
	2, // 4: events.FindManyResponse.events:type_name -> events.Event
	3, // 5: events.Events.FindOne:input_type -> events.FindOneRequest
	5, // 6: events.Events.FindMany:input_type -> events.FindManyRequest
	4, // 7: events.Events.FindOne:output_type -> events.FindOneResponse
	6, // 8: events.Events.FindMany:output_type -> events.FindManyResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pb_events_events_proto_init() }
func file_pb_events_events_proto_init() {
	if File_pb_events_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_events_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_pb_events_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_pb_events_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_pb_events_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_pb_events_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResponse); i {
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
		file_pb_events_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindManyRequest); i {
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
		file_pb_events_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindManyResponse); i {
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
	file_pb_events_events_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_events_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_events_events_proto_goTypes,
		DependencyIndexes: file_pb_events_events_proto_depIdxs,
		MessageInfos:      file_pb_events_events_proto_msgTypes,
	}.Build()
	File_pb_events_events_proto = out.File
	file_pb_events_events_proto_rawDesc = nil
	file_pb_events_events_proto_goTypes = nil
	file_pb_events_events_proto_depIdxs = nil
}
