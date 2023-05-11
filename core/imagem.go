package core

type Imagem struct {
	Id            int32           `json:"id"`
	Uuid          string          `json:"uuid"`
	Data          string          `json:"data"`
	Bucket        string          `json:"bucket"`
	Resolucao     string          `json:"resolucao"`
	TotalBandas   int             `json:"totalBandas"`
	TaxaCobertura float32         `json:"taxaCobertura"`
	Provedor      string          `json:"provedor"`
	ProvedorFonte string          `json:"provedorFonte"`
	Qualidade     ImagemQualidade `json:"qualidade"`
}

func NewImagem(id int32) Imagem {
	return Imagem{
		Id: id,
	}
}
