package core

type AnaliseBuilder struct {
	analise *Analise
}

func NewAnaliseBuilder(id int32) AnaliseBuilder {
	return AnaliseBuilder{
		analise: &Analise{
			Id: id,
		},
	}
}

func (b AnaliseBuilder) Cliente(cliente *Cliente) AnaliseBuilder {
	b.analise.Cliente = cliente
	return b
}

func (b AnaliseBuilder) Wkt(wkt string) AnaliseBuilder {
	b.analise.Wkt = wkt
	return b
}

func (b AnaliseBuilder) Uuid(uuid string) AnaliseBuilder {
	b.analise.Uuid = uuid
	return b
}

func (b AnaliseBuilder) Status(status int) AnaliseBuilder {
	b.analise.Status = status
	return b
}

func (b AnaliseBuilder) Tipo(tipo string) AnaliseBuilder {
	b.analise.Tipo = tipo
	return b
}

func (b AnaliseBuilder) ImagemAntes(imagemAntes Imagem) AnaliseBuilder {
	b.analise.ImagemAntes = imagemAntes
	return b
}

func (b AnaliseBuilder) ImagemDepois(imagemDepois Imagem) AnaliseBuilder {
	b.analise.ImagemDepois = imagemDepois
	return b
}

func (b AnaliseBuilder) AddDeteccao(deteccao Deteccao) AnaliseBuilder {
	b.analise.Deteccoes = append(b.analise.Deteccoes, deteccao)
	return b
}

func (b AnaliseBuilder) Ativo(ativo Ativo) AnaliseBuilder {
	b.analise.Ativo = ativo
	return b
}

func (b AlertaBuilder) AddEvento(evento AlertaEvento) AlertaBuilder {
	b.alerta.Eventos = append(b.alerta.Eventos, evento)
	return b
}

func (b AlertaBuilder) AddFase(fase AlertaFase) AlertaBuilder {
	b.alerta.Fases = append(b.alerta.Fases, fase)
	return b
}

func (b AnaliseBuilder) Build() Analise {
	return *b.analise
}
