package core

type ImagemQualidade struct {
	Id         int32  `json:"id"`
	Nome       string `json:"nome"`
	Color      string `json:"color"`
	Background string `json:"background"`
	Tipo       int    `json:"tipo"`
}

func NewImagemQualidade(id int32, nome string, color string, background string, tipo int) ImagemQualidade {
	return ImagemQualidade{id, nome, color, background, tipo}
}
