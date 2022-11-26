// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: notice/wrappers_value.proto

package notice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// 判断是否是默认值或者传零值
	Price         *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	SalePrice     *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`                    // float64
	Memo          *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`                                               // string
	OptionalPrice *int64                  `protobuf:"varint,6,opt,name=optional_price,json=optionalPrice,proto3,oneof" json:"optional_price,omitempty"` // 使用optional
	Info          *Book_Info              `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_wrappers_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_notice_wrappers_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_notice_wrappers_value_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() *wrapperspb.Int64Value {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Book) GetSalePrice() *wrapperspb.DoubleValue {
	if x != nil {
		return x.SalePrice
	}
	return nil
}

func (x *Book) GetMemo() *wrapperspb.StringValue {
	if x != nil {
		return x.Memo
	}
	return nil
}

func (x *Book) GetOptionalPrice() int64 {
	if x != nil && x.OptionalPrice != nil {
		return *x.OptionalPrice
	}
	return 0
}

func (x *Book) GetInfo() *Book_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type Book_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc string `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_wrappers_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_notice_wrappers_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_notice_wrappers_value_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Book_Info) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Book_Info) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_notice_wrappers_value_proto protoreflect.FileDescriptor

var file_notice_wrappers_value_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x2a,
	0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x1a, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x32, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4c,
	0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_wrappers_value_proto_rawDescOnce sync.Once
	file_notice_wrappers_value_proto_rawDescData = file_notice_wrappers_value_proto_rawDesc
)

func file_notice_wrappers_value_proto_rawDescGZIP() []byte {
	file_notice_wrappers_value_proto_rawDescOnce.Do(func() {
		file_notice_wrappers_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_wrappers_value_proto_rawDescData)
	})
	return file_notice_wrappers_value_proto_rawDescData
}

var file_notice_wrappers_value_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notice_wrappers_value_proto_goTypes = []interface{}{
	(*Book)(nil),                   // 0: notice.Book
	(*Book_Info)(nil),              // 1: notice.Book.Info
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(*wrapperspb.DoubleValue)(nil), // 3: google.protobuf.DoubleValue
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
}
var file_notice_wrappers_value_proto_depIdxs = []int32{
	2, // 0: notice.Book.price:type_name -> google.protobuf.Int64Value
	3, // 1: notice.Book.sale_price:type_name -> google.protobuf.DoubleValue
	4, // 2: notice.Book.memo:type_name -> google.protobuf.StringValue
	1, // 3: notice.Book.info:type_name -> notice.Book.Info
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_notice_wrappers_value_proto_init() }
func file_notice_wrappers_value_proto_init() {
	if File_notice_wrappers_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notice_wrappers_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_notice_wrappers_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book_Info); i {
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
	file_notice_wrappers_value_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notice_wrappers_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notice_wrappers_value_proto_goTypes,
		DependencyIndexes: file_notice_wrappers_value_proto_depIdxs,
		MessageInfos:      file_notice_wrappers_value_proto_msgTypes,
	}.Build()
	File_notice_wrappers_value_proto = out.File
	file_notice_wrappers_value_proto_rawDesc = nil
	file_notice_wrappers_value_proto_goTypes = nil
	file_notice_wrappers_value_proto_depIdxs = nil
}
