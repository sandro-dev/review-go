package intermediary

import (
	"fmt"
)

/*
	Adicione um método chamado Descricao à struct Livro (Questão 2, Fácil). O método deve
	retornar uma string formatada como "[Título] ([AnoPublicacao])". Na função main, crie uma
	instância de Livro e chame seu método Descricao, imprimindo o resultado.
*/

type Livro struct {
	Titulo        string
	AnoPublicacao int
}

func (l Livro) Descricao() string {
	description := fmt.Sprintf("[%s] (%d)", l.Titulo, l.AnoPublicacao)
	return description
}

func RunQuestion02() {

	livro := Livro{
		Titulo:        "Concurrency in Go",
		AnoPublicacao: 2017,
	}

	fmt.Println(livro.Descricao())
}
