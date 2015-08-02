package knife

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type KnifeSuite struct{}

var _ = Suite(&KnifeSuite{})

func (s *KnifeSuite) TestServerInit(c *C) {
	var server Server
	server.Init()
	c.Assert(server.port, Equals, _port)
}

func (s *KnifeSuite) TestServerInitWithPort(c *C) {
	var server Server
	server.Init(100)
	c.Assert(server.port, Equals, uint(100))
}
