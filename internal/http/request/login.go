package request

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Login struct {
	Email    string
	Password string
}

func (l Login) Validate() error {
	if err := validation.ValidateStruct(&l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required, validation.Length(PasswordMinLength, PasswordMaxLength)),
	); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
