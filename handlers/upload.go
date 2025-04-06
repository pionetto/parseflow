package upload

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"parseflow/config"
	"parseflow/utils"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const BatchSize = 1000
const NumWorkers = 4

var re = regexp.MustCompile(`\s{2,}`)

func UploadHandler(c *gin.Context) {
	start := time.Now()

	if config.DB == nil {
		log.Println("🔧 Inicializando banco de dados...")
		if err := config.ConnectAndPrepareDatabase(); err != nil {
			log.Printf("Erro ao inicializar o banco: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao preparar o banco de dados"})
			return
		}
	}

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

	cleanUploadFolder()

	c.JSON(http.StatusOK, gin.H{"message": "Upload concluído!", "tempo_execucao": duration.String()})
}

func processarArquivo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("❌ Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	linesChan := make(chan string, 1000)
	clientesChan := make(chan config.Cliente, 1000)

	var wg sync.WaitGroup

	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go worker(linesChan, clientesChan, &wg)
	}

	go batchInserter(clientesChan)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesChan <- line
	}
	close(linesChan)

	wg.Wait()
	close(clientesChan)

	if err := scanner.Err(); err != nil {
		log.Println("❌ Erro ao ler arquivo TXT:", err)
	}
}

func worker(linesChan <-chan string, clientesChan chan<- config.Cliente, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range linesChan {
		fields := re.Split(line, -1)

		if len(fields) < 8 {
			log.Printf("❌ Linha ignorada por ter %d colunas, esperado 8: %v", len(fields), fields)
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

		cpf, valido := utils.ValidateAndFormatCPF(fields[0])
		if !valido {
			log.Println("❌ CPF inválido:", fields[0])
			continue
		}

		var lojaMaisFrequente, lojaUltimaCompra *string
		if formattedCNPJ, _ := utils.ValidateAndFormatCNPJ(fields[6]); formattedCNPJ != nil {
			lojaMaisFrequente = formattedCNPJ
		}
		if formattedCNPJ, _ := utils.ValidateAndFormatCNPJ(fields[7]); formattedCNPJ != nil {
			lojaUltimaCompra = formattedCNPJ
		}

		cliente := config.Cliente{
			CPF:                *cpf,
			Private:            parseBool(fields[1]),
			Incompleto:         parseBool(fields[2]),
			DataUltimaCompra:   utils.NullifyString(fields[3]),
			TicketMedio:        ticketMedio,
			TicketUltimaCompra: ticketUltimaCompra,
			LojaMaisFrequente:  lojaMaisFrequente,
			LojaUltimaCompra:   lojaUltimaCompra,
		}

		clientesChan <- cliente
	}
}

func batchInserter(clientesChan <-chan config.Cliente) {
	var batch []config.Cliente

	for cliente := range clientesChan {
		batch = append(batch, cliente)

		if len(batch) >= BatchSize {
			insertBatch(batch)
			batch = nil
		}
	}

	if len(batch) > 0 {
		insertBatch(batch)
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

func cleanUploadFolder() {
	uploadDir := "./uploads"

	files, err := ioutil.ReadDir(uploadDir)
	if err != nil {
		log.Printf("❌ Erro ao ler diretório de uploads: %v", err)
		return
	}

	for _, file := range files {
		filePath := uploadDir + "/" + file.Name()
		err := os.Remove(filePath)
		if err != nil {
			log.Printf("❌ Erro ao remover arquivo %s: %v", filePath, err)
		} else {
			log.Printf("🧹 Arquivo removido: %s", filePath)
		}
	}
}
