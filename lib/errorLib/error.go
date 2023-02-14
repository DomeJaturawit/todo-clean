package errorLib

import (
	"fmt"
	"todo-clean/common"
	"todo-clean/delivery/model"
)

func WrapError(input string, inputErr error) (err error) {
	err = fmt.Errorf("%s error: %w", input, inputErr)
	return err
}

// TODO: can use package "github.com/go-playground/validator/v10" for input validation
func CheckEmptyStringCreateTodoRequest(input model.CreateTodoDeliveryRequest) (err error) {

	if input.Title == "" {
		return common.ErrTitleEmpty

	} else if input.Status == "" {
		return common.ErrStatusEmpty

	} else if input.Description == "" {
		return common.ErrDescriptionEmpty

	}

	return nil
}
