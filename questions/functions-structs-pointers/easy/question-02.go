package easy

import (
	"fmt"
)

/*
	Defina uma struct chamada Livro com os campos Titulo (string) e AnoPublicacao (int). Na
	função main, crie uma instância dessa struct representando um livro de sua escolha e
	imprima o Titulo do livro no console.

*/

type Livro struct {
	Titulo        string
	AnoPublicacao int
}

func RunQuestion02() {

	livro := Livro{
		Titulo:        "Concurrency in Go",
		AnoPublicacao: 2017,
	}

	fmt.Println(livro)
}
