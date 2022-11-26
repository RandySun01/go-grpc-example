package main

import (
	"fmt"
	"go-grpc-example/20protoLevelFiled/proto/notice"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

/*
@author RandySun
@create 2022-10-30-17:32
*/

// filedMaskDemo 使用google/protobuf/filed_mask.proto 实现部分更新数据
func fieldMaskDemo() {
	// client端
	paths := []string{"price", "info.name"} // 记录更新的字段路径
	updateReq := notice.UpdateBookRequest{
		Op: "randy",
		Book: &notice.Book{
			Price: &wrapperspb.Int64Value{Value: 88},
			Info: &notice.Book_Info{
				Name: "book name randy",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths}, // 提供那些字段需要更新
	}

	// server端实现部分更新字段
	// 根据路径信息, 获取更新的字段
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(updateReq.UpdateMask, generator.CamelCase)
	// 将数据读取到map[string]interface{}
	fmt.Printf("mask:%#v\n", mask)
	var bookDst = make(map[string]interface{})
	//fieldmask-utils支持读取到结构体等，更多用法可查看文档。
	// 通过修改的字段获取值放到bookDst中
	fieldmask_utils.StructToMap(mask, updateReq.Book, bookDst)
	// do update with bookDst
	fmt.Printf("bookDst:%#v\n", bookDst)
}

func main() {
	fieldMaskDemo()
}
