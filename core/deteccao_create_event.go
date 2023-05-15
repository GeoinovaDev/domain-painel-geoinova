package core

import "time"

type DeteccoesCreateEvent struct {
	CreatedAt time.Time  `json:"created_at"`
	Deteccoes []Deteccao `json:"deteccoes"`
}

func NewDeteccoesCreateEvent(deteccoes []Deteccao) DeteccoesCreateEvent {
	return DeteccoesCreateEvent{time.Now(), deteccoes}
}

func (e DeteccoesCreateEvent) Name() string {
	return DETECCOES_CREATE_EVENT
}
