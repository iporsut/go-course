package mrello

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	userRepo UserRepository
	cardRepo CardRepository
	hmacSalt []byte
}

type handlerOption func(*Handler)

func WithHMACSalt(hmacSalt []byte) handlerOption {
	return func(h *Handler) {
		h.hmacSalt = hmacSalt
	}
}

func WithUserRepository(userRepo UserRepository) handlerOption {
	return func(h *Handler) {
		h.userRepo = userRepo
	}
}

func WithCardRepository(cardRepo CardRepository) handlerOption {
	return func(h *Handler) {
		h.cardRepo = cardRepo
	}
}

func NewHandler(options ...handlerOption) *Handler {
	h := &Handler{}
	for _, optionalFunc := range options {
		optionalFunc(h)
	}

	return h
}

func (h *Handler) UserRegister(c *gin.Context) {
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
	type Request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=64"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindUserByEmail(c.Request.Context(), req.Email)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !user.IsPasswordMatch(req.Password) {
		log.Println("password not match")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		Email: user.Email,
	})

	tokenStr, err := token.SignedString(h.hmacSalt)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	// TODO: implement
	// leave it just wait for token to expire
}

func (h *Handler) CreateCard(c *gin.Context) {
	type Request struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := GetUserFromContext(c)

	card := user.NewCard(req.Title, req.Description)

	card, err := h.cardRepo.CreateCard(c.Request.Context(), card)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, card)
}

func (h *Handler) UpdateCard(c *gin.Context) {
	type Request struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	card, err := h.cardRepo.FindCardByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		if IsErrCode(err, ErrCodeNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	user := GetUserFromContext(c)

	card.Update(req.Title, req.Description, user.ID)

	card, err = h.cardRepo.SaveCard(c.Request.Context(), card)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, card)
}

func (h *Handler) MoveCard(c *gin.Context) {
	type Request struct {
		Column string `json:"column" binding:"required,oneof=todo doing done"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	card, err := h.cardRepo.FindCardByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		if IsErrCode(err, ErrCodeNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	user := GetUserFromContext(c)

	card.MoveToColumn(req.Column, user.ID)

	card, err = h.cardRepo.SaveCard(c.Request.Context(), card)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, card)

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

func (h *Handler) GetBoard(c *gin.Context) {
	cards, err := h.cardRepo.GetAllCards(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	board := CreateBoardFromCards(cards)

	c.JSON(http.StatusOK, board)
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
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

	user, err := h.userRepo.FindUserByEmail(c.Request.Context(), claims.Email)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("user", user)

	c.Next()
}

func (h *Handler) Frontend(c *gin.Context) {
	c.File("./frontend/index.html")
}

func GetUserFromContext(c *gin.Context) *User {
	v, ok := c.Get("user")
	if !ok {
		return nil
	}

	user, ok := v.(*User)
	if !ok {
		return nil
	}

	return user
}
