package controller

import "os"

// TODO - FILE 관련 함수들을 여기에다가 작성할 것.

// Path 함수는 현재 디렉토리 주소를 반환하는 함수
func Path() string {
	dir, err := os.Getwd()
	ErrorController(err)

	return dir + "\\"
}

// FileExist 함수는 파일이 존재하는 지 확인하는 함수
func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false // 파일 없음
	} else {
		return true // 파일 있음
	}
}
