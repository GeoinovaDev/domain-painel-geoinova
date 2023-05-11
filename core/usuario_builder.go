package core

type UsuarioBuilder struct {
	usuario *Usuario
}

func NewUsuarioBuilder(id int32, nome string) UsuarioBuilder {
	usuario := NewUsuario(id, nome)

	return UsuarioBuilder{
		usuario: &usuario,
	}
}

func (u UsuarioBuilder) Email(email string) UsuarioBuilder {
	u.usuario.Email = email
	return u
}

func (u UsuarioBuilder) Build() Usuario {
	return *u.usuario
}
