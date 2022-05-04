package helloproto

import (
	pb "go-grpc-example/01helloproto/proto"
	"io/ioutil"
	"log"
	"testing"

	"google.golang.org/protobuf/proto"
)

/*
@author RandySun
@create 2022-05-04-17:44
*/

//
//  TestPerson
//  @Description: 实例化对象
//  @param t
//
func TestPerson(t *testing.T) {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	t.Logf("id: %d name: %s", p.Id, p.Name)
}

//
//  TestWritePerson
//  @Description: person信息写入grpc_file文件中
//  @param t
//
func TestWritePerson(t *testing.T) {
	book := &pb.AddressBook{
		People: []*pb.Person{
			{
				Id:    1234,
				Name:  "hello word ！！",
				Email: "randy@example.com",
				Phones: []*pb.Person_PhoneNumber{
					{Number: "555-4321", Type: pb.Person_HOME},
				},
			},
		},
	}

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	// 将person写入grpc_file文件中
	if err := ioutil.WriteFile("grpc_file", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

//
//  TestReadPerson
//  @Description: 从grpc_file文件中读取person信息
//  @param t
//
func TestReadPerson(t *testing.T) {
	in, err := ioutil.ReadFile("grpc_file")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	t.Log(book.People[0])

}
