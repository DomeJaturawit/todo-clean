package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"todo-clean/common"
	"todo-clean/delivery/model"
	"todo-clean/domain"
)

// TODO: Fix Bug Send Empty Field Response "" But Not Keep Empty Data To DB
func (h newHandler) UpdateTodoHandler(ctx *gin.Context) {
	key := ctx.Param("id")
	request := new(model.UpdateTodoDeliveryRequest)
	id, err := uuid.Parse(key)
	if err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GinResponseError{
			Title: common.ErrFormat.Error(),
			Error: err.Error(),
		})
	}

	if err := ctx.BindJSON(request); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.GinResponseError{
			Title: common.ErrFormat.Error(),
			Error: err.Error(),
		})
		return
	}

	queryEntity := domain.QueryUpdateTodoEntity{ID: id}

	entity := domain.UpdateTodoEntity{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
	}

	response, err := h.usecase.UpdateTodoUseCase(ctx, queryEntity, &entity)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.GinResponseError{
			Title: common.ErrInternal.Error(),
			Error: err.Error(),
		})

	} else {
		ctx.JSON(http.StatusOK, response)
	}

}
