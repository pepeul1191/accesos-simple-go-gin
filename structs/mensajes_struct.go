package structs

type Error struct {
	TipoMensaje string   `json:"tipo_mensaje"`
	Mensaje     []string `json:"mensaje"`
}

type IdsNuevosTemp struct {
	Temporal string `json:"temporal"`
	NuevoId  int    `json:"nuevo_id"`
}
