package repositories_test

import (
	"regexp"

	"github.com/emmanuelperotto/pismo-test/api/repositories"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("OperationTypeRepository", func() {
	Context("GetByID", func() {
		It("executes the correct SQL query", func() {
			sql := `SELECT * FROM "operation_types" WHERE "operation_types"."id" = $1 ORDER BY "operation_types"."id" LIMIT 1`
			mock.ExpectQuery(regexp.QuoteMeta(sql))

			repositories.OperationTypeRepository.GetByID(1)
		})
	})
})
