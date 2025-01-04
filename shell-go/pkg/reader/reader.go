package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	reader io.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{
		reader,
	}
}

func (r *Reader) ReadLine() (string, error) {
	fmt.Fprint(os.Stdout, "🐓 ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("Error reading input: %v\n", err)
	}

	return input, nil
}
