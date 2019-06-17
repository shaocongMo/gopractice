package main

import(
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	args := Args{1000, 1}
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil{
		log.Fatal(err)
	}
	var reply int
	err = client.Call("Arith.Sum", args, &reply)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Reply: ", reply)
}