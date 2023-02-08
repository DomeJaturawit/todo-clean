package model

type CreateTodoDeliveryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewCreateTodoDeliveryRequest(title string, description string, status string) *CreateTodoDeliveryRequest {
	return &CreateTodoDeliveryRequest{Title: title, Description: description, Status: status}
}
