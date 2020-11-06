package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

type operationTypeRepoInterface interface {
	GetByID(id int) (*models.OperationType, error)
}

type operationTypeRepo struct{}

// GetByID tries to find an operationType given an ID. It can return an error if not found
func (repository operationTypeRepo) GetByID(id int) (*models.OperationType, error) {
	var operationType models.OperationType

	err := Repository.DB.First(&operationType, id).Error

	return &operationType, err
}

// OperationTypeRepository is a repository that wraps DB operations on OperationType table
var (
	OperationTypeRepository operationTypeRepoInterface = operationTypeRepo{}
)
