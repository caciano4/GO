package main

import "fmt"

// this class we learn how to declare a function and atrribute a value to a variable

func main() {
	age := getAge()

	fmt.Println("Idade:", age)
}

func getAge() int {
	var age int

	fmt.Print("Digite sua idade: ")
	fmt.Scanf("%d", &age)

	return age
}
