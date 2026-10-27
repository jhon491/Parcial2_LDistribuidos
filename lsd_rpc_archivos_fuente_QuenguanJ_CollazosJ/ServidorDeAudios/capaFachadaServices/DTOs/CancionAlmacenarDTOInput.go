package dtos

type CancionAlmacenarDTOInput struct {
	IdTipo string `json:"idTipo"` // 1: Musica, 2:Podcast, 3:Audiolibro, 4: RuidoBlanco
	Id     string `json:"id"`     // id del audio dentro de su tipo
}
