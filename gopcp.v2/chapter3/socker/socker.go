package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:9999"
	DELIMITER      = '\t'
	DEADLINE_SEC   = 10
)

func server() {
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("err: %s", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerLog("listern accept err: %s", err)
			return
		}
		go serverHandle(conn)
	}
}

func serverHandle(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()
	conn.SetDeadline(time.Now().Add(DEADLINE_SEC * time.Second))
	reader := bufio.NewReader(conn)
	rec_bytes, err := receive(reader)
	if err != nil {
		printServerLog("receive err : %s", err)
		return
	}
	num := binary.LittleEndian.Uint32(rec_bytes)
	fmt.Printf("server rec: %d\n", num)
	send_data := make([]byte, 8)
	result := num * 2
	binary.LittleEndian.PutUint64(send_data, uint64(result))
	n, err := send(conn, send_data)
	if err != nil {
		printServerLog("send err : %s", err)
		return
	} else if n > 0 {
		printServerLog("send result : %d", result)
	}
}

func client(id int, num uint32) {
	defer wg.Done()
	conn, err := net.Dial(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printClientLog(id, "err in dial : %s", err)
		return
	}
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(DEADLINE_SEC * time.Second))
	send_data := make([]byte, 4)
	binary.LittleEndian.PutUint32(send_data, num)
	n, err := send(conn, send_data)
	if err != nil {
		printClientLog(id, "send err: %s", err)
		return
	} else if n > 0 {
		printClientLog(id, "send num : %d", num)
	}
	reader := bufio.NewReader(conn)
	rec_bytes, err := receive(reader)
	if err != nil {
		printClientLog(id, "receive err: %s", err)
		return
	}
	re_num := binary.LittleEndian.Uint64(rec_bytes)
	printClientLog(id, "receive num : %d", re_num)
}

func send(conn net.Conn, data []byte) (int, error) {
	return conn.Write(append(data, DELIMITER))
}

func receive(reader *bufio.Reader) ([]byte, error) {
	for {
		readBytes, err := reader.ReadBytes(DELIMITER)
		if err != nil {
			return []byte{}, err
		}
		return readBytes[:len(readBytes)-1], nil
	}
}

func printServerLog(format string, args ...interface{}) {
	printLog("server", 0, format, args)
}

func printClientLog(id int, format string, args ...interface{}) {
	printLog("client", id, format, args)
}

func printLog(Name string, Id int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d] : %s ", Name, Id, fmt.Sprintf(format, args))
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go server()
	time.Sleep(500 * time.Microsecond)
	go client(1, 10)
	wg.Add(1)
	go client(2, 1100)
	wg.Wait()
}
