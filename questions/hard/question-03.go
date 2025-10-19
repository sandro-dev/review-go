package hard

/*
Tente criar um map onde a chave seja um slice de inteiros (map[[]int]string). O código irá
falhar na compilação. Em um comentário, explique por que Go não permite que slices (ou
maps e functions) sejam usados como chaves de um map.
 Qual é a regra fundamental para um tipo poder ser usado como chave?

*/

import (
	"fmt"
)

func RunQuestion03() {

	// mymap := make(map[[]int]string)
	/*
		Não vai funcionar porque, para ser chave em um map precisa ser um tipo primitivo.
		Um slice é um tipo composto
	*/

	// Explicação: Go exige que as chaves de um map sejam de um tipo "comparável".
	// Isso significa que o Go precisa ser capaz de determinar se duas chaves são
	// iguais ou diferentes usando o operador '=='. Tipos como int, string, bool,
	// ponteiros e structs (se todos os seus campos forem comparáveis) são
	// comparáveis. Tipos como slices, maps e functions não são comparáveis
	// porque não têm uma definição clara de igualdade (ex: dois slices são
	// iguais se tiverem os mesmos elementos, ou se apontarem para o mesmo
	// array subjacente?). Por essa ambiguidade, eles são proibidos como chaves.

	fmt.Println()

}
