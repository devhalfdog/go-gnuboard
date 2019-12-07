package controller

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// TODO - DATA 폴더와 관련된 데이터 컨트롤러 함수는 여기에다가 작성할 것.

type DBConfig struct {
	Database Database `json:"Database"` // 데이터베이스
	Table    Table    `json:"Table"`    // 테이블
}

type Database struct {
	Engine      string `json:"Engine"`      // PostgreSQL 엔진인지, MySQL 엔진인지.
	Host        string `json:"Host"`        // DB 호스트
	Port        string `json:"Port"`        // DB 포트
	User        string `json:"User"`        // DB 유저
	Password    string `json:"Password"`    // DB 패스워드
	Table       string `json:"Table"`       // DB 테이블
	TablePrefix string `json:"TablePrefix"` // DB 테이블 접두사
}

type Table struct {
	WritePrefix   string `json:"WritePrefix"`   // 게시판 테이블명 접두사
	Auth          string `json:"Auth"`          // 관리권한 설정 테이블
	Config        string `json:"Config"`        // 기본환경 설정 테이블
	Group         string `json:"Group"`         // 게시판 그룹 테이블
	GroupMember   string `json:"GroupMember"`   // 게시판 그룹 + 회원 테이블
	Board         string `json:"Board"`         // 게시판 설정 테이블
	BoardFile     string `json:"BoardFile"`     // 게시판 첨부파일 테이블
	BoardGood     string `json:"BoardGood"`     // 게시물 추천, 비추천 테이블
	BoardNew      string `json:"BoardNew"`      // 게시판 새글 테이블
	Login         string `json:"Login"`         // 로그인 테이블 (접속자 수)
	Mail          string `json:"Mail"`          // 회원 메일 테이블
	Member        string `json:"Member"`        // 회원 테이블
	Memo          string `json:"Memo"`          // 메모 테이블
	Poll          string `json:"Poll"`          // 투표 테이블
	PollEtc       string `json:"PollEtc"`       // 투표 기타의견 테이블
	Point         string `json:"Point"`         // 포인트 테이블
	Popular       string `json:"Popular"`       // 인기검색어 테이블
	Scrap         string `json:"Scrap"`         // 스크랩 테이블
	Visit         string `json:"Visit"`         // 방문자 테이블
	VisitSum      string `json:"VisitSum"`      // 방문자 합계 테이블
	UniqueID      string `json:"UniqueID"`      // 유니크한 값을 만드는 테이블
	Autosave      string `json:"Autosave"`      // 게시글 작성 시 일정시간마다 글을 임시저장하는 테이블
	CertHistory   string `json:"CertHistory"`   // 인증내역 테이블
	QAConfig      string `json:"QAConfig"`      // 1:1 문의 설정 테이블
	QAContent     string `json:"QAContent"`     // 1:1 문의 테이블
	Content       string `json:"content"`       // 내용(컨텐츠) 정보 테이블
	FAQ           string `json:"FAQ"`           // FAQ 테이블
	FAQMaster     string `json:"FAQMaster"`     // FAQ 마스터 테이블
	NewWin        string `json:"NewWin"`        // 새창 테이블
	Menu          string `json:"Menu"`          // 메뉴 테이블
	SocialProfile string `json:"SocialProfile"` // 소셜 로그인 테이블
}

// CreateDBConfig 함수는 dbconfig.json 파일이 없으면 생성하는 함수
func CreateDBConfig(db map[string]string, path string) bool {
	if exist := FileExist(path); exist != true {
		// dbconfig 파일이 없다면
		_, err := os.OpenFile(
			path,
			os.O_CREATE|os.O_RDWR,
			os.FileMode(0644),
		)
		ErrorController(err)

		WriteDBConfig(db, path)
	}
	return true
}

// WriteDBConfig 함수는 dbconfig.json 파일을 작성하는 함수
func WriteDBConfig(db map[string]string, path string) {
	// TODO - DBConfig 구조체를 JSON 패키지를 이용해서 data/dbconfig.json 에 작성할 것.
	dc := make([]DBConfig, 1)

	dc[0].Database.Engine = db["Database"]
	dc[0].Database.Host = db["Host"]
	dc[0].Database.Port = db["Port"]
	dc[0].Database.User = db["User"]
	dc[0].Database.Password = db["Password"]
	dc[0].Database.Table = db["DB"]
	dc[0].Database.TablePrefix = db["TablePrefix"]
	dc[0].Table.WritePrefix = dc[0].Database.TablePrefix + "write_"
	dc[0].Table.Auth = dc[0].Database.TablePrefix + "auth"
	dc[0].Table.Config = dc[0].Database.TablePrefix + "config"
	dc[0].Table.Group = dc[0].Database.TablePrefix + "group"
	dc[0].Table.GroupMember = dc[0].Database.TablePrefix + "group_member"
	dc[0].Table.Board = dc[0].Database.TablePrefix + "board"
	dc[0].Table.BoardFile = dc[0].Database.TablePrefix + "board_file"
	dc[0].Table.BoardGood = dc[0].Database.TablePrefix + "board_good"
	dc[0].Table.BoardNew = dc[0].Database.TablePrefix + "board_new"
	dc[0].Table.Login = dc[0].Database.TablePrefix + "login"
	dc[0].Table.Mail = dc[0].Database.TablePrefix + "mail"
	dc[0].Table.Member = dc[0].Database.TablePrefix + "member"
	dc[0].Table.Memo = dc[0].Database.TablePrefix + "memo"
	dc[0].Table.Poll = dc[0].Database.TablePrefix + "poll"
	dc[0].Table.PollEtc = dc[0].Database.TablePrefix + "poll_etc"
	dc[0].Table.Point = dc[0].Database.TablePrefix + "point"
	dc[0].Table.Popular = dc[0].Database.TablePrefix + "popular"
	dc[0].Table.Scrap = dc[0].Database.TablePrefix + "scrap"
	dc[0].Table.Visit = dc[0].Database.TablePrefix + "visit"
	dc[0].Table.VisitSum = dc[0].Database.TablePrefix + "visit_sum"
	dc[0].Table.UniqueID = dc[0].Database.TablePrefix + "uniqid"
	dc[0].Table.Autosave = dc[0].Database.TablePrefix + "autosave"
	dc[0].Table.CertHistory = dc[0].Database.TablePrefix + "cert_history"
	dc[0].Table.QAConfig = dc[0].Database.TablePrefix + "qa_config"
	dc[0].Table.QAContent = dc[0].Database.TablePrefix + "qa_content"
	dc[0].Table.Content = dc[0].Database.TablePrefix + "content"
	dc[0].Table.FAQ = dc[0].Database.TablePrefix + "faq"
	dc[0].Table.FAQMaster = dc[0].Database.TablePrefix + "faq_master"
	dc[0].Table.NewWin = dc[0].Database.TablePrefix + "new_win"
	dc[0].Table.Menu = dc[0].Database.TablePrefix + "menu"
	dc[0].Table.SocialProfile = dc[0].Database.TablePrefix + "member_social_profiles"

	doc, err := json.Marshal(dc)
	ErrorController(err)

	err = ioutil.WriteFile(path, doc, os.FileMode(0644)) // 파일 저장
	ErrorController(err)
}
