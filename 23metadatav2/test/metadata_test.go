package test

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/metadata"
)

/*
@author RandySun
@create 2022-05-15-15:44
*/
func TestMetadata(t *testing.T) {
	md := metadata.New(map[string]string{"go": "programming", "tour": "book"})
	fmt.Printf("%#v, \n%T \n ", md, md)
	fmt.Println(md.Get("go"))
	fmt.Println(md["go"])
}
