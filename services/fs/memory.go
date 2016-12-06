package fs

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"
)

var MemoryFileNotFound = errors.New("File not found")

//MemoryClient a very convenient client based on memory files, for tests, mocks
//and some other use cases
type MemoryClient struct {
	Files map[string]*MemoryFile
}

//NewMemoryClient returns a new MemoryClient
func NewMemoryClient() *MemoryClient {
	return &MemoryClient{
		Files: make(map[string]*MemoryFile, 0),
	}
}

func NewMemoryFile(name string, content []byte) *MemoryFile {
	return &MemoryFile{
		data:     content,
		BaseFile: BaseFile{filename: name},
		Content:  bytes.NewBuffer(content),
	}
}

func (c *MemoryClient) List() ([]string, error) {
	var list = make([]string, 0, len(c.Files))
	for i := range c.Files {
		list = append(list, i)
	}
	return list, nil
}

//Create creates a new MemoryFile, this file will remain on Files map until be
//closed and purged
func (c *MemoryClient) Create(filename string) (File, error) {
	c.Files[filename] = NewMemoryFile(filename, nil)
	return c.Files[filename], nil
}

func (c *MemoryClient) Open(filename string) (File, error) {
	if f, ok := c.Files[filename]; ok {
		f.Open()
		return f, nil
	}

	return nil, MemoryFileNotFound
}

func (c *MemoryClient) Stat(filename string) (FileInfo, error) {
	if _, ok := c.Files[filename]; !ok {
		return nil, fmt.Errorf(fmt.Sprintf("file not found: %s", filename))
	}

	return &MemoryFileInfo{c.Files[filename]}, nil
}

//Purge remove all files closed from Files map
func (c *MemoryClient) Purge() {
	for filename, f := range c.Files {
		if f.closed {
			delete(c.Files, filename)
		}
	}
}

type MemoryFile struct {
	data    []byte
	Content *bytes.Buffer
	BaseFile
}

//Write appends the contents of p to the File
func (f *MemoryFile) Read(p []byte) (int, error) {
	if f.IsClosed() {
		return 0, ErrClosed
	}

	return f.Content.Read(p)
}

//Write appends the contents of p to the File
func (f *MemoryFile) Write(p []byte) (int, error) {
	if f.IsClosed() {
		return 0, ErrClosed
	}

	b, err := f.Content.Write(p)
	if err != nil {
		return b, err
	}
	f.data = append(f.data, p...)
	return b, nil
}

//Close prevents more writes on this file
func (f *MemoryFile) clone() *MemoryFile {
	n := *f
	n.Content = bytes.NewBuffer(n.data)
	return &n
}

//Close prevents more writes on this file
func (f *MemoryFile) Close() error {
	f.closed = true
	return nil
}

func (f *MemoryFile) Open() error {
	f.closed = false
	return nil
}

type MemoryFileInfo struct {
	f *MemoryFile
}

// Name return the file name
func (fi *MemoryFileInfo) Name() string {
	return fi.f.GetFilename()
}

// Size return the file size
func (fi *MemoryFileInfo) Size() int64 {
	return int64(fi.f.Content.Len())
}

// Mode return a null FileMode
func (fi *MemoryFileInfo) Mode() os.FileMode {
	return os.FileMode(0)
}

// ModTime return modification time
func (*MemoryFileInfo) ModTime() time.Time {
	return time.Now()
}

// IsDir always return false
func (*MemoryFileInfo) IsDir() bool {
	return false
}

// Sys return nil always
func (*MemoryFileInfo) Sys() interface{} {
	return nil
}
