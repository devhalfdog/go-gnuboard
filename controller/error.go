package controller

import "fmt"

func ErrorController(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
