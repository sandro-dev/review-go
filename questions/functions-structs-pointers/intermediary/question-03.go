package intermediary

import (
	"fmt"
)

/*
	Crie uma função atualizarAnoLivro que recebe um ponteiro para uma struct Livro (*Livro) e
	um novoAno (int). A função deve atualizar o campo AnoPublicacao da struct original. Na main,
	crie uma instância de Livro, imprima seu ano, chame a função passando o endereço da
	instância e o novo ano, e depois imprima o ano novamente para confirmar a alteração.
*/

// type Livro struct {
// 	Titulo        string
// 	AnoPublicacao int
// }

func atualizarAnoLivro(l *Livro, novoAno int) {
	l.AnoPublicacao = novoAno
}

func RunQuestion03() {

	livro := Livro{
		Titulo:        "O óbvio que ignoramos",
		AnoPublicacao: 2020,
	}

	fmt.Printf("[%s] (%d) - ANTES\n", livro.Titulo, livro.AnoPublicacao)
	atualizarAnoLivro(&livro, 2025)

	fmt.Printf("[%s] (%d) - DEPOIS\n", livro.Titulo, livro.AnoPublicacao)

	fmt.Println()
}
