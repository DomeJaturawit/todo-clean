package delivery

import (
	"github.com/gin-gonic/gin"
	"todo-clean/common"
	"todo-clean/domain"
)

type newHandler struct {
	usecase domain.TodoUseCase
}

func NewHandler(restful *gin.Engine, usecase domain.TodoUseCase) {

	restfulHandler := newHandler{usecase: usecase}

	api := restful.Group(common.APIGroup)
	{
		api.POST(common.APICreatTodoPath, restfulHandler.CreateTodoHandler)
		api.GET(common.APIGetTodoPath, restfulHandler.GetTodoDelivery)
		api.GET(common.APIGetAllTodoPath, restfulHandler.GetTodoDelivery)
		api.PATCH(common.APIUpdateTodoPath, restfulHandler.UpdateTodoHandler)
		api.DELETE(common.APIDeleteTodoPath, restfulHandler.DeleteTodoHandler)
	}
}
