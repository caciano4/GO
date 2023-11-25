package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/waltercaciano/Projetos/CLASS_ROOM/GO/CLASS_ELEVENTH"))
	fmt.Printf("Server is running on port 9000")
	log.Fatal(http.ListenAndServe(":9000", fs))
}
