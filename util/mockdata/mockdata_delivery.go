package mockdata

import (
	"todo-clean/delivery/model"
)

func CreateTodoDeliveryRequestMockData() *model.CreateTodoDeliveryRequest {
	return model.NewCreateTodoDeliveryRequest(
		mockTitle,
		mockDescription,
		mockStatus,
	)
}

func UpdateTodoDeliveryRequestMockData() *model.UpdateTodoDeliveryRequest {
	return model.NewUpdateTodoDeliveryRequest(
		mockTitle,
		mockDescription,
		mockStatus,
	)
}
