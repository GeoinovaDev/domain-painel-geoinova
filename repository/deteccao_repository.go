package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type DeteccaoRepository interface {
	UpdateBucket(core.Deteccao) error
	FetchByAnaliseId(int32) ([]core.Deteccao, error)
}
