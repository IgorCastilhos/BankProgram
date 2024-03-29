package db

import (
	"context" //
	"database/sql"
	"github.com/IgorCastilhos/BankProgram/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// createRandomAccount cria uma conta com valores aleatórios e retorna essa conta.
// Utiliza funções do pacote utils para gerar valores aleatórios para os campos da conta.
// Realiza várias verificações para assegurar que a conta foi criada corretamente.
func createRandomAccount(t *testing.T) Account { //
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Test_CreateAccount testa se a função CreateAccount funciona conforme esperado.
// Verifica se a conta criada corresponde aos parâmetros fornecidos e se os campos obrigatórios estão presentes.
func Test_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// Test_GetAccount testa a função GetAccount.
// Cria uma conta, recupera essa conta e verifica se os dados recuperados correspondem aos dados da conta criada.
func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// TestUpdateAccount testa a função UpdateAccount.
// Cria uma conta, atualiza-a e verifica se os dados atualizados correspondem aos esperados.
func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// TestDeleteAccount testa a função DeleteAccount.
// Cria uma conta, deleta-a e verifica se ela realmente foi removida.
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

// TestListAccounts testa a função ListAccounts.
// Cria várias contas, lista uma quantidade específica delas e verifica se a lista contém a quantidade correta de contas.
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
