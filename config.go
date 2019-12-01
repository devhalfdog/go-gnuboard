package main

// TODO -- config.php 를 옮겨둘 것.

type GG5 struct {
	Version      string // go-gnuboard 버전
	Domain       string // 도메인 주소
	Port         string // 포트 주소
	HttpsDomain  string // 보안 도메인 주소, 없다면 공란
	CookieDomain string // 쿠키 도메인 주소
	Debug        bool   // 디버깅 모드
	DBEngine     string // 데이터베이스 엔진
	DBCharset    string // 데이터베이스 문자셋
}

// NewConfig 함수는 GG5 구조체를 생성하여 반환하는 함수
func NewConfig() *GG5 {
	return &GG5{
		Version:      "Go-Gnuboard Test",
		Domain:       "localhost",
		Port:         ":80",
		HttpsDomain:  "",
		CookieDomain: "",
		Debug:        true,
		DBEngine:     "",
		DBCharset:    "utf8mb4",
	}
}
