package main

// TODO -- config.php 를 옮겨둘 것.

var GG5_VERSION string
var GG5_GNUBOARD_VERSION string
var GG5_DOMAIN string
var GG5_HTTPS_DOMAIN string
var GG5_DEBUG bool
var GG5_DB_ENGINE string
var GG5_DB_CHARSET string

// config 초기화 함수
func init() {
	GG5_VERSION = "Golang 그누보드 5" // 버전
	GG5_GNUBOARD_VERSION = "0.0.1" // Golang 그누보드 버전
	GG5_DOMAIN = "" // 도메인 주소
	GG5_HTTPS_DOMAIN = "" // HTTPS 가 없다면 공란, 있다면 기재할 것.
	GG5_DEBUG = false // 실제 서버 운영 시에는 false
	GG5_DB_ENGINE = "" // 데이터베이스 테이블 엔진을 세팅. 만약 MyISAM 혹은 InnoDB 를 원한다면 MyISAM, InnoDB 로 바꿀 것.
	GG5_DB_CHARSET = "utf8mb4" // 데이터베이스 문자셋 세팅.


}