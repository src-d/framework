package configurable

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&S{})

type advancedConfiguration struct {
	BasicConfiguration
	Value string `default:"hola"`
}

func (s *S) TestStructWithBasicConfigurationEmbededCanBeInitialized(c *C) {
	config := &advancedConfiguration{}

	InitConfig(config)

	c.Assert(config.Value, Equals, "hola")
}
