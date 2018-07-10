package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
)

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index", gin.H{
		"constants": configs.Constants,
		"title":     "Home",
	})
}
