package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	FORM_MULTIPART   = "multipart/form-data"
	FORM_ENCODED     = "application/x-www-form-urlencoded"
	APPLICATION_JSON = "application/json"

	CONTENT_TYPE = "Content-Type"
)

// FileUploadInfo upload file info struct
type FileUploadInfo struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
	// 文件名
	FileName string
}

// Get request to target url http server
func Get(client *http.Client, reqUrl string, reqParams map[string]string, headers map[string]string) (*http.Response, error) {
	urlParams := url.Values{}
	Url, _ := url.Parse(reqUrl)
	for key, val := range reqParams {
		urlParams.Set(key, val)
	}

	Url.RawQuery = urlParams.Encode()
	urlPath := Url.String()

	httpRequest, _ := http.NewRequest(http.MethodGet, urlPath, nil)
	// 添加请求头
	for k, v := range headers {
		httpRequest.Header.Add(k, v)
	}
	// 发送请求
	resp, err := client.Do(httpRequest)
	return resp, err
}

// PostJson to post json format value to http server
func PostJson(client *http.Client, reqUrl string, reqParams map[string]string, headers map[string]string) (*http.Response, error) {
	return doPost(client, reqUrl, reqParams, APPLICATION_JSON, nil, headers)
}

// PostForm to post form field and value to http server
func PostForm(client *http.Client, reqUrl string, reqParams map[string]string, headers map[string]string) (*http.Response, error) {
	return doPost(client, reqUrl, reqParams, FORM_ENCODED, nil, headers)
}

// PostFile to post files to http server
func PostFile(client *http.Client, reqUrl string, reqParams map[string]string, files []FileUploadInfo, headers map[string]string) (*http.Response, error) {
	return doPost(client, reqUrl, reqParams, FORM_MULTIPART, files, headers)
}

func doPost(client *http.Client, reqUrl string, reqParams map[string]string, contentType string, files []FileUploadInfo, headers map[string]string) (*http.Response, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is nil")
	}
	requestBody, realContentType, err := buildRequestReader(reqParams, contentType, files)
	if err != nil {
		return nil, err
	}
	httpRequest, _ := http.NewRequest(http.MethodPost, reqUrl, requestBody)
	// 添加请求头
	httpRequest.Header.Add(CONTENT_TYPE, realContentType)
	for k, v := range headers {
		httpRequest.Header.Add(k, v)
	}
	// 发送请求
	resp, err := client.Do(httpRequest)
	return resp, err
}

func buildRequestReader(reqParams map[string]string, contentType string, files []FileUploadInfo) (io.Reader, string, error) {
	if strings.Contains(contentType, "json") {
		bytesData, _ := json.Marshal(reqParams)
		return bytes.NewReader(bytesData), contentType, nil
	} else if files != nil {
		body := &bytes.Buffer{}
		// 文件写入 body
		writer := multipart.NewWriter(body)
		for _, uploadFile := range files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				return nil, "", err
			}
			fileName := filepath.Base(uploadFile.Filepath)
			if uploadFile.FileName != "" {
				fileName = uploadFile.FileName
			}
			part, err := writer.CreateFormFile(uploadFile.Name, fileName)
			if err != nil {
				return nil, "", err
			}
			io.Copy(part, file)
			file.Close()
		}
		// 其他参数列表写入 body
		for k, v := range reqParams {
			if err := writer.WriteField(k, v); err != nil {
				return nil, "", err
			}
		}
		if err := writer.Close(); err != nil {
			return nil, "", err
		}
		// 上传文件需要自己专用的contentType
		return body, writer.FormDataContentType(), nil
	} else {
		urlValues := url.Values{}
		for key, val := range reqParams {
			urlValues.Set(key, val)
		}
		reqBody := urlValues.Encode()
		return strings.NewReader(reqBody), contentType, nil
	}
}
