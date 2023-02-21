package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
)

// template FS mock
type TemplateFS struct {
	Content     embed.FS
	Embbed      bool
	DelimsLeft  string
	DelimsRigth string
	FuncMap     template.FuncMap
}

func (t TemplateFS) Parse(nonEmbedPath string, patterns ...string) (*template.Template, error) {
	if t.Embbed {
		templ := template.Must(template.New("").Delims(t.DelimsLeft, t.DelimsRigth).Funcs(t.FuncMap).ParseFS(t.Content, patterns...))
		return templ, nil
	}
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(t.Content, pattern)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		for _, path := range list {
			vpath := nonEmbedPath + "/" + path
			filenames = append(filenames, vpath)
		}
	}
	templ := template.Must(template.New("").Delims(t.DelimsLeft, t.DelimsRigth).Funcs(t.FuncMap).ParseFiles(filenames...))
	return templ, nil
}
