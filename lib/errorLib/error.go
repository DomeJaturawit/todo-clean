package errorLib

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"todo-clean/common"
	"todo-clean/delivery/model"
)

func WrapError(input string, inputErr error) (err error) {
	err = fmt.Errorf("%s error: %w", input, inputErr)
	return err
}

func CheckEmptyInput(input model.CreateTodoDeliveryRequest) (err error) {

	validate := validator.New()

	err = validate.Struct(input)

	if err == nil {
		return
	}

	validationErr, ok := err.(validator.ValidationErrors)
	if !ok {

		return common.ErrFormat
	}

	for _, vErr := range validationErr {
		err := fmt.Errorf("%s has a value of '%v' ", vErr.Field(), vErr.Value())
		return err
	}

	return nil
}
