package structs

type KeyActivationStruct struct {
	UserId        int    `json:"user_id"`
	ActivationKey string `json:"activation_key"`
}

type KeyResetStruct struct {
	UserId   int    `json:"user_id"`
	ResetKey string `json:"reset_key"`
}
