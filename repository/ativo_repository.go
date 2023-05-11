package repository

import (
	"github.com/GeoinovaDev/domain-painel-geoinova/app"
	"github.com/GeoinovaDev/domain-painel-geoinova/core"
)

type AtivoRepository interface {
	FetchByUuid(string, ...app.IQueryOption) (core.Ativo, error)
	FetchById(int32, ...app.IQueryOption) (core.Ativo, error)
	FetchByClienteId(int32, ...app.IQueryOption) ([]core.Ativo, error)
	FetchWktByAtivoId(int) (string, error)
}
