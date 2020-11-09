package repositories_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepositories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repositories Suite")
}

var mock sqlmock.Sqlmock

var _ = BeforeEach(func() {
	var db *sql.DB
	var err error

	db, mock, err = sqlmock.New()
	Expect(err).ShouldNot(HaveOccurred())

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	Expect(err).ShouldNot(HaveOccurred())

	repositories.Repository.DB = gormDB
})

var _ = AfterEach(func() {
	err := mock.ExpectationsWereMet()
	Expect(err).ShouldNot(HaveOccurred())
})
