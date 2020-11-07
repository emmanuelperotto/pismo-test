package controllers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/emmanuelperotto/pismo-test/api/controllers"
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/services"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// SearchAccount mock
type searchAccountServiceMock struct {
	findByIDResult func(id int) (*models.Account, error)
}

func (mock searchAccountServiceMock) FindByID(id int) (*models.Account, error) {
	return mock.findByIDResult(id)
}

// CreateAccount mock
type createAccountServiceMock struct {
	createResult func(account *models.Account) (*models.Account, error)
}

func (mock createAccountServiceMock) Create(account *models.Account) (*models.Account, error) {
	return mock.createResult(account)
}

var _ = Describe("Accounts Controller", func() {
	Context("GetAccount", func() {
		mock := searchAccountServiceMock{}
		queryParams := map[string]string{
			"id": "1",
		}

		When("account exists", func() {
			BeforeEach(func() {
				mock.findByIDResult = func(id int) (*models.Account, error) {
					return &models.Account{ID: id}, nil
				}

				services.SearchAccount = mock
			})

			It("returns status OK", func() {
				request, _ := http.NewRequest("GET", "/accounts/1", nil)
				response := httptest.NewRecorder()
				request = mux.SetURLVars(request, queryParams)

				controllers.GetAccount(response, request)

				Expect(response.Code).To(Equal(http.StatusOK))
			})
		})

		When("account doesn't exist", func() {
			BeforeEach(func() {
				mock.findByIDResult = func(id int) (*models.Account, error) {
					return &models.Account{}, errors.New("Error")
				}

				services.SearchAccount = mock
			})

			It("returns status NOT FOUND", func() {
				request, _ := http.NewRequest("GET", "/accounts/1", nil)
				response := httptest.NewRecorder()
				request = mux.SetURLVars(request, queryParams)

				controllers.GetAccount(response, request)

				Expect(response.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	Context("CreateAccount", func() {
		mock := createAccountServiceMock{}
		requestBody := bytes.NewBuffer([]byte{})

		When("successfully registering an account", func() {

			BeforeEach(func() {
				mock.createResult = func(account *models.Account) (*models.Account, error) {
					return account, nil
				}

				services.CreateAccount = mock
			})

			It("returns a status code CREATED", func() {

				request, _ := http.NewRequest("POST", "/accounts", requestBody)
				response := httptest.NewRecorder()

				controllers.CreateAccount(response, request)

				Expect(response.Code).To(Equal(http.StatusCreated))
			})
		})

		When("failing to register an account", func() {
			BeforeEach(func() {
				mock.createResult = func(account *models.Account) (*models.Account, error) {
					return account, errors.New("Error")
				}

				services.CreateAccount = mock
			})

			It("returns a status code UNPROCESSABLE_ENTITY", func() {
				request, _ := http.NewRequest("POST", "/accounts", requestBody)
				response := httptest.NewRecorder()

				controllers.CreateAccount(response, request)

				Expect(response.Code).To(Equal(http.StatusUnprocessableEntity))
			})
		})
	})
})
