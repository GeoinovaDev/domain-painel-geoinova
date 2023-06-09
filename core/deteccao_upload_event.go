package core

import "time"

type DeteccaoAntesUploadEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

type DeteccaoDepoisUploadEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

type DeteccaoResultadoUploadEvent struct {
	CreatedAt time.Time `json:"created_at"`
	Deteccao  Deteccao  `json:"deteccao"`
}

func NewDeteccaoAntesUploadEvent(deteccao Deteccao) DeteccaoAntesUploadEvent {
	return DeteccaoAntesUploadEvent{time.Now(), deteccao}
}

func (e DeteccaoAntesUploadEvent) Name() string {
	return DETECCAO_ANTES_UPLOAD_EVENT
}

func NewDeteccaoDepoisUploadEvent(deteccao Deteccao) DeteccaoDepoisUploadEvent {
	return DeteccaoDepoisUploadEvent{time.Now(), deteccao}
}

func (e DeteccaoDepoisUploadEvent) Name() string {
	return DETECCAO_DEPOIS_UPLOAD_EVENT
}

func NewDeteccaoResultadoUploadEvent(deteccao Deteccao) DeteccaoResultadoUploadEvent {
	return DeteccaoResultadoUploadEvent{time.Now(), deteccao}
}

func (e DeteccaoResultadoUploadEvent) Name() string {
	return DETECCAO_RESULTADO_UPLOAD_EVENT
}
