package handler

import (
	"net/http"
	"simple-todo/inmemory"
	"simple-todo/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Storage *inmemory.TodoStorage
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	todo := h.Storage.GetByID(id)
	c.JSON(http.StatusOK, todo)
}

func (h *Handler) AddTodo(c *gin.Context) {
	var todo todo.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	todo = h.Storage.AddTodo(todo)

	c.JSON(http.StatusCreated, todo)
}

func (h *Handler) GetAllTodo(c *gin.Context) {
	todos := h.Storage.GetAll()
	c.JSON(http.StatusOK, todos)
}
