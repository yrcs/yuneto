// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: v1/hospital.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	pagination "github.com/yrcs/yuneto/third_party/pagination"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddHospitalSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegistrationNumber string  `protobuf:"bytes,3,opt,name=registrationNumber,proto3" json:"registrationNumber,omitempty"`
	ContactPerson      *string `protobuf:"bytes,4,opt,name=contactPerson,proto3,oneof" json:"contactPerson,omitempty"`
	ContactMobile      *string `protobuf:"bytes,5,opt,name=contactMobile,proto3,oneof" json:"contactMobile,omitempty"`
	Locked             *uint32 `protobuf:"varint,6,opt,name=locked,proto3,oneof" json:"locked,omitempty"`
	ApiUrl             string  `protobuf:"bytes,7,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
	SignatureKey       string  `protobuf:"bytes,8,opt,name=signatureKey,proto3" json:"signatureKey,omitempty"`
}

func (x *AddHospitalSettingRequest) Reset() {
	*x = AddHospitalSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHospitalSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHospitalSettingRequest) ProtoMessage() {}

func (x *AddHospitalSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHospitalSettingRequest.ProtoReflect.Descriptor instead.
func (*AddHospitalSettingRequest) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{0}
}

func (x *AddHospitalSettingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddHospitalSettingRequest) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *AddHospitalSettingRequest) GetContactPerson() string {
	if x != nil && x.ContactPerson != nil {
		return *x.ContactPerson
	}
	return ""
}

func (x *AddHospitalSettingRequest) GetContactMobile() string {
	if x != nil && x.ContactMobile != nil {
		return *x.ContactMobile
	}
	return ""
}

func (x *AddHospitalSettingRequest) GetLocked() uint32 {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return 0
}

func (x *AddHospitalSettingRequest) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *AddHospitalSettingRequest) GetSignatureKey() string {
	if x != nil {
		return x.SignatureKey
	}
	return ""
}

type EditHospitalSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name               *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	RegistrationNumber *string `protobuf:"bytes,3,opt,name=registrationNumber,proto3,oneof" json:"registrationNumber,omitempty"`
	ContactPerson      *string `protobuf:"bytes,4,opt,name=contactPerson,proto3,oneof" json:"contactPerson,omitempty"`
	ContactMobile      *string `protobuf:"bytes,5,opt,name=contactMobile,proto3,oneof" json:"contactMobile,omitempty"`
	Locked             *uint32 `protobuf:"varint,6,opt,name=locked,proto3,oneof" json:"locked,omitempty"`
	ApiUrl             *string `protobuf:"bytes,7,opt,name=apiUrl,proto3,oneof" json:"apiUrl,omitempty"`
	SignatureKey       *string `protobuf:"bytes,8,opt,name=signatureKey,proto3,oneof" json:"signatureKey,omitempty"`
}

func (x *EditHospitalSettingRequest) Reset() {
	*x = EditHospitalSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditHospitalSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditHospitalSettingRequest) ProtoMessage() {}

func (x *EditHospitalSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditHospitalSettingRequest.ProtoReflect.Descriptor instead.
func (*EditHospitalSettingRequest) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{1}
}

func (x *EditHospitalSettingRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditHospitalSettingRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *EditHospitalSettingRequest) GetRegistrationNumber() string {
	if x != nil && x.RegistrationNumber != nil {
		return *x.RegistrationNumber
	}
	return ""
}

func (x *EditHospitalSettingRequest) GetContactPerson() string {
	if x != nil && x.ContactPerson != nil {
		return *x.ContactPerson
	}
	return ""
}

func (x *EditHospitalSettingRequest) GetContactMobile() string {
	if x != nil && x.ContactMobile != nil {
		return *x.ContactMobile
	}
	return ""
}

func (x *EditHospitalSettingRequest) GetLocked() uint32 {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return 0
}

func (x *EditHospitalSettingRequest) GetApiUrl() string {
	if x != nil && x.ApiUrl != nil {
		return *x.ApiUrl
	}
	return ""
}

func (x *EditHospitalSettingRequest) GetSignatureKey() string {
	if x != nil && x.SignatureKey != nil {
		return *x.SignatureKey
	}
	return ""
}

type DeleteHospitalSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHospitalSettingRequest) Reset() {
	*x = DeleteHospitalSettingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHospitalSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHospitalSettingRequest) ProtoMessage() {}

func (x *DeleteHospitalSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHospitalSettingRequest.ProtoReflect.Descriptor instead.
func (*DeleteHospitalSettingRequest) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteHospitalSettingRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteHospitalSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteHospitalSettingsRequest) Reset() {
	*x = DeleteHospitalSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHospitalSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHospitalSettingsRequest) ProtoMessage() {}

func (x *DeleteHospitalSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHospitalSettingsRequest.ProtoReflect.Descriptor instead.
func (*DeleteHospitalSettingsRequest) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHospitalSettingsRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type HospitalSettingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Name               string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	RegistrationNumber string                 `protobuf:"bytes,5,opt,name=registrationNumber,proto3" json:"registrationNumber,omitempty"`
	ContactPerson      string                 `protobuf:"bytes,6,opt,name=contactPerson,proto3" json:"contactPerson,omitempty"`
	ContactMobile      string                 `protobuf:"bytes,7,opt,name=contactMobile,proto3" json:"contactMobile,omitempty"`
	Locked             uint32                 `protobuf:"varint,8,opt,name=locked,proto3" json:"locked,omitempty"`
	ApiUrl             string                 `protobuf:"bytes,9,opt,name=apiUrl,proto3" json:"apiUrl,omitempty"`
	SignatureKey       string                 `protobuf:"bytes,10,opt,name=signatureKey,proto3" json:"signatureKey,omitempty"`
}

func (x *HospitalSettingReply) Reset() {
	*x = HospitalSettingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HospitalSettingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HospitalSettingReply) ProtoMessage() {}

func (x *HospitalSettingReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HospitalSettingReply.ProtoReflect.Descriptor instead.
func (*HospitalSettingReply) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{4}
}

func (x *HospitalSettingReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HospitalSettingReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *HospitalSettingReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *HospitalSettingReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HospitalSettingReply) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *HospitalSettingReply) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *HospitalSettingReply) GetContactMobile() string {
	if x != nil {
		return x.ContactMobile
	}
	return ""
}

func (x *HospitalSettingReply) GetLocked() uint32 {
	if x != nil {
		return x.Locked
	}
	return 0
}

func (x *HospitalSettingReply) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *HospitalSettingReply) GetSignatureKey() string {
	if x != nil {
		return x.SignatureKey
	}
	return ""
}

type CommonAddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *CommonAddReply) Reset() {
	*x = CommonAddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonAddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonAddReply) ProtoMessage() {}

func (x *CommonAddReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonAddReply.ProtoReflect.Descriptor instead.
func (*CommonAddReply) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{5}
}

func (x *CommonAddReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommonAddReply) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CommonAddReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CommonEditReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *CommonEditReply) Reset() {
	*x = CommonEditReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_hospital_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonEditReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonEditReply) ProtoMessage() {}

func (x *CommonEditReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hospital_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonEditReply.ProtoReflect.Descriptor instead.
func (*CommonEditReply) Descriptor() ([]byte, []int) {
	return file_v1_hospital_proto_rawDescGZIP(), []int{6}
}

func (x *CommonEditReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommonEditReply) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_v1_hospital_proto protoreflect.FileDescriptor

var file_v1_hospital_proto_rawDesc = []byte{
	0x0a, 0x11, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x06, 0x18, 0x1e, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x39, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x12, 0x18, 0x16, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x06, 0x48, 0x00,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x33, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x98, 0x01, 0x0b, 0x48, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x30, 0x00,
	0x30, 0x01, 0x48, 0x02, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0xdd, 0x03, 0x0a, 0x1a, 0x45,
	0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x06, 0x18, 0x1e, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x12, 0x18, 0x16, 0x48, 0x01, 0x52,
	0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x06, 0x48, 0x02, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x0b, 0x48, 0x03, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x26, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x30, 0x00, 0x30, 0x01, 0x48, 0x04, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x61, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x61, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x0c,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x61, 0x70, 0x69, 0x55, 0x72, 0x6c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x37, 0x0a, 0x1c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x42, 0x0e, 0xfa, 0x42, 0x0b, 0x92, 0x01, 0x08, 0x08, 0x01, 0x22, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xfe, 0x02, 0x0a, 0x14, 0x48, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x70, 0x69, 0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69,
	0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5b,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xc8, 0x03, 0x0a, 0x08,
	0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x49, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x19, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x68, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5c,
	0x0a, 0x13, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5c, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x2a, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x58, 0x0a, 0x1a, 0x64, 0x65, 0x76, 0x2e, 0x79, 0x75,
	0x6e, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x72, 0x63, 0x73, 0x2f, 0x79, 0x75, 0x6e, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_hospital_proto_rawDescOnce sync.Once
	file_v1_hospital_proto_rawDescData = file_v1_hospital_proto_rawDesc
)

func file_v1_hospital_proto_rawDescGZIP() []byte {
	file_v1_hospital_proto_rawDescOnce.Do(func() {
		file_v1_hospital_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_hospital_proto_rawDescData)
	})
	return file_v1_hospital_proto_rawDescData
}

var file_v1_hospital_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_hospital_proto_goTypes = []interface{}{
	(*AddHospitalSettingRequest)(nil),     // 0: hospital.v1.AddHospitalSettingRequest
	(*EditHospitalSettingRequest)(nil),    // 1: hospital.v1.EditHospitalSettingRequest
	(*DeleteHospitalSettingRequest)(nil),  // 2: hospital.v1.DeleteHospitalSettingRequest
	(*DeleteHospitalSettingsRequest)(nil), // 3: hospital.v1.DeleteHospitalSettingsRequest
	(*HospitalSettingReply)(nil),          // 4: hospital.v1.HospitalSettingReply
	(*CommonAddReply)(nil),                // 5: hospital.v1.CommonAddReply
	(*CommonEditReply)(nil),               // 6: hospital.v1.CommonEditReply
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
	(*pagination.PagingRequest)(nil),      // 8: pagination.PagingRequest
	(*pagination.PagingReply)(nil),        // 9: pagination.PagingReply
	(*emptypb.Empty)(nil),                 // 10: google.protobuf.Empty
}
var file_v1_hospital_proto_depIdxs = []int32{
	7,  // 0: hospital.v1.HospitalSettingReply.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 1: hospital.v1.HospitalSettingReply.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: hospital.v1.CommonAddReply.createdAt:type_name -> google.protobuf.Timestamp
	7,  // 3: hospital.v1.CommonAddReply.updatedAt:type_name -> google.protobuf.Timestamp
	7,  // 4: hospital.v1.CommonEditReply.updatedAt:type_name -> google.protobuf.Timestamp
	8,  // 5: hospital.v1.Hospital.ListHospitalSetting:input_type -> pagination.PagingRequest
	0,  // 6: hospital.v1.Hospital.AddHospitalSetting:input_type -> hospital.v1.AddHospitalSettingRequest
	1,  // 7: hospital.v1.Hospital.EditHospitalSetting:input_type -> hospital.v1.EditHospitalSettingRequest
	2,  // 8: hospital.v1.Hospital.DeleteHospitalSetting:input_type -> hospital.v1.DeleteHospitalSettingRequest
	3,  // 9: hospital.v1.Hospital.DeleteHospitalSettings:input_type -> hospital.v1.DeleteHospitalSettingsRequest
	9,  // 10: hospital.v1.Hospital.ListHospitalSetting:output_type -> pagination.PagingReply
	5,  // 11: hospital.v1.Hospital.AddHospitalSetting:output_type -> hospital.v1.CommonAddReply
	6,  // 12: hospital.v1.Hospital.EditHospitalSetting:output_type -> hospital.v1.CommonEditReply
	10, // 13: hospital.v1.Hospital.DeleteHospitalSetting:output_type -> google.protobuf.Empty
	10, // 14: hospital.v1.Hospital.DeleteHospitalSettings:output_type -> google.protobuf.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_v1_hospital_proto_init() }
func file_v1_hospital_proto_init() {
	if File_v1_hospital_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_hospital_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHospitalSettingRequest); i {
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
		file_v1_hospital_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditHospitalSettingRequest); i {
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
		file_v1_hospital_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHospitalSettingRequest); i {
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
		file_v1_hospital_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHospitalSettingsRequest); i {
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
		file_v1_hospital_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HospitalSettingReply); i {
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
		file_v1_hospital_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonAddReply); i {
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
		file_v1_hospital_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonEditReply); i {
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
	file_v1_hospital_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_hospital_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_hospital_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_hospital_proto_goTypes,
		DependencyIndexes: file_v1_hospital_proto_depIdxs,
		MessageInfos:      file_v1_hospital_proto_msgTypes,
	}.Build()
	File_v1_hospital_proto = out.File
	file_v1_hospital_proto_rawDesc = nil
	file_v1_hospital_proto_goTypes = nil
	file_v1_hospital_proto_depIdxs = nil
}
