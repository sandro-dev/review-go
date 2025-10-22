package easy

import (
	"fmt"
)

/*
	Crie uma função chamada saudacao que recebe um parâmetro nome do tipo string. A função
	deve imprimir no console a mensagem "Olá, [nome]!". Chame essa função a partir da main
	passando seu próprio nome como argumento
*/

func greeting(name string) {
	fmt.Printf("Olá, %s!", name)
}

func RunQuestion01() {

	greeting("Sandro")

}
