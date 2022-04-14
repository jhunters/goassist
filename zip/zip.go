/*
 * @Author: Malin Xie
 * @Description:
 * @Date: 2022-02-10 18:53:19
 */
package zip

import (
	"bytes"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"github.com/yeka/zip"
)

type callbackFn func(filename, path string, data []byte)

func listFiles(dir, path string, callback callbackFn) error {

	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		data, _ := ioutil.ReadFile(dir)
		callback(fi.Name(), path, data)
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
			listFiles(dir+"/"+f.Name(), path+"/"+f.Name(), callback)
			continue
		}

		data, _ := ioutil.ReadFile(dir + "/" + f.Name())
		callback(f.Name(), path, data)
	}
	return nil
}

// Unzip unzip the target zip file to output directory
func Unzip(file, password, outputdir string) error {
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
		err = ioutil.WriteFile(outputdir+"/"+f.Name, buf, fs.ModeAppend)
		if err != nil {
			return err
		}
	}
	return nil
}

// ZipAll zip the target directory's all file into one zip file
func ZipAll(dir, zipfile, password string) error {
	buf, err := os.Create(zipfile)
	if err != nil {
		return err
	}
	zipw := zip.NewWriter(buf)
	defer zipw.Close()

	listFiles(dir, "", func(filename, path string, data []byte) {

		w, err := zipw.Encrypt(path+"/"+filename, password, zip.StandardEncryption)

		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(w, bytes.NewReader(data))
	})

	zipw.Flush()
	return nil
}
