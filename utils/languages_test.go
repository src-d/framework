package utils

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type UtilsSuite struct{}

var _ = Suite(&UtilsSuite{})

func (s *UtilsSuite) TestLanguages_Add(c *C) {
	l := make(Languages, 0)

	l.Add("foo", 10)
	c.Assert(l["foo"], Equals, float64(10))

	l.Add("foo", 10)
	c.Assert(l["foo"], Equals, float64(20))
}

func (s *UtilsSuite) TestLanguages_ContainsAnyOf(c *C) {
	l := Languages{}
	c.Assert(l.ContainsAnyOf(), Equals, false)
	c.Assert(l.ContainsAnyOf("foo"), Equals, false)
	c.Assert(l.ContainsAnyOf("foo", "bar"), Equals, false)

	l = Languages{}
	l.Add("foo", 0)
	c.Assert(l.ContainsAnyOf(), Equals, false)
	c.Assert(l.ContainsAnyOf("foo"), Equals, true)
	c.Assert(l.ContainsAnyOf("foo", "bar"), Equals, true)
	c.Assert(l.ContainsAnyOf("bar"), Equals, false)

	l = Languages{}
	l.Add("foo", 0)
	l.Add("bar", 0)
	c.Assert(l.ContainsAnyOf(), Equals, false)
	c.Assert(l.ContainsAnyOf("foo"), Equals, true)
	c.Assert(l.ContainsAnyOf("foo", "bar"), Equals, true)
	c.Assert(l.ContainsAnyOf("bar"), Equals, true)
	c.Assert(l.ContainsAnyOf("baz"), Equals, false)
	c.Assert(l.ContainsAnyOf("foo", "baz"), Equals, true)
}
