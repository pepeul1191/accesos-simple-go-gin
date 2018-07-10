package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
)

func GetPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":    "pong",
		"BASE_URL":   configs.Constants["BASE_URL"],
		"STATIC_URL": configs.Constants["STATIC_URL"],
	})
}

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
	r.GET("/pong", GetPong)
	r.Run() // listen and serve on 0.0.0.0:8080
}
