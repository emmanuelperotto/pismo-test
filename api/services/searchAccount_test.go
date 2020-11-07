package services_test

import (
	"errors"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type accountRepositoryMock struct {
	getAccountByIDResult func(id int) (*models.Account, error)
}

func (mock accountRepositoryMock) GetAccountByID(id int) (*models.Account, error) {
	return mock.getAccountByIDResult(id)
}

var _ = Describe("AccountService", func() {
	var mock accountRepositoryMock = accountRepositoryMock{}

	Context("FindAccountByID", func() {
		When("account doesn't exist", func() {
			BeforeEach(func() {
				mock = accountRepositoryMock{}
				mock.getAccountByIDResult = func(id int) (*models.Account, error) {
					account := models.Account{}
					return &account, errors.New("Not found")
				}

				repositories.AccountRepository = mock
			})

			It("returns no error", func() {
				_, err := services.SearchAccount.FindByID(3)
				Expect(err.Error()).To(Equal("Not found"))
			})

		})

		When("account exists", func() {
			BeforeEach(func() {
				mock.getAccountByIDResult = func(id int) (*models.Account, error) {
					account := models.Account{}
					return &account, nil
				}

				repositories.AccountRepository = mock
			})

			It("returns no error", func() {
				account := models.Account{
					ID: 1,
				}

				_, err := services.SearchAccount.FindByID(int(account.ID))
				Expect(err).To(BeNil())
			})
		})
	})
})
