package intermediary

import (
	"fmt"
)

/*
	Crie uma nova struct Autor com campos Nome (string) e Nacionalidade (string). Modifique a
	struct Livro para que ela contenha um campo AutorInfo do tipo Autor (composição). Na main,
	crie uma instância de Autor, depois uma instância de Livro associando o autor criado, e
	imprima o título do livro e o nome do seu autor (livro.AutorInfo.Nome).
*/

type Autor struct {
	Nome          string
	Nacionalidade string
}

type Book struct {
	Titulo        string
	AnoPublicacao int
	AutorInfo     Autor
}

func main() {

	autor := Autor{
		Nome:          "Jacob Petry",
		Nacionalidade: "Brasileiro",
	}

	livro := Book{
		Titulo:        "O óbvio que ignoramos",
		AnoPublicacao: 2020,
		AutorInfo:     autor,
	}

	fmt.Printf("[%s] - %s", livro.Titulo, livro.AutorInfo.Nome)
}
