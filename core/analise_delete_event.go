package core

import "time"

type AnaliseDeleteEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Analise   Analise   `json:"analise"`
}

func NewAnaliseDeleteEvent(analise Analise) AnaliseDeleteEvent {
	return AnaliseDeleteEvent{time.Now(), analise}
}

func (e AnaliseDeleteEvent) Name() string {
	return ANALISE_UPDATE_EVENT
}
