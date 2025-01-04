package command

import (
	"fmt"
	"os"
)

func CdHandler(args []string) (string, error) {
	c := args[0]

	if c == "~" {
		c = os.Getenv("HOME")
	}

	if err := os.Chdir(c); err != nil {
		return "", fmt.Errorf("cd: %s: No such file or directory", c)
	}

	return "", nil
}
