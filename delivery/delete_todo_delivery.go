package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"todo-clean/common"
	"todo-clean/delivery/model"
	"todo-clean/domain"
)

func (h newHandler) DeleteTodoHandler(ctx *gin.Context) {
	key := ctx.Param("id")

	id, err := uuid.Parse(key)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GinResponseError{
			Title: common.ErrFormat.Error(),
			Error: err.Error(),
		})
	}
	queryEntity := domain.DeleteTodoQueryEntity{ID: id}
	_, err = h.usecase.DeleteTodoUseCase(ctx, queryEntity)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GinResponseError{
			Title: common.ErrInternal.Error(),
			Error: err.Error(),
		})

	} else {
		ctx.JSON(http.StatusOK, model.GinDeleteResponse{
			ID:     queryEntity.ID,
			Status: common.DeleteSuccess,
		})
	}
}
