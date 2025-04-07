package utils_test

import (
	"parseflow/config"
	"parseflow/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) func() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao abrir banco em memória: %s", err)
	}

	err = db.AutoMigrate(&config.Cliente{})
	if err != nil {
		t.Fatalf("Erro ao migrar tabela: %s", err)
	}

	originalDB := config.DB
	config.DB = db

	cleanup := func() {
		config.DB = originalDB
	}

	return cleanup
}

func TestInsertBatch_Success(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	clientes := []config.Cliente{
		{CPF: "12345678900"},
	}

	err := utils.InsertBatch(clientes)
	assert.NoError(t, err)

	var count int64
	config.DB.Model(&config.Cliente{}).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestInsertBatch_Empty(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()

	clientes := []config.Cliente{}

	err := utils.InsertBatch(clientes)
	assert.NoError(t, err)

	var count int64
	config.DB.Model(&config.Cliente{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
