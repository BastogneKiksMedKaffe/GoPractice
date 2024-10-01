// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: bruh.proto

package GRPC_stc

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

type CourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CourseRequest) Reset() {
	*x = CourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseRequest) ProtoMessage() {}

func (x *CourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseRequest.ProtoReflect.Descriptor instead.
func (*CourseRequest) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{0}
}

func (x *CourseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{1}
}

func (x *StudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TeacherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TeacherRequest) Reset() {
	*x = TeacherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherRequest) ProtoMessage() {}

func (x *TeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherRequest.ProtoReflect.Descriptor instead.
func (*TeacherRequest) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{2}
}

func (x *TeacherRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CourseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CourseReply) Reset() {
	*x = CourseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseReply) ProtoMessage() {}

func (x *CourseReply) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseReply.ProtoReflect.Descriptor instead.
func (*CourseReply) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{3}
}

func (x *CourseReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StudentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StudentReply) Reset() {
	*x = StudentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentReply) ProtoMessage() {}

func (x *StudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentReply.ProtoReflect.Descriptor instead.
func (*StudentReply) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{4}
}

func (x *StudentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TeacherReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TeacherReply) Reset() {
	*x = TeacherReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bruh_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherReply) ProtoMessage() {}

func (x *TeacherReply) ProtoReflect() protoreflect.Message {
	mi := &file_bruh_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherReply.ProtoReflect.Descriptor instead.
func (*TeacherReply) Descriptor() ([]byte, []int) {
	return file_bruh_proto_rawDescGZIP(), []int{5}
}

func (x *TeacherReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_bruh_proto protoreflect.FileDescriptor

var file_bruh_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x72, 0x75, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a,
	0x0b, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x99, 0x01, 0x0a, 0x0a, 0x55,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x61, 0x79,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x61, 0x79,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x53, 0x61, 0x79,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x73,
	0x74, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bruh_proto_rawDescOnce sync.Once
	file_bruh_proto_rawDescData = file_bruh_proto_rawDesc
)

func file_bruh_proto_rawDescGZIP() []byte {
	file_bruh_proto_rawDescOnce.Do(func() {
		file_bruh_proto_rawDescData = protoimpl.X.CompressGZIP(file_bruh_proto_rawDescData)
	})
	return file_bruh_proto_rawDescData
}

var file_bruh_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_bruh_proto_goTypes = []any{
	(*CourseRequest)(nil),  // 0: CourseRequest
	(*StudentRequest)(nil), // 1: StudentRequest
	(*TeacherRequest)(nil), // 2: TeacherRequest
	(*CourseReply)(nil),    // 3: CourseReply
	(*StudentReply)(nil),   // 4: StudentReply
	(*TeacherReply)(nil),   // 5: TeacherReply
}
var file_bruh_proto_depIdxs = []int32{
	2, // 0: University.SayTeacher:input_type -> TeacherRequest
	1, // 1: University.SayStudent:input_type -> StudentRequest
	0, // 2: University.SayCourse:input_type -> CourseRequest
	5, // 3: University.SayTeacher:output_type -> TeacherReply
	4, // 4: University.SayStudent:output_type -> StudentReply
	3, // 5: University.SayCourse:output_type -> CourseReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bruh_proto_init() }
func file_bruh_proto_init() {
	if File_bruh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bruh_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CourseRequest); i {
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
		file_bruh_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StudentRequest); i {
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
		file_bruh_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TeacherRequest); i {
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
		file_bruh_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CourseReply); i {
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
		file_bruh_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StudentReply); i {
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
		file_bruh_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TeacherReply); i {
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
			RawDescriptor: file_bruh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bruh_proto_goTypes,
		DependencyIndexes: file_bruh_proto_depIdxs,
		MessageInfos:      file_bruh_proto_msgTypes,
	}.Build()
	File_bruh_proto = out.File
	file_bruh_proto_rawDesc = nil
	file_bruh_proto_goTypes = nil
	file_bruh_proto_depIdxs = nil
}
