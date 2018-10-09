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

func UserCreate(c *gin.Context) {
	var postData string = c.PostForm("data")
	data := &structs.TablePermissionStruct{}
	err := json.Unmarshal([]byte(postData), data)
	var newUserStruct structs.UserCreateStruct
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha creado el usuario, error de parseo del JSON",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		db := configs.Database()
		var user models.User
		var count int
		//error en validar usuario repetido
		err := db.Where("user = ?", newUserStruct.User).Find(&user).Count(&count).Error
		if err == nil {
			c.String(200, strconv.Itoa(count))
		} else if err.Error() == "record not found" {
			c.String(200, strconv.Itoa(0))
		} else {
			defer db.Close()
			fmt.Println(err)
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido valiar si el nombre de usuario es repetido",
					err.Error(),
				}}
			c.JSON(500, rpta)
		}
	}
}
