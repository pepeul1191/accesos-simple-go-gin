package routes

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func KeyActivationValidate(c *gin.Context) {
	var error bool = false
	var count int
	var errorStruct structs.Error
	var postData string = c.PostForm("data")
	data := &structs.KeyActivationStruct{}
	err := json.Unmarshal([]byte(postData), data)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"Sent parameter could not be parsed to integer",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var key models.UserKey
		db := configs.Database()
		fmt.Println("1 ++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(data.UserId)
		fmt.Println("2 ++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(data.ActivationKey)
		fmt.Println("3 ++++++++++++++++++++++++++++++++++++++++++++")
		err := db.Where("user_id = ? AND activation = ?", data.UserId, data.ActivationKey).Find(&key).Count(&count).Error
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
		if error == true {
			c.JSON(500, errorStruct)
		} else {
			c.JSON(200, count)
		}
	}
}
