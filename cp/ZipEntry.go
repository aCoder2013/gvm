package cp

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//todo:cache read classes
func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if len(className) > 0 {
		readCloser, err := zip.OpenReader(entry.absPath)
		if err != nil {
			panic(err)
		}
		defer readCloser.Close()
		for _, file := range readCloser.File {
			if file.Name == className {
				fileReader, err := file.Open()
				if err != nil {
					return nil, nil, err
				}
				defer fileReader.Close()
				fileBytes, err := ioutil.ReadAll(fileReader)
				return fileBytes, entry, nil
			}
		}
		return nil, nil, errors.New("class not found :" + className)
	} else {
		return nil, nil, errors.New("class name shouldn't be null or empty")
	}
}

func (self *ZipEntry) String() string {
	return self.absPath
}
