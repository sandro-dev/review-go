package easy

/*
Dado o slice numeros := []int{10, 20, 30, 40, 50}, utilize um laço for...range para iterar sobre
cada elemento e imprimi-lo no console.
*/

import (
	"fmt"
)

func RunQuestion04() {

	fmt.Println("---| Question 04 |---")
	numeros := []int{10, 20, 30, 40, 50}

	for _, value := range numeros {
		fmt.Println(value)
	}

	fmt.Println("O slice de números é", numeros)
}
