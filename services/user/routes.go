package user

import (
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

func (h *Handler) HandleLogin(c *gin.Context)   {}
func (h *Handler) HandleRgister(c *gin.Context) {}
