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

// Parse 此函数是模板文件系统(TemplateFS)的一部分。 可以处理embed与普通文件方式。 它接受一个路径参数，还有一个可变参数。
// 第一个if块检查Embbed字段的内容，如果它是true，它将使用文件内容（t.Content）和所提供的模式（patterns）创建一个模板。
// 否则，它将使用fs.Glob构建一个文件列表，其中包含给定模式（pattern）和nonEmbedPath参数指定的路径。 然后，他们使用文件名数组创建一个模板，并返回它
func (t TemplateFS) Parse(nonEmbedPath string, patterns ...string) (*template.Template, []string, error) {
	if t.Embbed {
		templ := template.Must(template.New("").Delims(t.DelimsLeft, t.DelimsRigth).Funcs(t.FuncMap).ParseFS(t.Content, patterns...))
		return templ, nil, nil
	}
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(t.Content, pattern)
		if err != nil {
			return nil, nil, err
		}
		if len(list) == 0 {
			return nil, nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		for _, path := range list {
			vpath := nonEmbedPath + "/" + path
			filenames = append(filenames, vpath)
		}
	}
	templ := template.Must(template.New("").Delims(t.DelimsLeft, t.DelimsRigth).Funcs(t.FuncMap).ParseFiles(filenames...))
	return templ, filenames, nil
}
