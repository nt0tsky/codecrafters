package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
)

type ResponseWriter interface {
	SetHeader(key string, value string)
	SetStatus(statusCode int, statusText string) error
	Write([]byte) (int, error)
}

type Handler func(wr ResponseWriter, req *Request, ctx context.Context) error

type Server interface {
	ListenAndServe(addr string) error
	HandleFunc(method string, path string, handler Handler)
}

type ServerOptions struct {
	Directory string
}

type httpServer struct {
	port      int
	routes    []*route
	directory string
}

type route struct {
	method  string
	regex   string
	path    string
	params  []string
	handler Handler
}

func NewHTTPServer(options ServerOptions) (Server, error) {
	if options.Directory != "" && !FileExists(options.Directory) {
		return nil, fmt.Errorf("directory %s does not exist", options.Directory)
	}

	return &httpServer{
		routes:    make([]*route, 0),
		directory: options.Directory,
	}, nil
}

func (s *httpServer) HandleFunc(method string, path string, handler Handler) {
	regex, params := pathToRegex(path)

	s.routes = append(s.routes, &route{
		method:  method,
		path:    path,
		regex:   regex,
		handler: handler,
		params:  params,
	})
}

func (s *httpServer) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to listen on address %v", addr))
	}
	defer l.Close()

	log.Printf("Server is running on address %s\n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go func() {
			if err := s.serve(conn); err != nil {
				fmt.Println("Error serving request:", err)
			}
		}()
	}
}

func (s *httpServer) serve(conn net.Conn) error {
	defer conn.Close()

	parser := NewRequestParser(conn)
	req, err := parser.ParseRequest()
	if err != nil {
		fmt.Printf("Failed to parse request: %v\n", err)
		return err
	}

	log.Printf("Incoming %s request %s from %s (proto: %s)", req.Method, req.Path, conn.RemoteAddr().String(), req.Proto)

	compression := parseAcceptEncoding(req.Headers["Accept-Encoding"])
	wr := NewResponseWriter(conn, compression)

	ctx := context.Background()
	route, params := s.matchRoute(req.Method, req.Path)
	if route == nil {
		wr.SetStatus(404, "Not Found")
		wr.Write([]byte(""))

		return nil
	}

	req.Params = params
	req.Headers["X-Directory"] = s.directory
	if err := route.handler(wr, req, ctx); err != nil {
		wr.SetStatus(500, "Internal Server Error")
		wr.Write([]byte(""))
	}

	return nil
}

func (s *httpServer) matchRoute(method string, path string) (*route, map[string]string) {
	for _, route := range s.routes {
		if route.method != method {
			continue
		}

		re := regexp.MustCompile(route.regex)
		matches := re.FindStringSubmatch(path)
		if len(matches) == 0 {
			continue
		}

		params := make(map[string]string)
		for i, paramName := range route.params {
			params[paramName] = matches[i+1]
		}

		return route, params
	}

	return nil, nil
}
