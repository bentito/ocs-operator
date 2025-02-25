// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: provider.proto

package providerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// OnboardConsumerRequest holds the required information to validate the consumer and create StorageConsumer
// resource on the StorageProvider cluster
type OnboardConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// onboardingTicket authenticates the storage consumer cluster
	OnboardingTicket string `protobuf:"bytes,1,opt,name=onboardingTicket,proto3" json:"onboardingTicket,omitempty"`
	// consumerName is the name of the consumer that is used to create the storageConsumer resource
	ConsumerName string `protobuf:"bytes,2,opt,name=consumerName,proto3" json:"consumerName,omitempty"`
	// capacity is the desired storage requested by the consumer cluster
	Capacity string `protobuf:"bytes,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *OnboardConsumerRequest) Reset() {
	*x = OnboardConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnboardConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnboardConsumerRequest) ProtoMessage() {}

func (x *OnboardConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnboardConsumerRequest.ProtoReflect.Descriptor instead.
func (*OnboardConsumerRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{0}
}

func (x *OnboardConsumerRequest) GetOnboardingTicket() string {
	if x != nil {
		return x.OnboardingTicket
	}
	return ""
}

func (x *OnboardConsumerRequest) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *OnboardConsumerRequest) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

// OnboardConsumerResponse holds the response for OnboardConsumer API request
type OnboardConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// K8s UID (UUID) of the consumer cluster
	StorageConsumerUUID string `protobuf:"bytes,1,opt,name=storageConsumerUUID,proto3" json:"storageConsumerUUID,omitempty"`
	// grantedCapacity is the storage granted by the provider cluster
	GrantedCapacity string `protobuf:"bytes,2,opt,name=grantedCapacity,proto3" json:"grantedCapacity,omitempty"`
}

func (x *OnboardConsumerResponse) Reset() {
	*x = OnboardConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnboardConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnboardConsumerResponse) ProtoMessage() {}

func (x *OnboardConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnboardConsumerResponse.ProtoReflect.Descriptor instead.
func (*OnboardConsumerResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{1}
}

func (x *OnboardConsumerResponse) GetStorageConsumerUUID() string {
	if x != nil {
		return x.StorageConsumerUUID
	}
	return ""
}

func (x *OnboardConsumerResponse) GetGrantedCapacity() string {
	if x != nil {
		return x.GrantedCapacity
	}
	return ""
}

// StorageConfigRequest holds the information required generate the json config for connecting to storage provider cluster
type StorageConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// K8s UID (UUID) of the consumer cluster
	StorageConsumerUUID string `protobuf:"bytes,1,opt,name=storageConsumerUUID,proto3" json:"storageConsumerUUID,omitempty"`
}

func (x *StorageConfigRequest) Reset() {
	*x = StorageConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageConfigRequest) ProtoMessage() {}

func (x *StorageConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageConfigRequest.ProtoReflect.Descriptor instead.
func (*StorageConfigRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{2}
}

func (x *StorageConfigRequest) GetStorageConsumerUUID() string {
	if x != nil {
		return x.StorageConsumerUUID
	}
	return ""
}

// ExternalResource holds the configuration data of the resources in external storage cluster
type ExternalResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the external storage cluster resource
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Kind of the external storage cluster resource
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// Data contains the contents of the external cluster resource
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ExternalResource) Reset() {
	*x = ExternalResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalResource) ProtoMessage() {}

func (x *ExternalResource) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalResource.ProtoReflect.Descriptor instead.
func (*ExternalResource) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{3}
}

func (x *ExternalResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExternalResource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ExternalResource) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// StorageConfigResponse holds the response for the GetStorageConfig API request
type StorageConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ExternalResource holds the configuration data of external storage cluster
	ExternalResource []*ExternalResource `protobuf:"bytes,1,rep,name=externalResource,proto3" json:"externalResource,omitempty"`
}

func (x *StorageConfigResponse) Reset() {
	*x = StorageConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageConfigResponse) ProtoMessage() {}

func (x *StorageConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageConfigResponse.ProtoReflect.Descriptor instead.
func (*StorageConfigResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{4}
}

func (x *StorageConfigResponse) GetExternalResource() []*ExternalResource {
	if x != nil {
		return x.ExternalResource
	}
	return nil
}

// OffboardConsumerRequest holds the required information to delete the StorageConsumer CR on the storage provider cluster.
type OffboardConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// K8s UID (UUID) of the consumer cluster
	StorageConsumerUUID string `protobuf:"bytes,1,opt,name=storageConsumerUUID,proto3" json:"storageConsumerUUID,omitempty"`
}

func (x *OffboardConsumerRequest) Reset() {
	*x = OffboardConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffboardConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffboardConsumerRequest) ProtoMessage() {}

func (x *OffboardConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffboardConsumerRequest.ProtoReflect.Descriptor instead.
func (*OffboardConsumerRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{5}
}

func (x *OffboardConsumerRequest) GetStorageConsumerUUID() string {
	if x != nil {
		return x.StorageConsumerUUID
	}
	return ""
}

// OffboardConsumerResponse holds the response for the OffboardConsumer API request
type OffboardConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OffboardConsumerResponse) Reset() {
	*x = OffboardConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffboardConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffboardConsumerResponse) ProtoMessage() {}

func (x *OffboardConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffboardConsumerResponse.ProtoReflect.Descriptor instead.
func (*OffboardConsumerResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{6}
}

// UpdateCapacityRequest holds the information required to increase or decrease the block pool size
// on the provider cluster
type UpdateCapacityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// K8s UID (UUID) of the consumer cluster
	StorageConsumerUUID string `protobuf:"bytes,1,opt,name=storageConsumerUUID,proto3" json:"storageConsumerUUID,omitempty"`
	// capacity is the desired storage requested by the consumer cluster
	Capacity string `protobuf:"bytes,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
}

