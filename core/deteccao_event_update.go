package core

import "time"

type DeteccoesUpdateEvent struct {
	CreatedAt time.Time  `json:"created_at"`
	Deteccoes []Deteccao `json:"deteccoes"`
}

func NewDeteccoesUpdateEvent(deteccoes []Deteccao) DeteccoesUpdateEvent {
	return DeteccoesUpdateEvent{time.Now(), deteccoes}
}

func (e DeteccoesUpdateEvent) Name() string {
	return DETECCOES_UPDATE_EVENT
}
