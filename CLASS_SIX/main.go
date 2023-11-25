package main

import "fmt"

func main() {
	// salarios := []int{1000, 2000, 3000, 4000, 5000}
	// digits := []int{1, 2, 3, 4, 5}
	name := "salarios"

	// fmt.Printf("isso é um teste %d, mais um p %s anskdlans syab%d", salarios[1], name, digits[2])

	for index, _ := range name {
		fmt.Println(index)
	}
}
