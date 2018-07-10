package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
)

func main() {
	r := gin.Default()
	configs.SetConstants()
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong??",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
