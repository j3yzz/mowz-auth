package user_role

import "github.com/j3yzz/mowz/internal/model"

type UserRole interface {
	Set(user model.UserWithId, role string) error
}
