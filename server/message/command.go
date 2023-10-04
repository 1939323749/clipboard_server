package message

import (
	"errors"
	"strings"
)

// Parse parses the command and returns the command name and parameters.
// /cmd params
func (c *Command) Parse() ([]byte, error) {
	if c.Raw[:1] != "/" {
		return nil, errors.New("invalid command")
	}
	c.Cmd = strings.Split(c.Raw, " ")[0][1:]
	c.Params = strings.TrimPrefix(c.Raw, "/"+c.Cmd+" ")
	return []byte(c.Cmd), nil
}

func (c *Command) GetType() string {
	return COMMAND
}
