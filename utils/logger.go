package utils

import (
	"log"
	"os"
)

const logFilePath = "logs/invalid_entries.log"

func InitLogDir() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}
}

func LogInvalidEntry(entryType, value string) {
	InitLogDir()

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erro ao abrir arquivo de log:", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Printf("Entrada inválida (%s): %s\n", entryType, value)
}
