package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func ValidateAndFormatCPF(cpf string) (string, bool) {
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		log.Println("CPF inválido:", cpf)
		return cpf, false
	}
	return cpf, true
}

func ValidateAndFormatCNPJ(cnpj string) (string, bool) {
	re := regexp.MustCompile(`\D`)
	cleanCNPJ := re.ReplaceAllString(cnpj, "")

	if len(cleanCNPJ) != 14 {
		return cnpj, false
	}
	return cleanCNPJ, true
}

func NullifyString(value string) *string {
	if strings.ToUpper(value) == "NULL" {
		return nil
	}
	upperValue := strings.ToUpper(value)
	return &upperValue
}

func NullifyFloat(value string) *float64 {
	if strings.ToUpper(value) == "NULL" {
		return nil
	}

	var result float64
	_, err := fmt.Sscanf(value, "%f", &result)
	if err != nil {
		return nil
	}
	return &result
}

func LogInvalidEntries(cpf, lojaMaisFrequente, lojaUltimaCompra string) {
	file, err := os.OpenFile("logs/invalid_entries.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	if cpf != "" {
		logger.Println("CPF inválido ignorado:", cpf)
	}
	if lojaMaisFrequente != "" {
		logger.Println("CNPJ da loja mais frequente inválido ignorado:", lojaMaisFrequente)
	}
	if lojaUltimaCompra != "" {
		logger.Println("CNPJ da loja última compra inválido ignorado:", lojaUltimaCompra)
	}
}
