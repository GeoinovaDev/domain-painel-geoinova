package core

type DeteccaoBuilder struct {
	deteccao *Deteccao
}

func NewDeteccaoBuilder(id int32) DeteccaoBuilder {
	return DeteccaoBuilder{
		deteccao: &Deteccao{
			Id: id,
		},
	}
}

func (b DeteccaoBuilder) Wkt(wkt string) DeteccaoBuilder {
	b.deteccao.Wkt = wkt
	return b
}

func (b DeteccaoBuilder) Classe(classe DeteccaoClasse) DeteccaoBuilder {
	b.deteccao.Classe = classe
	return b
}

func (b DeteccaoBuilder) Descricao(descricao string) DeteccaoBuilder {
	b.deteccao.Descricao = descricao
	return b
}

func (b DeteccaoBuilder) ImagemAntesUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemAntesUuid = uuid
	return b
}

func (b DeteccaoBuilder) ImagemDepoisUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemDepoisUuid = uuid
	return b
}

func (b DeteccaoBuilder) ImagemResultadoUuid(uuid string) DeteccaoBuilder {
	b.deteccao.ImagemResultadoUuid = uuid
	return b
}

func (b DeteccaoBuilder) AddCamada(camada Camada) DeteccaoBuilder {
	b.deteccao.Camadas = append(b.deteccao.Camadas, camada)
	return b
}

func (b DeteccaoBuilder) Build() Deteccao {
	return *b.deteccao
}
