package models

import (
	"github.com/NuttayotSukkum/acn/acn-authentication/internal/constants"
	"time"
)

type User struct {
	ID        string  `gorm:"primaryKey;"`
	FullName  string  `gorm:"size:50;"`
	Email     string  `gorm:"size:100;unique;not null;"`
	Password  string  `gorm:"size:100;unique;not null;"`
	CreatedAt string  `gorm:"size:50;not null;"`
	DeletedAt *string `gorm:"size:50;"`
	UpdatedAt *string `gorm:"size:50;"`
}

type UserModelBuilder struct {
	User *User
}

func NewUserModelBuilder() *UserModelBuilder {
	return &UserModelBuilder{User: &User{}}
}

func (u *UserModelBuilder) SetID(id string) *UserModelBuilder {
	u.User.ID = id
	return u
}

func (u *UserModelBuilder) SetFullName(fullName string) *UserModelBuilder {
	u.User.FullName = fullName
	return u
}

func (u *UserModelBuilder) SetEmail(email string) *UserModelBuilder {
	u.User.Email = email
	return u
}

func (u *UserModelBuilder) SetPassword(password string) *UserModelBuilder {
	u.User.Password = password
	return u
}

func (u *UserModelBuilder) SetCreatedDate(createdDate time.Time) *UserModelBuilder {
	u.User.CreatedAt = createdDate.Format(constants.DATE_TIME_FORMATTER) // Convert time.Time to primitive.DateTime
	return u
}

func (u *UserModelBuilder) SetDeletedDate(deletedDate *time.Time) *UserModelBuilder {
	if u.User.DeletedAt != nil {
		formatDate := deletedDate.Format(constants.DATE_TIME_FORMATTER)
		u.User.DeletedAt = &formatDate
	} else {
		u.User.DeletedAt = nil
	}
	return u
}

func (u *UserModelBuilder) SetUpdatedDate(updatedDate *time.Time) *UserModelBuilder {
	if u.User.UpdatedAt != nil {
		formatDate := updatedDate.Format(constants.DATE_TIME_FORMATTER)
		u.User.UpdatedAt = &formatDate
	} else {
		u.User.UpdatedAt = nil
	}
	return u
}

func (u *UserModelBuilder) Build() *User {
	return u.User
}

func (User) TableName() string {
	return "tbl_users"
}
