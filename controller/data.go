package controller

import "os"

// TODO - DATA 폴더와 관련된 데이터 컨트롤러 함수는 여기에다가 작성할 것.
// TODO -- DATA 관련 및 DB 관련 함수들을 추가할 것.

type DBConfig struct {
	Database []struct {
		Engine      string `json:"Engine"`      // PostgreSQL 엔진인지, MySQL 엔진인지.
		Host        string `json:"Host"`        // DB 호스트
		User        string `json:"User"`        // DB 유저
		Password    string `json:"Password"`    // DB 패스워드
		Table       string `json:"Table"`       // DB 테이블
		TablePrefix string `json:"TablePrefix"` // DB 테이블 접두사
	} `json:"Database"`
	Table []struct {
		WritePrefix   string `json:"WritePrefix"`   // 게시판 테이블명 접두사
		Auth          string `json:"Auth"`          // 관리권한 설정 테이블
		Config        string `json:"Config"`        // 기본환경 설정 테이블
		Group         string `json:"Group"`         // 게시판 그룹 테이블
		GroupMember   string `json:"GroupMember"`   // 게시판 그룹 + 회원 테이블
		Board         string `json:"Board"`         // 게시판 설정 테이블
		BoardFill     string `json:"BoardFill"`     // 게시판 첨부파일 테이블
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
	} `json:"Table"`
}

// CreateDBConfig 함수는 dbconfig.json 파일이 없으면 생성하는 함수
func CreateDBConfig(path string) bool {
	if exist := FileExist(path); exist != true {
		// dbconfig 파일이 없다면
		file, err := os.OpenFile(
			path,
			os.O_CREATE|os.O_RDWR,
			os.FileMode(0644),
		)
		ErrorController(err)

		defer file.Close()

		WriteDBConfig(file)
	}
	return true
}

// WriteDBConfig 함수는 dbconfig.json 파일을 작성하는 함수
func WriteDBConfig(f *os.File) {
	// TODO - DBConfig 구조체를 JSON 패키지를 이용해서 data/dbconfig.json 에 작성할 것.
}
