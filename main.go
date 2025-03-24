package main

import (
	"log"
	"parseflow/config"
	upload "parseflow/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	router := gin.Default()
	router.POST("/upload", upload.UploadHandler)

	log.Println("Servidor rodando na porta 8080")
	router.Run(":8080")
}
