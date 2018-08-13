package routes

import (
	"fmt"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/models"
	"github.com/ginv2/structs"
)

func DepartamentoListar(c *gin.Context) {
	var departamentos []models.Departamento
	db := configs.Database()
	if err := db.Find(&departamentos).Error; err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha podido listar los departamentos",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		defer db.Close()
		c.JSON(200, departamentos)
	}
}

func DepartamentoGuardar(c *gin.Context) {
	var postData string = c.PostForm("data")
	fmt.Println("1 ++++++++++++++++++++++++++++++++++++++")
	fmt.Println(postData)
	data := &structs.TableDepartamentoStruct{}
	err := json.Unmarshal([]byte(postData), data)
	if err != nil {
		fmt.Println(err.Error())
		//invalid character '\'' looking for beginning of object key string
	} else {
		for i := 0; i < len(data.Nuevos); i++ {
			fmt.Println("A. CREAR ++++++++++++++++++++++++++++++++++++++")
			var nuevo = models.Departamento{Nombre: data.Nuevos[i].Nombre}
			db := configs.Database()
			db.Create(&nuevo)
			fmt.Println(nuevo.ID)
		}
		for i := 0; i < len(data.Editados); i++ {
			editado := data.Editados[i]
			var departamentos []models.Departamento
			db := configs.Database()
			db.Model(&departamentos).Where("id = ?", editado.Id).Update("nombre", editado.Nombre)
		}
		for i := 0; i < len(data.Eliminados); i++ {
			eliminado := data.Eliminados[i]
			db := configs.Database()
			db.Where("id = ?", eliminado).Delete(&models.Departamento{})
		}
	}
	fmt.Println("2 ++++++++++++++++++++++++++++++++++++++")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
