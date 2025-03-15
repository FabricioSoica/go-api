package main

import (
	"fmt"
	"go-api/config"
	"go-api/models"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Erro ao rodar migrations:", err)
	} else {
		fmt.Println("✅ Migração concluída com sucesso!")
	}
}
