package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		return "$1$FSAop5Lx$LB1BSVAhkcvPBRRuXQJaJ1"
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Necessario a utilização do Parametro #1 Diretorio e #2 Porta")
		os.Exit(1)
	}

	directory := os.Args[1]
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", secret)

	// Move the Print statement outside the handler function
	fmt.Printf("Servidor ativo na porta: %s \n Diretorio: %s\n", port, directory)

	// Use r instead of &r.Request
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(directory)).ServeHTTP(w, &r.Request)
	}))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
