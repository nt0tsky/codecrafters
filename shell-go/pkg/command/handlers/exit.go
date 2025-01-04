package command

import "os"

func ExitHandler(args []string) (string, error) {
	os.Exit(0)
	return "", nil
}
