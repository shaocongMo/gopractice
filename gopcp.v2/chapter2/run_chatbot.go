package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"practice/gopcp.v2/chapter2/chatbot"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "The chatbot's name for dialogue.")
	// just like  chatbotName = "simple.en"
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewSimpleEN("simple.en", nil))
	chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	myChatbot := chatbot.Get(chatbotName)
	if myChatbot == nil {
		err := fmt.Errorf("Fatal error: unsupport chatbot named %s\n", chatbotName)
		checkError(nil, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := myChatbot.Begin()
	checkError(myChatbot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(myChatbot, err, true)
	name := input[:len(input)-1]
	hello := myChatbot.Hello(name)
	fmt.Println(hello)
	for {
		input, err := inputReader.ReadString('\n')
		checkError(myChatbot, err, false)
		input = input[:len(input)-1]
		saying, exit, err := myChatbot.Talk(input)
		fmt.Println(saying)
		if exit {
			myChatbot.End()
			os.Exit(0)
		}
		checkError(myChatbot, err, exit)
	}
}

func checkError(chatbot chatbot.ChatBot, err error, exit bool) bool {
	if err == nil {
		return false
	}
	if chatbot != nil {
		fmt.Println(chatbot.ReportError(err))
	} else {
		fmt.Println(err)
	}
	if exit {
		log.Println("exit")
		debug.PrintStack()
		os.Exit(1)
	}
	return true
}
