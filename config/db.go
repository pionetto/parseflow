package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Cliente struct {
	ID                 uint   `gorm:"primaryKey"`
	CPF                string `gorm:"size:14;not null"`
	Private            bool
	Incompleto         bool
	DataUltimaCompra   *string  `gorm:"type:date"`
	TicketMedio        *float64 `gorm:"type:numeric(10,2)"`
	TicketUltimaCompra *float64 `gorm:"type:numeric(10,2)"`
	LojaMaisFrequente  *string  `gorm:"size:255"`
	LojaUltimaCompra   *string  `gorm:"size:255"`
}

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env. Usando variáveis de ambiente do sistema.")
	}
}

func ConnectAndPrepareDatabase() error {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}

	if port == "" || user == "" || password == "" || dbname == "" {
		return fmt.Errorf("erro: variáveis de ambiente do banco não configuradas")
	}

	dsnPostgres := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)
	dbTemp, err := sql.Open("postgres", dsnPostgres)
	if err != nil {
		return fmt.Errorf("erro ao conectar no postgres: %v", err)
	}
	defer dbTemp.Close()

	_, err = dbTemp.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
	if err != nil {
		log.Println("Aviso: Banco já existe ou erro ao criar banco:", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados específico: %v", err)
	}

	DB = db

	if err := DB.AutoMigrate(&Cliente{}); err != nil {
		return fmt.Errorf("erro ao migrar tabela cliente: %v", err)
	}

	log.Println("Banco de dados conectado e migrado com sucesso!")
	return nil
}
