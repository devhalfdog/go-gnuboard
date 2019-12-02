package main

import (
	"github.com/hjson/hjson-go"
	"go-gnuboard/controller"
	"io/ioutil"
)

// TODO -- config.php 를 옮겨둘 것.

type GG5S struct {
	Version      string // go-gnuboard 버전
	Domain       string // 도메인 주소
	Port         string // 포트 주소
	HTTPSDomain  string // 보안 도메인 주소, 없다면 공란
	CookieDomain string // 쿠키 도메인 주소
	Debug        bool   // 디버깅 모드
	DB           string // DB 종류. MySQL, MariaDB, PostgreSQL 등등
	DBEngine     string // 데이터베이스 엔진
	DBCharset    string // 데이터베이스 문자셋
	Path         string // 디렉토리 주소
	DataFolder   string // 데이터 폴더 디렉토리 주소
}

// NewConfig 함수는 GG5 구조체를 생성하여 반환하는 함수
func NewConfig() *GG5S {
	var gg5 map[string]interface{}

	confAppFile, err := ioutil.ReadFile(controller.Path() + "conf\\app_conf.hjson")
	controller.ErrorController(err)
	if err = hjson.Unmarshal(confAppFile, &gg5); err != nil {
		panic(err)
	}

	return &GG5S{
		Version:      gg5["Version"].(string),
		Domain:       gg5["Domain"].(string),
		Port:         ":" + gg5["Port"].(string),
		HTTPSDomain:  gg5["HTTPSDomain"].(string),
		CookieDomain: gg5["CookieDomain"].(string),
		Debug:        gg5["Debug"].(bool),
		DB:           gg5["DB"].(string),
		DBEngine:     gg5["DBEngine"].(string),
		DBCharset:    gg5["DBCharset"].(string),
		Path:         controller.Path(),
		DataFolder:   controller.Path() + "data\\",
	}
}
