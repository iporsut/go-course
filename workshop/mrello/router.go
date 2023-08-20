package mrello

import (
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine, h *Handler) {
	e.POST("/register-user", h.UserRegister)

	e.GET("/app", h.Frontend)

	e.POST("/login", h.Login)
	e.POST("/logout", h.AuthMiddleware, h.Logout)

	e.POST("/cards", h.AuthMiddleware, h.CreateCard)
	e.PUT("/cards/:id", h.AuthMiddleware, h.UpdateCard)
	e.PUT("/cards/:id/move", h.AuthMiddleware, h.MoveCard)
	e.DELETE("/cards/:id", h.AuthMiddleware, h.DeleteCard)
	e.GET("/cards/:id/history", h.AuthMiddleware, h.ListCardEditHistory)

	e.GET("/board", h.AuthMiddleware, h.GetBoard)
}
