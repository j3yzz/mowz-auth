package user

import (
	"github.com/j3yzz/mowz/internal/model"
	"gorm.io/gorm"
)

var (
	ErrEmailDuplicateCode = 1062
)

type MysqlUser struct {
	DB *gorm.DB
}

func NewMysqlUser(db *gorm.DB) *MysqlUser {
	return &MysqlUser{
		DB: db,
	}
}

const Table = "users"

func (u *MysqlUser) Set(user model.User) error {
	uc := u.DB.Create(user)

	return uc.Error
}

func (u *MysqlUser) FindByEmail(email string) (model.UserWithId, error) {
	var user model.UserWithId

	query := u.DB.Table(Table).Find(&user, "email = ?", email)
	return user, query.Error
}

func (u *MysqlUser) FindById(id string) (model.UserWithId, error) {
	var user model.UserWithId
	query := u.DB.Table(Table).Find(&user, "id = ?", id)
	return user, query.Error
}
