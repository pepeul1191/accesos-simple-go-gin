package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func UserStateList(c *gin.Context) {
	var userStates []models.UserState
	db := configs.Database()
	if err := db.Find(&userStates).Error; err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"It was not possible to list the user states",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		defer db.Close()
		c.JSON(200, userStates)
	}
}

func UserStateGet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"Sent parameter could not be parsed to integer",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		db := configs.Database()
		var userStates []models.UserState
		if err := db.Where("id = ?", id).Select("id, name").Find(&userStates).Error; err != nil {
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"It was not possible to get the user state",
					err.Error(),
				}}
			c.JSON(500, rpta)
		} else {
			defer db.Close()
			c.JSON(200, userStates)
		}
	}
}
