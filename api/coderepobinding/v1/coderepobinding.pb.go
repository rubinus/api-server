// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: coderepobinding/v1/coderepobinding.proto

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

// Request message for listing items
type ListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the product to list repositories for
	ProductName string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
}

func (x *ListsRequest) Reset() {
	*x = ListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsRequest) ProtoMessage() {}

func (x *ListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[0]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{0}
}

func (x *ListsRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

// Define the GetRequest message, which includes the product_name and coderepo_binding_name fields.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName         string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`                   // The product_name field.
	CoderepoBindingName string `protobuf:"bytes,2,opt,name=coderepo_binding_name,proto3" json:"coderepo_binding_name,omitempty"` // The coderepo_binding_name field.
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[1]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetRequest) GetCoderepoBindingName() string {
	if x != nil {
		return x.CoderepoBindingName
	}
	return ""
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Code repo is authorized to this product or projects under it.
	Product string `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// CodeRepoBinding resource name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// If the project list is empty, it means that the code repo is authorized to the product.
	// If the project list has values, it means that the code repo is authorized to the specified projects.
	Projects []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	// Authorization Permissions, readwrite or readonly.
	Permissions string `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// Authorized Code Repository.
	Coderepo string `protobuf:"bytes,5,opt,name=coderepo,proto3" json:"coderepo,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[2]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{2}
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

func (x *GetReply) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *GetReply) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

func (x *GetReply) GetCoderepo() string {
	if x != nil {
		return x.Coderepo
	}
	return ""
}

// Define the ListsReply message, which includes the repeated items field.
type ListsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*GetReply `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // The items field.
}

func (x *ListsReply) Reset() {
	*x = ListsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListsReply) ProtoMessage() {}

func (x *ListsReply) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[3]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{3}
}

func (x *ListsReply) GetItems() []*GetReply {
	if x != nil {
		return x.Items
	}
	return nil
}

// Define the SaveRequest message, which includes the product_name, coderepo_binding_name, insecure_skip_check, and Body fields.
type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Products to which the resource belongs.
	ProductName string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
	// CodeRepoBinding resource name.
	CoderepoBindingName string `protobuf:"bytes,2,opt,name=coderepo_binding_name,proto3" json:"coderepo_binding_name,omitempty"`
	// Whether to skip global resource detection (not recommended).
	InsecureSkipCheck bool              `protobuf:"varint,3,opt,name=insecure_skip_check,proto3" json:"insecure_skip_check,omitempty"`
	Body              *SaveRequest_Body `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[4]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{4}
}

func (x *SaveRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SaveRequest) GetCoderepoBindingName() string {
	if x != nil {
		return x.CoderepoBindingName
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

// Define the SaveReply message, which includes the msg field.
type SaveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg is a message confirming the save.
	Msg string `protobuf:"bytes,1,opt,name=msg,json=message,proto3" json:"msg,omitempty"` // The msg field.
}

func (x *SaveReply) Reset() {
	*x = SaveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReply) ProtoMessage() {}

func (x *SaveReply) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[5]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{5}
}

func (x *SaveReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// Represents a request to delete a codeRepo manifest.
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName         string `protobuf:"bytes,1,opt,name=product_name,proto3" json:"product_name,omitempty"`
	CoderepoBindingName string `protobuf:"bytes,2,opt,name=coderepo_binding_name,proto3" json:"coderepo_binding_name,omitempty"`
	InsecureSkipCheck   bool   `protobuf:"varint,3,opt,name=insecure_skip_check,proto3" json:"insecure_skip_check,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[6]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *DeleteRequest) GetCoderepoBindingName() string {
	if x != nil {
		return x.CoderepoBindingName
	}
	return ""
}

func (x *DeleteRequest) GetInsecureSkipCheck() bool {
	if x != nil {
		return x.InsecureSkipCheck
	}
	return false
}

// Represents a response to a DeleteRequest message.
type DeleteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg is a message confirming the delete.
	Msg string `protobuf:"bytes,1,opt,name=msg,json=message,proto3" json:"msg,omitempty"`
}

func (x *DeleteReply) Reset() {
	*x = DeleteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReply) ProtoMessage() {}

func (x *DeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[7]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaveRequest_Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Code repo is authorized to this product or projects under it.
	Product string `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// If the project list is empty, it means that the code repo is authorized to the product.
	// If the project list has values, it means that the code repo is authorized to the specified projects.
	Projects []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	// Authorization Permissions, readwrite or readonly.
	Permissions string `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// Authorized Code Repository.
	Coderepo string `protobuf:"bytes,4,opt,name=coderepo,proto3" json:"coderepo,omitempty"`
}

func (x *SaveRequest_Body) Reset() {
	*x = SaveRequest_Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest_Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest_Body) ProtoMessage() {}

func (x *SaveRequest_Body) ProtoReflect() protoreflect.Message {
	mi := &file_coderepobinding_v1_coderepobinding_proto_msgTypes[8]
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
	return file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP(), []int{4, 0}
}

func (x *SaveRequest_Body) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *SaveRequest_Body) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *SaveRequest_Body) GetPermissions() string {
	if x != nil {
		return x.Permissions
	}
	return ""
}

func (x *SaveRequest_Body) GetCoderepo() string {
	if x != nil {
		return x.Coderepo
	}
	return ""
}

var File_coderepobinding_v1_coderepobinding_proto protoreflect.FileDescriptor

