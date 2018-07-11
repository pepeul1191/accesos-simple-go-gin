package models

import (
	"encoding/json"
	"strconv"
)

type Provincia struct {
	ID              int    `gorm:"primary_key" json:"id"`
	Nombre          string `gorm:"column:nombre" json:"nombre"`
	Departamento_id int    `gorm:"column:departamento_id" json:"departamento_id,omitempty"`
}

func (Provincia) TableName() string {
	return "provincias"
}

func (provincia *Provincia) ToJSON() ([]byte, error) {
	return json.Marshal(provincia)
}

func (provincia *Provincia) ToString() string {
	return "Provincia : { id : " + strconv.Itoa(provincia.ID) +
		", nombre : '" + provincia.Nombre + "' }"
}
