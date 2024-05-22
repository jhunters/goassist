package web_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/jhunters/goassist/web"
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
