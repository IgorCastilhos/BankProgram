package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/IgorCastilhos/BankProgram/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

// TestMain é o entry point para todos os testes de unidade em um pacote golang
// Este método especial é executado antes de qualquer teste unitário e é usado para configurar
// qualquer estado necessário para os testes.
func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("não pôde carregar configuração:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("não foi possível conectar ao banco de dados:", err)
	}

	testStore = NewStore(connPool)
	// Executa todos os testes de unidade definidos no pacote e sai com o status de saída correspondente
	// O método m.Run() executa todos os testes e retorna um código de status.
	os.Exit(m.Run())
}
