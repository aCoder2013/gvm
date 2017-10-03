package cp

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	if len(path) > 0 {
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		return &DirEntry{path}
	} else {
		panic("class path shouldn't be null")
	}
}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	if len(className) > 0 {
		bytes, err := ioutil.ReadFile(filepath.Join(entry.absPath, className))
		return bytes, entry, err
	} else {
		return nil, nil, errors.New("class name shouldn't be null or empty")
	}
}

func (self *DirEntry) String() string {
	return self.absPath
}
