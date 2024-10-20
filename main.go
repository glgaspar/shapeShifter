package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool
	Message string
	Data    interface{}
}

func main() {
	router := gin.Default()
	setRoutes(router)

	router.Run("0.0.0.0:8080")
}

func processor(c *gin.Context) {
	
}

func setRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) { c.Done() })
	router.POST("/", processor)
}
