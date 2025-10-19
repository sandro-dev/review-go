package intermediary

/*
Dado o slice letras := []string{"a", "b", "c", "d", "e", "f", "g"}, use a operação de "slicing"
([low:high]) para fazer o seguinte:
1. Crie um novo slice primeirasDuas contendo as duas primeiras letras.
2. Crie um novo slice meio contendo as letras de "c" a "e".
3. Crie um novo slice ultimasDuas contendo as duas últimas letras.
Imprima os três novos slices
*/

import (
	"fmt"
)

func RunQuestion04() {

	letras := []string{"a", "b", "c", "d", "e", "f", "g"}

	primeirasDuas := letras[:2]
	meio := letras[2:5]
	ultimasDuas := letras[5:]

	fmt.Println(primeirasDuas)
	fmt.Println(meio)
	fmt.Println(ultimasDuas)
}
