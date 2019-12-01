package main

import (
	"github.com/labstack/echo"
	"html/template"
	"net/http"
)

// TODO - RESTful 구현할 것
// TODO -- 접근 불가능한 폴더는 403 으로 넘길 것.

func router(e *echo.Echo) {
	renderer := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "assets")

	// GET
	e.GET("/", IndexPage)          // 인덱스 페이지
	e.GET("/install", InstallPage) // 인스톨 페이지

	// POST
	e.POST("/install/1", InstallProcess) // 인스톨 프로세스

	// 접근불가
	e.GET("/data/dbconfig.json", func(c echo.Context) error {
		return c.String(http.StatusForbidden, "403 Forbidden")
	})
}
