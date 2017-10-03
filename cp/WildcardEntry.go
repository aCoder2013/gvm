package cp

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) *CompositeEntry {
	baseDir := path[:len(path)-1]
	compositeEntrys := make([]Entry, 0, 8)
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntrys = append(compositeEntrys, jarEntry)
		}
		return nil
	})
	return &CompositeEntry{compositeEntrys}
}
