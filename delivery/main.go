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
		api.POST(common.APITodoCreatPath, restfulHandler.CreateTodoHandler)
		api.GET(common.APITodoGetPath, restfulHandler.GetTodoDelivery)
		api.GET(common.APIAllTodoGetPath, restfulHandler.GetTodoDelivery)
	}
}
