package fs

import (
	"io/ioutil"
	"time"

	. "gopkg.in/check.v1"
)

var _ = Suite(&GitHubCacheSuite{})

type GitHubCacheSuite struct {
	client *GitHubCache
}

func (s *GitHubCacheSuite) SetUpSuite(c *C) {
	client, err := NewGitHubCache(TokenFixture, RepositoryFixture, time.Second*2)
	if err != nil {
		c.Fatal(err)
	}
	for len(client.cache) == 0 {
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 3)
	s.client = client
}

func (s *GitHubCacheSuite) TestGithubClient_Open(c *C) {
	f, err := s.client.Open("CHANGELOG")
	c.Assert(err, IsNil)
	c.Assert(f.GetFilename(), Equals, "CHANGELOG")

	content, err := ioutil.ReadAll(f)
	c.Assert(err, IsNil)
	c.Assert(string(content), Equals, "Initial changelog\n")
}

func (s *GitHubCacheSuite) TestGithubClient_Stat(c *C) {
	f, err := s.client.Stat("CHANGELOG")
	c.Assert(err, IsNil)
	c.Assert(f.Name(), Equals, "CHANGELOG")
	c.Assert(f.Size(), Equals, int64(len([]byte("Initial changelog\n"))))
}
