package command

import command "github.com/codecrafters-io/shell-starter-go/pkg/command/handlers"

type Handler func(args []string) (string, error)

var BuiltInCommands = map[string]Handler{
	"cd":   command.CdHandler,
	"pwd":  command.PwdHandler,
	"echo": command.EchoHandler,
	"exit": command.ExitHandler,
}
