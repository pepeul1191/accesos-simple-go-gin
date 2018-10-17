package routes

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func UserSystemValidate(c *gin.Context) {
	var postData string = c.PostForm("data")
	data := &structs.UserSystemValidationStruct{}
	err := json.Unmarshal([]byte(postData), data)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha podido validar el usuario, error de parseo del JSON",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var error bool = false
		var found bool = true
		var errorStruct structs.Error
		var user string = data.User
		var systemId string = data.SystemId
		var pass string = data.Pass
		/*
			fmt.Println("1 ++++++++++++++++++++++++++++++++")
			fmt.Println(user)
			fmt.Println("2 ++++++++++++++++++++++++++++++++")
			fmt.Println(pass)
			fmt.Println("3 ++++++++++++++++++++++++++++++++")
			fmt.Println(systemId)
			fmt.Println("4 ++++++++++++++++++++++++++++++++")
		*/
		db := configs.Database()
		var state string = ""
		var userStateSystem models.UserStateSystem
		err := db.Where("user = ? AND pass = ? AND system_id = ?", user, pass, systemId).Find(&userStateSystem).Error
		if err != nil {
			if err.Error() == "record not found" {
				state = "inexistant"
				found = false
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
		if error == true {
			c.JSON(500, errorStruct)
		} else {
			if found {
				state = userStateSystem.State
			}
			c.JSON(200, state)
		}
	}
}

func UserCreate(c *gin.Context) {
	var error bool = false
	var idNewUser int
	var errorStruct structs.Error
	var postData string = c.PostForm("data")
	data := &structs.UserCreateStruct{}
	err := json.Unmarshal([]byte(postData), data)
	if err != nil {
		error = true
		errorStruct = structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha creado el usuario, error de parseo del JSON",
				err.Error(),
			}}
	} else {
		db := configs.Database()
		var user models.User
		var count int
		//1. validar si el nombre de usuario es repetido
		err := db.Where("user = ?", data.User).Find(&user).Count(&count).Error
		if err == nil {
			//1.1 Es repetido -> ERROR
			error = true
			errorStruct = structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"Usuario repetido",
				}}
		} else if err.Error() == "record not found" {
			//1.2 No es repetido -> OK
			//2. Validar si el correo de usuario es repetido
			err := db.Where("email = ?", data.Email).Find(&user).Count(&count).Error
			if err == nil {
				//2.1 Es repetido -> ERROR
				error = true
				errorStruct = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"Correo repetido",
					}}
			} else if err.Error() == "record not found" {
				//2.2 No es repetido -> OK
				//3. Encriptar pass y crear usuario
				var newUser = models.User{
					User:          data.User,
					Pass:          fmt.Sprintf("%x", data.Pass),
					Email:         data.Email,
					User_state_id: 1,
				}
				if err := db.Create(&newUser).Error; err != nil {
					error = true
					errorStruct = structs.Error{
						TipoMensaje: "error",
						Mensaje: []string{
							"No se ha crear el nuevo usuario",
							err.Error(),
						}}
				} else {
					idNewUser = newUser.ID
				}
			} else {
				defer db.Close()
				//fmt.Println(err)
				error = true
				errorStruct = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido valiar si el nombre de correo es repetido",
						err.Error(),
					}}
			}
		} else {
			defer db.Close()
			//fmt.Println(err)
			error = true
			errorStruct = structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido valiar si el nombre de usuario es repetido",
					err.Error(),
				}}
		}
	}
	if error == true {
		c.JSON(500, errorStruct)
	} else {
		c.String(200, strconv.Itoa(idNewUser))
	}
}
