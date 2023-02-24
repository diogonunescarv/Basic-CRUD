package main

import (
	"crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() // Contém todas as rotas para a conexão com o Banco de Dados

	// Quando a Rota de "/usuarios" receber um post ela irá iniciar alguma função
	router.HandleFunc("/usuarios", server.CriarUsurario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", server.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":5000", router))
}
