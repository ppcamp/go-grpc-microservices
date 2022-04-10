package utils

import (
	"fmt"
)

func Wraps(message string, err error) error {
	return fmt.Errorf("%s.\nOriginal error: %v", message, err)
}
