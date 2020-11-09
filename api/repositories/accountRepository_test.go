package repositories_test

import (
	"regexp"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("AccountRepository", func() {
	Context("GetAccountByID", func() {
		It("executes the correct SQL query", func() {
			sql := `SELECT * FROM "accounts" WHERE "accounts"."id" = $1 ORDER BY "accounts"."id" LIMIT 1`
			mock.ExpectQuery(regexp.QuoteMeta(sql))

			repositories.AccountRepository.GetAccountByID(1)
		})
	})

	Context("SaveAccountInDB", func() {
		account := models.Account{
			DocumentNumber: "12345678",
		}

		It("executes the correct SQL query", func() {
			sql := `INSERT INTO "accounts"`
			mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(account.DocumentNumber)

			repositories.AccountRepository.SaveAccountInDB(&account)
		})
	})

})
