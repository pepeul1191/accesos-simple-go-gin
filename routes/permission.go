package routes

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func PermissionSave(c *gin.Context) {
	var postData string = c.PostForm("data")
	data := &structs.TablePermissionStruct{}
	err := json.Unmarshal([]byte(postData), data)
	var nuevosId []structs.IdsNuevosTemp
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha guardar los permisos, error de parseo del JSON",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		tx := configs.Database().Begin()
		for i := 0; i < len(data.Nuevos); i++ {
			var nuevo = models.Permission{
				Name: data.Nuevos[i].Name}
			if err := tx.Create(&nuevo).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los permisos",
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
			var permisos []models.Permission
			if err := tx.Model(&permisos).Where("id = ?", editado.Id).Update(
				"Name", editado.Name).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los permisos",
						err.Error(),
					}}
				c.JSON(500, rpta)
			}
		}
		for i := 0; i < len(data.Eliminados); i++ {
			eliminado := data.Eliminados[i]
			if err := tx.Where("id = ?", eliminado).Delete(&models.Permission{}).Error; err != nil {
				tx.Rollback()
				rpta := structs.Error{
					TipoMensaje: "error",
					Mensaje: []string{
						"No se ha podido guardar los permisos",
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
					"No se ha podido guardar los permisos",
					nuevosId},
			})
		} else {
			tx.Close()
			if len(nuevosId) == 0 {
				c.JSON(200, gin.H{
					"tipo_mensaje": "success",
					"mensaje": []interface{}{
						"Se ha registrado los cambios en los permisos",
					},
				})
			} else {
				c.JSON(200, gin.H{
					"tipo_mensaje": "success",
					"mensaje": []interface{}{
						"Se ha registrado los cambios en los permisos",
						nuevosId},
				})
			}
		}
	}
}
