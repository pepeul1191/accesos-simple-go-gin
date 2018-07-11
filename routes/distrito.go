package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ginv2/configs"
	"github.com/ginv2/models"
	"github.com/ginv2/structs"
)

func DistritoListar(c *gin.Context) {
	provincia_id, err := strconv.ParseInt(c.Param("provincia_id"), 10, 64)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"El parámetro enviado no se pudo parsear a entero",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var distritos []models.Distrito
		if err := configs.Database().Where("provincia_id = ?", provincia_id).Select("id, nombre").Find(&distritos).Error; err != nil {
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido listar los distritos de la provincia",
					err.Error(),
				}}
			c.JSON(500, rpta)
		} else {
			c.JSON(200, distritos)
		}
	}
}

func DistritoBuscar(c *gin.Context) {
	var nombre string = c.Query("nombre")
	var distritos []models.DepartamentoProvinciaDistrito
	if err := configs.Database().Limit(10).Where("nombre LIKE ?", nombre+"%").Find(&distritos).Error; err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"No se ha podido buscar el nombre del distrito por id",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		c.JSON(200, distritos)
	}
}

func DistritoNombre(c *gin.Context) {
	distrito_id, err := strconv.ParseInt(c.Param("distrito_id"), 10, 64)
	if err != nil {
		rpta := structs.Error{
			TipoMensaje: "error",
			Mensaje: []string{
				"El parámetro enviado no se pudo parsear a entero",
				err.Error(),
			}}
		c.JSON(500, rpta)
	} else {
		var distrito models.DepartamentoProvinciaDistrito
		if err := configs.Database().Where("id = ?", distrito_id).First(&distrito).Error; err != nil {
			rpta := structs.Error{
				TipoMensaje: "error",
				Mensaje: []string{
					"No se ha podido buscar las coincidencias de nombre de distrito",
					err.Error(),
				}}
			c.JSON(500, rpta)
		} else {
			c.String(200, distrito.Nombre)
		}
	}
}
