package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func NewSimpleEN(name string, talk Talk) ChatBot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}

func (robot *simpleEN) Name() string {
	return robot.name
}

func (robot *simpleEN) Begin() (hello string, err error) {
	return "Please input your name: ", nil
}

func (robot *simpleEN) Hello(userName string) string {
	return fmt.Sprintf("Hello %s! I'm %s. What Can I do for you?", userName, robot.name)
}

func (robot *simpleEN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		saying = fmt.Sprintf("???")
	case "nothing", "bye":
		saying = fmt.Sprintf("Bye!")
		end = true
	default:
		saying = fmt.Sprintf("I didn't catch you.")
	}
	return
}

func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("robot %s has an error occurred: %s", robot.name, err)
}

func (robot *simpleEN) End() error {
	return nil
}
