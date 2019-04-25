package chatbot

import (
	"fmt"
	"strings"
)

type simpleCN struct {
	name string
	talk Talk
}

func NewSimpleCN(name string, talk Talk) ChatBot {
	return &simpleCN{
		name: name,
		talk: talk,
	}
}

func (robot *simpleCN) Name() string {
	return robot.name
}

func (robot *simpleCN) Begin() (hello string, err error) {
	return "请输入您的名字：", nil
}

func (robot *simpleCN) Hello(userName string) string {
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("您好 %s! 我叫 %s. 有什么能帮到您的吗？", userName, robot.name)
}

func (robot *simpleCN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		saying = fmt.Sprintf("你在说啥子呦？")
	case "没有", "再见":
		saying = fmt.Sprintf("后会有期!")
		end = true
	default:
		saying = fmt.Sprintf("我理解不了.")
	}
	return
}

func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("聊天机器人 %s 出现了一个错误: %s", robot.name, err)
}

func (robot *simpleCN) End() error {
	return nil
}
