package api

import (
	"database/sql"

	"github.com/FathiMohammadDev/shopping-cart/services/user"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	db   *sql.DB
	addr string
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{db, addr}
}

func (s *APIServer) Run() error {
	server := gin.Default()
	subrouter := server.Group("api/v1")
	
	userStore := user.NewUserStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	return server.Run(s.addr)
}
