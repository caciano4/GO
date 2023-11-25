package main

import (
	"fmt"
	"testing"
)

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {
	// it's a way to instantiate a struct
	// pessoa := Pessoa{nome: "João", idade: 20, salario: 1000}
	pessoa2 := new(Pessoa)
	pessoa2.idade = 21
	pessoa2.salario = 2000

	pessoa2.addSalaryP(200)
	fmt.Println(pessoa2.salario)
}

func (p *Pessoa) addSalaryP(bonus int) {
	p.salario += bonus
}

// It's still on test
func test_addSalaryP_success(t *testing.T) {
	pessoa := Pessoa{nome: "João", idade: 20, salario: 1000}
	pessoa2 := new(Pessoa)
	pessoa2.addSalaryP(200)
	fmt.Print("Passou no teste")
	if pessoa.salario != 1200 {
		t.Errorf("Expected salary to be 1200, but got %d", pessoa.salario)
	}
}
