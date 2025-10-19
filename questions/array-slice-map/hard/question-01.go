package hard

/*
Crie um slice usando make([]int, 0, 5), ou seja, com comprimento 0 e capacidade 5.
Em um laço for, adicione 10 números a este slice usando append.
Dentro do laço, a cada adição, imprima o len(), o cap() e o endereço de memória do primeiro elemento (&s[0]).
Observe e explique em um comentário quando e por que o endereço de memória muda.
*/

import (
	"fmt"
)

func RunQuestion01() {

	s := make([]int, 0, 5)

	for i := 0; i < 10; i++ {
		s = append(s, i)
		/*O endereço do primeiro elemento muda porque dobra a capacidade inicial e com isso é criado um slice novo*/
		fmt.Printf("i=%d -> len: %d, cap: %d, addr: %p\n", i, len(s), cap(s), &s[0])

		/*
			O endereço de memória do primeiro elemento permanece o mesmo
			enquanto houver capacidade no array subjacente (até i=4).
			Quando o sexto elemento (i=5) é adicionado, a capacidade (5) é excedida.
			O Go aloca um NOVO array subjacente, geralmente com o dobro da capacidade anterior (10),
			copia os elementos antigos para o novo e faz o slice apontar para ele.
			Essa realocação muda o endereço de memória base.
		*/
	}

	fmt.Println(s)
}
