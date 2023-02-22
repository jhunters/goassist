package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
