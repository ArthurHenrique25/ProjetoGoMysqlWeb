package server

import (
	"log"
	"net/http"

	"projeto/back-end/db" // ajuste conforme seu caminho real

	"github.com/joho/godotenv"
)

func StartServer() {
	// Carrega as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env:", err)
	}

	// Rota principal
	http.HandleFunc("/", db.Handler)

	log.Println("Servidor rodando em http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
