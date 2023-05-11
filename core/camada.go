package core

type CamadaCategoria struct {
	Id    int32  `json:"id"`
	Nome  string `json:"nome,omitempty"`
	Cor   string `json:"cor,omitempty"`
	Borda string `json:"borda,omitempty"`
}

type Camada struct {
	Id        int32             `json:"id"`
	Nome      string            `json:"nome,omitempty"`
	Wkt       string            `json:"wkt,omitempty"`
	Ativo     *Ativo            `json:"ativo,omitempty"`
	Categoria []CamadaCategoria `json:"categorias,omitempty"`
}

func NewCamada(id int32) Camada {
	return Camada{}
}

func NewCamadaCategoria(id int32, nome string, cor string, borda string) CamadaCategoria {
	return CamadaCategoria{id, nome, cor, borda}
}
