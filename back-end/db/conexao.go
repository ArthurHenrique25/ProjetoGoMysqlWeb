package db

import (
	"log"
	"os"
)

func ConexaoMysql() (string, string, string, string, string) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	base := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || porta == "" || base == "" {
		log.Fatal("Variáveis de ambiente do banco não estão definidas corretamente")
	}

	return user, password, host, porta, base
}
