package easy

import (
	"fmt"
)

/*
 */

func incrementarIdadeCopia(idadeAtual int) int {
	idadeAtual += 1
	fmt.Println("idadeAtual: ", idadeAtual)
	return idadeAtual
}

func RunQuestion05() {

	minhaIdade := 28
	fmt.Printf("minhaIdade antes %d \n", minhaIdade)

	/*
		É feita uma cópia de minhaIdade e passada a função incrementarIdadeCopia
		Golang por default é pass-by-value
	*/
	incrementarIdadeCopia(minhaIdade)
	fmt.Printf("minhaIdade depois %d \n", minhaIdade)

	// Explicação: O valor de minhaIdade não muda fora da função porque
	// Go, por padrão, passa argumentos para funções por VALOR (cópia).
	// A função 'incrementarIdadeCopia' recebe uma cópia de minhaIdade,
	// e qualquer modificação é feita apenas nessa cópia local.
}
