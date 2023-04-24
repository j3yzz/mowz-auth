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

func (u *MysqlUser) Set(user *model.UserWithId) error {
	uc := u.DB.Table(Table).Omit("id").Create(&user)
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

func (u *MysqlUser) FindUserWithRole(id string) (model.UserWithRole, error) {
	var user model.UserWithRole
	query := u.DB.Raw("select users.id, users.email, users.password, users.name, r.name as role from users "+
		"inner join user_roles ur on users.id = ur.user_id "+
		"inner join roles r on r.id = ur.role_id "+
		"where users.id = ?", id).First(&user)

	return user, query.Error
}
