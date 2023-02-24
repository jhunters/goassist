package client_test

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunters/goassist/web/client"
)

func ExamplePostFile() {

	files := []client.FileUploadInfo{
		{Name: "file", Filepath: `../../testresources/hello.txt`, FileName: "hello_upload.txt"},
	}

	var httpClient = &http.Client{}
	response, err := client.PostFile(httpClient, "http://localhost:8080/FileTest", map[string]string{}, files, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if response != nil {
		content, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		fmt.Println(string(content))
	}

}

func TestPostFile(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/FileTest", upload)

	server := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		time.Sleep(time.Second)

		files := []client.FileUploadInfo{
			{Name: "file", Filepath: `../../testresources/hello.txt`, FileName: "hello_upload.txt"},
		}

		var httpClient = &http.Client{}
		_, err := client.PostFile(httpClient, "http://localhost:8080/FileTest", map[string]string{}, files, map[string]string{})
		if err != nil {
			fmt.Println(err)
			return
		}

		server.Close()

	}()
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("%v", err)
	}

}

func http_resp(code int, msg string, w http.ResponseWriter) {
	Result := make(map[string]interface{})

	Result["code"] = code
	Result["msg"] = msg

	data, err := json.Marshal(Result)
	if err != nil {
		log.Printf("%v\n", err)
	}

	w.Write([]byte(string(data)))
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, nil)
		return
	}

	contentType := r.Header.Get("content-type")
	contentLen := r.ContentLength

	if !strings.Contains(contentType, "multipart/form-data") {
		http_resp(-1001, "The content-type must be multipart/form-data", w)
		return
	}
	//限制最大文件大小
	if contentLen >= 50*1024*1024 {
		http_resp(-1002, "Failed to large,limit 50MB", w)
		return
	}

	err := r.ParseMultipartForm(50 * 1024 * 1024)
	if err != nil {
		http_resp(-1003, "Failed to ParseMultipartForm", w)
		return
	}

	if len(r.MultipartForm.File) == 0 {
		http_resp(-1004, "File is NULL", w)
		return
	}

	Result := make(map[string]interface{})

	Result["code"] = 0

	FileCount := 0
	for _, headers := range r.MultipartForm.File {
		for _, header := range headers {
			log.Printf("recv file: %s\n", header.Filename)
			name := fmt.Sprintf("file%d_url", FileCount)
			Result[name] = (header.Filename)
		}
	}
	data, err := json.Marshal(Result)
	if err != nil {
		log.Printf("%v \n", err)
	}
	w.Write([]byte(string(data)))
}
