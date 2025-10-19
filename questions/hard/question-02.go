package hard

/*
Crie um slice original := []int{10, 20, 30, 40, 50}. A partir dele, crie um sub-slice subSlice :=
original[1:4] (que conterá [20, 30, 40]). Agora, modifique o primeiro elemento do subSlice
para 999. Imprima tanto o subSlice quanto o original. Explique o resultado. O que aconteceu
com o slice original?

*/

import (
	"fmt"
)

func RunQuestion02() {

	original := []int{10, 20, 30, 40, 50}
	subSlice := original[1:4]

	fmt.Println(original)
	fmt.Println(subSlice)

	subSlice[0] = 999
	fmt.Println()

	/*
		O Slice original foi modificado.
		Um sub-slice não é uma cópia de dados, ele é uma "visão" ou "um ponteiro"
		para uma seção do mesmo array subjacente do slice original
	*/
	fmt.Println(original)
	fmt.Println(subSlice)
}
