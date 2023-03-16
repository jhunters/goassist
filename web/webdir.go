package web

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
)

type WebDir struct {
	Prefix      string // relative path prefix read by direct from http filesytem
	EmbedPrefix string // relative path prefix read by embed mode
	Content     embed.FS
	Embbed      bool // flag to determin how to read
}

type webfile struct {
	io.Seeker
	fs.File
}

// Readdir implments Readdir mehtod for http.File
func (*webfile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

// Open implments Open method
func (d WebDir) Open(name string) (http.File, error) {
	if !d.Embbed { // if not embed mode
		ff := http.Dir(d.Prefix)
		f, err := ff.Open(name)
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	dir := d.EmbedPrefix
	if dir == "" {
		dir = "."
	}
	fullName := filepath.Join(dir, filepath.FromSlash(path.Clean(name)))
	f, err := d.Content.Open(fullName) // open from embed.FS
	if err != nil {
		return nil, err
	}
	wf := &webfile{File: f}
	return wf, nil
}
