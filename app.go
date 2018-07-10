package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/routes"
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
	configs.SetConstants()
	// configuracion de vistas
	/*
		r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
			Root:      "views",
			Extension: ".tpl",
			//Master:    "layouts/master",
			Partials: []string{"partials/blank_footer", "partials/blank_header"},
			Funcs: template.FuncMap{
				"sub": func(a, b int) int {
					return a - b
				},
				"copy": func() string {
					return time.Now().Format("2006")
				},
			},
			DisableCache: true,
		})
	*/
	r.HTMLRender = configs.GetViewSetup()
	// configuraciones de  archivos estáticos
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
	// rutas a otros arhcivos
	r.GET("/login", routes.LoginIndex)
	r.GET("/", routes.HomeIndex)
	r.Run() // listen and serve on 0.0.0.0:8080
}
