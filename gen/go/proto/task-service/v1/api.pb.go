// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: proto/task-service/v1/api.proto

package v1

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

type TaskInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskInfoRequest) Reset() {
	*x = TaskInfoRequest{}
	mi := &file_proto_task_service_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfoRequest) ProtoMessage() {}

func (x *TaskInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfoRequest.ProtoReflect.Descriptor instead.
func (*TaskInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *TaskInfoRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type TaskInfoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskInfoResponse) Reset() {
	*x = TaskInfoResponse{}
	mi := &file_proto_task_service_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfoResponse) ProtoMessage() {}

func (x *TaskInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfoResponse.ProtoReflect.Descriptor instead.
func (*TaskInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *TaskInfoResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskInfoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TaskInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TaskCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	mi := &file_proto_task_service_v1_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_v1_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_task_service_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *TaskCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type TaskCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskCreateResponse) Reset() {
	*x = TaskCreateResponse{}
	mi := &file_proto_task_service_v1_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateResponse) ProtoMessage() {}

func (x *TaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_service_v1_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_service_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *TaskCreateResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_proto_task_service_v1_api_proto protoreflect.FileDescriptor

var file_proto_task_service_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0x2a, 0x0a, 0x0f, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x49, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x12,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x32, 0x95, 0x01, 0x0a, 0x0b,
	0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x5d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x54, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x13, 0x54, 0x61, 0x73, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_task_service_v1_api_proto_rawDescOnce sync.Once
	file_proto_task_service_v1_api_proto_rawDescData = file_proto_task_service_v1_api_proto_rawDesc
)

func file_proto_task_service_v1_api_proto_rawDescGZIP() []byte {
	file_proto_task_service_v1_api_proto_rawDescOnce.Do(func() {
		file_proto_task_service_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_service_v1_api_proto_rawDescData)
	})
	return file_proto_task_service_v1_api_proto_rawDescData
}

var file_proto_task_service_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_task_service_v1_api_proto_goTypes = []any{
	(*TaskInfoRequest)(nil),    // 0: task.v1.TaskInfoRequest
	(*TaskInfoResponse)(nil),   // 1: task.v1.TaskInfoResponse
	(*TaskCreateRequest)(nil),  // 2: task.v1.TaskCreateRequest
	(*TaskCreateResponse)(nil), // 3: task.v1.TaskCreateResponse
}
var file_proto_task_service_v1_api_proto_depIdxs = []int32{
	0, // 0: task.v1.TaskService.TaskInfo:input_type -> task.v1.TaskInfoRequest
	2, // 1: task.v1.TaskService.TaskCreate:input_type -> task.v1.TaskCreateRequest
	1, // 2: task.v1.TaskService.TaskInfo:output_type -> task.v1.TaskInfoResponse
	3, // 3: task.v1.TaskService.TaskCreate:output_type -> task.v1.TaskCreateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_task_service_v1_api_proto_init() }
func file_proto_task_service_v1_api_proto_init() {
	if File_proto_task_service_v1_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_task_service_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_service_v1_api_proto_goTypes,
		DependencyIndexes: file_proto_task_service_v1_api_proto_depIdxs,
		MessageInfos:      file_proto_task_service_v1_api_proto_msgTypes,
	}.Build()
	File_proto_task_service_v1_api_proto = out.File
	file_proto_task_service_v1_api_proto_rawDesc = nil
	file_proto_task_service_v1_api_proto_goTypes = nil
	file_proto_task_service_v1_api_proto_depIdxs = nil
}
