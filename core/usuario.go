package core

const (
	UsuarioStatusAtivo      = 1
	UsuarioStatusDesativado = 0
)

type Usuario struct {
	Id      int32   `json:"id"`
	Nome    string  `json:"nome"`
	Email   string  `json:"email"`
	Senha   string  `json:"-"`
	Status  int     `json:"-"`
	Cliente Cliente `json:"-"`
}

func NewUsuario(id int32, nome string) Usuario {
	usuario := Usuario{
		Id:   id,
		Nome: nome,
	}

	return usuario
}
