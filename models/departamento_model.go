package models

import (
	"encoding/json"
	"strconv"
)

type Departamento struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Nombre string `gorm:"column:nombre" json:"nombre"`
}

func (Departamento) TableName() string {
	return "departamentos"
}

func (departamento *Departamento) ToJSON() ([]byte, error) {
	return json.Marshal(departamento)
}

func (departamento *Departamento) ToString() string {
	return "Departamento : { id : " + strconv.Itoa(departamento.ID) +
		", nombre : '" + departamento.Nombre + "' }"
}
