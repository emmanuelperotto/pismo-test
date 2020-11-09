package repositories_test

import (
	"regexp"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("TransactionRepository", func() {
	Context("SaveTransactionInDB", func() {
		transaction := models.Transaction{
			AccountID:       1,
			AmountCents:     1000,
			OperationTypeID: 1,
		}

		It("executes the correct SQL query", func() {
			sql := `INSERT INTO "transactions"`
			mock.ExpectQuery(regexp.QuoteMeta(sql)).
				WithArgs(transaction.AmountCents, transaction.AccountID, transaction.OperationTypeID)

			repositories.TransactionRepository.SaveTransactionInDB(&transaction)
		})
	})
})
