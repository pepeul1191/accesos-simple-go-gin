package structs

type PermissionEditadoStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PermissionNuevoStruct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TablePermissionStruct struct {
	Nuevos     []PermissionNuevoStruct   `json:"nuevos"`
	Editados   []PermissionEditadoStruct `json:"editados"`
	Eliminados []int                     `json:"eliminados"`
}
