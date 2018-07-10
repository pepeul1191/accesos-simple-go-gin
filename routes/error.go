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
		url := configs.Constants["BASE_URL"] + "error/access/404"
		c.Redirect(http.StatusMovedPermanently, url)
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
	numeroError := c.Param("error")
	mensaje := ""
	descripcion := ""
	error := ""
	switch numeroError {
	case "404":
		error = "404"
		mensaje = "Archivo no encontrado"
		descripcion = "La página que busca no se encuentra en el servidor"
	case "501":
		error = "404"
		mensaje = "Página en Contrucción"
		descripcion = "Lamentamos el incoveniente, estamos trabajando en ello."
	case "5050":
		error = "5050"
		mensaje = "Acceso restringido"
		descripcion = "No cuenta con los privilegios necesarios"
	case "505":
		error = "505"
		mensaje = "Acceso restringido"
		descripcion = "Necesita estar logueado"
	case "8080":
		error = "8080"
		mensaje = "Tiempo de la sesion agotado"
		descripcion = "Vuelva a ingresar al sistema"
	default:
		error = "404"
		mensaje = "Archivo no encontrado"
		descripcion = "La página que busca no se encuentra en el servidor"
	}
	c.HTML(404, "error/access", gin.H{
		"constants":   configs.Constants,
		"title":       "Error",
		"csss":        helpers.ErrorAccessCSS(),
		"jss":         helpers.ErrorAccessJS(),
		"mensaje":     mensaje,
		"descripcion": descripcion,
		"error":       error,
	})
}
