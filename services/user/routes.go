package user

import (
	"net/http"

	"github.com/FathiMohammadDev/shopping-cart/services/auth"
	"github.com/FathiMohammadDev/shopping-cart/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.HandleRgister)
	router.POST("/login", h.HandleLogin)
}

func (h *Handler) HandleLogin(c *gin.Context) {}

func (h *Handler) HandleRgister(c *gin.Context) {
	var payload types.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//hashed password
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// create user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//write success res
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfuly."})
}
