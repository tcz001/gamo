package knife

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const _port uint = 50000

//Server is the game Server instance
type Server struct {
	*grpc.Server
	port uint
	lis  net.Listener
}

//Init is to initialize a Serve instance
func (s *Server) Init(port ...uint) {
	if port != nil {
		s.port = port[0]
	} else {
		s.port = _port
	}
	s.Server = grpc.NewServer()
}

//Stop is to start a Serve instance
func (s *Server) Stop() error {
	return nil
}

//Start is to start a Serve instance
func (s *Server) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s.lis = lis
	err = s.Serve(s.lis)
}
