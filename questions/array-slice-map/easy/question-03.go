package easy

/*
Crie um map onde as chaves são do tipo string (nome do aluno) e os valores do tipo float64
(nota). Inicialize o map com os seguintes dados: "Ana" com nota 8.5 e "João" com nota 9.0.
Em seguida, acesse e imprima a nota de Ana
*/

import (
	"fmt"
)

func RunQuestion03() {

	fmt.Println("---| Question 03 |---")

	notes := map[string]float64{"Ana": 8.5, "João": 9.0}

	fmt.Println(notes)
	fmt.Println("A nota de Ana é", notes["Ana"])

}
