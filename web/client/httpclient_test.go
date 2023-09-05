package client_test

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhunters/goassist/web"
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

func newLocalListener() net.Listener {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	return l
}

func TestPostFile(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/FileTest", upload)

	listener := newLocalListener()

	server := &http.Server{Handler: mux}

	go func() {
		time.Sleep(time.Second)

		files := []client.FileUploadInfo{
			{Name: "file", Filepath: `../../testresources/hello.txt`, FileName: "hello_upload.txt"},
		}

		var httpClient = &http.Client{}
		url := "http://" + listener.Addr().String()
		_, err := client.PostFile(httpClient, url, map[string]string{}, files, map[string]string{})
		if err != nil {
			fmt.Println(err)
			return
		}

		server.Close()

	}()
	err := server.Serve(listener)
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
			name := fmt.Sprintf("file%d", FileCount)
			Result[name] = header.Filename
			FileCount++
		}
	}
	data, err := json.Marshal(Result)
	if err != nil {
		log.Printf("%v \n", err)
	}
	fmt.Println(string(data))
	w.Write([]byte(string(data)))
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	var httpClient = &http.Client{}
	res, err := client.Get(httpClient, ts.URL, map[string]string{}, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}
	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	res.Body.Close()
	fmt.Printf("%s", greeting)
}

func TestStreamClientRequest(t *testing.T) {
	eventStream := func(r *http.Request, ch chan<- string) {
		// write data
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("data: %s\n", time.Now().String())
			time.Sleep(time.Second)
		}
		close(ch)
	}

	ts := httptest.NewServer(http.HandlerFunc(web.EventStreamHandler(eventStream)))
	defer ts.Close()

	var httpClient = &http.Client{}
	resp, err := client.Get(httpClient, ts.URL, map[string]string{}, map[string]string{})
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("EOF reached")
				return
			}
			panic(err)
		}

		fmt.Printf("Received: %s", line)
	}
}
