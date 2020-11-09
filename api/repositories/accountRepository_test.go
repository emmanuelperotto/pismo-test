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
			DocumentNumber:            "12345678",
			AvailableCreditLimitCents: 100000,
		}

		It("executes the correct SQL query", func() {
			sql := `INSERT INTO "accounts"`
			mock.ExpectQuery(regexp.QuoteMeta(sql)).WithArgs(account.DocumentNumber, account.AvailableCreditLimitCents)

			repositories.AccountRepository.SaveAccountInDB(&account)
		})
	})

	// Context("Update", func() {
	// 	account := models.Account{
	// 		ID:                        1,
	// 		DocumentNumber:            "12345678",
	// 		AvailableCreditLimitCents: 500000,
	// 	}

	// 	FIt("executes the correct SQL query", func() {
	// 		sql := `UPDATE "accounts" SET "created_at"='0000-00-00 00:00:00',"updated_at"='2020-11-09 18:29:47.98',"document_number"='12345678',"available_credit_limit_cents"=500000 WHERE "id" = 1`
	// 		mock.ExpectQuery(regexp.QuoteMeta(sql))

	// 		repositories.AccountRepository.Update(&account)
	// 	})
	// })
})
