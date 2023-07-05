package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type DeteccaoRepository interface {
	Inserir(int, core.Deteccao) error
	Update(core.Deteccao) error
	Delete(core.Deteccao) error
	UpdateBucket(core.Deteccao) error
	UpdateAntesBucket(core.Deteccao) error
	UpdateDepoisBucket(core.Deteccao) error
	UpdateResultadoBucket(core.Deteccao) error
	FetchByAnaliseId(int32) ([]core.Deteccao, error)
}
