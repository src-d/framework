package container

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

type simpleTestConfig struct {
	MyFieldWithoutDefault     string
	MyFieldWithDefault        string `default:"mydefault"`
	MyFieldFromEnv            string `envconfig:"MY_ENV_VAR"`
	MyFieldFromEnvWithDefault string `envconfig:"MY_WINNING_ENV_VAR" default:"defaulted"`
}

func (c simpleTestConfig) EtcdConfigurable() bool {
	return false
}

func (c simpleTestConfig) EtcdServers() string {
	return ""
}
func (c simpleTestConfig) EtcdEnvironment() string {
	return ""
}

func (s *S) TestInitializesFieldWithoutDefault(c *C) {
	config := &simpleTestConfig{}
	InitConfig(config)
	c.Assert(config.MyFieldWithoutDefault, Equals, "")
}

func (s *S) TestInitializesFieldWithDefault(c *C) {
	config := &simpleTestConfig{}
	InitConfig(config)

	c.Assert(config.MyFieldWithDefault, Equals, "mydefault")
}

func (s *S) TestInitializesFieldFromEnvironment(c *C) {
	expectedString := "my expected string"
	os.Setenv("MY_ENV_VAR", expectedString)
	config := &simpleTestConfig{}
	InitConfig(config)

	c.Assert(config.MyFieldFromEnv, Equals, expectedString)
	os.Unsetenv("MY_ENV_VAR")
}

func (s *S) TestEnvironmentValueWinsEvenIfThereIsAlsoDefault(c *C) {
	expectedString := "my expected string"
	os.Setenv("MY_WINNING_ENV_VAR", expectedString)
	config := &simpleTestConfig{}
	InitConfig(config)

	c.Assert(config.MyFieldFromEnvWithDefault, Equals, expectedString)
	os.Unsetenv("MY_WINNING_ENV_VAR")
}

func (s *S) TestDefaultIsAppliedIfThereIsNoEnvVar(c *C) {
	os.Unsetenv("MY_WINNING_ENV_VAR")
	config := &simpleTestConfig{}
	InitConfig(config)

	c.Assert(config.MyFieldFromEnvWithDefault, Equals, "defaulted")
}
