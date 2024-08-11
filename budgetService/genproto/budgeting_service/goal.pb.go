// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: goal.proto

package budgeting_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Goal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,4,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,5,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt     string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string  `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Goal) Reset() {
	*x = Goal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goal) ProtoMessage() {}

func (x *Goal) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goal.ProtoReflect.Descriptor instead.
func (*Goal) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{0}
}

func (x *Goal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Goal) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Goal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Goal) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *Goal) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *Goal) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Goal) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Goal) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Goal) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GoalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TargetAmount  float32 `protobuf:"fixed32,3,opt,name=target_amount,json=targetAmount,proto3" json:"target_amount,omitempty"`
	CurrentAmount float32 `protobuf:"fixed32,4,opt,name=current_amount,json=currentAmount,proto3" json:"current_amount,omitempty"`
	Deadline      string  `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GoalRequest) Reset() {
	*x = GoalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalRequest) ProtoMessage() {}

func (x *GoalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalRequest.ProtoReflect.Descriptor instead.
func (*GoalRequest) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{1}
}

func (x *GoalRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoalRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoalRequest) GetTargetAmount() float32 {
	if x != nil {
		return x.TargetAmount
	}
	return 0
}

func (x *GoalRequest) GetCurrentAmount() float32 {
	if x != nil {
		return x.CurrentAmount
	}
	return 0
}

func (x *GoalRequest) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *GoalRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Goals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goals []*Goal `protobuf:"bytes,1,rep,name=goals,proto3" json:"goals,omitempty"`
	Count int32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Goals) Reset() {
	*x = Goals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Goals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Goals) ProtoMessage() {}

func (x *Goals) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Goals.ProtoReflect.Descriptor instead.
func (*Goals) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{2}
}

func (x *Goals) GetGoals() []*Goal {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *Goals) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GoalProgressReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goal     *Goal   `protobuf:"bytes,1,opt,name=goal,proto3" json:"goal,omitempty"`
	Progress float32 `protobuf:"fixed32,2,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *GoalProgressReport) Reset() {
	*x = GoalProgressReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalProgressReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalProgressReport) ProtoMessage() {}

func (x *GoalProgressReport) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalProgressReport.ProtoReflect.Descriptor instead.
func (*GoalProgressReport) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{3}
}

func (x *GoalProgressReport) GetGoal() *Goal {
	if x != nil {
		return x.Goal
	}
	return nil
}

func (x *GoalProgressReport) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

type GoalProgressesReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoalProgressesReport []*GoalProgressReport `protobuf:"bytes,1,rep,name=goal_progresses_report,json=goalProgressesReport,proto3" json:"goal_progresses_report,omitempty"`
}

func (x *GoalProgressesReport) Reset() {
	*x = GoalProgressesReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoalProgressesReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoalProgressesReport) ProtoMessage() {}

func (x *GoalProgressesReport) ProtoReflect() protoreflect.Message {
	mi := &file_goal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoalProgressesReport.ProtoReflect.Descriptor instead.
func (*GoalProgressesReport) Descriptor() ([]byte, []int) {
	return file_goal_proto_rawDescGZIP(), []int{4}
}

func (x *GoalProgressesReport) GetGoalProgressesReport() []*GoalProgressReport {
	if x != nil {
		return x.GoalProgressesReport
	}
	return nil
}

var File_goal_proto protoreflect.FileDescriptor

var file_goal_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x02, 0x0a, 0x04, 0x47, 0x6f, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xba, 0x01, 0x0a,
	0x0b, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a, 0x05, 0x47, 0x6f, 0x61,
	0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x12, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x67,
	0x6f, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x47, 0x6f, 0x61, 0x6c,
	0x52, 0x04, 0x67, 0x6f, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x49, 0x0a, 0x16, 0x67, 0x6f,
	0x61, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x6f, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x14, 0x67, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x84, 0x02, 0x0a, 0x0b, 0x47, 0x6f, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x12, 0x0c, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x05, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x05, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x1a, 0x05, 0x2e,
	0x47, 0x6f, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x12,
	0x0b, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x05, 0x2e, 0x47,
	0x6f, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f,
	0x61, 0x6c, 0x73, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x3b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x15, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_goal_proto_rawDescOnce sync.Once
	file_goal_proto_rawDescData = file_goal_proto_rawDesc
)

func file_goal_proto_rawDescGZIP() []byte {
	file_goal_proto_rawDescOnce.Do(func() {
		file_goal_proto_rawDescData = protoimpl.X.CompressGZIP(file_goal_proto_rawDescData)
	})
	return file_goal_proto_rawDescData
}

var file_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_goal_proto_goTypes = []interface{}{
	(*Goal)(nil),                 // 0: Goal
	(*GoalRequest)(nil),          // 1: GoalRequest
	(*Goals)(nil),                // 2: Goals
	(*GoalProgressReport)(nil),   // 3: GoalProgressReport
	(*GoalProgressesReport)(nil), // 4: GoalProgressesReport
	(*PrimaryKey)(nil),           // 5: PrimaryKey
	(*GetListRequest)(nil),       // 6: GetListRequest
	(*empty.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_goal_proto_depIdxs = []int32{
	0, // 0: Goals.goals:type_name -> Goal
	0, // 1: GoalProgressReport.goal:type_name -> Goal
	3, // 2: GoalProgressesReport.goal_progresses_report:type_name -> GoalProgressReport
	1, // 3: GoalService.CreateGoal:input_type -> GoalRequest
	0, // 4: GoalService.UpdateGoal:input_type -> Goal
	5, // 5: GoalService.GetGoal:input_type -> PrimaryKey
	6, // 6: GoalService.GetListGoals:input_type -> GetListRequest
	5, // 7: GoalService.DeleteGoal:input_type -> PrimaryKey
	5, // 8: GoalService.GetGoalReportProgress:input_type -> PrimaryKey
	0, // 9: GoalService.CreateGoal:output_type -> Goal
	0, // 10: GoalService.UpdateGoal:output_type -> Goal
	0, // 11: GoalService.GetGoal:output_type -> Goal
	2, // 12: GoalService.GetListGoals:output_type -> Goals
	7, // 13: GoalService.DeleteGoal:output_type -> google.protobuf.Empty
	4, // 14: GoalService.GetGoalReportProgress:output_type -> GoalProgressesReport
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_goal_proto_init() }
func file_goal_proto_init() {
	if File_goal_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_goal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goal); i {
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
		file_goal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalRequest); i {
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
		file_goal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Goals); i {
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
		file_goal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalProgressReport); i {
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
		file_goal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoalProgressesReport); i {
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
			RawDescriptor: file_goal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goal_proto_goTypes,
		DependencyIndexes: file_goal_proto_depIdxs,
		MessageInfos:      file_goal_proto_msgTypes,
	}.Build()
	File_goal_proto = out.File
	file_goal_proto_rawDesc = nil
	file_goal_proto_goTypes = nil
	file_goal_proto_depIdxs = nil
}
