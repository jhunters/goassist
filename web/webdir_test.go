package web_test

import (
	"embed"
	"log"
	"net/http"
	"time"

	"github.com/jhunters/goassist/web"
)

//go:embed web_test/*
var content embed.FS

func ExampleWebDir() {
	// web could use WebDir to develop by embed mode or direct mode(files modify aware on running)
	webdir := web.WebDir{Prefix: "./web_test", EmbedPrefix: "./web_test", Content: content, Embbed: true} // embed mode files modify not aware on runing

	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(webdir)) // visit http://localhost:8080/web.html
	server := &http.Server{Addr: ":8080", Handler: mutex}
	go func() {
		time.Sleep(5 * time.Second)
		server.Close()
	}()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// to integrate with gin
	// import "github.com/gin-gonic/gin"
	// router := gin.Default()
	// router.StaticFS("/", webdir)

	// router.Run(":8
}
