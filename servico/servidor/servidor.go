package main

import (
	"github.com/gorilla/mux"
	"mega/servico/agente"
	"mega/servico/banco"
	"mega/servico/custom"
	"mega/servico/login"
	"net/http"
	"os"
)

func main() {
	porta := os.Getenv("SERVICO_PORTA")

	if porta == "" {
		panic("Favor informar uma porta na variável de ambiente SERVICO_PORTA")
	}

	banco.Conectar()

	r := mux.NewRouter()

	agente.Inicializa(r)
	custom.Inicializa(r)
	login.Inicializa(r)

	http.Handle("/api/", r)
	http.Handle("/login", r)

	fs := http.Dir("public/")
	http.Handle("/", http.FileServer(fs))

	http.ListenAndServe(":"+porta, nil)

}
