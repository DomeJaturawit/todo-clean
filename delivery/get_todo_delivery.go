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

	//TODO: Check ID Format
	id, _ := uuid.Parse(key)

	if key != "" {
		response, err := h.usecase.GetTodoUseCase(ctx, &id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GinResponseError{
				Title: common.ErrInternal.Error(),
				Error: err.Error(),
			})

		}

		ctx.JSON(http.StatusOK, response)

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
