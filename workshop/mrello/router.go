package mrello

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func Router(e *gin.Engine, h *Handler) {
	e.GET("/app", h.Frontend)
	e.POST("/users/register", h.UserRegistration)

	e.POST("/users/login", h.Login)
	e.POST("/users/logout", h.AuthMiddleware, h.Logout)

	e.POST("/cards", h.AuthMiddleware, h.CreateCard)
	e.PUT("/cards/:id", h.AuthMiddleware, h.UpdateCard)
	e.PUT("/cards/:id/move", h.AuthMiddleware, h.MoveCard)
	e.DELETE("/cards/:id", h.AuthMiddleware, h.DeleteCard)
	e.GET("/cards/:id/history", h.AuthMiddleware, h.ListCardEditHistory)
}
