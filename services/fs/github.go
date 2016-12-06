package fs

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/sourcegraph/go-vcsurl"

	"golang.org/x/oauth2"
)

//GithubClient a filesystem on top of Github repositories
type GithubClient struct {
	client *github.Client
	vcs    *vcsurl.RepoInfo
}

func NewGithubClient(token, repositoryURL string) (*GithubClient, error) {
	vcs, err := vcsurl.Parse(repositoryURL)
	if err != nil {
		return nil, err
	}

	tc := oauth2.NewClient(oauth2.NoContext, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	return &GithubClient{
		client: github.NewClient(tc),
		vcs:    vcs,
	}, nil
}

func (c *GithubClient) List() ([]string, error) {
	return nil, ErrNotSupported
}

func (c *GithubClient) Create(filename string) (File, error) {
	return nil, ErrReadOnly
}

func (c *GithubClient) Open(filename string) (File, error) {
	opts := &github.RepositoryContentGetOptions{
		Ref: "master",
	}

	content, _, r, err := c.client.Repositories.GetContents(
		c.vcs.Username, c.vcs.Name, filename, opts,
	)

	if err != nil {
		return nil, err
	}

	if r.Remaining < 100 {
		fmt.Printf("Low Github request level: %d remaining of %d", r.Remaining, r.Limit)
	}

	return NewGithubFile(content), nil
}

func (c *GithubClient) Stat(filename string) (FileInfo, error) {
	f, err := c.Open(filename)
	if err != nil {
		return nil, err
	}

	return &GithubFileInfo{f: f.(*GithubFile)}, nil
}

type GithubFile struct {
	reader *bytes.Buffer
	BaseFile
	size int
}

func NewGithubFile(c *github.RepositoryContent) *GithubFile {
	b, _ := c.Decode()

	file := &GithubFile{}
	file.reader = bytes.NewBuffer(b)
	file.filename = *c.Name
	file.size = len(b)

	return file
}

func (f *GithubFile) Read(p []byte) (int, error) {
	if f.IsClosed() {
		return 0, ErrClosed
	}

	return f.reader.Read(p)
}

func (f *GithubFile) Write(p []byte) (int, error) {
	return -1, ErrReadOnly
}

func (f *GithubFile) Close() error {
	f.closed = true
	return nil
}

func (f *GithubFile) Open() error {
	f.closed = false
	return nil
}

type GithubFileInfo struct {
	f *GithubFile
}

// Name return the file name
func (fi *GithubFileInfo) Name() string {
	return fi.f.GetFilename()
}

// Size return the file size
func (fi *GithubFileInfo) Size() int64 {
	return int64(fi.f.size)
}

// Mode return a null FileMode
func (fi *GithubFileInfo) Mode() os.FileMode {
	return os.FileMode(0)
}

// ModTime return modification time
func (*GithubFileInfo) ModTime() time.Time {
	return time.Now()
}

// IsDir always return false
func (*GithubFileInfo) IsDir() bool {
	return false
}

// Sys return nil always
func (*GithubFileInfo) Sys() interface{} {
	return nil
}
