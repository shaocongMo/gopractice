package chatbot

import (
	"errors"
)

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type ChatBot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	ErrInvalidChatbotName = errors.New("Invalid chatbot name")
	ErrInvalidChatbot     = errors.New("Invalid chatbot")
	ErrExistingChatbot    = errors.New("Existing chatbot")
)

var chatbotMap = map[string]ChatBot{}

func Register(chatbot ChatBot) error {
	if chatbot == nil {
		return ErrInvalidChatbot
	}
	if chatbot.Name() == "" {
		return ErrInvalidChatbotName
	}
	if chatbotMap[chatbot.Name()] != nil {
		return ErrExistingChatbot
	}
	chatbotMap[chatbot.Name()] = chatbot
	return nil
}

func Get(name string) ChatBot {
	return chatbotMap[name]
}
