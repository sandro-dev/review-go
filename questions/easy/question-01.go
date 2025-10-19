package easy

/*
Questão 1 (Array): Dias da Semana
Declare um array de strings com tamanho 5 chamado diasUteis. Inicialize-o com os cinco dias
úteis da semana. Em seguida, imprima o valor que está na terceira posição do array
*/

import (
	"fmt"
)

func RunQuestion01() {

	fmt.Println("---| Question 01 |---")

	diasUteis := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println("O valor que está na terceira posição do array diasUteis é", diasUteis[2])
}
