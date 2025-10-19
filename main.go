package main

/*
Dado o slice de palavras palavras := []string{"go", "java", "go", "python", "go", "java"}, crie um
map que conte a frequência de cada palavra. O map final deve ter as palavras como chaves e
a contagem como valor. Imprima o map de frequência no final. (Resultado esperado:
map[go:3 java:2 python:1])
*/

import (
	"fmt"
)

func main() {

	palavras := []string{"go", "java", "go", "python", "go", "java"}

	frequencia := map[string]int{}

	for _, value := range palavras {
		_, exists := frequencia[value]

		if !exists {
			frequencia[value] = 1
		} else {
			frequencia[value] += 1
		}
	}

	fmt.Println(frequencia)
}
