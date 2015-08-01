package knife

import (
	"fmt"
	"net"
)

//Server is the game Server instance
type Server struct {
	port string
	lis  net.Listener
}

const _port = ":50000"

//Init is to initialize a Serve instance
func (s *Server) Init() error {
	lis, err := net.Listen("tcp", _port)
	if err != nil {
		return err
	}
	fmt.Println("listening on ", lis.Addr())
	s.port = _port
	s.lis = lis
	return nil
}
