package utils

func Must[T any](param T, err error) T {
	if err != nil {
		panic(err)
	}
	return param
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
