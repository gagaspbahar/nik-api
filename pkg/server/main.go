package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start(port string) error
	App() *gin.Engine
}

type Server struct {
	app *gin.Engine
}

func (s *Server) App() *gin.Engine {
	return s.app
}

func (s *Server) Start(port string) error {
	log.Printf("Starting server on port %s", port)
	return s.app.Run(port)
}

func NewServer() *Server {
	return &Server{
		app: gin.Default(),
	}
}
