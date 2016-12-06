package email

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type EmailsSuite struct{}

var _ = Suite(&EmailsSuite{})

func (s *EmailsSuite) TestValidEmailsAreNotIgnored(c *C) {
	c.Assert(IsIgnored("foo@qux.com"), Equals, false)
	c.Assert(IsIgnored("foo+bar@qux.com"), Equals, false)
	c.Assert(IsIgnored("foo+bar@gmail.com"), Equals, false)
}

func (s *EmailsSuite) TestPlaceholderDomainsAreIgnored(c *C) {
	c.Assert(IsIgnored("john@example.com"), Equals, true)
	c.Assert(IsIgnored("john@test.com"), Equals, true)
}

func (s *EmailsSuite) TestSingleLabelDomainsAreIgnored(c *C) {
	c.Assert(IsIgnored("john@example"), Equals, true)
	c.Assert(IsIgnored("john@localhost"), Equals, true)
	c.Assert(IsIgnored("john@mymachine"), Equals, true)
}

func (s *EmailsSuite) TestLocalDomainsAreIgnored(c *C) {
	c.Assert(IsIgnored("john@my.local"), Equals, true)
	c.Assert(IsIgnored("john@localhost.localdomain"), Equals, true)
}

func (s *EmailsSuite) TestNonIndividualEmailsAreIgnored(c *C) {
	c.Assert(IsIgnored("nobody@android.com"), Equals, true)
	c.Assert(IsIgnored("badger@gitter.im"), Equals, true)
}

func (s *EmailsSuite) TestMalformedEmailsDoNotCrash(c *C) {
	// Undefined behaviour on malformed emails, but don't panic
	IsIgnored("example@example.com@@@@")
	IsIgnored("example@@@@example.com")
	IsIgnored("example.com")
	IsIgnored("gmail.com")
	IsIgnored("")
}
