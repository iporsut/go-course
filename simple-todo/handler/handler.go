package handler

import (
	"net/http"
	"simple-todo/inmemory"
	"simple-todo/postgresql"
	"simple-todo/todo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Storage *inmemory.TodoStorage
	DB      *postgresql.DB
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.DB.GetByID(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *Handler) AddTodo(c *gin.Context) {
	todo := new(todo.Todo)
	if err := c.ShouldBind(todo); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	todo, err := h.DB.AddTodo(c.Request.Context(), todo)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusCreated, todo)
}

func (h *Handler) GetAllTodo(c *gin.Context) {
	todos, err := h.DB.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, todos)
}
