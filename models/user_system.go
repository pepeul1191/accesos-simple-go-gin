package models

import (
	"encoding/json"
)

type UserSystem struct {
	ID       int `gorm:"primary_key" json:"id"`
	SystemId int `gorm:"column:system_id" json:"system_id"`
	UserId   int `gorm:"column:user_id" json:"user_id"`
}

func (UserSystem) TableName() string {
	return "users_systems"
}

func (userSystem *UserSystem) ToJSON() ([]byte, error) {
	return json.Marshal(userSystem)
}
