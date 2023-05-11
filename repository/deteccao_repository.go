package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type DeteccaoRepository interface {
	UpdateBucket(int32, string, string, string) error
	FetchByAnaliseId(int32) ([]core.Deteccao, error)
}
