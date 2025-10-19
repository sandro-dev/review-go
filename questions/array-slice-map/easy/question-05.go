package easy

/*
Crie um map chamado estoque para armazenar a quantidade de produtos (chave string, valor
int). Adicione os seguintes itens: "Caneta" com 100 unidades e "Lápis" com 150. Em seguida,
adicione um novo produto, "Borracha", com 200 unidades
*/

import (
	"fmt"
)

func RunQuestion05() {

	estoque := map[string]int{
		"Caneta": 100,
		"Lápis":  150,
	}

	estoque["Borracha"] = 200

	fmt.Println("O estoque é", estoque)
}
