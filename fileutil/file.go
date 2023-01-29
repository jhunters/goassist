package fileutil

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	COMMON_FILE_MODE = 0644 // common file mode , value as '-rw-r--r--' by unix os
)

// file info watch call back function process during in #ListFiles
type callbackFn func(filename, path string, data []byte)

func ListFiles(dir string, callback callbackFn) (reterr error) {
	fileSystem := os.DirFS(dir)
	return fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(filepath.Join(dir, path))
		if err != nil {
			return err
		}
		callback(path, dir, data)
		return nil
	})

}
