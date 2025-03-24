package upload

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"parseflow/config"
	"parseflow/utils"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const BatchSize = 1000

var re = regexp.MustCompile("\\s{2,}")

func UploadHandler(c *gin.Context) {
	start := time.Now()

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao fazer upload do arquivo"})
		return
	}

	log.Println("📂 Arquivo recebido:", file.Filename)

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.Mkdir(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar diretório de uploads"})
			return
		}
	}

	tempFile, err := os.CreateTemp(uploadDir, "upload-*.txt")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar arquivo temporário"})
		return
	}
	defer tempFile.Close()

	if err := c.SaveUploadedFile(file, tempFile.Name()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o arquivo"})
		return
	}

	log.Println("📊 Iniciando processamento do arquivo:", tempFile.Name())
	processarArquivo(tempFile.Name())

	duration := time.Since(start)
	log.Printf("✅ Upload e processamento concluídos em %v", duration)

	c.JSON(http.StatusOK, gin.H{"message": "Upload concluído!", "tempo_execucao": duration.String()})
}

func processarArquivo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("❌ Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var clientes []config.Cliente

	for scanner.Scan() {
		line := scanner.Text()
		fields := re.Split(line, -1)

		if len(fields) < 8 {
			log.Printf("❌ Linha ignorada por ter %d colunas, esperado 8: %v", len(fields), fields)
			continue
		}

		log.Println("📌 Processando linha:", fields)

		cpf, valido := utils.ValidateAndFormatCPF(fields[0])
		if !valido {
			log.Println("❌ CPF inválido:", fields[0])
			continue
		}

		ticketMedio := utils.ParseFloatWithComma(fields[4])
		if ticketMedio == nil {
			log.Println("❌ Erro ao converter Ticket Médio:", fields[4])
			continue
		}

		ticketUltimaCompra := utils.ParseFloatWithComma(fields[5])
		if ticketUltimaCompra == nil {
			log.Println("❌ Erro ao converter Ticket Última Compra:", fields[5])
			continue
		}

		log.Println("✅ Linha válida:", fields)

		cliente := config.Cliente{
			CPF:                cpf,
			Private:            parseBool(fields[1]),
			Incompleto:         parseBool(fields[2]),
			DataUltimaCompra:   utils.NullifyString(fields[3]),
			TicketMedio:        ticketMedio,
			TicketUltimaCompra: ticketUltimaCompra,
			LojaMaisFrequente:  utils.NullifyString(fields[6]),
			LojaUltimaCompra:   utils.NullifyString(fields[7]),
		}

		clientes = append(clientes, cliente)

		if len(clientes) >= BatchSize {
			insertBatch(clientes)
			clientes = nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("❌ Erro ao ler arquivo TXT:", err)
		return
	}

	if len(clientes) > 0 {
		insertBatch(clientes)
	}
}

func insertBatch(clientes []config.Cliente) {
	if err := utils.InsertBatch(clientes); err != nil {
		log.Println("❌ Erro ao inserir batch:", err)
	}
}

func parseBool(value string) bool {
	val, err := strconv.ParseBool(strings.TrimSpace(value))
	if err != nil {
		return false
	}
	return val
}
