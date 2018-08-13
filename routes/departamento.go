package routes

import (
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
	var nuevosId []structs.IdsNuevosTemp
	if err != nil {
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
			if err := tx.Create(&nuevo).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los departamentos",
						err.Error(),
					}}
				c.JSON(500, rpta)
			}
			nuevosId = append(nuevosId,
				structs.IdsNuevosTemp{
					Temporal: data.Nuevos[i].Id,
					NuevoId:  nuevo.ID})
		}
		for i := 0; i < len(data.Editados); i++ {
			editado := data.Editados[i]
			var departamentos []models.Departamento
			if err := tx.Model(&departamentos).Where("id = ?", editado.Id).Update(
				"nombre", editado.Nombre).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los departamentos",
						err.Error(),
					}}
				c.JSON(500, rpta)
			}
		}
		for i := 0; i < len(data.Eliminados); i++ {
			eliminado := data.Eliminados[i]
			if err := tx.Where("id = ?", eliminado).Delete(&models.Departamento{}).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los departamentos",
						err.Error(),
					}}
				c.JSON(500, rpta)
			}
		}
		tx.Commit()
		if tx.Error != nil {
			tx.Rollback()
			tx.Close()
			c.JSON(500, gin.H{
				"tipo_mensaje": "error",
				"mensaje": []interface{}{
					"No se ha podido guardar los departamentos",
					nuevosId},
			})
		} else {
			tx.Close()
			c.JSON(200, gin.H{
				"tipo_mensaje": "success",
				"mensaje": []interface{}{
					"Se ha registrado los cambios en los departamentos",
					nuevosId},
			})
		}
	}
}
