package core

type ImagemBuilder struct {
	imagem *Imagem
}

func NewImagemBuilder(id int32) ImagemBuilder {
	return ImagemBuilder{
		imagem: &Imagem{
			Id: id,
		},
	}
}

func (b ImagemBuilder) TotalBanda(totalBandas int) ImagemBuilder {
	b.imagem.TotalBandas = totalBandas
	return b
}

func (b ImagemBuilder) Uuid(uuid string) ImagemBuilder {
	b.imagem.Uuid = uuid
	return b
}

func (b ImagemBuilder) Data(data string) ImagemBuilder {
	b.imagem.Data = data
	return b
}

func (b ImagemBuilder) Bucket(bucket string) ImagemBuilder {
	b.imagem.Bucket = bucket
	return b
}

func (b ImagemBuilder) Resolucao(resolucao string) ImagemBuilder {
	b.imagem.Resolucao = resolucao
	return b
}

func (b ImagemBuilder) Provedor(provedor string) ImagemBuilder {
	b.imagem.Provedor = provedor
	return b
}

func (b ImagemBuilder) TaxaCobertura(tx float32) ImagemBuilder {
	b.imagem.TaxaCobertura = tx
	return b
}

func (b ImagemBuilder) ProvedorFonte(fonte string) ImagemBuilder {
	b.imagem.ProvedorFonte = fonte
	return b
}

func (b ImagemBuilder) Qualidade(qualidade ImagemQualidade) ImagemBuilder {
	b.imagem.Qualidade = qualidade
	return b
}

func (b ImagemBuilder) Build() Imagem {
	return *b.imagem
}
