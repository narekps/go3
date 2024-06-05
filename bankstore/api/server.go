package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/narekps/go3/bankstore/db/sqlc"
	"github.com/narekps/go3/bankstore/util"
)

type Server struct {
	config util.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts", server.updateAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
