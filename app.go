package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/routes"
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
	r.HTMLRender = configs.GetViewSetup()
	// configuraciones de  archivos estáticos
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	// middleware personalizado
	r.Use(configs.BeforeAll())
	// rutas
	// home
	r.GET("/", routes.HomeIndex)
	r.GET("/pong", GetPong)
	// login
	r.GET("/login", routes.LoginIndex)
	// error
	r.GET("/error/access/:error", routes.ErrorAccess)
	// REST
	/*
		r.GET("/departamento/listar", routes.DepartamentoListar)
		r.POST("/departamento/guardar", routes.DepartamentoGuardar)
		r.GET("/provincia/listar/:departamento_id", routes.ProvinciaListar)
		r.GET("/distrito/buscar", routes.DistritoBuscar)
		r.GET("/distrito/listar/:provincia_id", routes.DistritoListar)
		r.GET("/distrito/nombre/:distrito_id", routes.DistritoNombre)
	*/
	r.GET("/user_state/list", routes.UserStateList)
	r.GET("/user_state/get/:id", routes.UserStateGet)
	r.POST("/permission/save", routes.PermissionSave)
	r.POST("/user/create", routes.UserCreate)
	r.POST("/user/update_state", routes.UserStateUpdate)
	r.POST("/user/update_pass", routes.UserPassUpdate)
	r.POST("/user/delete/:id", routes.UserDelete)
	r.POST("/user_system/validate", routes.UserSystemValidate)
	r.POST("/key/activation/validate", routes.KeyActivationValidate)
	r.POST("/key/reset/validate", routes.KeyResetValidate)
	r.POST("/key/reset_by_email", routes.KeyResetByEmail)
	r.POST("/key/activation/update_by_user_id", routes.KeyActivationUpdateByUserId)
	r.POST("/key/reset/update_by_user_id", routes.KeyResetUpdateByUserId)
	// ruta por default
	r.NoRoute(routes.ErrorNoRoute)
	r.Run(":4100") // listen and serve on 0.0.0.0:8080
}
