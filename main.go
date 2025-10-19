package main

import (
	"fmt"

	question01 "github.com/sandro-dev/review-go/questions/array-slice-map/easy"
)

/*

 */

type Comentario struct {
	Autor string
	Texto string
}

type Post struct {
	ID          int
	Titulo      string
	Conteudo    string
	Comentarios []Comentario
}

func main() {

	question01.RunQuestion05()

	fmt.Println()

}
