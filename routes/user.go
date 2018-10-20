package routes

import (
	"encoding/json"
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
	var rptaStruct structs.KeyActivationStruct
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
					"repeated_user",
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
						"repeated_email",
					}}
			} else if err.Error() == "record not found" {
				//2.2 No es repetido -> OK
				//3. Crear usuario
				var newUser = models.User{
					User:          data.User,
					Pass:          data.Pass,
					Email:         data.Email,
					User_state_id: 6,
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
					//4. Crear asociación usuario sistema
					var newUserSystem = models.UserSystem{
						SystemId: data.SystemId,
						UserId:   idNewUser,
					}
					if err2 := db.Create(&newUserSystem).Error; err2 != nil {
						error = true
						errorStruct = structs.Error{
							TipoMensaje: "error",
							Mensaje: []string{
								"No se ha crear el nuevo usuario",
								err2.Error(),
							}}
					} else {
						//5. Crear key de activación y asociar
						var activationKey = configs.RandStringNumber(40)
						var newUserKey = models.UserKey{
							Activation: activationKey,
							UserId:     idNewUser,
						}
						if err3 := db.Create(&newUserKey).Error; err3 != nil {
							error = true
							errorStruct = structs.Error{
								TipoMensaje: "error",
								Mensaje: []string{
									"No se ha crear el nuevo usuario",
									err3.Error(),
								}}
						} else {
							rptaStruct = structs.KeyActivationStruct{
								UserId:        idNewUser,
								ActivationKey: activationKey,
							}
						}
					}
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
		c.JSON(200, rptaStruct)
	}
}

func UserDelete(c *gin.Context) {
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
		var user []models.User
		if err := db.Where("id = ?", id).Delete(user).Error; err != nil {
			var status int
			var rpta structs.Error
			if err.Error() == "record not found" {
				rpta = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"User to delete does not exist",
						err.Error(),
					}}
				status = 404
			} else {
				rpta = structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"An error has occurred removing the user",
						err.Error(),
					}}
				status = 500
			}
			c.JSON(status, rpta)
		} else {
			defer db.Close()
			c.JSON(200, "User deleted")
		}
	}
}
