package model

import "github.com/google/uuid"

type CreateTodoDeliveryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type GetTodoDeliveryRequest struct {
	ID uuid.UUID `json:"id"`
}

type GinResponseError struct {
	Title string `json:"title"`
	Error string `json:"error"`
}

func NewCreateTodoDeliveryRequest(title string, description string, status string) *CreateTodoDeliveryRequest {
	return &CreateTodoDeliveryRequest{Title: title, Description: description, Status: status}
}
