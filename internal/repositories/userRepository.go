package repositories

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Save(u *models.User) error {
	return repo.db.Create(u).Error
}

func (repo *UserRepository) QueryByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) QueryByID(id string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Errorf("Error:%v", err)
		return nil, err
	}
	return &user, nil
}
