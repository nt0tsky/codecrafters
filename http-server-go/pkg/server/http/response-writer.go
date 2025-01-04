package http

import (
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type responseWriter struct {
	conn        net.Conn
	statusCode  int
	statusText  string
	headers     map[string]string
	compression []string
}

func NewResponseWriter(conn net.Conn, compression []string) ResponseWriter {
	return &responseWriter{
		conn:        conn,
		headers:     map[string]string{"Content-Type": "text/plain"},
		compression: compression,
		statusCode:  200,
		statusText:  "OK",
	}
}

func (w *responseWriter) SetHeader(key string, value string) {
	w.headers[key] = value
}

func (w *responseWriter) SetStatus(statusCode int, statusText string) error {
	w.statusCode = statusCode
	w.statusText = statusText

	return nil
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if len(w.compression) > 0 {
		enc, cb, err := w.compress(b)
		if err == nil {
			b = cb
			w.SetHeader("Content-Encoding", enc)
		}
	}

	w.SetHeader("Content-Length", strconv.Itoa(len(b)))
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", w.statusCode, w.statusText)
	if _, err := w.conn.Write([]byte(statusLine)); err != nil {
		return 0, fmt.Errorf("failed to write status line: %w", err)
	}

	for key, value := range w.headers {
		headerLine := fmt.Sprintf("%s: %s\r\n", key, value)
		if _, err := w.conn.Write([]byte(headerLine)); err != nil {
			return 0, fmt.Errorf("failed to write headers: %w", err)
		}
	}

	if _, err := w.conn.Write([]byte("\r\n")); err != nil {
		return 0, fmt.Errorf("failed to write header end: %w", err)
	}

	if _, err := w.conn.Write(b); err != nil {
		return 0, fmt.Errorf("failed to write body: %w", err)
	}

	return len(b), nil
}

func (w *responseWriter) compress(b []byte) (string, []byte, error) {
	compressors := map[string]func([]byte) ([]byte, error){
		"gzip": w.compressGzip,
	}

	for _, c := range w.compression {
		if compressFn, supported := compressors[c]; supported {
			cb, err := compressFn(b)
			return c, cb, err
		}
	}

	return "", nil, errors.New("compression not supported")
}

func (w *responseWriter) compressGzip(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(b); err != nil {
		return nil, fmt.Errorf("gzip compression failed while writing body: %w", err)
	}
	if err := zw.Close(); err != nil {
		return nil, fmt.Errorf("gzip compression failed while closing writer: %w", err)
	}

	return buf.Bytes(), nil
}
