package main

import (
	routes "./routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)

	router.Run(port)
}
