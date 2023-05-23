// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: project/v1/project.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Defines the GetRequest message which is used to retrieve a specific project.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the product the project belongs to.
	ProductName string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
	// The name of the project being retrieved.
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,proto3" json:"project_name,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

// Defines the GetReply message which is used to return a specific project.
type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the product the project belongs to.
	Product string `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// The name of the project.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The language used in the project.
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{1}
}

func (x *GetReply) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *GetReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetReply) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

// Defines the ListsRequest message which is used to retrieve a list of projects.
type ListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the product the projects belong to.
	ProductName string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
}

func (x *ListsRequest) Reset() {
	*x = ListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsRequest) ProtoMessage() {}

func (x *ListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsRequest.ProtoReflect.Descriptor instead.
func (*ListsRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{2}
}

func (x *ListsRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

// Defines the ListsReply message which is used to return a list of projects.
type ListsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of projects being returned.
	Items []*GetReply `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListsReply) Reset() {
	*x = ListsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsReply) ProtoMessage() {}

func (x *ListsReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListsReply.ProtoReflect.Descriptor instead.
func (*ListsReply) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{3}
}

func (x *ListsReply) GetItems() []*GetReply {
	if x != nil {
		return x.Items
	}
	return nil
}

// Defines the SaveRequest message which is used to create or update a project.
type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the product the project belongs to.
	ProductName string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
	// The name of the project being created or updated.
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,proto3" json:"project_name,omitempty"`
	// Whether or not to skip validation.
	InsecureSkipCheck bool `protobuf:"varint,4,opt,name=insecure_skip_check,json=insecureSkipCheck,proto3" json:"insecure_skip_check,omitempty"`
	// The request body for the project.
	Body *SaveRequest_Body `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{4}
}

func (x *SaveRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SaveRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *SaveRequest) GetInsecureSkipCheck() bool {
	if x != nil {
		return x.InsecureSkipCheck
	}
	return false
}

func (x *SaveRequest) GetBody() *SaveRequest_Body {
	if x != nil {
		return x.Body
	}
	return nil
}

// Defines the SaveReply message which is used to return a message after creating or updating a project.
type SaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message being returned.
	Msg string `protobuf:"bytes,1,opt,name=msg,json=message,proto3" json:"msg,omitempty"`
}

func (x *SaveReply) Reset() {
	*x = SaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReply) ProtoMessage() {}

func (x *SaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveReply.ProtoReflect.Descriptor instead.
func (*SaveReply) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{5}
}

func (x *SaveReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// Defines the DeleteRequest message which is used to delete a project.
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the project being deleted.
	ProjectName string `protobuf:"bytes,1,opt,name=project_name,proto3" json:"project_name,omitempty"`
	// The name of the product the project belongs to.
	ProductName string `protobuf:"bytes,2,opt,name=product_name,proto3" json:"product_name,omitempty"`
	// Whether or not to skip validation.
	InsecureSkipCheck bool `protobuf:"varint,3,opt,name=insecure_skip_check,json=insecureSkipCheck,proto3" json:"insecure_skip_check,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *DeleteRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *DeleteRequest) GetInsecureSkipCheck() bool {
	if x != nil {
		return x.InsecureSkipCheck
	}
	return false
}

// Defines the SaveReply message which is used to return a message after deleting a project.
type DeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message being returned.
	Msg string `protobuf:"bytes,1,opt,name=msg,json=message,proto3" json:"msg,omitempty"`
}

func (x *DeleteReply) Reset() {
	*x = DeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReply) ProtoMessage() {}

func (x *DeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReply.ProtoReflect.Descriptor instead.
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// The request body for the project.
type SaveRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The language used in the project.
	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *SaveRequest_Body) Reset() {
	*x = SaveRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_project_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest_Body) ProtoMessage() {}

