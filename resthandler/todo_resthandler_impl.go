package resthandler

import (
	"backend-mytodo/service"
	"backend-mytodo/shareddomain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoResthandlerImpl struct {
	service service.TodoService
}

type ResponseWithError struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func NewTodoResthandler(service service.TodoService) TodoResthandler {
	return &TodoResthandlerImpl{service}
}

func (handler *TodoResthandlerImpl) Create(c *gin.Context) {
	var request shareddomain.RequestCreate
	if err := c.ShouldBindJSON(&request); err != nil {
		response := ResponseWithError{
			Code:    http.StatusUnprocessableEntity,
			Success: false,
			Message: "gagal menambahkan todo",
			Error:   err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err := handler.service.Create(request); err != nil {
		response := ResponseWithError{
			Code:    http.StatusUnprocessableEntity,
			Success: false,
			Message: "gagal menambahkan todo",
			Error:   err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := shareddomain.ResponseWithoutData{
		Code:    http.StatusCreated,
		Success: true,
		Message: "berhasil menambahkan todo",
	}
	c.JSON(http.StatusCreated, response)
}