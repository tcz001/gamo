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
	server.Init("")
	c.Assert(server.port, Equals, _port)
}

func (s *KnifeSuite) TestServerStart(c *C) {
	var dup Server
	err := dup.Start()
	c.Assert(err, IsNil)
	c.Assert(dup.lis, NotNil)
}

func (s *KnifeSuite) TestServerStartError(c *C) {
	var dup Server
	dup.Init(":invalid port")
	err := dup.Start()
	c.Assert(err, NotNil)
	c.Assert(dup.lis, IsNil)
}
