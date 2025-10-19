package easy

/*
Questão 2 (Slice): Lista de Compras
Crie um slice de strings vazio chamado listaDeCompras. Utilize a função append para
adicionar os seguintes itens: "Leite", "Pão", "Café". Por fim, imprima o slice completo e o seu
tamanho utilizando a função len()
*/

import (
	"fmt"
)

func RunQuestion02() {

	fmt.Println("---| Question 02 |---")

	// var listaDeCompras []string
	listaDeCompras := []string{}

	listaDeCompras = append(listaDeCompras, "Leite")
	listaDeCompras = append(listaDeCompras, "Pão")
	listaDeCompras = append(listaDeCompras, "Café")

	fmt.Println(listaDeCompras)
	fmt.Println("O tamanho do slice listaDeCompras é:", len(listaDeCompras))
}
