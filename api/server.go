package api

import (
	"github.com/georgechristman/golang-okta-rest-api-viper-config-demo/util"
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

	router.GET("/login", server.listPersons)

	authRoutes := router.Group("/").Use(authMiddleware(server.config))
	authRoutes.GET("/persons", server.listPersons)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
