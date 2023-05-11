package core

import "time"

const (
	ALERTA_FASE_ABERTA = "AB"
)

type AlertaFase struct {
	Id       int32     `json:"id"`
	Status   string    `json:"status"`
	DateTime time.Time `json:"datetime"`
}

func NewAlertaFase(status string) AlertaFase {
	return AlertaFase{
		Status:   status,
		DateTime: time.Now(),
	}
}
