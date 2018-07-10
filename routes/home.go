package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/helpers"
)

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{
		"constants": configs.Constants,
		"title":     "Home",
		"csss":      helpers.HomeIndexCSS(),
		"jss":       helpers.HomeIndexJS(),
	})
}
