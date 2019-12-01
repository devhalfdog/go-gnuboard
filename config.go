package main

import (
	"go-gnuboard/controller"
)

// TODO -- config.php 를 옮겨둘 것.

type GG5S struct {
	Version      string // go-gnuboard 버전
	Domain       string // 도메인 주소
	Port         string // 포트 주소
	HttpsDomain  string // 보안 도메인 주소, 없다면 공란
	CookieDomain string // 쿠키 도메인 주소
	Debug        bool   // 디버깅 모드
	DBEngine     string // 데이터베이스 엔진
	DBCharset    string // 데이터베이스 문자셋
	Path         string // 디렉토리 주소
	DataFolder   string // 데이터 폴더 디렉토리 주소
}

// NewConfig 함수는 GG5 구조체를 생성하여 반환하는 함수
func NewConfig() *GG5S {
	return &GG5S{
		Version:      "Go-Gnuboard Test",
		Domain:       "localhost",
		Port:         ":80",
		HttpsDomain:  "",
		CookieDomain: "",
		Debug:        true,
		DBEngine:     "",
		DBCharset:    "utf8mb4",
		Path:         controller.Path(),
		DataFolder:   controller.Path() + "\\data\\",
	}
}
