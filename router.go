package main

import (
	"github.com/labstack/echo"
	"html/template"
)

func router(e *echo.Echo) {
	renderer := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	// GET
	e.GET("/", IndexPage)
}
