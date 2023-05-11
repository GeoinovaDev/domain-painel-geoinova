package core

type AtivoBuilder struct {
	ativo *Ativo
}

func NewAtivoBuilder(id int32, nome string) AtivoBuilder {
	ativo := NewAtivo(id, nome)
	return AtivoBuilder{&ativo}
}

func (b AtivoBuilder) Wkt(wkt string) AtivoBuilder {
	b.ativo.Wkt = wkt
	return b
}

func (b AtivoBuilder) Cliente(cliente *Cliente) AtivoBuilder {
	b.ativo.Cliente = cliente
	return b
}

func (b AtivoBuilder) SubAtivo(subAtivo *SubAtivo) AtivoBuilder {
	b.ativo.SubAtivo = subAtivo
	return b
}

func (b AtivoBuilder) Build() Ativo {
	return *b.ativo
}
