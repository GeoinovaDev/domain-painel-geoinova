package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type AnaliseRepository interface {
	Insert(core.Analise) (core.Analise, error)
	Update(core.Analise) (core.Analise, error)
	Delete(core.Analise) error
	Fetch(id int32) (core.Analise, error)
	FetchWkt(analiseId int32) (string, error)
	FetchByAtivoId(ativoId int32) ([]core.Analise, error)
}
