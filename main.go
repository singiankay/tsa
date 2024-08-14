package main

import (
	"github.com/gin-gonic/gin"
	"github.com/singiankay/tsa/config"
	"github.com/singiankay/tsa/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.ContactRoute(router)
	router.Run(":8080")
}