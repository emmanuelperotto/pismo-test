package services_test

import (
	"errors"
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

// AccountRepositoryMock
type accountRepoMock struct{}

func (mock accountRepoMock) Update(account *models.Account) (*models.Account, error) {
	return account, nil
}

func (mock accountRepoMock) GetAccountByID(id int) (*models.Account, error) {
	return &models.Account{
		ID:                        1,
		DocumentNumber:            "123",
		AvailableCreditLimitCents: 1000,
	}, nil
}

func (mock accountRepoMock) SaveAccountInDB(*models.Account) (*models.Account, error) {
	return &models.Account{}, nil
}

// Tests
var _ = Describe("Create Transaction", func() {
	Context("Create", func() {
		var operationType models.OperationType
		var transaction models.Transaction
		operationTypeRepoMock := operationTypeRepoMock{}
		transactionRepoMock := transactionRepositoryMock{}
		accountRepoMock := accountRepoMock{}

		BeforeEach(func() {
			operationTypeRepoMock.getByIDResult = func(id int) (*models.OperationType, error) {
				return &operationType, nil
			}

			repositories.TransactionRepository = transactionRepoMock
			repositories.OperationTypeRepository = operationTypeRepoMock
			repositories.AccountRepository = accountRepoMock
		})

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

		When("OperationTypeID doesn't exist in the database", func() {
			BeforeEach(func() {
				operationTypeRepoMock.getByIDResult = func(id int) (*models.OperationType, error) {
					return &operationType, errors.New("not found")
				}
				repositories.OperationTypeRepository = operationTypeRepoMock
			})

			It("returns an error", func() {
				transaction := models.Transaction{
					AmountCents:     1000,
					AccountID:       1,
					OperationTypeID: 1,
				}

				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err.Error()).To(Equal("OperationTypeID not found"))
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

		When("account has enough limit", func() {
			BeforeEach(func() {
				transaction = models.Transaction{
					AmountCents:     -1000,
					AccountID:       1,
					OperationTypeID: operationType.ID,
				}
			})

			It("returns no error", func() {
				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err).To(BeNil())
			})

			// It("returns correct limit", func() {
			// 	services.CreateTransaction.Create(&transaction)

			// 	Expect(account.AvailableCreditLimitCents).To(Equal(1))
			// })
		})

		When("account doesn't enough limit", func() {
			BeforeEach(func() {
				transaction = models.Transaction{
					AmountCents:     -1500,
					AccountID:       1,
					OperationTypeID: operationType.ID,
				}
			})

			It("returns error", func() {
				_, err := services.CreateTransaction.Create(&transaction)

				Expect(err.Error()).To(Equal("Invalid limit"))
			})
		})
	})
})
