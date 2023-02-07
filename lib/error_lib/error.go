package error_lib

import "fmt"

func WrapError(input string, inputErr error) (err error) {
	err = fmt.Errorf("%s error: %w", input, inputErr)
	return err

}
