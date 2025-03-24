package utils

import (
	"log"
	"strconv"
	"strings"
)

func ParseFloatWithComma(value string) *float64 {
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println("Erro ao converter número:", value)
		return nil
	}
	return &parsed
}
