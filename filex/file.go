package filex

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type callbackFn func(filename, path string, data []byte)

// list all files and directies by recursion way
func ListFiles(dir string, callback callbackFn) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		data, _ := ioutil.ReadFile(dir)
		callback(fi.Name(), dir, data)
		return nil
	}

	f, err := os.Open(dir)
	defer f.Close()

	directry, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range directry {
		if f.IsDir() {
			ListFiles(filepath.Join(dir, f.Name()), callback)
			continue
		}
		data, _ := ioutil.ReadFile(filepath.Join(dir, f.Name()))
		callback(f.Name(), dir, data)
	}
	return nil
}
