package core

type Cliente struct {
	Id        int32  `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Descricao string `json:"descricao"`
	Status    string `json:"status"`
}

func NewCliente(id int32, nome string) *Cliente {
	return &Cliente{id, nome, "", ""}
}
