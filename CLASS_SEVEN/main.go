package main

// I faced some problems to import the package from another folder, so I had to use the full path
import (
	funcionario "command-line-arguments/Users/waltercaciano/Projetos/CLASS_ROOM/GO/CLASS_SEVEN/funcionario/Pessoa.go"
	"fmt"
)

func main() {
	pessoa := &funcionario.Pessoa{
		nome: "Pessoa",
		idade: 30
	}

	fmt.Println(pessoa.nome)
}
