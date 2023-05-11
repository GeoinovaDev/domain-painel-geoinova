package core

type Analise struct {
	Id           int32      `json:"id"`
	Uuid         string     `json:"uuid,omitempty"`
	Wkt          string     `json:"wkt,omitempty"`
	Usuario      Usuario    `json:"usuario,omitempty"`
	Cliente      *Cliente   `json:"cliente,omitempty"`
	Ativo        Ativo      `json:"ativo,omitempty"`
	Status       int        `json:"status,omitempty"`
	Tipo         string     `json:"tipo,omitempty"`
	TotalArea    float32    `json:"totalArea,omitempty"`
	ImagemAntes  Imagem     `json:"imagemAntes,omitempty"`
	ImagemDepois Imagem     `json:"imagemDepois,omitempty"`
	Deteccoes    []Deteccao `json:"deteccoes,omitempty"`
	Camadas      []Camada   `json:"camadas,omitempty"`
}
