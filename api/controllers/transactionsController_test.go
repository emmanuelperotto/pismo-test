package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/emmanuelperotto/pismo-test/api/controllers"
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type createTransactionMock struct {
	createResult func(transaction *models.Transaction) (*models.Transaction, error)
}

func (mock createTransactionMock) Create(transaction *models.Transaction) (*models.Transaction, error) {
	return mock.createResult(transaction)
}

var _ = Describe("Transactions Controller", func() {
	Context("CreateTransaction", func() {
		mock := createTransactionMock{}
		requestBody := bytes.NewBuffer([]byte{})

		When("successfully registering a transaction", func() {
			BeforeEach(func() {
				mock.createResult = func(transaction *models.Transaction) (*models.Transaction, error) {
					return transaction, nil
				}

				services.CreateTransaction = mock
			})

			It("returns a status code CREATED", func() {
				request, _ := http.NewRequest("POST", "/transactions", requestBody)
				response := httptest.NewRecorder()

				controllers.CreateTransaction(response, request)

				Expect(response.Code).To(Equal(http.StatusCreated))
			})
		})

		When("failing to register an transaction", func() {
			BeforeEach(func() {
				mock.createResult = func(transaction *models.Transaction) (*models.Transaction, error) {
					return transaction, errors.New("Error")
				}

				services.CreateTransaction = mock
			})

			It("returns a status code UNPROCESSABLE_ENTITY", func() {
				request, _ := http.NewRequest("POST", "/transactions", requestBody)
				response := httptest.NewRecorder()

				controllers.CreateTransaction(response, request)

				Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})
	})
})
