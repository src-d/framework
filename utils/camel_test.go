package utils

import . "gopkg.in/check.v1"

func (s *UtilsSuite) TestCamelToSnake(c *C) {
	c.Assert(CamelToSnake("Hello"), Equals, "hello")
	c.Assert(CamelToSnake("HelloWorld"), Equals, "hello_world")
	c.Assert(CamelToSnake(""), Equals, "")
}

func (s *UtilsSuite) TestSnakeToCamel(c *C) {
	c.Assert(SnakeToCamel("hello", false), Equals, "hello")
	c.Assert(SnakeToCamel("hello", true), Equals, "Hello")
	c.Assert(SnakeToCamel("hello_world", false), Equals, "helloWorld")
	c.Assert(SnakeToCamel("hello_world", true), Equals, "HelloWorld")
	c.Assert(SnakeToCamel("", false), Equals, "")
	c.Assert(SnakeToCamel("", true), Equals, "")
}
