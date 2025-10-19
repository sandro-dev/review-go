package hard

/*
Por padrão, a iteração sobre um map não tem uma ordem garantida. Dado o map capitais :=
map[string]string{"SP": "São Paulo", "RJ": "Rio de Janeiro", "MG": "Belo Horizonte", "BA":
"Salvador"}, escreva um código que imprima as capitais em ordem alfabética de suas siglas
(chaves).
Dica: Você precisará extrair as chaves para um slice, ordenar o slice e depois iterar sobre ele
*/

import (
	"fmt"
	"sort"
)

func RunQuestion04() {

	capitais := map[string]string{
		"SP": "São Paulo",
		"RJ": "Rio de Janeiro",
		"MG": "Belo Horizonte",
		"BA": "Salvador",
	}

	fmt.Println(capitais)
	keys := make([]string, 0, len(capitais))

	for _, v := range capitais {
		keys = append(keys, v)
		sort.Strings(keys)
	}

	fmt.Println(keys)

	/*
		// 1. Criar um slice para armazenar as chaves.
		chaves := make([]string, 0, len(capitais))
		// 2. Extrair as chaves para o slice.
		for k := range capitais {
			chaves = append(chaves, k)
		}
		// 3. Ordenar o slice de chaves.
		sort.Strings(chaves)
		// 4. Iterar sobre o slice ORDENADO para acessar o map.
		fmt.Println("Capitais em ordem alfabética da sigla:")
		for _, sigla := range chaves {
			fmt.Printf("%s: %s\n", sigla, capitais[sigla])
		}
	*/

}
