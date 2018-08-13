package structs

type Error struct {
	TipoMensaje string   `json:"tipo_mensaje"`
	Mensaje     []string `json:"mensaje"`
}
