// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: notice/filed_mask_value.proto

package notice

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//  UpdateBookRequest 更新书籍的消息
type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 操作人
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// 要更新的书籍信息
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	// 记录要更新的字段
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_filed_mask_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notice_filed_mask_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_notice_filed_mask_value_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateBookRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_notice_filed_mask_value_proto protoreflect.FileDescriptor

var file_notice_filed_mask_value_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x2f, 0x5a, 0x2d, 0x67,
	0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x32,
	0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_filed_mask_value_proto_rawDescOnce sync.Once
	file_notice_filed_mask_value_proto_rawDescData = file_notice_filed_mask_value_proto_rawDesc
)

func file_notice_filed_mask_value_proto_rawDescGZIP() []byte {
	file_notice_filed_mask_value_proto_rawDescOnce.Do(func() {
		file_notice_filed_mask_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_filed_mask_value_proto_rawDescData)
	})
	return file_notice_filed_mask_value_proto_rawDescData
}

var file_notice_filed_mask_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notice_filed_mask_value_proto_goTypes = []interface{}{
	(*UpdateBookRequest)(nil),     // 0: notice.UpdateBookRequest
	(*Book)(nil),                  // 1: notice.Book
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
}
var file_notice_filed_mask_value_proto_depIdxs = []int32{
	1, // 0: notice.UpdateBookRequest.book:type_name -> notice.Book
	2, // 1: notice.UpdateBookRequest.update_mask:type_name -> google.protobuf.FieldMask
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notice_filed_mask_value_proto_init() }
func file_notice_filed_mask_value_proto_init() {
	if File_notice_filed_mask_value_proto != nil {
		return
	}
	file_notice_wrappers_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_notice_filed_mask_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
			RawDescriptor: file_notice_filed_mask_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notice_filed_mask_value_proto_goTypes,
		DependencyIndexes: file_notice_filed_mask_value_proto_depIdxs,
		MessageInfos:      file_notice_filed_mask_value_proto_msgTypes,
	}.Build()
	File_notice_filed_mask_value_proto = out.File
	file_notice_filed_mask_value_proto_rawDesc = nil
	file_notice_filed_mask_value_proto_goTypes = nil
	file_notice_filed_mask_value_proto_depIdxs = nil
}