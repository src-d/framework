package utils

import . "gopkg.in/check.v1"

func (s *UtilsSuite) TestReverseStringListMap(c *C) {
	m := map[string][]string{
		"A": {"a", "b", "c"},
		"B": {"a", "b"},
		"C": {"a"},
	}

	r := ReverseStringListMap(m)

	c.Assert(r, HasLen, 3)
	c.Assert(r["a"], HasLen, 3)
	c.Assert(r["b"], HasLen, 2)
	c.Assert(r["c"], HasLen, 1)
	c.Assert(r["A"], IsNil)
	c.Assert(r["B"], IsNil)
	c.Assert(r["C"], IsNil)
}
