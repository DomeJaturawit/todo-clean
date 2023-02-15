package model

import "github.com/google/uuid"

type CreateTodoDeliveryRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

type UpdateTodoDeliveryRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

func NewUpdateTodoDeliveryRequest(title string, description string, status string) *UpdateTodoDeliveryRequest {
	return &UpdateTodoDeliveryRequest{Title: title, Description: description, Status: status}
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
