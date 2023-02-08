package delivery

import (
	"github.com/gin-gonic/gin"
	"todo-clean/common"
	"todo-clean/domain"
)

type newHandler struct {
	usecase domain.TodoUseCaseInterface
}

func NewHandler(restful *gin.Engine, usecase domain.TodoUseCaseInterface) {

	restfulHandler := newHandler{usecase: usecase}

	//TODO Make it const
	api := restful.Group(common.APIGroup)
	{
		api.POST(common.APITodoCreatPath, restfulHandler.CreateTodoHandler)
	}
}
