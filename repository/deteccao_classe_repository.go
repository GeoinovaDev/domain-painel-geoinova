package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type DeteccaoClasseRepository interface {
	FetchAll() ([]core.DeteccaoClasse, error)
}
