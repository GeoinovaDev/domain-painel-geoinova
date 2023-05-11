package core

type DeteccaoClasse struct {
	Id    int32  `json:"id"`
	Nome  string `json:"nome"`
	Tipo  string `json:"tipo"`
	Color string `json:"color"`
}

func NewDeteccaoClasse(id int32, nome string, tipo string, color string) DeteccaoClasse {
	return DeteccaoClasse{id, nome, tipo, color}
}
