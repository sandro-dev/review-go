package intermediary

/*
Escreva uma função imprimeSlice que recebe um slice de inteiros ([]int) como parâmetro e o
imprime. Na sua função main, crie um array de 3 inteiros e um slice de 3 inteiros. Tente
chamar a função imprimeSlice com ambos. Explique (em um comentário no seu código) por
que a chamada com o array resulta em um erro de compilação.
*/

import (
	"fmt"
)

func imprimeSlices(intNumbers []int) {
	fmt.Println(intNumbers)
}

func RunQuestion03A() {

	intSlice := []int{1, 2, 3}
	// intArray := [3]int{2, 4, 6}
	imprimeSlices(intSlice)
	// imprimeSlices(intArray) retorna erro porque um array é um slice de tamanho fixo
	// ao definir no tipo []int, espera-se um slice e não um array de tamanho fixo
	// ao fixar o tamanho [3]int, passa a aceitar somente arrays
	// uma solução é definir uma constraint e aceitar ela como parametro

}
