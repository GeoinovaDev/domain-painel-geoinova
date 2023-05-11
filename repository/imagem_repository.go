package repository

import (
	"github.com/GeoinovaDev/domain-painel-geoinova/app"
	"github.com/GeoinovaDev/domain-painel-geoinova/core"
)

type ImagemRepository interface {
	FetchByAtivoId(int32, ...app.IQueryOption) ([]core.Imagem, error)
}
