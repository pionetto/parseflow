package upload_test

import (
	"parseflow/config"
	upload "parseflow/handlers"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBool(t *testing.T) {
	assert.True(t, upload.ParseBool("true"))
	assert.True(t, upload.ParseBool(" TRUE "))
	assert.False(t, upload.ParseBool("false"))
	assert.False(t, upload.ParseBool("invalid"))
}

func TestWorkerIgnoresInvalidLines(t *testing.T) {
	lines := make(chan string, 1)
	clientes := make(chan config.Cliente, 1)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer close(done)
		lines <- "Linha inválida sem separadores suficientes"
		close(lines)
		upload.Worker(lines, clientes, &wg)
	}()

	wg.Wait()

	select {
	case c := <-clientes:
		t.Errorf("Esperava não receber cliente, mas recebeu: %+v", c)
	case <-done:

	}
}
