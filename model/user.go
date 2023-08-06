package model

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id              *string    `json:"id"`
	PhoneNumber     *string    `json:"phone_number"`
	FullName        *string    `json:"full_name"`
	Password        *string    `json:"-"`
	PasswordSalt    *string    `json:"-"`
	SuccessfulLogin *int       `json:"successful_login"`
	CreatedAt       *time.Time `json:"created_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.PhoneNumber, validation.Required, validation.Length(10, 13), validation.Match(regexp.MustCompile(`^\+62\d+$`))),
		validation.Field(&u.FullName, validation.Required, validation.Length(3, 60), is.Alpha),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 64), validation.Match(regexp.MustCompile(`^(?=.*[A-Z])(?=.*\d)(?=.*[^A-Za-z0-9]).{8,}$`))),
	)
}
