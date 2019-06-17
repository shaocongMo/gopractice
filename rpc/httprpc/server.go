package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct{
	A, B int
}

type Arith int

func (this *Arith) Sum(args Args, total *int) error {
	fmt.Println("args: ", args)
	*total = args.A + args.B
	return nil
}

func main() {
	fmt.Println("http rpc server start")
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}