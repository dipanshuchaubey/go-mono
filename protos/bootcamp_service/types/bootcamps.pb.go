// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: protos/bootcamp_service/types/bootcamps.proto

package types

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

type BootcampInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BootcampId  string   `protobuf:"bytes,1,opt,name=bootcamp_id,json=bootcampId,proto3" json:"bootcamp_id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Website     string   `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	NameSlug    string   `protobuf:"bytes,6,opt,name=name_slug,json=nameSlug,proto3" json:"name_slug,omitempty"`
	Careers     []string `protobuf:"bytes,7,rep,name=careers,proto3" json:"careers,omitempty"`
}

func (x *BootcampInfo) Reset() {
	*x = BootcampInfo{}
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootcampInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootcampInfo) ProtoMessage() {}

func (x *BootcampInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootcampInfo.ProtoReflect.Descriptor instead.
func (*BootcampInfo) Descriptor() ([]byte, []int) {
	return file_protos_bootcamp_service_types_bootcamps_proto_rawDescGZIP(), []int{0}
}

func (x *BootcampInfo) GetBootcampId() string {
	if x != nil {
		return x.BootcampId
	}
	return ""
}

func (x *BootcampInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BootcampInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BootcampInfo) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *BootcampInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BootcampInfo) GetNameSlug() string {
	if x != nil {
		return x.NameSlug
	}
	return ""
}

func (x *BootcampInfo) GetCareers() []string {
	if x != nil {
		return x.Careers
	}
	return nil
}

type CourseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId    string `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CourseInfo) Reset() {
	*x = CourseInfo{}
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CourseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseInfo) ProtoMessage() {}

func (x *CourseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseInfo.ProtoReflect.Descriptor instead.
func (*CourseInfo) Descriptor() ([]byte, []int) {
	return file_protos_bootcamp_service_types_bootcamps_proto_rawDescGZIP(), []int{1}
}

func (x *CourseInfo) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CourseInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CourseInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewId string `protobuf:"bytes,1,opt,name=review_id,json=reviewId,proto3" json:"review_id,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Message  string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Rating   int32  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bootcamp_service_types_bootcamps_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_protos_bootcamp_service_types_bootcamps_proto_rawDescGZIP(), []int{2}
}

func (x *Review) GetReviewId() string {
	if x != nil {
		return x.ReviewId
	}
	return ""
}

func (x *Review) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Review) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Review) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_protos_bootcamp_service_types_bootcamps_proto protoreflect.FileDescriptor

var file_protos_bootcamp_service_types_bootcamps_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x74,
	0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x73, 0x22, 0x61, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x06,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x63, 0x61, 0x72, 0x74, 0x68, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_bootcamp_service_types_bootcamps_proto_rawDescOnce sync.Once
	file_protos_bootcamp_service_types_bootcamps_proto_rawDescData = file_protos_bootcamp_service_types_bootcamps_proto_rawDesc
)

func file_protos_bootcamp_service_types_bootcamps_proto_rawDescGZIP() []byte {
	file_protos_bootcamp_service_types_bootcamps_proto_rawDescOnce.Do(func() {
		file_protos_bootcamp_service_types_bootcamps_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_bootcamp_service_types_bootcamps_proto_rawDescData)
	})
	return file_protos_bootcamp_service_types_bootcamps_proto_rawDescData
}

var file_protos_bootcamp_service_types_bootcamps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_bootcamp_service_types_bootcamps_proto_goTypes = []any{
	(*BootcampInfo)(nil), // 0: bootcamp_service.types.BootcampInfo
	(*CourseInfo)(nil),   // 1: bootcamp_service.types.CourseInfo
	(*Review)(nil),       // 2: bootcamp_service.types.Review
}
var file_protos_bootcamp_service_types_bootcamps_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_bootcamp_service_types_bootcamps_proto_init() }
func file_protos_bootcamp_service_types_bootcamps_proto_init() {
	if File_protos_bootcamp_service_types_bootcamps_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_bootcamp_service_types_bootcamps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_bootcamp_service_types_bootcamps_proto_goTypes,
		DependencyIndexes: file_protos_bootcamp_service_types_bootcamps_proto_depIdxs,
		MessageInfos:      file_protos_bootcamp_service_types_bootcamps_proto_msgTypes,
	}.Build()
	File_protos_bootcamp_service_types_bootcamps_proto = out.File
	file_protos_bootcamp_service_types_bootcamps_proto_rawDesc = nil
	file_protos_bootcamp_service_types_bootcamps_proto_goTypes = nil
	file_protos_bootcamp_service_types_bootcamps_proto_depIdxs = nil
}