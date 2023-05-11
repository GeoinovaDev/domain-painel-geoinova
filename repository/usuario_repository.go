package repository

import "github.com/GeoinovaDev/domain-painel-geoinova/core"

type UsuarioRepository interface {
	Fetch(int32) (core.Usuario, error)
}
