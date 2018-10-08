package models

import (
	"encoding/json"
	"strconv"
)

type UserState struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (UserState) TableName() string {
	return "user_states"
}

func (userState *UserState) ToJSON() ([]byte, error) {
	return json.Marshal(userState)
}

func (userState *UserState) ToString() string {
	return "UserState : { id : " + strconv.Itoa(userState.ID) +
		", name : '" + userState.Name + "' }"
}
