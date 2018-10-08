package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/helpers"
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
