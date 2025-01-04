package command

import "os"

func PwdHandler(args []string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return pwd, nil
}
