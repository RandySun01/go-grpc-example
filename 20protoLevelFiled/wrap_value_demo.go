package main

import (
	"fmt"
	"go-grpc-example/20protoLevelFiled/proto/notice"

	"github.com/golang/protobuf/proto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

/*
@author RandySun
@create 2022-10-30-17:32
*/

// wrapValue 使用google/protobuf/wrappers.proto
func wrapValueDemo() {
	// client端
	book := notice.Book{
		Title: "默认值判断",
		Price: &wrapperspb.Int64Value{Value: 8888},
		Memo:  &wrapperspb.StringValue{Value: "学不死就往死里学"},
	}
	// server端判断
	if book.GetPrice() == nil {
		// 没有给price赋值
	} else {
		// price被赋值了直接使用
		price := book.GetPrice() // 获取到的是结构体
		fmt.Println("price: ", price.GetValue())
	}

	if book.GetMemo() != nil {
		// 有值直接使用
		fmt.Println("memo: ", book.GetMemo().GetValue())
	} else {
		fmt.Println("memo没有被设置值")
	}

	if book.GetSalePrice() != nil {
		// 有值直接使用
		fmt.Println("sale price: ", book.GetSalePrice().GetValue())
	} else {
		fmt.Println("sale price没有被设置值")

	}
}

//
//  optionalDemo
//  @Description: 判断是否传值
//
func optionalDemo() {
	// client 客户端
	book := notice.Book{
		Title:         "optional value",
		OptionalPrice: proto.Int64(8888), // sql.NullInt64或者*int64
	}
	book2 := notice.Book{
		Title: "optional value",
	}

	// server 端判断
	// 判断book.OptionalPrice有没有被赋值
	if book.OptionalPrice == nil {
		// 没有被赋值
		fmt.Println("not price")
	} else {
		fmt.Printf("book with price: %#v\n", book.GetOptionalPrice())
	}

	if book2.OptionalPrice == nil {
		// 没有被赋值
		fmt.Println("not price")
	} else {
		fmt.Printf("book with price: %#v", book.GetOptionalPrice())
	}
}
func main() {
	//wrapValueDemo()
	optionalDemo()
}
