package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"todo-clean/common"
	"todo-clean/delivery/model"
)

func (h newHandler) GetTodoDelivery(ctx *gin.Context) {
	key := ctx.Param("id")

	if key != "" {
		id, err := uuid.Parse(key)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GinResponseError{
				Title: common.ErrFormat.Error(),
				Error: err.Error(),
			})
		}
		response, err := h.usecase.GetTodoUseCase(ctx, &id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.GinResponseError{
				Title: common.ErrDataNotFound.Error(),
				Error: err.Error(),
			})

		} else {
			ctx.JSON(http.StatusOK, response)
		}

	} else {
		response, err := h.usecase.GetTodoUseCase(ctx, nil)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GinResponseError{
				Title: common.ErrInternal.Error(),
				Error: err.Error(),
			})

		}
		ctx.JSON(http.StatusOK, response)
	}
}
