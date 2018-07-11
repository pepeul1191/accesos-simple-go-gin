package models

import (
	"encoding/json"
	"strconv"
)

type Distrito struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Nombre       string `gorm:"column:nombre" json:"nombre"`
	Provincia_id int    `gorm:"column:provincia_id" json:"provincia_id,omitempty"`
}

func (Distrito) TableName() string {
	return "distritos"
}

func (distrito *Distrito) ToJSON() ([]byte, error) {
	return json.Marshal(distrito)
}

func (distrito *Distrito) ToString() string {
	return "Distrito : { id : " + strconv.Itoa(distrito.ID) +
		", nombre : '" + distrito.Nombre + "' }"
}
