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
	err := server.Init()
	c.Assert(err, IsNil)
	c.Assert(server.port, Equals, _port)
	c.Assert(server.lis, NotNil)
}
