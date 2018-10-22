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
				"Error de parseo de user_id",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var key models.UserKey
		db := configs.Database()
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
	var userId string = c.PostForm("user_id")
	var key models.UserKey
	var resetKey = configs.RandStringNumber(40)
	db := configs.Database()
	if err2 := db.Model(&key).Where("user_id = ?", userId).Update(
		"Reset", resetKey).Error; err2 != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha podido actualizar el reset key",
				err2.Error(),
			}}
		defer db.Close()
		c.JSON(500, rpta)
	}
	defer db.Close()
	c.JSON(200, resetKey)
}
