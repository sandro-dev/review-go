package intermediary

/*
Dado o slice numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, crie um novo slice chamado pares que
contenha apenas os números pares do slice original. Imprima o slice pares no final.
*/

import (
	"fmt"
)

func RunQuestion01() {

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares := []int{}

	for _, value := range numeros {
		if value%2 == 0 {
			pares = append(pares, value)
		}
	}

	fmt.Println(pares)
}
