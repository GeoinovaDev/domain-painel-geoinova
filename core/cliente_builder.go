package core

type ClienteBuilder struct {
	cliente *Cliente
}

func NewClienteBuilder(id int32, nome string) ClienteBuilder {
	return ClienteBuilder{NewCliente(id, nome)}
}

func (b ClienteBuilder) WithStatus(status string) ClienteBuilder {
	b.cliente.Status = status

	return b
}

func (b ClienteBuilder) WithDescricao(descricao string) ClienteBuilder {
	b.cliente.Descricao = descricao

	return b
}

func (b ClienteBuilder) Build() *Cliente {
	return b.cliente
}
