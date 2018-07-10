package main

import (
	"net/http"

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
	// cargando constantes
	//configs.SetConstants()
	// configuraciones de vistas y archivos estáticos
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	// middleware personalizado
	r.Use(configs.BeforeAll())
	// rutas
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong??",
		})
	})
	r.GET("/pong", GetPong)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
			"title": "Home",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
