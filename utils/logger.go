package utils

import (
	"log"
	"os"
	"path/filepath"
)

func InitLogDir() {
	logDir := "logs"

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Println("❌ Erro ao criar diretório logs:", err)
		} else {
			log.Println("🗂️ Diretório logs criado com sucesso.")
		}
	}
}

func logInvalid(filename, entryType, value string) {
	InitLogDir()

	filePath := filepath.Join("logs", filename)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("❌ Erro ao abrir arquivo de log:", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Printf("Entrada inválida (%s): %s\n", entryType, value)
}

func LogInvalidCPF(cpf string) {
	logInvalid("cpf_invalido.txt", "CPF", cpf)
}

func LogInvalidCNPJ(cnpj string) {
	logInvalid("cnpj_invalido.txt", "CNPJ", cnpj)
}
