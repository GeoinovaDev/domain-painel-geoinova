package core

type ClienteBuilder struct {
	cliente *Cliente
}

func NewClienteBuilder(id int32, nome string) ClienteBuilder {
	return ClienteBuilder{NewCliente(id, nome)}
}

func (b ClienteBuilder) Build() *Cliente {
	return b.cliente
}
