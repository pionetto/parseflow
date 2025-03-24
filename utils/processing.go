package utils

import (
	"log"
	"parseflow/config"
)

func InsertBatch(clientes []config.Cliente) error {
	if len(clientes) == 0 {
		log.Println("Nenhum cliente a ser inserido.")
		return nil
	}

	log.Printf("Tentando inserir %d registros no banco de dados...\n", len(clientes))

	err := config.DB.Create(&clientes).Error
	if err != nil {
		log.Println("Erro ao inserir batch:", err)
		return err
	}

	log.Printf("Inseridos %d registros no banco de dados com sucesso.\n", len(clientes))
	return nil
}
