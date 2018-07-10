package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/helpers"
	"github.com/ginv2/structs"
)

func ErrorNoRoute(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.Redirect(http.StatusMovedPermanently, configs.Constants["BASE_URL"]+"error/access/404")
	} else {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"Recurso no disponible",
				"Error 404",
			}}
		c.JSON(404, rpta)
	}
}

func ErrorAccess(c *gin.Context) {
	c.HTML(http.StatusOK, "error/access", gin.H{
		"constants": configs.Constants,
		"title":     "Error",
		"csss":      helpers.ErrorAccessCSS(),
		"jss":       helpers.ErrorAccessJS(),
		"mensaje":   "",
	})
}
