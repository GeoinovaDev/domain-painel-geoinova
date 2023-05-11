package models

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type AnaliseRequest struct {
	Id           int32           `json:"id"`
	Wkt          string          `json:"wkt,omitempty"`
	Ativo        core.Ativo      `json:"ativo,omitempty"`
	Status       int             `json:"status,omitempty"`
	Tipo         string          `json:"tipo,omitempty"`
	TotalArea    float32         `json:"totalArea,omitempty"`
	ImagemAntes  core.Imagem     `json:"imagemAntes,omitempty"`
	ImagemDepois core.Imagem     `json:"imagemDepois,omitempty"`
	Deteccoes    []core.Deteccao `json:"deteccoes,omitempty"`
	Camadas      []core.Camada   `json:"camadas,omitempty"`
}

func (a AnaliseRequest) ToAnalise() core.Analise {
	return core.Analise{
		Id:           a.Id,
		Wkt:          a.Wkt,
		Ativo:        a.Ativo,
		Status:       a.Status,
		Tipo:         a.Tipo,
		TotalArea:    a.TotalArea,
		ImagemAntes:  a.ImagemAntes,
		ImagemDepois: a.ImagemDepois,
		Deteccoes:    a.Deteccoes,
		Camadas:      a.Camadas,
	}
}
