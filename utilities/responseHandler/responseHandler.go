package responseHandler

func ErrorHandler(e error) {
	if e != nil {
		panic(e.Error())
	}
}