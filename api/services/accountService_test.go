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

func (mock accountRepositoryMock) SaveAccountInDB(account *models.Account) (*models.Account, error) {
	return account, nil
}

func (mock accountRepositoryMock) GetAccountByID(id int) (*models.Account, error) {
	return mock.getAccountByIDResult(id)
}

var _ = Describe("AccountService", func() {
	var mock accountRepositoryMock = accountRepositoryMock{}

	Context("CreateAccount", func() {
		BeforeEach(func() {
			repositories.AccountRepository = mock
		})

		When("account is valid", func() {
			It("returns no error", func() {
				account := models.Account{
					DocumentNumber: "12345679",
				}

				_, err := services.AccountService.CreateAccount(&account)
				Expect(err).To(BeNil())
			})
		})

		When("account doesn't have a DocumentNumber", func() {
			account := models.Account{
				DocumentNumber: "",
			}

			It("returns an error", func() {
				_, err := services.AccountService.CreateAccount(&account)
				Expect(err.Error()).To(Equal("DocumentNumber can't be empty"))
			})
		})

		When("account doesn't have a numeric DocumentNumber", func() {
			account := models.Account{
				DocumentNumber: "HelloWorld",
			}

			It("returns an error", func() {
				_, err := services.AccountService.CreateAccount(&account)
				Expect(err.Error()).To(Equal("DocumentNumber must contain only numbers"))
			})
		})
	})

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
				_, err := services.AccountService.FindAccountByID(3)
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

				_, err := services.AccountService.FindAccountByID(int(account.ID))
				Expect(err).To(BeNil())
			})
		})
	})
})
