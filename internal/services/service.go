package services

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
)

type Services interface {
	CreateUser(user *models.User) error
	LoginUser(email, password string) (*models.User, error)
}
