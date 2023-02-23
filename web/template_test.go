package web_test

import (
	"bytes"
	"embed"
	"os"
	"testing"

	"github.com/jhunters/goassist/web"
	. "github.com/smartystreets/goconvey/convey"
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

func TestTemplateFS(t *testing.T) {
	Convey("TestTemplateFS", t, func() {
		webTemplate := web.TemplateFS{Content: htmlDir, Embbed: false, DelimsLeft: "${{", DelimsRigth: "}}"}
		templ, err := webTemplate.Parse("./", "web_test/*.html")
		So(err, ShouldBeNil)

		data := struct {
			Title string
			Body  string
		}{
			Title: "My page",
			Body:  "here is html content",
		}

		var b bytes.Buffer
		templ.ExecuteTemplate(&b, "web2.html", data)
		So(b.Bytes(), ShouldNotBeEmpty)
		So(bytes.Contains(b.Bytes(), []byte(data.Title)), ShouldBeTrue)
		So(bytes.Contains(b.Bytes(), []byte(data.Body)), ShouldBeTrue)
	})

	Convey("TestTemplateFS embed", t, func() {
		webTemplate := web.TemplateFS{Content: htmlDir, Embbed: true, DelimsLeft: "${{", DelimsRigth: "}}"}
		templ, err := webTemplate.Parse("./", "web_test/*.html")
		So(err, ShouldBeNil)

		data := struct {
			Title string
			Body  string
		}{
			Title: "My page",
			Body:  "here is html content",
		}

		var b bytes.Buffer
		templ.ExecuteTemplate(&b, "web2.html", data)
		So(b.Bytes(), ShouldNotBeEmpty)
		So(bytes.Contains(b.Bytes(), []byte(data.Title)), ShouldBeTrue)
		So(bytes.Contains(b.Bytes(), []byte(data.Body)), ShouldBeTrue)
	})
}
