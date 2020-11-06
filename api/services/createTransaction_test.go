package services_test

import (
	"fmt"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TransactionRepoMock
type transactionRepositoryMock struct{}

func (mock transactionRepositoryMock) SaveTransactionInDB(transaction *models.Transaction) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

// OperationTypeRepoMock
type operationTypeRepoMock struct {
	getByIDResult func(id int) (*models.OperationType, error)
}

func (mock operationTypeRepoMock) GetByID(id int) (*models.OperationType, error) {
	return mock.getByIDResult(id)
}

// Tests
var _ = Describe("Create Transaction", func() {
	Context("Create", func() {
		var operationType models.OperationType
		var transaction models.Transaction

		operationTypeRepoMock := operationTypeRepoMock{}
		transactionRepoMock := transactionRepositoryMock{}
		operationTypeRepoMock.getByIDResult = func(id int) (*models.OperationType, error) {
			return &operationType, nil
		}

		repositories.TransactionRepository = transactionRepoMock
		repositories.OperationTypeRepository = operationTypeRepoMock

		When("transaction is valid", func() {
			It("returns no error", func() {
				transaction := models.Transaction{
					AmountCents:     1000,
					AccountID:       1,
					OperationTypeID: 1,
				}
				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err).To(BeNil())
			})
		})

		When("transaction doesn't have an AccountID", func() {
			It("returns an error", func() {
				transaction := models.Transaction{
					AmountCents:     1000,
					OperationTypeID: 1,
				}
				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err.Error()).To(Equal("AccountID is required"))
			})
		})

		When("transaction doesn't have an OperationTypeID", func() {
			It("returns an error", func() {
				transaction := models.Transaction{
					AmountCents: 1000,
					AccountID:   1,
				}

				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err.Error()).To(Equal("OperationTypeID is required"))
			})
		})

		When("transaction AmountCents should be positive but has negative value", func() {

			BeforeEach(func() {
				operationType = models.OperationType{
					ID:                    1,
					ShouldWithdrawBalance: false,
					Description:           "Pagamento",
				}
				transaction = models.Transaction{
					AmountCents:     -1000,
					AccountID:       1,
					OperationTypeID: operationType.ID,
				}
			})

			It("returns an error", func() {
				_, err := services.CreateTransaction.Create(&transaction)
				expectedMessage := fmt.Sprintf("AmountCents must be positive when registering: %s", operationType.Description)

				Expect(err.Error()).To(Equal(expectedMessage))
			})
		})

		When("transaction AmountCents should be negative but has positive value", func() {
			BeforeEach(func() {
				operationType = models.OperationType{
					ID:                    1,
					ShouldWithdrawBalance: true,
					Description:           "Saque",
				}
				transaction = models.Transaction{
					AmountCents:     1000,
					AccountID:       1,
					OperationTypeID: operationType.ID,
				}
			})

			It("returns an error", func() {
				_, err := services.CreateTransaction.Create(&transaction)
				expectedMessage := fmt.Sprintf("AmountCents must be negative when registering: %s", operationType.Description)

				Expect(err.Error()).To(Equal(expectedMessage))
			})
		})
	})
})
