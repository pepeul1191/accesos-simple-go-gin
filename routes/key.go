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

func KeyResetByUserId(c *gin.Context) {
	userId, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	var rptaError structs.Error
	var status int
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
		err2 := db.Where("user_id = ?", userId).Find(&key).Error
		if err2 != nil {
			if err2.Error() == "record not found" {
				rptaError = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"User not found",
						err2.Error(),
					}}
				status = 404
			} else {
				rptaError = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"Unable to update the reset key",
						err2.Error(),
					}}
				status = 500
			}
			defer db.Close()
			c.JSON(status, rptaError)
		} else {
			var resetKey = configs.RandStringNumber(40)
			key.Reset = resetKey
			db.Model(&key).Update("reset", resetKey)
			defer db.Close()
			c.JSON(200, resetKey)
		}
	}
}

func KeyResetByEmail(c *gin.Context) {
	var email string = c.PostForm("email")
	var rptaError structs.Error
	var status int
	db := configs.Database()
	var user models.User
	var key models.UserKey
	err := db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			rptaError = structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"User not found",
					err.Error(),
				}}
			status = 404
		} else {
			rptaError = structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"Unable to find the user",
					err.Error(),
				}}
			status = 500
		}
		defer db.Close()
		c.JSON(status, rptaError)
	} else {
		err2 := db.Where("user_id = ?", user.ID).Find(&key).Error
		if err2 != nil {
			if err2.Error() == "record not found" {
				rptaError = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"User Key not found",
						err2.Error(),
					}}
				status = 404
			} else {
				rptaError = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"Unable to update the reset user key",
						err2.Error(),
					}}
				status = 500
			}
			defer db.Close()
			c.JSON(status, rptaError)
		} else {
			var resetKey = configs.RandStringNumber(40)
			key.Reset = resetKey
			db.Model(&key).Update("reset", resetKey)
			defer db.Close()
			c.JSON(200, resetKey)
		}
	}
}
