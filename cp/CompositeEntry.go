package cp

import (
	"errors"
	"strings"
)

type CompositeEntry struct {
	listOfEntry []Entry
}

func newCompositeEntry(path string) *CompositeEntry {
	listOfEntry := make([]Entry, 0, 8)
	for _, subPath := range strings.Split(path, pathSeparator) {
		entry := newEntry(subPath)
		listOfEntry = append(listOfEntry, entry)
	}
	return &CompositeEntry{listOfEntry: listOfEntry}
}

func (entry *CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	if len(className) > 0 {
		for _, entry := range entry.listOfEntry {
			bytes, e, err := entry.readClass(className)
			if err == nil {
				return bytes, e, nil
			}
		}
		return nil, nil, errors.New("class not found :" + className)
	} else {
		return nil, nil, errors.New("class name shouldn't be null or empty")
	}
}

func (self *CompositeEntry) String() string {
	strs := make([]string, len(self.listOfEntry))

	for i, entry := range self.listOfEntry {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathSeparator)
}
