package services

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/configs"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/models"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/pkg/middleware"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/pkg/utils"
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/repositories"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	logger "go.uber.org/zap"
	"time"
)

type UserRepo struct {
	UserRepo repositories.Repository
}

func NewUserService(userRepo repositories.Repository) *UserRepo {
	return &UserRepo{UserRepo: userRepo}
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	if user.FullName == "" || user.Password == "" || user.Email == "" {
		log.Errorf("user name: %v or email: %v or password :%v is empty", user.FullName, user.Email, user.Password)
		return errors.New("user name or email or password is empty")
	}
	haspass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	var req = models.NewUserModelBuilder().SetID(uuid.NewString()).
		SetFullName(user.FullName).
		SetPassword(haspass).
		SetEmail(user.Email).
		SetCreatedDate(time.Now()).
		SetUpdatedDate(nil).
		SetDeletedDate(nil).
		Build()
	if err := repo.UserRepo.Save(req); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) LoginUser(email, password string, cfg configs.Configs) (*string, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password is empty")
	}
	userExist, err := repo.UserRepo.QueryByEmail(email)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if userExist == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := utils.VerifyPassword(userExist.Password, password); err != true {
		return nil, errors.New("invalid password")
	}
	token, err := middleware.Token(userExist.ID, cfg.App.HmacSecret)
	if err != nil {
		return nil, err
	}
	logger.String("token:", *token)
	log.Errorf("token: %v", token)
	return token, nil
}

func (repo *UserRepo) VerifyUser(id string) (*models.User, error) {
	log.Errorf("Not Found user: %s", id)
	user, err := repo.UserRepo.QueryByID(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}
