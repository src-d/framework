package fs

import (
	"io/ioutil"

	. "gopkg.in/check.v1"
)

const RepositoryFixture = "github.com/src-d/git-fixture"
const TokenFixture = "87cd2d7ee6ccfc14b2f1e579869d89299c654779"

func (s *WritersSuite) TestGithubClient_Open(c *C) {
	client, err := NewGithubClient(TokenFixture, RepositoryFixture)
	c.Assert(err, IsNil)

	f, err := client.Open("CHANGELOG")
	c.Assert(err, IsNil)
	c.Assert(f.GetFilename(), Equals, "CHANGELOG")

	content, err := ioutil.ReadAll(f)
	c.Assert(err, IsNil)
	c.Assert(string(content), Equals, "Initial changelog\n")
}

func (s *WritersSuite) TestGithubClient_Stat(c *C) {
	client, err := NewGithubClient(TokenFixture, RepositoryFixture)
	c.Assert(err, IsNil)

	f, err := client.Stat("CHANGELOG")
	c.Assert(err, IsNil)
	c.Assert(f.Name(), Equals, "CHANGELOG")
	c.Assert(f.Size(), Equals, int64(len([]byte("Initial changelog\n"))))
}
