package http

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

type RequestParser struct {
	reader *bufio.Reader
}

func NewRequestParser(conn net.Conn) *RequestParser {
	return &RequestParser{
		reader: bufio.NewReader(conn),
	}
}

func (rp *RequestParser) ParseRequest() (*Request, error) {
	method, path, proto, err := rp.parseRequestLine()
	if err != nil {
		return nil, fmt.Errorf("failed to parse request line: %w", err)
	}

	headers, err := rp.parseHeaders()
	if err != nil {
		return nil, fmt.Errorf("failed to parse headers: %w", err)
	}

	body, err := rp.parseBody(headers)
	if err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}

	return &Request{
		Method:  method,
		Path:    path,
		Proto:   proto,
		Headers: headers,
		Body:    body,
	}, nil
}

func (rp *RequestParser) parseRequestLine() (method, path, proto string, err error) {
	line, err := rp.reader.ReadString('\n')
	if err != nil {
		return "", "", "", fmt.Errorf("failed to read request line: %w", err)
	}

	parts := strings.Fields(strings.TrimSpace(line))
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("malformed request line: %q", line)
	}

	return parts[0], parts[1], parts[2], nil
}

func (rp *RequestParser) parseHeaders() (map[string]string, error) {
	headers := make(map[string]string)

	for {
		line, err := rp.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to read header line: %w", err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed header line: %q", line)
		}

		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return headers, nil
}

func (rp *RequestParser) parseBody(headers map[string]string) ([]byte, error) {
	contentLengthStr := headers["Content-Length"]
	if contentLengthStr == "" {
		return nil, nil
	}

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return nil, fmt.Errorf("invalid Content-Length: %q", contentLengthStr)
	}

	body := make([]byte, contentLength)
	_, err = io.ReadFull(rp.reader, body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	return body, nil
}
