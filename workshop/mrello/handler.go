package mrello

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	userRepo UserRepository
	hmacSalt []byte
}

type handerOption func(*Handler)

func WithHMACSalt(hmacSalt []byte) handerOption {
	return func(h *Handler) {
		h.hmacSalt = hmacSalt
	}
}

func WithUserRepository(userRepo UserRepository) handerOption {
	return func(h *Handler) {
		h.userRepo = userRepo
	}
}

func NewHandler(options ...handerOption) *Handler {
	h := &Handler{}
	for _, option := range options {
		option(h)
	}

	return h
}

func (h *Handler) UserRegistration(c *gin.Context) {
	type Request struct {
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required,min=8,max=64"`
		PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	bh, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	passwordHash := string(bh)

	user, err := h.userRepo.FindUserByEmail(c.Request.Context(), req.Email)
	if err != nil && !IsErrCode(err, ErrCodeNotFound) {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	if _, err := h.userRepo.CreateUser(c.Request.Context(), req.Email, passwordHash); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) Login(c *gin.Context) {
	// TODO: implement
	// 1. get email and password from request body
	// 2. find user by email
	// 3. compare password hash
	// 4. generate jwt token
	// 5. return token
}

func (h *Handler) Logout(c *gin.Context) {
	// TODO: implement
	// leave it just wait for token to expire
}

func (h *Handler) CreateCard(c *gin.Context) {
	// TODO: implement
	// 1. get card data from request body
	// 2. create card
	// 3. return card
}

func (h *Handler) UpdateCard(c *gin.Context) {
	// TODO: implement
	// 1. get card data from request body
	// 2. update card
	// 3. return card
}

func (h *Handler) MoveCard(c *gin.Context) {
	// TODO: implement
	// 1. get card data from request body
	// 2. move card
	// 3. return card
}

func (h *Handler) DeleteCard(c *gin.Context) {
	// TODO: implement
	// 1. delete card
	// 2. return card
}

func (h *Handler) ListCardEditHistory(c *gin.Context) {
	// TODO: implement
	// 1. get card id from path
	// 2. list card edit history

}

func (h *Handler) AuthMiddleware(c *gin.Context) {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	strToken := strings.TrimPrefix(bearerToken, "Bearer ")

	var claims CustomClaims
	token, err := jwt.ParseWithClaims(strToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return h.hmacSalt, nil
	})

	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func (h *Handler) Frontend(c *gin.Context) {
	c.File("./frontend/index.html")
}
