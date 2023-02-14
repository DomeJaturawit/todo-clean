package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-clean/common"
	"todo-clean/delivery/model"
	"todo-clean/domain"
	"todo-clean/lib/errorLib"
)

// TODO: Move to /model
type GinResponseError struct {
	Title string `json:"title"`
	Error string `json:"error"`
}

func (h newHandler) CreateTodoHandler(ctx *gin.Context) {
	var req model.CreateTodoDeliveryRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, GinResponseError{
			Title: common.ErrFormat.Error(),
			Error: err.Error(),
		})
		return
	}

	if err := errorLib.CheckEmptyStringCreateTodoRequest(req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, GinResponseError{
			Title: common.ErrFormat.Error(),
			Error: err.Error(),
		})
		return
	}

	todo := domain.CreateTodoInputEntity{
		Title:       req.Title,
		Status:      req.Status,
		Description: req.Description,
	}

	tx := ctx.Request.Context()
	resp, err := h.usecase.CreateTodoUseCase(tx, todo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, GinResponseError{
			Title: common.ErrInternal.Error(),
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
