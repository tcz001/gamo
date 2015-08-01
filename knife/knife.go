package knife

import (
	"net"

	"google.golang.org/grpc"
)

//Server is the game Server instance
type Server struct {
	*grpc.Server
	port string
	lis  net.Listener
}

const _port = ":50000"

//Init is to initialize a Serve instance
func (s *Server) Init(port string) {
	if port != "" {
		s.port = port
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
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	s.lis = lis
	return nil
}
