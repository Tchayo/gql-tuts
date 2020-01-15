package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := "localhost"
	port := "7777"
	pathGQL := "/graphql"
	r := gin.Default()
	// Setup a route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})
	// Inform the user where the server is listening
	log.Println("Running @ http://" + host + ":" + port + pathGQL)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(r.Run(host + ":" + port))

}
