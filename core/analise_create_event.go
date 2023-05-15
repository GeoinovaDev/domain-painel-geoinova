package core

import "time"

type AnaliseCreateEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Analise   Analise   `json:"analise"`
}

func NewAnaliseCreateEvent(analise Analise) AnaliseCreateEvent {
	return AnaliseCreateEvent{time.Now(), analise}
}

func (e AnaliseCreateEvent) Name() string {
	return ANALISE_CREATE_EVENT
}
