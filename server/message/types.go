package message

import (
	"errors"
	"strings"
)

const (
	MESSAGE  = "message"
	COMMAND  = "command"
	RESPONSE = "response"
)

type Message struct {
	Msg string
}

type Command struct {
	Raw    string
	Cmd    string
	Params string
}

type Response struct {
	Msg string
}

type SendAble interface {
	Parse() ([]byte, error)
	GetType() string
}

func NewSendAble(msg string) SendAble {
	if _, err := newCommand(msg); err == nil {
		cmd, err := newCommand(msg)
		if err != nil {
			return nil
		}
		return cmd
	}
	return &Message{msg}
}

func newCommand(msg string) (*Command, error) {
	if msg[:1] != "/" {
		return nil, errors.New("invalid command")
	}
	cmd := strings.Split(msg, " ")[0][1:]
	params := strings.TrimPrefix(msg, "/"+cmd+" ")
	return &Command{msg, cmd, params}, nil
}
