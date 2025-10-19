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

type arrayOrSlice interface {
	[]int | [3]int
}

func imprimeSlice[T []int | [3]int](intNumbers T) {
	fmt.Println(intNumbers)
}

func RunQuestion03B() {

	intSlice := []int{1, 2, 3}
	intArray := [3]int{2, 4, 6}
	imprimeSlice(intSlice)
	imprimeSlice(intArray)
	// fmt.Println()
}
