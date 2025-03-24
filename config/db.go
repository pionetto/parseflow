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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Erro: Uma ou mais variáveis de ambiente do banco de dados não foram definidas.")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", host, port, user, password)
	dbTemp, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados postgres:", err)
	}
	defer dbTemp.Close()

	_, err = dbTemp.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
	if err != nil {
		log.Println("Aviso: O banco já existe ou não pode ser criado:", err)
	}

	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	DB = db

	log.Println("Executando migração...")
	err = DB.AutoMigrate(&Cliente{})
	if err != nil {
		log.Fatal("Erro ao migrar a estrutura do banco de dados:", err)
	}

	log.Println("Banco de dados conectado e migrado!")
}
