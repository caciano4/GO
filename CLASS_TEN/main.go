package main

import "fmt"

func addValue(a *int) {
	*a = 300
}
func main() {

	var a int = 1
	var b *int = &a

	addValue(b)

	fmt.Printf("Esse é o valor de a: %d \n", a)
	fmt.Printf("Esse é o valor do endereço de a: %d \n", &a)
	fmt.Printf("Esse é o valor do endereço de b: %d \n", b)
	fmt.Printf("Esse é o valor de b: %d \n", *b)

}
