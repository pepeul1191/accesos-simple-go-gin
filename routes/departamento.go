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
	data := &structs.TableDepartamentoStruct{}
	err := json.Unmarshal([]byte(postData), data)
	var nuevosId []int
	if err != nil {
		fmt.Println(err.Error())
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha guardar los departamentos, error de parseo del JSON",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		tx := configs.Database().Begin()
		for i := 0; i < len(data.Nuevos); i++ {
			var nuevo = models.Departamento{
				Nombre: data.Nuevos[i].Nombre}
			tx.Create(&nuevo)
			fmt.Println(nuevo.ID)
			//temp = {'temporal' : temp_id, 'nuevo_id' : s.id}
			nuevosId = append(nuevosId, nuevo.ID)
		}
		for i := 0; i < len(data.Editados); i++ {
			editado := data.Editados[i]
			var departamentos []models.Departamento
			tx.Model(&departamentos).Where("id = ?", editado.Id).Update(
				"nombre", editado.Nombre)
		}
		for i := 0; i < len(data.Eliminados); i++ {
			eliminado := data.Eliminados[i]
			tx.Where("id = ?", eliminado).Delete(&models.Departamento{})
		}
		if tx.Error != nil {
			tx.Rollback()
			tx.Close()
			//errors := tx.Error
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido listar los departamentos",
					err.Error(),
				}}
			c.JSON(500, rpta)
		} else {
			tx.Commit()
			tx.Close()
			mensaje := []interface{}{
				"Se ha registrado los cambios en los departamentos",
				nuevosId}
			c.JSON(200, gin.H{
				"tipo_mensaje": "success",
				"mensaje":      mensaje,
			})
		}
	}
}
