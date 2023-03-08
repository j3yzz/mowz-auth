package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	PasswordMinLength = 6
	PasswordMaxLength = 0
)

type Register struct {
	Name     string
	Email    string
	Password string
}

func (r Register) Validate() error {
	if err := validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(PasswordMinLength, PasswordMaxLength)),
	); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
