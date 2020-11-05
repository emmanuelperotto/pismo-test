package services_test

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type accountRepositoryMock struct{}

func (mock accountRepositoryMock) SaveAccountInDB(account *models.Account) (*models.Account, error) {
	return account, nil
}

var _ = Describe("AccountService", func() {

	Context("CreateAccount", func() {
		repositories.AccountRepository = accountRepositoryMock{}

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
				Expect(err.Error()).To(Equal("DocumentNumber must be a number"))
			})
		})
	})
})
