package fs

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/github"

	"golang.org/x/oauth2"

	"gopkg.in/inconshreveable/log15.v2"
)

var (
	ErrEmptyEtag = errors.New("empty ETag header")
)

// NewGitHubCache creates a new github client with cached contents that auto updates
func NewGitHubCache(token, url string, check time.Duration) (g *GitHubCache, err error) {
	repo := strings.Split(url, "/")
	if len(repo) != 3 {
		return nil, fmt.Errorf("Invalid url: %q", url)
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic occured: %v", r)
		}
	}()
	client := github.NewClient(oauth2.NewClient(oauth2.NoContext,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})))
	g = &GitHubCache{
		exit:    make(chan struct{}),
		m:       new(sync.RWMutex),
		handler: repo[1],
		repo:    repo[2],
		etag:    `0000000000000000000000000000000000000000`,
		client:  client,
	}
	if err := g.checkUpdate(g.etag); err != nil {
		return nil, err
	}
	go func() {
		tick := time.Tick(check)
		for {
			select {
			case <-tick:
				g.checkUpdate(g.etag)
			case <-g.exit:
				return
			}
		}
	}()
	return
}

type GitHubCache struct {
	exit    chan struct{}
	m       *sync.RWMutex
	client  *github.Client
	handler string
	repo    string
	etag    string
	cache   map[string]*MemoryFile
}

// Close send an exit signal to goroutine avoid memory leaks
func (g *GitHubCache) List() ([]string, error) {
	var list = make([]string, 0, len(g.cache))
	for i := range g.cache {
		list = append(list, i)
	}
	return list, nil
}

// Close send an exit signal to goroutine avoid memory leaks
func (g *GitHubCache) Close() {
	g.exit <- struct{}{}
}

// Create returns an error
func (g *GitHubCache) Create(name string) (File, error) {
	return nil, ErrReadOnly
}

// Open returns a File for reading purposes
func (g *GitHubCache) Open(name string) (File, error) {
	g.m.RLock()
	defer g.m.RUnlock()
	f, ok := g.cache[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return f.clone(), nil
}

// Stat returns file infos
func (g *GitHubCache) Stat(name string) (FileInfo, error) {
	g.m.RLock()
	defer g.m.RUnlock()
	f, ok := g.cache[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return &MemoryFileInfo{f}, nil
}

func (g *GitHubCache) checkUpdate(last string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log15.Warn("GithubCache", "panic", r)
			return
		}
		switch true {
		case err != nil:
			log15.Warn("GithubCache", "error", err)
		case g.etag != last:
			last = g.etag
		}
	}()
	finalUrl, _, err := g.client.Repositories.GetArchiveLink(g.handler, g.repo, github.Zipball, nil)
	if err != nil {
		return err
	}
	req, err := g.client.NewRequest("GET", finalUrl.String(), nil)
	if err != nil {
		return err
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		return fmt.Errorf("Request to %q failed: %s", finalUrl, r.Status)
	}
	etag := strings.Trim(r.Header.Get("ETag"), `"`)
	if etag == "" {
		err = ErrEmptyEtag
		return
	}
	if etag == g.etag {
		return nil
	}
	g.etag = etag
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return g.doUpdate(b)
}

func (g *GitHubCache) doUpdate(b []byte) error {
	g.m.Lock()
	defer g.m.Unlock()
	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return err
	}
	g.cache = make(map[string]*MemoryFile)
	for _, f := range r.File {
		if strings.HasSuffix(f.Name, "/") {
			continue
		}
		n := f.Name[strings.Index(f.Name, "/")+1:]
		r, err := f.Open()
		if err != nil {
			return fmt.Errorf("can't open %q: %s", n, err)
		}
		b, err = ioutil.ReadAll(r)
		if err != nil {
			return fmt.Errorf("can't read %q: %s", n, err)
		}
		g.cache[n] = NewMemoryFile(n, b)
	}
	return nil
}
