package structs

type UserCreateStruct struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
	SystemId int    `json:"system_id"`
}

type UserKeyCreatedStruct struct {
	UserId        int    `json:"user_id"`
	ActivationKey string `json:"activation_key"`
}
