package ecosystem

import (
	"testing"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
)

type EcosystemSuite struct {
	session *mgo.Session
	ch      chan bool
}

var _ = Suite(&EcosystemSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *EcosystemSuite) TestEcosystems_Add(c *C) {
	e := make(Ecosystems, 0)

	e.Add("foo", 10)
	c.Assert(e["foo"], Equals, float64(10))

	e.Add("foo", 10)
	c.Assert(e["foo"], Equals, float64(20))
}

func (s *EcosystemSuite) TestEcosystems_SuitableLanguage(c *C) {
	e := make(Ecosystems, 0)
	e.Add(Django, 20)
	e.Add(Tensorflow, 20)

	r := e.SuitableLanguages()
	c.Assert(r, HasLen, 2)
	c.Assert(r["Python"], HasLen, 2)
	c.Assert(r["C++"], HasLen, 1)
}