func (x *UpdateCapacityRequest) Reset() {
	*x = UpdateCapacityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCapacityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCapacityRequest) ProtoMessage() {}

func (x *UpdateCapacityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCapacityRequest.ProtoReflect.Descriptor instead.
func (*UpdateCapacityRequest) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCapacityRequest) GetStorageConsumerUUID() string {
	if x != nil {
		return x.StorageConsumerUUID
	}
	return ""
}

func (x *UpdateCapacityRequest) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

// UpdateCapacityResponse holds the response for UpdateCapacity API request
type UpdateCapacityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// grantedCapacity is the storage granted by the provider cluster
	GrantedCapacity string `protobuf:"bytes,2,opt,name=grantedCapacity,proto3" json:"grantedCapacity,omitempty"`
}

func (x *UpdateCapacityResponse) Reset() {
	*x = UpdateCapacityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCapacityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCapacityResponse) ProtoMessage() {}

func (x *UpdateCapacityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCapacityResponse.ProtoReflect.Descriptor instead.
func (*UpdateCapacityResponse) Descriptor() ([]byte, []int) {
	return file_provider_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCapacityResponse) GetGrantedCapacity() string {
	if x != nil {
		return x.GrantedCapacity
	}
	return ""
}

var File_provider_proto protoreflect.FileDescriptor

var file_provider_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x16, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x22, 0x75, 0x0a, 0x17, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x28, 0x0a, 0x0f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x14, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x55, 0x55, 0x49, 0x44, 0x22, 0x4e, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x5f, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x17, 0x4f, 0x66, 0x66, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x55,
	0x49, 0x44, 0x22, 0x1a, 0x0a, 0x18, 0x4f, 0x66, 0x66, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x32, 0xf2, 0x02, 0x0a, 0x0b, 0x4f, 0x43,
	0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x0f, 0x4f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x4f, 0x66,
	0x66, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x66, 0x66,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61,
	0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_proto_rawDescOnce sync.Once
	file_provider_proto_rawDescData = file_provider_proto_rawDesc
)

func file_provider_proto_rawDescGZIP() []byte {
	file_provider_proto_rawDescOnce.Do(func() {
		file_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_proto_rawDescData)
	})
	return file_provider_proto_rawDescData
}

var file_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_provider_proto_goTypes = []interface{}{
	(*OnboardConsumerRequest)(nil),   // 0: provider.OnboardConsumerRequest
	(*OnboardConsumerResponse)(nil),  // 1: provider.OnboardConsumerResponse
	(*StorageConfigRequest)(nil),     // 2: provider.StorageConfigRequest
	(*ExternalResource)(nil),         // 3: provider.ExternalResource
	(*StorageConfigResponse)(nil),    // 4: provider.StorageConfigResponse
	(*OffboardConsumerRequest)(nil),  // 5: provider.OffboardConsumerRequest
	(*OffboardConsumerResponse)(nil), // 6: provider.OffboardConsumerResponse
	(*UpdateCapacityRequest)(nil),    // 7: provider.UpdateCapacityRequest
	(*UpdateCapacityResponse)(nil),   // 8: provider.UpdateCapacityResponse
}
var file_provider_proto_depIdxs = []int32{
	3, // 0: provider.StorageConfigResponse.externalResource:type_name -> provider.ExternalResource
	0, // 1: provider.OCSProvider.OnboardConsumer:input_type -> provider.OnboardConsumerRequest
	2, // 2: provider.OCSProvider.GetStorageConfig:input_type -> provider.StorageConfigRequest
	5, // 3: provider.OCSProvider.OffboardConsumer:input_type -> provider.OffboardConsumerRequest
	7, // 4: provider.OCSProvider.UpdateCapacity:input_type -> provider.UpdateCapacityRequest
	1, // 5: provider.OCSProvider.OnboardConsumer:output_type -> provider.OnboardConsumerResponse
	4, // 6: provider.OCSProvider.GetStorageConfig:output_type -> provider.StorageConfigResponse
	6, // 7: provider.OCSProvider.OffboardConsumer:output_type -> provider.OffboardConsumerResponse
	8, // 8: provider.OCSProvider.UpdateCapacity:output_type -> provider.UpdateCapacityResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_provider_proto_init() }
func file_provider_proto_init() {
	if File_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnboardConsumerRequest); i {
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
		file_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnboardConsumerResponse); i {
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
		file_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageConfigRequest); i {
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
		file_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalResource); i {
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
		file_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageConfigResponse); i {
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
		file_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffboardConsumerRequest); i {
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
		file_provider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffboardConsumerResponse); i {
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
		file_provider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCapacityRequest); i {
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
		file_provider_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCapacityResponse); i {
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
			RawDescriptor: file_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_proto_goTypes,
		DependencyIndexes: file_provider_proto_depIdxs,
		MessageInfos:      file_provider_proto_msgTypes,
	}.Build()
	File_provider_proto = out.File
	file_provider_proto_rawDesc = nil
	file_provider_proto_goTypes = nil
	file_provider_proto_depIdxs = nil
}
