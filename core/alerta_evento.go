package core

import "time"

const (
	ALERTA_EVENTO_CRIADO = "CR"
)

type AlertaEvento struct {
	Id       int32     `json:"id"`
	Status   string    `json:"status"`
	Param    string    `json:"param"`
	DateTime time.Time `json:"datetime"`
}

func NewAlertaEvento(status string, param string) AlertaEvento {
	return AlertaEvento{
		Status:   status,
		Param:    param,
		DateTime: time.Now(),
	}
}
