package main

import (
	"fmt"
)

func main() {
	// name, salary := os.Getenv("NAME"), os.Getenv("SALARY")
	name, salary, bonus := "Caciano", 30000, 10000
	setName(name)
	newSalary, bonusValue := relatorio(salary, bonus)

	fmt.Println("Novo Salario:", newSalary, "Valor do Bonus:", bonusValue)
}

func setName(name string) {
	fmt.Println("Ola", name)
}

func addSalary(currentSalary int, bonus int) {
	fmt.Print("Seu novo salario é: ")
	fmt.Println(currentSalary + bonus)
}

func relatorio(currentSalary int, bonus int) (int, int) {
	return currentSalary + bonus, bonus
}
