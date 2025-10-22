package easy

import (
	"fmt"
)

/*
	Declare uma variável idade do tipo int e atribua a ela sua idade. Em seguida, crie uma variável
	ponteiroIdade que armazene o endereço de memória da variável idade. Imprima o valor de
	idade, o endereço armazenado em ponteiroIdade, e o valor contido no endereço apontado
	por ponteiroIdade (usando o operador de dereferência *).
*/

func RunQuestion03() {

	idade := 28
	ponteiroIdade := &idade

	fmt.Println(idade)
	fmt.Println(ponteiroIdade)
	fmt.Println(*ponteiroIdade)
}
