package main

import (
	"log"
	"fmt"

	"github.com/golang/protobuf/proto"
	"gopractice/rpc/go-protobuf/example"
)

func main() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type: proto.Int32(17),
		Reps: []int64{1, 2, 3},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println("data: ", data)

	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil{
		log.Fatal("unmarshaling error: ", err)
	}

	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q",
			test.GetLabel(), newTest.GetLabel())
	}

}