func (x *SaveRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_project_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest_Body.ProtoReflect.Descriptor instead.
func (*SaveRequest_Body) Descriptor() ([]byte, []int) {
	return file_project_v1_project_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SaveRequest_Body) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_project_v1_project_proto protoreflect.FileDescriptor

var file_project_v1_project_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x54, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xdf, 0x01,
	0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x22, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0x21, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x23, 0x0a, 0x0b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xb4, 0x05, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0xa4, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x60, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0xba, 0x47, 0x1e, 0x2a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x12, 0x9d, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0xba, 0x47, 0x20, 0x2a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x12, 0xae, 0x01, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x67, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3f, 0x22, 0x37, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0xba, 0x47, 0x1f, 0x2a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xb0, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x63, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x2a, 0x37, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0xba, 0x47, 0x21, 0x2a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x42, 0xcc, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x75, 0x74, 0x65, 0x73, 0x2d, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0xba, 0x47, 0x93, 0x01, 0x12, 0x56, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20,
	0x41, 0x50, 0x49, 0x2a, 0x40, 0x0a, 0x12, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69,
	0x63, 0x65, 0x6e, 0x73, 0x65, 0x20, 0x32, 0x2e, 0x30, 0x12, 0x2a, 0x68, 0x74, 0x74, 0x70, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x2d, 0x32, 0x2e, 0x30, 0x32, 0x05, 0x30, 0x2e, 0x33, 0x2e, 0x30, 0x2a, 0x39, 0x3a, 0x37,
	0x0a, 0x35, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27,
	0x0a, 0x25, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2a,
	0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_v1_project_proto_rawDescOnce sync.Once
	file_project_v1_project_proto_rawDescData = file_project_v1_project_proto_rawDesc
)

func file_project_v1_project_proto_rawDescGZIP() []byte {
	file_project_v1_project_proto_rawDescOnce.Do(func() {
		file_project_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_v1_project_proto_rawDescData)
	})
	return file_project_v1_project_proto_rawDescData
}

var file_project_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_project_v1_project_proto_goTypes = []interface{}{
	(*GetRequest)(nil),       // 0: api.project.v1.GetRequest
	(*GetReply)(nil),         // 1: api.project.v1.GetReply
	(*ListsRequest)(nil),     // 2: api.project.v1.ListsRequest
	(*ListsReply)(nil),       // 3: api.project.v1.ListsReply
	(*SaveRequest)(nil),      // 4: api.project.v1.SaveRequest
	(*SaveReply)(nil),        // 5: api.project.v1.SaveReply
	(*DeleteRequest)(nil),    // 6: api.project.v1.DeleteRequest
	(*DeleteReply)(nil),      // 7: api.project.v1.DeleteReply
	(*SaveRequest_Body)(nil), // 8: api.project.v1.SaveRequest.Body
}
var file_project_v1_project_proto_depIdxs = []int32{
	1, // 0: api.project.v1.ListsReply.items:type_name -> api.project.v1.GetReply
	8, // 1: api.project.v1.SaveRequest.body:type_name -> api.project.v1.SaveRequest.Body
	0, // 2: api.project.v1.Project.GetProject:input_type -> api.project.v1.GetRequest
	2, // 3: api.project.v1.Project.ListProjects:input_type -> api.project.v1.ListsRequest
	4, // 4: api.project.v1.Project.SaveProject:input_type -> api.project.v1.SaveRequest
	6, // 5: api.project.v1.Project.DeleteProject:input_type -> api.project.v1.DeleteRequest
	1, // 6: api.project.v1.Project.GetProject:output_type -> api.project.v1.GetReply
	3, // 7: api.project.v1.Project.ListProjects:output_type -> api.project.v1.ListsReply
	5, // 8: api.project.v1.Project.SaveProject:output_type -> api.project.v1.SaveReply
	7, // 9: api.project.v1.Project.DeleteProject:output_type -> api.project.v1.DeleteReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_project_v1_project_proto_init() }
func file_project_v1_project_proto_init() {
	if File_project_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_project_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
		file_project_v1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsRequest); i {
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
		file_project_v1_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListsReply); i {
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
		file_project_v1_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_project_v1_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveReply); i {
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
		file_project_v1_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_project_v1_project_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReply); i {
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
		file_project_v1_project_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest_Body); i {
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
			RawDescriptor: file_project_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_v1_project_proto_goTypes,
		DependencyIndexes: file_project_v1_project_proto_depIdxs,
		MessageInfos:      file_project_v1_project_proto_msgTypes,
	}.Build()
	File_project_v1_project_proto = out.File
	file_project_v1_project_proto_rawDesc = nil
	file_project_v1_project_proto_goTypes = nil
	file_project_v1_project_proto_depIdxs = nil
}
