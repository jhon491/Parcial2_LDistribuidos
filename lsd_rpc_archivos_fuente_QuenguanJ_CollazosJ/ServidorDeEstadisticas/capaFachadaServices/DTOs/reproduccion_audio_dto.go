package dtos

type ReproduccionAudioDTO struct {
	TituloAudio string `json:"titulo_audio"`
	TipoAudio   string `json:"tipo_audio"`
	FechaHora   string `json:"fecha_hora"`
}
