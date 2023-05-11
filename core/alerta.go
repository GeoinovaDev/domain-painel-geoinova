package core

const (
	ALERTA_STATUS_ABERTA = "AB"
)

type Alerta struct {
	Id            int32          `json:"id,omitempty"`
	Uuid          string         `json:"uuid,omitempty"`
	Order         uint           `json:"order,omitempty"`
	Observacao    string         `json:"observacao,omitempty"`
	Responsavel   string         `json:"responsavel,omitempty"`
	Justificativa string         `json:"justificativa,omitempty"`
	Status        string         `json:"status,omitempty"`
	Deteccao      Deteccao       `json:"deteccao,omitempty"`
	Analise       Analise        `json:"analise,omitempty"`
	Eventos       []AlertaEvento `json:"eventos"`
	Fases         []AlertaFase   `json:"fases"`
}
