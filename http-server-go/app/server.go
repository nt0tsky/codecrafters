package main

import (
	"context"
	"flag"
	"github.com/codecrafters-io/http-server-starter-go/pkg/server/http"
	"log"
	"os"
	"strconv"
)

func parseDirectory() string {
	dir := flag.String("directory", "", "The directory to use (e.g., /tmp/)")
	flag.Parse()

	return *dir
}

func main() {
	s, err := http.NewHTTPServer(http.ServerOptions{
		Directory: parseDirectory(),
	})
	if err != nil {
		log.Fatal(err)
	}

	s.HandleFunc("GET", "/", func(wr http.ResponseWriter, req *http.Request, ctx context.Context) error {
		if _, err := wr.Write([]byte("")); err != nil {
			return err
		}

		return nil
	})

	s.HandleFunc("GET", "/echo/:value", func(wr http.ResponseWriter, req *http.Request, ctx context.Context) error {
		if _, err := wr.Write([]byte(req.Params["value"])); err != nil {
			return err
		}

		return nil
	})

	s.HandleFunc("GET", "/user-agent", func(wr http.ResponseWriter, req *http.Request, ctx context.Context) error {
		if _, err := wr.Write([]byte(req.Headers["User-Agent"])); err != nil {
			return err
		}

		return nil
	})

	s.HandleFunc("POST", "/files/:file", func(wr http.ResponseWriter, req *http.Request, ctx context.Context) error {
		filePath := http.GetFilePath(req.Headers, req.Params["file"])

		os.WriteFile(filePath, req.Body, os.ModePerm)

		wr.SetStatus(201, "Created")
		if _, err := wr.Write([]byte("")); err != nil {
			return err
		}

		return nil
	})

	s.HandleFunc("GET", "/files/:file", func(wr http.ResponseWriter, req *http.Request, ctx context.Context) error {
		filePath := http.GetFilePath(req.Headers, req.Params["file"])
		if !http.FileExists(filePath) {
			wr.SetStatus(404, "Not Found")
			if _, err := wr.Write([]byte("")); err != nil {
				return err
			}
		}

		bytes, err := os.ReadFile(filePath)
		if err != nil {
			wr.SetStatus(404, "Not Found")
			if _, err := wr.Write([]byte("")); err != nil {
				return err
			}
		}

		wr.SetHeader("Content-Type", "application/octet-stream")
		wr.SetHeader("Content-Length", strconv.Itoa(len(bytes)))
		if _, err := wr.Write(bytes); err != nil {
			return err
		}

		return nil
	})

	if err := s.ListenAndServe(":4221"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
