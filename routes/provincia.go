package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swp/access/configs"
	"github.com/swp/access/models"
	"github.com/swp/access/structs"
)

func ProvinciaListar(c *gin.Context) {
	departamento_id, err := strconv.ParseInt(c.Param("departamento_id"), 10, 64)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"El parámetro enviado no se pudo parsear a entero",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var provincias []models.Provincia
		db := configs.Database()
		if err := db.Where("departamento_id = ?", departamento_id).Select("id, nombre").Find(&provincias).Error; err != nil {
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido listar las provincias del departamento",
					err.Error(),
				}}
			c.JSON(500, rpta)
		} else {
			defer db.Close()
			c.JSON(200, provincias)
		}
	}
}
