package intermediary

import (
	"fmt"
)

/*
	Baseado na Questão 5 (Fácil), crie uma nova função incrementarIdadeReferencia que recebe
	um ponteiro para um inteiro (*int). Dentro da função, use o operador de dereferência para
	incrementar o valor original em 1. Na main, declare minhaIdade, chame a nova função
	passando o endereço de minhaIdade (&minhaIdade) e, depois da chamada, imprima
	minhaIdade. Observe a diferença no resultado em comparação com a questão anterior

*/

func incrementarIdadeReferencia(idade *int) {
	*idade += 1
}

func RunQuestion01() {

	minhaIdade := 28
	fmt.Println("minhaIdade antes:", minhaIdade)

	incrementarIdadeReferencia(&minhaIdade)
	fmt.Println("minhaIdade depois:", minhaIdade)

	fmt.Println()
}
