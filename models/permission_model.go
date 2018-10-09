package models

import (
	"encoding/json"
)

type Permission struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (permission *Permission) ToJSON() ([]byte, error) {
	return json.Marshal(permission)
}
