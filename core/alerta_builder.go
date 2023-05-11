package core

type AlertaBuilder struct {
	alerta *Alerta
}

func NewAlertaBuilder(id int32) AlertaBuilder {
	return AlertaBuilder{
		alerta: &Alerta{
			Id: id,
		},
	}
}

func (b AlertaBuilder) Uuid(uuid string) AlertaBuilder {
	b.alerta.Uuid = uuid
	return b
}

func (b AlertaBuilder) Order(order uint) AlertaBuilder {
	b.alerta.Order = order
	return b
}

func (b AlertaBuilder) Observacao(obs string) AlertaBuilder {
	b.alerta.Observacao = obs
	return b
}

func (b AlertaBuilder) Responsavel(responsavel string) AlertaBuilder {
	b.alerta.Responsavel = responsavel
	return b
}

func (b AlertaBuilder) Justificativa(justificativa string) AlertaBuilder {
	b.alerta.Justificativa = justificativa
	return b
}

func (b AlertaBuilder) Status(status string) AlertaBuilder {
	b.alerta.Status = status
	return b
}

func (b AlertaBuilder) Deteccao(deteccao Deteccao) AlertaBuilder {
	b.alerta.Deteccao = deteccao
	return b
}

func (b AlertaBuilder) Analise(analise Analise) AlertaBuilder {
	b.alerta.Analise = analise
	return b
}

func (b AlertaBuilder) Build() Alerta {
	return *b.alerta
}
