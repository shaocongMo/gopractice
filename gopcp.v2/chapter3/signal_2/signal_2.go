package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	sendSignal()
}

func sendSignal() {
	var cmds = []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "mysql"),
	}
	results, err := runCmds(cmds)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
}

func runCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("This cmd slice is invalid!")
	}
	var output []byte
	first := true
	for _, cmd := range cmds {
		if cmd == nil || len((*cmd).Args) == 0 {
			return nil, errors.New("Has a invalied cmd!")
		}
		if !first {
			var inputBuf bytes.Buffer
			inputBuf.Write(output)
			cmd.Stdin = &inputBuf
		} else {
			first = false
		}
		var outputBuf bytes.Buffer
		cmd.Stdout = &outputBuf
		if err := cmd.Start(); err != nil {
			return nil, errors.New(fmt.Sprintf("an error: %s occurred by cmd start: %s", err, (*cmd).Args))
		}
		if err := cmd.Wait(); err != nil {
			return nil, errors.New(fmt.Sprintf("an error: %s occurred by cmd wait: %s", err, (*cmd).Args))
		}
		output = outputBuf.Bytes()
	}
	var lines []string
	var readBuf bytes.Buffer
	readBuf.Write(output)
	for {
		line, err := readBuf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
