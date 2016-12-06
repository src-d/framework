package fs

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"cloud.google.com/go/storage"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/option"
)

type GCSConfig struct {
	Key    []byte
	Scope  string
	Bucket string
}

type GCSClient struct {
	root   string
	config GCSConfig
	ctx    context.Context
	client *storage.Client
	bucket *storage.BucketHandle
}

func NewGCSClient(root string, config GCSConfig) (*GCSClient, error) {
	ctx := context.Background()

	var options []option.ClientOption
	if len(config.Key) > 0 {
		jwt, err := google.JWTConfigFromJSON(
			config.Key,
			config.Scope,
		)
		if err != nil {
			return nil, err
		}

		token := option.WithTokenSource(jwt.TokenSource(ctx))
		options = append(options, token)
	}

	client, err := storage.NewClient(ctx, options...)
	if err != nil {
		return nil, err
	}

	return &GCSClient{
		root:   strings.TrimPrefix(path.Clean(root), "/"),
		config: config,
		ctx:    ctx,
		client: client,
		bucket: client.Bucket(config.Bucket),
	}, nil
}

func (g *GCSClient) Close() error {
	return g.client.Close()
}

func (c *GCSClient) Remove(filename string) error {
	filename = path.Join(c.root, path.Clean(filename))
	obj := c.bucket.Object(filename)

	return obj.Delete(c.ctx)
}

// Create returns a write-only file
func (c *GCSClient) Create(filename string) (File, error) {
	filename = path.Join(c.root, path.Clean(filename))
	obj := c.bucket.Object(filename)

	ow := obj.NewWriter(c.ctx)
	ow.ContentType = "text/plain"
	ow.Metadata = map[string]string{
		"x-goog-meta-srcd": "foo",
	}

	file := &GCSFile{w: ow}
	file.filename = filename

	return file, nil
}

// Create returns a read-write file
func (c *GCSClient) Open(filename string) (File, error) {
	filename = path.Join(c.root, path.Clean(filename))
	obj := c.bucket.Object(filename)

	r, err := obj.NewReader(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("open error %q", err)
	}

	file := &GCSFile{r: r}
	file.filename = filename

	return file, nil
}

func (c *GCSClient) Stat(filename string) (FileInfo, error) {
	filename = path.Join(c.root, path.Clean(filename))
	obj := c.bucket.Object(filename)

	attr, err := obj.Attrs(c.ctx)
	if err != nil {
		return nil, fmt.Errorf("stat error %q", err)
	}

	fi := &GCSFileInfo{}
	fi.filename = filename
	fi.size = attr.Size
	fi.modTime = attr.Updated

	return fi, nil
}

func (c *GCSClient) List() ([]string, error) {
	return nil, ErrNotSupported
}

type GCSFile struct {
	BaseFile
	r *storage.Reader
	w *storage.Writer
}

func (g *GCSFile) Read(p []byte) (n int, err error)  { return g.r.Read(p) }
func (g *GCSFile) Write(p []byte) (n int, err error) { return g.w.Write(p) }

func (g *GCSFile) Close() (err error) {
	g.closed = true
	if g.r != nil {
		return g.r.Close()
	}
	if g.w != nil {
		return g.w.Close()
	}

	return nil
}

type GCSFileInfo struct {
	filename string
	size     int64
	modTime  time.Time
}

func (fi *GCSFileInfo) Name() string       { return fi.filename }
func (fi *GCSFileInfo) Size() int64        { return fi.size }
func (fi *GCSFileInfo) Mode() os.FileMode  { return os.FileMode(0) }
func (fi *GCSFileInfo) ModTime() time.Time { return fi.modTime }
func (fi *GCSFileInfo) IsDir() bool        { return false }
func (fi *GCSFileInfo) Sys() interface{}   { return nil }
