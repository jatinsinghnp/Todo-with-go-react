package main

import (
	"my/app/config"
	"my/app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.Connect()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	routes.TodoRoute(router)

	router.Run()
}
