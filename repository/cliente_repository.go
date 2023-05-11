package repository

import (
	"github.com/GeoinovaDev/domain-painel-geoinova/app"
	"github.com/GeoinovaDev/domain-painel-geoinova/core"
)

type ClienteRepository interface {
	FetchAll(...app.IQueryOption) ([]core.Cliente, error)
}
