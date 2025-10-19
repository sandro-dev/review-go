package main

import (
	"fmt"
)

/*
	Modele a estrutura de um blog simples.
	Crie uma struct Post com os campos ID (int), Titulo (string) e Conteudo (string).
	Em seguida, crie uma struct Comentario com os campos Autor (string) e Texto (string).
	Modifique a struct Post para que ela também contenha um slice de Comentarios.
	Por fim, crie um map chamado blog onde a chave é o ID do post (int) e o valor é a própria struct Post.
	Popule o blog com dois posts, cada um com alguns comentários, e então itere
	sobre o map para imprimir cada post e seus respectivos comentários de forma organizada.
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

	post1 := Post{
		ID:       1,
		Titulo:   "Let's talk about golang",
		Conteudo: "Golang is a static typed language",
		Comentarios: []Comentario{
			{
				Autor: "John",
				Texto: "It's a excelent post!",
			},
			{
				Autor: "Marie",
				Texto: "Awesome point",
			},
		},
	}

	post2 := Post{
		ID:       2,
		Titulo:   "Golang is a performatic language",
		Conteudo: "Golang was develop by Google's Team at 2009 and it's a performatic language",
		Comentarios: []Comentario{
			{
				Autor: "Sandro",
				Texto: "I like golang, it's a simple languege, a small number of reserved words",
			},
			{
				Autor: "Mariana",
				Texto: "I like the go away",
			},
			{
				Autor: "Sarah",
				Texto: "Good! It's a simple language",
			},
		},
	}

	blog := map[int]Post{}

	blog[post1.ID] = post1
	blog[post2.ID] = post2

	for _, post := range blog {
		fmt.Printf("\n -----| Post #%d |----- \n", post.ID)
		fmt.Println("Título:", post.Titulo)
		fmt.Println("Conteúdo:", post.Conteudo)
		fmt.Printf("Comentários do post #%d: \n", post.ID)

		for key, comment := range post.Comentarios {
			fmt.Printf("\n ---Comment %d--- \n", key)
			fmt.Println("\t Autor: ", comment.Autor)
			fmt.Println("\t Comentário: ", comment.Texto)
		}

	}

	// fmt.Println(blog)

}
