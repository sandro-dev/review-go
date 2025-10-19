package hard

/*
Modele a estrutura de um blog simples. Crie uma struct Post com os campos ID (int), Titulo
(string) e Conteudo (string). Em seguida, crie uma struct Comentario com os campos Autor
(string) e Texto (string). Modifique a struct Post para que ela também contenha um slice de
Comentarios.
Por fim, crie um map chamado blog onde a chave é o ID do post (int) e o valor é a própria
struct Post. Popule o blog com dois posts, cada um com alguns comentários, e então itere
sobre o map para imprimir cada post e seus respectivos comentários de forma organizada.
*/

import (
	"fmt"
)

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

func RunQuestion05() {

	blog := map[int]Post{}

	blog[1] = Post{
		ID:          1,
		Titulo:      "Title 1",
		Conteudo:    "Conteúdo 1",
		Comentarios: append([]Comentario{}, Comentario{Autor: "Autor 1", Texto: "Texto 1"}),
	}

	blog[2] = Post{
		ID:       2,
		Titulo:   "Title 2",
		Conteudo: "Conteúdo 2",
		Comentarios: append(
			[]Comentario{},
			Comentario{Autor: "Autor 2", Texto: "Texto 2"},
			Comentario{Autor: "Autor 3", Texto: "Texto 3"}),
	}

	for k, v := range blog {
		fmt.Println(k, v)
	}

	/*
		post1 := Post{
			ID:       1,
			Titulo:   "Introdução a Go",
			Conteudo: "Go é uma linguagem de programação incrível...",
			Comentarios: []Comentario{
				{Autor: "Ana", Texto: "Ótimo post!"},
				{Autor: "João", Texto: "Aprendi muito."},
			},
		}
		post2 := Post{
			ID:       2,
			Titulo:   "Concorrência em Go",
			Conteudo: "Goroutines e Channels são fantásticos...",
			Comentarios: []Comentario{
				{Autor: "Maria", Texto: "Preciso estudar mais sobre isso."},
			},
		}
		// 2. Criar e popular o map do blog
		blog := make(map[int]Post)
		blog[post1.ID] = post1
		blog[post2.ID] = post2
		// 3. Iterar e imprimir
		for _, post := range blog {
			fmt.Printf("--- Post %d: %s ---\n", post.ID, post.Titulo)
			fmt.Println(post.Conteudo)
			fmt.Println("Comentários:")
			if len(post.Comentarios) == 0 {
				fmt.Println(" (Nenhum comentário)")
			} else {
				for _, comentario := range post.Comentarios {
					fmt.Printf(" - %s: %s\n", comentario.Autor, comentario.Texto)
				}
			}
			fmt.Println()
		}
	*/

}
