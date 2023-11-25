package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the correct file path based on your project structure
	filePath := "./file-teste.txt"

	// Open the file
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Print("Error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close() // Defer the closing of the file to ensure it's closed properly

	// Print the file type
	fmt.Print(file)

	fmt.Println("Successfully Opened file-teste.txt")

}
