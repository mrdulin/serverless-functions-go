package repositories

import (
	"serverless-functions-go/domain/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
}
