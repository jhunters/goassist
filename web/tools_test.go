package web_test

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunters/goassist/web"
	"github.com/jhunters/goassist/web/client"
)

func ExampleOpenBrowser() {
	httpAddr := "localhost:8080"
	url := "http://" + httpAddr
	if web.OpenBrowser(url) && waitServer(url) {
		log.Printf("A browser window should open. If not, please visit %s", url)
	} else {
		log.Printf("Please open your web browser and visit %s", url)
	}
	fmt.Println("ok")

	// Ouput:
	// ok
}

// waitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

// ExampleEventStreamHandler 是一个演示函数，用于展示如何使用 web.EventStreamHandler 来处理 Server-Sent Events（服务器发送事件）
func ExampleEventStreamHandler() {
	// 定义一个事件流回调函数， 使用流的方式向客户端发送数据
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
