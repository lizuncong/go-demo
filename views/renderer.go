package views

import (
	"fmt"
	"html/template"
)

// NewRenderer()
//   -> 解析所有页面模板
//   -> 存进 Renderer.pages  { "user/list": template}
//   -> 返回 *Renderer

// renderer.HTML(...)
//   -> 根据 page 名字找模板
//   -> 找不到就报错
//   -> 找到就执行 layout 模板
//   -> 把 HTML 写回浏览器

type Renderer struct {
	pages map[string]*template.Template
}

func NewRenderer() (*Renderer, error) {
	pageFiles := map[string][]string{
		"users/list": {
			"templates/layouts/base.html",
		},
	}
	pages := map[string]*template.Template{}
	for page, files := range pageFiles {
		tmpl, err := template.ParseFiles(files...)
		fmt.Println("parse files....", &tmpl, err)
		if err != nil {
			return nil, err
		}
		pages[page] = tmpl
	}
	return &Renderer{pages}, nil
}
