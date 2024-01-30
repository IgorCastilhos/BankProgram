package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"                                                            // Define o driver do bd, neste caso PostgreSQL.
	dbSource = "postgresql://root:secret@localhost:5432/bankProject?sslmode=disable" // Define a conexão com o bd.
)

var testQueries *Queries // Declara uma var global para armazenar as queries de teste

// TestMain é o entry point para todos os testes de unidade em um pacote golang
// Este método especial é executado antes de qualquer teste unitário e é usado para configurar
// qualquer estado necessário para os testes.
func TestMain(m *testing.M) {

	// Cria uma conexão com o banco de dados
	// Em caso de falha na conexão, o programa encerra com log.Fatal.
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	// Inicializa testQueries com a conexão estabelecida
	// Isso permite o uso de testQueries em testes de unidade subsequentes.
	testQueries = New(conn)

	// Executa todos os testes de unidade definidos no pacote e sai com o status de saída correspondente
	// O método m.Run() executa todos os testes e retorna um código de status.
	os.Exit(m.Run())
}
