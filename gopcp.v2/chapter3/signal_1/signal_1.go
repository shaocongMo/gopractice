package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sigRecv := make(chan os.Signal, 1)
	signal.Notify(sigRecv)
	for sig := range sigRecv {
		fmt.Printf("Received a signal: %s\n", sig)
	}
}
