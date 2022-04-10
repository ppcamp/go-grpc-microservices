package utils

import (
	"fmt"
)

func Wraps(message string, err error) error {
	return fmt.Errorf("%s.\nOriginal error: %v", message, err)
}

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
