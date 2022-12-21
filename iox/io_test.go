package iox_test

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 8)

	// buf is used here...
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	// ... reused here also. No need to allocate an extra buffer.
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

func TestWalkDir(t *testing.T) {
	root := "D:\\rivercross"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if d.IsDir() {
			fmt.Println(path, d.Type())
		}
		return nil
	})
}
