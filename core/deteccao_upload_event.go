package core

import "time"

type DeteccaoAntesUpdateEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

type DeteccaoDepoisUpdateEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

type DeteccaoResultadoUpdateEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

func NewDeteccaoAntesUpdateEvent(deteccao Deteccao) DeteccaoAntesUpdateEvent {
	return DeteccaoAntesUpdateEvent{time.Now(), deteccao}
}

func (e DeteccaoAntesUpdateEvent) Name() string {
	return DETECCOES_DELETE_EVENT
}

func NewDeteccaoDepoisUpdateEvent(deteccao Deteccao) DeteccaoDepoisUpdateEvent {
	return DeteccaoDepoisUpdateEvent{time.Now(), deteccao}
}

func (e DeteccaoDepoisUpdateEvent) Name() string {
	return DETECCOES_DELETE_EVENT
}

func NewDeteccaoResultadoUpdateEvent(deteccao Deteccao) DeteccaoResultadoUpdateEvent {
	return DeteccaoResultadoUpdateEvent{time.Now(), deteccao}
}

func (e DeteccaoResultadoUpdateEvent) Name() string {
	return DETECCOES_DELETE_EVENT
}
