package web_test

import (
	"fmt"
	"log"
	"net/http"
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
