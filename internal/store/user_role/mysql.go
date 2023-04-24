package user_role

import (
	"github.com/j3yzz/mowz/internal/model"
	"gorm.io/gorm"
)

type MysqlUserRole struct {
	DB *gorm.DB
}

func NewMysqlUserRole(db *gorm.DB) *MysqlUserRole {
	return &MysqlUserRole{DB: db}
}

const Table = "user_roles"
const RoleTable = "roles"

func (ur *MysqlUserRole) Set(user model.UserWithId, role string) error {
	var rm model.Role
	ur.DB.Table(RoleTable).First(&rm, "name = ?", role)

	userRoleModel := model.UserRole{
		RoleId: rm.Id,
		UserId: user.Id,
	}

	c := ur.DB.Omit("id").Create(&userRoleModel)
	return c.Error
}
