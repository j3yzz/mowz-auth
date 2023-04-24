package model

type UserRole struct {
	Id     int `json:"id" gorm:"primary_key"`
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}
