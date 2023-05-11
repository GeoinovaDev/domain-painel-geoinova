package core

import "time"

type DeteccoesDeleteEvent struct {
	CreatedAt time.Time  `json:"created_at"`
	Deteccoes []Deteccao `json:"deteccoes"`
}

func NewDeteccoesDeleteEvent(deteccoes []Deteccao) DeteccoesDeleteEvent {
	return DeteccoesDeleteEvent{time.Now(), deteccoes}
}

func (e DeteccoesDeleteEvent) Name() string {
	return DETECCOES_DELETE_EVENT
}
