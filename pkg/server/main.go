package server

import (
	"log"

	"github.com/cmelgarejo/go-gql-server/internal/handlers"
	"github.com/gin-gonic/gin"
)

// HOST declare server host
var HOST string

// PORT declare server port
var PORT string

func init() {
	HOST = "localhost"
	PORT = "7777"
}

// Run : run server
func Run() {
	r := gin.Default()
	// Setup routes
	r.GET("/ping", handlers.Ping())
	log.Println("Running @ http://" + HOST + ":" + PORT)
	log.Fatalln(r.Run(HOST + "+" + PORT))
}
