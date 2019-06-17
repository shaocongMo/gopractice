package main

import (
	"fmt"
	"net"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	fmt.Println("tcp rpc server start")
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr , err := net.ResolveTCPAddr("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go jsonrpc.ServeConn(conn)
	}
}