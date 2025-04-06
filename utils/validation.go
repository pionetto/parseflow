package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ValidateAndFormatCPF(cpf string) (*string, bool) {
	re := regexp.MustCompile(`\D`)
	cleanCPF := re.ReplaceAllString(cpf, "")

	if len(cleanCPF) != 11 {
		LogInvalidCPF(cpf)
		return nil, false
	}

	return &cleanCPF, true
}

func ValidateAndFormatCNPJ(cnpj string) (*string, bool) {
	re := regexp.MustCompile(`\D`)
	cleanCNPJ := re.ReplaceAllString(cnpj, "")

	if len(cleanCNPJ) != 14 {
		LogInvalidCNPJ(cnpj)
		return nil, false
	}

	return &cleanCNPJ, true
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

func ParseFloatWithComma(value string) *float64 {
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}
	return &floatValue
}
