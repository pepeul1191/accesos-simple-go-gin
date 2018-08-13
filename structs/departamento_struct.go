package structs

type DepartamentoEditadoStruct struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type DepartamentoNuevoStruct struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
}

type TableDepartamentoStruct struct {
	Nuevos     []DepartamentoNuevoStruct   `json:"nuevos"`
	Editados   []DepartamentoEditadoStruct `json:"editados"`
	Eliminados []string                    `json:"eliminados"`
}
