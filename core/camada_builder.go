package core

type CamadaBuilder struct {
	camada *Camada
}

func NewCamadaBuilder(id int32) CamadaBuilder {
	return CamadaBuilder{&Camada{
		Id: id,
	}}
}

func (b CamadaBuilder) Nome(nome string) CamadaBuilder {
	b.camada.Nome = nome
	return b
}

func (b CamadaBuilder) Wkt(wkt string) CamadaBuilder {
	b.camada.Wkt = wkt
	return b
}

func (b CamadaBuilder) Ativo(ativo *Ativo) CamadaBuilder {
	b.camada.Ativo = ativo
	return b
}

func (b CamadaBuilder) AddCategoria(categoria CamadaCategoria) CamadaBuilder {
	b.camada.Categoria = append(b.camada.Categoria, categoria)
	return b
}

func (b CamadaBuilder) Build() Camada {
	return *b.camada
}
