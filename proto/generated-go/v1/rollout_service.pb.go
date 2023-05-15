// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: v1/rollout_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the plan to retrieve.
	// Format: projects/{project}/plans/{plan}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPlanRequest) Reset() {
	*x = GetPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanRequest) ProtoMessage() {}

func (x *GetPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanRequest.ProtoReflect.Descriptor instead.
func (*GetPlanRequest) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent, which owns this collection of plans.
	// Format: projects/{project}
	// Use "projects/-" to list all plans from all projects.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of plans to return. The service may return fewer than
	// this value.
	// If unspecified, at most 50 plans will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListPlans` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListPlans` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPlansRequest) Reset() {
	*x = ListPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlansRequest) ProtoMessage() {}

func (x *ListPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlansRequest.ProtoReflect.Descriptor instead.
func (*ListPlansRequest) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPlansRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListPlansRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPlansRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPlansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The plans from the specified request.
	Plans []*Plan `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPlansResponse) Reset() {
	*x = ListPlansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPlansResponse) ProtoMessage() {}

func (x *ListPlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPlansResponse.ProtoReflect.Descriptor instead.
func (*ListPlansResponse) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPlansResponse) GetPlans() []*Plan {
	if x != nil {
		return x.Plans
	}
	return nil
}

func (x *ListPlansResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreatePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent project where this plan will be created.
	// Format: projects/{project}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The plan to create.
	Plan *Plan `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *CreatePlanRequest) Reset() {
	*x = CreatePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanRequest) ProtoMessage() {}

func (x *CreatePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlanRequest.ProtoReflect.Descriptor instead.
func (*CreatePlanRequest) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePlanRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreatePlanRequest) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type UpdatePlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The plan to update.
	//
	// The plan's `name` field is used to identify the plan to update.
	// Format: projects/{project}/plans/{plan}
	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	// The list of fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdatePlanRequest) Reset() {
	*x = UpdatePlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlanRequest) ProtoMessage() {}

func (x *UpdatePlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlanRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlanRequest) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePlanRequest) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *UpdatePlanRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the plan.
	// `plan` is a system generated ID.
	// Format: projects/{project}/plans/{plan}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The system-assigned, unique identifier for a resource.
	Uid         string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rollout_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rollout_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_v1_rollout_service_proto_rawDescGZIP(), []int{5}
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plan) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Plan) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Plan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_v1_rollout_service_proto protoreflect.FileDescriptor

var file_v1_rollout_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x79, 0x74, 0x65,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x64, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x5c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x81, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73,
	0x6b, 0x22, 0x69, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xed, 0x03, 0x0a,
	0x0e, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x67, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x2c, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x7a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70,
	0x6c, 0x61, 0x6e, 0x73, 0x12, 0x6f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x12, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x2e, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x3a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1e, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x43, 0xda, 0x41, 0x10, 0x70, 0x6c, 0x61, 0x6e,
	0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2a, 0x3a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x32, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x42, 0x11, 0x5a, 0x0f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_rollout_service_proto_rawDescOnce sync.Once
	file_v1_rollout_service_proto_rawDescData = file_v1_rollout_service_proto_rawDesc
)

func file_v1_rollout_service_proto_rawDescGZIP() []byte {
	file_v1_rollout_service_proto_rawDescOnce.Do(func() {
		file_v1_rollout_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_rollout_service_proto_rawDescData)
	})
	return file_v1_rollout_service_proto_rawDescData
}

var file_v1_rollout_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_rollout_service_proto_goTypes = []interface{}{
	(*GetPlanRequest)(nil),        // 0: bytebase.v1.GetPlanRequest
	(*ListPlansRequest)(nil),      // 1: bytebase.v1.ListPlansRequest
	(*ListPlansResponse)(nil),     // 2: bytebase.v1.ListPlansResponse
	(*CreatePlanRequest)(nil),     // 3: bytebase.v1.CreatePlanRequest
	(*UpdatePlanRequest)(nil),     // 4: bytebase.v1.UpdatePlanRequest
	(*Plan)(nil),                  // 5: bytebase.v1.Plan
	(*fieldmaskpb.FieldMask)(nil), // 6: google.protobuf.FieldMask
}
var file_v1_rollout_service_proto_depIdxs = []int32{
	5, // 0: bytebase.v1.ListPlansResponse.plans:type_name -> bytebase.v1.Plan
	5, // 1: bytebase.v1.CreatePlanRequest.plan:type_name -> bytebase.v1.Plan
	5, // 2: bytebase.v1.UpdatePlanRequest.plan:type_name -> bytebase.v1.Plan
	6, // 3: bytebase.v1.UpdatePlanRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 4: bytebase.v1.RolloutService.GetPlan:input_type -> bytebase.v1.GetPlanRequest
	1, // 5: bytebase.v1.RolloutService.ListPlans:input_type -> bytebase.v1.ListPlansRequest
	3, // 6: bytebase.v1.RolloutService.CreatePlan:input_type -> bytebase.v1.CreatePlanRequest
	4, // 7: bytebase.v1.RolloutService.UpdatePlan:input_type -> bytebase.v1.UpdatePlanRequest
	5, // 8: bytebase.v1.RolloutService.GetPlan:output_type -> bytebase.v1.Plan
	2, // 9: bytebase.v1.RolloutService.ListPlans:output_type -> bytebase.v1.ListPlansResponse
	5, // 10: bytebase.v1.RolloutService.CreatePlan:output_type -> bytebase.v1.Plan
	5, // 11: bytebase.v1.RolloutService.UpdatePlan:output_type -> bytebase.v1.Plan
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_rollout_service_proto_init() }
func file_v1_rollout_service_proto_init() {
	if File_v1_rollout_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_rollout_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanRequest); i {
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
		file_v1_rollout_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlansRequest); i {
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
		file_v1_rollout_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPlansResponse); i {
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
		file_v1_rollout_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlanRequest); i {
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
		file_v1_rollout_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlanRequest); i {
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
		file_v1_rollout_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
			RawDescriptor: file_v1_rollout_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_rollout_service_proto_goTypes,
		DependencyIndexes: file_v1_rollout_service_proto_depIdxs,
		MessageInfos:      file_v1_rollout_service_proto_msgTypes,
	}.Build()
	File_v1_rollout_service_proto = out.File
	file_v1_rollout_service_proto_rawDesc = nil
	file_v1_rollout_service_proto_goTypes = nil
	file_v1_rollout_service_proto_depIdxs = nil
}