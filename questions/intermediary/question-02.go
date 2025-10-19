package intermediary

/*
Crie um map para representar permissões de usuários (chave string para o nome, valor bool
para indicar se é admin). Adicione os usuários "carlos" (admin: true) e "beatriz" (admin: false).
Escreva um código que verifique se um usuário chamado "lucas" existe no map. Se ele não
existir, imprima "Usuário 'lucas' não encontrado.". Se existir, imprima seu status de admin.
*/

import (
	"fmt"
)

func main() {

	permissions := map[string]bool{
		"carlos":  true,
		"beatriz": false,
	}

	isAdmin, ok := permissions["lucas"]

	if !ok {
		println("Usuário 'lucas' não encontrado.")
	} else {
		println("isAdmin:", isAdmin)
	}

	fmt.Println(permissions)
}
