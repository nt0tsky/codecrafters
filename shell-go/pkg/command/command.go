package command

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/pkg/tokenizer"
)

type Command struct {
	Command string
	Args    []string
}

func NewCommand(fields *tokenizer.Tokenizer) (*Command, error) {
	if fields.Command == "" {
		return nil, fmt.Errorf("invalid command")
	}

	return &Command{
		Command: fields.Command,
		Args:    fields.Args,
	}, nil
}
