/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-10 18:53:19
 */
package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jhunters/goassist/filex"
	"github.com/yeka/zip"
)

// GZIP do gzip action by gzip package
func GZIP(b []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)
	defer w.Close()

	_, err := w.Write(b)
	w.Flush()

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// GUNZIP do unzip action by gzip package
func GUNZIP(b []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.Write(b)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)

	return undatas, nil
}

// UnzipFile unzip the target zip file to output directory
func UnzipFile(file, password, outputdir string) error {
	r, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.IsEncrypted() {
			f.SetPassword(password)
		}

		r, err := f.Open()
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		r.Close()
		err = ioutil.WriteFile(filepath.Join(outputdir, f.Name), buf, fs.ModeAppend)
		if err != nil {
			return err
		}
	}
	return nil
}

// ZipDir zip the target directory's all file into one zip file
func ZipDir(dir, zipfile, password string) error {
	buf, err := os.Create(zipfile)
	if err != nil {
		return err
	}
	zipw := zip.NewWriter(buf)
	defer zipw.Close()

	filex.ListFiles(dir, func(filename, path string, data []byte) {
		w, err := zipw.Encrypt(filepath.Join(path, filename), password, zip.StandardEncryption)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(w, bytes.NewReader(data))
	})

	zipw.Flush()
	return nil
}
