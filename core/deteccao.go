package core

type Deteccao struct {
	Id                  int32          `json:"id,omitempty"`
	Wkt                 string         `json:"wkt,omitempty"`
	Classe              DeteccaoClasse `json:"classe,omitempty"`
	Descricao           string         `json:"descricao,omitempty"`
	Camadas             []Camada       `json:"camadas,omitempty"`
	ImagemAntesUuid     string         `json:"ImagemAntesUuid"`
	ImagemDepoisUuid    string         `json:"ImagemDepoisUuid"`
	ImagemResultadoUuid string         `json:"ImagemResultadoUuid"`
}
