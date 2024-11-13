package repositories

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
)

type Repository interface {
	Save(u *models.User) error
	QueryByEmail(email string) (*models.User, error)
	QueryByID(id string) (*models.User, error)
}
