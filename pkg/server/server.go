package server

import (
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
)

type Server struct {
	ln      net.Listener
	handler func(net.Conn)
}

func NewServer(addr string) *Server {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return &Server{
		ln: lis,
	}
}

func (s *Server) AddConnectionHandler(handler func(net.Conn)) {
	s.handler = handler
}

func (s *Server) Start() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Info(err)
			continue
		}

		_ = ants.Submit(func() {
			s.handler(conn)
		})
	}
}
