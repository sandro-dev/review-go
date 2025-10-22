package easy

import (
	"fmt"
)

/*
	Crie uma função chamada exibirDetalhesLivro que recebe um parâmetro do tipo Livro (a
	struct da Questão 2). A função deve imprimir o título e o ano de publicação do livro
	formatados. Chame essa função a partir da main passando a instância do livro que você
	criou.
*/

func exibirDetalhesLivro(l Livro) {
	fmt.Printf("Título: %s, Ano de Publicação: %d", l.Titulo, l.AnoPublicacao)
}

func RunQuestion04() {

	livro := Livro{
		Titulo:        "Concurrency in Go",
		AnoPublicacao: 2017,
	}

	exibirDetalhesLivro(livro)
	fmt.Println()
}