var file_coderepobinding_v1_coderepobinding_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x15, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x44, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xf9,
	0x02, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f,
	0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x9f, 0x01, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xfa, 0x42,
	0x17, 0x72, 0x15, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x6f, 0x6e, 0x6c, 0x79, 0x52, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x21, 0x0a, 0x09, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9b, 0x01,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x6e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x23, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0xf8, 0x06, 0x0a, 0x0f, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0xd5, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x12, 0x48, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x72,
	0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0xba, 0x47, 0x26, 0x2a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xc5, 0x01, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x63, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70,
	0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xba, 0x47, 0x28, 0x2a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x00, 0x12, 0xe0, 0x01, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x80, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x50, 0x22, 0x48, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x3a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0xba, 0x47, 0x27,
	0x2a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xe1, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x4a, 0x2a, 0x48, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0xba,
	0x47, 0x29, 0x2a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5a, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x42, 0xe4, 0x01, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x75, 0x74, 0x65,
	0x73, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x62, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0xba, 0x47, 0xa3, 0x01, 0x12,
	0x66, 0x0a, 0x1b, 0x43, 0x6f, 0x64, 0x65, 0x20, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x20, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x41, 0x50, 0x49, 0x2a, 0x40,
	0x0a, 0x12, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x20, 0x32, 0x2e, 0x30, 0x12, 0x2a, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x73, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x2d, 0x32, 0x2e, 0x30,
	0x32, 0x05, 0x30, 0x2e, 0x33, 0x2e, 0x30, 0x2a, 0x39, 0x3a, 0x37, 0x0a, 0x35, 0x0a, 0x0a, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x25, 0x0a, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2a, 0x06, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coderepobinding_v1_coderepobinding_proto_rawDescOnce sync.Once
	file_coderepobinding_v1_coderepobinding_proto_rawDescData = file_coderepobinding_v1_coderepobinding_proto_rawDesc
)

func file_coderepobinding_v1_coderepobinding_proto_rawDescGZIP() []byte {
	file_coderepobinding_v1_coderepobinding_proto_rawDescOnce.Do(func() {
		file_coderepobinding_v1_coderepobinding_proto_rawDescData = protoimpl.X.CompressGZIP(file_coderepobinding_v1_coderepobinding_proto_rawDescData)
	})
	return file_coderepobinding_v1_coderepobinding_proto_rawDescData
}

var file_coderepobinding_v1_coderepobinding_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_coderepobinding_v1_coderepobinding_proto_goTypes = []interface{}{
	(*ListsRequest)(nil),     // 0: api.coderepobinding.v1.ListsRequest
	(*GetRequest)(nil),       // 1: api.coderepobinding.v1.GetRequest
	(*GetReply)(nil),         // 2: api.coderepobinding.v1.GetReply
	(*ListsReply)(nil),       // 3: api.coderepobinding.v1.ListsReply
	(*SaveRequest)(nil),      // 4: api.coderepobinding.v1.SaveRequest
	(*SaveReply)(nil),        // 5: api.coderepobinding.v1.SaveReply
	(*DeleteRequest)(nil),    // 6: api.coderepobinding.v1.DeleteRequest
	(*DeleteReply)(nil),      // 7: api.coderepobinding.v1.DeleteReply
	(*SaveRequest_Body)(nil), // 8: api.coderepobinding.v1.SaveRequest.Body
}
var file_coderepobinding_v1_coderepobinding_proto_depIdxs = []int32{
	2, // 0: api.coderepobinding.v1.ListsReply.items:type_name -> api.coderepobinding.v1.GetReply
	8, // 1: api.coderepobinding.v1.SaveRequest.body:type_name -> api.coderepobinding.v1.SaveRequest.Body
	1, // 2: api.coderepobinding.v1.CodeRepoBinding.GetCodeRepoBinding:input_type -> api.coderepobinding.v1.GetRequest
	0, // 3: api.coderepobinding.v1.CodeRepoBinding.ListCodeRepoBindings:input_type -> api.coderepobinding.v1.ListsRequest
	4, // 4: api.coderepobinding.v1.CodeRepoBinding.SaveCodeRepoBinding:input_type -> api.coderepobinding.v1.SaveRequest
	6, // 5: api.coderepobinding.v1.CodeRepoBinding.DeleteCodeRepoBinding:input_type -> api.coderepobinding.v1.DeleteRequest
	2, // 6: api.coderepobinding.v1.CodeRepoBinding.GetCodeRepoBinding:output_type -> api.coderepobinding.v1.GetReply
	3, // 7: api.coderepobinding.v1.CodeRepoBinding.ListCodeRepoBindings:output_type -> api.coderepobinding.v1.ListsReply
	5, // 8: api.coderepobinding.v1.CodeRepoBinding.SaveCodeRepoBinding:output_type -> api.coderepobinding.v1.SaveReply
	7, // 9: api.coderepobinding.v1.CodeRepoBinding.DeleteCodeRepoBinding:output_type -> api.coderepobinding.v1.DeleteReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coderepobinding_v1_coderepobinding_proto_init() }
func file_coderepobinding_v1_coderepobinding_proto_init() {
	if File_coderepobinding_v1_coderepobinding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_coderepobinding_v1_coderepobinding_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_coderepobinding_v1_coderepobinding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coderepobinding_v1_coderepobinding_proto_goTypes,
		DependencyIndexes: file_coderepobinding_v1_coderepobinding_proto_depIdxs,
		MessageInfos:      file_coderepobinding_v1_coderepobinding_proto_msgTypes,
	}.Build()
	File_coderepobinding_v1_coderepobinding_proto = out.File
	file_coderepobinding_v1_coderepobinding_proto_rawDesc = nil
	file_coderepobinding_v1_coderepobinding_proto_goTypes = nil
	file_coderepobinding_v1_coderepobinding_proto_depIdxs = nil
}
