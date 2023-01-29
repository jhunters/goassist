package fileutil_test

import (
	"fmt"
	"testing"

	"github.com/jhunters/goassist/fileutil"
)

func TestListFiles(t *testing.T) {

	fileutil.ListFiles(`.`, func(filename, path string, data []byte) {
		fmt.Println(path, filename)
	})
}
