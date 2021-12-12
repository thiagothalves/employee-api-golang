package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thiagothalves/employee-api-golang/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{"5000", gin.Default()}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))

}
