package models

import (
	"encoding/json"
)

type UserStateSystem struct {
	ID       int    `gorm:"primary_key" json:"id"`
	SystemId int    `gorm:"column:system_id" json:"system_id"`
	System   string `gorm:"column:system" json:"system"`
	StateId  int    `gorm:"column:user_state_id" json:"user_state_id"`
	State    string `gorm:"column:state" json:"state"`
	UserId   string `gorm:"column:user_id" json:"user_id"`
	User     string `gorm:"column:user" json:"user"`
	Pass     string `gorm:"column:pass" json:"pass"`
}

func (UserStateSystem) TableName() string {
	return "vw_users_states_systems"
}

func (userStateSystem *UserStateSystem) ToJSON() ([]byte, error) {
	return json.Marshal(userStateSystem)
}
