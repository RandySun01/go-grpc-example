package main

import (
	"fmt"
	"go-grpc-example/20protoLevelFiled/proto/notice"
)

/*
@author RandySun
@create 2022-10-30-17:32
*/

// oneOf 使用示例
func oneOfDemo() {
	// client端
	reqEmail := &notice.NoticeReaderRequest{
		Msg: "邮箱通知",
		NoticeWay: &notice.NoticeReaderRequest_Email{
			Email: "RandySun@qq.com",
		},
	}
	//reqPhone := &notice.NoticeReaderRequest{
	//	Msg: "手机通知",
	//	NoticeWay: &notice.NoticeReaderRequest_Phone{
	//		Phone: "157895",
	//	},
	//}
	//fmt.Println(reqEmail)
	//fmt.Println(reqPhone)
	//// server端获取发送类型
	responseEmail := reqEmail

	// 类型断言
	switch v := responseEmail.NoticeWay.(type) {
	case *notice.NoticeReaderRequest_Email:
		noticeWithEmail(v) // 发送邮件通知
	case *notice.NoticeReaderRequest_Phone:
		noticeWithPhone(v) // 发送手机通知
	}
}

// 发送通知相关的功能函数
func noticeWithEmail(in *notice.NoticeReaderRequest_Email) {
	fmt.Printf("notice reader by email:%v\n", in.Email)
}

func noticeWithPhone(in *notice.NoticeReaderRequest_Phone) {
	fmt.Printf("notice reader by phone:%v\n", in.Phone)
}

func main() {
	oneOfDemo()
}
