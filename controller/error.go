package controller

func ErrorController(err error) {
	if err != nil {
		panic(err)
	}
}
