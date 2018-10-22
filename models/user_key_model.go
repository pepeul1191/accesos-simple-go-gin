package models

import (
	"encoding/json"
)

type UserKey struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Activation string `gorm:"column:activation" json:"activation"`
	Reset      string `gorm:"column:reset" json:"reset"`
	UserId     int    `gorm:"column:user_id" json:"user_id"`
}

func (UserKey) TableName() string {
	return "users_keys"
}

func (userKey *UserKey) ToJSON() ([]byte, error) {
	return json.Marshal(userKey)
}
