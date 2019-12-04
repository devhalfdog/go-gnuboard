package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// TODO -- Header 작성

// InstallProcessHeader 미들웨어 함수는 Header를 추가합니다.
func InstallProcessHeader(next echo.HandlerFunc) echo.HandlerFunc {
	// TODO --- Header가 일정 페이지에서만 작동하게 하는 거 찾을 것.
	return func(c echo.Context) error {
		c.Response().Header().Set("Expires", "0")
		c.Response().Header().Set(echo.HeaderLastModified, time.Now().Format("Mon, 02 Dec 2019 09:56:15 GMT")) // Mon, 02 Dec 2019 09:56:15 GMT
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		c.Response().Header().Set("Cache-Control", "pre-check=0, post-check=0, max-age=0")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=utf-8")
		c.Response().Header().Set("X-Robots-Tag", "noindex")

		return next(c)
	}
}

// TODO - 인스톨 프로세스 작성할 것.
func InstallProcess(c echo.Context) error {
	c.Response().Header().Set("Expires", "0")
	c.Response().Header().Set(echo.HeaderLastModified, time.Now().Format("Mon, 02 Dec 2019 09:56:15 GMT")) // Mon, 02 Dec 2019 09:56:15 GMT
	c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	c.Response().Header().Set("Cache-Control", "pre-check=0, post-check=0, max-age=0")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set(echo.HeaderContentType, "text/html; charset=utf-8")
	c.Response().Header().Set("X-Robots-Tag", "noindex")
	c.Response().WriteHeader(http.StatusOK)

	if agree := c.FormValue("agree"); agree != "동의함" {
		return c.HTML(http.StatusOK, `<!doctype html><html><head><meta charset="UTF-8"><title>Go-Gnuboard 설치하기</title>
		<link rel="stylesheet" href="css/install/install.css"></head><body>
		<div class="ins_inner"><p>라이센스(License) 내용에 동의하셔야 설치를 계속하실 수 있습니다.</p><div class="inner_btn"><a href="./">뒤로가기</a></div></div>`)
	}

	return c.Render(http.StatusOK, "install/installer1", map[string]interface{}{
		"Title":   "Go-Gnuboard 설치하기",
		"Version": GG5.Version,
	})
}

func InstallDBProcess(c echo.Context) error {
	a := c.FormValue("database")
	return c.String(http.StatusOK, a)
}
