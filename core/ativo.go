package core

type Ativo struct {
	Id       int32     `json:"id"`
	Nome     string    `json:"nome,omitempty"`
	Wkt      string    `json:"wkt,omitempty"`
	SubAtivo *SubAtivo `json:"subAtivo,omitempty"`
	Cliente  *Cliente  `json:"cliente,omitempty"`
}

func NewAtivo(id int32, nome string) Ativo {
	ativo := Ativo{
		Id:   id,
		Nome: nome,
	}

	return ativo
}
