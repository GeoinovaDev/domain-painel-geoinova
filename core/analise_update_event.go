package core

import "time"

type AnaliseUpdateEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Analise   Analise   `json:"analise"`
}

func NewAnaliseUpdateEvent(analise Analise) AnaliseUpdateEvent {
	return AnaliseUpdateEvent{time.Now(), analise}
}

func (e AnaliseUpdateEvent) Name() string {
	return ANALISE_UPDATE_EVENT
}
