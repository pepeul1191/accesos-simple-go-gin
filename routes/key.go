package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func KeyActivationValidate(c *gin.Context) {
	var error bool = false
	var count int
	var errorStruct structs.Error
	userId, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	var activationKey string = c.PostForm("activation_key")
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"Parsing error of user_id",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		db := configs.Database()
		var key models.UserKey
		err := db.Where("user_id = ? AND activation = ?", userId, activationKey).Find(&key).Count(&count).Error
		if err != nil {
			if err.Error() == "record not found" {
				count = 0
			} else {
				error = true
				errorStruct = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"An error occurred while validating the user",
						err.Error(),
					}}
			}
		}
		defer db.Close()
		if error == true {
			c.JSON(500, errorStruct)
		} else {
			c.JSON(200, count)
		}
	}
}

func KeyReset(c *gin.Context) {
	userId, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	var resetKey = configs.RandStringNumber(40)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"Parsing error of user_id",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		db := configs.Database()
		var key models.UserKey
		if err2 := db.Model(&key).Where("user_id = ?", userId).Update(
			"Reset", resetKey).Error; err2 != nil {
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"Unable to update the reset key",
					err2.Error(),
				}}
			defer db.Close()
			c.JSON(500, rpta)
		}
		defer db.Close()
		c.JSON(200, resetKey)
	}
}
