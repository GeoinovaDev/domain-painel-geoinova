package repository

import (
	"github.com/GeoinovaDev/domain-painel-geoinova/app"
	"github.com/GeoinovaDev/domain-painel-geoinova/core"
)

type CamadaRepository interface {
	FetchByAtivoId(int32, ...app.IQueryOption) ([]core.Camada, error)
	FetchWkt(int32) (string, error)
}
