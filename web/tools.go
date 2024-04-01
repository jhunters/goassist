package web

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

// OpenBrowser tries to open the URL in a browser, and returns
// whether it succeed.
func OpenBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func EventStreamHandler(onEvent func(*http.Request, chan<- string)) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		// 判断响应是否支持流
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		// 设置响应头信息
		w.Header().Set(HTTP_HEADER_CONTENT_TYPE, HTTP_HEADER_CONTENT_TYPE_ES)
		w.Header().Set(HTTP_HEADER_CACHE_CONTROL, HTTP_HEADER_CONTENT_Control_NO_CACHE)
		w.Header().Set(HTTP_HEADER_CONNECTION, HTTP_HEADER_KEEPALIVE)
		w.Header().Set(HTTP_HEADER_ACCESS_CONTROL_ALLOW_ORIGIN, "*")

		// 创建字符串通道
		ch := make(chan string)
		// 启动事件处理协程
		go onEvent(r, ch)

		// 循环从通道接收数据并输出到响应中
		for {
			v, ok := <-ch
			if !ok {
				// 如果通道已关闭，则终止流输出
				// if closed channel then break stream output
				return
			}

			// 将接收到的事件输出到响应中
			fmt.Fprintf(w, "%v", v)

			// 刷新缓冲区，将数据立即发送给客户端
			flusher.Flush()
		}
	}

	return f
}
