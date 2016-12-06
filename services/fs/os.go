package fs

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// OSClient a filesystem based on OSClient
type OSClient struct {
	RootDir string
}

// NewOSClient returns a new OSClient
func NewOSClient(rootDir string) *OSClient {
	return &OSClient{
		RootDir: rootDir,
	}
}

// Create creates a new GlusterFSFile
func (c *OSClient) Create(filename string) (File, error) {
	fullpath := path.Join(c.RootDir, filename)

	dir := filepath.Dir(fullpath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}
	}

	f, err := os.Create(fullpath)
	if err != nil {
		return nil, err
	}

	return &OSFile{
		BaseFile: BaseFile{filename: fullpath},
		file:     f,
	}, nil
}
func (c *OSClient) List() ([]string, error) {
	l, err := ioutil.ReadDir(c.RootDir)
	if err != nil {
		return nil, err
	}
	var s = make([]string, 0, len(l))
	for _, f := range l {
		s = append(s, f.Name())
	}
	return s, nil
}

func (c *OSClient) Open(filename string) (File, error) {
	fullpath := path.Join(c.RootDir, filename)

	f, err := os.Open(fullpath)
	if err != nil {
		return nil, err
	}

	return &OSFile{
		BaseFile: BaseFile{filename: fullpath},
		file:     f,
	}, nil
}

func (c *OSClient) Stat(filename string) (FileInfo, error) {
	fullpath := path.Join(c.RootDir, filename)

	return os.Stat(fullpath)
}

type OSFile struct {
	file *os.File
	BaseFile
}

func (f *OSFile) Read(p []byte) (int, error) {
	return f.file.Read(p)
}

func (f *OSFile) Write(p []byte) (int, error) {
	return f.file.Write(p)
}

func (f *OSFile) Close() error {
	f.closed = true

	return f.file.Close()
}
