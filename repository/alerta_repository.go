package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type AlertaRepository interface {
	Insert(core.Alerta) (core.Alerta, error)
}
