package models

import (
	"encoding/json"
	"strconv"
)

type DepartamentoProvinciaDistrito struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Nombre string `gorm:"column:nombre" json:"nombre"`
}

func (DepartamentoProvinciaDistrito) TableName() string {
	return "vw_distrito_provincia_departamento"
}

func (distrito *DepartamentoProvinciaDistrito) ToJSON() ([]byte, error) {
	return json.Marshal(distrito)
}

func (distrito *DepartamentoProvinciaDistrito) ToString() string {
	return "DepartamentoProvinciaDistrito : { id : " + strconv.Itoa(distrito.ID) +
		", nombre : '" + distrito.Nombre + "' }"
}
