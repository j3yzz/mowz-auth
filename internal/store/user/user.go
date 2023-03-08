package user

import (
	"github.com/j3yzz/mowz/internal/model"
)

type User interface {
	Set(user model.User) error
	FindByEmail(email string) (model.UserWithId, error)
}
