package models

import (
	"encoding/json"
)

type User struct {
	ID            int    `gorm:"primary_key" json:"id"`
	User          string `gorm:"column:user" json:"user"`
	Pass          string `gorm:"column:pass" json:"pass"`
	Email         string `gorm:"column:email" json:"email"`
	Count         string `gorm:"column:count"`
	User_state_id int    `gorm:"column:user_state_id" json:"user_state_id,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) ToJSON() ([]byte, error) {
	return json.Marshal(user)
}
