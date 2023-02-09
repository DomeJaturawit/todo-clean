package error_lib

import (
	"fmt"
	"todo-clean/common"
	"todo-clean/delivery/model"
)

func WrapError(input string, inputErr error) (err error) {
	err = fmt.Errorf("%s error: %w", input, inputErr)
	return err
}

func CheckEmptyStringCreateTodoRequest(input model.CreateTodoDeliveryRequest) (err error) {

	if input.Title == "" {
		return WrapError(common.ErrTitleEmpty.Error(), err)

	} else if input.Status == "" {
		return WrapError(common.ErrStatusEmpty.Error(), err)

	} else if input.Description == "" {
		return WrapError(common.ErrDescriptionEmpty.Error(), err)

	}

	return nil
}
