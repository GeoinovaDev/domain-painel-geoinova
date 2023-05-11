package core

type Cliente struct {
	Id   int32  `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
}

func NewCliente(id int32, nome string) *Cliente {
	return &Cliente{id, nome}
}
