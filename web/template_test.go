package web_test

import (
	"embed"
	"os"

	"github.com/jhunters/goassist/web"
)

//go:embed web_test/*
var htmlDir embed.FS

func ExampleTemplateFS() {
	// web could use TemplateFS to develop by embed mode or direct mode(files modify aware on running)
	webTemplate := web.TemplateFS{Content: htmlDir, Embbed: false, DelimsLeft: "${{", DelimsRigth: "}}"}
	templ, err := webTemplate.Parse("./", "web_test/*.html")
	if err != nil {
		return
	}

	data := struct {
		Title string
		Body  string
	}{
		Title: "My page",
		Body:  "here is html content",
	}

	templ.ExecuteTemplate(os.Stdout, "web2.html", data)

	// to integrate with gin
	// import "github.com/gin-gonic/gin"
	// router := gin.Default()
	// router.SetHTMLTemplate(templ)

	// router.Run(":8080")

}
