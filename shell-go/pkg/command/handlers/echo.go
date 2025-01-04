package command

import (
	"fmt"
	"strings"
)

func EchoHandler(args []string) (string, error) {
	return fmt.Sprintf("%v", strings.Join(args, " ")), nil
}
