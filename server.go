/*
	PHP 그누보드를 Golang 으로 이식하기 위한 프로젝트
*/

package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// TODO -- 디자인은 HTML/CSS/JAVASCRIPT 를 이용할 것.
// TODO -- 웹에서 쉽게 위젯을 설치, 사용할 수 있게 할 것.

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main()  {
	e := echo.New()

	renderer := &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	e.Renderer = renderer

	e.Logger.Fatal(e.Start(":80"))
}
