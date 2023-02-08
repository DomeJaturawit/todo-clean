package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-clean/common"
	"todo-clean/delivery/model"
	"todo-clean/domain"
	"todo-clean/lib/error_lib"
)

func (h newHandler) CreateTodoHandler(c *gin.Context) {
	var req model.CreateTodoDeliveryRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"title": common.ErrFormat.Error(),
			"error": err.Error(),
		})

	}

	if err := error_lib.CheckEmptyStringCreateTodoRequest(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"title": common.ErrFormat.Error(),
			"error": err.Error(),
		})
	}

	todo := domain.CreateTodoEntityRequest{
		Title:       req.Title,
		Status:      req.Status,
		Description: req.Description,
	}

	ctx := c.Request.Context()
	resp, err := h.usecase.CreateTodoUseCase(ctx, todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"title": common.ErrInternal.Error(),
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, resp)
}
