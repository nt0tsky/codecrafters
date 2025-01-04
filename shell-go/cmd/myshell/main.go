package main

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/pkg/command"
	"github.com/codecrafters-io/shell-starter-go/pkg/reader"
	"github.com/codecrafters-io/shell-starter-go/pkg/tokenizer"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	r := reader.NewReader(os.Stdin)
	cm := command.NewManager(command.BuiltInCommands)

	fmt.Println("        _")
	fmt.Println("      <(o )___")
	fmt.Println("       ( ._> /")
	fmt.Println("        `---'")
	for {
		text, err := r.ReadLine()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		f := tokenizer.NewTokenizer(text)
		c, err := command.NewCommand(f)
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		cm.Run(c.Command, c.Args)
	}
}
