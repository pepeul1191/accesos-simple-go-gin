package structs

type KeyActivationStruct struct {
	UserId        int    `json:"user_id"`
	ActivationKey string `json:"activation_key"`
}
