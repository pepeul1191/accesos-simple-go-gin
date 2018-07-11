package routes

import (
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
