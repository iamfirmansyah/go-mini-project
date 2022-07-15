package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorIfDataEmpty(data interface{}) bool {
	if data == nil {
		return true
	}

	return false
}
