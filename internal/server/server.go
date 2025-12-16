package server

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"

	"github.com/mdnewmandev/httpfromtcp/internal/response"
)

type Server struct {
	listener net.Listener
	closed   atomic.Bool
}

func Serve(port int) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	s := &Server{
		listener: listener,
	}
	go s.listen()

	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true)
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}

func (s *Server) listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if s.closed.Load() {
				return
			} else {
				log.Printf("Error accepting connection: %v", err)
				continue
			}
		}
		go s.handle(conn)	
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	response.WriteStatusLine(conn, response.StatusCode200)
	headers := response.GetDefaultHeaders(0)
	err := response.WriteHeaders(conn, headers)
	if err != nil {
		log.Printf("Error writing headers: %v", err)
	}
}