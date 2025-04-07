package utils_test

import (
	"testing"

	"parseflow/utils"

	"github.com/stretchr/testify/assert"
)

func TestValidateAndFormatCPF(t *testing.T) {
	t.Run("valid CPF", func(t *testing.T) {
		validCPF := "529.982.247-25"
		formattedCPF, ok := utils.ValidateAndFormatCPF(validCPF)

		assert.True(t, ok)
		assert.NotNil(t, formattedCPF)
		assert.Equal(t, "52998224725", *formattedCPF)
	})

	t.Run("invalid CPF - wrong length", func(t *testing.T) {
		invalidCPF := "123.456.789-0"
		formattedCPF, ok := utils.ValidateAndFormatCPF(invalidCPF)

		assert.False(t, ok)
		assert.Nil(t, formattedCPF)
	})
}

func TestValidateAndFormatCNPJ(t *testing.T) {
	t.Run("valid CNPJ", func(t *testing.T) {
		validCNPJ := "11.222.333/0001-81"
		formattedCNPJ, ok := utils.ValidateAndFormatCNPJ(validCNPJ)

		assert.True(t, ok)
		assert.NotNil(t, formattedCNPJ)
		assert.Equal(t, "11222333000181", *formattedCNPJ)
	})

	t.Run("invalid CNPJ - wrong length", func(t *testing.T) {
		invalidCNPJ := "12.345.678/0001-9"
		formattedCNPJ, ok := utils.ValidateAndFormatCNPJ(invalidCNPJ)

		assert.False(t, ok)
		assert.Nil(t, formattedCNPJ)
	})
}

func TestNullifyString(t *testing.T) {
	t.Run("NULL string", func(t *testing.T) {
		val := utils.NullifyString("NULL")
		assert.Nil(t, val)
	})

	t.Run("non-NULL string", func(t *testing.T) {
		val := utils.NullifyString("example")
		assert.NotNil(t, val)
		assert.Equal(t, "EXAMPLE", *val)
	})
}

func TestNullifyFloat(t *testing.T) {
	t.Run("NULL float", func(t *testing.T) {
		val := utils.NullifyFloat("NULL")
		assert.Nil(t, val)
	})

	t.Run("valid float", func(t *testing.T) {
		val := utils.NullifyFloat("123.45")
		assert.NotNil(t, val)
		assert.InDelta(t, 123.45, *val, 0.0001)
	})

	t.Run("invalid float", func(t *testing.T) {
		val := utils.NullifyFloat("invalid")
		assert.Nil(t, val)
	})
}

func TestParseFloatWithComma(t *testing.T) {
	t.Run("valid number with comma", func(t *testing.T) {
		val := utils.ParseFloatWithComma("1.234,56")
		assert.NotNil(t, val)
		assert.InDelta(t, 1234.56, *val, 0.0001)
	})

	t.Run("invalid number", func(t *testing.T) {
		val := utils.ParseFloatWithComma("abc,def")
		assert.Nil(t, val)
	})
}
