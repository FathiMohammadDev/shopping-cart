package api

import (
	"database/sql"

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
	

	
	return server.Run(s.addr)
}
