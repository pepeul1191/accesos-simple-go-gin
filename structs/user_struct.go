package structs

type UserCreateStruct struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
	SystemId int    `json:"system_id"`
}

type UserGetStruct struct {
	User        string `json:"user"`
	Email       string `json:"email"`
	UserStateId int    `json:"state_id"`
}

type UserUpdateStruct struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	UserStateId string `json:"user_state_id"`
}
