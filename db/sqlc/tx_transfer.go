package db

import "context"

// TransferTxParams contém os parâmetros de input da transação de transferência
type TransferTxParams struct {
	FromAccountID int64 `json:"fromAccountID"` // ID conta que enviará o dinheiro
	ToAccountID   int64 `json:"toAccountID"`   // ID da conta que receberá o dinheiro
	Amount        int64 `json:"amount"`        // Valor da quantidade enviada
}

// TransferTxResult é o resultado da transação de transferência
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`    // O registro de transferência criado
	FromAccount Account  `json:"fromAccount"` // A conta de origem, depois do seu saldo ser atualizado
	ToAccount   Account  `json:"toAccount"`   // A conta destinatária, após ser atualizada
	FromEntry   Entry    `json:"fromEntry"`   // Registra que o dinheiro está saindo
	ToEntry     Entry    `json:"toEntry"`     // Registra que o dinheiro está entrando
}

// TransferTx performa uma transferência de dinheiro de uma conta para outra.
// Cria um registro de transferência, adiciona registro de movimentações às contas e atualiza o saldo das contas em uma única transação
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	// Resultado vazio
	var result TransferTxResult

	// Cria e executa uma nova transação no banco de dados.
	// A função de callback abaixo, que é o segundo parâmetro, se torna uma closure, pois ela está acessando a variável result da função exterior,
	// no caso, TransferTx e também o argumento arg em TransferTx.
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Cria uma transferência
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// Adicionando registros de entrada e saída
		// Cria a entrada 1
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// Cria a entrada 2
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// Atualiza saldo da conta 1 ou conta 2
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 int64, amount1 int64, accountID2 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{

		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{

		ID:     accountID2,
		Amount: amount2,
	})
	return
}
