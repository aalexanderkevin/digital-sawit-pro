// This file contains types that are used in the repository layer.
package repository

import (
	"digital-sawit-pro/model"
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type UserGetFilter struct {
	Id          *string
	PhoneNumber *string
}

type User struct {
	Id              *string
	PhoneNumber     *string
	FullName        *string
	Password        *string
	PasswordSalt    *string
	SuccessfulLogin *int
	CreatedAt       *time.Time
}

func (u User) FromModel(data model.User) *User {
	return &User{
		Id:              data.Id,
		PhoneNumber:     data.PhoneNumber,
		FullName:        data.FullName,
		Password:        data.Password,
		PasswordSalt:    data.PasswordSalt,
		SuccessfulLogin: data.SuccessfulLogin,
		CreatedAt:       data.CreatedAt,
	}
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id:              u.Id,
		PhoneNumber:     u.PhoneNumber,
		FullName:        u.FullName,
		Password:        u.Password,
		PasswordSalt:    u.PasswordSalt,
		SuccessfulLogin: u.SuccessfulLogin,
		CreatedAt:       u.CreatedAt,
	}
}

func (u User) GetID() *string {
	return u.Id
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.Id == nil {
		db.Statement.SetColumn("id", ksuid.New().String())
	}

	if u.SuccessfulLogin == nil {
		db.Statement.SetColumn("successful_login", 0)
	}
	return nil
}
