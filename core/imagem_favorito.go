package core

type ImagemFavorito struct {
	Id      string
	Usuario Usuario
	Imagem  Imagem
}

func NewImagemFavorito(id string, usuario Usuario, imagem Imagem) ImagemFavorito {
	return ImagemFavorito{id, usuario, imagem}
}
