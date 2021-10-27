package api

import (
	"github.com/georgechristman/golang-okta-viper-config-demo/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/")
	authRoutes.POST("/accounts", server.createAccount)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start() error {
	return server.router.Run()
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
