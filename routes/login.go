package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/helpers"
)

func LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index", gin.H{
		"constants": configs.Constants,
		"title":     "Bienvenido",
		"csss":      helpers.LoginIndexCSS(),
		"jss":       helpers.LoginIndexJS(),
		"mensaje":   "",
	})
}
