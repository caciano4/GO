package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	exception()

	agora := time.Now()
	route := os.Args[1]
	get, err := http.Get(route)
	if err != nil {
		println("Error: ", err.Error())
		os.Exit(1)
	}

	decorrido := time.Since(agora).Seconds()

	fmt.Printf("Tempo decorrido: [%f], Status da Request: [%s]", decorrido, get.Status)

}

func exception() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered in f", r)
		}
	}()

	if err := os.MkdirAll("tmp/subdir", 0755); err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		panic("usage: " + os.Args[0] + " <route>")
	}

	panic("Error")
}
