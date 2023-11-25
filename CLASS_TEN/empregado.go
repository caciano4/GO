package main

import "fmt"

type Empregado struct {
	nome    string
	idade   int
	salario int
}

func main() {
	empregado := Empregado{"walter", 30, 15000}

	ptrEmpregado := &empregado

	fmt.Print(empregado.nome, "\n", empregado.salario, "\n")

	empregado.nome = "walter white"
	fmt.Print(ptrEmpregado.nome, "\n", ptrEmpregado.salario, "\n")
}
