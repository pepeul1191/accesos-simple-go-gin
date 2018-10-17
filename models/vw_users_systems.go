package models

import (
	"encoding/json"
)

type UserSystem struct {
	ID       int    `gorm:"primary_key" json:"id"`
	SystemId int    `gorm:"column:system_id" json:"system_id"`
	System   string `gorm:"column:system" json:"system"`
	UserId   string `gorm:"column:user_id" json:"user_id"`
	User     string `gorm:"column:user" json:"user"`
	Pass     string `gorm:"column:pass" json:"pass"`
}

func (UserSystem) TableName() string {
	return "vw_users_systems"
}

func (userSystem *UserSystem) ToJSON() ([]byte, error) {
	return json.Marshal(userSystem)
}
