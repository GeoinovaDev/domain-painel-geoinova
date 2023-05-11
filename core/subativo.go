package core

type SubAtivo struct {
	Id   int32  `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
}

func NewSubAtivo(id int32, nome string) *SubAtivo {
	return &SubAtivo{id, nome}
}
